package quote

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// AftermarketQuote — 시간외 호가
type AftermarketQuote struct {
	Symbol    string  `json:"symbol"`    // 종목 심볼
	BidSize   int64   `json:"bidSize"`   // 매수 호가 수량
	BidPrice  float64 `json:"bidPrice"`  // 매수 호가
	AskSize   int64   `json:"askSize"`   // 매도 호가 수량
	AskPrice  float64 `json:"askPrice"`  // 매도 호가
	Volume    int64   `json:"volume"`    // 거래량
	Timestamp int64   `json:"timestamp"` // 시각 (Unix epoch ms)
}

// AftermarketTrade — 시간외 체결
type AftermarketTrade struct {
	Symbol    string  `json:"symbol"`    // 종목 심볼
	Price     float64 `json:"price"`     // 체결가
	TradeSize int64   `json:"tradeSize"` // 체결 수량
	Timestamp int64   `json:"timestamp"` // 시각 (Unix epoch ms)
}

// AftermarketQuote 는 종목의 시간외 호가를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) AftermarketQuote(ctx context.Context, symbol string) (*AftermarketQuote, error) {
	return fetch.OneBySymbol[AftermarketQuote](ctx, c.http, "/stable/aftermarket-quote", symbol)
}

// AftermarketTrade 는 종목의 시간외 체결을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) AftermarketTrade(ctx context.Context, symbol string) (*AftermarketTrade, error) {
	return fetch.OneBySymbol[AftermarketTrade](ctx, c.http, "/stable/aftermarket-trade", symbol)
}

// BatchAftermarketQuote 는 여러 종목의 시간외 호가를 한 번에 조회한다.
func (c *Client) BatchAftermarketQuote(ctx context.Context, symbols ...string) ([]AftermarketQuote, error) {
	return fetch.ListBySymbols[AftermarketQuote](ctx, c.http, "/stable/batch-aftermarket-quote", symbols)
}

// BatchAftermarketTrade 는 여러 종목의 시간외 체결을 한 번에 조회한다.
func (c *Client) BatchAftermarketTrade(ctx context.Context, symbols ...string) ([]AftermarketTrade, error) {
	return fetch.ListBySymbols[AftermarketTrade](ctx, c.http, "/stable/batch-aftermarket-trade", symbols)
}
