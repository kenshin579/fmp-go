# FMP Go SDK — Search 그룹 (v0.5.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** FMP `search` 카테고리 7 endpoint 를 신규 `search/` 패키지로 추가하고 v0.5.0 준비. 전체 API 커버리지 캠페인 3번째 그룹.

**Architecture:** `internal/fetch` 공유 helper 사용. 검색은 모두 다건 → `[]T` 반환. symbol/name 검색은 struct 공유. screener 는 19 필터 `ScreenerParams` + `toMap()`. 모든 struct 필드 한국어 주석.

**Tech Stack:** Go 1.25 generics / `internal/fetch` / fixture 단위테스트 + build-tag 통합. `unset GOROOT` 필요시.

**Spec:** `docs/superpowers/specs/2026-06-02-search-group-design.md`
**Repo / Branch:** `github.com/kenshin579/fmp-go`, branch `feature/search-group` (spec 커밋 이미).

**확정된 사실 (조사 완료):**
- 7 path: `/stable/search-symbol?query=`, `search-name?query=`, `search-cik?cik=`, `search-cusip?cusip=`, `search-isin?isin=`, `search-exchange-variants?symbol=`, `company-screener`(19 필터).
- 응답 shape: symbol/name `{symbol,name,currency,exchangeFullName,exchange}` 동일; cik `{symbol,companyName,cik,exchangeFullName,exchange,currency}`; cusip `{symbol,companyName,cusip,marketCap(float)}`; isin `{symbol,name,isin,marketCap(int)}`; exchange-variants profile 유사; screener `{symbol,companyName,marketCap(null),sector,industry,beta(null),price,lastAnnualDividend(null),volume,exchange,exchangeShortName,country,isEtf,isFund,isActivelyTrading}`.
- **marketCap 타입**: isin=int64, cusip=float64 (확인됨).
- `internal/fetch`: `ListBySymbol[T](ctx, hc, path, symbol)`(symbol 가드+키), `List[T](ctx, hc, path, params)`(임의 params, 가드 없음).
- root `client.go`: `Quote *quote.Client` 필드 + `c.Quote = quote.New(hc)` — Search 도 동형.
- 기존 테스트 헬퍼 패턴: `newTestClient`/`newCapturingClient`(company/quote 참고).

---

## File Structure
- Create: `search/client.go` — Client + New.
- Create: `search/search.go` + `_test.go` + testdata — SymbolSearchResult + SearchSymbol/SearchName.
- Create: `search/identifiers.go` + `_test.go` + testdata — CIK/CUSIP/ISIN.
- Create: `search/variants.go` + `_test.go` + testdata — ExchangeVariant.
- Create: `search/screener.go` + `_test.go` + testdata — ScreenerParams + ScreenerResult + CompanyScreener.
- Modify: `client.go` / `client_test.go` — Search 와이어.
- Modify: `README.md`; Create `examples/search/main.go`; Modify `integration_test.go`.

---

## Task 1: search 패키지 기반 + SearchSymbol/SearchName (TDD)

**Files:** Create `search/client.go`, `search/search.go`, `search/search_test.go`, `search/testdata/search-symbol-aapl.json`, `search/testdata/search-name-apple.json`.

- [ ] **Step 1: client.go**

Create `search/client.go`:
```go
// Package search 는 FMP 검색(search) API sub-client.
// fmp.Client.Search 로 접근.
package search

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 검색 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
```

- [ ] **Step 2: search.go**

Create `search/search.go`:
```go
package search

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SymbolSearchResult — 심볼/회사명 검색 결과 (search-symbol / search-name 공용)
type SymbolSearchResult struct {
	Symbol           string `json:"symbol"`           // 종목 심볼
	Name             string `json:"name"`             // 종목/회사명
	Currency         string `json:"currency"`         // 통화
	ExchangeFullName string `json:"exchangeFullName"` // 거래소 전체명
	Exchange         string `json:"exchange"`         // 거래소 코드
}

// SearchSymbol 은 심볼(티커)로 종목을 검색한다.
func (c *Client) SearchSymbol(ctx context.Context, query string) ([]SymbolSearchResult, error) {
	if strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("fmp: query must not be empty")
	}
	return fetch.List[SymbolSearchResult](ctx, c.http, "/stable/search-symbol", map[string]string{"query": query})
}

// SearchName 은 회사명으로 종목을 검색한다.
func (c *Client) SearchName(ctx context.Context, query string) ([]SymbolSearchResult, error) {
	if strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("fmp: query must not be empty")
	}
	return fetch.List[SymbolSearchResult](ctx, c.http, "/stable/search-name", map[string]string{"query": query})
}
```

