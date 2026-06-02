# Reports 그룹 (v0.11.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `reports/` 패키지 7 endpoint(as-reported 3 + full 1 + latest 1 + dates 1 + 10-K JSON 1).

**Architecture:** internal/fetch(List/ListBySymbol) + 패키지 helper `asReportedParams`. as-reported 3종은 공유 `AsReportedStatement`(Data map[string]json.Number), full 은 `AsReportedFull`(Data map[string]any), 10-K JSON 은 `[]map[string]any` 원시. 구조체는 스펙 verbatim.

**Tech Stack:** Go 1.25 generics, encoding/json(json.Number), internal/fetch.

참고: `unset GOROOT`. 커밋 한국어 `feat(reports): ...`. 테스트 헬퍼는 calendar/metrics 패턴을 패키지 내부에 정의.

---

### Task 1: reports 패키지 scaffold + as-reported 4종

**Files:** Create `reports/client.go`, `reports/as_reported.go`, `reports/as_reported_test.go`, testdata `income-statement-as-reported.json`, `balance-sheet-statement-as-reported.json`, `cash-flow-statement-as-reported.json`, `financial-statement-full-as-reported.json`.

- [ ] **Step 1:** `reports/client.go`:
```go
// Package reports 는 FMP 보고서 API sub-client (as-reported/latest/dates/10-K).
// fmp.Client.Reports 로 접근.
package reports

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 보고서 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// asReportedParams 는 symbol(필수) + period(비어있지 않으면) + limit(>0) 쿼리 맵.
func asReportedParams(symbol, period string, limit int) map[string]string {
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
- [ ] **Step 2:** `reports/as_reported.go` — 스펙의 `AsReportedStatement`(Data map[string]json.Number) + `AsReportedFull`(Data map[string]any) struct + 4 메서드. import: context, fmt, strings, encoding/json, internal/fetch.
```go
func (c *Client) IncomeStatementAsReported(ctx context.Context, symbol, period string, limit int) ([]AsReportedStatement, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[AsReportedStatement](ctx, c.http, "/stable/income-statement-as-reported", asReportedParams(symbol, period, limit))
}
// BalanceSheetStatementAsReported → /stable/balance-sheet-statement-as-reported (동일 패턴, []AsReportedStatement)
// CashFlowStatementAsReported → /stable/cash-flow-statement-as-reported (동일)
// FinancialStatementFullAsReported → /stable/financial-statement-full-as-reported ([]AsReportedFull)
```
4개 모두 빈 symbol 가드 + fetch.List + asReportedParams.
- [ ] **Step 3:** fixtures(1건 배열, AAPL). income/balance/cashflow-as-reported: `reportedCurrency: null`, `fiscalYear: 2024`, `data` 에 대표 키 몇 개(income 은 `grossprofit`, `netincomeloss` 정수 + `earningspersharediluted` 소수). full: data 에 숫자 키(`netincomeloss`) + 문자열 키(`documenttype`:"10-K") + 불리언/문자열 혼합 1개.
- [ ] **Step 4:** `reports/as_reported_test.go` — 테스트 헬퍼(newTestClient/newCapturingClient/capturedReq) 를 calendar/dividends_test.go 패턴대로 reports 패키지에 정의. 테스트:
  - IncomeStatementAsReported 파싱: len==1, ReportedCurrency==nil, FiscalYear==2024, Data["grossprofit"].Int64() 성공 & >0, Data["earningspersharediluted"].Float64() 성공.
  - FinancialStatementFullAsReported 파싱: Data["documenttype"]=="10-K"(string), Data["netincomeloss"] 숫자 존재.
  - delegation(capturing): IncomeStatementAsReported(ctx,"AAPL","annual",1) → path "/stable/income-statement-as-reported", query symbol/period/limit.
  - 빈 symbol 가드 1건.
- [ ] **Step 5:** `unset GOROOT && go test ./reports/ && go vet ./reports/ && gofmt -l reports/`. 커밋 `feat(reports): as-reported 재무제표 4종`.

### Task 2: latest + dates + 10-K JSON

**Files:** Create `reports/reports.go`, `reports/reports_test.go`, testdata `latest-financial-statements.json`, `financial-reports-dates.json`, `financial-reports-json.json`.

- [ ] **Step 1:** `reports/reports.go` — 스펙 `LatestFinancialStatement` + `FinancialReportDate` struct + 3 메서드:
```go
package reports

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

