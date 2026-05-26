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
