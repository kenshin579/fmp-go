# FMP Go SDK — Analyst 그룹 (v0.7.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** FMP `analyst` 카테고리 8 endpoint 를 신규 `analyst/` 패키지로 추가하고 v0.7.0 준비.

**Architecture:** `internal/fetch` 사용. 단일 *T 4개(OneBySymbol) + list []T 3개(ListBySymbol) + financial-estimates(List+params). Rating 은 snapshot/historical 공유. FinancialEstimate 는 카탈로그 예시 없어 FMP 공개 shape 합성(통합 검증). 모든 struct 필드 한국어 주석.

**Tech Stack:** Go 1.25 / `internal/fetch` / fixture 단위테스트 + build-tag 통합. `unset GOROOT` 필요시.

**Spec:** `docs/superpowers/specs/2026-06-02-analyst-group-design.md`
**Repo / Branch:** `github.com/kenshin579/fmp-go`, branch `feature/analyst-group` (spec 커밋 이미).

**확정된 사실:**
- path: `/stable/grades?symbol=`, `/stable/grades-consensus?symbol=`, `/stable/grades-historical?symbol=`, `/stable/ratings-snapshot?symbol=`, `/stable/ratings-historical?symbol=`, `/stable/price-target-consensus?symbol=`, `/stable/price-target-summary?symbol=`, `/stable/analyst-estimates?symbol=&period=&page=`.
- 응답 shape 확인(grades/grades-consensus/grades-historical/ratings-snapshot/ratings-historical/price-target-consensus/price-target-summary). analyst-estimates 만 예시 없음 → 합성.
- HistoricalGrade 에 StrongBuy 없음(Buy/Hold/Sell/StrongSell). Rating snapshot=date 없음/historical=date 있음. PriceTargetSummary.Publishers 는 JSON 배열 문자열.
- `internal/fetch`: `OneBySymbol[T]`/`ListBySymbol[T]`/`List[T]`.
- root 와이어: `News *news.Client` 필드 + `c.News = news.New(hc)` 패턴.

---

## File Structure
- Create: `analyst/client.go` — Client + New.
- Create: `analyst/grades.go` + `_test.go` + testdata — Grade/GradesConsensus/HistoricalGrade.
- Create: `analyst/ratings.go` + `_test.go` + testdata — Rating(공유).
- Create: `analyst/price_target.go` + `_test.go` + testdata — PriceTargetConsensus/Summary.
- Create: `analyst/estimates.go` + `_test.go` + testdata — FinancialEstimate.
- Modify: `client.go` / `client_test.go` — Analyst 와이어.
- Modify: `README.md`; Create `examples/analyst/main.go`; Modify `integration_test.go`.

---

## Task 1: analyst 패키지 기반 + grades 3개 (TDD)

**Files:** Create `analyst/client.go`, `analyst/grades.go`, `analyst/grades_test.go`, `analyst/testdata/grades.json`, `analyst/testdata/grades-consensus.json`, `analyst/testdata/grades-historical.json`.

- [ ] **Step 1: client.go**
```go
// Package analyst 는 FMP 애널리스트(analyst) API sub-client.
// fmp.Client.Analyst 로 접근.
package analyst

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 애널리스트 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
```

