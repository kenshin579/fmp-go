# FMP Go SDK — Calendar 그룹 (v0.8.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** FMP `calendar` 카테고리 9 endpoint 를 신규 `calendar/` 패키지로 추가하고 v0.8.0 준비.

**Architecture:** `internal/fetch` 사용. calendar/ipos 7개 = `(from, to string)` + `dateRange` 헬퍼 + `fetch.List`. company 변형 3개 = `(symbol)` + `fetch.ListBySymbol`. dividend/earning/split struct 는 calendar+company 공유. nullable 포인터(Earning eps/revenue, IPO shares/priceRange/marketCap). 전 struct 필드 한국어 주석.

**Tech Stack:** Go 1.25 / `internal/fetch` / fixture 단위테스트 + build-tag 통합. `unset GOROOT` 필요시.

**Spec:** `docs/superpowers/specs/2026-06-02-calendar-group-design.md`
**Repo / Branch:** `github.com/kenshin579/fmp-go`, branch `feature/calendar-group` (spec 커밋 이미).

**확정된 사실:**
- path: dividends-calendar/earnings-calendar/ipos-calendar/ipos-disclosure/ipos-prospectus/splits-calendar = from/to params, dividends/earnings/splits = symbol.
- 응답 shape 확인. Earning eps/revenue nullable, IPO shares/priceRange/marketCap nullable. dividend/earning/split 의 calendar/company 변형 동일 shape.
- `internal/fetch`: `List[T](ctx, hc, path, params)`, `ListBySymbol[T](ctx, hc, path, symbol)`.
- root 와이어: `Analyst *analyst.Client` 필드 + `c.Analyst = analyst.New(hc)` 패턴.

---

## File Structure
- Create: `calendar/client.go` — Client + New + dateRange 헬퍼.
- Create: `calendar/dividends.go` + `_test.go` + testdata — Dividend.
- Create: `calendar/earnings.go` + `_test.go` + testdata — Earning(nullable).
- Create: `calendar/ipos.go` + `_test.go` + testdata — IPO/IPODisclosure/IPOProspectus.
- Create: `calendar/splits.go` + `_test.go` + testdata — Split.
- Modify: `client.go` / `client_test.go`; `README.md`; Create `examples/calendar/main.go`; Modify `integration_test.go`.

---

## Task 1: 패키지 기반 + dateRange + dividends (TDD)

**Files:** Create `calendar/client.go`, `calendar/dividends.go`, `calendar/dividends_test.go`, `calendar/testdata/dividends-calendar.json`, `calendar/testdata/dividends-company.json`.

- [ ] **Step 1: client.go (Client + New + dateRange)**
```go
// Package calendar 는 FMP 캘린더(calendar) API sub-client.
// fmp.Client.Calendar 로 접근.
package calendar

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 캘린더 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// dateRange 는 from/to(YYYY-MM-DD) 를 쿼리 맵으로 만든다. 빈 값은 제외.
func dateRange(from, to string) map[string]string {
	m := map[string]string{}
	if from != "" {
		m["from"] = from
	}
	if to != "" {
		m["to"] = to
	}
	return m
}
```

- [ ] **Step 2: dividends.go**
```go
package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Dividend — 배당 (dividends-calendar / dividends 공용)
type Dividend struct {
	Symbol          string  `json:"symbol"`          // 종목 심볼
	Date            string  `json:"date"`            // 배당락일
	RecordDate      string  `json:"recordDate"`      // 기준일
	PaymentDate     string  `json:"paymentDate"`     // 지급일
	DeclarationDate string  `json:"declarationDate"` // 선언일
	AdjDividend     float64 `json:"adjDividend"`     // 수정 배당금
	Dividend        float64 `json:"dividend"`        // 배당금
	Yield           float64 `json:"yield"`           // 배당수익률 (%)
	Frequency       string  `json:"frequency"`       // 배당 주기 (예: Quarterly)
}

// DividendsCalendar 는 기간 내 전체 배당 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) DividendsCalendar(ctx context.Context, from, to string) ([]Dividend, error) {
	return fetch.List[Dividend](ctx, c.http, "/stable/dividends-calendar", dateRange(from, to))
}

// CompanyDividends 는 종목의 배당 이력을 조회한다.
func (c *Client) CompanyDividends(ctx context.Context, symbol string) ([]Dividend, error) {
	return fetch.ListBySymbol[Dividend](ctx, c.http, "/stable/dividends", symbol)
}
```

