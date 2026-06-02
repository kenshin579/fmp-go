# FMP Go SDK — News 그룹 (v0.6.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** FMP `news` 카테고리 10 endpoint 를 신규 `news/` 패키지로 추가하고 v0.6.0 준비.

**Architecture:** `internal/fetch` 사용. 9 endpoint 가 `Article` shape 공유, fmp-articles 만 `FMPArticle`. latest 6 = `(page int)` + List{page}, search 4 = `(symbols ...string)` + ListBySymbols. 전부 `[]T`.

**Tech Stack:** Go 1.25 / `internal/fetch` / fixture 단위테스트 + build-tag 통합. `unset GOROOT` 필요시.

**Spec:** `docs/superpowers/specs/2026-06-02-news-group-design.md`
**Repo / Branch:** `github.com/kenshin579/fmp-go`, branch `feature/news-group` (spec 커밋 이미).

**확정된 사실:**
- Article shape `{symbol, publishedDate, publisher, title, image, site, text, url}` — stock/crypto/forex/press-releases/general + 각 search(9개) 공용. general 은 `symbol: null`(→ "").
- FMPArticle shape `{title, date, content, tickers, image, link, author, site}`.
- path: latest `/stable/news/{stock,crypto,forex,general,press-releases}-latest?page=`, fmp `/stable/fmp-articles?page=`, search `/stable/news/{stock,crypto,forex,press-releases}?symbols=`.
- `internal/fetch`: `List[T](ctx, hc, path, params)`, `ListBySymbols[T](ctx, hc, path, symbols []string)`.
- root client 와이어 패턴: `Search *search.Client` 필드 + `c.Search = search.New(hc)`.

---

## File Structure
- Create: `news/client.go` — Client + New.
- Create: `news/article.go` — Article struct.
- Create: `news/latest.go` + `_test.go` + testdata — 5 latest 메서드.
- Create: `news/search.go` + `_test.go` + testdata — 4 search 메서드.
- Create: `news/fmp_articles.go` + `_test.go` + testdata — FMPArticle + FMPArticles.
- Modify: `client.go` / `client_test.go` — News 와이어.
- Modify: `README.md`; Create `examples/news/main.go`; Modify `integration_test.go`.

---

## Task 1: news 패키지 기반 + Article + latest 5개 (TDD)

**Files:** Create `news/client.go`, `news/article.go`, `news/latest.go`, `news/latest_test.go`, `news/testdata/stock-news-latest.json`, `news/testdata/general-news-latest.json`.

- [ ] **Step 1: client.go**
```go
// Package news 는 FMP 뉴스(news) API sub-client.
// fmp.Client.News 로 접근.
package news

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 뉴스 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
```

- [ ] **Step 2: article.go**
```go
package news

// Article — 뉴스/보도자료 기사 (stock/crypto/forex/press-releases/general + 각 search 공용)
type Article struct {
	Symbol        string `json:"symbol"`        // 관련 종목 심볼 (general 뉴스는 빈 문자열)
	PublishedDate string `json:"publishedDate"` // 게시 일시 (YYYY-MM-DD HH:MM:SS)
	Publisher     string `json:"publisher"`     // 발행처 (예: Seeking Alpha)
	Title         string `json:"title"`         // 기사 제목
	Image         string `json:"image"`         // 대표 이미지 URL
	Site          string `json:"site"`          // 출처 사이트 도메인
	Text          string `json:"text"`          // 기사 본문 요약
	URL           string `json:"url"`           // 원문 URL
}
```

- [ ] **Step 3: latest.go**
```go
package news

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

func pageParams(page int) map[string]string {
	return map[string]string{"page": strconv.Itoa(page)}
}

// StockNewsLatest 는 최신 주식 뉴스를 페이지 단위로 조회한다.
func (c *Client) StockNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/stock-latest", pageParams(page))
}

// CryptoNewsLatest 는 최신 암호화폐 뉴스를 조회한다.
func (c *Client) CryptoNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/crypto-latest", pageParams(page))
}

// ForexNewsLatest 는 최신 외환 뉴스를 조회한다.
func (c *Client) ForexNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/forex-latest", pageParams(page))
}

// GeneralNewsLatest 는 최신 일반 경제 뉴스를 조회한다.
func (c *Client) GeneralNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/general-latest", pageParams(page))
}

// PressReleasesLatest 는 최신 보도자료를 조회한다.
func (c *Client) PressReleasesLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/press-releases-latest", pageParams(page))
}
```

