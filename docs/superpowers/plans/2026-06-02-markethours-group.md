# Market Hours 그룹 (v0.16.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `markethours/` 패키지 3 endpoint, 2 구조체.

**Architecture:** internal/fetch(List). ExchangeHours(2 메서드 공유) + ExchangeHoliday(adj 시각 nullable). 구조체는 스펙 verbatim.

참고: `unset GOROOT`. 커밋 한국어 `feat(markethours): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 + 3 endpoint

**Files:** Create `markethours/client.go`, `markethours/markethours.go`, `markethours/markethours_test.go`, testdata `exchange-market-hours.json`, `all-exchange-market-hours.json`, `holidays-by-exchange.json`.

- [ ] **Step 1:** `markethours/client.go`:
```go
// Package markethours 는 FMP 거래소 운영시간/휴장일 API sub-client.
// fmp.Client.MarketHours 로 접근.
package markethours

import "github.com/kenshin579/fmp-go/internal/httpclient"

type Client struct {
	http *httpclient.Client
}

func New(http *httpclient.Client) *Client { return &Client{http: http} }
```
- [ ] **Step 2:** `markethours/markethours.go` — 스펙의 `ExchangeHours`/`ExchangeHoliday` struct + 3 메서드. import: context, fmt, strings, internal/fetch.
```go
func (c *Client) ExchangeMarketHours(ctx context.Context, exchange string) ([]ExchangeHours, error) {
	if strings.TrimSpace(exchange) == "" {
		return nil, fmt.Errorf("fmp: exchange must not be empty")
	}
	return fetch.List[ExchangeHours](ctx, c.http, "/stable/exchange-market-hours", map[string]string{"exchange": exchange})
}
func (c *Client) AllExchangeMarketHours(ctx context.Context) ([]ExchangeHours, error) {
	return fetch.List[ExchangeHours](ctx, c.http, "/stable/all-exchange-market-hours", nil)
}
func (c *Client) HolidaysByExchange(ctx context.Context, exchange, from, to string) ([]ExchangeHoliday, error) {
	if strings.TrimSpace(exchange) == "" {
		return nil, fmt.Errorf("fmp: exchange must not be empty")
	}
	q := map[string]string{"exchange": exchange}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return fetch.List[ExchangeHoliday](ctx, c.http, "/stable/holidays-by-exchange", q)
}
```
- [ ] **Step 3:** fixtures. exchange-market-hours `[{exchange:"NASDAQ",name:"NASDAQ Global Market",openingHour:"09:30 AM -04:00",closingHour:"04:00 PM -04:00",timezone:"America/New_York",isMarketOpen:true}]`, all-exchange-market-hours(같은 shape 2건), holidays-by-exchange `[{exchange:"NASDAQ",date:"2026-04-03",name:"Good Friday",isClosed:true,adjOpenTime:null,adjCloseTime:null}]`.
- [ ] **Step 4:** `markethours_test.go`(헬퍼 calendar 패턴 정의): ExchangeMarketHours("NASDAQ") 파싱(IsMarketOpen==true, Timezone!="") + delegation(path+exchange) + 빈 exchange 가드 / AllExchangeMarketHours 파싱(len>=1) + path / HolidaysByExchange("NASDAQ","2025-01-01","2026-01-01") 파싱(IsClosed==true, AdjOpenTime==nil, AdjCloseTime==nil) + delegation(path+exchange+from+to) + 빈 exchange 가드.
- [ ] **Step 5:** `unset GOROOT && go test ./markethours/ && go vet ./markethours/ && gofmt -l markethours/`. 커밋 `feat(markethours): 거래소 운영시간/휴장일 3종`.

### Task 2: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/markethours/main.go`.

- [ ] **Step 1:** `client.go` — import `markethours`, struct 에 `MarketHours *markethours.Client`, NewClient 에 `c.MarketHours = markethours.New(hc)`.
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasMarketHours`.
- [ ] **Step 3:** README 표 행 신규: `| Market Hours | \`client.MarketHours\` | ExchangeMarketHours, AllExchangeMarketHours, HolidaysByExchange — 3 endpoint |`.
- [ ] **Step 4:** `examples/markethours/main.go` — NewClientFromEnv → ExchangeMarketHours("NASDAQ") 개장 여부 출력 + AllExchangeMarketHours 건수.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_MarketHours`: AllExchangeMarketHours len>0 / ExchangeMarketHours("NASDAQ") len>0 & rows[0].Exchange!="" / HolidaysByExchange("NASDAQ","","") err 만 체크.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(markethours): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- ExchangeHours 2 메서드 공유. ExchangeHoliday adj 시각 nullable.
- ExchangeMarketHours/HolidaysByExchange exchange 가드, AllExchangeMarketHours nil.
