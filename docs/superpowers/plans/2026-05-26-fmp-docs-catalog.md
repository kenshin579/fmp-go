# FMP API docs 카탈로그 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** FMP stable API 전체(≈274개) 엔드포인트 문서를 `fmp-go/docs/api/<category>/<slug>.md`로 충실히 변환한 카탈로그를 생성한다.

**Architecture:** `tools/gendocs/`의 self-contained Node + Playwright 크롤러가 ① `sitemap-3.xml`에서 doc URL을 열거 → ② 각 페이지를 렌더링(JS 실행)해 `__NEXT_DATA__.props.pageProps.doc`(메타데이터)와 렌더된 응답 예시 JSON을 추출 → ③ md로 변환해 카테고리(`doc.code`)별 디렉토리에 기록 → ④ `docs/api/README.md` 인덱스 생성. 순수 변환 로직(`lib.mjs`)은 `node:test`로 단위테스트, 크롤링은 소규모(LIMIT) 검증 후 전체 실행.

**Tech Stack:** Node.js (ESM), `@playwright/test`(chromium), `node:test`/`node:assert`. Go 코드 없음(생성기는 fmp-go 레포의 docs 생성 도구).

**Spec:** `docs/superpowers/specs/2026-05-26-fmp-docs-catalog-design.md`

**Repo:** `github.com/kenshin579/fmp-go` — `/Users/frankoh/src/workspace_moneyflow/fmp-go`, branch `feature/sdk-foundation`.

**확정된 사전 조사:**
- 열거: `https://site.financialmodelingprep.com/sitemap-3.xml`의 `<loc>` 중 `.../developer/docs/stable/<slug>` ≈ 274개. (요청에 브라우저 User-Agent 필요 — 403 회피.)
- 메타데이터: `__NEXT_DATA__.props.pageProps.doc` = `{title, description, urls:[endpointURL], params:{query:{header,rows}}, code(카테고리), subcode, content(HTML), pageURL(slug)}`.
- 응답 예시: 렌더된 DOM의 `<code>` 요소 중 텍스트가 `[` 또는 `{`로 시작하는 것(라인번호 거터는 별도 `<code>`라 제외됨). 응답은 배열 형태(`[ { "symbol":"AAPL", ... } ]`) 확인됨. API 키 불필요(문서 샘플).

---

## File Structure

```
fmp-go/tools/gendocs/
├── package.json        # ESM, @playwright/test 의존, "gen"/"test" 스크립트
├── .gitignore          # node_modules
├── lib.mjs             # 순수 변환 로직 (paramsTable/htmlToText/renderMarkdown/categoryDir)
├── lib.test.mjs        # node:test 단위테스트
├── gendocs.mjs         # 크롤러(열거 + Playwright 렌더 + 추출 + 파일 기록 + 인덱스)
└── failures.log        # (런타임 산출) 실패 URL
```
출력: `fmp-go/docs/api/<category>/<slug>.md` + `fmp-go/docs/api/README.md`.

---

## Task 1: gendocs 도구 스캐폴딩

**Files:**
- Create: `tools/gendocs/package.json`
- Create: `tools/gendocs/.gitignore`

- [ ] **Step 1: package.json 작성**

Create `tools/gendocs/package.json`:
```json
{
  "name": "fmp-gendocs",
  "private": true,
  "type": "module",
  "version": "0.0.0",
  "description": "FMP API docs → markdown catalog generator",
  "scripts": {
    "gen": "node gendocs.mjs",
    "test": "node --test"
  },
  "devDependencies": {
    "@playwright/test": "^1.60.0"
  }
}
```

- [ ] **Step 2: .gitignore 작성**

Create `tools/gendocs/.gitignore`:
```
node_modules/
failures.log
```