- [ ] **Step 4: fixtures**

`news/testdata/stock-news-latest.json`:
```json
[
  {
    "symbol": "INSG",
    "publishedDate": "2025-02-03 23:53:40",
    "publisher": "Seeking Alpha",
    "title": "Q4 Earnings Release Looms For Inseego",
    "image": "https://images.financialmodelingprep.com/news/x.jpg",
    "site": "seekingalpha.com",
    "text": "Inseego's Q3 beat was largely due to a one-time gain.",
    "url": "https://seekingalpha.com/article/4754485"
  }
]
```
`news/testdata/general-news-latest.json` (symbol null):
```json
[
  {
    "symbol": null,
    "publishedDate": "2025-02-03 23:51:37",
    "publisher": "CNBC",
    "title": "Asia tech stocks rise after Trump pauses tariffs",
    "image": "https://images.financialmodelingprep.com/news/y.jpg",
    "site": "cnbc.com",
    "text": "Gains in Asian tech companies were broad-based.",
    "url": "https://www.cnbc.com/2025/02/04/asia-tech.html"
  }
]
```

- [ ] **Step 5: latest_test.go (test helper 포함)**
```go
package news

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

func TestStockNewsLatest_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/stock-news-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.StockNewsLatest(context.Background(), 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "INSG" || rows[0].Title == "" || rows[0].URL == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/news/stock-latest" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}

func TestGeneralNewsLatest_NullSymbolDecodes(t *testing.T) {
	raw, _ := os.ReadFile("testdata/general-news-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.GeneralNewsLatest(context.Background(), 1)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "" {
		t.Errorf("null symbol should decode to empty string, got %q", rows[0].Symbol)
	}
	if rows[0].Publisher != "CNBC" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestCryptoForexPressLatest_Delegate(t *testing.T) {
	cases := []struct {
		name string
		call func(c *Client) ([]Article, error)
		path string
	}{
		{"crypto", func(c *Client) ([]Article, error) { return c.CryptoNewsLatest(context.Background(), 0) }, "/stable/news/crypto-latest"},
		{"forex", func(c *Client) ([]Article, error) { return c.ForexNewsLatest(context.Background(), 0) }, "/stable/news/forex-latest"},
		{"press", func(c *Client) ([]Article, error) { return c.PressReleasesLatest(context.Background(), 0) }, "/stable/news/press-releases-latest"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, cap, cleanup := newCapturingClient(t, `[]`)
			defer cleanup()
			if _, err := tc.call(c); err != nil {
				t.Fatalf("%s: %v", tc.name, err)
			}
			if cap.path != tc.path {
				t.Errorf("path=%q want %q", cap.path, tc.path)
			}
		})
	}
}
```

- [ ] **Step 6: 통과 확인**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./news/ -v && go vet ./news/ && gofmt -l news/
```
Expected: PASS, vet clean, gofmt 빈 출력.

- [ ] **Step 7: Commit**
```bash
git add news/client.go news/article.go news/latest.go news/latest_test.go news/testdata/stock-news-latest.json news/testdata/general-news-latest.json
git commit -m "feat(news): 패키지 기반 + Article + latest 5개"
```

---

## Task 2: search 뉴스 4개 (TDD)

**Files:** Create `news/search.go`, `news/search_test.go`, `news/testdata/search-stock-news.json`.

- [ ] **Step 1: search.go**
```go
package news

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SearchStockNews 는 종목별 주식 뉴스를 조회한다.
func (c *Client) SearchStockNews(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/stock", symbols)
}

// SearchCryptoNews 는 종목별 암호화폐 뉴스를 조회한다.
func (c *Client) SearchCryptoNews(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/crypto", symbols)
}

// SearchForexNews 는 통화쌍별 외환 뉴스를 조회한다.
func (c *Client) SearchForexNews(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/forex", symbols)
}