- [ ] **Step 2: grades.go**
```go
package analyst

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Grade — 개별 애널리스트 등급 변경 (grades)
type Grade struct {
	Symbol         string `json:"symbol"`         // 종목 심볼
	Date           string `json:"date"`           // 등급 변경일
	GradingCompany string `json:"gradingCompany"` // 평가 기관
	PreviousGrade  string `json:"previousGrade"`  // 이전 등급
	NewGrade       string `json:"newGrade"`       // 신규 등급
	Action         string `json:"action"`         // 조치 (maintain/upgrade/downgrade)
}

// GradesConsensus — 등급 컨센서스 집계 (grades-consensus)
type GradesConsensus struct {
	Symbol     string `json:"symbol"`     // 종목 심볼
	StrongBuy  int    `json:"strongBuy"`  // 적극 매수 수
	Buy        int    `json:"buy"`        // 매수
	Hold       int    `json:"hold"`       // 보유
	Sell       int    `json:"sell"`       // 매도
	StrongSell int    `json:"strongSell"` // 적극 매도
	Consensus  string `json:"consensus"`  // 종합 컨센서스
}

// HistoricalGrade — 일자별 등급 분포 (grades-historical). FMP 응답에 StrongBuy 없음.
type HistoricalGrade struct {
	Symbol                   string `json:"symbol"`                   // 종목 심볼
	Date                     string `json:"date"`                     // 기준일
	AnalystRatingsBuy        int    `json:"analystRatingsBuy"`        // 매수 의견 수
	AnalystRatingsHold       int    `json:"analystRatingsHold"`       // 보유
	AnalystRatingsSell       int    `json:"analystRatingsSell"`       // 매도
	AnalystRatingsStrongSell int    `json:"analystRatingsStrongSell"` // 적극 매도
}

// Grades 는 종목의 개별 애널리스트 등급 변경 이력을 조회한다.
func (c *Client) Grades(ctx context.Context, symbol string) ([]Grade, error) {
	return fetch.ListBySymbol[Grade](ctx, c.http, "/stable/grades", symbol)
}

// GradesConsensus 는 종목의 등급 컨센서스 집계를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) GradesConsensus(ctx context.Context, symbol string) (*GradesConsensus, error) {
	return fetch.OneBySymbol[GradesConsensus](ctx, c.http, "/stable/grades-consensus", symbol)
}

// HistoricalGrades 는 종목의 일자별 등급 분포를 조회한다.
func (c *Client) HistoricalGrades(ctx context.Context, symbol string) ([]HistoricalGrade, error) {
	return fetch.ListBySymbol[HistoricalGrade](ctx, c.http, "/stable/grades-historical", symbol)
}
```

- [ ] **Step 3: fixtures**

`analyst/testdata/grades.json`:
```json
[{ "symbol": "AAPL", "date": "2025-01-31", "gradingCompany": "Morgan Stanley", "previousGrade": "Overweight", "newGrade": "Overweight", "action": "maintain" }]
```
`analyst/testdata/grades-consensus.json`:
```json
[{ "symbol": "AAPL", "strongBuy": 1, "buy": 29, "hold": 11, "sell": 4, "strongSell": 0, "consensus": "Buy" }]
```
`analyst/testdata/grades-historical.json`:
```json
[
  { "symbol": "AAPL", "date": "2025-02-01", "analystRatingsBuy": 8, "analystRatingsHold": 14, "analystRatingsSell": 2, "analystRatingsStrongSell": 2 },
  { "symbol": "AAPL", "date": "2025-01-01", "analystRatingsBuy": 7, "analystRatingsHold": 15, "analystRatingsSell": 3, "analystRatingsStrongSell": 1 }
]
```

- [ ] **Step 4: grades_test.go (test helper 포함)**
```go
package analyst

import (
	"context"
	"errors"
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

func TestGrades_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/grades.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.Grades(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].GradingCompany == "" || rows[0].Action != "maintain" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/grades" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}

func TestGradesConsensus_ParsesSingle(t *testing.T) {
	raw, _ := os.ReadFile("testdata/grades-consensus.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	g, err := c.GradesConsensus(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("GradesConsensus: %v", err)
	}
	if g.Buy != 29 || g.Consensus != "Buy" {
		t.Errorf("not parsed: %+v", g)
	}
}

func TestGradesConsensus_EmptyNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.GradesConsensus(context.Background(), "NOPE"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestHistoricalGrades_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/grades-historical.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HistoricalGrades(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].AnalystRatingsBuy != 8 || rows[0].Date == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
```

