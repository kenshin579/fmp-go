# Insider Trades 그룹 (v0.17.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `insidertrades/` 패키지 6 endpoint, 5 구조체.

**Architecture:** internal/fetch(List). InsiderTrade(latest+search 공유) + SearchParams. 나머지 4 struct. 구조체는 스펙 verbatim. AcquisitionOwnership 전 필드 string.

참고: `unset GOROOT`. 커밋 한국어 `feat(insidertrades): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 scaffold + InsiderTrade(latest+search) + SearchParams

**Files:** Create `insidertrades/client.go`, `insidertrades/trades.go`, `insidertrades/trades_test.go`, testdata `latest.json`, `search.json`.

- [ ] **Step 1:** `insidertrades/client.go`:
```go
// Package insidertrades 는 FMP 내부자 거래 API sub-client.
// fmp.Client.InsiderTrades 로 접근.
package insidertrades

import "github.com/kenshin579/fmp-go/internal/httpclient"

type Client struct {
	http *httpclient.Client
}

func New(http *httpclient.Client) *Client { return &Client{http: http} }
```
- [ ] **Step 2:** `insidertrades/trades.go` — 스펙의 `InsiderTrade`(16필드) struct + `SearchParams` struct + queryParams 메서드 + 2 메서드. import: context, strconv, internal/fetch.
```go
// SearchParams.queryParams 는 빈 문자열 제외, page 항상 포함(0 허용), limit>0.
func (p SearchParams) queryParams() map[string]string {
	q := map[string]string{"page": strconv.Itoa(p.Page)}
	if p.Limit > 0 {
		q["limit"] = strconv.Itoa(p.Limit)
	}
	if p.Symbol != "" {
		q["symbol"] = p.Symbol
	}
	if p.ReportingCik != "" {
		q["reportingCik"] = p.ReportingCik
	}
	if p.CompanyCik != "" {
		q["companyCik"] = p.CompanyCik
	}
	if p.TransactionType != "" {
		q["transactionType"] = p.TransactionType
	}
	return q
}

func (c *Client) LatestInsiderTrades(ctx context.Context, date string, page, limit int) ([]InsiderTrade, error) {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	if date != "" {
		q["date"] = date
	}
	return fetch.List[InsiderTrade](ctx, c.http, "/stable/insider-trading/latest", q)
}