// SearchPressReleases 는 종목별 보도자료를 조회한다.
func (c *Client) SearchPressReleases(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/press-releases", symbols)
}
```

- [ ] **Step 2: fixture**

`news/testdata/search-stock-news.json`:
```json
[
  {
    "symbol": "AAPL",
    "publishedDate": "2025-02-03 21:05:14",
    "publisher": "Zacks Investment Research",
    "title": "Apple & China Tariffs: A Closer Look",
    "image": "https://images.financialmodelingprep.com/news/z.jpg",
    "site": "zacks.com",
    "text": "Tariffs have been the talk of the town.",
    "url": "https://www.zacks.com/stock/news/2408814"
  }
]
```

- [ ] **Step 3: search_test.go**
```go
package news

import (
	"context"
	"os"
	"testing"
)

func TestSearchStockNews_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-stock-news.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchStockNews(context.Background(), "AAPL", "MSFT")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/news/stock" || cap.query.Get("symbols") != "AAPL,MSFT" {
		t.Errorf("delegation: path=%q symbols=%q", cap.path, cap.query.Get("symbols"))
	}
}

func TestSearchNews_EmptySymbolsGuard(t *testing.T) {
	c, cleanup := newTestClient(t, 200, `[]`)
	defer cleanup()
	if _, err := c.SearchStockNews(context.Background()); err == nil {
		t.Fatal("want empty symbols guard")
	}
}

func TestSearchCryptoForexPress_Delegate(t *testing.T) {
	cases := []struct {
		name string
		call func(c *Client) ([]Article, error)
		path string
	}{
		{"crypto", func(c *Client) ([]Article, error) { return c.SearchCryptoNews(context.Background(), "BTCUSD") }, "/stable/news/crypto"},
		{"forex", func(c *Client) ([]Article, error) { return c.SearchForexNews(context.Background(), "EURUSD") }, "/stable/news/forex"},
		{"press", func(c *Client) ([]Article, error) { return c.SearchPressReleases(context.Background(), "AAPL") }, "/stable/news/press-releases"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, cap, cleanup := newCapturingClient(t, `[]`)
			defer cleanup()
			if _, err := tc.call(c); err != nil {
				t.Fatalf("%s: %v", tc.name, err)
			}
			if cap.path != tc.path {
				t.Errorf("path=%q want %q", cap.path, tc.path)
			}
		})
	}
}
```

- [ ] **Step 4: 통과 확인** — `go test ./news/ && go vet ./news/ && gofmt -l news/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add news/search.go news/search_test.go news/testdata/search-stock-news.json
git commit -m "feat(news): SearchStockNews/CryptoNews/ForexNews/PressReleases"
```

---

## Task 3: fmp-articles (TDD)

**Files:** Create `news/fmp_articles.go`, `news/fmp_articles_test.go`, `news/testdata/fmp-articles.json`.

- [ ] **Step 1: fmp_articles.go**
```go
package news

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// FMPArticle — FMP 자체 작성 기사 (fmp-articles)
type FMPArticle struct {
	Title   string `json:"title"`   // 제목
	Date    string `json:"date"`    // 작성 일시
	Content string `json:"content"` // 본문 (HTML)
	Tickers string `json:"tickers"` // 관련 티커 (예: "NYSE:MRK")
	Image   string `json:"image"`   // 이미지 URL
	Link    string `json:"link"`    // FMP 기사 링크
	Author  string `json:"author"`  // 작성자
	Site    string `json:"site"`    // 출처 (Financial Modeling Prep)
}

