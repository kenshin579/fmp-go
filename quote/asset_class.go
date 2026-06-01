package quote

import "context"

// ExchangeQuotes 는 특정 거래소의 전체 종목 시세를 조회한다 (예: NASDAQ).
func (c *Client) ExchangeQuotes(ctx context.Context, exchange string) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-exchange-quote", map[string]string{"exchange": exchange})
}

// IndexQuotes 는 전체 지수 시세를 조회한다.
func (c *Client) IndexQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-index-quotes", nil)
}

// CommodityQuotes 는 전체 원자재 시세를 조회한다.
func (c *Client) CommodityQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-commodity-quotes", nil)
}

// CryptoQuotes 는 전체 암호화폐 시세를 조회한다.
func (c *Client) CryptoQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-crypto-quotes", nil)
}

// ETFQuotes 는 전체 ETF 시세를 조회한다.
func (c *Client) ETFQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-etf-quotes", nil)
}

// ForexQuotes 는 전체 외환 시세를 조회한다.
func (c *Client) ForexQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-forex-quotes", nil)
}

// MutualFundQuotes 는 전체 뮤추얼펀드 시세를 조회한다.
func (c *Client) MutualFundQuotes(ctx context.Context) ([]QuoteShort, error) {
	return fetchList[QuoteShort](ctx, c, "/stable/batch-mutualfund-quotes", nil)
}
