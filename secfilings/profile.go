package secfilings

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CompanyProfile — SEC 회사 전체 프로필 (sec-profile)
type CompanyProfile struct {
	Symbol                  string  `json:"symbol"`
	CIK                     string  `json:"cik"`
	RegistrantName          string  `json:"registrantName"`
	SICCode                 string  `json:"sicCode"`
	SICDescription          string  `json:"sicDescription"`
	SICGroup                string  `json:"sicGroup"`
	ISIN                    string  `json:"isin"`
	BusinessAddress         string  `json:"businessAddress"`
	MailingAddress          string  `json:"mailingAddress"`
	PhoneNumber             string  `json:"phoneNumber"`
	PostalCode              string  `json:"postalCode"`
	City                    string  `json:"city"`
	State                   string  `json:"state"`
	Country                 string  `json:"country"`
	Description             string  `json:"description"`
	CEO                     string  `json:"ceo"`
	Website                 string  `json:"website"`
	Exchange                string  `json:"exchange"`
	StateLocation           string  `json:"stateLocation"`
	StateOfIncorporation    string  `json:"stateOfIncorporation"`
	FiscalYearEnd           string  `json:"fiscalYearEnd"`
	IPODate                 string  `json:"ipoDate"`
	Employees               string  `json:"employees"`
	SECFilingsURL           string  `json:"secFilingsUrl"`
	TaxIdentificationNumber string  `json:"taxIdentificationNumber"`
	FiftyTwoWeekRange       string  `json:"fiftyTwoWeekRange"`
	IsActive                bool    `json:"isActive"`
	AssetType               string  `json:"assetType"`
	OpenFigiComposite       string  `json:"openFigiComposite"`
	PriceCurrency           string  `json:"priceCurrency"`
	MarketSector            string  `json:"marketSector"`
	SecurityType            *string `json:"securityType"`
	IsEtf                   bool    `json:"isEtf"`
	IsAdr                   bool    `json:"isAdr"`
	IsFund                  bool    `json:"isFund"`
}

// Profile 은 SEC 회사 전체 프로필을 조회한다. symbol 또는 cik 중 하나 필수.
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
