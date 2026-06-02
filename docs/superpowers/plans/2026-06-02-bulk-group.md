# Bulk 그룹 (v0.29.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** internal/httpclient 에 GetRaw 추가 + 신규 `bulk/` 패키지 18 endpoint(원시 CSV `[]byte` 반환).

**Architecture:** httpclient.get 추출(요청/에러매핑 공유) → GetJSON+GetRaw. bulk 메서드는 GetRaw 로 원시 바이트 반환. 파라미터 4형태.

참고: `unset GOROOT`. 커밋 한국어 `feat(bulk): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의(단, CSV 바디 반환).

---

### Task 1: httpclient.GetRaw + bulk scaffold + statement bulk 6종

**Files:** Modify `internal/httpclient/client.go`(get 추출 + GetRaw 추가); add `internal/httpclient/raw_test.go`(GetRaw 단위테스트). Create `bulk/client.go`, `bulk/statements.go`, `bulk/statements_test.go`.

- [ ] **Step 1:** `internal/httpclient/client.go` 리팩터 — 비공개 `get(ctx, path, params) ([]byte, error)` 추출(현 GetJSON 의 요청 빌드 + 비-200 매핑 + 200-envelope 매핑 + 바디 반환). `GetJSON` 은 `body, err := c.get(...)` 후 `json.Unmarshal(body, out)`. 신규:
```go
// GetRaw 는 apikey 주입 GET 후 응답 바디를 원시 바이트로 반환한다(CSV 등 비-JSON 용).
func (c *Client) GetRaw(ctx context.Context, path string, params map[string]string) ([]byte, error) {
	return c.get(ctx, path, params)
}
```
기존 GetJSON 동작/시그니처 보존. `go test ./internal/httpclient/` 기존 테스트 회귀 없어야 함.
- [ ] **Step 2:** `internal/httpclient/raw_test.go` — GetRaw 단위테스트: mock 서버가 CSV `"symbol,price\nAAPL,225.5\n"` 반환 → GetRaw 가 동일 바이트 반환. 비-200(예: 401 + `{"Error Message":"x"}`) → *APIError.
- [ ] **Step 3:** `bulk/client.go`:
```go
// Package bulk 는 FMP 대량 CSV export API sub-client. 모든 메서드는 원시 CSV 바이트를 반환한다.
// fmp.Client.Bulk 로 접근.
package bulk

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

type Client struct {
	http *httpclient.Client
}

func New(http *httpclient.Client) *Client { return &Client{http: http} }

