package secfilings

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// LatestFiling — 최신 SEC 공시 (sec-filings-financials / sec-filings-8k 공유)
type LatestFiling struct {
	Symbol        string `json:"symbol"`
	CIK           string `json:"cik"`
	FilingDate    string `json:"filingDate"`
	AcceptedDate  string `json:"acceptedDate"`
	FormType      string `json:"formType"`
	HasFinancials bool   `json:"hasFinancials"`
	Link          string `json:"link"`
	FinalLink     string `json:"finalLink"`
}

// FilingSearchResult — SEC 공시 검색 결과 (sec-filings-search/{symbol,cik,form-type} 공유)
type FilingSearchResult struct {
	Symbol       string `json:"symbol"`
	CIK          string `json:"cik"`
	FilingDate   string `json:"filingDate"`
	AcceptedDate string `json:"acceptedDate"`
	FormType     string `json:"formType"`
	Link         string `json:"link"`
	FinalLink    string `json:"finalLink"`
}

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