- [ ] **Step 3: fixtures**

`calendar/testdata/dividends-calendar.json`:
```json
[{ "symbol": "1D0.SI", "date": "2025-02-04", "recordDate": "", "paymentDate": "", "declarationDate": "", "adjDividend": 0.01, "dividend": 0.01, "yield": 6.25, "frequency": "Semi-Annual" }]
```
`calendar/testdata/dividends-company.json`:
```json
[{ "symbol": "AAPL", "date": "2025-02-10", "recordDate": "2025-02-10", "paymentDate": "2025-02-13", "declarationDate": "2025-01-30", "adjDividend": 0.25, "dividend": 0.25, "yield": 0.4295, "frequency": "Quarterly" }]
```

- [ ] **Step 4: dividends_test.go (test helper 포함)**
```go
package calendar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func newTestClient(t *testing.T, status int, body string) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, srv.Close
}

type capturedReq struct {
	path  string
	query url.Values
}

func newCapturingClient(t *testing.T, body string) (*Client, *capturedReq, func()) {
	t.Helper()
	cap := &capturedReq{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cap.path = r.URL.Path
		cap.query = r.URL.Query()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, cap, srv.Close
}

func TestDateRange_OmitsEmpty(t *testing.T) {
	m := dateRange("2025-01-01", "")
	if m["from"] != "2025-01-01" {
		t.Errorf("from=%q", m["from"])
	}
	if _, ok := m["to"]; ok {
		t.Error("empty to should be omitted")
	}
	if len(dateRange("", "")) != 0 {
		t.Error("both empty should yield empty map")
	}
}

func TestDividendsCalendar_DelegatesDateRange(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dividends-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.DividendsCalendar(context.Background(), "2025-02-01", "2025-02-28")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Frequency != "Semi-Annual" || rows[0].Yield <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/dividends-calendar" || cap.query.Get("from") != "2025-02-01" || cap.query.Get("to") != "2025-02-28" {
		t.Errorf("delegation: path=%q from=%q to=%q", cap.path, cap.query.Get("from"), cap.query.Get("to"))
	}
}

func TestCompanyDividends_DelegatesSymbol(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dividends-company.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.CompanyDividends(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].PaymentDate == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/dividends" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}
```

- [ ] **Step 5: 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./calendar/ -v && go vet ./calendar/ && gofmt -l calendar/
```
Expected: PASS, vet clean, gofmt 빈 출력.

- [ ] **Step 6: Commit**
```bash
git add calendar/client.go calendar/dividends.go calendar/dividends_test.go calendar/testdata/dividends-calendar.json calendar/testdata/dividends-company.json
git commit -m "feat(calendar): 패키지 기반 + dateRange + Dividends(calendar/company)"
```

---

## Task 2: earnings (nullable eps/revenue) (TDD)

**Files:** Create `calendar/earnings.go`, `calendar/earnings_test.go`, `calendar/testdata/earnings-calendar.json`, `calendar/testdata/earnings-company.json`.

- [ ] **Step 1: earnings.go**
```go
package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Earning — 실적 (earnings-calendar / earnings 공용). 미래 실적은 actual/estimated null.
type Earning struct {
	Symbol           string   `json:"symbol"`           // 종목 심볼
	Date             string   `json:"date"`             // 실적 발표일
	EpsActual        *float64 `json:"epsActual"`        // 실제 EPS(결측 가능)
	EpsEstimated     *float64 `json:"epsEstimated"`     // 추정 EPS(결측 가능)
	RevenueActual    *int64   `json:"revenueActual"`    // 실제 매출(결측 가능)
	RevenueEstimated *int64   `json:"revenueEstimated"` // 추정 매출(결측 가능)
	LastUpdated      string   `json:"lastUpdated"`      // 최종 갱신일
}

