# FMP Go SDK — Company 그룹 완성 + internal/fetch (v0.4.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** company 카테고리의 남은 16 endpoint 추가(17/17 완성) + generic helper 를 `internal/fetch` 공유 패키지로 hoist(quote 리팩터 포함). v0.4.0.

**Architecture:** 신규 `internal/fetch` 가 OneBySymbol/ListBySymbol/ListBySymbols/One/List 5 helper 제공. quote 가 로컬 helper 제거하고 fetch.* 호출(기존 테스트로 회귀 보증). company 가 fetch.* 로 16 메서드 구현, struct 재사용 적극. 모든 struct 필드 한국어 주석, fixture + delegation 테스트.

**Tech Stack:** Go 1.25+ generics / `internal/httpclient` / fixture 단위테스트 + build-tag 통합. `unset GOROOT` 필요시.

**Spec:** `docs/superpowers/specs/2026-06-02-company-group-design.md`
**Repo / Branch:** `github.com/kenshin579/fmp-go`, branch `feature/company-group` (spec 커밋 이미).

**확정된 사실 (조사 완료):**
- company 17 endpoint 중 Profile(`/stable/profile?symbol=`) 만 구현됨. 남은 16개 path 확인 완료(아래 매핑).
- 응답 shape 확인: market-cap `{symbol,date,marketCap}`; peers `{symbol,companyName,price,mktCap}`; employee-count `{symbol,cik,acceptanceTime,periodOfReport,companyName,formType,filingDate,employeeCount,source}`; executives `{title,name,pay(null),currencyPay,gender,yearBorn(null),titleSince(null),active}`; executive-compensation `{cik,symbol,companyName,filingDate,acceptedDate,nameAndPosition,year,salary,bonus,stockAward,optionAward,incentivePlanCompensation,allOtherCompensation,total}`; benchmark `{industryTitle,year,averageCompensation}`; company-notes `{cik,symbol,title,exchange}`; mergers `{symbol,companyName,cik,targetedCompanyName,targetedCik,targetedSymbol,transactionDate,acceptedDate,link}`; delisted `{symbol,companyName,exchange,ipoDate,delistedDate}`; profile-cik → 기존 Profile shape.
- **shares-float / all-shares-float 는 카탈로그에 응답 예시 없음** → FMP 공개 shape `{symbol,date,freeFloat,floatShares,outstandingShares}` 로 합성 + 통합테스트(라이브) 검증. 가정과 다르면 조정.
- 기존 `company/profile.go`: `Profile` struct + `func (c *Client) Profile(ctx, symbol) (*Profile, error)` (strings.TrimSpace 가드 + ErrNotFound). `company/client.go`: `Client{http}` + `New`.
- quote: `quote/client.go` 에 로컬 `fetchOne`/`fetchBatch`/`fetchList`; 공개 메서드들이 이를 호출. 18 테스트(fixture/delegation/가드) 존재.
- `httpclient.GetJSON(ctx, path, map[string]string, out any) error` + `httpclient.ErrNotFound`.

---

## File Structure
- Create: `internal/fetch/fetch.go` — 5 generic helper.
- Create: `internal/fetch/fetch_test.go` — helper 단위테스트.
- Modify: `quote/client.go` — 로컬 helper 제거.
- Modify: `quote/quote.go`/`short.go`/`change.go`/`aftermarket.go`/`asset_class.go` — fetch.* 호출로 교체.
- Modify: `company/profile.go` — `ProfileByCIK` 추가 (+ Profile 의 fetch 사용 통일).
- Create: `company/market_cap.go` + `_test.go` + testdata — MarketCap.
- Create: `company/shares_float.go` + `_test.go` + testdata — SharesFloat.
- Create: `company/employees.go` + `_test.go` + testdata — EmployeeCount.
- Create: `company/executives.go` + `_test.go` + testdata — Executive/ExecutiveCompensation/Benchmark.
- Create: `company/peers.go` + `_test.go` + testdata — Peer.
- Create: `company/notes.go` + `_test.go` + testdata — CompanyNote.
- Create: `company/mergers.go` + `_test.go` + testdata — MergerAcquisition.
- Create: `company/delisted.go` + `_test.go` + testdata — DelistedCompany.
- Modify: `README.md` — Company 행 17 endpoint.
- Create: `examples/company/main.go`.
- Modify: `integration_test.go` — company 통합.

---

## Task 1: `internal/fetch` 공유 helper (TDD)

**Files:** Create `internal/fetch/fetch.go`, `internal/fetch/fetch_test.go`.

- [ ] **Step 1: fetch.go 작성**

Create `internal/fetch/fetch.go`:
```go
// Package fetch 는 FMP sub-client 들이 공유하는 generic 조회 helper.
// 모든 그룹 패키지(quote/company/...)가 이 helper 로 endpoint 를 구현한다.
package fetch

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// OneBySymbol — {symbol} 단일 레코드. 빈 symbol 가드, 빈 배열 → httpclient.ErrNotFound.
func OneBySymbol[T any](ctx context.Context, hc *httpclient.Client, path, symbol string) (*T, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return One[T](ctx, hc, path, map[string]string{"symbol": symbol})
}

// ListBySymbol — {symbol} 리스트(시계열/다건). 빈 symbol 가드.
func ListBySymbol[T any](ctx context.Context, hc *httpclient.Client, path, symbol string) ([]T, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return List[T](ctx, hc, path, map[string]string{"symbol": symbol})
}

// ListBySymbols — {symbols:쉼표 join} 배치 리스트. 빈 symbols 가드.
func ListBySymbols[T any](ctx context.Context, hc *httpclient.Client, path string, symbols []string) ([]T, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("fmp: symbols must not be empty")
	}
	return List[T](ctx, hc, path, map[string]string{"symbols": strings.Join(symbols, ",")})
}

// One — 임의 params 단일 레코드. 빈 배열 → httpclient.ErrNotFound.
func One[T any](ctx context.Context, hc *httpclient.Client, path string, params map[string]string) (*T, error) {
	var out []T
	if err := hc.GetJSON(ctx, path, params, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return &out[0], nil
}

// List — 임의 params 리스트.
func List[T any](ctx context.Context, hc *httpclient.Client, path string, params map[string]string) ([]T, error) {
	var out []T
	err := hc.GetJSON(ctx, path, params, &out)
	return out, err
}
```

