# FMP Go SDK — 기반 + Company Profile 설계

- 작성일: 2026-05-26
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go` (워크스페이스 `fmp-go/`, 이미 클론됨 — 빈 README + 초기 커밋만)
- 토픽: Financial Modeling Prep(FMP) API의 Go 클라이언트 라이브러리 — 확장 가능한 기반 + 첫 카테고리(Company Profile)

## 배경 / 목적

moneyflow 종목 상세의 **해외(US) 회사 개요**(대표자·설립일·홈페이지·기업 서술·섹터)는
국내(DART)와 달리 동등한 무료 정형 소스가 없어 미구현 상태다. 데이터 소스로 FMP를
채택하되, moneyflow에서 FMP를 직접 호출하지 않고 **재사용 가능한 Go SDK**(opendart·
korea-investment-stock와 동급의 독립 라이브러리)로 만들어 다른 프로젝트/개발자도 쓸 수
있게 한다.

**장기 목표는 FMP 전체 API 커버리지**다. 단, FMP는 신뢰할 만한 공식 OpenAPI 스펙을
공개하지 않으므로(비공식 커뮤니티 스펙만 존재, 최신성·정확성 미보장), codegen 대신
**opendart 방식의 수작업 점진 확장**으로 간다: 확장 가능한 뼈대를 세우고 카테고리 단위로
손으로 채워 전체 API로 키운다. 본 스펙의 범위는 **기반 + Company 카테고리(Profile
엔드포인트)** = v0.1.0이다.

## 결정 사항 (브레인스토밍)

- **빌드 전략**: 수작업 점진 확장(opendart 패턴). codegen 미채택(공식 스펙 부재 → 품질·
  유지보수 리스크). 전체 API는 아키텍처가 지향하는 목표이며 카테고리별 누적 릴리스.
- **API 기준선**: FMP **stable** 엔드포인트(`https://financialmodelingprep.com/stable/...`).
  legacy v3/v4는 점차 deprecated이라 채택하지 않음.
- **인증**: `FMP_API_KEY` 환경변수 + `fmp.NewClient(apiKey)` / `fmp.NewClientFromEnv()`.
  API 키는 모든 요청에 `apikey` 쿼리 파라미터로 자동 주입.
- **v1 커버리지**: `company` 서비스 패키지의 **Profile** 엔드포인트 하나.
- **docs md 생성**: 별도 서브프로젝트 A0(FMP API docs 카탈로그)에서 전체 엔드포인트를 md로
  먼저 캡처한다(설계 `2026-05-26-fmp-docs-catalog-design.md`). 본 SDK 구현은 그 카탈로그의
  `docs/api/company/profile-symbol.md`를 1차 참조로 쓰고, 응답 타입은 실호출 fixture로 확정.
- **문서/계획 위치**: `fmp-go/docs/superpowers/specs|plans/`.

## 아키텍처 (opendart 패턴 차용)

```
fmp-go/
├── go.mod                  # module github.com/kenshin579/fmp-go, go 1.25+
├── client.go               # fmp.Client — 단일 진입점, 서비스 서브클라이언트 보유
├── config.go               # Option (timeout/baseURL/httpClient)
├── from_env.go             # NewClientFromEnv() — FMP_API_KEY
├── errors.go               # APIError, ErrNotFound 등 sentinel/타입
├── company/
│   ├── client.go           # company.New(hc) *Client; Profile(ctx, symbol)
│   ├── profile.go          # Profile 구조체 + 매핑
│   └── profile_test.go     # fixture 기반 파싱/매핑 테스트
├── internal/
│   └── httpclient/
│       ├── client.go       # baseURL, apikey 주입, timeout, GET+JSON 디코딩, 에러 매핑
│       └── client_test.go  # httptest.Server 스텁 단위테스트
├── docs/api/company/profile-symbol.md   # (A0 카탈로그에서 생성)
├── examples/               # 실행 가능한 사용 예시
├── scripts/release.sh      # 태그 + GitHub Release
├── integration_test.go     # build tag integration, FMP_API_KEY 있을 때만
└── README.md               # 설치/사용/커버리지 표
```

- **`fmp.Client`**: 단일 진입점. 내부에 `*httpclient.Client`를 두고 서비스 서브클라이언트를
  필드로 노출. v1은 `Company *company.Client` 하나. 이후 `Quote`/`Statements`/… 누적.
  ```go
  type Client struct {
      http    *httpclient.Client
      Company *company.Client
  }
  func NewClient(apiKey string, opts ...Option) (*Client, error)
  func NewClientFromEnv(opts ...Option) (*Client, error) // FMP_API_KEY
  ```
- **`internal/httpclient`**: baseURL 기본 `https://financialmodelingprep.com`. 모든 요청에
  `apikey` 쿼리 자동 주입. timeout 기본 30s, 재시도 없음(v1). `GET(ctx, path, query)` →
  JSON 디코딩. 비-200 또는 FMP 에러 바디를 `errors.go` 타입으로 매핑.
- **`company` 패키지**: 도메인 경계. `company.New(hc)` → `*company.Client`. 메서드
  `Profile(ctx, symbol string) (*Profile, error)`.

## v1 데이터 흐름 — Company Profile

