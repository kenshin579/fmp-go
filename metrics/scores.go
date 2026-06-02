package metrics

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// FinancialScores — 재무 건전성 점수 (financial-scores)
type FinancialScores struct {
	Symbol           string  `json:"symbol"`           // 종목 심볼
	ReportedCurrency string  `json:"reportedCurrency"` // 보고 통화
	AltmanZScore     float64 `json:"altmanZScore"`     // Altman Z-Score
	PiotroskiScore   int     `json:"piotroskiScore"`   // Piotroski 점수(0~9)
	WorkingCapital   int64   `json:"workingCapital"`   // 운전자본
	TotalAssets      int64   `json:"totalAssets"`      // 총자산
	RetainedEarnings int64   `json:"retainedEarnings"` // 이익잉여금
	EBIT             int64   `json:"ebit"`             // EBIT
	MarketCap        int64   `json:"marketCap"`        // 시가총액
	TotalLiabilities int64   `json:"totalLiabilities"` // 총부채
	Revenue          int64   `json:"revenue"`          // 매출
}

// FinancialScores 는 종목의 재무 건전성 점수를 조회한다. 없으면 httpclient.ErrNotFound.
func (c *Client) FinancialScores(ctx context.Context, symbol string) (*FinancialScores, error) {
	return fetch.OneBySymbol[FinancialScores](ctx, c.http, "/stable/financial-scores", symbol)
}