- [ ] **Step 5: 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./analyst/ -v && go vet ./analyst/ && gofmt -l analyst/
```
Expected: PASS, vet clean, gofmt 빈 출력.

- [ ] **Step 6: Commit**
```bash
git add analyst/client.go analyst/grades.go analyst/grades_test.go analyst/testdata/grades.json analyst/testdata/grades-consensus.json analyst/testdata/grades-historical.json
git commit -m "feat(analyst): 패키지 기반 + Grades/GradesConsensus/HistoricalGrades"
```

---

## Task 2: ratings (snapshot + historical, Rating 공유) (TDD)

**Files:** Create `analyst/ratings.go`, `analyst/ratings_test.go`, `analyst/testdata/ratings-snapshot.json`, `analyst/testdata/ratings-historical.json`.

- [ ] **Step 1: ratings.go**
```go
package analyst

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Rating — 종합 평가 점수 (ratings-snapshot / ratings-historical 공용). snapshot 은 Date "".
type Rating struct {
	Symbol                  string `json:"symbol"`                  // 종목 심볼
	Date                    string `json:"date"`                    // 기준일(snapshot 은 빈 문자열)
	Rating                  string `json:"rating"`                  // 등급 (예: A-)
	OverallScore            int    `json:"overallScore"`            // 종합 점수
	DiscountedCashFlowScore int    `json:"discountedCashFlowScore"` // DCF 점수
	ReturnOnEquityScore     int    `json:"returnOnEquityScore"`     // ROE 점수
	ReturnOnAssetsScore     int    `json:"returnOnAssetsScore"`     // ROA 점수
	DebtToEquityScore       int    `json:"debtToEquityScore"`       // 부채비율 점수
	PriceToEarningsScore    int    `json:"priceToEarningsScore"`    // PER 점수
	PriceToBookScore        int    `json:"priceToBookScore"`        // PBR 점수
}

// RatingsSnapshot 은 종목의 현재 종합 평가 점수를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) RatingsSnapshot(ctx context.Context, symbol string) (*Rating, error) {
	return fetch.OneBySymbol[Rating](ctx, c.http, "/stable/ratings-snapshot", symbol)
}

// HistoricalRatings 는 종목의 평가 점수 시계열을 조회한다.
func (c *Client) HistoricalRatings(ctx context.Context, symbol string) ([]Rating, error) {
	return fetch.ListBySymbol[Rating](ctx, c.http, "/stable/ratings-historical", symbol)
}
```

- [ ] **Step 2: fixtures**

`analyst/testdata/ratings-snapshot.json` (date 없음):
```json
[{ "symbol": "AAPL", "rating": "A-", "overallScore": 4, "discountedCashFlowScore": 3, "returnOnEquityScore": 5, "returnOnAssetsScore": 5, "debtToEquityScore": 4, "priceToEarningsScore": 2, "priceToBookScore": 1 }]
```
`analyst/testdata/ratings-historical.json` (date 있음):
```json
[
  { "symbol": "AAPL", "date": "2025-02-04", "rating": "A-", "overallScore": 4, "discountedCashFlowScore": 3, "returnOnEquityScore": 5, "returnOnAssetsScore": 5, "debtToEquityScore": 4, "priceToEarningsScore": 2, "priceToBookScore": 1 },
  { "symbol": "AAPL", "date": "2025-02-03", "rating": "B+", "overallScore": 4, "discountedCashFlowScore": 3, "returnOnEquityScore": 4, "returnOnAssetsScore": 5, "debtToEquityScore": 4, "priceToEarningsScore": 2, "priceToBookScore": 1 }
]
```

- [ ] **Step 3: ratings_test.go**
```go
package analyst

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestRatingsSnapshot_ParsesNoDateField(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ratings-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	r, err := c.RatingsSnapshot(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("RatingsSnapshot: %v", err)
	}
	if r.Rating != "A-" || r.OverallScore != 4 {
		t.Errorf("not parsed: %+v", r)
	}
	if r.Date != "" {
		t.Errorf("snapshot Date should be empty, got %q", r.Date)
	}
}

