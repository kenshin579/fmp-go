package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Split — 주식 분할 (splits-calendar / splits 공용)
type Split struct {
	Symbol      string `json:"symbol"`      // 종목 심볼
	Date        string `json:"date"`        // 분할 기준일
	Numerator   int    `json:"numerator"`   // 분할 비율 분자
	Denominator int    `json:"denominator"` // 분할 비율 분모
}

// SplitsCalendar 는 기간 내 주식 분할 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) SplitsCalendar(ctx context.Context, from, to string) ([]Split, error) {
	return fetch.List[Split](ctx, c.http, "/stable/splits-calendar", dateRange(from, to))
}

// CompanySplits 는 특정 종목의 주식 분할 이력을 조회한다.
func (c *Client) CompanySplits(ctx context.Context, symbol string) ([]Split, error) {
	return fetch.ListBySymbol[Split](ctx, c.http, "/stable/splits", symbol)
}
