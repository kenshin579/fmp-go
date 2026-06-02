package fundraisers

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CrowdfundingCampaign — 크라우드펀딩 캠페인 (crowdfunding-offerings-latest / crowdfunding-offerings 공유).
type CrowdfundingCampaign struct {
	CIK                                       string  `json:"cik"`
	CompanyName                               string  `json:"companyName"`
	Date                                      string  `json:"date"`
	FilingDate                                string  `json:"filingDate"`
	AcceptedDate                              string  `json:"acceptedDate"`
	FormType                                  string  `json:"formType"`
	FormSignification                         string  `json:"formSignification"`
	NameOfIssuer                              string  `json:"nameOfIssuer"`
	LegalStatusForm                           string  `json:"legalStatusForm"`
	JurisdictionOrganization                  string  `json:"jurisdictionOrganization"`
	IssuerStreet                              string  `json:"issuerStreet"`
	IssuerCity                                string  `json:"issuerCity"`
	IssuerStateOrCountry                      string  `json:"issuerStateOrCountry"`
	IssuerZipCode                             string  `json:"issuerZipCode"`
	IssuerWebsite                             string  `json:"issuerWebsite"`
	IntermediaryCompanyName                   string  `json:"intermediaryCompanyName"`
	IntermediaryCommissionCik                 string  `json:"intermediaryCommissionCik"`
	IntermediaryCommissionFileNumber          string  `json:"intermediaryCommissionFileNumber"`
	CompensationAmount                        string  `json:"compensationAmount"`
	FinancialInterest                         string  `json:"financialInterest"`
	SecurityOfferedType                       string  `json:"securityOfferedType"`
	SecurityOfferedOtherDescription           string  `json:"securityOfferedOtherDescription"`
	NumberOfSecurityOffered                   int64   `json:"numberOfSecurityOffered"`
	OfferingPrice                             int64   `json:"offeringPrice"`
	OfferingAmount                            int64   `json:"offeringAmount"`
	OverSubscriptionAccepted                  string  `json:"overSubscriptionAccepted"`
	OverSubscriptionAllocationType            string  `json:"overSubscriptionAllocationType"`
	MaximumOfferingAmount                     int64   `json:"maximumOfferingAmount"`
	OfferingDeadlineDate                      string  `json:"offeringDeadlineDate"`
	CurrentNumberOfEmployees                  int64   `json:"currentNumberOfEmployees"`
	TotalAssetMostRecentFiscalYear            float64 `json:"totalAssetMostRecentFiscalYear"`
	TotalAssetPriorFiscalYear                 float64 `json:"totalAssetPriorFiscalYear"`
	CashAndCashEquiValentMostRecentFiscalYear float64 `json:"cashAndCashEquiValentMostRecentFiscalYear"`
	CashAndCashEquiValentPriorFiscalYear      float64 `json:"cashAndCashEquiValentPriorFiscalYear"`
	AccountsReceivableMostRecentFiscalYear    float64 `json:"accountsReceivableMostRecentFiscalYear"`
	AccountsReceivablePriorFiscalYear         float64 `json:"accountsReceivablePriorFiscalYear"`
	ShortTermDebtMostRecentFiscalYear         float64 `json:"shortTermDebtMostRecentFiscalYear"`
	ShortTermDebtPriorFiscalYear              float64 `json:"shortTermDebtPriorFiscalYear"`
	LongTermDebtMostRecentFiscalYear          float64 `json:"longTermDebtMostRecentFiscalYear"`
	LongTermDebtPriorFiscalYear               float64 `json:"longTermDebtPriorFiscalYear"`
	RevenueMostRecentFiscalYear               float64 `json:"revenueMostRecentFiscalYear"`
	RevenuePriorFiscalYear                    float64 `json:"revenuePriorFiscalYear"`
	CostGoodsSoldMostRecentFiscalYear         float64 `json:"costGoodsSoldMostRecentFiscalYear"`
	CostGoodsSoldPriorFiscalYear              float64 `json:"costGoodsSoldPriorFiscalYear"`
	TaxesPaidMostRecentFiscalYear             float64 `json:"taxesPaidMostRecentFiscalYear"`
	TaxesPaidPriorFiscalYear                  float64 `json:"taxesPaidPriorFiscalYear"`
	NetIncomeMostRecentFiscalYear             float64 `json:"netIncomeMostRecentFiscalYear"`
	NetIncomePriorFiscalYear                  float64 `json:"netIncomePriorFiscalYear"`
}

// FundraiserSearchResult — 펀드레이저 검색 결과 (크라우드펀딩/지분공모 검색 공유).
type FundraiserSearchResult struct {
	CIK  string `json:"cik"`  // CIK
	Name string `json:"name"` // 회사/발행인명
	Date string `json:"date"` // 일자(null 가능 → 빈 문자열)
}

// LatestCrowdfunding 은 최신 크라우드펀딩 캠페인 목록을 반환한다.
func (c *Client) LatestCrowdfunding(ctx context.Context, page, limit int) ([]CrowdfundingCampaign, error) {
	return fetch.List[CrowdfundingCampaign](ctx, c.http, "/stable/crowdfunding-offerings-latest", pageParams(page, limit))
}

// CrowdfundingByCIK 은 CIK 로 크라우드펀딩 캠페인을 조회한다.
func (c *Client) CrowdfundingByCIK(ctx context.Context, cik string) ([]CrowdfundingCampaign, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[CrowdfundingCampaign](ctx, c.http, "/stable/crowdfunding-offerings", map[string]string{"cik": cik})
}

// CrowdfundingSearch 는 회사명으로 크라우드펀딩 캠페인을 검색한다.
func (c *Client) CrowdfundingSearch(ctx context.Context, name string) ([]FundraiserSearchResult, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[FundraiserSearchResult](ctx, c.http, "/stable/crowdfunding-offerings-search", map[string]string{"name": name})
}
