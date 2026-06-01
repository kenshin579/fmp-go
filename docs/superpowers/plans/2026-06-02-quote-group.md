# FMP Go SDK — Quote 그룹 (v0.3.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** FMP `quote` 카테고리 16 endpoint 를 신규 `quote/` 패키지로 추가하고 v0.3.0 릴리스 준비. 전체 API 커버리지 캠페인의 첫 그룹 — 이후 모든 그룹 PR 의 템플릿.

**Architecture:** `statements/` 구조 미러링 — 그룹 패키지 + 그룹별 파일. 16 메서드의 반복을 줄이기 위해 quote 패키지 안에 generic helper(`fetchOne[T]`/`fetchBatch[T]`/`fetchList[T]`) 3개를 두고, 공개 메서드는 한 줄 위임. 모든 응답 struct 필드에 한국어 주석.

**Tech Stack:** Go 1.25+ / `internal/httpclient` (기존) / fixture 기반 단위테스트 + build-tag 통합테스트. `unset GOROOT` 필요시.

**Spec:** `docs/superpowers/specs/2026-06-02-quote-group-design.md`
**Repo / Branch:** `github.com/kenshin579/fmp-go`, branch `feature/quote-group` (spec 커밋 `5dee8b6` 이미).

**확정된 사실 (조사 완료):**
- 16 endpoint path 전부 확인(카탈로그 GET 줄): quote / quote-short / stock-price-change / aftermarket-quote / aftermarket-trade / batch-quote / batch-quote-short / batch-aftermarket-quote / batch-aftermarket-trade / batch-exchange-quote(?exchange=) / batch-index-quotes / batch-commodity-quotes / batch-crypto-quotes / batch-etf-quotes / batch-forex-quotes / batch-mutualfund-quotes.
- 기존 패턴: `company.Profile` = 단일 `*T` + `strings.TrimSpace` 빈 가드 + 빈 배열 `httpclient.ErrNotFound`. `statements` = `[]T` + fixture 테스트(`newTestClient(t, status, body)` httptest 헬퍼).
- `httpclient.GetJSON(ctx, path, map[string]string, out any) error` + `httpclient.ErrNotFound`.
- root `client.go` 의 `NewClient` 가 `c.Company/Statements/Ratios = X.New(hc)` 와이어.
- asset-class 응답 shape: crypto/forex 는 `{symbol,price,change,volume}` = QuoteShort 확인. exchange 는 full Quote 가정(구현 시 fixture 확정). index/commodity/etf/mutualfund 도 QuoteShort 가정 — fixture 로 확정, 추가 필드 시 별도 struct.

---

## File Structure
- Create: `quote/client.go` — `Client` + `New` + generic helper 3개 + 빈 인자 sentinel.
- Create: `quote/quote.go` — `Quote` struct + `Quote()` + `BatchQuote()`.
- Create: `quote/quote_test.go` + `quote/testdata/quote-aapl.json`, `batch-quote.json`.
- Create: `quote/short.go` — `QuoteShort` struct + `QuoteShort()` + `BatchQuoteShort()`.
- Create: `quote/short_test.go` + testdata.
- Create: `quote/change.go` — `PriceChange` struct + `PriceChange()`.
- Create: `quote/change_test.go` + testdata.
- Create: `quote/aftermarket.go` — `AftermarketQuote`/`AftermarketTrade` struct + 4 메서드.
- Create: `quote/aftermarket_test.go` + testdata.
- Create: `quote/asset_class.go` — `ExchangeQuotes` + 6 asset-class 메서드.
- Create: `quote/asset_class_test.go` + testdata.
- Modify: `client.go` — `Quote *quote.Client` 필드 + 와이어.
- Modify: `client_test.go` — `TestNewClient_HasQuote`.
- Modify: `README.md` — 커버리지 표 Quote 행.
- Create: `examples/quote/main.go`.
- Modify: `integration_test.go` — quote 통합 케이스.

---

## Task 1: quote 패키지 기반 + Quote/BatchQuote (TDD)

**Files:** Create `quote/client.go`, `quote/quote.go`, `quote/quote_test.go`, `quote/testdata/quote-aapl.json`, `quote/testdata/batch-quote.json`.

- [ ] **Step 1: fixture 작성**

