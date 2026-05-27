// FMP API docs 카탈로그 생성기.
// 사용: node gendocs.mjs            (전체)
//       LIMIT=3 node gendocs.mjs    (앞 3개만 — 검증용)
//       RETRY=1 node gendocs.mjs    (failures.log 의 실패 URL만 재처리)
import { chromium } from '@playwright/test';
import { writeFile, mkdir, readFile, readdir } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { renderMarkdown, categoryDir } from './lib.mjs';

const UA = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36';
const SITEMAP = 'https://site.financialmodelingprep.com/sitemap-3.xml';
const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT_ROOT = path.resolve(HERE, '../../docs/api'); // tools/gendocs → repo/docs/api
const DELAY_MS = 1000;

async function fetchSitemapUrls() {
  const res = await fetch(SITEMAP, { headers: { 'User-Agent': UA } });
  if (!res.ok) throw new Error(`sitemap fetch failed: ${res.status}`);
  const xml = await res.text();
  // 영문(locale 접두사 없는) stable 문서 URL만 추출 — 총 274개.
  const urls = [...xml.matchAll(/<loc>\s*(https:\/\/site\.financialmodelingprep\.com\/developer\/docs\/stable\/[a-z0-9-]+)\s*<\/loc>/g)].map((m) => m[1]);
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
    // 첫 글자가 [ 또는 { 인 <code> 블록을 JSON 응답 예시로 간주(라인번호 거터는 제외됨)
    for (const c of document.querySelectorAll('code')) {
      const t = (c.innerText || '').trim();
      if (t.startsWith('[') || t.startsWith('{')) { resp = t; break; }
    }
    return { doc, resp };
  });
}

// writeIndex 는 OUT_ROOT 를 디스크에서 스캔해 README.md 를 재구성한다.
// 증분/재시도 실행 후에도 전체 엔드포인트 목록이 유지된다.
async function writeIndex() {
  let categories;
  try {
    categories = await readdir(OUT_ROOT, { withFileTypes: true });
  } catch {
    categories = [];
  }

  const index = {}; // category -> [{title, slug}]

  for (const entry of categories) {
    if (!entry.isDirectory()) continue;
    const cat = entry.name;
    const catDir = path.join(OUT_ROOT, cat);
    let files;
    try {
      files = await readdir(catDir);
    } catch {
      continue;
    }
    const mdFiles = files.filter((f) => f.endsWith('.md') && f !== 'README.md');
    for (const f of mdFiles) {
      const slug = f.slice(0, -3); // .md 제거
      let title = slug;
      try {
        const content = await readFile(path.join(catDir, f), 'utf8');
        const firstLine = content.split('\n')[0] || '';
        if (firstLine.startsWith('# ')) title = firstLine.slice(2).trim();
      } catch {
        // 읽기 실패 시 slug 그대로 사용
      }
      (index[cat] ||= []).push({ title, slug });
    }
  }

  const totalCount = Object.values(index).reduce((s, arr) => s + arr.length, 0);
  const idx = ['# FMP API 문서 카탈로그', '', `총 ${totalCount}개 엔드포인트. FMP stable API 문서를 자동 변환(tools/gendocs).`, ''];
  for (const cat of Object.keys(index).sort()) {
    idx.push(`## ${cat}`, '');
    for (const e of index[cat].sort((a, b) => a.slug.localeCompare(b.slug))) {
      idx.push(`- [${e.title}](${cat}/${e.slug}.md)`);
    }
    idx.push('');
  }
  await mkdir(OUT_ROOT, { recursive: true });
  await writeFile(path.join(OUT_ROOT, 'README.md'), idx.join('\n'));
  return totalCount;
}

async function main() {
  const limit = process.env.LIMIT ? parseInt(process.env.LIMIT, 10) : Infinity;
  if (process.env.LIMIT && Number.isNaN(limit)) {
    throw new Error(`invalid LIMIT: ${process.env.LIMIT}`);
  }

  let urls;
  if (process.env.RETRY) {
    const raw = await readFile(path.join(HERE, 'failures.log'), 'utf8');
    urls = raw.split('\n').map((l) => l.trim()).filter(Boolean);
    console.log(`retry mode: ${urls.length} URLs from failures.log`);
  } else {
    urls = await fetchSitemapUrls();
    if (Number.isFinite(limit)) urls = urls.slice(0, limit);
    console.log(`enumerated ${urls.length} doc URLs`);
  }

  const browser = await chromium.launch();
  try {
    const ctx = await browser.newContext({ userAgent: UA });
    const page = await ctx.newPage();

    const failures = [];
    let ok = 0;

    for (const url of urls) {
      try {
        let extracted = { doc: null, resp: null };
        for (let attempt = 0; attempt < 3; attempt++) {
          if (attempt > 0) await page.waitForTimeout(3000); // 백오프(스로틀 완화)
          try {
            // domcontentloaded(빠름) + 고정 대기. networkidle 은 bulk 페이지에서
            // 네트워크가 계속 활성이라 60s 타임아웃 → 누락. __NEXT_DATA__ 는 SSR HTML 에
            // 있고 응답 예시는 JS 렌더라 고정 대기로 충분.
            await page.goto(url, { waitUntil: 'domcontentloaded', timeout: 30000 });
            await page.waitForTimeout(2500);
            extracted = await extractDoc(page);
            if (extracted.doc) break;
          } catch (e) {
            if (attempt === 2) console.warn(`  retry failed: ${url} (${e.message})`);
          }
        }
        const { doc, resp } = extracted;
        if (!doc) { failures.push(url); console.warn(`  MISS doc: ${url}`); continue; }

        const slug = doc.pageURL || url.split('/').pop();
        const cat = categoryDir(doc);
        const dir = path.join(OUT_ROOT, cat);
        await mkdir(dir, { recursive: true });
        await writeFile(path.join(dir, `${slug}.md`), renderMarkdown(doc, resp, url));
        ok++;
        if (ok % 25 === 0) console.log(`  ...${ok} written`);
      } catch (e) {
        failures.push(url);
        console.warn(`  ERROR: ${url} (${e.message})`);
      } finally {
        await new Promise((r) => setTimeout(r, DELAY_MS));
      }
    }

    // 디스크 스캔으로 전체 인덱스 재구성 (증분/재시도 실행에도 완전한 목록 유지)
    const totalIndexed = await writeIndex();
    await writeFile(path.join(HERE, 'failures.log'), failures.join('\n'));

    console.log(`done: ${ok} ok, ${failures.length} failed (see failures.log)`);
    console.log(`index: ${totalIndexed} total entries written to docs/api/README.md`);
  } finally {
    await browser.close();
  }
}

main().catch((e) => { console.error(e); process.exit(1); });
