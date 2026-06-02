# SEC Filings 그룹 (v0.27.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `secfilings/` 패키지 12 endpoint, 5 구조체.

**Architecture:** internal/fetch(List). 구조체는 스펙 verbatim. path family(sec-filings-search vs sec-filings-company-search) 주의. CompanySearchResult 5개 공유.

참고: `unset GOROOT`. 커밋 한국어 `feat(secfilings): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 + 필링 5종 (latest 2 + search 3)

**Files:** Create `secfilings/client.go`, `secfilings/filings.go`, `secfilings/filings_test.go`, testdata `financials-latest.json`, `8k-latest.json`, `search-symbol.json`.

- [ ] **Step 1:** `secfilings/client.go`:
```go
// Package secfilings 는 FMP SEC 공시 검색/분류/프로필 API sub-client.
// fmp.Client.SECFilings 로 접근.
package secfilings

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

type Client struct {
	http *httpclient.Client
}

func New(http *httpclient.Client) *Client { return &Client{http: http} }

// filingParams: 주param + from/to + page/limit.
func filingParams(key, val, from, to string, page, limit int) map[string]string {
	q := map[string]string{"from": from, "to": to, "page": strconv.Itoa(page)}
	if key != "" {
		q[key] = val
	}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
```
- [ ] **Step 2:** `secfilings/filings.go` — 스펙 `LatestFiling`/`FilingSearchResult` struct + 5 메서드. import: context, fmt, strings, internal/fetch.
```go
func (c *Client) LatestFinancials(ctx context.Context, from, to string, page, limit int) ([]LatestFiling, error) {
	if strings.TrimSpace(from) == "" || strings.TrimSpace(to) == "" {
		return nil, fmt.Errorf("fmp: from, to must not be empty")
	}
	return fetch.List[LatestFiling](ctx, c.http, "/stable/sec-filings-financials", filingParams("", "", from, to, page, limit))
}
func (c *Client) Latest8K(ctx context.Context, from, to string, page, limit int) ([]LatestFiling, error) {
	if strings.TrimSpace(from) == "" || strings.TrimSpace(to) == "" {
		return nil, fmt.Errorf("fmp: from, to must not be empty")
	}
	return fetch.List[LatestFiling](ctx, c.http, "/stable/sec-filings-8k", filingParams("", "", from, to, page, limit))
}
func (c *Client) SearchBySymbol(ctx context.Context, symbol, from, to string, page, limit int) ([]FilingSearchResult, error) {
	if strings.TrimSpace(symbol) == "" || strings.TrimSpace(from) == "" || strings.TrimSpace(to) == "" {
		return nil, fmt.Errorf("fmp: symbol, from, to must not be empty")
	}
	return fetch.List[FilingSearchResult](ctx, c.http, "/stable/sec-filings-search/symbol", filingParams("symbol", symbol, from, to, page, limit))
}
func (c *Client) SearchByCIK(ctx context.Context, cik, from, to string, page, limit int) ([]FilingSearchResult, error) {
	if strings.TrimSpace(cik) == "" || strings.TrimSpace(from) == "" || strings.TrimSpace(to) == "" {
		return nil, fmt.Errorf("fmp: cik, from, to must not be empty")
	}
	return fetch.List[FilingSearchResult](ctx, c.http, "/stable/sec-filings-search/cik", filingParams("cik", cik, from, to, page, limit))
}
func (c *Client) SearchByFormType(ctx context.Context, formType, from, to string, page, limit int) ([]FilingSearchResult, error) {
	if strings.TrimSpace(formType) == "" || strings.TrimSpace(from) == "" || strings.TrimSpace(to) == "" {
		return nil, fmt.Errorf("fmp: formType, from, to must not be empty")
	}
	return fetch.List[FilingSearchResult](ctx, c.http, "/stable/sec-filings-search/form-type", filingParams("formType", formType, from, to, page, limit))
}
```
- [ ] **Step 3:** fixtures: financials-latest.json `[{symbol,cik,filingDate:"2024-03-01 00:00:00",acceptedDate,formType:"10-K",hasFinancials:true,link,finalLink}]`, 8k-latest.json(hasFinancials:false), search-symbol.json(hasFinancials 키 없음).
- [ ] **Step 4:** `filings_test.go`(헬퍼 calendar 패턴 정의): LatestFinancials 파싱(HasFinancials==true) + delegation(path+from/to/page) + 빈 from 가드 / Latest8K 파싱(HasFinancials==false) + path / SearchBySymbol 파싱(FormType!="") + delegation(path `/stable/sec-filings-search/symbol`+symbol/from/to) + 빈 symbol 가드 / SearchByCIK path / SearchByFormType path.
- [ ] **Step 5:** `unset GOROOT && go test ./secfilings/ && go vet ./secfilings/ && gofmt -l secfilings/`. 커밋 `feat(secfilings): 최신 필링 2종 + 검색 3종`.

### Task 2: 회사검색 + 산업분류 6종

**Files:** Create `secfilings/company.go`, `secfilings/company_test.go`, testdata `company-search.json`, `industry-classification-list.json`, `industry-classification-search.json`.

- [ ] **Step 1:** `secfilings/company.go` — 스펙 `CompanySearchResult`/`IndustryClassification` struct + 6 메서드:
```go
// SearchByName(ctx, company) → company 가드 + {company}, /stable/sec-filings-company-search/name
// CompanySearchBySymbol(ctx, symbol) → symbol 가드 + {symbol}, /stable/sec-filings-company-search/symbol
// CompanySearchByCIK(ctx, cik) → cik 가드 + {cik}, /stable/sec-filings-company-search/cik
// IndustryClassificationList(ctx, industryTitle, sicCode) → 빈값 제외 맵, /stable/standard-industrial-classification-list ([]IndustryClassification)
// IndustryClassificationSearch(ctx, symbol, cik, sicCode) → 빈값 제외 맵, /stable/industry-classification-search ([]CompanySearchResult)
// AllIndustryClassification(ctx, page, limit) → {page,limit?}, /stable/all-industry-classification ([]CompanySearchResult)
```
import: context, fmt, strconv, strings, internal/fetch.
- [ ] **Step 2:** fixtures: company-search.json `[{symbol:"AAPL",name:"Apple Inc",cik,sicCode:"3571",industryTitle,businessAddress,phoneNumber}]`, industry-classification-list.json `[{office,sicCode:"100",industryTitle}]`, industry-classification-search.json(CompanySearchResult shape, businessAddress 리스트형 문자열).
- [ ] **Step 3:** `company_test.go`(헬퍼 재사용): SearchByName 파싱(Name!="", SICCode string) + delegation(path+company) + 빈 company 가드 / CompanySearchBySymbol delegation(path+symbol) / IndustryClassificationList 파싱(Office!="", SICCode=="100") + delegation(path+industryTitle) / IndustryClassificationSearch 파싱 + delegation(path) / AllIndustryClassification delegation(path+page/limit).
- [ ] **Step 4:** `go test ./secfilings/ && go vet && gofmt -l`. 커밋 `feat(secfilings): 회사검색 3종 + 산업분류 3종`.

### Task 3: CompanyProfile (sec-profile, 35필드)

**Files:** Create `secfilings/profile.go`, `secfilings/profile_test.go`, testdata `profile.json`.

- [ ] **Step 1:** `secfilings/profile.go` — 스펙 `CompanyProfile`(35필드, SecurityType *string) struct + 메서드:
```go
func (c *Client) Profile(ctx context.Context, symbol, cik string) ([]CompanyProfile, error) {
	if strings.TrimSpace(symbol) == "" && strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: symbol or cik must be provided")
	}
	q := map[string]string{}
	if symbol != "" {
		q["symbol"] = symbol
	}
	if cik != "" {
		q["cik"] = cik
	}
	return fetch.List[CompanyProfile](ctx, c.http, "/stable/sec-profile", q)
}
```
- [ ] **Step 2:** fixture profile.json — 1건, 35필드 전부, securityType:null(SecurityType nil 검증), isActive/isEtf/isAdr/isFund bool, employees:"164000" 문자열, fiftyTwoWeekRange:"164.08 - 260.1".
- [ ] **Step 3:** `profile_test.go`(헬퍼 재사용): Profile("AAPL","") 파싱(RegistrantName!="", IsActive bool, SecurityType==nil, Employees=="164000" string) + delegation(path+symbol) / 빈 symbol+빈 cik 가드 / Profile("","320193") delegation(path+cik).
- [ ] **Step 4:** `go test ./secfilings/ && go vet && gofmt -l`. 커밋 `feat(secfilings): CompanyProfile (sec-profile)`.

### Task 4: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/secfilings/main.go`.