// FMPArticles 는 FMP 자체 작성 기사를 페이지 단위로 조회한다.
func (c *Client) FMPArticles(ctx context.Context, page int) ([]FMPArticle, error) {
	return fetch.List[FMPArticle](ctx, c.http, "/stable/fmp-articles", map[string]string{"page": strconv.Itoa(page)})
}
```

- [ ] **Step 2: fixture**

`news/testdata/fmp-articles.json`:
```json
[
  {
    "title": "Merck Shares Plunge 8% as Weak Guidance Overshadows Strong Revenue Growth",
    "date": "2025-02-04 09:33:00",
    "content": "<p>Merck & Co (NYSE:MRK) saw its stock sink over 8% in pre-market today.</p>",
    "tickers": "NYSE:MRK",
    "image": "https://cdn.financialmodelingprep.com/images/fmp-x.jpg",
    "link": "https://financialmodelingprep.com/market-news/fmp-merck-shares-plunge",
    "author": "Davit Kirakosyan",
    "site": "Financial Modeling Prep"
  }
]
```

- [ ] **Step 3: fmp_articles_test.go**
```go
package news

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestFMPArticles_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/fmp-articles.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.FMPArticles(context.Background(), 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Title == "" || rows[0].Tickers != "NYSE:MRK" || rows[0].Author == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/fmp-articles" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
	_ = http.StatusOK
}
```

- [ ] **Step 4: 통과 확인** — `go test ./news/ && go vet ./news/ && gofmt -l news/`. PASS.

- [ ] **Step 5: Commit**
```bash
git add news/fmp_articles.go news/fmp_articles_test.go news/testdata/fmp-articles.json
git commit -m "feat(news): FMPArticles"
```

---

## Task 4: 루트 와이어 + README + examples + 통합 + 검증 (TDD)

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/news/main.go`.

- [ ] **Step 1: client_test.go 실패 테스트**
```go
func TestNewClient_HasNews(t *testing.T) {
	c, err := NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.News == nil {
		t.Fatal("News sub-client is nil")
	}
}
```

- [ ] **Step 2: client.go 와이어** — import `"github.com/kenshin579/fmp-go/news"`; `Client` 에 `News *news.Client // 뉴스(주식/암호화폐/외환/보도자료/일반)` 필드; `NewClient` 에 `c.News = news.New(hc)`.

- [ ] **Step 3: 통과** — `go build ./... && go test . -run TestNewClient_HasNews`. PASS.

- [ ] **Step 4: README + examples**

README 커버리지 표 행 추가:
```markdown
| News | `client.News` | StockNewsLatest, CryptoNewsLatest, ForexNewsLatest, GeneralNewsLatest, PressReleasesLatest, Search(Stock/Crypto/Forex/PressReleases)News, FMPArticles — 10 endpoint |
```

Create `examples/news/main.go`:
```go
// 실행: FMP_API_KEY=... go run examples/news/main.go
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

	hits, err := c.News.SearchStockNews(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	for i, a := range hits {
		if i >= 3 {
			break
		}
		fmt.Printf("[%s] %s — %s\n", a.PublishedDate, a.Title, a.Site)
	}

	latest, err := c.News.StockNewsLatest(ctx, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("latest stock news: %d개\n", len(latest))
}
```

- [ ] **Step 5: integration_test.go 에 news 케이스**
```go
func TestIntegration_News(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.News.StockNewsLatest(ctx, 0); err != nil || len(rows) == 0 {
		t.Errorf("StockNewsLatest: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.News.SearchStockNews(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("SearchStockNews: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.News.FMPArticles(ctx, 0); err != nil || len(rows) == 0 {
		t.Errorf("FMPArticles: err=%v len=%d", err, len(rows))
	}
}
```

- [ ] **Step 6: 전체 검증**
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./...
go test -tags integration ./... -run TestIntegration_News -v 2>&1 | tail -10
gofmt -l .
```
Expected: 단위 전체 PASS, gofmt clean. 통합 key 없으면 skip.

- [ ] **Step 7: Commit**
```bash
git add client.go client_test.go README.md examples/news/main.go integration_test.go
git commit -m "feat(news): 루트 Client 와이어 + README + examples + 통합"
```

---

## 자기 점검 메모 (작성자용)
- 9 endpoint Article 공유, fmp-articles 만 FMPArticle. 전부 []T.
- latest 6 = `(page int)` + `fetch.List{page}`. search 4 = `(symbols ...string)` + `fetch.ListBySymbols`(가드 내장).
- `pageParams(page)` 헬퍼는 latest.go 에 정의 → fmp_articles.go 는 별도 inline(파일 분리 — 같은 패키지지만 단순 inline 유지).
- general-news `symbol: null` → string "" 디코딩 (Go encoding/json 안전). fixture 로 검증.
- test helper(newTestClient/newCapturingClient/capturedReq)는 Task 1 latest_test.go 정의 → Task 2-3 재사용(동일 패키지).
- 와이어: `News *news.Client` + `c.News = news.New(hc)`.
- 모든 struct 필드 한국어 주석.