func TestHistoricalRatings_ParsesWithDate(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ratings-historical.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.HistoricalRatings(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Date == "" || rows[0].PriceToBookScore != 1 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/ratings-historical" {
		t.Errorf("path=%q", cap.path)
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./analyst/ && go vet ./analyst/ && gofmt -l analyst/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add analyst/ratings.go analyst/ratings_test.go analyst/testdata/ratings-snapshot.json analyst/testdata/ratings-historical.json
git commit -m "feat(analyst): RatingsSnapshot + HistoricalRatings (Rating 공유)"
```

---

## Task 3: price target (consensus + summary) (TDD)

**Files:** Create `analyst/price_target.go`, `analyst/price_target_test.go`, `analyst/testdata/price-target-consensus.json`, `analyst/testdata/price-target-summary.json`.

- [ ] **Step 1: price_target.go**
```go
package analyst

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// PriceTargetConsensus — 목표주가 컨센서스 (price-target-consensus)
type PriceTargetConsensus struct {
	Symbol          string  `json:"symbol"`          // 종목 심볼
	TargetHigh      float64 `json:"targetHigh"`      // 최고 목표가
	TargetLow       float64 `json:"targetLow"`       // 최저 목표가
	TargetConsensus float64 `json:"targetConsensus"` // 평균 목표가
	TargetMedian    float64 `json:"targetMedian"`    // 중앙값 목표가
}

// PriceTargetSummary — 목표주가 요약 (price-target-summary)
type PriceTargetSummary struct {
	Symbol                    string  `json:"symbol"`                    // 종목 심볼
	LastMonthCount            int     `json:"lastMonthCount"`            // 최근 1개월 리포트 수
	LastMonthAvgPriceTarget   float64 `json:"lastMonthAvgPriceTarget"`   // 최근 1개월 평균 목표가
	LastQuarterCount          int     `json:"lastQuarterCount"`          // 최근 분기 리포트 수
	LastQuarterAvgPriceTarget float64 `json:"lastQuarterAvgPriceTarget"` // 최근 분기 평균 목표가
	LastYearCount             int     `json:"lastYearCount"`             // 최근 1년 리포트 수
	LastYearAvgPriceTarget    float64 `json:"lastYearAvgPriceTarget"`    // 최근 1년 평균 목표가
	AllTimeCount              int     `json:"allTimeCount"`              // 전체 리포트 수
	AllTimeAvgPriceTarget     float64 `json:"allTimeAvgPriceTarget"`     // 전체 평균 목표가
	Publishers                string  `json:"publishers"`                // 발행처 목록(JSON 배열 문자열)
}

// PriceTargetConsensus 는 종목의 목표주가 컨센서스를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) PriceTargetConsensus(ctx context.Context, symbol string) (*PriceTargetConsensus, error) {
	return fetch.OneBySymbol[PriceTargetConsensus](ctx, c.http, "/stable/price-target-consensus", symbol)
}

// PriceTargetSummary 는 종목의 목표주가 요약(기간별 평균)을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) PriceTargetSummary(ctx context.Context, symbol string) (*PriceTargetSummary, error) {
	return fetch.OneBySymbol[PriceTargetSummary](ctx, c.http, "/stable/price-target-summary", symbol)
}
```

- [ ] **Step 2: fixtures**

`analyst/testdata/price-target-consensus.json`:
```json
[{ "symbol": "AAPL", "targetHigh": 300, "targetLow": 200, "targetConsensus": 251.7, "targetMedian": 258 }]
```
`analyst/testdata/price-target-summary.json`:
```json
[{ "symbol": "AAPL", "lastMonthCount": 1, "lastMonthAvgPriceTarget": 200.75, "lastQuarterCount": 3, "lastQuarterAvgPriceTarget": 204.2, "lastYearCount": 48, "lastYearAvgPriceTarget": 232.99, "allTimeCount": 167, "allTimeAvgPriceTarget": 201.21, "publishers": "[\"Benzinga\",\"TheFly\",\"Barrons\"]" }]
```

- [ ] **Step 3: price_target_test.go**
```go
package analyst

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestPriceTargetConsensus_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/price-target-consensus.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	pt, err := c.PriceTargetConsensus(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("PriceTargetConsensus: %v", err)
	}
	if pt.TargetHigh <= 0 || pt.TargetConsensus <= 0 || pt.TargetMedian <= 0 {
		t.Errorf("not parsed: %+v", pt)
	}
}

