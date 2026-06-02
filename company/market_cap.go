package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// MarketCap — 시가총액 (market-capitalization / historical / batch 공용)
type MarketCap struct {
	Symbol    string `json:"symbol"`    // 종목 심볼
	Date      string `json:"date"`      // 기준일 (YYYY-MM-DD)
	MarketCap int64  `json:"marketCap"` // 시가총액
}

// MarketCap 은 종목의 현재 시가총액을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) MarketCap(ctx context.Context, symbol string) (*MarketCap, error) {
	return fetch.OneBySymbol[MarketCap](ctx, c.http, "/stable/market-capitalization", symbol)
}

// HistoricalMarketCap 은 종목의 시가총액 시계열을 조회한다.
func (c *Client) HistoricalMarketCap(ctx context.Context, symbol string) ([]MarketCap, error) {
	return fetch.ListBySymbol[MarketCap](ctx, c.http, "/stable/historical-market-capitalization", symbol)
}

// BatchMarketCap 은 여러 종목의 시가총액을 한 번에 조회한다.
func (c *Client) BatchMarketCap(ctx context.Context, symbols ...string) ([]MarketCap, error) {
	return fetch.ListBySymbols[MarketCap](ctx, c.http, "/stable/market-capitalization-batch", symbols)
}
