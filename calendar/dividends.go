package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Dividend — 배당 (dividends-calendar / dividends 공용)
type Dividend struct {
	Symbol          string  `json:"symbol"`          // 종목 심볼
	Date            string  `json:"date"`            // 배당락일
	RecordDate      string  `json:"recordDate"`      // 기준일
	PaymentDate     string  `json:"paymentDate"`     // 지급일
	DeclarationDate string  `json:"declarationDate"` // 선언일
	AdjDividend     float64 `json:"adjDividend"`     // 수정 배당금
	Dividend        float64 `json:"dividend"`        // 배당금
	Yield           float64 `json:"yield"`           // 배당수익률 (%)
	Frequency       string  `json:"frequency"`       // 배당 주기 (예: Quarterly)
}

// DividendsCalendar 는 기간 내 전체 배당 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) DividendsCalendar(ctx context.Context, from, to string) ([]Dividend, error) {
	return fetch.List[Dividend](ctx, c.http, "/stable/dividends-calendar", dateRange(from, to))
}

// CompanyDividends 는 종목의 배당 이력을 조회한다.
func (c *Client) CompanyDividends(ctx context.Context, symbol string) ([]Dividend, error) {
	return fetch.ListBySymbol[Dividend](ctx, c.http, "/stable/dividends", symbol)
}
