package quote

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// PriceChange — 기간별 등락률(%) — stock-price-change
type PriceChange struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	D1     float64 `json:"1D"`     // 1일 등락률 (%)
	D5     float64 `json:"5D"`     // 5일
	M1     float64 `json:"1M"`     // 1개월
	M3     float64 `json:"3M"`     // 3개월
	M6     float64 `json:"6M"`     // 6개월
	YTD    float64 `json:"ytd"`    // 연초 대비
	Y1     float64 `json:"1Y"`     // 1년
	Y3     float64 `json:"3Y"`     // 3년
	Y5     float64 `json:"5Y"`     // 5년
	Y10    float64 `json:"10Y"`    // 10년
	Max    float64 `json:"max"`    // 상장 이후 전체
}

// PriceChange 는 종목의 기간별 등락률을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) PriceChange(ctx context.Context, symbol string) (*PriceChange, error) {
	return fetch.OneBySymbol[PriceChange](ctx, c.http, "/stable/stock-price-change", symbol)
}
