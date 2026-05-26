# FMP API docs 카탈로그 설계

- 작성일: 2026-05-26
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go` (워크스페이스 `fmp-go/`, branch `feature/sdk-foundation`)
- 토픽: Financial Modeling Prep(FMP) stable API 전체 엔드포인트 문서를 md로 충실히 변환한 카탈로그

## 배경 / 목적

FMP Go SDK(서브프로젝트 A)를 카테고리 단위로 점진 구현하기에 앞서, **FMP stable API 전체
문서를 md 카탈로그로 먼저 갖춘다.** 목적은 "문서를 있는 그대로 md로 충실히 변환"한 전체
API 레퍼런스를 확보하는 것. 이후 SDK 구현 시 이 카탈로그를 1차 참조로 쓴다.

작업 순서(재분해):
- **A0 (본 스펙)** — FMP 전체 API docs 카탈로그 생성. 코드(SDK) 없음.
- **A** — FMP Go SDK: moneyflow 필요분(Company Profile) 먼저, 이후 점진 확장.
  (계획 `docs/superpowers/plans/2026-05-26-fmp-go-sdk-foundation.md` 존재 — docs 캡처
  스텝은 본 카탈로그 참조로 단순화 예정.)
- **B** — moneyflow 통합. SDK v0.1.0 릴리스 후.

## 사전 조사 결과 (확정 사실)

- **열거 소스**: `https://site.financialmodelingprep.com/sitemap.xml`은 sitemap index이며,
  `sitemap-3.xml`에 stable 문서 URL **274개**(`/developer/docs/stable/<slug>`)가 전부 들어
  있다. (UA 헤더 필요 — 자동 요청은 403, 브라우저 UA는 200.)
- **페이지 구조**: 각 doc 페이지는 Next.js SSR. `<script id="__NEXT_DATA__">`의
  `props.pageProps.doc`에 구조화 메타데이터가 있다:
  - `title` (예: "Company Profile Data")
  - `description` (요약)
  - `urls` (실제 엔드포인트 URL+예시 파라미터, 예:
    `["https://financialmodelingprep.com/stable/profile?symbol=AAPL"]`)
  - `params` (`{query:{header:[...], rows:[[...]]}}` 파라미터 표)
  - `code` (**카테고리**, 예: `company`)
  - `subcode` (예: `profile`), `schemaName`, `content`(HTML 설명), `pageURL`(slug)
- **응답 예시**: 응답 JSON은 SSR HTML엔 없고 **클라이언트 JS가 렌더**한다("Response" 코드
  블록). 따라서 "문서 그대로" 캡처하려면 **JS를 실행하는 헤드리스 브라우저(Playwright)**가
  필요. 렌더 결과에 샘플 응답 JSON이 포함됨(문서용 샘플이라 API 키 불필요).

## 결정 사항 (브레인스토밍)

- **생성 방식**: Node + Playwright 크롤러로 274개 페이지를 렌더링해 md로 변환(접근법 A).
  curl+`__NEXT_DATA__`만으로는 응답 예시가 빠져 "문서 그대로" 요건 미달이므로 채택 안 함.
- **응답 필드 포함**: 포함. 단 **라이브 데이터 API 호출이 아니라** 문서 페이지가 렌더하는
  샘플 응답을 캡처(키 불필요, premium 402 영향 없음).
- **생성기 위치**: `fmp-go/tools/gendocs/`에 self-contained Node 프로젝트(`package.json` +
  `gendocs.mjs`). Go 레포지만 docs 생성 전용 도구라 분리 디렉토리에 둔다.
- **출력**: `fmp-go/docs/api/<category>/<slug>.md` (+ `docs/api/README.md` 인덱스).

## 생성기 설계 (`tools/gendocs/`)

```
fmp-go/tools/gendocs/
├── package.json        # @playwright/test 의존, "gen" 스크립트
├── gendocs.mjs         # 크롤러 + md 렌더러
└── failures.log        # (런타임) 실패 URL 기록
```

흐름:
1. **열거**: `https://site.financialmodelingprep.com/sitemap-3.xml`을 브라우저 UA로 받아
   `/developer/docs/stable/<slug>` URL을 모두 추출(중복 제거). 기대 ≈ 274개.
