package analyst

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// PriceTargetConsensus — 목표주가 컨센서스 (price-target-consensus)
type PriceTargetConsensus struct {
	Symbol          string  `json:"symbol"`          // 종목 심볼
	TargetHigh      float64 `json:"targetHigh"`      // 최고 목표가
	TargetLow       float64 `json:"targetLow"`       // 최저 목표가
	TargetConsensus float64 `json:"targetConsensus"` // 평균 목표가
	TargetMedian    float64 `json:"targetMedian"`    // 중앙값 목표가
}

// PriceTargetSummary — 목표주가 요약 (price-target-summary)
type PriceTargetSummary struct {
	Symbol                    string  `json:"symbol"`                    // 종목 심볼
	LastMonthCount            int     `json:"lastMonthCount"`            // 최근 1개월 리포트 수
	LastMonthAvgPriceTarget   float64 `json:"lastMonthAvgPriceTarget"`   // 최근 1개월 평균 목표가
	LastQuarterCount          int     `json:"lastQuarterCount"`          // 최근 분기 리포트 수
	LastQuarterAvgPriceTarget float64 `json:"lastQuarterAvgPriceTarget"` // 최근 분기 평균 목표가
	LastYearCount             int     `json:"lastYearCount"`             // 최근 1년 리포트 수
	LastYearAvgPriceTarget    float64 `json:"lastYearAvgPriceTarget"`    // 최근 1년 평균 목표가
	AllTimeCount              int     `json:"allTimeCount"`              // 전체 리포트 수
	AllTimeAvgPriceTarget     float64 `json:"allTimeAvgPriceTarget"`     // 전체 평균 목표가
	Publishers                string  `json:"publishers"`                // 발행처 목록(JSON 배열 문자열)
}

// PriceTargetConsensus 는 종목의 목표주가 컨센서스를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) PriceTargetConsensus(ctx context.Context, symbol string) (*PriceTargetConsensus, error) {
	return fetch.OneBySymbol[PriceTargetConsensus](ctx, c.http, "/stable/price-target-consensus", symbol)
}

// PriceTargetSummary 는 종목의 목표주가 요약(기간별 평균)을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) PriceTargetSummary(ctx context.Context, symbol string) (*PriceTargetSummary, error) {
	return fetch.OneBySymbol[PriceTargetSummary](ctx, c.http, "/stable/price-target-summary", symbol)
}
