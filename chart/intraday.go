package chart

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// IntradayBar — 분/시간봉 (historical-chart/{interval} 6종 공유). symbol 없음.
type IntradayBar struct {
	Date   string  `json:"date"`   // 일시
	Open   float64 `json:"open"`   // 시가
	Low    float64 `json:"low"`    // 저가
	High   float64 `json:"high"`   // 고가
	Close  float64 `json:"close"`  // 종가
	Volume int64   `json:"volume"` // 거래량
}

// Intraday1Min 은 1분봉을 조회한다.
func (c *Client) Intraday1Min(ctx context.Context, symbol, from, to string, nonadjusted bool) ([]IntradayBar, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[IntradayBar](ctx, c.http, "/stable/historical-chart/1min", intradayParams(symbol, from, to, nonadjusted))
}

// Intraday5Min 은 5분봉을 조회한다.
func (c *Client) Intraday5Min(ctx context.Context, symbol, from, to string, nonadjusted bool) ([]IntradayBar, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[IntradayBar](ctx, c.http, "/stable/historical-chart/5min", intradayParams(symbol, from, to, nonadjusted))
}

// Intraday15Min 은 15분봉을 조회한다.
func (c *Client) Intraday15Min(ctx context.Context, symbol, from, to string, nonadjusted bool) ([]IntradayBar, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[IntradayBar](ctx, c.http, "/stable/historical-chart/15min", intradayParams(symbol, from, to, nonadjusted))
}

// Intraday30Min 은 30분봉을 조회한다.
func (c *Client) Intraday30Min(ctx context.Context, symbol, from, to string, nonadjusted bool) ([]IntradayBar, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[IntradayBar](ctx, c.http, "/stable/historical-chart/30min", intradayParams(symbol, from, to, nonadjusted))
}

// Intraday1Hour 는 1시간봉을 조회한다.
func (c *Client) Intraday1Hour(ctx context.Context, symbol, from, to string, nonadjusted bool) ([]IntradayBar, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[IntradayBar](ctx, c.http, "/stable/historical-chart/1hour", intradayParams(symbol, from, to, nonadjusted))
}

// Intraday4Hour 는 4시간봉을 조회한다.
func (c *Client) Intraday4Hour(ctx context.Context, symbol, from, to string, nonadjusted bool) ([]IntradayBar, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[IntradayBar](ctx, c.http, "/stable/historical-chart/4hour", intradayParams(symbol, from, to, nonadjusted))
}
