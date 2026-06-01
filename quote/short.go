package quote

import "context"

// QuoteShort — 경량 시세 (quote-short / batch-quote-short / 자산군 배치 공용)
type QuoteShort struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	Price  float64 `json:"price"`  // 현재가
	Change float64 `json:"change"` // 전일 대비 등락액
	Volume int64   `json:"volume"` // 거래량
}

// QuoteShort 는 종목의 경량 시세를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) QuoteShort(ctx context.Context, symbol string) (*QuoteShort, error) {
	return fetchOne[QuoteShort](ctx, c, "/stable/quote-short", symbol)
}

// BatchQuoteShort 는 여러 종목의 경량 시세를 한 번에 조회한다.
func (c *Client) BatchQuoteShort(ctx context.Context, symbols ...string) ([]QuoteShort, error) {
	return fetchBatch[QuoteShort](ctx, c, "/stable/batch-quote-short", symbols)
}
