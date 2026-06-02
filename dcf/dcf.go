package dcf

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// DCFValue — 단순 DCF 내재가치 (discounted-cash-flow / levered-discounted-cash-flow 공유)
type DCFValue struct {
	Symbol     string  `json:"symbol"`      // 종목 심볼
	Date       string  `json:"date"`        // 기준일
	DCF        float64 `json:"dcf"`         // 주당 내재가치
	StockPrice float64 `json:"Stock Price"` // 현재가(FMP 키 "Stock Price")
}

// DiscountedCashFlow 는 종목의 단순 DCF 내재가치를 조회한다.
func (c *Client) DiscountedCashFlow(ctx context.Context, symbol string) ([]DCFValue, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[DCFValue](ctx, c.http, "/stable/discounted-cash-flow", map[string]string{"symbol": symbol})
}

// LeveredDiscountedCashFlow 는 종목의 레버드 DCF 내재가치를 조회한다.
func (c *Client) LeveredDiscountedCashFlow(ctx context.Context, symbol string) ([]DCFValue, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[DCFValue](ctx, c.http, "/stable/levered-discounted-cash-flow", map[string]string{"symbol": symbol})
}