- [ ] **Step 3: fixtures**

`search/testdata/search-symbol-aapl.json`:
```json
[{ "symbol": "AAPL", "name": "Apple Inc.", "currency": "USD", "exchangeFullName": "NASDAQ Global Select", "exchange": "NASDAQ" }]
```
`search/testdata/search-name-apple.json`:
```json
[
  { "symbol": "AAPL", "name": "Apple Inc.", "currency": "USD", "exchangeFullName": "NASDAQ Global Select", "exchange": "NASDAQ" },
  { "symbol": "APLE", "name": "Apple Hospitality REIT, Inc.", "currency": "USD", "exchangeFullName": "New York Stock Exchange", "exchange": "NYSE" }
]
```

- [ ] **Step 4: search_test.go (test helper 포함)**

Create `search/search_test.go`:
```go
package search

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

func TestSearchSymbol_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-symbol-aapl.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchSymbol(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].Exchange != "NASDAQ" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/search-symbol" || cap.query.Get("query") != "AAPL" {
		t.Errorf("delegation: path=%q query=%q", cap.path, cap.query.Get("query"))
	}
	if _, err := c.SearchSymbol(context.Background(), "  "); err == nil {
		t.Fatal("want empty query guard")
	}
}

func TestSearchName_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-name-apple.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchName(context.Background(), "Apple")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/search-name" || cap.query.Get("query") != "Apple" {
		t.Errorf("delegation: path=%q query=%q", cap.path, cap.query.Get("query"))
	}
}
```

- [ ] **Step 5: 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./search/ -v && go vet ./search/ && gofmt -l search/
```
Expected: 2 테스트 PASS, vet clean, gofmt 빈 출력.

- [ ] **Step 6: Commit**
```bash
git add search/client.go search/search.go search/search_test.go search/testdata/search-symbol-aapl.json search/testdata/search-name-apple.json
git commit -m "feat(search): 패키지 기반 + SearchSymbol/SearchName"
```

---

## Task 2: identifiers — CIK/CUSIP/ISIN (TDD)

**Files:** Create `search/identifiers.go`, `search/identifiers_test.go`, 3 testdata.

- [ ] **Step 1: identifiers.go**

Create `search/identifiers.go`:
```go
package search

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CIKSearchResult — CIK 검색 결과
type CIKSearchResult struct {
	Symbol           string `json:"symbol"`           // 종목 심볼
	CompanyName      string `json:"companyName"`      // 회사명
	CIK              string `json:"cik"`              // SEC CIK
	ExchangeFullName string `json:"exchangeFullName"` // 거래소 전체명
	Exchange         string `json:"exchange"`         // 거래소 코드
	Currency         string `json:"currency"`         // 통화
}

// CUSIPSearchResult — CUSIP 검색 결과
type CUSIPSearchResult struct {
	Symbol      string  `json:"symbol"`      // 종목 심볼
	CompanyName string  `json:"companyName"` // 회사명
	CUSIP       string  `json:"cusip"`       // CUSIP 코드
	MarketCap   float64 `json:"marketCap"`   // 시가총액
}

// ISINSearchResult — ISIN 검색 결과
type ISINSearchResult struct {
	Symbol    string `json:"symbol"`    // 종목 심볼
	Name      string `json:"name"`      // 회사명
	ISIN      string `json:"isin"`      // ISIN 코드
	MarketCap int64  `json:"marketCap"` // 시가총액
}

// SearchCIK 는 SEC CIK 로 종목을 검색한다.
func (c *Client) SearchCIK(ctx context.Context, cik string) ([]CIKSearchResult, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[CIKSearchResult](ctx, c.http, "/stable/search-cik", map[string]string{"cik": cik})
}

// SearchCUSIP 는 CUSIP 코드로 종목을 검색한다.
func (c *Client) SearchCUSIP(ctx context.Context, cusip string) ([]CUSIPSearchResult, error) {
	if strings.TrimSpace(cusip) == "" {
		return nil, fmt.Errorf("fmp: cusip must not be empty")
	}
	return fetch.List[CUSIPSearchResult](ctx, c.http, "/stable/search-cusip", map[string]string{"cusip": cusip})
}

