# Senate/House 그룹 (v0.20.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `senate/` 패키지 6 endpoint, 단일 CongressTrade 구조체.

**Architecture:** internal/fetch(List) + pageParams helper. CongressTrade(6 endpoint 공유, 15필드 string). 구조체는 스펙 verbatim.

참고: `unset GOROOT`. 커밋 한국어 `feat(senate): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 + 6 endpoint

**Files:** Create `senate/client.go`, `senate/senate.go`, `senate/senate_test.go`, testdata `senate-latest.json`, `senate-trades.json`, `house-latest.json`.

- [ ] **Step 1:** `senate/client.go`:
```go
// Package senate 는 FMP 의회(상원/하원) 거래 공시 API sub-client.
// fmp.Client.Senate 로 접근.
package senate

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

type Client struct {
	http *httpclient.Client
}

func New(http *httpclient.Client) *Client { return &Client{http: http} }

func pageParams(page, limit int) map[string]string {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
```
- [ ] **Step 2:** `senate/senate.go` — 스펙 `CongressTrade` struct + 6 메서드. import: context, fmt, strings, internal/fetch.
```go
func (c *Client) SenateLatest(ctx context.Context, page, limit int) ([]CongressTrade, error) {
	return fetch.List[CongressTrade](ctx, c.http, "/stable/senate-latest", pageParams(page, limit))
}
func (c *Client) SenateTrades(ctx context.Context, symbol string, page, limit int) ([]CongressTrade, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := pageParams(page, limit)
	q["symbol"] = symbol
	return fetch.List[CongressTrade](ctx, c.http, "/stable/senate-trades", q)
}
func (c *Client) SenateTradesByName(ctx context.Context, name string) ([]CongressTrade, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[CongressTrade](ctx, c.http, "/stable/senate-trades-by-name", map[string]string{"name": name})
}
func (c *Client) HouseLatest(ctx context.Context, page, limit int) ([]CongressTrade, error) {
	return fetch.List[CongressTrade](ctx, c.http, "/stable/house-latest", pageParams(page, limit))
}
func (c *Client) HouseTrades(ctx context.Context, symbol string, page, limit int) ([]CongressTrade, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := pageParams(page, limit)
	q["symbol"] = symbol
	return fetch.List[CongressTrade](ctx, c.http, "/stable/house-trades", q)
}
func (c *Client) HouseTradesByName(ctx context.Context, name string) ([]CongressTrade, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[CongressTrade](ctx, c.http, "/stable/house-trades-by-name", map[string]string{"name": name})
}
```
- [ ] **Step 3:** fixtures: senate-trades.json `[{15필드 전부, amount:"$1,001 - $15,000", capitalGainsOver200USD:"False", type:"Purchase", symbol:"AAPL"}]`, senate-latest.json(capitalGainsOver200USD 키 없이 14필드), house-latest.json(district:"CA31" 등 하원형).
- [ ] **Step 4:** `senate_test.go`(헬퍼 calendar 패턴 정의): SenateTrades 파싱(Amount=="$1,001 - $15,000", CapitalGainsOver200USD!="", Symbol=="AAPL") / SenateLatest 파싱(capitalGainsOver200USD 키 없으면 빈 문자열 — len>0, Symbol!="") / delegation: SenateLatest(0,10) path `/stable/senate-latest`+page/limit, SenateTrades("AAPL",0,5) path+symbol, SenateTradesByName("Moran") path+name, HouseLatest path, HouseTrades("AAPL",0,5) path, HouseTradesByName("Pelosi") path / 가드 SenateTrades 빈 symbol + SenateTradesByName 빈 name.
- [ ] **Step 5:** `unset GOROOT && go test ./senate/ && go vet ./senate/ && gofmt -l senate/`. 커밋 `feat(senate): 상원/하원 거래 공시 6종`.

### Task 2: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/senate/main.go`.

- [ ] **Step 1:** `client.go` — import `senate`, struct 에 `Senate *senate.Client`, NewClient 에 `c.Senate = senate.New(hc)`. (gofmt -w client.go.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasSenate`.
- [ ] **Step 3:** README 표 행 신규: `| Senate/House | \`client.Senate\` | SenateLatest, SenateTrades, SenateTradesByName, HouseLatest, HouseTrades, HouseTradesByName — 6 endpoint |`.
- [ ] **Step 4:** `examples/senate/main.go` — NewClientFromEnv → SenateLatest(0,5) 첫 건 + HouseLatest(0,5) 건수 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Senate`: SenateLatest(0,10) len>0 / HouseLatest(0,10) len>0 / SenateTrades("AAPL",0,5) err 체크.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(senate): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 6 endpoint, 단일 CongressTrade. amount/capitalGains 문자열.
- 쿼리 3패턴(latest/trades/by-name). trades symbol 가드, by-name name 가드.