- [ ] **Step 3: 의존성 + 브라우저 설치**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go/tools/gendocs
npm install
npx playwright install chromium
```
Expected: `node_modules/` 생성, chromium 다운로드 완료(이미 있으면 즉시 종료).

- [ ] **Step 4: Commit**

```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
git add tools/gendocs/package.json tools/gendocs/.gitignore tools/gendocs/package-lock.json
git commit -m "chore(gendocs): Node+Playwright 도구 스캐폴딩"
```

---

## Task 2: 순수 변환 로직 (lib.mjs) + 단위테스트

**Files:**
- Create: `tools/gendocs/lib.mjs`
- Create: `tools/gendocs/lib.test.mjs`

- [ ] **Step 1: 실패하는 테스트 작성**

Create `tools/gendocs/lib.test.mjs`:
````js
import { test } from 'node:test';
import assert from 'node:assert';
import { paramsTable, htmlToText, categoryDir, renderMarkdown } from './lib.mjs';

test('paramsTable renders a markdown table', () => {
  const params = { query: { header: ['Query Parameter', 'Type', 'Example'], rows: [['symbol*', 'string', 'AAPL']] } };
  const md = paramsTable(params);
  assert.match(md, /\| Query Parameter \| Type \| Example \|/);
  assert.match(md, /\| symbol\* \| string \| AAPL \|/);
  assert.match(md, /\| --- \| --- \| --- \|/);
});

test('paramsTable returns empty string when no params', () => {
  assert.strictEqual(paramsTable(undefined), '');
  assert.strictEqual(paramsTable({ query: { header: [], rows: [] } }), '');
});

test('htmlToText strips tags and decodes entities', () => {
  const txt = htmlToText('<p>Key &amp; metrics</p><ul><li>One</li><li>Two</li></ul>');
  assert.match(txt, /Key & metrics/);
  assert.match(txt, /- One/);
  assert.match(txt, /- Two/);
  assert.doesNotMatch(txt, /<[^>]+>/);
});

test('categoryDir uses code, falls back to _uncategorized', () => {
  assert.strictEqual(categoryDir({ code: 'company' }), 'company');
  assert.strictEqual(categoryDir({ code: '' }), '_uncategorized');
  assert.strictEqual(categoryDir({}), '_uncategorized');
});

test('renderMarkdown composes all sections', () => {
  const doc = {
    title: 'Company Profile Data',
    description: 'Access company profile.',
    urls: ['https://financialmodelingprep.com/stable/profile?symbol=AAPL'],
    params: { query: { header: ['Query Parameter', 'Type', 'Example'], rows: [['symbol*', 'string', 'AAPL']] } },
    content: '<p>Details</p>',
    code: 'company',
    pageURL: 'profile-symbol',
  };
  const md = renderMarkdown(doc, '[\n  { "symbol": "AAPL" }\n]', 'https://site.financialmodelingprep.com/developer/docs/stable/profile-symbol');
  assert.match(md, /^# Company Profile Data/m);
  assert.match(md, /## Endpoint/);
  assert.match(md, /`GET https:\/\/financialmodelingprep\.com\/stable\/profile\?symbol=AAPL`/);
  assert.match(md, /## Parameters/);
  assert.match(md, /## Description/);
  assert.match(md, /## Response \(example\)/);
  assert.match(md, /```json/);
  assert.match(md, /"symbol": "AAPL"/);
  assert.match(md, /카테고리: company/);
});

test('renderMarkdown notes missing response example', () => {
  const md = renderMarkdown({ title: 'X', pageURL: 'x', code: 'c' }, null, 'http://u');
  assert.match(md, /응답 예시를 찾지 못함/);
});
````

- [ ] **Step 2: 테스트 실행 — 실패 확인**

Run: `cd /Users/frankoh/src/workspace_moneyflow/fmp-go/tools/gendocs && node --test`
Expected: FAIL — `lib.mjs` 없음(모듈 import 에러).

- [ ] **Step 3: lib.mjs 구현**

Create `tools/gendocs/lib.mjs`:
````js
// 순수 변환 로직 — Playwright/네트워크 의존 없음(단위테스트 대상).

// paramsTable 은 doc.params({query:{header,rows}}) 를 markdown 표로. 없으면 빈 문자열.
export function paramsTable(params) {
  const q = params?.query;
  if (!q || !Array.isArray(q.header) || q.header.length === 0 || !Array.isArray(q.rows) || q.rows.length === 0) {
    return '';
  }
  const head = `| ${q.header.join(' | ')} |`;
  const sep = `| ${q.header.map(() => '---').join(' | ')} |`;
  const body = q.rows.map((r) => `| ${r.join(' | ')} |`).join('\n');
  return [head, sep, body].join('\n');
}