// EarningsCalendar 는 기간 내 전체 실적 발표 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) EarningsCalendar(ctx context.Context, from, to string) ([]Earning, error) {
	return fetch.List[Earning](ctx, c.http, "/stable/earnings-calendar", dateRange(from, to))
}

// CompanyEarnings 는 종목의 실적 이력/예정을 조회한다.
func (c *Client) CompanyEarnings(ctx context.Context, symbol string) ([]Earning, error) {
	return fetch.ListBySymbol[Earning](ctx, c.http, "/stable/earnings", symbol)
}
```

- [ ] **Step 2: fixtures**

`calendar/testdata/earnings-calendar.json` (값 있음):
```json
[{ "symbol": "KEC.NS", "date": "2024-11-04", "epsActual": 3.32, "epsEstimated": 4.97, "revenueActual": 51133100000, "revenueEstimated": 44687400000, "lastUpdated": "2024-12-08" }]
```
`calendar/testdata/earnings-company.json` (null 포함):
```json
[
  { "symbol": "AAPL", "date": "2025-10-29", "epsActual": null, "epsEstimated": null, "revenueActual": null, "revenueEstimated": null, "lastUpdated": "2025-02-04" },
  { "symbol": "AAPL", "date": "2025-01-30", "epsActual": 2.4, "epsEstimated": 2.35, "revenueActual": 124300000000, "revenueEstimated": 121000000000, "lastUpdated": "2025-01-31" }
]
```

- [ ] **Step 3: earnings_test.go**
```go
package calendar

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestEarningsCalendar_ParsesValues(t *testing.T) {
	raw, _ := os.ReadFile("testdata/earnings-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.EarningsCalendar(context.Background(), "2024-11-01", "2024-11-30")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].EpsActual == nil || *rows[0].EpsActual != 3.32 {
		t.Errorf("EpsActual: %v", rows[0].EpsActual)
	}
	if rows[0].RevenueActual == nil || *rows[0].RevenueActual <= 0 {
		t.Errorf("RevenueActual: %v", rows[0].RevenueActual)
	}
	if cap.path != "/stable/earnings-calendar" || cap.query.Get("from") != "2024-11-01" {
		t.Errorf("delegation: path=%q from=%q", cap.path, cap.query.Get("from"))
	}
}

