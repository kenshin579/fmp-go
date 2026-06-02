package technicals

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Bar — 기술지표 응답 공통 OHLCV 블록(임베디드).
type Bar struct {
	Date   string  `json:"date"`   // 일시(YYYY-MM-DD HH:MM:SS)
	Open   float64 `json:"open"`   // 시가
	High   float64 `json:"high"`   // 고가
	Low    float64 `json:"low"`    // 저가
	Close  float64 `json:"close"`  // 종가
	Volume int64   `json:"volume"` // 거래량
}

// SMA — 단순 이동평균 (technical-indicators/sma)
type SMA struct {
	Bar
	SMA float64 `json:"sma"` // 단순 이동평균
}

// EMA — 지수 이동평균 (technical-indicators/ema)
type EMA struct {
	Bar
	EMA float64 `json:"ema"` // 지수 이동평균
}

// WMA — 가중 이동평균 (technical-indicators/wma)
type WMA struct {
	Bar
	WMA float64 `json:"wma"` // 가중 이동평균
}

// DEMA — 이중 지수 이동평균 (technical-indicators/dema)
type DEMA struct {
	Bar
	DEMA float64 `json:"dema"` // 이중 지수 이동평균
}

// TEMA — 삼중 지수 이동평균 (technical-indicators/tema)
type TEMA struct {
	Bar
	TEMA float64 `json:"tema"` // 삼중 지수 이동평균
}

// SMA 는 단순 이동평균을 조회한다.
func (c *Client) SMA(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]SMA, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[SMA](ctx, c.http, "/stable/technical-indicators/sma", indicatorParams(symbol, periodLength, timeframe, from, to))
}

// EMA 는 지수 이동평균을 조회한다.
func (c *Client) EMA(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]EMA, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[EMA](ctx, c.http, "/stable/technical-indicators/ema", indicatorParams(symbol, periodLength, timeframe, from, to))
}

// WMA 는 가중 이동평균을 조회한다.
func (c *Client) WMA(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]WMA, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[WMA](ctx, c.http, "/stable/technical-indicators/wma", indicatorParams(symbol, periodLength, timeframe, from, to))
}

// DEMA 는 이중 지수 이동평균을 조회한다.
func (c *Client) DEMA(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]DEMA, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[DEMA](ctx, c.http, "/stable/technical-indicators/dema", indicatorParams(symbol, periodLength, timeframe, from, to))
}

// TEMA 는 삼중 지수 이동평균을 조회한다.
func (c *Client) TEMA(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]TEMA, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[TEMA](ctx, c.http, "/stable/technical-indicators/tema", indicatorParams(symbol, periodLength, timeframe, from, to))
}
