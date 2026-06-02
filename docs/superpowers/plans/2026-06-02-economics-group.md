# Economics 그룹 (v0.15.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `economics/` 패키지 4 endpoint, 4 구조체.

**Architecture:** internal/fetch(List) + fromToParams helper. 구조체는 스펙 verbatim. EconomicCalendarEvent 의 Estimate/Unit 는 nullable 포인터.

**Tech Stack:** Go 1.25 generics, internal/fetch.

참고: `unset GOROOT`. 커밋 한국어 `feat(economics): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 + 4 endpoint

**Files:** Create `economics/client.go`, `economics/economics.go`, `economics/economics_test.go`, testdata `treasury-rates.json`, `economic-indicators.json`, `economic-calendar.json`, `market-risk-premium.json`.

- [ ] **Step 1:** `economics/client.go` — Client + New + helper:
```go
// Package economics 는 FMP 경제 API sub-client (국채/지표/캘린더/리스크프리미엄).
package economics

import "github.com/kenshin579/fmp-go/internal/httpclient"

type Client struct{ http *httpclient.Client }

func New(http *httpclient.Client) *Client { return &Client{http: http} }

func fromToParams(from, to string) map[string]string {
	q := map[string]string{}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return q
}
```
- [ ] **Step 2:** `economics/economics.go` — 스펙의 4 struct(TreasuryRate/EconomicIndicator/EconomicCalendarEvent/RiskPremium) + 4 메서드. import: context, fmt, strings, internal/fetch.
```go
func (c *Client) TreasuryRates(ctx context.Context, from, to string) ([]TreasuryRate, error) {
	return fetch.List[TreasuryRate](ctx, c.http, "/stable/treasury-rates", fromToParams(from, to))
}
func (c *Client) EconomicIndicators(ctx context.Context, name, from, to string) ([]EconomicIndicator, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	q := fromToParams(from, to)
	q["name"] = name
	return fetch.List[EconomicIndicator](ctx, c.http, "/stable/economic-indicators", q)
}
func (c *Client) EconomicCalendar(ctx context.Context, country, from, to string) ([]EconomicCalendarEvent, error) {
	q := fromToParams(from, to)
	if country != "" {
		q["country"] = country
	}
	return fetch.List[EconomicCalendarEvent](ctx, c.http, "/stable/economic-calendar", q)
}
func (c *Client) MarketRiskPremium(ctx context.Context) ([]RiskPremium, error) {
	return fetch.List[RiskPremium](ctx, c.http, "/stable/market-risk-premium", nil)
}
```
- [ ] **Step 3:** fixtures(1건 배열). treasury-rates {date, month1..year30 전부}, economic-indicators `[{name:"GDP",date,value}]`, economic-calendar `[{date:"2026-04-08 23:50:00",country:"US",event,currency:"USD",previous:1.2,estimate:null,actual:1.3,change:0.1,impact:"Low",changePercentage:8.3,unit:null}]`, market-risk-premium `[{country:"United States",continent:"North America",countryRiskPremium:0,totalEquityRiskPremium:4.6}]`.
- [ ] **Step 4:** `economics_test.go`(헬퍼 calendar 패턴 정의): TreasuryRates 파싱(Year10!=0) + delegation(path+from/to) / EconomicIndicators("CPI","","") delegation(path+name) + 빈 name 가드 / EconomicCalendar 파싱(Estimate==nil, Unit==nil, Previous!=0) + delegation(path+country) / MarketRiskPremium 파싱(TotalEquityRiskPremium!=0) + path.
- [ ] **Step 5:** `unset GOROOT && go test ./economics/ && go vet ./economics/ && gofmt -l economics/`. 커밋 `feat(economics): 국채/지표/캘린더/리스크프리미엄 4종`.

### Task 2: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/economics/main.go`.

- [ ] **Step 1:** `client.go` — import `economics`, struct 에 `Economics *economics.Client`, NewClient 에 `c.Economics = economics.New(hc)`.
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasEconomics`.
- [ ] **Step 3:** README 표 행 신규: `| Economics | \`client.Economics\` | TreasuryRates, EconomicIndicators, EconomicCalendar, MarketRiskPremium — 4 endpoint |`.
- [ ] **Step 4:** `examples/economics/main.go` — NewClientFromEnv → TreasuryRates("","") 첫 행 10년물 출력 + MarketRiskPremium 건수.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Economics`: TreasuryRates("","") len>0 / EconomicIndicators("GDP","","") len>0 / MarketRiskPremium len>0.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(economics): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 4 endpoint 1 패키지. path 단수 economic-indicators/economic-calendar 주의.
- EconomicCalendarEvent Estimate/Unit nullable. EconomicIndicators name 가드. MarketRiskPremium nil params.