// SearchISIN 는 ISIN 코드로 종목을 검색한다.
func (c *Client) SearchISIN(ctx context.Context, isin string) ([]ISINSearchResult, error) {
	if strings.TrimSpace(isin) == "" {
		return nil, fmt.Errorf("fmp: isin must not be empty")
	}
	return fetch.List[ISINSearchResult](ctx, c.http, "/stable/search-isin", map[string]string{"isin": isin})
}
```

- [ ] **Step 2: fixtures**

`search/testdata/search-cik.json`:
```json
[{ "symbol": "AAPL", "companyName": "Apple Inc.", "cik": "0000320193", "exchangeFullName": "NASDAQ Global Select", "exchange": "NASDAQ", "currency": "USD" }]
```
`search/testdata/search-cusip.json`:
```json
[{ "symbol": "AAPL.NE", "companyName": "Apple Inc.", "cusip": "037833100", "marketCap": 5156676087644.16 }]
```
`search/testdata/search-isin.json`:
```json
[{ "symbol": "AAPL", "name": "Apple Inc.", "isin": "US0378331005", "marketCap": 3900351299800 }]
```

- [ ] **Step 3: identifiers_test.go**

Create `search/identifiers_test.go`:
```go
package search

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSearchCIK_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-cik.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchCIK(context.Background(), "320193")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].CIK == "" || rows[0].Symbol != "AAPL" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/search-cik" || cap.query.Get("cik") != "320193" {
		t.Errorf("delegation: path=%q cik=%q", cap.path, cap.query.Get("cik"))
	}
	if _, err := c.SearchCIK(context.Background(), " "); err == nil {
		t.Fatal("want empty cik guard")
	}
}

func TestSearchCUSIP_ParsesFloatMarketCap(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-cusip.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchCUSIP(context.Background(), "037833100")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].CUSIP != "037833100" || rows[0].MarketCap <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestSearchISIN_ParsesIntMarketCap(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-isin.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchISIN(context.Background(), "US0378331005")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].ISIN != "US0378331005" || rows[0].MarketCap <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./search/ && go vet ./search/ && gofmt -l search/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add search/identifiers.go search/identifiers_test.go search/testdata/search-cik.json search/testdata/search-cusip.json search/testdata/search-isin.json
git commit -m "feat(search): SearchCIK/CUSIP/ISIN"
```

---

## Task 3: exchange-variants (TDD)

**Files:** Create `search/variants.go`, `search/variants_test.go`, `search/testdata/exchange-variants-aapl.json`.

- [ ] **Step 0: 카탈로그 ExchangeVariant 전체 필드 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
awk '/Response/,/출처/' docs/api/search/search-exchange-variants.md | grep -E '":'
```
응답 전체 필드를 확인해 아래 struct 에 누락 없이 반영(profile 유사 — 추가 필드 있으면 한국어 주석과 함께 추가). 가정과 다르면 status 에 보고.

- [ ] **Step 1: variants.go**

Create `search/variants.go` (Step 0 결과로 필드 보강):
```go
package search

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// ExchangeVariant — 거래소별 심볼 변형 (search-exchange-variants). profile 유사.
type ExchangeVariant struct {
	Symbol      string  `json:"symbol"`      // 종목 심볼
	Price       float64 `json:"price"`       // 현재가
	Beta        float64 `json:"beta"`        // 베타
	VolAvg      int64   `json:"volAvg"`      // 평균 거래량
	MktCap      int64   `json:"mktCap"`      // 시가총액
	LastDiv     float64 `json:"lastDiv"`     // 최근 배당
	Range       string  `json:"range"`       // 52주 범위
	Changes     float64 `json:"changes"`     // 등락액
	CompanyName string  `json:"companyName"` // 회사명
	Currency    string  `json:"currency"`    // 통화
	// Step 0 에서 추가 필드 확인 시 여기 보강
}

// SearchExchangeVariants 는 한 종목의 거래소별 변형(다른 거래소 상장)을 조회한다.
func (c *Client) SearchExchangeVariants(ctx context.Context, symbol string) ([]ExchangeVariant, error) {
	return fetch.ListBySymbol[ExchangeVariant](ctx, c.http, "/stable/search-exchange-variants", symbol)
}
```

- [ ] **Step 2: fixture**