2. **페이지 처리**(각 URL):
   - Playwright `page.goto(url, {waitUntil:'networkidle'})` + 짧은 대기(응답 블록 렌더).
   - `__NEXT_DATA__`에서 `doc` 객체 파싱 → `title, description, urls, params, code, subcode,
     content, pageURL`.
   - 렌더된 DOM에서 "Response" 코드 블록의 응답 예시 JSON 텍스트 추출(라인번호 거터 제외,
     코드 텍스트만).
   - md 렌더 → `docs/api/<code>/<pageURL>.md` 기록(디렉토리 자동 생성, 덮어쓰기).
3. **정중함/안정성**: 요청 간 딜레이(기본 500ms), 페이지 실패 시 1회 재시도, 최종 실패 URL은
   `failures.log`에 기록하고 계속 진행. 전 과정 멱등(재실행 시 덮어쓰기).
4. **인덱스**: 처리 완료 후 카테고리(`code`)별로 묶어 `docs/api/README.md`에 카테고리 헤더 +
   엔드포인트 목록(제목 + 상대 링크) 생성.

실행: `cd tools/gendocs && npm install && npm run gen` (Playwright 브라우저 설치 필요 시
`npx playwright install chromium`).

## md 템플릿 (엔드포인트당)

```markdown
# <title>

<description>

## Endpoint

`GET <urls[0]>`

## Parameters

| <params.query.header[0]> | <header[1]> | <header[2]> |
|---|---|---|
| <rows...>                |             |             |

## Description

<content 를 HTML→md 변환(또는 텍스트화)>

## Response (example)

​```json
<렌더된 응답 예시 JSON>
​```

> 출처: <doc page URL> · 카테고리: <code>
```

- `params`가 없으면 Parameters 섹션 생략. `content`가 비면 Description 생략. 응답 블록을 못
  찾으면 "Response (example)" 섹션에 "(문서에서 응답 예시를 찾지 못함)" 표기.

## 카테고리 조직

- 디렉토리 = `doc.code`(예: `company`, `quote`, `statements`, `calendar`, `etf`, `crypto`,
  `forex`, `economics`, `sec-filings` 등 FMP가 부여하는 코드 그대로). 파일명 = `doc.pageURL`
  (slug). 경로: `docs/api/<code>/<slug>.md`.
- `code`가 비어 있는 페이지는 `docs/api/_uncategorized/<slug>.md`로 두고 failures/리뷰 대상.

## 에러 / 부분 처리

- 페이지 구조가 예상과 다르거나(`__NEXT_DATA__.doc` 부재) 응답 블록이 없으면, 가능한 필드만
  으로 md를 만들고 누락은 명시 표기. 완전 실패(네트워크/타임아웃, 재시도 후)는
  `failures.log`에 URL 기록 후 다음으로 진행 — 전체 작업을 중단시키지 않는다.
- 274개 중 일부 누락은 허용(재실행으로 보강). 성공/실패 카운트를 콘솔에 요약.

## 검증

- 생성 후 `find docs/api -name '*.md' | wc -l`이 열거된 URL 수에 근접(예: ≥ 260)하는지 확인.
- `failures.log`가 비었거나 소수인지 확인, 남으면 재실행.
- 샘플 검수: `docs/api/company/profile-symbol.md`에 엔드포인트 URL·파라미터 표·응답 예시
  JSON이 모두 포함됐는지 육안 확인. `docs/api/README.md` 카테고리 인덱스 링크 유효성 점검.

## 범위 밖 / 후속

- SDK 구현(서브프로젝트 A) 및 moneyflow 통합(B) — 본 스펙 범위 아님.
- 응답 스키마의 타입 엄밀 검증 — 카탈로그는 "문서 그대로"의 샘플 응답까지. SDK 구현 시
  실호출 fixture로 타입 확정(SDK 계획 Task 2).
- legacy(v3/v4) 문서 — stable만 대상.
- 생성기의 CI 자동화/스케줄 — 수동 재실행으로 충분, 후속.

## 위험 / 주의

- FMP docs 사이트는 자동 요청을 403 차단 → Playwright/HTTP 모두 브라우저 User-Agent 사용.
- 응답 예시는 **클라이언트 렌더** → 반드시 JS 실행(headless browser). 단순 curl로는 응답 누락.
- 274회 렌더링은 수 분~수십 분 소요. 딜레이로 사이트에 과부하 주지 않도록 함.
- Playwright 브라우저(chromium) 설치 필요 — 환경에 따라 `npx playwright install chromium`.
- 사이트 구조(`__NEXT_DATA__.doc` 스키마, "Response" 블록 DOM)는 FMP가 바꿀 수 있음 →
  생성기는 구조 변경 시 셀렉터/경로만 수정하면 되도록 파싱부를 작게 분리.
