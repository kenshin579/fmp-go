package chart

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// EODLight — EOD 경량 시세 (historical-price-eod/light)
type EODLight struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	Date   string  `json:"date"`   // 일자
	Price  float64 `json:"price"`  // 종가
	Volume int64   `json:"volume"` // 거래량
}

// EODFull — EOD 전체 시세 (historical-price-eod/full)
type EODFull struct {
	Symbol        string  `json:"symbol"`        // 종목 심볼
	Date          string  `json:"date"`          // 일자
	Open          float64 `json:"open"`          // 시가
	High          float64 `json:"high"`          // 고가
	Low           float64 `json:"low"`           // 저가
	Close         float64 `json:"close"`         // 종가
	Volume        int64   `json:"volume"`        // 거래량
	Change        float64 `json:"change"`        // 변동액
	ChangePercent float64 `json:"changePercent"` // 변동률(%)
	VWAP          float64 `json:"vwap"`          // 거래량가중평균가
}

// EODAdjusted — 조정 시세 (dividend-adjusted / non-split-adjusted 공유)
type EODAdjusted struct {
	Symbol   string  `json:"symbol"`   // 종목 심볼
	Date     string  `json:"date"`     // 일자
	AdjOpen  float64 `json:"adjOpen"`  // 조정 시가
	AdjHigh  float64 `json:"adjHigh"`  // 조정 고가
	AdjLow   float64 `json:"adjLow"`   // 조정 저가
	AdjClose float64 `json:"adjClose"` // 조정 종가
	Volume   int64   `json:"volume"`   // 거래량
}

// HistoricalPriceEODLight 는 종목의 EOD 경량 시세를 조회한다.
func (c *Client) HistoricalPriceEODLight(ctx context.Context, symbol, from, to string) ([]EODLight, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EODLight](ctx, c.http, "/stable/historical-price-eod/light", eodParams(symbol, from, to))
}

// HistoricalPriceEODFull 는 종목의 EOD 전체 시세를 조회한다.
func (c *Client) HistoricalPriceEODFull(ctx context.Context, symbol, from, to string) ([]EODFull, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EODFull](ctx, c.http, "/stable/historical-price-eod/full", eodParams(symbol, from, to))
}

// HistoricalPriceEODDividendAdjusted 는 배당 조정 EOD 시세를 조회한다.
func (c *Client) HistoricalPriceEODDividendAdjusted(ctx context.Context, symbol, from, to string) ([]EODAdjusted, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EODAdjusted](ctx, c.http, "/stable/historical-price-eod/dividend-adjusted", eodParams(symbol, from, to))
}

// HistoricalPriceEODNonSplitAdjusted 는 분할 미조정 EOD 시세를 조회한다.
func (c *Client) HistoricalPriceEODNonSplitAdjusted(ctx context.Context, symbol, from, to string) ([]EODAdjusted, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EODAdjusted](ctx, c.http, "/stable/historical-price-eod/non-split-adjusted", eodParams(symbol, from, to))
}
