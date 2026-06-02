package fundraisers

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// EquityOffering — 지분공모 (fundraising-latest / fundraising 공유).
type EquityOffering struct {
	CIK                                    string `json:"cik"`
	CompanyName                            string `json:"companyName"`
	Date                                   string `json:"date"`
	FilingDate                             string `json:"filingDate"`
	AcceptedDate                           string `json:"acceptedDate"`
	FormType                               string `json:"formType"`
	FormSignification                      string `json:"formSignification"`
	EntityName                             string `json:"entityName"`
	IssuerStreet                           string `json:"issuerStreet"`
	IssuerCity                             string `json:"issuerCity"`
	IssuerStateOrCountry                   string `json:"issuerStateOrCountry"`
	IssuerStateOrCountryDescription        string `json:"issuerStateOrCountryDescription"`
	IssuerZipCode                          string `json:"issuerZipCode"`
	IssuerPhoneNumber                      string `json:"issuerPhoneNumber"`
	JurisdictionOfIncorporation            string `json:"jurisdictionOfIncorporation"`
	EntityType                             string `json:"entityType"`
	IncorporatedWithinFiveYears            *bool  `json:"incorporatedWithinFiveYears"`
	YearOfIncorporation                    string `json:"yearOfIncorporation"`
	RelatedPersonFirstName                 string `json:"relatedPersonFirstName"`
	RelatedPersonLastName                  string `json:"relatedPersonLastName"`
	RelatedPersonStreet                    string `json:"relatedPersonStreet"`
	RelatedPersonCity                      string `json:"relatedPersonCity"`
	RelatedPersonStateOrCountry            string `json:"relatedPersonStateOrCountry"`
	RelatedPersonStateOrCountryDescription string `json:"relatedPersonStateOrCountryDescription"`
	RelatedPersonZipCode                   string `json:"relatedPersonZipCode"`
	RelatedPersonRelationship              string `json:"relatedPersonRelationship"`
	IndustryGroupType                      string `json:"industryGroupType"`
	RevenueRange                           string `json:"revenueRange"`
	FederalExemptionsExclusions            string `json:"federalExemptionsExclusions"`
	IsAmendment                            bool   `json:"isAmendment"`
	DateOfFirstSale                        string `json:"dateOfFirstSale"`
	DurationOfOfferingIsMoreThanYear       bool   `json:"durationOfOfferingIsMoreThanYear"`
	SecuritiesOfferedAreOfEquityType       *bool  `json:"securitiesOfferedAreOfEquityType"`
	IsBusinessCombinationTransaction       bool   `json:"isBusinessCombinationTransaction"`
	MinimumInvestmentAccepted              int64  `json:"minimumInvestmentAccepted"`
	TotalOfferingAmount                    int64  `json:"totalOfferingAmount"`
	TotalAmountSold                        int64  `json:"totalAmountSold"`
	TotalAmountRemaining                   int64  `json:"totalAmountRemaining"`
	HasNonAccreditedInvestors              bool   `json:"hasNonAccreditedInvestors"`
	TotalNumberAlreadyInvested             int64  `json:"totalNumberAlreadyInvested"`
	SalesCommissions                       int64  `json:"salesCommissions"`
	FindersFees                            int64  `json:"findersFees"`
	GrossProceedsUsed                      int64  `json:"grossProceedsUsed"`
}

// LatestEquityOffering 은 최신 지분공모 목록을 반환한다.
func (c *Client) LatestEquityOffering(ctx context.Context, page, limit int, cik string) ([]EquityOffering, error) {
	q := pageParams(page, limit)
	if cik != "" {
		q["cik"] = cik
	}
	return fetch.List[EquityOffering](ctx, c.http, "/stable/fundraising-latest", q)
}

// EquityOfferingByCIK 은 CIK 로 지분공모를 조회한다.
func (c *Client) EquityOfferingByCIK(ctx context.Context, cik string) ([]EquityOffering, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[EquityOffering](ctx, c.http, "/stable/fundraising", map[string]string{"cik": cik})
}

// EquityOfferingSearch 는 회사명으로 지분공모를 검색한다.
func (c *Client) EquityOfferingSearch(ctx context.Context, name string) ([]FundraiserSearchResult, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[FundraiserSearchResult](ctx, c.http, "/stable/fundraising-search", map[string]string{"name": name})
}