// htmlToText 는 doc.content HTML 을 평문 텍스트로(태그 제거 + 기본 엔티티 디코딩).
export function htmlToText(html) {
  if (!html) return '';
  return String(html)
    .replace(/<li[^>]*>/gi, '- ')
    .replace(/<\/(p|li|ul|ol|h[1-6]|div)>/gi, '\n')
    .replace(/<br\s*\/?>/gi, '\n')
    .replace(/<[^>]+>/g, '')
    .replace(/&rsquo;/g, "'").replace(/&lsquo;/g, "'")
    .replace(/&ldquo;/g, '"').replace(/&rdquo;/g, '"')
    .replace(/&amp;/g, '&').replace(/&lt;/g, '<').replace(/&gt;/g, '>')
    .replace(/&nbsp;/g, ' ').replace(/&#39;/g, "'").replace(/&quot;/g, '"')
    .replace(/[ \t]+\n/g, '\n')
    .replace(/\n{3,}/g, '\n\n')
    .trim();
}

// categoryDir 은 doc.code 를 디렉토리명으로, 비면 _uncategorized.
export function categoryDir(doc) {
  const code = (doc && typeof doc.code === 'string') ? doc.code.trim() : '';
  return code || '_uncategorized';
}

// renderMarkdown 은 doc + 응답예시 + 출처URL 을 카탈로그 md 로 변환.
export function renderMarkdown(doc, responseExample, sourceUrl) {
  const lines = [];
  lines.push(`# ${doc.title || doc.pageURL || 'Untitled'}`, '');
  if (doc.description) lines.push(doc.description, '');

  const endpoint = Array.isArray(doc.urls) && doc.urls.length ? doc.urls[0] : '';
  if (endpoint) lines.push('## Endpoint', '', '`GET ' + endpoint + '`', '');

  const pt = paramsTable(doc.params);
  if (pt) lines.push('## Parameters', '', pt, '');

  const desc = htmlToText(doc.content);
  if (desc) lines.push('## Description', '', desc, '');

  lines.push('## Response (example)', '');
  if (responseExample && responseExample.trim()) {
    lines.push('```json', responseExample.trim(), '```', '');
  } else {
    lines.push('(문서에서 응답 예시를 찾지 못함)', '');
  }

  lines.push(`> 출처: ${sourceUrl} · 카테고리: ${categoryDir(doc)}`, '');
  return lines.join('\n');
}
````

- [ ] **Step 4: 테스트 통과 확인**

Run: `node --test`
Expected: 모든 테스트 PASS.

- [ ] **Step 5: Commit**

```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
git add tools/gendocs/lib.mjs tools/gendocs/lib.test.mjs
git commit -m "feat(gendocs): 순수 md 변환 로직 + 단위테스트"
```

---

## Task 3: 크롤러 (gendocs.mjs) + 소규모 검증

**Files:**
- Create: `tools/gendocs/gendocs.mjs`

- [ ] **Step 1: 크롤러 구현**

Create `tools/gendocs/gendocs.mjs`:
```js
// FMP API docs 카탈로그 생성기.
// 사용: node gendocs.mjs            (전체)
//       LIMIT=3 node gendocs.mjs    (앞 3개만 — 검증용)
import { chromium } from '@playwright/test';
import { writeFile, mkdir } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { renderMarkdown, categoryDir } from './lib.mjs';

const UA = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36';
const SITEMAP = 'https://site.financialmodelingprep.com/sitemap-3.xml';
const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT_ROOT = path.resolve(HERE, '../../docs/api'); // tools/gendocs → repo/docs/api
const DELAY_MS = 500;

async function fetchSitemapUrls() {
  const res = await fetch(SITEMAP, { headers: { 'User-Agent': UA } });
  if (!res.ok) throw new Error(`sitemap fetch failed: ${res.status}`);
  const xml = await res.text();
  const urls = [...xml.matchAll(/<loc>\s*([^<\s]*\/developer\/docs\/stable\/[a-z0-9-]+)\s*<\/loc>/g)].map((m) => m[1]);
  return [...new Set(urls)];
}

async function extractDoc(page) {
  return await page.evaluate(() => {
    const el = document.getElementById('__NEXT_DATA__');
    let doc = null;
    if (el) {
      try { doc = JSON.parse(el.textContent)?.props?.pageProps?.doc ?? null; } catch (_) { doc = null; }
    }
    let resp = null;
    for (const c of document.querySelectorAll('code')) {
      const t = (c.innerText || '').trim();
      if (t.startsWith('[') || t.startsWith('{')) { resp = t; break; }
    }
    return { doc, resp };
  });
}

async function main() {
  const limit = process.env.LIMIT ? parseInt(process.env.LIMIT, 10) : Infinity;
  let urls = await fetchSitemapUrls();
  if (Number.isFinite(limit)) urls = urls.slice(0, limit);
  console.log(`enumerated ${urls.length} doc URLs`);

  const browser = await chromium.launch();
  const ctx = await browser.newContext({ userAgent: UA });
  const page = await ctx.newPage();

  const failures = [];
  const index = {}; // category -> [{title, slug}]
  let ok = 0;

  for (const url of urls) {
    let extracted = { doc: null, resp: null };
    for (let attempt = 0; attempt < 2; attempt++) {
      try {
        await page.goto(url, { waitUntil: 'networkidle', timeout: 60000 });
        await page.waitForTimeout(1500);
        extracted = await extractDoc(page);
        if (extracted.doc) break;
      } catch (e) {
        if (attempt === 1) console.warn(`  retry failed: ${url} (${e.message})`);
      }
    }
    const { doc, resp } = extracted;
    if (!doc) { failures.push(url); console.warn(`  MISS doc: ${url}`); await new Promise((r) => setTimeout(r, DELAY_MS)); continue; }

    const slug = doc.pageURL || url.split('/').pop();
    const cat = categoryDir(doc);
    const dir = path.join(OUT_ROOT, cat);
    await mkdir(dir, { recursive: true });
    await writeFile(path.join(dir, `${slug}.md`), renderMarkdown(doc, resp, url));
    (index[cat] ||= []).push({ title: doc.title || slug, slug });
    ok++;
    if (ok % 25 === 0) console.log(`  ...${ok} written`);
    await new Promise((r) => setTimeout(r, DELAY_MS));
  }

  await browser.close();

  // 인덱스 README
  const idx = ['# FMP API 문서 카탈로그', '', `총 ${ok}개 엔드포인트. FMP stable API 문서를 자동 변환(tools/gendocs).`, ''];
  for (const cat of Object.keys(index).sort()) {
    idx.push(`## ${cat}`, '');
    for (const e of index[cat].sort((a, b) => a.slug.localeCompare(b.slug))) {
      idx.push(`- [${e.title}](${cat}/${e.slug}.md)`);
    }
    idx.push('');
  }
  await mkdir(OUT_ROOT, { recursive: true });
  await writeFile(path.join(OUT_ROOT, 'README.md'), idx.join('\n'));
  await writeFile(path.join(HERE, 'failures.log'), failures.join('\n'));

  console.log(`done: ${ok} ok, ${failures.length} failed (see failures.log)`);
}

