package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Earning — 실적 (earnings-calendar / earnings 공용). 미래 실적은 actual/estimated null.
type Earning struct {
	Symbol           string   `json:"symbol"`           // 종목 심볼
	Date             string   `json:"date"`             // 실적 발표일
	EpsActual        *float64 `json:"epsActual"`        // 실제 EPS(결측 가능)
	EpsEstimated     *float64 `json:"epsEstimated"`     // 추정 EPS(결측 가능)
	RevenueActual    *int64   `json:"revenueActual"`    // 실제 매출(결측 가능)
	RevenueEstimated *int64   `json:"revenueEstimated"` // 추정 매출(결측 가능)
	LastUpdated      string   `json:"lastUpdated"`      // 최종 갱신일
}

// EarningsCalendar 는 기간 내 전체 실적 발표 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) EarningsCalendar(ctx context.Context, from, to string) ([]Earning, error) {
	return fetch.List[Earning](ctx, c.http, "/stable/earnings-calendar", dateRange(from, to))
}

// CompanyEarnings 는 종목의 실적 이력/예정을 조회한다.
func (c *Client) CompanyEarnings(ctx context.Context, symbol string) ([]Earning, error) {
	return fetch.ListBySymbol[Earning](ctx, c.http, "/stable/earnings", symbol)
}