`quote/testdata/quote-aapl.json` (카탈로그 `docs/api/quote/quote.md` 응답 예시):
```json
[
  {
    "symbol": "AAPL",
    "name": "Apple Inc.",
    "price": 232.8,
    "changePercentage": 2.1008,
    "change": 4.79,
    "volume": 44489128,
    "dayLow": 226.65,
    "dayHigh": 233.13,
    "yearHigh": 260.1,
    "yearLow": 164.08,
    "marketCap": 3500823120000,
    "priceAvg50": 240.2278,
    "priceAvg200": 219.98755,
    "exchange": "NASDAQ",
    "open": 227.2,
    "previousClose": 228.01,
    "timestamp": 1738702801
  }
]
```

`quote/testdata/batch-quote.json` (2건 — AAPL + MSFT 합성):
```json
[
  {
    "symbol": "AAPL", "name": "Apple Inc.", "price": 232.8, "changePercentage": 2.1008,
    "change": 4.79, "volume": 44489128, "dayLow": 226.65, "dayHigh": 233.13,
    "yearHigh": 260.1, "yearLow": 164.08, "marketCap": 3500823120000,
    "priceAvg50": 240.2278, "priceAvg200": 219.98755, "exchange": "NASDAQ",
    "open": 227.2, "previousClose": 228.01, "timestamp": 1738702801
  },
  {
    "symbol": "MSFT", "name": "Microsoft Corporation", "price": 410.5, "changePercentage": 0.8,
    "change": 3.25, "volume": 18000000, "dayLow": 405.0, "dayHigh": 412.0,
    "yearHigh": 468.35, "yearLow": 366.5, "marketCap": 3050000000000,
    "priceAvg50": 420.1, "priceAvg200": 425.3, "exchange": "NASDAQ",
    "open": 407.0, "previousClose": 407.25, "timestamp": 1738702801
  }
]
```

- [ ] **Step 2: client.go 작성 (helper + sentinel)**

Create `quote/client.go`:
```go
// Package quote 는 FMP 시세(quote) API sub-client.
// fmp.Client.Quote 로 접근.
package quote

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 시세 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// fetchOne 은 단일 심볼 조회 공통 — FMP 배열 응답의 첫 요소를 *T 로 반환.
// 빈 symbol → 에러, 빈 배열 → httpclient.ErrNotFound.
func fetchOne[T any](ctx context.Context, c *Client, path, symbol string) (*T, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []T
	if err := c.http.GetJSON(ctx, path, map[string]string{"symbol": symbol}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return &out[0], nil
}

// fetchBatch 은 symbols 배치 조회 공통 — 쉼표 join.
func fetchBatch[T any](ctx context.Context, c *Client, path string, symbols []string) ([]T, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("fmp: symbols must not be empty")
	}
	var out []T
	err := c.http.GetJSON(ctx, path, map[string]string{"symbols": strings.Join(symbols, ",")}, &out)
	return out, err
}

// fetchList 은 단일 키 또는 무파라미터 리스트 조회 공통.
func fetchList[T any](ctx context.Context, c *Client, path string, params map[string]string) ([]T, error) {
	var out []T
	err := c.http.GetJSON(ctx, path, params, &out)
	return out, err
}
```

- [ ] **Step 3: quote.go 작성**

Create `quote/quote.go`:
```go
package quote

import "context"

// Quote — 전체 시세 (quote / batch-quote / exchange-quote 공용)
type Quote struct {
	Symbol           string  `json:"symbol"`           // 종목 심볼 (예: AAPL)
	Name             string  `json:"name"`             // 종목명
	Price            float64 `json:"price"`            // 현재가
	ChangePercentage float64 `json:"changePercentage"` // 등락률 (%)
	Change           float64 `json:"change"`           // 전일 대비 등락액
	Volume           int64   `json:"volume"`           // 거래량
	DayLow           float64 `json:"dayLow"`           // 당일 저가
	DayHigh          float64 `json:"dayHigh"`          // 당일 고가
	YearHigh         float64 `json:"yearHigh"`         // 52주 최고가
	YearLow          float64 `json:"yearLow"`          // 52주 최저가
	MarketCap        int64   `json:"marketCap"`        // 시가총액
	PriceAvg50       float64 `json:"priceAvg50"`       // 50일 이동평균가
	PriceAvg200      float64 `json:"priceAvg200"`      // 200일 이동평균가
	Exchange         string  `json:"exchange"`         // 거래소 (예: NASDAQ)
	Open             float64 `json:"open"`             // 시가
	PreviousClose    float64 `json:"previousClose"`    // 전일 종가
	Timestamp        int64   `json:"timestamp"`        // 시세 시각 (Unix epoch sec)
}

// Quote 는 종목의 전체 시세를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) Quote(ctx context.Context, symbol string) (*Quote, error) {
	return fetchOne[Quote](ctx, c, "/stable/quote", symbol)
}

// BatchQuote 는 여러 종목의 전체 시세를 한 번에 조회한다.
func (c *Client) BatchQuote(ctx context.Context, symbols ...string) ([]Quote, error) {
	return fetchBatch[Quote](ctx, c, "/stable/batch-quote", symbols)
}
```