`search/testdata/exchange-variants-aapl.json` (Step 0 의 실제 필드 반영, 최소 아래):
```json
[{ "symbol": "AAPL", "price": 262.82, "beta": 1.109, "volAvg": 47424558, "mktCap": 3900351299800, "lastDiv": 1.04, "range": "169.21-288.62", "changes": 3.24, "companyName": "Apple Inc.", "currency": "USD" }]
```

- [ ] **Step 3: variants_test.go**

Create `search/variants_test.go`:
```go
package search

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSearchExchangeVariants_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/exchange-variants-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchExchangeVariants(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].MktCap <= 0 || rows[0].CompanyName == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	// 빈 symbol 가드(ListBySymbol 내장)
	if _, err := c.SearchExchangeVariants(context.Background(), "  "); err == nil {
		t.Fatal("want empty symbol guard")
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./search/ && go vet ./search/ && gofmt -l search/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add search/variants.go search/variants_test.go search/testdata/exchange-variants-aapl.json
git commit -m "feat(search): SearchExchangeVariants"
```

---

## Task 4: screener — ScreenerParams + CompanyScreener (TDD)

**Files:** Create `search/screener.go`, `search/screener_test.go`, `search/testdata/screener.json`.

- [ ] **Step 1: screener.go**

Create `search/screener.go`:
```go
package search

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// ScreenerParams — company-screener 필터. 빈 값/0/nil 은 쿼리에서 생략.
// 숫자 0 = 미지정(생략). boolean 은 *bool(false 와 미지정 구분).
type ScreenerParams struct {
	MarketCapMoreThan      int64
	MarketCapLowerThan     int64
	Sector                 string
	Industry               string
	BetaMoreThan           float64
	BetaLowerThan          float64
	PriceMoreThan          float64
	PriceLowerThan         float64
	DividendMoreThan       float64
	DividendLowerThan      float64
	VolumeMoreThan         int64
	VolumeLowerThan        int64
	Exchange               string
	Country                string
	IsEtf                  *bool
	IsFund                 *bool
	IsActivelyTrading      *bool
	Limit                  int
	IncludeAllShareClasses *bool
}

// toMap 은 비제로/non-nil 필터만 쿼리 맵으로 만든다.
func (p ScreenerParams) toMap() map[string]string {
	m := map[string]string{}
	putI := func(k string, v int64) {
		if v != 0 {
			m[k] = strconv.FormatInt(v, 10)
		}
	}
	putF := func(k string, v float64) {
		if v != 0 {
			m[k] = strconv.FormatFloat(v, 'f', -1, 64)
		}
	}
	putS := func(k, v string) {
		if v != "" {
			m[k] = v
		}
	}
	putB := func(k string, v *bool) {
		if v != nil {
			m[k] = strconv.FormatBool(*v)
		}
	}
	putI("marketCapMoreThan", p.MarketCapMoreThan)
	putI("marketCapLowerThan", p.MarketCapLowerThan)
	putS("sector", p.Sector)
	putS("industry", p.Industry)
	putF("betaMoreThan", p.BetaMoreThan)
	putF("betaLowerThan", p.BetaLowerThan)
	putF("priceMoreThan", p.PriceMoreThan)
	putF("priceLowerThan", p.PriceLowerThan)
	putF("dividendMoreThan", p.DividendMoreThan)
	putF("dividendLowerThan", p.DividendLowerThan)
	putI("volumeMoreThan", p.VolumeMoreThan)
	putI("volumeLowerThan", p.VolumeLowerThan)
	putS("exchange", p.Exchange)
	putS("country", p.Country)
	putB("isEtf", p.IsEtf)
	putB("isFund", p.IsFund)
	putB("isActivelyTrading", p.IsActivelyTrading)
	if p.Limit > 0 {
		m["limit"] = strconv.Itoa(p.Limit)
	}
	putB("includeAllShareClasses", p.IncludeAllShareClasses)
	return m
}

// ScreenerResult — 스크리너 결과. nullable(MarketCap/Beta/LastAnnualDividend) → 포인터.
type ScreenerResult struct {
	Symbol             string   `json:"symbol"`             // 종목 심볼
	CompanyName        string   `json:"companyName"`        // 회사명
	MarketCap          *int64   `json:"marketCap"`          // 시가총액(결측 가능)
	Sector             string   `json:"sector"`             // 섹터
	Industry           string   `json:"industry"`           // 산업
	Beta               *float64 `json:"beta"`               // 베타(결측 가능)
	Price              float64  `json:"price"`              // 현재가
	LastAnnualDividend *float64 `json:"lastAnnualDividend"` // 최근 연간 배당(결측 가능)
	Volume             int64    `json:"volume"`             // 거래량
	Exchange           string   `json:"exchange"`           // 거래소
	ExchangeShortName  string   `json:"exchangeShortName"`  // 거래소 약칭
	Country            string   `json:"country"`            // 국가
	IsEtf              bool     `json:"isEtf"`              // ETF 여부
	IsFund             bool     `json:"isFund"`             // 펀드 여부
	IsActivelyTrading  bool     `json:"isActivelyTrading"`  // 거래 활성 여부
}

// CompanyScreener 는 다중 필터로 종목을 스크리닝한다. 빈 params 는 전체 스크리닝.
func (c *Client) CompanyScreener(ctx context.Context, params ScreenerParams) ([]ScreenerResult, error) {
	return fetch.List[ScreenerResult](ctx, c.http, "/stable/company-screener", params.toMap())
}
```

