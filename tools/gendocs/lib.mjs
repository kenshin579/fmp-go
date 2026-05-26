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