- [ ] **Step 4: quote_test.go 작성**

Create `quote/quote_test.go`:
```go
package quote

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
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

func TestQuote_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/quote-aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	q, err := c.Quote(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("Quote: %v", err)
	}
	if q.Symbol != "AAPL" {
		t.Errorf("Symbol = %q", q.Symbol)
	}
	if q.Price <= 0 || q.MarketCap <= 0 || q.Volume <= 0 {
		t.Errorf("core numeric fields not parsed: %+v", q)
	}
	if q.Exchange != "NASDAQ" {
		t.Errorf("Exchange = %q", q.Exchange)
	}
	if q.Timestamp == 0 {
		t.Error("Timestamp not parsed")
	}
}

func TestQuote_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.Quote(context.Background(), "NOPE"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestQuote_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.Quote(context.Background(), "  "); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

func TestBatchQuote_ParsesFixtureAndEmptyGuard(t *testing.T) {
	raw, err := os.ReadFile("testdata/batch-quote.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil || len(arr) < 2 {
		t.Fatalf("fixture must have >=2 items")
	}
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.BatchQuote(context.Background(), "AAPL", "MSFT")
	if err != nil {
		t.Fatalf("BatchQuote: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("len = %d, want 2", len(rows))
	}
	if rows[1].Symbol != "MSFT" {
		t.Errorf("rows[1].Symbol = %q", rows[1].Symbol)
	}

	// 빈 symbols 가드
	if _, err := c.BatchQuote(context.Background()); err == nil {
		t.Fatal("expected error for empty symbols")
	}
}
```

- [ ] **Step 5: 실패 → 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./quote/ -run 'TestQuote|TestBatchQuote' -v
go vet ./quote/
```
Expected: 4 테스트 PASS, vet clean.

- [ ] **Step 6: Commit**
```bash
git add quote/client.go quote/quote.go quote/quote_test.go quote/testdata/quote-aapl.json quote/testdata/batch-quote.json
git commit -m "feat(quote): 패키지 기반 + Quote/BatchQuote"
```

---

## Task 2: QuoteShort/BatchQuoteShort (TDD)

**Files:** Create `quote/short.go`, `quote/short_test.go`, `quote/testdata/quote-short-aapl.json`, `quote/testdata/batch-quote-short.json`.

- [ ] **Step 1: fixture**

`quote/testdata/quote-short-aapl.json`:
```json
[{ "symbol": "AAPL", "price": 232.8, "change": 4.79, "volume": 44489128 }]
```
`quote/testdata/batch-quote-short.json`:
```json
[
  { "symbol": "AAPL", "price": 232.8, "change": 4.79, "volume": 44489128 },
  { "symbol": "MSFT", "price": 410.5, "change": 3.25, "volume": 18000000 }
]
```

- [ ] **Step 2: short.go**

Create `quote/short.go`:
```go
package quote

import "context"

// QuoteShort — 경량 시세 (quote-short / batch-quote-short / 자산군 배치 공용)
type QuoteShort struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	Price  float64 `json:"price"`  // 현재가
	Change float64 `json:"change"` // 전일 대비 등락액
	Volume int64   `json:"volume"` // 거래량
}

// QuoteShort 는 종목의 경량 시세를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) QuoteShort(ctx context.Context, symbol string) (*QuoteShort, error) {
	return fetchOne[QuoteShort](ctx, c, "/stable/quote-short", symbol)
}