- [ ] **Step 2: fixture**

`search/testdata/screener.json` (nullable null/값 혼합):
```json
[
  {
    "symbol": "WIMA", "companyName": "WisdomTree International Adaptive Moving Average Fund",
    "marketCap": null, "sector": "Financial Services", "industry": "Asset Management",
    "beta": null, "price": 41.0956, "lastAnnualDividend": null, "volume": 2979,
    "exchange": "NASDAQ Global Market", "exchangeShortName": "NASDAQ", "country": "US",
    "isEtf": false, "isFund": true, "isActivelyTrading": true
  },
  {
    "symbol": "AAPL", "companyName": "Apple Inc.",
    "marketCap": 3900351299800, "sector": "Technology", "industry": "Consumer Electronics",
    "beta": 1.109, "price": 262.82, "lastAnnualDividend": 1.04, "volume": 47424558,
    "exchange": "NASDAQ Global Select", "exchangeShortName": "NASDAQ", "country": "US",
    "isEtf": false, "isFund": false, "isActivelyTrading": true
  }
]
```

- [ ] **Step 3: screener_test.go**

Create `search/screener_test.go`:
```go
package search

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestScreenerParams_ToMap(t *testing.T) {
	f := false
	tr := true
	p := ScreenerParams{
		MarketCapMoreThan: 1000000,
		Sector:            "Technology",
		PriceMoreThan:     10.5,
		IsEtf:             &f,
		IsActivelyTrading: &tr,
		Limit:             50,
	}
	m := p.toMap()
	if m["marketCapMoreThan"] != "1000000" {
		t.Errorf("marketCapMoreThan=%q", m["marketCapMoreThan"])
	}
	if m["sector"] != "Technology" {
		t.Errorf("sector=%q", m["sector"])
	}
	if m["priceMoreThan"] != "10.5" {
		t.Errorf("priceMoreThan=%q", m["priceMoreThan"])
	}
	if m["isEtf"] != "false" {
		t.Errorf("isEtf=%q want false (포인터 false 구분)", m["isEtf"])
	}
	if m["isActivelyTrading"] != "true" {
		t.Errorf("isActivelyTrading=%q", m["isActivelyTrading"])
	}
	if m["limit"] != "50" {
		t.Errorf("limit=%q", m["limit"])
	}
	// 미지정 필드는 키 없음
	if _, ok := m["industry"]; ok {
		t.Error("industry should be omitted")
	}
	if _, ok := m["isFund"]; ok {
		t.Error("isFund(nil) should be omitted")
	}
}

func TestScreenerParams_ToMap_Empty(t *testing.T) {
	if len(ScreenerParams{}.toMap()) != 0 {
		t.Error("empty params should yield empty map")
	}
}

func TestCompanyScreener_ParsesNullableAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/screener.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	tr := true
	rows, err := c.CompanyScreener(context.Background(), ScreenerParams{Sector: "Technology", IsActivelyTrading: &tr, Limit: 5})
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	// row0: nullable null → nil
	if rows[0].MarketCap != nil || rows[0].Beta != nil || rows[0].LastAnnualDividend != nil {
		t.Errorf("row0 nullables should be nil: %+v", rows[0])
	}
	// row1: nullable 값 존재
	if rows[1].MarketCap == nil || *rows[1].MarketCap <= 0 {
		t.Errorf("row1 MarketCap should be set")
	}
	if rows[1].Beta == nil || rows[1].LastAnnualDividend == nil {
		t.Errorf("row1 Beta/Dividend should be set")
	}
	// delegation
	if cap.path != "/stable/company-screener" || cap.query.Get("sector") != "Technology" || cap.query.Get("limit") != "5" {
		t.Errorf("delegation: path=%q sector=%q limit=%q", cap.path, cap.query.Get("sector"), cap.query.Get("limit"))
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./search/ && go vet ./search/ && gofmt -l search/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add search/screener.go search/screener_test.go search/testdata/screener.json
git commit -m "feat(search): CompanyScreener + ScreenerParams(19 필터)"
```