func (c *Client) SearchInsiderTrades(ctx context.Context, p SearchParams) ([]InsiderTrade, error) {
	return fetch.List[InsiderTrade](ctx, c.http, "/stable/insider-trading/search", p.queryParams())
}
```
- [ ] **Step 3:** fixtures `latest.json`/`search.json` — 1건 InsiderTrade(16필드 모두, securitiesTransacted/price 비0, directOrIndirect:"D", formType:"4", transactionType:"P-Purchase").
- [ ] **Step 4:** `trades_test.go`(헬퍼 calendar 패턴 정의): LatestInsiderTrades("",0,5) 파싱(SecuritiesTransacted!=0, Price!=0, DirectOrIndirect!="", FormType!="") + delegation(path+page/limit) / SearchInsiderTrades(SearchParams{Symbol:"AAPL",TransactionType:"P-Purchase",Limit:10}) delegation(path `/stable/insider-trading/search`+symbol/transactionType/page/limit).
- [ ] **Step 5:** `unset GOROOT && go test ./insidertrades/ && go vet ./insidertrades/ && gofmt -l insidertrades/`. 커밋 `feat(insidertrades): InsiderTrade(latest/search) + SearchParams`.

### Task 2: TransactionTypes + Statistics + AcquisitionOwnership + SearchReportingName

**Files:** Create `insidertrades/misc.go`, `insidertrades/misc_test.go`, testdata `transaction-type.json`, `statistics.json`, `acquisition-ownership.json`, `reporting-name.json`.

- [ ] **Step 1:** `misc.go` — 스펙 `InsiderTransactionType`/`TradeStatistics`/`AcquisitionOwnership`/`ReportingName` struct + 4 메서드. import: context, fmt, strconv, strings, internal/fetch.
```go
func (c *Client) TransactionTypes(ctx context.Context) ([]InsiderTransactionType, error) {
	return fetch.List[InsiderTransactionType](ctx, c.http, "/stable/insider-trading-transaction-type", nil)
}
func (c *Client) Statistics(ctx context.Context, symbol string) ([]TradeStatistics, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[TradeStatistics](ctx, c.http, "/stable/insider-trading/statistics", map[string]string{"symbol": symbol})
}
func (c *Client) AcquisitionOwnership(ctx context.Context, symbol string, limit int) ([]AcquisitionOwnership, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := map[string]string{"symbol": symbol}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[AcquisitionOwnership](ctx, c.http, "/stable/acquisition-of-beneficial-ownership", q)
}
func (c *Client) SearchReportingName(ctx context.Context, name string) ([]ReportingName, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[ReportingName](ctx, c.http, "/stable/insider-trading/reporting-name", map[string]string{"name": name})
}
```
(메서드명 `AcquisitionOwnership` == 타입명, Go 에서 합법.)
- [ ] **Step 2:** fixtures: transaction-type `[{transactionType:"A-Award"}]`, statistics `[{symbol,cik,year,quarter,acquiredTransactions,disposedTransactions,acquiredDisposedRatio,totalAcquired,totalDisposed,averageAcquired,averageDisposed,totalPurchases,totalSales}]`(averageAcquired 소수), acquisition-ownership `[{...15필드 전부 문자열, percentOfClass:"5.1"}]`, reporting-name `[{reportingCik,reportingName}]`.
- [ ] **Step 3:** `misc_test.go`(헬퍼 재사용): TransactionTypes 파싱(TransactionType!="") + path / Statistics("AAPL") 파싱(AverageAcquired!=0, CIK!="") + delegation(path+symbol) + 빈 symbol 가드 / AcquisitionOwnership("AAPL",5) 파싱(PercentOfClass!="") + delegation(path+symbol/limit) + 빈 symbol 가드 / SearchReportingName("Cook") 파싱(ReportingName!="") + delegation(path+name) + 빈 name 가드.
- [ ] **Step 4:** `go test ./insidertrades/ && go vet && gofmt -l`. 커밋 `feat(insidertrades): TransactionTypes + Statistics + AcquisitionOwnership + SearchReportingName`.

### Task 3: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/insidertrades/main.go`.

- [ ] **Step 1:** `client.go` — import `insidertrades`, struct 에 `InsiderTrades *insidertrades.Client`, NewClient 에 `c.InsiderTrades = insidertrades.New(hc)`. (gofmt -w client.go 로 정렬.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasInsiderTrades`.
- [ ] **Step 3:** README 표 행 신규: `| Insider Trades | \`client.InsiderTrades\` | LatestInsiderTrades, SearchInsiderTrades, TransactionTypes, Statistics, AcquisitionOwnership, SearchReportingName — 6 endpoint |`.
- [ ] **Step 4:** `examples/insidertrades/main.go` — NewClientFromEnv → SearchInsiderTrades(SearchParams{Symbol:"AAPL",Limit:5}) 건수 + Statistics("AAPL") 첫 행 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_InsiderTrades`: LatestInsiderTrades("",0,5) len>0 / SearchInsiderTrades({Symbol:"AAPL",Limit:5}) / Statistics("AAPL") / TransactionTypes len>0. import insidertrades(SearchParams 사용).
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(insidertrades): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 6 endpoint: InsiderTrade(latest/search)=T1, 나머지 4=T2, 와이어/문서=T3.
- AcquisitionOwnership 전 필드 string. InsiderTrade directOrIndirect/formType 포함.
- 통합테스트 import 에 insidertrades 필요(SearchParams).
