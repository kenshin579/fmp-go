# Fundraisers 그룹 (v0.24.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `fundraisers/` 패키지 6 endpoint, 3 구조체(큰 2개는 카탈로그 도출).

**Architecture:** internal/fetch(List) + pageParams. CrowdfundingCampaign/EquityOffering 는 카탈로그 JSON 전수 매핑, FundraiserSearchResult 공유.

참고: `unset GOROOT`. 커밋 한국어 `feat(fundraisers): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 + Crowdfunding 3종 + FundraiserSearchResult

**Files:** Create `fundraisers/client.go`, `fundraisers/crowdfunding.go`, `fundraisers/crowdfunding_test.go`, testdata `crowdfunding-latest.json`, `crowdfunding-search.json`.

- [ ] **Step 1:** `fundraisers/client.go`:
```go
// Package fundraisers 는 FMP 크라우드펀딩/지분공모 API sub-client.
// fmp.Client.Fundraisers 로 접근.
package fundraisers

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
- [ ] **Step 2:** READ `docs/api/Fundraisers/latest-crowdfunding.md` JSON 예시. `fundraisers/crowdfunding.go` 에 `CrowdfundingCampaign` struct(예시의 모든 키 매핑) + `FundraiserSearchResult` struct(스펙 verbatim) + 3 메서드. 타입 규칙(스펙): 회계 재무 float64, compensationAmount/financialInterest string, offering/number 류 int64, `cashAndCashEquiValent...` 오타 키 보존, 텍스트 string. import: context, fmt, strings, internal/fetch.
```go
func (c *Client) LatestCrowdfunding(ctx context.Context, page, limit int) ([]CrowdfundingCampaign, error) {
	return fetch.List[CrowdfundingCampaign](ctx, c.http, "/stable/crowdfunding-offerings-latest", pageParams(page, limit))
}
func (c *Client) CrowdfundingByCIK(ctx context.Context, cik string) ([]CrowdfundingCampaign, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[CrowdfundingCampaign](ctx, c.http, "/stable/crowdfunding-offerings", map[string]string{"cik": cik})
}
func (c *Client) CrowdfundingSearch(ctx context.Context, name string) ([]FundraiserSearchResult, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[FundraiserSearchResult](ctx, c.http, "/stable/crowdfunding-offerings-search", map[string]string{"name": name})
}
```
- [ ] **Step 3:** fixtures: crowdfunding-latest.json = 카탈로그 latest-crowdfunding.md JSON 예시 그대로(배열). crowdfunding-search.json `[{cik:"0001234567",name:"Acme Inc",date:null}]`.
- [ ] **Step 4:** `crowdfunding_test.go`(헬퍼 calendar 패턴 정의): LatestCrowdfunding 파싱(대표 — CompanyName!="", 회계 float 필드 파싱, compensationAmount string) + delegation(path+page/limit) / CrowdfundingByCIK delegation(path+cik) + 빈 cik 가드 / CrowdfundingSearch 파싱(Date null→"" , Name!="") + delegation(path+name) + 빈 name 가드.
- [ ] **Step 5:** `unset GOROOT && go test ./fundraisers/ && go vet ./fundraisers/ && gofmt -l fundraisers/`. 커밋 `feat(fundraisers): Crowdfunding 3종 + 검색결과`.

### Task 2: Equity Offering 3종

**Files:** Create `fundraisers/equity.go`, `fundraisers/equity_test.go`, testdata `equity-latest.json`.

- [ ] **Step 1:** READ `docs/api/Fundraisers/latest-equity-offering.md` JSON 예시. `fundraisers/equity.go` 에 `EquityOffering` struct(예시의 모든 키 매핑) + 3 메서드. 타입 규칙(스펙): 금액/카운트 int64, bool 4종 bool, `incorporatedWithinFiveYears`/`securitiesOfferedAreOfEquityType` 는 `*bool`(nullable), 텍스트 string. import: context, fmt, strconv, strings, internal/fetch.
```go
func (c *Client) LatestEquityOffering(ctx context.Context, page, limit int, cik string) ([]EquityOffering, error) {
	q := pageParams(page, limit)
	if cik != "" {
		q["cik"] = cik
	}
	return fetch.List[EquityOffering](ctx, c.http, "/stable/fundraising-latest", q)
}
func (c *Client) EquityOfferingByCIK(ctx context.Context, cik string) ([]EquityOffering, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[EquityOffering](ctx, c.http, "/stable/fundraising", map[string]string{"cik": cik})
}
func (c *Client) EquityOfferingSearch(ctx context.Context, name string) ([]FundraiserSearchResult, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[FundraiserSearchResult](ctx, c.http, "/stable/fundraising-search", map[string]string{"name": name})
}
```
- [ ] **Step 2:** fixture equity-latest.json = 카탈로그 latest-equity-offering.md JSON 예시 그대로(배열). 단 nullable bool 키(securitiesOfferedAreOfEquityType 등)가 예시에서 null 이면 그대로 null 유지(없으면 한 건은 null 로 조정해 *bool nil 검증 가능하게).
- [ ] **Step 3:** `equity_test.go`(헬퍼 재사용): LatestEquityOffering 파싱(TotalOfferingAmount int64, IsAmendment bool, 그리고 nullable *bool 필드가 null 이면 nil) + delegation(path+page/limit) / EquityOfferingByCIK delegation(path+cik) + 빈 cik 가드 / EquityOfferingSearch delegation(path+name) + 빈 name 가드.
- [ ] **Step 4:** `go test ./fundraisers/ && go vet && gofmt -l`. 커밋 `feat(fundraisers): Equity Offering 3종`.

### Task 3: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/fundraisers/main.go`.

- [ ] **Step 1:** `client.go` — import `fundraisers`, struct 에 `Fundraisers *fundraisers.Client`, NewClient 에 `c.Fundraisers = fundraisers.New(hc)`. (gofmt -w client.go.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasFundraisers`.
- [ ] **Step 3:** README 표 행 신규: `| Fundraisers | \`client.Fundraisers\` | LatestCrowdfunding, CrowdfundingByCIK, CrowdfundingSearch, LatestEquityOffering, EquityOfferingByCIK, EquityOfferingSearch — 6 endpoint |`.
- [ ] **Step 4:** `examples/fundraisers/main.go` — NewClientFromEnv → LatestCrowdfunding(0,5) 건수 + LatestEquityOffering(0,5,"") 건수 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Fundraisers`: LatestCrowdfunding(0,5) err 체크 / LatestEquityOffering(0,5,"") err 체크 / EquityOfferingSearch("Tesla") err 체크.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(fundraisers): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 6 endpoint: Crowdfunding 3=T1, Equity 3=T2, 와이어/문서=T3.
- 큰 struct 카탈로그 전수 매핑. Crowdfunding 회계 float64, cashAndCashEquiValent 오타. EquityOffering 2개 *bool nullable.
- FundraiserSearchResult 공유.