---

## Task 5: 루트 와이어 + README + examples + 통합 + 검증 (TDD)

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/search/main.go`.

- [ ] **Step 1: client_test.go 실패 테스트**

`client_test.go` 에 추가:
```go
func TestNewClient_HasSearch(t *testing.T) {
	c, err := NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Search == nil {
		t.Fatal("Search sub-client is nil")
	}
}
```

- [ ] **Step 2: client.go 와이어** — import `"github.com/kenshin579/fmp-go/search"`; `Client` 에 `Search *search.Client // 검색(심볼/식별자/스크리너)` 필드; `NewClient` 에 `c.Search = search.New(hc)`.

- [ ] **Step 3: 통과** — `go build ./... && go test . -run TestNewClient_HasSearch`. PASS.

- [ ] **Step 4: README + examples**

README 커버리지 표에 행 추가:
```markdown
| Search | `client.Search` | SearchSymbol, SearchName, SearchCIK, SearchCUSIP, SearchISIN, SearchExchangeVariants, CompanyScreener — 7 endpoint |
```

Create `examples/search/main.go`:
```go
// 실행: FMP_API_KEY=... go run examples/search/main.go
package main

import (
	"context"
	"fmt"
	"log"

	fmp "github.com/kenshin579/fmp-go"
	"github.com/kenshin579/fmp-go/search"
)

func main() {
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	hits, err := c.Search.SearchSymbol(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	for _, h := range hits {
		fmt.Printf("%s — %s (%s)\n", h.Symbol, h.Name, h.Exchange)
	}

	tech, err := c.Search.CompanyScreener(ctx, search.ScreenerParams{Sector: "Technology", Limit: 5})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tech screener: %d개\n", len(tech))
}
```

- [ ] **Step 5: integration_test.go 에 search 케이스**

`integration_test.go` 에 추가:
```go
func TestIntegration_Search(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Search.SearchSymbol(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("SearchSymbol: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Search.SearchName(ctx, "Apple"); err != nil || len(rows) == 0 {
		t.Errorf("SearchName: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Search.CompanyScreener(ctx, search.ScreenerParams{Sector: "Technology", Limit: 5}); err != nil || len(rows) == 0 {
		t.Errorf("CompanyScreener: err=%v len=%d", err, len(rows))
	}
}
```
(integration_test.go 에 `"github.com/kenshin579/fmp-go/search"` import 추가.)

- [ ] **Step 6: 전체 검증**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./...
go test -tags integration ./... -run TestIntegration_Search -v 2>&1 | tail -10
gofmt -l .
```
Expected: 단위 전체 PASS, gofmt clean. 통합 key 없으면 skip.

- [ ] **Step 7: Commit**
```bash
git add client.go client_test.go README.md examples/search/main.go integration_test.go
git commit -m "feat(search): 루트 Client 와이어 + README + examples + 통합"
```

---

## 자기 점검 메모 (작성자용)
- 검색 전부 `[]T`(단일 *T 없음). query/cik/cusip/isin 빈 가드 inline + `fetch.List`. exchange-variants 는 `fetch.ListBySymbol`(symbol 가드 내장).
- struct 공유: SymbolSearchResult(symbol/name 검색).
- marketCap 타입: CUSIP float64 / ISIN int64 (확인됨). 혼동 주의.
- ScreenerParams: 숫자 0=생략, boolean `*bool`(false/미지정 구분), Limit>0 만 포함. toMap 단위테스트로 검증.
- ScreenerResult nullable(MarketCap/Beta/LastAnnualDividend) 포인터 — null→nil fixture 검증.
- test helper(newTestClient/newCapturingClient/capturedReq)는 Task 1 search_test.go 정의 → Task 2-4 재사용(동일 패키지).
- ExchangeVariant 정확 필드는 Task 3 Step 0 카탈로그로 확정.
- 와이어: `Search *search.Client` + `c.Search = search.New(hc)`.