func (c *Client) LatestFinancialStatements(ctx context.Context, page, limit int) ([]LatestFinancialStatement, error) {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[LatestFinancialStatement](ctx, c.http, "/stable/latest-financial-statements", q)
}

func (c *Client) FinancialReportDates(ctx context.Context, symbol string) ([]FinancialReportDate, error) {
	return fetch.ListBySymbol[FinancialReportDate](ctx, c.http, "/stable/financial-reports-dates", symbol)
}

func (c *Client) FinancialReportJSON(ctx context.Context, symbol string, year int, period string) ([]map[string]any, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := map[string]string{"symbol": symbol, "year": strconv.Itoa(year), "period": period}
	return fetch.List[map[string]any](ctx, c.http, "/stable/financial-reports-json", q)
}
```
- [ ] **Step 2:** fixtures. latest: 1건 `{symbol, calendarYear:2024, period:"Q4", date, dateAdded:"2025-03-13 17:03:59"}`. dates: 1건 `{symbol:"AAPL", fiscalYear:2022, period:"FY", linkXlsx:"https://...", linkJson:"https://..."}`. financial-reports-json: 1건 배열 `[{"symbol":"AAPL","period":"FY","year":"2022","Cover Page":[...간단...]}]`.
- [ ] **Step 3:** `reports/reports_test.go`(헬퍼 재사용):
  - LatestFinancialStatements 파싱: CalendarYear==2024, DateAdded!="" ; delegation path+page/limit.
  - FinancialReportDates 파싱: LinkXlsx!="", LinkJson!="" ; path.
  - FinancialReportJSON 파싱: len==1, m["symbol"]=="AAPL", m["year"]=="2022"(string), "Cover Page" 키 존재 ; delegation path+symbol/year/period. 빈 symbol 가드.
- [ ] **Step 4:** `go test ./reports/ && go vet && gofmt -l`. 커밋 `feat(reports): LatestFinancialStatements + FinancialReportDates + FinancialReportJSON`.

### Task 3: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/reports/main.go`.

- [ ] **Step 1:** `client.go` — import `reports`, struct 에 `Reports *reports.Client`, NewClient 에 `c.Reports = reports.New(hc)`.
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasReports`.
- [ ] **Step 3:** README 표 Reports 행 신규: `| Reports | \`client.Reports\` | IncomeStatementAsReported, BalanceSheetStatementAsReported, CashFlowStatementAsReported, FinancialStatementFullAsReported, LatestFinancialStatements, FinancialReportDates, FinancialReportJSON — 7 endpoint |`.
- [ ] **Step 4:** `examples/reports/main.go` — NewClientFromEnv → IncomeStatementAsReported(AAPL,annual,1) data 키 개수 출력 + FinancialReportDates(AAPL) 첫 링크 출력. (examples/metrics/main.go 패턴 참고.)
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Reports`: IncomeStatementAsReported(AAPL,annual,1) len(Data)>0 & Data["grossprofit"] 존재 / LatestFinancialStatements(0,5) len>0 / FinancialReportDates(AAPL) len>0 & LinkJson!="" / FinancialReportJSON(AAPL,2022,"FY") len>0 → t.Logf 첫 원소 키 일부. import 불필요(positional).
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(reports): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 7 endpoint 매핑: as-reported 4=T1, latest/dates/10-K=T2, 와이어/문서=T3.
- AsReportedStatement(json.Number) vs AsReportedFull(any) 구분 — 문자열 혼재 때문.
- LatestFinancialStatements 는 symbol 없음(page/limit), calendarYear 사용.
- FinancialReportJSON 은 []map[string]any 원시, year 쿼리는 문자열로 보내되 응답 year 도 문자열.
- xlsx 제외 — 범위 밖.
