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