// BatchQuoteShort 는 여러 종목의 경량 시세를 한 번에 조회한다.
func (c *Client) BatchQuoteShort(ctx context.Context, symbols ...string) ([]QuoteShort, error) {
	return fetchBatch[QuoteShort](ctx, c, "/stable/batch-quote-short", symbols)
}
```

- [ ] **Step 3: short_test.go**

Create `quote/short_test.go`:
```go
package quote

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func TestQuoteShort_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/quote-short-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	q, err := c.QuoteShort(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("QuoteShort: %v", err)
	}
	if q.Symbol != "AAPL" || q.Price <= 0 || q.Volume <= 0 {
		t.Errorf("not parsed: %+v", q)
	}
}

func TestQuoteShort_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.QuoteShort(context.Background(), "NOPE"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestBatchQuoteShort_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/batch-quote-short.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.BatchQuoteShort(context.Background(), "AAPL", "MSFT")
	if err != nil {
		t.Fatalf("BatchQuoteShort: %v", err)
	}
	if len(rows) != 2 || rows[1].Symbol != "MSFT" {
		t.Errorf("rows = %+v", rows)
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./quote/ -run Short -v && go vet ./quote/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add quote/short.go quote/short_test.go quote/testdata/quote-short-aapl.json quote/testdata/batch-quote-short.json
git commit -m "feat(quote): QuoteShort/BatchQuoteShort"
```

---

## Task 3: PriceChange (TDD)

**Files:** Create `quote/change.go`, `quote/change_test.go`, `quote/testdata/price-change-aapl.json`.

- [ ] **Step 1: fixture**

`quote/testdata/price-change-aapl.json`:
```json
[
  {
    "symbol": "AAPL",
    "1D": 2.1008, "5D": -2.45946, "1M": -4.33925, "3M": 4.86014, "6M": 5.88556,
    "ytd": -4.53147, "1Y": 24.04092, "3Y": 35.04264, "5Y": 192.05871,
    "10Y": 678.8558, "max": 181279.04168
  }
]
```

- [ ] **Step 2: change.go**

Create `quote/change.go`:
```go
package quote

import "context"

// PriceChange — 기간별 등락률(%) — stock-price-change
type PriceChange struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	D1     float64 `json:"1D"`     // 1일 등락률 (%)
	D5     float64 `json:"5D"`     // 5일
	M1     float64 `json:"1M"`     // 1개월
	M3     float64 `json:"3M"`     // 3개월
	M6     float64 `json:"6M"`     // 6개월
	YTD    float64 `json:"ytd"`    // 연초 대비
	Y1     float64 `json:"1Y"`     // 1년
	Y3     float64 `json:"3Y"`     // 3년
	Y5     float64 `json:"5Y"`     // 5년
	Y10    float64 `json:"10Y"`    // 10년
	Max    float64 `json:"max"`    // 상장 이후 전체
}

// PriceChange 는 종목의 기간별 등락률을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) PriceChange(ctx context.Context, symbol string) (*PriceChange, error) {
	return fetchOne[PriceChange](ctx, c, "/stable/stock-price-change", symbol)
}
```

- [ ] **Step 3: change_test.go**

Create `quote/change_test.go`:
```go
package quote

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestPriceChange_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/price-change-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	pc, err := c.PriceChange(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("PriceChange: %v", err)
	}
	if pc.Symbol != "AAPL" {
		t.Errorf("Symbol = %q", pc.Symbol)
	}
	// 숫자 시작 JSON 키 매핑 검증
	if pc.D1 == 0 || pc.Y1 == 0 || pc.Max == 0 {
		t.Errorf("period fields not parsed: %+v", pc)
	}
	if pc.YTD == 0 {
		t.Error("YTD not parsed")
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./quote/ -run PriceChange -v && go vet ./quote/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add quote/change.go quote/change_test.go quote/testdata/price-change-aapl.json
git commit -m "feat(quote): PriceChange (기간별 등락률)"
```

---

## Task 4: Aftermarket (quote/trade + batch) (TDD)

**Files:** Create `quote/aftermarket.go`, `quote/aftermarket_test.go`, `quote/testdata/aftermarket-quote-aapl.json`, `aftermarket-trade-aapl.json`, `batch-aftermarket-quote.json`, `batch-aftermarket-trade.json`.

- [ ] **Step 1: fixtures**

`aftermarket-quote-aapl.json`:
```json
[{ "symbol": "AAPL", "bidSize": 1, "bidPrice": 232.45, "askSize": 3, "askPrice": 232.64, "volume": 41647042, "timestamp": 1738715334311 }]
```
`aftermarket-trade-aapl.json`:
```json
[{ "symbol": "AAPL", "price": 232.53, "tradeSize": 132, "timestamp": 1738715334311 }]
```
`batch-aftermarket-quote.json`:
```json
[
  { "symbol": "AAPL", "bidSize": 1, "bidPrice": 232.45, "askSize": 3, "askPrice": 232.64, "volume": 41647042, "timestamp": 1738715334311 },
  { "symbol": "MSFT", "bidSize": 2, "bidPrice": 410.1, "askSize": 1, "askPrice": 410.4, "volume": 9000000, "timestamp": 1738715334311 }
]
```
`batch-aftermarket-trade.json`:
```json
[
  { "symbol": "AAPL", "price": 232.53, "tradeSize": 132, "timestamp": 1738715334311 },
  { "symbol": "MSFT", "price": 410.2, "tradeSize": 50, "timestamp": 1738715334311 }
]
```

- [ ] **Step 2: aftermarket.go**

Create `quote/aftermarket.go`:
```go
package quote

import "context"

// AftermarketQuote — 시간외 호가
type AftermarketQuote struct {
	Symbol    string  `json:"symbol"`    // 종목 심볼
	BidSize   int64   `json:"bidSize"`   // 매수 호가 수량
	BidPrice  float64 `json:"bidPrice"`  // 매수 호가
	AskSize   int64   `json:"askSize"`   // 매도 호가 수량
	AskPrice  float64 `json:"askPrice"`  // 매도 호가
	Volume    int64   `json:"volume"`    // 거래량
	Timestamp int64   `json:"timestamp"` // 시각 (Unix epoch ms)
}

// AftermarketTrade — 시간외 체결
type AftermarketTrade struct {
	Symbol    string  `json:"symbol"`    // 종목 심볼
	Price     float64 `json:"price"`     // 체결가
	TradeSize int64   `json:"tradeSize"` // 체결 수량
	Timestamp int64   `json:"timestamp"` // 시각 (Unix epoch ms)
}

// AftermarketQuote 는 종목의 시간외 호가를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) AftermarketQuote(ctx context.Context, symbol string) (*AftermarketQuote, error) {
	return fetchOne[AftermarketQuote](ctx, c, "/stable/aftermarket-quote", symbol)
}

