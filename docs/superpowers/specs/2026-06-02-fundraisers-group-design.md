# FMP Go SDK — Fundraisers 그룹 (v0.24.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/fundraisers-group`
- 토픽: FMP `Fundraisers` 카테고리 6 endpoint(크라우드펀딩 + 지분공모). 캠페인 22번째 그룹.

## 결정 사항
- 신규 `fundraisers/` 패키지, internal/fetch. 3 구조체.
- `CrowdfundingCampaign`(48필드, 카탈로그 도출), `EquityOffering`(43필드, 카탈로그 도출), `FundraiserSearchResult`(공유 {cik,name,date}).
- 큰 구조체 2개는 필드 수가 많아 **카탈로그 JSON 에서 전수 매핑**(COT 선례).
- **타입 주의**: Crowdfunding 회계 재무필드(totalAsset/cash/debt/revenue/costGoodsSold/taxesPaid/netIncome × MostRecent/Prior)는 `float64`(소수·음수). compensationAmount/financialInterest 는 free text `string`(이름에 Amount 있어도). offering/maximum/number 류는 int64. JSON 키 오타 `cashAndCashEquiValent...` 보존.
- EquityOffering: `incorporatedWithinFiveYears`, `securitiesOfferedAreOfEquityType` 는 nullable → `*bool`. 나머지 bool 은 bool. 금액/카운트 int64. nullable 문자열은 plain string(null→"").
- 릴리스 `v0.24.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `LatestCrowdfunding(ctx, page, limit)` | `/stable/crowdfunding-offerings-latest` | page,limit | `[]CrowdfundingCampaign` |
| `CrowdfundingByCIK(ctx, cik)` | `/stable/crowdfunding-offerings` | cik(필수) | `[]CrowdfundingCampaign` |
| `CrowdfundingSearch(ctx, name)` | `/stable/crowdfunding-offerings-search` | name(필수) | `[]FundraiserSearchResult` |
| `LatestEquityOffering(ctx, page, limit, cik)` | `/stable/fundraising-latest` | page,limit,cik(선택) | `[]EquityOffering` |
| `EquityOfferingByCIK(ctx, cik)` | `/stable/fundraising` | cik(필수) | `[]EquityOffering` |
| `EquityOfferingSearch(ctx, name)` | `/stable/fundraising-search` | name(필수) | `[]FundraiserSearchResult` |

파일: `fundraisers/client.go`(New + pageParams helper), `fundraisers/crowdfunding.go`(CrowdfundingCampaign + FundraiserSearchResult + 3 crowdfunding method), `fundraisers/equity.go`(EquityOffering + 3 equity method).
- by-cik/search: cik/name 빈값 가드. latest: pageParams (+ LatestEquityOffering 은 cik 선택 추가).

## 루트 Client 와이어
```go
COT         *cot.Client
Fundraisers *fundraisers.Client // 크라우드펀딩/지분공모
```
`c.Fundraisers = fundraisers.New(hc)`. `TestNewClient_HasFundraisers`.

## 응답 타입

### FundraiserSearchResult (crowdfunding-offerings-search / fundraising-search 공유)
```go
// FundraiserSearchResult — 펀드레이저 검색 결과 (크라우드펀딩/지분공모 검색 공유).
type FundraiserSearchResult struct {
	CIK  string `json:"cik"`  // CIK
	Name string `json:"name"` // 회사/발행인명
	Date string `json:"date"` // 일자(null 가능 → 빈 문자열)
}
```

### CrowdfundingCampaign (crowdfunding-offerings-latest / crowdfunding-offerings) — 48필드, 카탈로그 도출
구현 시 `docs/api/Fundraisers/latest-crowdfunding.md` JSON 예시의 모든 키 매핑. 타입 규칙:
- 텍스트/주소/이름/website/formType/legalStatus 등 → string (nullable 포함, null→"")
- compensationAmount, financialInterest → string(free text)
- numberOfSecurityOffered, offeringPrice, offeringAmount, maximumOfferingAmount, currentNumberOfEmployees → int64
- 회계 재무(total/cash/accountsReceivable/shortTermDebt/longTermDebt/revenue/costGoodsSold/taxesPaid/netIncome × MostRecent/Prior) → float64
- JSON 키 오타 `cashAndCashEquiValentMostRecentFiscalYear`/`...PriorFiscalYear` 보존.
- overSubscriptionAccepted → string("Y"/"N").

### EquityOffering (fundraising-latest / fundraising) — 43필드, 카탈로그 도출
구현 시 `docs/api/Fundraisers/latest-equity-offering.md` JSON 예시의 모든 키 매핑. 타입 규칙:
- 텍스트/주소/relatedPerson*/entityName/industryGroupType 등 → string
- 금액/카운트(minimumInvestmentAccepted/totalOfferingAmount/totalAmountSold/totalAmountRemaining/totalNumberAlreadyInvested/salesCommissions/findersFees/grossProceedsUsed) → int64
- bool: isAmendment/durationOfOfferingIsMoreThanYear/isBusinessCombinationTransaction/hasNonAccreditedInvestors → bool
- **nullable bool**: incorporatedWithinFiveYears, securitiesOfferedAreOfEquityType → `*bool`

## 시그니처 규칙
- LatestCrowdfunding: `(ctx, page, limit int)` → pageParams.
- CrowdfundingByCIK/EquityOfferingByCIK: `(ctx, cik string)` → cik 가드 + {cik}.
- CrowdfundingSearch/EquityOfferingSearch: `(ctx, name string)` → name 가드 + {name}.
- LatestEquityOffering: `(ctx, page, limit int, cik string)` → pageParams + cik(비어있지 않으면).

## 테스트
- fixture 단위: CrowdfundingCampaign(회계 float 필드, compensationAmount string, cashAndCashEquiValent 오타 키), EquityOffering(*bool null→nil + bool, int64 금액), FundraiserSearchResult(date null→"").
  - 큰 struct fixture 는 카탈로그 예시 사용. 대표 필드 검증.
- delegation: LatestCrowdfunding(0,10) path+page/limit / CrowdfundingByCIK("0001234") path+cik / CrowdfundingSearch("Acme") path+name / LatestEquityOffering(0,10,"") path / EquityOfferingByCIK path / EquityOfferingSearch path.
- 가드: CrowdfundingByCIK 빈 cik, CrowdfundingSearch 빈 name, EquityOfferingByCIK 빈 cik, EquityOfferingSearch 빈 name(대표 2~3건).
- 통합: LatestCrowdfunding(0,5) / LatestEquityOffering(0,5,"") / EquityOfferingSearch("Tesla") err 체크.

## 문서 / 릴리스
- README Fundraisers 행(6 endpoint).
- `examples/fundraisers/main.go` — LatestCrowdfunding + LatestEquityOffering.
- 릴리스 `v0.24.0`.

## 범위 밖 / 위험
- 큰 struct 는 카탈로그 전수 매핑(누락 키 없도록). 회계 필드 float64(소수/음수).
- nullable bool(*bool) — null 구분. nullable string 은 plain(null→"").
- 다음 그룹: form13F / etfAndMutualFunds / secFilings / crypto·forex·commodity / bulk.
