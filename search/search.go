package search

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SymbolSearchResult — 심볼/회사명 검색 결과 (search-symbol / search-name 공용)
type SymbolSearchResult struct {
	Symbol           string `json:"symbol"`           // 종목 심볼
	Name             string `json:"name"`             // 종목/회사명
	Currency         string `json:"currency"`         // 통화
	ExchangeFullName string `json:"exchangeFullName"` // 거래소 전체명
	Exchange         string `json:"exchange"`         // 거래소 코드
}

// SearchSymbol 은 심볼(티커)로 종목을 검색한다.
func (c *Client) SearchSymbol(ctx context.Context, query string) ([]SymbolSearchResult, error) {
	if strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("fmp: query must not be empty")
	}
	return fetch.List[SymbolSearchResult](ctx, c.http, "/stable/search-symbol", map[string]string{"query": query})
}

// SearchName 은 회사명으로 종목을 검색한다.
func (c *Client) SearchName(ctx context.Context, query string) ([]SymbolSearchResult, error) {
	if strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("fmp: query must not be empty")
	}
	return fetch.List[SymbolSearchResult](ctx, c.http, "/stable/search-name", map[string]string{"query": query})
}