func TestPriceTargetSummary_ParsesPublishersString(t *testing.T) {
	raw, _ := os.ReadFile("testdata/price-target-summary.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	s, err := c.PriceTargetSummary(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("PriceTargetSummary: %v", err)
	}
	if s.AllTimeCount != 167 || s.LastYearAvgPriceTarget <= 0 {
		t.Errorf("not parsed: %+v", s)
	}
	// publishers 는 JSON 배열이 문자열로 인코딩됨 — 문자열 그대로 보존
	if !strings.HasPrefix(s.Publishers, "[") || !strings.Contains(s.Publishers, "Benzinga") {
		t.Errorf("publishers should be JSON-array string: %q", s.Publishers)
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./analyst/ && go vet ./analyst/ && gofmt -l analyst/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add analyst/price_target.go analyst/price_target_test.go analyst/testdata/price-target-consensus.json analyst/testdata/price-target-summary.json
git commit -m "feat(analyst): PriceTargetConsensus + PriceTargetSummary"
```

---

## Task 4: financial-estimates (params, 합성 struct) (TDD)

**Files:** Create `analyst/estimates.go`, `analyst/estimates_test.go`, `analyst/testdata/analyst-estimates.json`.

> 주의: analyst-estimates 카탈로그 응답 예시 없음 → 아래 struct/fixture 는 FMP 공개 shape 합성. Task 5 통합테스트(라이브)에서 실 shape 확인 후 다르면 조정 + status 보고.

- [ ] **Step 1: estimates.go**
```go
package analyst

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// FinancialEstimate — 애널리스트 재무 추정 (analyst-estimates). 카탈로그 예시 없음 → FMP 공개 shape 합성.
type FinancialEstimate struct {
	Symbol             string  `json:"symbol"`             // 종목 심볼
	Date               string  `json:"date"`               // 추정 기준일
	RevenueLow         int64   `json:"revenueLow"`         // 매출 추정 최저
	RevenueHigh        int64   `json:"revenueHigh"`        // 매출 추정 최고
	RevenueAvg         int64   `json:"revenueAvg"`         // 매출 추정 평균
	EbitdaLow          int64   `json:"ebitdaLow"`          // EBITDA 최저
	EbitdaHigh         int64   `json:"ebitdaHigh"`         // EBITDA 최고
	EbitdaAvg          int64   `json:"ebitdaAvg"`          // EBITDA 평균
	EbitLow            int64   `json:"ebitLow"`            // EBIT 최저
	EbitHigh           int64   `json:"ebitHigh"`           // EBIT 최고
	EbitAvg            int64   `json:"ebitAvg"`            // EBIT 평균
	NetIncomeLow       int64   `json:"netIncomeLow"`       // 순이익 최저
	NetIncomeHigh      int64   `json:"netIncomeHigh"`      // 순이익 최고
	NetIncomeAvg       int64   `json:"netIncomeAvg"`       // 순이익 평균
	SgaExpenseLow      int64   `json:"sgaExpenseLow"`      // 판관비 최저
	SgaExpenseHigh     int64   `json:"sgaExpenseHigh"`     // 판관비 최고
	SgaExpenseAvg      int64   `json:"sgaExpenseAvg"`      // 판관비 평균
	EpsLow             float64 `json:"epsLow"`             // EPS 최저
	EpsHigh            float64 `json:"epsHigh"`            // EPS 최고
	EpsAvg             float64 `json:"epsAvg"`             // EPS 평균
	NumAnalystsRevenue int     `json:"numAnalystsRevenue"` // 매출 추정 애널리스트 수
	NumAnalystsEps     int     `json:"numAnalystsEps"`     // EPS 추정 애널리스트 수
}

// FinancialEstimates 는 종목의 애널리스트 재무 추정을 조회한다.
// period 는 "annual" 또는 "quarter". page 는 0부터.
func (c *Client) FinancialEstimates(ctx context.Context, symbol, period string, page int) ([]FinancialEstimate, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	if strings.TrimSpace(period) == "" {
		return nil, fmt.Errorf("fmp: period must not be empty (annual|quarter)")
	}
	return fetch.List[FinancialEstimate](ctx, c.http, "/stable/analyst-estimates", map[string]string{
		"symbol": symbol,
		"period": period,
		"page":   strconv.Itoa(page),
	})
}
```

- [ ] **Step 2: fixture (합성)**

`analyst/testdata/analyst-estimates.json`:
```json
[
  {
    "symbol": "AAPL", "date": "2025-09-30",
    "revenueLow": 400000000000, "revenueHigh": 420000000000, "revenueAvg": 410000000000,
    "ebitdaLow": 130000000000, "ebitdaHigh": 140000000000, "ebitdaAvg": 135000000000,
    "ebitLow": 120000000000, "ebitHigh": 128000000000, "ebitAvg": 124000000000,
    "netIncomeLow": 95000000000, "netIncomeHigh": 105000000000, "netIncomeAvg": 100000000000,
    "sgaExpenseLow": 25000000000, "sgaExpenseHigh": 27000000000, "sgaExpenseAvg": 26000000000,
    "epsLow": 6.5, "epsHigh": 7.1, "epsAvg": 6.8,
    "numAnalystsRevenue": 14, "numAnalystsEps": 16
  }
]
```

- [ ] **Step 3: estimates_test.go**
```go
package analyst

import (
	"context"
	"os"
	"testing"
)

func TestFinancialEstimates_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/analyst-estimates.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.FinancialEstimates(context.Background(), "AAPL", "annual", 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].RevenueAvg <= 0 || rows[0].EpsAvg <= 0 || rows[0].NumAnalystsEps <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/analyst-estimates" || cap.query.Get("symbol") != "AAPL" || cap.query.Get("period") != "annual" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q symbol=%q period=%q page=%q", cap.path, cap.query.Get("symbol"), cap.query.Get("period"), cap.query.Get("page"))
	}
}