main().catch((e) => { console.error(e); process.exit(1); });
```

- [ ] **Step 2: 소규모 검증 실행 (LIMIT=3)**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go/tools/gendocs
LIMIT=3 npm run gen
```
Expected: `enumerated 3 doc URLs` → `done: 3 ok, 0 failed`. `fmp-go/docs/api/` 아래 3개 md 생성.

- [ ] **Step 3: 생성물 육안 확인**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
find docs/api -name '*.md' | head
# 가능하면 company/profile-symbol.md 가 포함됐는지(앞 3개에 없을 수 있음) — 없으면 단건 확인:
LIMIT=1 node tools/gendocs/gendocs.mjs >/dev/null 2>&1 || true
```
생성된 md 하나를 열어 제목·`## Endpoint`(GET URL)·`## Parameters` 표·`## Response (example)`의 ```json``` 블록이 모두 들어있는지 확인한다. 응답 블록이 비어 "(문서에서 응답 예시를 찾지 못함)"이면 `page.waitForTimeout`을 2500ms로 늘려 재시도(렌더 지연). 확인 후 검증 산출물은 다음 태스크의 전체 실행이 덮어쓴다.

- [ ] **Step 4: Commit (크롤러 코드만)**

검증으로 생성된 `docs/api/`의 일부 md는 아직 커밋하지 않는다(전체 실행 후 Task 4에서 일괄 커밋). 생성된 부분 결과는 정리:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
rm -rf docs/api
git add tools/gendocs/gendocs.mjs
git commit -m "feat(gendocs): FMP docs 크롤러 — sitemap 열거 + 렌더 추출 + md 생성"
```

---

## Task 4: 전체 카탈로그 생성

**Files:**
- Create: `docs/api/<category>/<slug>.md` (다수)
- Create: `docs/api/README.md`

- [ ] **Step 1: 전체 실행**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go/tools/gendocs
npm run gen
```
Expected: `enumerated ~274 doc URLs` → 진행 로그(`...25 written` 등) → `done: N ok, M failed`. 수 분~수십 분 소요. `failures.log` 확인.

