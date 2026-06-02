package company

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SharesFloat — 유통 주식 수(float)
type SharesFloat struct {
	Symbol            string  `json:"symbol"`            // 종목 심볼
	Date              string  `json:"date"`              // 기준일
	FreeFloat         float64 `json:"freeFloat"`         // 유통 비율 (%)
	FloatShares       int64   `json:"floatShares"`       // 유통 주식 수
	OutstandingShares int64   `json:"outstandingShares"` // 발행 주식 총수
}

// SharesFloat 은 종목의 유통 주식 정보를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) SharesFloat(ctx context.Context, symbol string) (*SharesFloat, error) {
	return fetch.OneBySymbol[SharesFloat](ctx, c.http, "/stable/shares-float", symbol)
}

// AllSharesFloat 은 전체 종목의 유통 주식 정보를 페이지 단위로 조회한다.
func (c *Client) AllSharesFloat(ctx context.Context, page int) ([]SharesFloat, error) {
	return fetch.List[SharesFloat](ctx, c.http, "/stable/shares-float-all", map[string]string{"page": strconv.Itoa(page)})
}