// AftermarketTrade 는 종목의 시간외 체결을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) AftermarketTrade(ctx context.Context, symbol string) (*AftermarketTrade, error) {
	return fetchOne[AftermarketTrade](ctx, c, "/stable/aftermarket-trade", symbol)
}

// BatchAftermarketQuote 는 여러 종목의 시간외 호가를 한 번에 조회한다.
func (c *Client) BatchAftermarketQuote(ctx context.Context, symbols ...string) ([]AftermarketQuote, error) {
	return fetchBatch[AftermarketQuote](ctx, c, "/stable/batch-aftermarket-quote", symbols)
}

// BatchAftermarketTrade 는 여러 종목의 시간외 체결을 한 번에 조회한다.
func (c *Client) BatchAftermarketTrade(ctx context.Context, symbols ...string) ([]AftermarketTrade, error) {
	return fetchBatch[AftermarketTrade](ctx, c, "/stable/batch-aftermarket-trade", symbols)
}
```

- [ ] **Step 3: aftermarket_test.go**

Create `quote/aftermarket_test.go`:
```go
package quote

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestAftermarketQuote_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/aftermarket-quote-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	q, err := c.AftermarketQuote(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("AftermarketQuote: %v", err)
	}
	if q.Symbol != "AAPL" || q.BidPrice <= 0 || q.AskPrice <= 0 || q.Timestamp == 0 {
		t.Errorf("not parsed: %+v", q)
	}
}

func TestAftermarketTrade_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/aftermarket-trade-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	tr, err := c.AftermarketTrade(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("AftermarketTrade: %v", err)
	}
	if tr.Symbol != "AAPL" || tr.Price <= 0 || tr.TradeSize <= 0 {
		t.Errorf("not parsed: %+v", tr)
	}
}

