package secfilings

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CompanySearchResult — 회사 검색/산업분류 결과 (company-search + industry-classification 공유).
type CompanySearchResult struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	CIK             string `json:"cik"`
	SICCode         string `json:"sicCode"`
	IndustryTitle   string `json:"industryTitle"`
	BusinessAddress string `json:"businessAddress"`
	PhoneNumber     string `json:"phoneNumber"`
}

// IndustryClassification — SIC 산업분류 목록 (standard-industrial-classification-list)
type IndustryClassification struct {
	Office        string `json:"office"`
	SICCode       string `json:"sicCode"`
	IndustryTitle string `json:"industryTitle"`
}

func (c *Client) SearchByName(ctx context.Context, company string) ([]CompanySearchResult, error) {
	if strings.TrimSpace(company) == "" {
		return nil, fmt.Errorf("fmp: company must not be empty")
	}
	return fetch.List[CompanySearchResult](ctx, c.http, "/stable/sec-filings-company-search/name", map[string]string{"company": company})
}
func (c *Client) CompanySearchBySymbol(ctx context.Context, symbol string) ([]CompanySearchResult, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[CompanySearchResult](ctx, c.http, "/stable/sec-filings-company-search/symbol", map[string]string{"symbol": symbol})
}
func (c *Client) CompanySearchByCIK(ctx context.Context, cik string) ([]CompanySearchResult, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[CompanySearchResult](ctx, c.http, "/stable/sec-filings-company-search/cik", map[string]string{"cik": cik})
}
func (c *Client) IndustryClassificationList(ctx context.Context, industryTitle, sicCode string) ([]IndustryClassification, error) {
	q := map[string]string{}
	if industryTitle != "" {
		q["industryTitle"] = industryTitle
	}
	if sicCode != "" {
		q["sicCode"] = sicCode
	}
	return fetch.List[IndustryClassification](ctx, c.http, "/stable/standard-industrial-classification-list", q)
}
func (c *Client) IndustryClassificationSearch(ctx context.Context, symbol, cik, sicCode string) ([]CompanySearchResult, error) {
	q := map[string]string{}
	if symbol != "" {
		q["symbol"] = symbol
	}
	if cik != "" {
		q["cik"] = cik
	}
	if sicCode != "" {
		q["sicCode"] = sicCode
	}
	return fetch.List[CompanySearchResult](ctx, c.http, "/stable/industry-classification-search", q)
}
func (c *Client) AllIndustryClassification(ctx context.Context, page, limit int) ([]CompanySearchResult, error) {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[CompanySearchResult](ctx, c.http, "/stable/all-industry-classification", q)
}