- [ ] **Step 1:** `client.go` — import `secfilings`, struct 에 `SECFilings *secfilings.Client`, NewClient 에 `c.SECFilings = secfilings.New(hc)`. (gofmt -w client.go.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasSECFilings`.
- [ ] **Step 3:** README 표 행 신규: `| SEC Filings | \`client.SECFilings\` | LatestFinancials, Latest8K, SearchBySymbol, SearchByCIK, SearchByFormType, SearchByName, CompanySearchBySymbol, CompanySearchByCIK, Profile, IndustryClassificationList, IndustryClassificationSearch, AllIndustryClassification — 12 endpoint |`.
- [ ] **Step 4:** `examples/secfilings/main.go` — NewClientFromEnv → Profile("AAPL","") 회사명/CEO + Latest8K(최근 from/to, 0, 5) 건수 출력(time.Now 기준 from/to).
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_SECFilings`: Profile("AAPL","") len>0 / IndustryClassificationList("","") len>0 / SearchBySymbol("AAPL", from, to, 0, 5) err 체크(time 사용). import time(이미 있음).
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(secfilings): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 12 endpoint: 필링 5=T1, 회사/산업 6=T2, Profile=T3, 와이어/문서=T4.
- LatestFiling(hasFinancials) vs FilingSearchResult(없음). CompanySearchResult 5개 공유. CompanyProfile securityType *string.
- path family 주의. sicCode/employees string.