func TestBatchAftermarket_ParsesFixtures(t *testing.T) {
	rawQ, _ := os.ReadFile("testdata/batch-aftermarket-quote.json")
	cQ, cleanupQ := newTestClient(t, http.StatusOK, string(rawQ))
	defer cleanupQ()
	qs, err := cQ.BatchAftermarketQuote(context.Background(), "AAPL", "MSFT")
	if err != nil || len(qs) != 2 || qs[1].Symbol != "MSFT" {
		t.Fatalf("BatchAftermarketQuote: err=%v rows=%+v", err, qs)
	}

	rawT, _ := os.ReadFile("testdata/batch-aftermarket-trade.json")
	cT, cleanupT := newTestClient(t, http.StatusOK, string(rawT))
	defer cleanupT()
	ts, err := cT.BatchAftermarketTrade(context.Background(), "AAPL", "MSFT")
	if err != nil || len(ts) != 2 || ts[1].Symbol != "MSFT" {
		t.Fatalf("BatchAftermarketTrade: err=%v rows=%+v", err, ts)
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./quote/ -run Aftermarket -v && go vet ./quote/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add quote/aftermarket.go quote/aftermarket_test.go quote/testdata/aftermarket-quote-aapl.json quote/testdata/aftermarket-trade-aapl.json quote/testdata/batch-aftermarket-quote.json quote/testdata/batch-aftermarket-trade.json
git commit -m "feat(quote): Aftermarket quote/trade + batch"
```

---

## Task 5: 거래소/자산군 배치 (TDD)

**Files:** Create `quote/asset_class.go`, `quote/asset_class_test.go`, `quote/testdata/exchange-quotes.json`, `quote/testdata/crypto-quotes.json`.

- [ ] **Step 1: fixtures**

`quote/testdata/exchange-quotes.json` (full Quote shape, 2건):
```json
[
  {
    "symbol": "AAPL", "name": "Apple Inc.", "price": 232.8, "changePercentage": 2.1,
    "change": 4.79, "volume": 44489128, "dayLow": 226.65, "dayHigh": 233.13,
    "yearHigh": 260.1, "yearLow": 164.08, "marketCap": 3500823120000,
    "priceAvg50": 240.2, "priceAvg200": 219.9, "exchange": "NASDAQ",
    "open": 227.2, "previousClose": 228.01, "timestamp": 1738702801
  },
  {
    "symbol": "MSFT", "name": "Microsoft Corporation", "price": 410.5, "changePercentage": 0.8,
    "change": 3.25, "volume": 18000000, "dayLow": 405.0, "dayHigh": 412.0,
    "yearHigh": 468.35, "yearLow": 366.5, "marketCap": 3050000000000,
    "priceAvg50": 420.1, "priceAvg200": 425.3, "exchange": "NASDAQ",
    "open": 407.0, "previousClose": 407.25, "timestamp": 1738702801
  }
]
```
`quote/testdata/crypto-quotes.json` (QuoteShort shape):
```json
[
  { "symbol": "BTCUSD", "price": 96000.12, "change": 1200.5, "volume": 35000 },
  { "symbol": "ETHUSD", "price": 3400.55, "change": -45.2, "volume": 120000 }
]
```

- [ ] **Step 2: asset_class.go**

Create `quote/asset_class.go`:
```go
package quote

import "context"

// ExchangeQuotes 는 특정 거래소의 전체 종목 시세를 조회한다 (예: NASDAQ).
func (c *Client) ExchangeQuotes(ctx context.Context, exchange string) ([]Quote, error) {
	return fetchList[Quote](ctx, c, "/stable/batch-exchange-quote", map[string]string{"exchange": exchange})
}

// IndexQuotes 는 전체 지수 시세를 조회한다.
func (c *Client) IndexQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-index-quotes", nil)
}

// CommodityQuotes 는 전체 원자재 시세를 조회한다.
func (c *Client) CommodityQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-commodity-quotes", nil)
}

// CryptoQuotes 는 전체 암호화폐 시세를 조회한다.
func (c *Client) CryptoQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-crypto-quotes", nil)
}

// ETFQuotes 는 전체 ETF 시세를 조회한다.
func (c *Client) ETFQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-etf-quotes", nil)
}

// ForexQuotes 는 전체 외환 시세를 조회한다.
func (c *Client) ForexQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-forex-quotes", nil)
}

