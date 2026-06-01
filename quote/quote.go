package quote

import "context"

// Quote — 전체 시세 (quote / batch-quote / exchange-quote 공용)
type Quote struct {
	Symbol           string  `json:"symbol"`           // 종목 심볼 (예: AAPL)
	Name             string  `json:"name"`             // 종목명
	Price            float64 `json:"price"`            // 현재가
	ChangePercentage float64 `json:"changePercentage"` // 등락률 (%)
	Change           float64 `json:"change"`           // 전일 대비 등락액
	Volume           int64   `json:"volume"`           // 거래량
	DayLow           float64 `json:"dayLow"`           // 당일 저가
	DayHigh          float64 `json:"dayHigh"`          // 당일 고가
	YearHigh         float64 `json:"yearHigh"`         // 52주 최고가
	YearLow          float64 `json:"yearLow"`          // 52주 최저가
	MarketCap        int64   `json:"marketCap"`        // 시가총액
	PriceAvg50       float64 `json:"priceAvg50"`       // 50일 이동평균가
	PriceAvg200      float64 `json:"priceAvg200"`      // 200일 이동평균가
	Exchange         string  `json:"exchange"`         // 거래소 (예: NASDAQ)
	Open             float64 `json:"open"`             // 시가
	PreviousClose    float64 `json:"previousClose"`    // 전일 종가
	Timestamp        int64   `json:"timestamp"`        // 시세 시각 (Unix epoch sec)
}

// Quote 는 종목의 전체 시세를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) Quote(ctx context.Context, symbol string) (*Quote, error) {
	return fetchOne[Quote](ctx, c, "/stable/quote", symbol)
}

// BatchQuote 는 여러 종목의 전체 시세를 한 번에 조회한다.
func (c *Client) BatchQuote(ctx context.Context, symbols ...string) ([]Quote, error) {
	return fetchBatch[Quote](ctx, c, "/stable/batch-quote", symbols)
}
