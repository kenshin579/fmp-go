# FMP Go SDK — Bulk 그룹 (v0.29.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/bulk-group`
- 토픽: FMP `bulk` 카테고리 18 endpoint(대량 CSV export). 캠페인 27번째(마지막) 그룹.

## 배경 / 결정 사항
- bulk endpoint 는 실제 **CSV(text/csv) 대량 export** (문서 예시는 JSON 렌더이나 수치가 전부 문자열 → CSV 유래). 거대 파일이라 타입 구조체 대신 **원시 바이트 반환**이 적합.
- `internal/httpclient` 에 `GetRaw(ctx, path, params) ([]byte, error)` 추가(공통 요청/에러매핑 추출, JSON 디코딩 생략). GetJSON 동작 보존.
- 신규 `bulk/` 패키지. 18 메서드 모두 `([]byte, error)`(원시 CSV) 반환.
- 파라미터 4형태: `part`(2), `date`(1), `year`+`period`(6), `year`(1), 무파라미터(8).
- 릴리스 `v0.29.0`.

## httpclient 변경
```go
// get 은 apikey 주입 GET 후 바디([]byte)를 반환. 비-200/에러 envelope → *APIError.
func (c *Client) get(ctx context.Context, path string, params map[string]string) ([]byte, error) { ... }
// GetJSON 은 get 후 json.Unmarshal.
// GetRaw 는 get 결과(원시 바이트)를 그대로 반환.
func (c *Client) GetRaw(ctx context.Context, path string, params map[string]string) ([]byte, error) {
	return c.get(ctx, path, params)
}
```
기존 GetJSON 의 동작(비-200 매핑, 200-envelope 매핑, 디코딩)은 그대로. `get` 은 비-200 + 200-envelope 검사까지 수행하고 바디 반환. (CSV 바디는 json.Unmarshal envelope 검사에서 무시됨 → 정상.)

## 패키지 구조 + endpoint 매핑
| 메서드 | path | 파라미터 |
|---|---|---|
| `Profile(ctx, part)` | `/stable/profile-bulk` | part(필수) |
| `ETFHolder(ctx, part)` | `/stable/etf-holder-bulk` | part(필수) |
| `EOD(ctx, date)` | `/stable/eod-bulk` | date(필수) |
| `IncomeStatement(ctx, year, period)` | `/stable/income-statement-bulk` | year/period(필수) |
| `IncomeStatementGrowth(ctx, year, period)` | `/stable/income-statement-growth-bulk` | year/period |
| `BalanceSheetStatement(ctx, year, period)` | `/stable/balance-sheet-statement-bulk` | year/period |
| `BalanceSheetStatementGrowth(ctx, year, period)` | `/stable/balance-sheet-statement-growth-bulk` | year/period |
| `CashFlowStatement(ctx, year, period)` | `/stable/cash-flow-statement-bulk` | year/period |
| `CashFlowStatementGrowth(ctx, year, period)` | `/stable/cash-flow-statement-growth-bulk` | year/period |
| `EarningsSurprises(ctx, year)` | `/stable/earnings-surprises-bulk` | year(필수) |
| `RatiosTTM(ctx)` | `/stable/ratios-ttm-bulk` | — |
| `KeyMetricsTTM(ctx)` | `/stable/key-metrics-ttm-bulk` | — |
| `Scores(ctx)` | `/stable/scores-bulk` | — |
| `DCF(ctx)` | `/stable/dcf-bulk` | — |
| `Peers(ctx)` | `/stable/peers-bulk` | — |
| `PriceTargetSummary(ctx)` | `/stable/price-target-summary-bulk` | — |
| `Rating(ctx)` | `/stable/rating-bulk` | — |
| `UpgradesDowngradesConsensus(ctx)` | `/stable/upgrades-downgrades-consensus-bulk` | — |

파일: `bulk/client.go`(New + helpers), `bulk/statements.go`(year+period 6), `bulk/misc.go`(part/date/year/none 12).
- 모든 메서드 `([]byte, error)`. 필수 파라미터 빈값 가드.
- 반환은 원시 CSV 바이트 — 호출자가 `encoding/csv` 등으로 파싱.

## 루트 Client 와이어
```go
Assets *assets.Client
Bulk   *bulk.Client // 대량 CSV export
```
`c.Bulk = bulk.New(hc)`. `TestNewClient_HasBulk`.

## 시그니처 규칙
- 무파라미터 8: `(ctx)` → `c.http.GetRaw(ctx, path, nil)`.
- part 2: `(ctx, part string)` → part 가드 + {part}.
- date 1: `(ctx, date string)` → date 가드 + {date}.
- year+period 6: `(ctx, year, period string)` → 둘 다 가드 + {year, period}.
- year 1: `(ctx, year string)` → year 가드 + {year}.

## 테스트
- httpclient: GetRaw 단위테스트(mock 서버 CSV 바디 → 바이트 그대로; 비-200 → APIError).
- bulk: 각 메서드 mock 서버에 CSV 텍스트 반환 → `[]byte` 비어있지 않고 헤더 포함 확인. delegation: Profile(part) path+part / EOD(date) path+date / IncomeStatement(year,period) path+쿼리 / Scores path(무파라미터).
- 가드: Profile 빈 part, EOD 빈 date, IncomeStatement 빈 year, EarningsSurprises 빈 year (대표).
- 통합(`//go:build integration`): EOD(어제 날짜) len>0(또는 err 체크) / Scores len>0 / Peers len>0 — CSV 바이트 비어있지 않음 확인.

## 문서 / 릴리스
- README Bulk 행(18 endpoint) + 비고: 원시 CSV(`[]byte`) 반환.
- `examples/bulk/main.go` — Scores + EOD(어제) 바이트 길이 출력.
- 릴리스 `v0.29.0`. 이 그룹으로 FMP 카테고리 커버리지 캠페인 사실상 완료.

## 범위 밖 / 위험
- bulk 는 타입 파싱 안 함(원시 CSV) — 18종 CSV 스키마 파싱은 호출자/후속 과제.
- 만약 일부 bulk 가 JSON 이면 GetRaw 바이트로도 동작(호출자가 처리). content-type 무관하게 바이트 반환.
- 캠페인 마무리 그룹.
