package metrics

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// EnterpriseValue — 기업가치 (enterprise-values)
type EnterpriseValue struct {
	Symbol                      string  `json:"symbol"`                      // 종목 심볼
	Date                        string  `json:"date"`                        // 기준일
	StockPrice                  float64 `json:"stockPrice"`                  // 주가
	NumberOfShares              int64   `json:"numberOfShares"`              // 발행주식수
	MarketCapitalization        int64   `json:"marketCapitalization"`        // 시가총액
	MinusCashAndCashEquivalents int64   `json:"minusCashAndCashEquivalents"` // (-)현금성자산
	AddTotalDebt                int64   `json:"addTotalDebt"`                // (+)총부채
	EnterpriseValue             int64   `json:"enterpriseValue"`             // 기업가치(EV)
}

// EnterpriseValues 는 종목의 기업가치 시계열을 조회한다.
func (c *Client) EnterpriseValues(ctx context.Context, symbol, period string, limit int) ([]EnterpriseValue, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EnterpriseValue](ctx, c.http, "/stable/enterprise-values", listParams(symbol, period, limit))
}