// MutualFundQuotes 는 전체 뮤추얼펀드 시세를 조회한다.
func (c *Client) MutualFundQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-mutualfund-quotes", nil)
}
```

> **구현 시 검증**: ExchangeQuotes 가 full `Quote` shape 인지, asset-class(index/commodity/etf/forex/mutualfund)가 `QuoteShort` 인지 카탈로그 `docs/api/quote/full-*.md` 응답 예시로 재확인. 다르면 그 그룹만 별도 struct + fixture 추가.

- [ ] **Step 3: asset_class_test.go**

Create `quote/asset_class_test.go`:
```go
package quote

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestExchangeQuotes_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/exchange-quotes.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.ExchangeQuotes(context.Background(), "NASDAQ")
	if err != nil {
		t.Fatalf("ExchangeQuotes: %v", err)
	}
	if len(rows) != 2 || rows[0].Symbol != "AAPL" || rows[0].MarketCap <= 0 {
		t.Errorf("rows = %+v", rows)
	}
}

func TestCryptoQuotes_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crypto-quotes.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CryptoQuotes(context.Background())
	if err != nil {
		t.Fatalf("CryptoQuotes: %v", err)
	}
	if len(rows) != 2 || rows[0].Symbol != "BTCUSD" || rows[0].Price <= 0 {
		t.Errorf("rows = %+v", rows)
	}
}
```

(index/commodity/etf/forex/mutualfund 은 CryptoQuotes 와 동일 shape·경로 패턴이라 대표 1개(crypto)만 fixture 검증. 나머지는 통합 테스트(Task 8)에서 라이브 확인.)

- [ ] **Step 4: 통과 확인** — `go test ./quote/ && go vet ./quote/`. 전체 PASS.

- [ ] **Step 5: Commit**
```bash
git add quote/asset_class.go quote/asset_class_test.go quote/testdata/exchange-quotes.json quote/testdata/crypto-quotes.json
git commit -m "feat(quote): 거래소/자산군 배치 시세 (exchange/index/commodity/crypto/etf/forex/mutualfund)"
```

---

## Task 6: 루트 Client 와이어 (TDD)

**Files:** Modify `client.go`, `client_test.go`.

- [ ] **Step 1: client_test.go 에 실패 테스트 추가**

`client_test.go` 에 추가:
```go
func TestNewClient_HasQuote(t *testing.T) {
	c, err := NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Quote == nil {
		t.Fatal("Quote sub-client is nil")
	}
}
```

- [ ] **Step 2: 실패 확인** — `go test . -run TestNewClient_HasQuote` → `c.Quote undefined` 컴파일 실패.

- [ ] **Step 3: client.go 와이어**

import 에 `"github.com/kenshin579/fmp-go/quote"` 추가. `Client` 구조체에 필드 추가:
```go
	Quote      *quote.Client      // 시세(실시간/배치/자산군)
```
`NewClient` 에 추가:
```go
	c.Quote = quote.New(hc)
```

- [ ] **Step 4: 통과 확인** — `go build ./... && go vet ./... && go test .`. PASS.

- [ ] **Step 5: Commit**
```bash
git add client.go client_test.go
git commit -m "feat(quote): 루트 Client 에 Quote sub-client 와이어"
```

---

## Task 7: README 커버리지 + examples

**Files:** Modify `README.md`; Create `examples/quote/main.go`.

- [ ] **Step 1: README 커버리지 표 갱신**

`README.md` 의 커버리지 표에 행 추가 (기존 Company/Statements/Ratios 행 아래):
```markdown
| Quote | `client.Quote` | Quote, QuoteShort, PriceChange, Aftermarket(Quote/Trade), Batch(Quote/Short/Aftermarket), 자산군(Exchange/Index/Commodity/Crypto/ETF/Forex/MutualFund) — `/stable/quote` 외 16 endpoint |
```
(기존 표 형식에 맞춰 컬럼 수 일치. 표가 Company 한 줄만 있으면 Statements/Ratios 행도 같이 채워 최신화.)

- [ ] **Step 2: examples/quote/main.go**

Create `examples/quote/main.go`:
```go
//go:build ignore

