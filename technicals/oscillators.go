package technicals

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// RSI — 상대강도지수 (technical-indicators/rsi)
type RSI struct {
	Bar
	RSI float64 `json:"rsi"` // 상대강도지수
}

// StandardDeviation — 표준편차 (technical-indicators/standarddeviation)
type StandardDeviation struct {
	Bar
	StandardDeviation float64 `json:"standardDeviation"` // 표준편차
}

// Williams — 윌리엄스 %R (technical-indicators/williams)
type Williams struct {
	Bar
	Williams float64 `json:"williams"` // 윌리엄스 %R
}

// ADX — 평균방향지수 (technical-indicators/adx)
type ADX struct {
	Bar
	ADX float64 `json:"adx"` // 평균방향지수
}

// RSI 는 상대강도지수를 조회한다.
func (c *Client) RSI(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]RSI, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[RSI](ctx, c.http, "/stable/technical-indicators/rsi", indicatorParams(symbol, periodLength, timeframe, from, to))
}

// StandardDeviation 은 표준편차를 조회한다. (path 소문자 standarddeviation, JSON 키 standardDeviation)
func (c *Client) StandardDeviation(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]StandardDeviation, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[StandardDeviation](ctx, c.http, "/stable/technical-indicators/standarddeviation", indicatorParams(symbol, periodLength, timeframe, from, to))
}

// Williams 는 윌리엄스 %R 을 조회한다.
func (c *Client) Williams(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]Williams, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[Williams](ctx, c.http, "/stable/technical-indicators/williams", indicatorParams(symbol, periodLength, timeframe, from, to))
}

// ADX 는 평균방향지수를 조회한다.
func (c *Client) ADX(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]ADX, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[ADX](ctx, c.http, "/stable/technical-indicators/adx", indicatorParams(symbol, periodLength, timeframe, from, to))
}
