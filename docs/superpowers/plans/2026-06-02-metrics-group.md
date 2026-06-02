# Metrics 그룹 (v0.10.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `metrics/` 패키지 7 endpoint + 기존 `ratios/` 패키지 RatiosTTM 1개 = 8 endpoint 추가.

**Architecture:** metrics 패키지는 `internal/fetch`(calendar/analyst 컨벤션) + 패키지 helper `listParams`. RatiosTTM 은 ratios 패키지 기존 스타일(GetJSON+ErrNotFound) 따름. 구조체/태그는 `docs/superpowers/specs/2026-06-02-metrics-group-design.md` verbatim.

**Tech Stack:** Go 1.25 generics, internal/fetch, internal/httpclient.

참고: `unset GOROOT` 후 go 실행. 커밋 한국어 `feat(metrics): ...` / `feat(ratios): ...`. 구조체 전체 정의는 스펙 참조(서브에이전트 프롬프트에 verbatim 포함).

---

### Task 1: metrics 패키지 scaffold + KeyMetrics + KeyMetricsTTM

**Files:** Create `metrics/client.go`, `metrics/key_metrics.go`, `metrics/key_metrics_test.go`, `metrics/testdata/key-metrics.json`, `metrics/testdata/key-metrics-ttm.json`.

- [ ] **Step 1:** `metrics/client.go`:
```go
// Package metrics 는 FMP 지표 API sub-client (key-metrics/scores/owner-earnings/EV/segments).
// fmp.Client.Metrics 로 접근.
package metrics

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 지표 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// listParams 는 symbol(필수) + period(비어있지 않으면) + limit(>0) 쿼리 맵.
func listParams(symbol, period string, limit int) map[string]string {
	q := map[string]string{"symbol": symbol}
	if period != "" {
		q["period"] = period
	}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
```
- [ ] **Step 2:** `metrics/key_metrics.go` — 스펙의 `KeyMetrics`(45필드) + `KeyMetricsTTM`(43필드) struct + 메서드:
```go
func (c *Client) KeyMetrics(ctx context.Context, symbol, period string, limit int) ([]KeyMetrics, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[KeyMetrics](ctx, c.http, "/stable/key-metrics", listParams(symbol, period, limit))
}
func (c *Client) KeyMetricsTTM(ctx context.Context, symbol string) ([]KeyMetricsTTM, error) {
	return fetch.ListBySymbol[KeyMetricsTTM](ctx, c.http, "/stable/key-metrics-ttm", symbol)
}
```
import: context, fmt, strings, internal/fetch.
- [ ] **Step 3:** fixtures(1건 배열, AAPL 현실값; key-metrics 는 freeCashFlowToFirm 소수, marketCap 정수). 테스트: KeyMetrics 파싱(MarketCap!=0, FreeCashFlowToFirm!=0, ResearchAndDevelopementToRevenue 태그 파싱) + delegation(path `/stable/key-metrics`, 쿼리 symbol/period/limit) + 빈 symbol 가드. KeyMetricsTTM 파싱(EnterpriseValueTTM!=0) + path.
- [ ] **Step 4:** `unset GOROOT && go test ./metrics/ && go vet ./metrics/ && gofmt -l metrics/`. 커밋 `feat(metrics): 패키지 기반 + KeyMetrics + KeyMetricsTTM`.

### Task 2: FinancialScores + OwnerEarnings + EnterpriseValues

**Files:** Create `metrics/scores.go`, `metrics/owner_earnings.go`, `metrics/enterprise_values.go`, `metrics/scores_test.go`, `metrics/financials_test.go`, testdata `financial-scores.json`, `owner-earnings.json`, `enterprise-values.json`.

- [ ] **Step 1:** `scores.go` — 스펙 `FinancialScores` struct + `func (c *Client) FinancialScores(ctx, symbol string) (*FinancialScores, error) { return fetch.OneBySymbol[FinancialScores](ctx, c.http, "/stable/financial-scores", symbol) }`.
- [ ] **Step 2:** `owner_earnings.go` — 스펙 `OwnerEarning` struct + `OwnerEarnings(ctx, symbol string, limit int)` (guard + List + listParams(symbol,"",limit), path `/stable/owner-earnings`).
- [ ] **Step 3:** `enterprise_values.go` — 스펙 `EnterpriseValue` struct + `EnterpriseValues(ctx, symbol, period string, limit int)` (guard + List + listParams, path `/stable/enterprise-values`).
- [ ] **Step 4:** fixtures + 테스트: FinancialScores 파싱(PiotroskiScore int, AltmanZScore float) + 빈 배열 ErrNotFound + 빈 symbol 가드 / OwnerEarning 파싱(OwnersEarnings, OwnersEarningsPerShare) / EnterpriseValue 파싱(EnterpriseValue, StockPrice) + delegation path.
- [ ] **Step 5:** `go test ./metrics/ && go vet && gofmt -l`. 커밋 `feat(metrics): FinancialScores + OwnerEarnings + EnterpriseValues`.

### Task 3: Revenue Segmentation (지역/제품 공유)

