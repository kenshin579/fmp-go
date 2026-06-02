# FMP Go SDK — Company 그룹 완성 + `internal/fetch` 공유화 (v0.4.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/company-group`
- 토픽: company 카테고리의 남은 16 endpoint 추가(Profile 은 v0.1.0 에 구현됨) + generic helper 를 `internal/fetch` 공유 패키지로 hoist(quote 리팩터 포함). 전체 API 커버리지 캠페인 2번째 그룹.

## 배경 / 목적

quote 그룹(v0.3.0)으로 16 endpoint + 템플릿(필드 주석/시그니처 규칙/fixture+delegation 테스트/README/examples)이 확립됨. 2번째 그룹으로 **company 완성** 선정 — 이미 Profile 로 시작한 그룹을 17/17 로 마무리해 "덜 된 부분 구현" 의도에 직결, 한 그룹 100% 커버.

동시에 quote 에서 패키지-로컬이던 helper 를 **`internal/fetch` 공유 패키지**로 hoist — 2번째 그룹부터 같은 패턴이 반복되므로 지금 공유화해 이후 26 그룹의 중복을 원천 차단(브레인스토밍 확정).

## 결정 사항 (브레인스토밍)

- **helper 위치**: 신규 `internal/fetch` 공유 패키지. quote 도 이걸 쓰게 리팩터(같은 PR). company 및 이후 그룹은 helper 복제 없이 `fetch.*` 호출.
- **범위**: company 남은 16 endpoint **전부**. struct 재사용 적극(MarketCap/SharesFloat/EmployeeCount/MergerAcquisition/Profile 공용).
- **템플릿 계승**: 필드 한국어 주석, 단일 `*T`(ErrNotFound)/리스트 `[]T` 시그니처, fixture + delegation 테스트, README/examples.
- **릴리스**: `v0.4.0`.

## 섹션 1 — 공유 helper `internal/fetch` + quote 리팩터

신규 `internal/fetch` 패키지:
```go
package fetch

// OneBySymbol — {symbol} 단일 레코드. 빈 symbol 가드, 빈 배열 → httpclient.ErrNotFound.
func OneBySymbol[T any](ctx context.Context, hc *httpclient.Client, path, symbol string) (*T, error)

// ListBySymbol — {symbol} 리스트(시계열/다건). 빈 symbol 가드.
func ListBySymbol[T any](ctx context.Context, hc *httpclient.Client, path, symbol string) ([]T, error)

// ListBySymbols — {symbols:쉼표 join} 배치 리스트. 빈 symbols 가드.
func ListBySymbols[T any](ctx context.Context, hc *httpclient.Client, path string, symbols []string) ([]T, error)

// One — 임의 params 단일 레코드. 빈 배열 → ErrNotFound.
func One[T any](ctx context.Context, hc *httpclient.Client, path string, params map[string]string) (*T, error)

// List — 임의 params 리스트.
func List[T any](ctx context.Context, hc *httpclient.Client, path string, params map[string]string) ([]T, error)
```
- 빈 symbol/symbols 가드는 OneBySymbol/ListBySymbol/ListBySymbols 내장. cik/page/name 등 임의 파라미터 가드는 `One`/`List` 호출 메서드 단 inline.
- `httpclient` 는 internal 패키지이므로 `internal/fetch` 가 import 가능.

**quote 리팩터** (같은 PR):
- quote 의 로컬 `fetchOne`/`fetchBatch`/`fetchList` 제거.
- 공개 메서드가 `fetch.*` 직접 호출:
  - `Quote()` → `fetch.OneBySymbol[Quote](ctx, c.http, "/stable/quote", symbol)`
  - `BatchQuote()` → `fetch.ListBySymbols[Quote](ctx, c.http, "/stable/batch-quote", symbols)`
  - `QuoteShort()`/`BatchQuoteShort()`/`PriceChange()`/`AftermarketQuote()`/`AftermarketTrade()`/`BatchAftermarket*()` → 동일 대응
  - `ExchangeQuotes()` → `fetch.List[QuoteShort](ctx, c.http, "/stable/batch-exchange-quote", map[string]string{"exchange": exchange})`
  - `IndexQuotes()` 등 자산군 → `fetch.List[QuoteShort](ctx, c.http, path, nil)`
- quote 의 기존 18 테스트(fixture/delegation/가드)는 공개 메서드 대상 → **무수정 통과**(리팩터 안전망).

## 섹션 2 — company 16 endpoint → 메서드/struct 매핑