func TestCompanyEarnings_NullableNilVsValue(t *testing.T) {
	raw, _ := os.ReadFile("testdata/earnings-company.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CompanyEarnings(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	// row0: 미래 실적 → 전부 nil
	if rows[0].EpsActual != nil || rows[0].RevenueActual != nil {
		t.Errorf("row0 nullables should be nil: %+v", rows[0])
	}
	// row1: 값 존재
	if rows[1].EpsActual == nil || *rows[1].EpsActual != 2.4 {
		t.Errorf("row1 EpsActual should be set")
	}
	if rows[1].RevenueActual == nil || *rows[1].RevenueActual <= 0 {
		t.Errorf("row1 RevenueActual should be set")
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./calendar/ && go vet ./calendar/ && gofmt -l calendar/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add calendar/earnings.go calendar/earnings_test.go calendar/testdata/earnings-calendar.json calendar/testdata/earnings-company.json
git commit -m "feat(calendar): Earnings(calendar/company, nullable eps/revenue)"
```

---

## Task 3: ipos (calendar + disclosure + prospectus) (TDD)

**Files:** Create `calendar/ipos.go`, `calendar/ipos_test.go`, `calendar/testdata/ipos-calendar.json`, `calendar/testdata/ipos-disclosure.json`, `calendar/testdata/ipos-prospectus.json`.

- [ ] **Step 1: ipos.go**
```go
package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// IPO — IPO 일정 (ipos-calendar). shares/priceRange/marketCap 결측 가능.
type IPO struct {
	Symbol     string  `json:"symbol"`     // 종목 심볼
	Date       string  `json:"date"`       // IPO 일자
	Daa        string  `json:"daa"`        // 공시 일시 (ISO8601)
	Company    string  `json:"company"`    // 회사명
	Exchange   string  `json:"exchange"`   // 거래소
	Actions    string  `json:"actions"`    // 상태 (예: Expected)
	Shares     *int64  `json:"shares"`     // 공모 주식 수(결측 가능)
	PriceRange *string `json:"priceRange"` // 공모가 범위(결측 가능)
	MarketCap  *int64  `json:"marketCap"`  // 시가총액(결측 가능)
}

// IPODisclosure — IPO 공시 서류 (ipos-disclosure)
type IPODisclosure struct {
	Symbol            string `json:"symbol"`            // 종목 심볼
	FilingDate        string `json:"filingDate"`        // 제출일
	AcceptedDate      string `json:"acceptedDate"`      // 수리일
	EffectivenessDate string `json:"effectivenessDate"` // 효력 발생일
	CIK               string `json:"cik"`               // SEC CIK
	Form              string `json:"form"`              // 공시 양식 (예: CERT)
	URL               string `json:"url"`               // 원문 URL
}

// IPOProspectus — IPO 투자설명서 (ipos-prospectus)
type IPOProspectus struct {
	Symbol                          string  `json:"symbol"`                          // 종목 심볼
	AcceptedDate                    string  `json:"acceptedDate"`                    // 수리일
	FilingDate                      string  `json:"filingDate"`                      // 제출일
	IPODate                         string  `json:"ipoDate"`                         // IPO 일자
	CIK                             string  `json:"cik"`                             // SEC CIK
	PricePublicPerShare             float64 `json:"pricePublicPerShare"`             // 주당 공모가
	PricePublicTotal                float64 `json:"pricePublicTotal"`                // 총 공모금액
	DiscountsAndCommissionsPerShare float64 `json:"discountsAndCommissionsPerShare"` // 주당 인수수수료
	DiscountsAndCommissionsTotal    float64 `json:"discountsAndCommissionsTotal"`    // 총 인수수수료
	ProceedsBeforeExpensesPerShare  float64 `json:"proceedsBeforeExpensesPerShare"`  // 주당 순수취금(비용 전)
	ProceedsBeforeExpensesTotal     float64 `json:"proceedsBeforeExpensesTotal"`     // 총 순수취금(비용 전)
	Form                            string  `json:"form"`                            // 공시 양식 (예: 424B4)
	URL                             string  `json:"url"`                             // 원문 URL
}

// IPOsCalendar 는 기간 내 IPO 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) IPOsCalendar(ctx context.Context, from, to string) ([]IPO, error) {
	return fetch.List[IPO](ctx, c.http, "/stable/ipos-calendar", dateRange(from, to))
}

// IPODisclosures 는 기간 내 IPO 공시 서류를 조회한다.
func (c *Client) IPODisclosures(ctx context.Context, from, to string) ([]IPODisclosure, error) {
	return fetch.List[IPODisclosure](ctx, c.http, "/stable/ipos-disclosure", dateRange(from, to))
}

// IPOProspectuses 는 기간 내 IPO 투자설명서를 조회한다.
func (c *Client) IPOProspectuses(ctx context.Context, from, to string) ([]IPOProspectus, error) {
	return fetch.List[IPOProspectus](ctx, c.http, "/stable/ipos-prospectus", dateRange(from, to))
}
```

- [ ] **Step 2: fixtures**

`calendar/testdata/ipos-calendar.json` (null 포함):
```json
[{ "symbol": "PEVC", "date": "2025-02-03", "daa": "2025-02-03T05:00:00.000Z", "company": "Pacer Funds Trust", "exchange": "NYSE", "actions": "Expected", "shares": null, "priceRange": null, "marketCap": null }]
```
`calendar/testdata/ipos-disclosure.json`:
```json
[{ "symbol": "SCHM", "filingDate": "2025-02-03", "acceptedDate": "2025-02-03", "effectivenessDate": "2025-02-03", "cik": "0001454889", "form": "CERT", "url": "https://www.sec.gov/Archives/edgar/data/1454889/x.pdf" }]
```
`calendar/testdata/ipos-prospectus.json`:
```json
[{ "symbol": "ATAK", "acceptedDate": "2025-02-03", "filingDate": "2025-02-03", "ipoDate": "2022-03-20", "cik": "0001883788", "pricePublicPerShare": 0.78, "pricePublicTotal": 4649936.72, "discountsAndCommissionsPerShare": 0.04, "discountsAndCommissionsTotal": 254909.67, "proceedsBeforeExpensesPerShare": 0.74, "proceedsBeforeExpensesTotal": 4395207.05, "form": "424B4", "url": "https://www.sec.gov/Archives/edgar/data/1883788/x.htm" }]
```

- [ ] **Step 3: ipos_test.go**
```go
package calendar

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestIPOsCalendar_NullableFields(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ipos-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.IPOsCalendar(context.Background(), "2025-02-01", "2025-02-28")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Company == "" || rows[0].Exchange != "NYSE" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	// null → nil
	if rows[0].Shares != nil || rows[0].PriceRange != nil || rows[0].MarketCap != nil {
		t.Errorf("nullables should be nil: %+v", rows[0])
	}
	if cap.path != "/stable/ipos-calendar" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestIPODisclosures_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ipos-disclosure.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.IPODisclosures(context.Background(), "", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Form != "CERT" || rows[0].URL == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestIPOProspectuses_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ipos-prospectus.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.IPOProspectuses(context.Background(), "", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].PricePublicPerShare <= 0 || rows[0].ProceedsBeforeExpensesTotal <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./calendar/ && go vet ./calendar/ && gofmt -l calendar/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add calendar/ipos.go calendar/ipos_test.go calendar/testdata/ipos-calendar.json calendar/testdata/ipos-disclosure.json calendar/testdata/ipos-prospectus.json
git commit -m "feat(calendar): IPOsCalendar + IPODisclosures + IPOProspectuses"
```

---

## Task 4: splits (TDD)

**Files:** Create `calendar/splits.go`, `calendar/splits_test.go`, `calendar/testdata/splits-calendar.json`, `calendar/testdata/splits-company.json`.

- [ ] **Step 1: splits.go**
```go
package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Split — 주식 분할 (splits-calendar / splits 공용)
type Split struct {
	Symbol      string `json:"symbol"`      // 종목 심볼
	Date        string `json:"date"`        // 분할 기준일
	Numerator   int    `json:"numerator"`   // 분할 비율 분자
	Denominator int    `json:"denominator"` // 분할 비율 분모
}

// SplitsCalendar 는 기간 내 주식 분할 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) SplitsCalendar(ctx context.Context, from, to string) ([]Split, error) {
	return fetch.List[Split](ctx, c.http, "/stable/splits-calendar", dateRange(from, to))
}

// CompanySplits 는 종목의 주식 분할 이력을 조회한다.
func (c *Client) CompanySplits(ctx context.Context, symbol string) ([]Split, error) {
	return fetch.ListBySymbol[Split](ctx, c.http, "/stable/splits", symbol)
}
```

- [ ] **Step 2: fixtures**

`calendar/testdata/splits-calendar.json`:
```json
[{ "symbol": "EYEN", "date": "2025-02-03", "numerator": 1, "denominator": 80 }]
```
`calendar/testdata/splits-company.json`:
```json
[{ "symbol": "AAPL", "date": "2020-08-31", "numerator": 4, "denominator": 1 }]
```

- [ ] **Step 3: splits_test.go**
```go
package calendar

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSplitsCalendar_DelegatesDateRange(t *testing.T) {
	raw, _ := os.ReadFile("testdata/splits-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SplitsCalendar(context.Background(), "2025-02-01", "2025-02-28")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Numerator != 1 || rows[0].Denominator != 80 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/splits-calendar" || cap.query.Get("from") != "2025-02-01" {
		t.Errorf("delegation: path=%q from=%q", cap.path, cap.query.Get("from"))
	}
}

func TestCompanySplits_DelegatesSymbol(t *testing.T) {
	raw, _ := os.ReadFile("testdata/splits-company.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.CompanySplits(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Numerator != 4 || rows[0].Denominator != 1 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/splits" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./calendar/ && go vet ./calendar/ && gofmt -l calendar/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add calendar/splits.go calendar/splits_test.go calendar/testdata/splits-calendar.json calendar/testdata/splits-company.json
git commit -m "feat(calendar): Splits(calendar/company)"
```

---

## Task 5: 루트 와이어 + README + examples + 통합 + 검증 (TDD)

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/calendar/main.go`.

- [ ] **Step 1: client_test.go 실패 테스트**
```go
func TestNewClient_HasCalendar(t *testing.T) {
	c, err := NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Calendar == nil {
		t.Fatal("Calendar sub-client is nil")
	}
}
```

- [ ] **Step 2: client.go 와이어** — import `"github.com/kenshin579/fmp-go/calendar"`; `Client` 에 `Calendar *calendar.Client // 캘린더(배당/실적/IPO/분할)` 필드; `NewClient` 에 `c.Calendar = calendar.New(hc)`.

- [ ] **Step 3: 통과** — `go build ./... && go test . -run TestNewClient_HasCalendar`. PASS.

- [ ] **Step 4: README + examples**

README 커버리지 표 행 추가:
```markdown
| Calendar | `client.Calendar` | DividendsCalendar, CompanyDividends, EarningsCalendar, CompanyEarnings, IPOsCalendar, IPODisclosures, IPOProspectuses, SplitsCalendar, CompanySplits — 9 endpoint |
```

Create `examples/calendar/main.go`:
```go
// 실행: FMP_API_KEY=... go run examples/calendar/main.go
package main

import (
	"context"
	"fmt"
	"log"

	fmp "github.com/kenshin579/fmp-go"
)

func main() {
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	earns, err := c.Calendar.EarningsCalendar(ctx, "2025-02-01", "2025-02-07")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("이번 주 실적 발표: %d건\n", len(earns))

	divs, err := c.Calendar.CompanyDividends(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	if len(divs) > 0 {
		fmt.Printf("AAPL 최근 배당: %.2f (%s)\n", divs[0].Dividend, divs[0].Date)
	}
}
```

- [ ] **Step 5: integration_test.go 에 calendar 케이스**
```go
func TestIntegration_Calendar(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Calendar.EarningsCalendar(ctx, "2025-02-01", "2025-02-07"); err != nil {
		t.Errorf("EarningsCalendar: %v", err)
	} else {
		t.Logf("EarningsCalendar: %d건", len(rows))
	}
	if rows, err := c.Calendar.CompanyDividends(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("CompanyDividends: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Calendar.IPOsCalendar(ctx, "2025-02-01", "2025-02-28"); err != nil {
		t.Errorf("IPOsCalendar: %v", err)
	} else if len(rows) > 0 {
		t.Logf("IPO[0]: %+v", rows[0]) // priceRange 실 타입 확인
	}
}
```

- [ ] **Step 6: 전체 검증**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./...
go test -tags integration ./... -run TestIntegration_Calendar -v 2>&1 | tail -12
gofmt -l .
```
Expected: 단위 전체 PASS, gofmt clean. 통합 key 없으면 skip.

- [ ] **Step 7: Commit**
```bash
git add client.go client_test.go README.md examples/calendar/main.go integration_test.go
git commit -m "feat(calendar): 루트 Client 와이어 + README + examples + 통합"
```

---

## 자기 점검 메모 (작성자용)
- struct 공유: Dividend/Earning/Split 의 calendar+company 변형.
- nullable 포인터: Earning(eps/revenue), IPO(shares/priceRange/marketCap). null→nil fixture 검증.
- calendar/ipos = `(from, to)` + dateRange(빈값 제외). company = `(symbol)` + ListBySymbol.
- IPO.PriceRange *string — 카탈로그 예시 null 뿐, 통합 로그로 실 타입 확인(다르면 조정).
- test helper(newTestClient/newCapturingClient/capturedReq + dateRange 테스트)는 Task 1 dividends_test.go 정의 → Task 2-4 재사용(동일 패키지).
- 와이어: `Calendar *calendar.Client` + `c.Calendar = calendar.New(hc)`.