**Files:** Create `metrics/segments.go`, `metrics/segments_test.go`, testdata `revenue-geographic-segmentation.json`, `revenue-product-segmentation.json`.

- [ ] **Step 1:** `segments.go` — 스펙 `RevenueSegment` struct(`Data map[string]int64`, `FiscalYear int`, `ReportedCurrency *string`) + 2 메서드:
```go
func (c *Client) RevenueGeographicSegmentation(ctx context.Context, symbol, period string) ([]RevenueSegment, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[RevenueSegment](ctx, c.http, "/stable/revenue-geographic-segmentation", listParams(symbol, period, 0))
}
func (c *Client) RevenueProductSegmentation(ctx context.Context, symbol, period string) ([]RevenueSegment, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[RevenueSegment](ctx, c.http, "/stable/revenue-product-segmentation", listParams(symbol, period, 0))
}
```
- [ ] **Step 2:** fixtures(스펙 verbatim JSON: AAPL, reportedCurrency null, fiscalYear 2024 숫자, data 다중 키). 테스트: 두 메서드 파싱 — len==1, FiscalYear==2024, ReportedCurrency==nil, len(Data)>0 및 특정 키 값 검증(예: geographic "Americas Segment", product "iPhone") + delegation path 각각 + 빈 symbol 가드 1건.
- [ ] **Step 3:** `go test ./metrics/ && go vet && gofmt -l`. 커밋 `feat(metrics): Revenue 지역/제품 세그먼트`.

### Task 4: ratios.RatiosTTM (기존 패키지 확장)

**Files:** Create `ratios/ratios_ttm.go`, `ratios/ratios_ttm_test.go`, `ratios/testdata/ratios-ttm.json`. ratios 패키지에 testdata 디렉터리 없으면 생성.

- [ ] **Step 1:** `ratios/ratios_ttm.go` — 스펙 `RatioTTM`(59필드) struct + 메서드(ratios 패키지 기존 스타일):
```go
package ratios

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// RatiosTTM 는 종목의 TTM 재무비율을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) RatiosTTM(ctx context.Context, symbol string) ([]RatioTTM, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []RatioTTM
	if err := c.http.GetJSON(ctx, "/stable/ratios-ttm", map[string]string{"symbol": symbol}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
```
(RatioTTM struct 를 같은 파일에 둔다.)
- [ ] **Step 2:** fixture(1건, EnterpriseValueTTM 정수, 나머지 비율) + 테스트: 파싱(GrossProfitMarginTTM!=0, EnterpriseValueTTM!=0, PriceToEarningsRatioTTM!=0) + path `/stable/ratios-ttm` + symbol 쿼리 + 빈 symbol 가드 + 빈 배열 ErrNotFound. 기존 ratios 테스트 패턴 참고(`ratios/ratios_test.go`).
- [ ] **Step 3:** `go test ./ratios/ && go vet ./ratios/ && gofmt -l ratios/`. 커밋 `feat(ratios): RatiosTTM (ratios-ttm)`.

### Task 5: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/metrics/main.go`.

- [ ] **Step 1:** `client.go` — import `metrics`, struct 에 `Metrics *metrics.Client` 추가, NewClient 에 `c.Metrics = metrics.New(hc)`.
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasMetrics`(c.Metrics != nil).
- [ ] **Step 3:** README 커버리지 표 — Metrics 행 신규: `| Metrics | \`client.Metrics\` | KeyMetrics, KeyMetricsTTM, FinancialScores, OwnerEarnings, EnterpriseValues, RevenueGeographicSegmentation, RevenueProductSegmentation — 7 endpoint |`. Ratios 행을 `Ratios, RatiosTTM — 2 endpoint` 로 갱신.
- [ ] **Step 4:** `examples/metrics/main.go` — NewClientFromEnv → KeyMetrics(AAPL,annual,1) 첫 행 일부 출력 + FinancialScores(AAPL) Altman/Piotroski 출력. (분석 예시 패턴은 examples/analyst/main.go 참고.)
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Metrics`: KeyMetrics(AAPL,annual,2) MarketCap!=0 / FinancialScores(AAPL) PiotroskiScore 0~9 / RevenueProductSegmentation(AAPL,annual) len(Data)>0 / c.Ratios.RatiosTTM(AAPL). import metrics 불필요(메서드는 positional). t.Logf 로 KeyMetrics[0] 로그(절대값 타입 확인).
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 전부 통과. 커밋 `feat(metrics): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 8 endpoint 매핑: KeyMetrics/TTM=T1, scores/owner/EV=T2, segments×2=T3, RatiosTTM=T4, 와이어/문서=T5.
- FinancialScores 만 단일 *T(OneBySymbol). 나머지 list.
- RevenueSegment: fiscalYear int, reportedCurrency *string, data map[string]int64.
- ratios 는 internal/fetch 미사용(기존 스타일 유지), metrics 는 internal/fetch 사용.
- 타입 주의: KeyMetrics.FreeCashFlowToFirm float64(소수), FinancialScores.PiotroskiScore int.