| 파일 | 메서드 | path | helper | 반환 |
|---|---|---|---|---|
| `profile.go`(기존) | `ProfileByCIK(ctx, cik)` | `/stable/profile-cik` | One{cik} | `*Profile` (재사용) |
| `market_cap.go` | `MarketCap(ctx, symbol)` | `/stable/market-capitalization` | OneBySymbol | `*MarketCap` |
| | `HistoricalMarketCap(ctx, symbol)` | `/stable/historical-market-capitalization` | ListBySymbol | `[]MarketCap` |
| | `BatchMarketCap(ctx, symbols...)` | `/stable/market-capitalization-batch` | ListBySymbols | `[]MarketCap` |
| `shares_float.go` | `SharesFloat(ctx, symbol)` | `/stable/shares-float` | OneBySymbol | `*SharesFloat` |
| | `AllSharesFloat(ctx, page)` | `/stable/shares-float-all` | List{page} | `[]SharesFloat` |
| `employees.go` | `EmployeeCount(ctx, symbol)` | `/stable/employee-count` | ListBySymbol | `[]EmployeeCount` |
| | `HistoricalEmployeeCount(ctx, symbol)` | `/stable/historical-employee-count` | ListBySymbol | `[]EmployeeCount` |
| `executives.go` | `KeyExecutives(ctx, symbol)` | `/stable/key-executives` | ListBySymbol | `[]Executive` |
| | `ExecutiveCompensation(ctx, symbol)` | `/stable/governance-executive-compensation` | ListBySymbol | `[]ExecutiveCompensation` |
| | `ExecutiveCompensationBenchmark(ctx, year)` | `/stable/executive-compensation-benchmark` | List{year} | `[]ExecutiveCompensationBenchmark` |
| `peers.go` | `StockPeers(ctx, symbol)` | `/stable/stock-peers` | ListBySymbol | `[]Peer` |
| `notes.go` | `CompanyNotes(ctx, symbol)` | `/stable/company-notes` | ListBySymbol | `[]CompanyNote` |
| `mergers.go` | `LatestMergersAcquisitions(ctx, page)` | `/stable/mergers-acquisitions-latest` | List{page} | `[]MergerAcquisition` |
| | `SearchMergersAcquisitions(ctx, name)` | `/stable/mergers-acquisitions-search` | List{name} | `[]MergerAcquisition` |
| `delisted.go` | `DelistedCompanies(ctx, page)` | `/stable/delisted-companies` | List{page} | `[]DelistedCompany` |

**신규 struct 10개**: `MarketCap` / `SharesFloat` / `EmployeeCount` / `Executive` / `ExecutiveCompensation` / `ExecutiveCompensationBenchmark` / `Peer` / `CompanyNote` / `MergerAcquisition` / `DelistedCompany`. + 기존 `Profile` 재사용.

**struct 재사용**: market-cap/historical/batch → `MarketCap`. shares-float/all → `SharesFloat`. employee-count/historical → `EmployeeCount`. mergers latest/search → `MergerAcquisition`. profile-cik → `Profile`.

**새 시그니처 패턴**:
- 페이지네이션: `(ctx, page int)` → `map[string]string{"page": strconv.Itoa(page)}`.
- cik/name: `(ctx, cik string)` / `(ctx, name string)` — 빈 가드 + `{"cik"/"name": ...}`.
- benchmark: `(ctx, year int)` → `{"year": ...}` (파라미터 형태 구현 시 카탈로그 확정).

## 섹션 3 — struct 정책 · 테스트 · 문서 · 릴리스

### 응답 struct
- 10개 신규 struct 필드는 구현 시 각 `docs/api/company/*.md` 응답 예시 + fixture 와 1:1(faithful). 전 필드 한국어 주석.
- 예시 형태(샘플 기준, 구현 시 확정):
  - `MarketCap{ Symbol, Date string; MarketCap int64 }`
  - `Peer{ Symbol, CompanyName string; Price float64; MktCap int64 }`
  - `EmployeeCount{ Symbol, CIK, CompanyName, FormType, FilingDate, PeriodOfReport, AcceptanceTime, Source string; EmployeeCount int64 }`
  - `DelistedCompany{ Symbol, CompanyName, Exchange, IPODate, DelistedDate string }`
  - `Executive{ Title, Name, CurrencyPay, Gender string; Pay, YearBorn, TitleSince *int64; Active bool }` — nullable 숫자는 `*int64`/`*float64`(fixture 의 null 확인 후 결정).
- nullable 필드(pay/yearBorn 등) 포인터 vs zero 값 — fixture 기반 결정.

### 테스트
- 각 endpoint fixture 단위테스트(faithful 디코딩).
- delegation 테스트: 그룹당 대표 2~3개(symbol/cik/page/name 파라미터 매핑 검증).
- 에러 경로: 단일 빈 배열 ErrNotFound, 빈 symbol/cik/name 가드.
- 통합(`//go:build integration`): FMP_API_KEY 있으면 AAPL market-cap/peers/executives/employee-count + delisted(page) sample.

### `internal/fetch` 자체 테스트
- helper 5종 단위테스트 — httptest stub 으로 path/params/빈배열 ErrNotFound/가드 검증. 공유 helper 가 한 번 검증되면 그룹은 path 매핑만 확인하면 됨.

### 문서 / 릴리스
- README 커버리지 표 Company 행 갱신(17 endpoint 전체).
- `examples/company/main.go` 신규(또는 기존 profile 예시 확장).
- 릴리스 `v0.4.0`.

## 범위 밖 / 후속
- 나머지 26 그룹은 별도 PR(다음: search → news → analyst → calendar → statements 확장 → ...).
- moneyflow 통합은 필요 시점(`go get ...@v0.4.0`).

## 위험 / 주의
- benchmark/shares-float/all-shares-float 의 정확한 파라미터·응답 shape 는 구현 시 카탈로그 확정 — 가정과 다르면 조정(quote 의 exchange=short 케이스처럼).
- nullable 필드(executives pay/yearBorn) 포인터 처리 — fixture 의 null 확인 후 결정.
- quote 리팩터가 같은 PR — quote 18 테스트로 회귀 보증. `internal/fetch` 가 httpclient 를 import(둘 다 internal — 순환 없음 확인).
- 카탈로그 파일명 ≠ 실제 path(예: company-executives.md → `/stable/key-executives`) — 각 .md 의 `GET` 줄 정확히 따름.
