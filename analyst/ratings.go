package analyst

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Rating — 종합 평가 점수 (ratings-snapshot / ratings-historical 공용). snapshot 은 Date "".
type Rating struct {
	Symbol                  string `json:"symbol"`                  // 종목 심볼
	Date                    string `json:"date"`                    // 기준일(snapshot 은 빈 문자열)
	Rating                  string `json:"rating"`                  // 등급 (예: A-)
	OverallScore            int    `json:"overallScore"`            // 종합 점수
	DiscountedCashFlowScore int    `json:"discountedCashFlowScore"` // DCF 점수
	ReturnOnEquityScore     int    `json:"returnOnEquityScore"`     // ROE 점수
	ReturnOnAssetsScore     int    `json:"returnOnAssetsScore"`     // ROA 점수
	DebtToEquityScore       int    `json:"debtToEquityScore"`       // 부채비율 점수
	PriceToEarningsScore    int    `json:"priceToEarningsScore"`    // PER 점수
	PriceToBookScore        int    `json:"priceToBookScore"`        // PBR 점수
}

// RatingsSnapshot 은 종목의 현재 종합 평가 점수를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) RatingsSnapshot(ctx context.Context, symbol string) (*Rating, error) {
	return fetch.OneBySymbol[Rating](ctx, c.http, "/stable/ratings-snapshot", symbol)
}

// HistoricalRatings 는 종목의 평가 점수 시계열을 조회한다.
func (c *Client) HistoricalRatings(ctx context.Context, symbol string) ([]Rating, error) {
	return fetch.ListBySymbol[Rating](ctx, c.http, "/stable/ratings-historical", symbol)
}