- [ ] **Step 2: fetch_test.go 작성**

Create `internal/fetch/fetch_test.go`:
```go
package fetch

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

type rec struct {
	Symbol string `json:"symbol"`
	V      int    `json:"v"`
}

func newHC(t *testing.T, body string, capture *captured) *httpclient.Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if capture != nil {
			capture.path = r.URL.Path
			capture.query = r.URL.Query()
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL})
}

type captured struct {
	path  string
	query url.Values
}

func TestOneBySymbol_ParsesAndDelegates(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[{"symbol":"AAPL","v":1}]`, cap)
	got, err := OneBySymbol[rec](context.Background(), hc, "/stable/x", "AAPL")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if got.Symbol != "AAPL" || got.V != 1 {
		t.Errorf("got %+v", got)
	}
	if cap.path != "/stable/x" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}

func TestOneBySymbol_EmptyArrayNotFound(t *testing.T) {
	hc := newHC(t, `[]`, nil)
	if _, err := OneBySymbol[rec](context.Background(), hc, "/x", "AAPL"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestOneBySymbol_EmptySymbolGuard(t *testing.T) {
	hc := newHC(t, `[]`, nil)
	if _, err := OneBySymbol[rec](context.Background(), hc, "/x", "  "); err == nil {
		t.Fatal("want guard error")
	}
}

func TestListBySymbols_JoinsAndGuards(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[{"symbol":"A"},{"symbol":"B"}]`, cap)
	rows, err := ListBySymbols[rec](context.Background(), hc, "/x", []string{"A", "B"})
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.query.Get("symbols") != "A,B" {
		t.Errorf("symbols=%q want A,B", cap.query.Get("symbols"))
	}
	if _, err := ListBySymbols[rec](context.Background(), hc, "/x", nil); err == nil {
		t.Fatal("want empty symbols guard")
	}
}

func TestList_ArbitraryParamsAndEmptyOK(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[]`, cap)
	rows, err := List[rec](context.Background(), hc, "/x", map[string]string{"page": "2"})
	if err != nil {
		t.Fatalf("empty list should not error: %v", err)
	}
	if len(rows) != 0 {
		t.Errorf("rows=%+v", rows)
	}
	if cap.query.Get("page") != "2" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
}

func TestListBySymbol_ParsesAndDelegates(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[{"symbol":"AAPL","v":1},{"symbol":"AAPL","v":2}]`, cap)
	rows, err := ListBySymbol[rec](context.Background(), hc, "/x", "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
}
```

- [ ] **Step 3: 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./internal/fetch/ -v && go vet ./internal/fetch/
```
Expected: 6 테스트 PASS, vet clean.

- [ ] **Step 4: Commit**
```bash
git add internal/fetch/
git commit -m "feat(fetch): 공유 generic helper (OneBySymbol/ListBySymbol/ListBySymbols/One/List)"
```

---

## Task 2: quote 리팩터 — fetch.* 사용 (회귀 안전)

**Files:** Modify `quote/client.go`, `quote/quote.go`, `quote/short.go`, `quote/change.go`, `quote/aftermarket.go`, `quote/asset_class.go`.

- [ ] **Step 1: quote/client.go 에서 로컬 helper 제거**

`quote/client.go` 의 `fetchOne`/`fetchBatch`/`fetchList` 3 함수 삭제. `Client`/`New` 만 남김. import 정리(`fmt`,`strings`,`httpclient` 가 더 안 쓰이면 제거 — 단 New 가 httpclient 쓰면 유지). 결과:
```go
package quote

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 시세 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
```

- [ ] **Step 2: 각 메서드를 fetch.* 호출로 교체**

`quote/quote.go`:
```go
package quote

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Quote struct 는 그대로 유지(필드/주석 변경 없음).

func (c *Client) Quote(ctx context.Context, symbol string) (*Quote, error) {
	return fetch.OneBySymbol[Quote](ctx, c.http, "/stable/quote", symbol)
}

func (c *Client) BatchQuote(ctx context.Context, symbols ...string) ([]Quote, error) {
	return fetch.ListBySymbols[Quote](ctx, c.http, "/stable/batch-quote", symbols)
}
```

`quote/short.go` — `fetch.OneBySymbol[QuoteShort](..., "/stable/quote-short", symbol)` / `fetch.ListBySymbols[QuoteShort](..., "/stable/batch-quote-short", symbols)`.

`quote/change.go` — `fetch.OneBySymbol[PriceChange](..., "/stable/stock-price-change", symbol)`.

`quote/aftermarket.go`:
- `AftermarketQuote` → `fetch.OneBySymbol[AftermarketQuote](..., "/stable/aftermarket-quote", symbol)`
- `AftermarketTrade` → `fetch.OneBySymbol[AftermarketTrade](..., "/stable/aftermarket-trade", symbol)`
- `BatchAftermarketQuote` → `fetch.ListBySymbols[AftermarketQuote](..., "/stable/batch-aftermarket-quote", symbols)`
- `BatchAftermarketTrade` → `fetch.ListBySymbols[AftermarketTrade](..., "/stable/batch-aftermarket-trade", symbols)`

`quote/asset_class.go`:
- `ExchangeQuotes` → `fetch.List[QuoteShort](ctx, c.http, "/stable/batch-exchange-quote", map[string]string{"exchange": exchange})`
- `IndexQuotes`/`CommodityQuotes`/`CryptoQuotes`/`ETFQuotes`/`ForexQuotes`/`MutualFundQuotes` → `fetch.List[QuoteShort](ctx, c.http, "<path>", nil)`

각 파일에 `import ("context"; "github.com/kenshin579/fmp-go/internal/fetch")` 정리.

- [ ] **Step 3: 통과 확인 (기존 18 테스트 무수정 회귀)**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./quote/ && go test ./quote/ -v 2>&1 | tail -25
```
Expected: 기존 18 테스트 전부 PASS(테스트 파일 무수정). vet clean.

- [ ] **Step 4: Commit**
```bash
git add quote/
git commit -m "refactor(quote): 로컬 helper 제거하고 internal/fetch 사용"
```

---

## Task 3: company profile-cik + market-cap (TDD)

**Files:** Modify `company/profile.go`; Create `company/market_cap.go`, `company/market_cap_test.go`, `company/testdata/market-cap-aapl.json`, `company/testdata/historical-market-cap-aapl.json`, `company/testdata/batch-market-cap.json`, `company/testdata/profile-cik-aapl.json`.

- [ ] **Step 1: profile.go — ProfileByCIK 추가**

`company/profile.go` 에 추가(기존 Profile struct/메서드 유지). 파일 상단 import 에 `"github.com/kenshin579/fmp-go/internal/fetch"`, `"strings"`, `"fmt"` 확인:
```go
// ProfileByCIK 는 CIK 로 회사 프로필을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) ProfileByCIK(ctx context.Context, cik string) (*Profile, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.One[Profile](ctx, c.http, "/stable/profile-cik", map[string]string{"cik": cik})
}
```
(기존 `Profile` 메서드는 그대로 두되, 선택적으로 `fetch.OneBySymbol` 로 통일 가능 — 본 task 범위에선 ProfileByCIK 만 추가, 기존 Profile 미수정.)

fixture `company/testdata/profile-cik-aapl.json` (profile shape 일부):
```json
[{ "symbol": "AAPL", "price": 262.82, "marketCap": 3900351299800, "companyName": "Apple Inc.", "currency": "USD", "cik": "0000320193", "isin": "US0378331005", "exchange": "NASDAQ" }]
```

- [ ] **Step 2: market_cap.go**

Create `company/market_cap.go`:
```go
package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// MarketCap — 시가총액 (market-capitalization / historical / batch 공용)
type MarketCap struct {
	Symbol    string `json:"symbol"`    // 종목 심볼
	Date      string `json:"date"`      // 기준일 (YYYY-MM-DD)
	MarketCap int64  `json:"marketCap"` // 시가총액
}

// MarketCap 은 종목의 현재 시가총액을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) MarketCap(ctx context.Context, symbol string) (*MarketCap, error) {
	return fetch.OneBySymbol[MarketCap](ctx, c.http, "/stable/market-capitalization", symbol)
}

// HistoricalMarketCap 은 종목의 시가총액 시계열을 조회한다.
func (c *Client) HistoricalMarketCap(ctx context.Context, symbol string) ([]MarketCap, error) {
	return fetch.ListBySymbol[MarketCap](ctx, c.http, "/stable/historical-market-capitalization", symbol)
}

// BatchMarketCap 은 여러 종목의 시가총액을 한 번에 조회한다.
func (c *Client) BatchMarketCap(ctx context.Context, symbols ...string) ([]MarketCap, error) {
	return fetch.ListBySymbols[MarketCap](ctx, c.http, "/stable/market-capitalization-batch", symbols)
}
```

fixtures:
`company/testdata/market-cap-aapl.json`:
```json
[{ "symbol": "AAPL", "date": "2025-10-24", "marketCap": 3900351299800 }]
```
`company/testdata/historical-market-cap-aapl.json`:
```json
[
  { "symbol": "AAPL", "date": "2026-04-08", "marketCap": 3818298106199 },
  { "symbol": "AAPL", "date": "2026-04-07", "marketCap": 3800000000000 }
]
```
`company/testdata/batch-market-cap.json`:
```json
[
  { "symbol": "AAPL", "date": "2025-10-24", "marketCap": 3900351299800 },
  { "symbol": "MSFT", "date": "2025-10-24", "marketCap": 3050000000000 }
]
```

- [ ] **Step 3: market_cap_test.go**

Create `company/market_cap_test.go`:
```go
package company

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

type capturedReq struct {
	path  string
	query url.Values
}

func TestMarketCap_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/market-cap-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	m, err := c.MarketCap(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("MarketCap: %v", err)
	}
	if m.Symbol != "AAPL" || m.MarketCap <= 0 || m.Date == "" {
		t.Errorf("not parsed: %+v", m)
	}
}

func TestHistoricalMarketCap_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/historical-market-cap-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HistoricalMarketCap(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
}

func TestBatchMarketCap_DelegatesSymbols(t *testing.T) {
	raw, _ := os.ReadFile("testdata/batch-market-cap.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.BatchMarketCap(context.Background(), "AAPL", "MSFT")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/market-capitalization-batch" || cap.query.Get("symbols") != "AAPL,MSFT" {
		t.Errorf("delegation: path=%q symbols=%q", cap.path, cap.query.Get("symbols"))
	}
}

func TestProfileByCIK_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/profile-cik-aapl.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	p, err := c.ProfileByCIK(context.Background(), "320193")
	if err != nil {
		t.Fatalf("ProfileByCIK: %v", err)
	}
	if p.Symbol != "AAPL" {
		t.Errorf("Symbol=%q", p.Symbol)
	}
	if cap.path != "/stable/profile-cik" || cap.query.Get("cik") != "320193" {
		t.Errorf("delegation: path=%q cik=%q", cap.path, cap.query.Get("cik"))
	}
	if _, err := c.ProfileByCIK(context.Background(), "  "); err == nil {
		t.Fatal("want empty cik guard")
	}
}
```

- [ ] **Step 4: 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./company/ -v 2>&1 | tail -25 && go vet ./company/
```
Expected: 신규 테스트 + 기존 profile 테스트 PASS, vet clean.

- [ ] **Step 5: Commit**
```bash
git add company/profile.go company/market_cap.go company/market_cap_test.go company/testdata/market-cap-aapl.json company/testdata/historical-market-cap-aapl.json company/testdata/batch-market-cap.json company/testdata/profile-cik-aapl.json
git commit -m "feat(company): ProfileByCIK + MarketCap(현재/historical/batch)"
```

---

## Task 4: shares-float + employee-count (TDD)

**Files:** Create `company/shares_float.go`, `company/employees.go`, 각 `_test.go`, testdata.

- [ ] **Step 1: shares_float.go**

> **주의**: shares-float / all-shares-float 는 카탈로그에 응답 예시가 없음. 아래 struct 는 FMP 공개 docs shape 기반 합성 — Step 4 통합테스트(라이브)에서 실제 필드 확인하고, 다르면 struct/fixture 조정 후 status 에 보고.

Create `company/shares_float.go`:
```go
package company

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SharesFloat — 유통 주식 수(float)
type SharesFloat struct {
	Symbol            string  `json:"symbol"`            // 종목 심볼
	Date              string  `json:"date"`              // 기준일
	FreeFloat         float64 `json:"freeFloat"`         // 유통 비율 (%)
	FloatShares       int64   `json:"floatShares"`       // 유통 주식 수
	OutstandingShares int64   `json:"outstandingShares"` // 발행 주식 총수
}

// SharesFloat 은 종목의 유통 주식 정보를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) SharesFloat(ctx context.Context, symbol string) (*SharesFloat, error) {
	return fetch.OneBySymbol[SharesFloat](ctx, c.http, "/stable/shares-float", symbol)
}

// AllSharesFloat 은 전체 종목의 유통 주식 정보를 페이지 단위로 조회한다.
func (c *Client) AllSharesFloat(ctx context.Context, page int) ([]SharesFloat, error) {
	return fetch.List[SharesFloat](ctx, c.http, "/stable/shares-float-all", map[string]string{"page": strconv.Itoa(page)})
}
```

fixtures (합성):
`company/testdata/shares-float-aapl.json`:
```json
[{ "symbol": "AAPL", "date": "2025-10-24", "freeFloat": 99.92, "floatShares": 14800000000, "outstandingShares": 14840000000 }]
```
`company/testdata/all-shares-float.json`:
```json
[
  { "symbol": "AAPL", "date": "2025-10-24", "freeFloat": 99.92, "floatShares": 14800000000, "outstandingShares": 14840000000 },
  { "symbol": "MSFT", "date": "2025-10-24", "freeFloat": 99.7, "floatShares": 7400000000, "outstandingShares": 7430000000 }
]
```

- [ ] **Step 2: employees.go**

Create `company/employees.go`:
```go
package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// EmployeeCount — 직원 수 (employee-count / historical 공용, SEC 공시 기준)
type EmployeeCount struct {
	Symbol         string `json:"symbol"`         // 종목 심볼
	CIK            string `json:"cik"`            // SEC CIK
	AcceptanceTime string `json:"acceptanceTime"` // 공시 수리 시각
	PeriodOfReport string `json:"periodOfReport"` // 보고 기준일
	CompanyName    string `json:"companyName"`    // 회사명
	FormType       string `json:"formType"`       // 공시 양식 (예: 10-K)
	FilingDate     string `json:"filingDate"`     // 공시일
	EmployeeCount  int64  `json:"employeeCount"`  // 직원 수
	Source         string `json:"source"`         // 원문 URL
}

// EmployeeCount 는 종목의 최신 직원 수를 조회한다.
func (c *Client) EmployeeCount(ctx context.Context, symbol string) ([]EmployeeCount, error) {
	return fetch.ListBySymbol[EmployeeCount](ctx, c.http, "/stable/employee-count", symbol)
}

// HistoricalEmployeeCount 는 종목의 직원 수 시계열을 조회한다.
func (c *Client) HistoricalEmployeeCount(ctx context.Context, symbol string) ([]EmployeeCount, error) {
	return fetch.ListBySymbol[EmployeeCount](ctx, c.http, "/stable/historical-employee-count", symbol)
}
```

fixtures:
`company/testdata/employee-count-aapl.json`:
```json
[{ "symbol": "AAPL", "cik": "0000320193", "acceptanceTime": "2025-10-31 06:01:26", "periodOfReport": "2025-09-27", "companyName": "Apple Inc.", "formType": "10-K", "filingDate": "2025-10-31", "employeeCount": 166000, "source": "https://www.sec.gov/..." }]
```
`company/testdata/historical-employee-count-aapl.json`:
```json
[
  { "symbol": "AAPL", "cik": "0000320193", "periodOfReport": "2025-09-27", "companyName": "Apple Inc.", "formType": "10-K", "filingDate": "2025-10-31", "employeeCount": 166000, "source": "x" },
  { "symbol": "AAPL", "cik": "0000320193", "periodOfReport": "2024-09-28", "companyName": "Apple Inc.", "formType": "10-K", "filingDate": "2024-11-01", "employeeCount": 164000, "source": "x" }
]
```

- [ ] **Step 3: 테스트**

Create `company/shares_float_test.go`:
```go
package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSharesFloat_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/shares-float-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	s, err := c.SharesFloat(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("SharesFloat: %v", err)
	}
	if s.Symbol != "AAPL" || s.OutstandingShares <= 0 {
		t.Errorf("not parsed: %+v", s)
	}
}

