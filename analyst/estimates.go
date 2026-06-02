package analyst

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// FinancialEstimate — 애널리스트 재무 추정 (analyst-estimates). 카탈로그 예시 없음 → FMP 공개 shape 합성.
type FinancialEstimate struct {
	Symbol             string  `json:"symbol"`             // 종목 심볼
	Date               string  `json:"date"`               // 추정 기준일
	RevenueLow         int64   `json:"revenueLow"`         // 매출 추정 최저
	RevenueHigh        int64   `json:"revenueHigh"`        // 매출 추정 최고
	RevenueAvg         int64   `json:"revenueAvg"`         // 매출 추정 평균
	EbitdaLow          int64   `json:"ebitdaLow"`          // EBITDA 최저
	EbitdaHigh         int64   `json:"ebitdaHigh"`         // EBITDA 최고
	EbitdaAvg          int64   `json:"ebitdaAvg"`          // EBITDA 평균
	EbitLow            int64   `json:"ebitLow"`            // EBIT 최저
	EbitHigh           int64   `json:"ebitHigh"`           // EBIT 최고
	EbitAvg            int64   `json:"ebitAvg"`            // EBIT 평균
	NetIncomeLow       int64   `json:"netIncomeLow"`       // 순이익 최저
	NetIncomeHigh      int64   `json:"netIncomeHigh"`      // 순이익 최고
	NetIncomeAvg       int64   `json:"netIncomeAvg"`       // 순이익 평균
	SgaExpenseLow      int64   `json:"sgaExpenseLow"`      // 판관비 최저
	SgaExpenseHigh     int64   `json:"sgaExpenseHigh"`     // 판관비 최고
	SgaExpenseAvg      int64   `json:"sgaExpenseAvg"`      // 판관비 평균
	EpsLow             float64 `json:"epsLow"`             // EPS 최저
	EpsHigh            float64 `json:"epsHigh"`            // EPS 최고
	EpsAvg             float64 `json:"epsAvg"`             // EPS 평균
	NumAnalystsRevenue int     `json:"numAnalystsRevenue"` // 매출 추정 애널리스트 수
	NumAnalystsEps     int     `json:"numAnalystsEps"`     // EPS 추정 애널리스트 수
}

// FinancialEstimates 는 종목의 애널리스트 재무 추정을 조회한다.
// period 는 "annual" 또는 "quarter". page 는 0부터.
func (c *Client) FinancialEstimates(ctx context.Context, symbol, period string, page int) ([]FinancialEstimate, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	if strings.TrimSpace(period) == "" {
		return nil, fmt.Errorf("fmp: period must not be empty (annual|quarter)")
	}
	return fetch.List[FinancialEstimate](ctx, c.http, "/stable/analyst-estimates", map[string]string{
		"symbol": symbol,
		"period": period,
		"page":   strconv.Itoa(page),
	})
}
