package search

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CIKSearchResult — CIK 검색 결과
type CIKSearchResult struct {
	Symbol           string `json:"symbol"`           // 종목 심볼
	CompanyName      string `json:"companyName"`      // 회사명
	CIK              string `json:"cik"`              // SEC CIK
	ExchangeFullName string `json:"exchangeFullName"` // 거래소 전체명
	Exchange         string `json:"exchange"`         // 거래소 코드
	Currency         string `json:"currency"`         // 통화
}

// CUSIPSearchResult — CUSIP 검색 결과
type CUSIPSearchResult struct {
	Symbol      string  `json:"symbol"`      // 종목 심볼
	CompanyName string  `json:"companyName"` // 회사명
	CUSIP       string  `json:"cusip"`       // CUSIP 코드
	MarketCap   float64 `json:"marketCap"`   // 시가총액
}

// ISINSearchResult — ISIN 검색 결과
type ISINSearchResult struct {
	Symbol    string `json:"symbol"`    // 종목 심볼
	Name      string `json:"name"`      // 회사명
	ISIN      string `json:"isin"`      // ISIN 코드
	MarketCap int64  `json:"marketCap"` // 시가총액
}

// SearchCIK 는 SEC CIK 로 종목을 검색한다.
func (c *Client) SearchCIK(ctx context.Context, cik string) ([]CIKSearchResult, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[CIKSearchResult](ctx, c.http, "/stable/search-cik", map[string]string{"cik": cik})
}

// SearchCUSIP 는 CUSIP 코드로 종목을 검색한다.
func (c *Client) SearchCUSIP(ctx context.Context, cusip string) ([]CUSIPSearchResult, error) {
	if strings.TrimSpace(cusip) == "" {
		return nil, fmt.Errorf("fmp: cusip must not be empty")
	}
	return fetch.List[CUSIPSearchResult](ctx, c.http, "/stable/search-cusip", map[string]string{"cusip": cusip})
}

// SearchISIN 는 ISIN 코드로 종목을 검색한다.
func (c *Client) SearchISIN(ctx context.Context, isin string) ([]ISINSearchResult, error) {
	if strings.TrimSpace(isin) == "" {
		return nil, fmt.Errorf("fmp: isin must not be empty")
	}
	return fetch.List[ISINSearchResult](ctx, c.http, "/stable/search-isin", map[string]string{"isin": isin})
}