func TestAllSharesFloat_DelegatesPage(t *testing.T) {
	raw, _ := os.ReadFile("testdata/all-shares-float.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.AllSharesFloat(context.Background(), 2)
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/shares-float-all" || cap.query.Get("page") != "2" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}
```

Create `company/employees_test.go`:
```go
package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestEmployeeCount_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/employee-count-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.EmployeeCount(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].EmployeeCount <= 0 || rows[0].FormType == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestHistoricalEmployeeCount_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/historical-employee-count-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HistoricalEmployeeCount(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
}
```

- [ ] **Step 4: 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./company/ && go vet ./company/
# (선택) shares-float 실 shape 확인:
# FMP_API_KEY=... go test -tags integration ./... -run TestIntegration_Company -v   # Task 7 후 가능
```
Expected: 전체 PASS, vet clean.

- [ ] **Step 5: Commit**
```bash
git add company/shares_float.go company/shares_float_test.go company/employees.go company/employees_test.go company/testdata/shares-float-aapl.json company/testdata/all-shares-float.json company/testdata/employee-count-aapl.json company/testdata/historical-employee-count-aapl.json
git commit -m "feat(company): SharesFloat(+all) + EmployeeCount(+historical)"
```

---

## Task 5: executives (key-executives + compensation + benchmark) (TDD)

**Files:** Create `company/executives.go`, `company/executives_test.go`, testdata 3개.

- [ ] **Step 1: executives.go**

Create `company/executives.go`:
```go
package company

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Executive — 주요 임원 (key-executives). pay/yearBorn/titleSince 는 결측 가능(null) → 포인터.
type Executive struct {
	Title       string   `json:"title"`       // 직책
	Name        string   `json:"name"`        // 이름
	Pay         *int64   `json:"pay"`         // 보수(결측 가능)
	CurrencyPay string   `json:"currencyPay"` // 보수 통화
	Gender      string   `json:"gender"`      // 성별
	YearBorn    *int     `json:"yearBorn"`    // 출생연도(결측 가능)
	TitleSince  *int64   `json:"titleSince"`  // 현 직책 부임(결측 가능)
	Active      bool     `json:"active"`      // 재직 여부
}

// ExecutiveCompensation — 임원 보수 공시 (governance-executive-compensation)
type ExecutiveCompensation struct {
	CIK                       string `json:"cik"`                       // SEC CIK
	Symbol                    string `json:"symbol"`                    // 종목 심볼
	CompanyName               string `json:"companyName"`               // 회사명
	FilingDate                string `json:"filingDate"`                // 공시일
	AcceptedDate              string `json:"acceptedDate"`              // 수리일시
	NameAndPosition           string `json:"nameAndPosition"`           // 이름·직책
	Year                      int    `json:"year"`                      // 회계연도
	Salary                    int64  `json:"salary"`                    // 급여
	Bonus                     int64  `json:"bonus"`                     // 상여
	StockAward                int64  `json:"stockAward"`                // 주식 보상
	OptionAward               int64  `json:"optionAward"`               // 옵션 보상
	IncentivePlanCompensation int64  `json:"incentivePlanCompensation"` // 성과급
	AllOtherCompensation      int64  `json:"allOtherCompensation"`      // 기타 보상
	Total                     int64  `json:"total"`                     // 총 보수
}

// ExecutiveCompensationBenchmark — 산업별 평균 임원 보수
type ExecutiveCompensationBenchmark struct {
	IndustryTitle       string  `json:"industryTitle"`       // 산업 분류
	Year                int     `json:"year"`                // 연도
	AverageCompensation float64 `json:"averageCompensation"` // 평균 보수
}

// KeyExecutives 는 종목의 주요 임원 목록을 조회한다.
func (c *Client) KeyExecutives(ctx context.Context, symbol string) ([]Executive, error) {
	return fetch.ListBySymbol[Executive](ctx, c.http, "/stable/key-executives", symbol)
}

// ExecutiveCompensation 은 종목의 임원 보수 공시를 조회한다.
func (c *Client) ExecutiveCompensation(ctx context.Context, symbol string) ([]ExecutiveCompensation, error) {
	return fetch.ListBySymbol[ExecutiveCompensation](ctx, c.http, "/stable/governance-executive-compensation", symbol)
}

// ExecutiveCompensationBenchmark 은 연도별 산업 평균 임원 보수를 조회한다.
func (c *Client) ExecutiveCompensationBenchmark(ctx context.Context, year int) ([]ExecutiveCompensationBenchmark, error) {
	params := map[string]string{}
	if year > 0 {
		params["year"] = strconv.Itoa(year)
	}
	return fetch.List[ExecutiveCompensationBenchmark](ctx, c.http, "/stable/executive-compensation-benchmark", params)
}
```

> **주의**: benchmark 파라미터가 `year` 가 맞는지 구현 시 카탈로그 `executive-compensation-benchmark.md` 의 Parameters 절 확인. 다르면 조정.

fixtures:
`company/testdata/key-executives-aapl.json`:
```json
[
  { "title": "Chief Executive Officer", "name": "Tim Cook", "pay": 16520856, "currencyPay": "USD", "gender": "male", "yearBorn": 1960, "titleSince": 1378944000, "active": true },
  { "title": "Senior Vice President of Worldwide Marketing", "name": "Greg Joswiak", "pay": null, "currencyPay": "USD", "gender": "male", "yearBorn": null, "titleSince": null, "active": true }
]
```
`company/testdata/executive-compensation-aapl.json`:
```json
[{ "cik": "0000320193", "symbol": "AAPL", "companyName": "Apple Inc.", "filingDate": "2026-01-08", "acceptedDate": "2026-01-08 16:31:36", "nameAndPosition": "Tim Cook Chief Executive Officer", "year": 2025, "salary": 3000000, "bonus": 0, "stockAward": 57535293, "optionAward": 0, "incentivePlanCompensation": 12000000, "allOtherCompensation": 1759518, "total": 74294811 }]
```
`company/testdata/executive-compensation-benchmark.json`:
```json
[
  { "industryTitle": "ABRASIVE, ASBESTOS & MISC NONMETALLIC MINERAL PRODS", "year": 2024, "averageCompensation": 784407.5555555555 },
  { "industryTitle": "ACCIDENT & HEALTH INSURANCE", "year": 2024, "averageCompensation": 1200000.0 }
]
```

- [ ] **Step 2: executives_test.go**

Create `company/executives_test.go`:
```go
package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestKeyExecutives_ParsesFixtureWithNullable(t *testing.T) {
	raw, _ := os.ReadFile("testdata/key-executives-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.KeyExecutives(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	// 첫 행: pay/yearBorn 존재
	if rows[0].Pay == nil || *rows[0].Pay <= 0 {
		t.Errorf("row0 Pay should be set: %+v", rows[0])
	}
	if rows[0].YearBorn == nil || *rows[0].YearBorn != 1960 {
		t.Errorf("row0 YearBorn: %v", rows[0].YearBorn)
	}
	// 둘째 행: pay/yearBorn null → nil 포인터
	if rows[1].Pay != nil {
		t.Errorf("row1 Pay should be nil, got %v", *rows[1].Pay)
	}
	if rows[1].YearBorn != nil {
		t.Errorf("row1 YearBorn should be nil")
	}
}

func TestExecutiveCompensation_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/executive-compensation-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.ExecutiveCompensation(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Total <= 0 || rows[0].Year == 0 || rows[0].NameAndPosition == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestExecutiveCompensationBenchmark_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/executive-compensation-benchmark.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.ExecutiveCompensationBenchmark(context.Background(), 2024)
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].AverageCompensation <= 0 || rows[0].IndustryTitle == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/executive-compensation-benchmark" || cap.query.Get("year") != "2024" {
		t.Errorf("delegation: path=%q year=%q", cap.path, cap.query.Get("year"))
	}
}
```

- [ ] **Step 3: 통과 확인** — `go test ./company/ && go vet ./company/`. PASS.

- [ ] **Step 4: Commit**
```bash
git add company/executives.go company/executives_test.go company/testdata/key-executives-aapl.json company/testdata/executive-compensation-aapl.json company/testdata/executive-compensation-benchmark.json
git commit -m "feat(company): KeyExecutives + ExecutiveCompensation + Benchmark"
```

---

## Task 6: peers + notes + mergers + delisted (TDD)

**Files:** Create `company/peers.go`, `company/notes.go`, `company/mergers.go`, `company/delisted.go`, 각 `_test.go`, testdata.

- [ ] **Step 1: 4개 파일 작성**

`company/peers.go`:
```go
package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Peer — 동종업계 비교 종목
type Peer struct {
	Symbol      string  `json:"symbol"`      // 종목 심볼
	CompanyName string  `json:"companyName"` // 회사명
	Price       float64 `json:"price"`       // 현재가
	MktCap      int64   `json:"mktCap"`      // 시가총액
}

// StockPeers 는 종목의 동종업계 비교 종목을 조회한다.
func (c *Client) StockPeers(ctx context.Context, symbol string) ([]Peer, error) {
	return fetch.ListBySymbol[Peer](ctx, c.http, "/stable/stock-peers", symbol)
}
```

`company/notes.go`:
```go
package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CompanyNote — 회사 채권/노트 발행 정보
type CompanyNote struct {
	CIK      string `json:"cik"`      // SEC CIK
	Symbol   string `json:"symbol"`   // 종목 심볼
	Title    string `json:"title"`    // 노트명 (예: 0.000% Notes due 2025)
	Exchange string `json:"exchange"` // 거래소
}

// CompanyNotes 는 회사가 발행한 채권/노트 목록을 조회한다.
func (c *Client) CompanyNotes(ctx context.Context, symbol string) ([]CompanyNote, error) {
	return fetch.ListBySymbol[CompanyNote](ctx, c.http, "/stable/company-notes", symbol)
}
```

`company/mergers.go`:
```go
package company

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// MergerAcquisition — M&A 공시 (latest / search 공용)
type MergerAcquisition struct {
	Symbol              string `json:"symbol"`              // 인수 회사 심볼
	CompanyName         string `json:"companyName"`         // 인수 회사명
	CIK                 string `json:"cik"`                 // 인수 회사 CIK
	TargetedCompanyName string `json:"targetedCompanyName"` // 피인수 회사명
	TargetedCik         string `json:"targetedCik"`         // 피인수 회사 CIK
	TargetedSymbol      string `json:"targetedSymbol"`      // 피인수 회사 심볼
	TransactionDate     string `json:"transactionDate"`     // 거래일
	AcceptedDate        string `json:"acceptedDate"`        // 수리일시
	Link                string `json:"link"`                // 원문 URL
}

// LatestMergersAcquisitions 는 최신 M&A 공시를 페이지 단위로 조회한다.
func (c *Client) LatestMergersAcquisitions(ctx context.Context, page int) ([]MergerAcquisition, error) {
	return fetch.List[MergerAcquisition](ctx, c.http, "/stable/mergers-acquisitions-latest", map[string]string{"page": strconv.Itoa(page)})
}

// SearchMergersAcquisitions 는 회사명으로 M&A 공시를 검색한다.
func (c *Client) SearchMergersAcquisitions(ctx context.Context, name string) ([]MergerAcquisition, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[MergerAcquisition](ctx, c.http, "/stable/mergers-acquisitions-search", map[string]string{"name": name})
}
```

`company/delisted.go`:
```go
package company

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// DelistedCompany — 상장폐지 종목
type DelistedCompany struct {
	Symbol       string `json:"symbol"`       // 종목 심볼
	CompanyName  string `json:"companyName"`  // 회사명
	Exchange     string `json:"exchange"`     // 거래소
	IPODate      string `json:"ipoDate"`      // 상장일
	DelistedDate string `json:"delistedDate"` // 상장폐지일
}

// DelistedCompanies 는 상장폐지 종목을 페이지 단위로 조회한다.
func (c *Client) DelistedCompanies(ctx context.Context, page int) ([]DelistedCompany, error) {
	return fetch.List[DelistedCompany](ctx, c.http, "/stable/delisted-companies", map[string]string{"page": strconv.Itoa(page)})
}
```

- [ ] **Step 2: fixtures**

`company/testdata/peers-aapl.json`:
```json
[
  { "symbol": "GOOGL", "companyName": "Alphabet Inc.", "price": 317.32, "mktCap": 3838620208180 },
  { "symbol": "MSFT", "companyName": "Microsoft Corporation", "price": 410.5, "mktCap": 3050000000000 }
]
```
`company/testdata/company-notes-aapl.json`:
```json
[{ "cik": "0000320193", "symbol": "AAPL", "title": "0.000% Notes due 2025", "exchange": "NASDAQ" }]
```
`company/testdata/mergers-latest.json`:
```json
[{ "symbol": "ALGT", "companyName": "Allegiant Travel CO", "cik": "0001362468", "targetedCompanyName": "Sun Country Airlines Holdings, Inc.", "targetedCik": "0001743907", "targetedSymbol": "SNCY", "transactionDate": "2026-03-27", "acceptedDate": "2026-03-27 17:15:41", "link": "https://www.sec.gov/..." }]
```
`company/testdata/delisted-companies.json`:
```json
[{ "symbol": "5CV.DE", "companyName": "CureVac N.V.", "exchange": "XETRA", "ipoDate": "2020-08-25", "delistedDate": "2026-12-05" }]
```

- [ ] **Step 3: 테스트**

Create `company/peers_test.go`:
```go
package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestStockPeers_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/peers-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.StockPeers(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol == "" || rows[0].MktCap <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
```

Create `company/notes_test.go`:
```go
package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestCompanyNotes_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/company-notes-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CompanyNotes(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Title == "" || rows[0].Exchange == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
```

Create `company/mergers_test.go`:
```go
package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestLatestMergersAcquisitions_DelegatesPage(t *testing.T) {
	raw, _ := os.ReadFile("testdata/mergers-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.LatestMergersAcquisitions(context.Background(), 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].TargetedSymbol == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/mergers-acquisitions-latest" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}

func TestSearchMergersAcquisitions_DelegatesNameAndGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/mergers-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	if _, err := c.SearchMergersAcquisitions(context.Background(), "Apple"); err != nil {
		t.Fatalf("Search: %v", err)
	}
	if cap.path != "/stable/mergers-acquisitions-search" || cap.query.Get("name") != "Apple" {
		t.Errorf("delegation: path=%q name=%q", cap.path, cap.query.Get("name"))
	}
	if _, err := c.SearchMergersAcquisitions(context.Background(), "  "); err == nil {
		t.Fatal("want empty name guard")
	}
}
```

Create `company/delisted_test.go`:
```go
package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestDelistedCompanies_DelegatesPage(t *testing.T) {
	raw, _ := os.ReadFile("testdata/delisted-companies.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.DelistedCompanies(context.Background(), 1)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].DelistedDate == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/delisted-companies" || cap.query.Get("page") != "1" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./company/ && go vet ./company/`. 전체 PASS.

- [ ] **Step 5: Commit**
```bash
git add company/peers.go company/notes.go company/mergers.go company/delisted.go company/peers_test.go company/notes_test.go company/mergers_test.go company/delisted_test.go company/testdata/peers-aapl.json company/testdata/company-notes-aapl.json company/testdata/mergers-latest.json company/testdata/delisted-companies.json
git commit -m "feat(company): StockPeers + CompanyNotes + Mergers(latest/search) + DelistedCompanies"
```

---

## Task 7: README + examples + 통합 + 전체 검증

**Files:** Modify `README.md`, `integration_test.go`; Create `examples/company/main.go`.

- [ ] **Step 1: README 커버리지 Company 행 갱신**

`README.md` 커버리지 표의 Company 행을 17 endpoint 반영으로 교체:
```markdown
| Company | `client.Company` | Profile, ProfileByCIK, MarketCap(+historical/batch), SharesFloat(+all), EmployeeCount(+historical), KeyExecutives, ExecutiveCompensation(+benchmark), StockPeers, CompanyNotes, Mergers(latest/search), DelistedCompanies — 17 endpoint |
```

- [ ] **Step 2: examples/company/main.go**

Create `examples/company/main.go` (기존 examples 무태그 컨벤션):
```go
// 실행: FMP_API_KEY=... go run examples/company/main.go
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

	mc, err := c.Company.MarketCap(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL marketCap=%d (%s)\n", mc.MarketCap, mc.Date)

	peers, err := c.Company.StockPeers(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("peers: %d개\n", len(peers))

	execs, err := c.Company.KeyExecutives(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range execs {
		fmt.Printf("  %s — %s\n", e.Title, e.Name)
	}
}
```
(`examples/profile/main.go` 와 같은 package main 무태그 — 두 example 이 같은 디렉토리가 아니므로 충돌 없음. `examples/quote/main.go` 와도 별 디렉토리.)

- [ ] **Step 3: integration_test.go 에 company 케이스**

`integration_test.go` (기존 `package fmp_test` + `//go:build integration` + `NewClientFromEnv`) 에 추가:
```go
func TestIntegration_Company(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if mc, err := c.Company.MarketCap(ctx, "AAPL"); err != nil || mc.MarketCap <= 0 {
		t.Errorf("MarketCap: err=%v mc=%+v", err, mc)
	}
	if rows, err := c.Company.StockPeers(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("StockPeers: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Company.KeyExecutives(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("KeyExecutives: err=%v len=%d", err, len(rows))
	}
	// shares-float 실 shape 확인(카탈로그 예시 없던 endpoint)
	if sf, err := c.Company.SharesFloat(ctx, "AAPL"); err != nil {
		t.Errorf("SharesFloat: %v", err)
	} else {
		t.Logf("SharesFloat AAPL: %+v", sf) // 실 필드 확인용
	}
	if rows, err := c.Company.DelistedCompanies(ctx, 0); err != nil || len(rows) == 0 {
		t.Errorf("DelistedCompanies: err=%v len=%d", err, len(rows))
	}
}
```

- [ ] **Step 4: 전체 검증**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./...
go test -tags integration ./... -run 'TestIntegration_Company' -v 2>&1 | tail -15   # FMP_API_KEY 있으면 실호출(특히 SharesFloat shape 로그 확인)
gofmt -l .
```
Expected: 단위 전체 PASS, gofmt clean. 통합은 key 있으면 PASS(+SharesFloat 실 shape 로그) / 없으면 skip.
> SharesFloat 실 shape 가 합성 struct 와 다르면 `company/shares_float.go` struct + fixture 조정 후 재커밋.

- [ ] **Step 5: Commit**
```bash
git add README.md examples/company/main.go integration_test.go
git commit -m "docs(company): README 커버리지 + examples + 통합 테스트"
```

---

## 자기 점검 메모 (작성자용)
- **internal/fetch hoist**: quote 리팩터(Task 2)는 기존 18 테스트가 무수정 통과해야 안전. 실패 시 helper 시그니처/위임 재확인.
- **struct 재사용**: MarketCap(현재 *T / historical·batch []T), SharesFloat(*T / all []T), EmployeeCount(둘 다 []T), MergerAcquisition(latest·search []T), Profile(profile-cik *T).
- **nullable 필드**: Executive 의 pay/yearBorn/titleSince → 포인터(*int64/*int). fixture 가 null/값 양쪽 케이스 검증.
- **카탈로그 예시 없는 endpoint**: shares-float/all-shares-float — 합성 struct + 통합테스트 로그로 실 shape 확인. benchmark 파라미터(year)도 구현 시 카탈로그 Parameters 절 확인.
- **path 정확성**: company-executives → `/stable/key-executives`, executive-compensation → `/stable/governance-executive-compensation`, peers → `/stable/stock-peers`, profile-cik → `/stable/profile-cik`. 모두 카탈로그 GET 줄 기준 확정.
- **페이지네이션**: page int → `strconv.Itoa(page)`. limit 은 FMP 기본값 사용(미전달).
- **company test helper**: `newTestClient` + `newCapturingClient` + `capturedReq` 는 Task 3 에서 정의(market_cap_test.go), 이후 task 들이 같은 패키지라 재사용.
- **delegation 테스트**: 공유 fetch 가 Task 1 에서 검증되므로 company 는 path/특수파라미터(cik/page/name/symbols) 매핑만 대표 확인.