- 요청: `GET /stable/profile?symbol={symbol}&apikey={key}`.
- 응답: FMP는 **객체 배열**을 반환(단일 종목이면 요소 1개). SDK는 배열을 디코딩해 **첫
  요소를 `*Profile`로 반환**, 빈 배열이면 `ErrNotFound`.
- `Profile` 구조체: FMP 응답 필드를 타입대로 매핑. 최소 포함 필드(moneyflow가 소비):
  `Symbol, CompanyName, CEO, IPODate, Website, Description, Sector, Industry, Country,
  Exchange, ExchangeFullName, Image, Currency` + 시세/식별자성 필드(`Price, MarketCap,
  CIK, ISIN, CUSIP, FullTimeEmployees, Phone, Address, City, State, Zip, IsEtf,
  IsActivelyTrading, IsAdr, IsFund` 등). 정확한 전체 필드·타입은 구현 시 카탈로그 + 실호출
  fixture로 확정.
- 금액/비율은 적절한 수치 타입. FMP의 빈 문자열/`null` 허용 필드는 zero-value로 안전 디코딩.

## 에러 처리

- `errors.go`:
  - `ErrNotFound` (sentinel) — 빈 배열/결과 없음.
  - `APIError{StatusCode int, Message string}` — 비-200 응답 또는 FMP 에러 바디
    (예: `{"Error Message": "..."}`). 402(요금제 한도)·429(rate limit)·401(키 오류)을
    StatusCode로 식별 가능.
- `NewClient`는 apiKey 빈 값이면 에러. `NewClientFromEnv`는 `FMP_API_KEY` 미설정 시 에러.
- httpclient는 네트워크/디코딩 에러를 `%w`로 래핑해 전달.

## 테스트

- **`internal/httpclient`**: `httptest.Server` 스텁 — apikey 쿼리 주입 확인, JSON 디코딩,
  비-200/에러바디 → `APIError` 매핑, 타임아웃.
- **`company.Profile`**: 저장한 실제 응답 fixture(JSON)로 파싱·매핑 단위테스트. 배열→단일
  변환, 빈 배열→`ErrNotFound`, 필드 매핑 정확성(특히 moneyflow 소비 필드).
- **`integration_test.go`** (build tag `integration`): `FMP_API_KEY` 있을 때만 실제 AAPL
  프로필 조회로 계약 검증. 기본 `go test ./...`에서 제외(비용·키 의존 회피).

## 릴리스 & 문서

- `scripts/release.sh` — opendart/korea-investment-stock와 동일 절차: main/clean 검증 →
  `go build/vet/test` → 모듈 검증 → `git tag vX.Y.Z` push → `gh release create
  --generate-notes`.
- `README.md` — 설치(`go get github.com/kenshin579/fmp-go@v0.1.0`), 사용 예시
  (`NewClientFromEnv` → `Company.Profile`), 커버리지 표(v1: Company / Profile), 인증 안내.
- `examples/` — 실행 가능한 프로필 조회 예시.
- 완료 후 `main`에 머지하고 **`v0.1.0` 태그 릴리스** → 서브프로젝트 B(moneyflow 통합)가
  이 태그를 `go.mod`에 추가해 소비.

## 범위 밖 (후속)

- **서브프로젝트 A0 — FMP API docs 카탈로그**: 전체 엔드포인트 문서를 md로 변환(본 SDK보다
  먼저 수행). 별도 spec(`2026-05-26-fmp-docs-catalog-design.md`).
- **서브프로젝트 B — moneyflow 해외 회사 개요 통합**: `stock_company_us` 테이블 + US
  `CompanySource`(FMP) + `GetDetail` US 회사 개요 채움. 별도 spec→plan 사이클. 본 SDK
  v0.1.0 릴리스 후 진행.
- FMP의 다른 카테고리(Quote/Statements/Calendar/ETF/Crypto/Forex/Insider/SEC 등) — 전체
  API 목표의 후속 누적 릴리스. 본 스펙 범위 아님.
- 재시도/백오프, rate-limit 자동 스로틀링, 캐싱 — v1 미포함(소비 측 moneyflow가 lazy+TTL
  캐싱 담당). 필요 시 후속.

## 미확정 / 후속

- **Profile 전체 필드·타입 확정**: 구현 단계에서 A0 카탈로그의 `profile-symbol.md` + 실제
  AAPL 응답 fixture로 `Profile` 구조체를 확정한다. 본 스펙은 moneyflow 소비 필드 보장에 초점.
- **stable `/profile` 응답이 배열인지 단일 객체인지** 실제 호출로 최종 확인(설계는 배열
  가정 — legacy v3가 배열이었음). 단일 객체면 매핑만 조정.
- **숫자 필드의 빈 문자열/null 처리**: FMP가 일부 필드를 `""`로 줄 수 있어, 수치 타입은
  관용 디코딩(빈→0) 필요 여부를 fixture로 고정.

## 위험 / 주의

- FMP docs 사이트는 자동 요청을 403 차단 → docs 캡처는 브라우저 User-Agent 사용(A0에서 처리).
- FMP legacy(v3/v4) vs stable 혼재 — 본 SDK는 stable만. 문서/URL에서 버전 혼동 금지.
- 라이브러리 미릴리스(태그 없음) 상태로는 moneyflow가 소비 불가 → B 진행 전 **v0.1.0 태그
  필수**(moneyflow CLAUDE.md의 내부 Go 라이브러리 소비 규칙과 동일).