func TestFinancialEstimates_Guards(t *testing.T) {
	c, cleanup := newTestClient(t, 200, `[]`)
	defer cleanup()
	if _, err := c.FinancialEstimates(context.Background(), "  ", "annual", 0); err == nil {
		t.Error("want empty symbol guard")
	}
	if _, err := c.FinancialEstimates(context.Background(), "AAPL", "", 0); err == nil {
		t.Error("want empty period guard")
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./analyst/ && go vet ./analyst/ && gofmt -l analyst/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add analyst/estimates.go analyst/estimates_test.go analyst/testdata/analyst-estimates.json
git commit -m "feat(analyst): FinancialEstimates (period/page params, 합성 struct)"
```

---

## Task 5: 루트 와이어 + README + examples + 통합 + 검증 (TDD)

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/analyst/main.go`.

- [ ] **Step 1: client_test.go 실패 테스트**
```go
func TestNewClient_HasAnalyst(t *testing.T) {
	c, err := NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Analyst == nil {
		t.Fatal("Analyst sub-client is nil")
	}
}
```

- [ ] **Step 2: client.go 와이어** — import `"github.com/kenshin579/fmp-go/analyst"`; `Client` 에 `Analyst *analyst.Client // 애널리스트(등급/목표주가/추정)` 필드; `NewClient` 에 `c.Analyst = analyst.New(hc)`.

- [ ] **Step 3: 통과** — `go build ./... && go test . -run TestNewClient_HasAnalyst`. PASS.

- [ ] **Step 4: README + examples**

README 커버리지 표 행 추가:
```markdown
| Analyst | `client.Analyst` | Grades, GradesConsensus, HistoricalGrades, RatingsSnapshot, HistoricalRatings, PriceTargetConsensus, PriceTargetSummary, FinancialEstimates — 8 endpoint |
```

Create `examples/analyst/main.go`:
```go
// 실행: FMP_API_KEY=... go run examples/analyst/main.go
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

	pt, err := c.Analyst.PriceTargetConsensus(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 목표가 컨센서스: %.2f (고 %.2f / 저 %.2f)\n", pt.TargetConsensus, pt.TargetHigh, pt.TargetLow)

	g, err := c.Analyst.GradesConsensus(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("등급 컨센서스: %s (매수 %d / 보유 %d / 매도 %d)\n", g.Consensus, g.Buy, g.Hold, g.Sell)
}
```

- [ ] **Step 5: integration_test.go 에 analyst 케이스**
```go
func TestIntegration_Analyst(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if g, err := c.Analyst.GradesConsensus(ctx, "AAPL"); err != nil || g.Symbol != "AAPL" {
		t.Errorf("GradesConsensus: err=%v g=%+v", err, g)
	}
	if pt, err := c.Analyst.PriceTargetConsensus(ctx, "AAPL"); err != nil || pt.TargetConsensus <= 0 {
		t.Errorf("PriceTargetConsensus: err=%v pt=%+v", err, pt)
	}
	if rows, err := c.Analyst.FinancialEstimates(ctx, "AAPL", "annual", 0); err != nil || len(rows) == 0 {
		t.Errorf("FinancialEstimates: err=%v len=%d", err, len(rows))
	} else {
		t.Logf("FinancialEstimate[0]: %+v", rows[0]) // 합성 struct 실 shape 확인
	}
}
```

- [ ] **Step 6: 전체 검증**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./...
go test -tags integration ./... -run TestIntegration_Analyst -v 2>&1 | tail -12
gofmt -l .
```
Expected: 단위 전체 PASS, gofmt clean. 통합 key 없으면 skip.
> FinancialEstimate 실 shape 가 합성과 다르면 `analyst/estimates.go` struct + fixture 조정 후 재커밋.

- [ ] **Step 7: Commit**
```bash
git add client.go client_test.go README.md examples/analyst/main.go integration_test.go
git commit -m "feat(analyst): 루트 Client 와이어 + README + examples + 통합"
```

---

## 자기 점검 메모 (작성자용)
- 단일 *T 4개(GradesConsensus/RatingsSnapshot/PriceTargetConsensus/PriceTargetSummary) → OneBySymbol(빈 symbol 가드 + ErrNotFound). list 3개(Grades/HistoricalGrades/HistoricalRatings) → ListBySymbol.
- Rating 공유(snapshot date "" / historical date 값) — 양쪽 fixture 검증.
- FinancialEstimate 합성 — 통합 로그로 실 shape 확인(shares-float 선례).
- financial-estimates 빈 symbol/period 가드 + {symbol,period,page}.
- PriceTargetSummary.Publishers = JSON 배열 문자열(string 그대로).
- HistoricalGrade StrongBuy 없음(FMP 응답 그대로).
- test helper(newTestClient/newCapturingClient/capturedReq)는 Task 1 grades_test.go 정의 → Task 2-4 재사용(동일 패키지).
- 와이어: `Analyst *analyst.Client` + `c.Analyst = analyst.New(hc)`.
