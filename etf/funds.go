package etf

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// FundDisclosureHolderSearch — 펀드 공시 보유자 검색 (funds/disclosure-holders-search)
type FundDisclosureHolderSearch struct {
	Symbol              string `json:"symbol"`
	CIK                 string `json:"cik"`
	ClassID             string `json:"classId"`
	SeriesID            string `json:"seriesId"`
	EntityName          string `json:"entityName"`
	EntityOrgType       string `json:"entityOrgType"`
	SeriesName          string `json:"seriesName"`
	ClassName           string `json:"className"`
	ReportingFileNumber string `json:"reportingFileNumber"`
	Address             string `json:"address"`
	City                string `json:"city"`
	ZipCode             string `json:"zipCode"`
	State               string `json:"state"`
}

// FundDisclosureDate — 펀드 공시 가용 일자 (funds/disclosure-dates)
type FundDisclosureDate struct {
	Date    string `json:"date"`
	Year    int    `json:"year"`
	Quarter int    `json:"quarter"`
}

// FundDisclosureHolderLatest — 최신 펀드 보유자 (funds/disclosure-holders-latest). 키 weightPercent.
type FundDisclosureHolderLatest struct {
	CIK           string  `json:"cik"`
	Holder        string  `json:"holder"`
	Shares        int64   `json:"shares"`
	DateReported  string  `json:"dateReported"`
	Change        int64   `json:"change"`
	WeightPercent float64 `json:"weightPercent"`
}

// MutualFundDisclosure — 뮤추얼펀드 보유 공시 (funds/disclosure). is* 는 "N"/"Y" 문자열.
type MutualFundDisclosure struct {
	CIK                 string  `json:"cik"`
	Date                string  `json:"date"`
	AcceptedDate        string  `json:"acceptedDate"`
	Symbol              string  `json:"symbol"`
	Name                string  `json:"name"`
	LEI                 string  `json:"lei"`
	Title               string  `json:"title"`
	CUSIP               string  `json:"cusip"`
	ISIN                string  `json:"isin"`
	Balance             int64   `json:"balance"`
	Units               string  `json:"units"`
	CurCd               string  `json:"cur_cd"`
	ValUsd              float64 `json:"valUsd"`
	PctVal              float64 `json:"pctVal"`
	PayoffProfile       string  `json:"payoffProfile"`
	AssetCat            string  `json:"assetCat"`
	IssuerCat           string  `json:"issuerCat"`
	InvCountry          string  `json:"invCountry"`
	IsRestrictedSec     string  `json:"isRestrictedSec"`
	FairValLevel        string  `json:"fairValLevel"`
	IsCashCollateral    string  `json:"isCashCollateral"`
	IsNonCashCollateral string  `json:"isNonCashCollateral"`
	IsLoanByFund        string  `json:"isLoanByFund"`
}

func (c *Client) DisclosureHoldersSearch(ctx context.Context, name string) ([]FundDisclosureHolderSearch, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[FundDisclosureHolderSearch](ctx, c.http, "/stable/funds/disclosure-holders-search", map[string]string{"name": name})
}

func (c *Client) DisclosureDates(ctx context.Context, symbol, cik string) ([]FundDisclosureDate, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := map[string]string{"symbol": symbol}
	if cik != "" {
		q["cik"] = cik
	}
	return fetch.List[FundDisclosureDate](ctx, c.http, "/stable/funds/disclosure-dates", q)
}

func (c *Client) LatestDisclosureHolders(ctx context.Context, symbol string) ([]FundDisclosureHolderLatest, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[FundDisclosureHolderLatest](ctx, c.http, "/stable/funds/disclosure-holders-latest", map[string]string{"symbol": symbol})
}

func (c *Client) Disclosure(ctx context.Context, symbol, year, quarter, cik string) ([]MutualFundDisclosure, error) {
	if strings.TrimSpace(symbol) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: symbol, year, quarter must not be empty")
	}
	q := map[string]string{"symbol": symbol, "year": year, "quarter": quarter}
	if cik != "" {
		q["cik"] = cik
	}
	return fetch.List[MutualFundDisclosure](ctx, c.http, "/stable/funds/disclosure", q)
}