- [ ] **Step 2: 실패 재시도(있으면)**

`failures.log`에 URL이 남았으면 `page.waitForTimeout`을 2500ms로 늘리거나 1회 더 `npm run gen` 재실행(멱등 — 덮어쓰기). 실패가 소수(예: ≤5)면 허용하고 그 목록을 보고에 남긴다.

- [ ] **Step 3: 산출물 점검**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
echo "md files: $(find docs/api -name '*.md' -not -name 'README.md' | wc -l)"
echo "categories: $(find docs/api -mindepth 1 -maxdepth 1 -type d | wc -l)"
ls docs/api
head -40 docs/api/company/profile-symbol.md
```
Expected: md 파일 수가 열거 수에 근접(예 ≥ 260), 카테고리 디렉토리 다수, `company/profile-symbol.md`에 엔드포인트·파라미터·응답 예시 포함. `docs/api/README.md` 인덱스 존재.

- [ ] **Step 4: Commit (카탈로그)**

```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
git add docs/api
git commit -m "docs(catalog): FMP stable API 전체 엔드포인트 문서 카탈로그 생성"
```

---

## Task 5: 전체 검증

**Files:** 없음(검증 전용)

- [ ] **Step 1: 단위테스트 재확인**

Run: `cd /Users/frankoh/src/workspace_moneyflow/fmp-go/tools/gendocs && node --test`
Expected: lib 단위테스트 전부 PASS.

- [ ] **Step 2: 카탈로그 무결성 점검**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
# README 인덱스의 링크가 실제 파일을 가리키는지 표본 점검
grep -oE '\]\([^)]+\.md\)' docs/api/README.md | head -5
# 빈 응답("찾지 못함")이 과도하지 않은지
echo "missing-response files: $(grep -rl '응답 예시를 찾지 못함' docs/api | wc -l)"
```
Expected: 인덱스 링크 유효, "찾지 못함" 파일이 소수. 과도하면(예 > 20) 렌더 대기시간을 늘려 재생성.

- [ ] **Step 3: 커밋 상태 확인**

Run: `git log --oneline -8 && git status --short`
Expected: Task 1~4 커밋이 `feature/sdk-foundation`에 순서대로, 워킹트리 클린(node_modules/failures.log는 .gitignore).

---

## 자기 점검 메모 (작성자용)

- **node_modules 미커밋**: `.gitignore`로 제외. `package-lock.json`은 커밋(재현성).
- **응답 추출**: `<code>` 중 `[`/`{` 시작 텍스트. 라인번호 거터(`1\n2\n...`)는 `[`/`{`로 시작 안 해 자동 제외. 렌더 지연 시 빈 응답 → 대기시간 상향으로 대응.
- **카테고리**: `doc.code` 그대로. 빈 값 → `_uncategorized`(리뷰 대상).
- **멱등**: 재실행은 덮어쓰기. 실패는 `failures.log` 누적.
- **범위**: 카탈로그 생성까지. SDK 구현(A)·moneyflow 통합(B)은 별도 plan.
- **정중함**: 요청 간 500ms 딜레이로 사이트 과부하 방지.
```