// yearPeriod 는 year+period 필수 bulk 공통 구현.
func (c *Client) yearPeriod(ctx context.Context, path, year, period string) ([]byte, error) {
	if strings.TrimSpace(year) == "" || strings.TrimSpace(period) == "" {
		return nil, fmt.Errorf("fmp: year, period must not be empty")
	}
	return c.http.GetRaw(ctx, path, map[string]string{"year": year, "period": period})
}
```
- [ ] **Step 4:** `bulk/statements.go` — 6 메서드:
```go
func (c *Client) IncomeStatement(ctx context.Context, year, period string) ([]byte, error) {
	return c.yearPeriod(ctx, "/stable/income-statement-bulk", year, period)
}
// IncomeStatementGrowth → /stable/income-statement-growth-bulk
// BalanceSheetStatement → /stable/balance-sheet-statement-bulk
// BalanceSheetStatementGrowth → /stable/balance-sheet-statement-growth-bulk
// CashFlowStatement → /stable/cash-flow-statement-bulk
// CashFlowStatementGrowth → /stable/cash-flow-statement-growth-bulk
```
- [ ] **Step 5:** `bulk/statements_test.go`(헬퍼 calendar 패턴 정의하되 CSV 바디 반환): IncomeStatement("2024","FY") 파싱(반환 []byte 비어있지 않고 헤더 "symbol" 또는 "date" 포함) + delegation(path+year/period) + 빈 year 가드. 나머지 5종 delegation path 각각.
- [ ] **Step 6:** `unset GOROOT && go test ./internal/httpclient/ ./bulk/ && go vet ./bulk/ && gofmt -l bulk/ internal/httpclient/`. 커밋 `feat(bulk): httpclient.GetRaw + 재무제표 bulk 6종`.

### Task 2: 나머지 12종 (part 2 + date 1 + year 1 + 무파라미터 8)

**Files:** Create `bulk/misc.go`, `bulk/misc_test.go`.

- [ ] **Step 1:** `bulk/misc.go` — 12 메서드. part/date/year 가드, 무파라미터 nil:
```go
func (c *Client) Profile(ctx context.Context, part string) ([]byte, error) {
	if strings.TrimSpace(part) == "" {
		return nil, fmt.Errorf("fmp: part must not be empty")
	}
	return c.http.GetRaw(ctx, "/stable/profile-bulk", map[string]string{"part": part})
}
// ETFHolder(part) → /stable/etf-holder-bulk (part 가드)
// EOD(date) → /stable/eod-bulk (date 가드, {date})
// EarningsSurprises(year) → /stable/earnings-surprises-bulk (year 가드, {year})
// RatiosTTM(ctx) → /stable/ratios-ttm-bulk (nil)
// KeyMetricsTTM(ctx) → /stable/key-metrics-ttm-bulk (nil)
// Scores(ctx) → /stable/scores-bulk (nil)
// DCF(ctx) → /stable/dcf-bulk (nil)
// Peers(ctx) → /stable/peers-bulk (nil)
// PriceTargetSummary(ctx) → /stable/price-target-summary-bulk (nil)
// Rating(ctx) → /stable/rating-bulk (nil)
// UpgradesDowngradesConsensus(ctx) → /stable/upgrades-downgrades-consensus-bulk (nil)
```
import: context, fmt, strings, internal/httpclient(이미 client.go). misc.go 는 context/fmt/strings 필요.
- [ ] **Step 2:** `bulk/misc_test.go`(헬퍼 재사용): Profile("0") 파싱([]byte 비어있지 않음) + delegation(path+part) + 빈 part 가드 / EOD("2024-10-22") delegation(path+date) + 빈 date 가드 / EarningsSurprises("2024") delegation(path+year) + 빈 year 가드 / Scores delegation(path, 무파라미터) / Peers·DCF·Rating 등 대표 path 1~2건.
- [ ] **Step 3:** `go test ./bulk/ && go vet && gofmt -l`. 커밋 `feat(bulk): part/date/year/무파라미터 12종`.

### Task 3: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/bulk/main.go`.

- [ ] **Step 1:** `client.go` — import `bulk`, struct 에 `Bulk *bulk.Client`, NewClient 에 `c.Bulk = bulk.New(hc)`. (gofmt -w client.go.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasBulk`.
- [ ] **Step 3:** README 표 행 신규: `| Bulk | \`client.Bulk\` | Profile, ETFHolder, EOD, IncomeStatement(+Growth), BalanceSheetStatement(+Growth), CashFlowStatement(+Growth), EarningsSurprises, RatiosTTM, KeyMetricsTTM, Scores, DCF, Peers, PriceTargetSummary, Rating, UpgradesDowngradesConsensus — 18 endpoint (원시 CSV []byte 반환) |`.
- [ ] **Step 4:** `examples/bulk/main.go` — NewClientFromEnv → Scores(ctx) 바이트 길이 + EOD(어제 날짜) 바이트 길이 출력(time 사용).
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Bulk`: Scores(ctx) len>0 / Peers(ctx) len>0 / EOD(어제) err 체크(time 사용).
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(bulk): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- httpclient.GetRaw 추가(get 추출, GetJSON 회귀 없음).
- 18 메서드 전부 ([]byte, error) 원시 CSV. 파라미터 4형태 가드.
- 마지막 캠페인 그룹.