// 실행: FMP_API_KEY=... go run examples/quote/main.go
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

	q, err := c.Quote.Quote(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL %.2f (%.2f%%) vol=%d\n", q.Price, q.ChangePercentage, q.Volume)

	batch, err := c.Quote.BatchQuote(ctx, "AAPL", "MSFT", "GOOGL")
	if err != nil {
		log.Fatal(err)
	}
	for _, b := range batch {
		fmt.Printf("  %s %.2f\n", b.Symbol, b.Price)
	}

	pc, err := c.Quote.PriceChange(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 1Y=%.2f%% YTD=%.2f%%\n", pc.Y1, pc.YTD)

	crypto, err := c.Quote.CryptoQuotes(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("crypto quotes: %d개\n", len(crypto))
}
```
(`//go:build ignore` 로 일반 빌드 제외 — 기존 `examples/profile` 패턴 확인 후 동일하게. profile 예시가 build tag 없으면 그 방식 따름.)

- [ ] **Step 3: 검증** — `go build ./... && go vet ./...`. PASS. (examples 가 build ignore 면 `go vet ./examples/...` 별도 확인.)

- [ ] **Step 4: Commit**
```bash
git add README.md examples/quote/main.go
git commit -m "docs(quote): README 커버리지 + examples/quote"
```

---

## Task 8: 통합 테스트 + 전체 검증

**Files:** Modify `integration_test.go`.

- [ ] **Step 1: 통합 테스트 추가**

`integration_test.go` (build tag `//go:build integration`) 에 quote 케이스 추가. 기존 통합 테스트 스타일(FMP_API_KEY skip) 따름:
```go
func TestIntegration_Quote(t *testing.T) {
	key := os.Getenv("FMP_API_KEY")
	if key == "" {
		t.Skip("FMP_API_KEY 미설정 — 통합 테스트 skip")
	}
	c, err := fmp.NewClient(key)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	q, err := c.Quote.Quote(ctx, "AAPL")
	if err != nil {
		t.Fatalf("Quote: %v", err)
	}
	if q.Symbol != "AAPL" || q.Price <= 0 {
		t.Errorf("quote = %+v", q)
	}

	if _, err := c.Quote.QuoteShort(ctx, "AAPL"); err != nil {
		t.Errorf("QuoteShort: %v", err)
	}
	if _, err := c.Quote.PriceChange(ctx, "AAPL"); err != nil {
		t.Errorf("PriceChange: %v", err)
	}
	if rows, err := c.Quote.BatchQuote(ctx, "AAPL", "MSFT"); err != nil || len(rows) == 0 {
		t.Errorf("BatchQuote: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Quote.CryptoQuotes(ctx); err != nil || len(rows) == 0 {
		t.Errorf("CryptoQuotes: err=%v len=%d", err, len(rows))
	}
}
```
(기존 `integration_test.go` 의 import/패키지명 확인 후 맞춤. 파일 없으면 신규 생성 — build tag + package 선언.)

- [ ] **Step 2: 전체 검증**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./...
go test -tags integration ./... -run TestIntegration_Quote -v   # FMP_API_KEY 있으면 실호출, 없으면 skip
gofmt -l .   # 빈 출력이어야
```
Expected: 단위 전체 PASS, gofmt clean. 통합은 key 있으면 PASS / 없으면 skip.

- [ ] **Step 3: Commit**
```bash
git add integration_test.go
git commit -m "test(quote): 통합 테스트 (FMP_API_KEY 라이브)"
```

---

## 자기 점검 메모 (작성자용)
- **generic helper 도입**: `fetchOne[T]`/`fetchBatch[T]`/`fetchList[T]` — 16 메서드 한 줄 위임. 기존 statements/company 는 명시 코드였으나 16 endpoint 반복 회피 위해 generics 채택. 이후 그룹들도 같은 helper 패턴 재사용 가능(추후 internal 로 hoist 고려 — 본 PR 범위 밖).
- **단일 vs 배치 반환 타입**: 단일 `*T`(ErrNotFound), 배치/자산군 `[]T`(빈 배열 허용). Company.Profile / Statements 패턴 일관.
- **asset-class shape 가정**: exchange=full Quote, 나머지=QuoteShort. Task 5 에서 카탈로그로 재확인 — 다르면 별도 struct.
- **path 정확성**: stock-price-change(quote-change 파일), batch-*-quotes(full-* 파일). 카탈로그 GET 줄 기준 — 본 plan 의 path 는 검증 완료.
- **PriceChange 숫자 키**: `1D` 등 → `D1`/`M1`/`Y1` Go 필드 + 명시 json 태그. 테스트가 매핑 검증.
- **fixture 출처**: quote/quote-short/price-change/aftermarket 은 카탈로그 실응답 예시. batch/exchange 는 동일 shape 2건 합성(MSFT 추가).
- **examples build tag**: 기존 `examples/profile` 방식 확인 후 동일 적용(`//go:build ignore` 또는 무태그).
- **NewClientFromEnv 존재 확인**: examples/통합테스트가 사용 — `from_env.go` 에 있음(README 사용 예시에 등장).
