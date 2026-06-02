# Directory 그룹 (v0.14.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `directory/` 패키지 11 endpoint, 10 구조체.

**Architecture:** internal/fetch(List). 무파라미터 9 + CIKList(page/limit) + SymbolChangesList(invalid/limit). 구조체는 스펙 verbatim. SymbolName 은 etf/actively 공유.

**Tech Stack:** Go 1.25 generics, internal/fetch.

참고: `unset GOROOT`. 커밋 한국어 `feat(directory): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의. nil params 안전.

---

### Task 1: 패키지 scaffold + 심볼 목록 7종

**Files:** Create `directory/client.go`, `directory/symbols.go`, `directory/symbols_test.go`, testdata `stock-list.json`, `financial-symbol-list.json`, `cik-list.json`, `symbol-change.json`, `etf-list.json`, `actively-trading-list.json`, `earnings-transcript-list.json`.

- [ ] **Step 1:** `directory/client.go`:
```go
// Package directory 는 FMP 목록 API sub-client (심볼/거래소/섹터/산업/국가).
// fmp.Client.Directory 로 접근.
package directory

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 목록 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
```
- [ ] **Step 2:** `directory/symbols.go` — 스펙의 `SymbolName`/`CompanySymbol`/`FinancialSymbol`/`CIKEntry`/`SymbolChange`/`TranscriptSymbol` struct + 7 메서드. import: context, strconv, internal/fetch.
```go
func (c *Client) CompanySymbolsList(ctx context.Context) ([]CompanySymbol, error) {
	return fetch.List[CompanySymbol](ctx, c.http, "/stable/stock-list", nil)
}
func (c *Client) FinancialSymbolsList(ctx context.Context) ([]FinancialSymbol, error) {
	return fetch.List[FinancialSymbol](ctx, c.http, "/stable/financial-statement-symbol-list", nil)
}
func (c *Client) CIKList(ctx context.Context, page, limit int) ([]CIKEntry, error) {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[CIKEntry](ctx, c.http, "/stable/cik-list", q)
}
func (c *Client) SymbolChangesList(ctx context.Context, invalid bool, limit int) ([]SymbolChange, error) {
	q := map[string]string{"invalid": strconv.FormatBool(invalid)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[SymbolChange](ctx, c.http, "/stable/symbol-change", q)
}
func (c *Client) ETFsList(ctx context.Context) ([]SymbolName, error) {
	return fetch.List[SymbolName](ctx, c.http, "/stable/etf-list", nil)
}
func (c *Client) ActivelyTradingList(ctx context.Context) ([]SymbolName, error) {
	return fetch.List[SymbolName](ctx, c.http, "/stable/actively-trading-list", nil)
}
func (c *Client) EarningsTranscriptList(ctx context.Context) ([]TranscriptSymbol, error) {
	return fetch.List[TranscriptSymbol](ctx, c.http, "/stable/earnings-transcript-list", nil)
}
```
- [ ] **Step 3:** fixtures(각 1~2건). stock-list `[{symbol,companyName}]`, financial-symbol-list `[{symbol,companyName,tradingCurrency,reportingCurrency}]`, cik-list `[{cik:"0002036063",companyName}]`, symbol-change `[{date,companyName,oldSymbol,newSymbol}]`, etf-list `[{symbol,name}]`, actively-trading-list `[{symbol,name}]`, earnings-transcript-list `[{symbol,companyName,noOfTranscripts:"16"}]`.
- [ ] **Step 4:** `symbols_test.go`(헬퍼 calendar 패턴 정의): CompanySymbolsList 파싱(Symbol/CompanyName) + delegation path / CIKList(0,10) 파싱(CIK=="0002036063" 0-padded 유지) + delegation path+page/limit / SymbolChangesList(true,10) delegation path+invalid=="true"/limit / ETFsList & ActivelyTradingList SymbolName 파싱 + path 각각 / EarningsTranscriptList NoOfTranscripts=="16" 문자열.
- [ ] **Step 5:** `unset GOROOT && go test ./directory/ && go vet ./directory/ && gofmt -l directory/`. 커밋 `feat(directory): 심볼 목록 7종`.

### Task 2: available-* 4종

**Files:** Create `directory/available.go`, `directory/available_test.go`, testdata `available-exchanges.json`, `available-sectors.json`, `available-industries.json`, `available-countries.json`.

- [ ] **Step 1:** `available.go` — 스펙 `Exchange`/`Sector`/`Industry`/`Country` struct + 4 메서드(전부 nil params):
```go
func (c *Client) AvailableExchanges(ctx context.Context) ([]Exchange, error) {
	return fetch.List[Exchange](ctx, c.http, "/stable/available-exchanges", nil)
}
// AvailableSectors → /stable/available-sectors ([]Sector)
// AvailableIndustries → /stable/available-industries ([]Industry)
// AvailableCountries → /stable/available-countries ([]Country)
```
- [ ] **Step 2:** fixtures: available-exchanges `[{exchange:"AMEX",name,countryName,countryCode:"US",symbolSuffix:"N/A",delay:"Real-time"}]`, available-sectors `[{sector:"Basic Materials"}]`, available-industries `[{industry:"Steel"}]`, available-countries `[{country:"US"}]`.
- [ ] **Step 3:** `available_test.go`(헬퍼 재사용): AvailableExchanges 파싱(Exchange=="AMEX", CountryCode=="US", Delay!="") + delegation path / AvailableSectors 파싱(Sector!="") + path / AvailableIndustries(Industry!="") / AvailableCountries(Country!="").
- [ ] **Step 4:** `go test ./directory/ && go vet && gofmt -l`. 커밋 `feat(directory): available 거래소/섹터/산업/국가 4종`.

### Task 3: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/directory/main.go`.

- [ ] **Step 1:** `client.go` — import `directory`, struct 에 `Directory *directory.Client`, NewClient 에 `c.Directory = directory.New(hc)`.
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasDirectory`.
- [ ] **Step 3:** README 표 Directory 행 신규: `| Directory | \`client.Directory\` | CompanySymbolsList, FinancialSymbolsList, CIKList, SymbolChangesList, ETFsList, ActivelyTradingList, EarningsTranscriptList, AvailableExchanges, AvailableSectors, AvailableIndustries, AvailableCountries — 11 endpoint |`.
- [ ] **Step 4:** `examples/directory/main.go` — NewClientFromEnv → AvailableExchanges 건수 + AvailableSectors 목록 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Directory`: CompanySymbolsList len>0 / AvailableExchanges len>0 & rows[0].Exchange!="" / AvailableSectors len>0.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(directory): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 11 endpoint: 심볼 7=T1, available 4=T2, 와이어/문서=T3.
- SymbolName 공유(etf/actively). CIK 0-padded·noOfTranscripts 문자열 유지.
- CIKList/SymbolChangesList 만 파라미터, 나머지 nil.
