package metrics

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// RevenueSegment — 매출 세그먼트 (지역/제품 공용). data 는 세그먼트명→매출 동적 맵.
type RevenueSegment struct {
	Symbol           string           `json:"symbol"`           // 종목 심볼
	FiscalYear       int              `json:"fiscalYear"`       // 회계연도(숫자)
	Period           string           `json:"period"`           // 기간 (FY/Q1..)
	ReportedCurrency *string          `json:"reportedCurrency"` // 보고 통화(null 가능)
	Date             string           `json:"date"`             // 기준일
	Data             map[string]int64 `json:"data"`             // 세그먼트명 → 매출액
}

// RevenueGeographicSegmentation 은 종목의 지역별 매출 세그먼트를 조회한다.
func (c *Client) RevenueGeographicSegmentation(ctx context.Context, symbol, period string) ([]RevenueSegment, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[RevenueSegment](ctx, c.http, "/stable/revenue-geographic-segmentation", listParams(symbol, period, 0))
}

// RevenueProductSegmentation 은 종목의 제품별 매출 세그먼트를 조회한다.
func (c *Client) RevenueProductSegmentation(ctx context.Context, symbol, period string) ([]RevenueSegment, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[RevenueSegment](ctx, c.http, "/stable/revenue-product-segmentation", listParams(symbol, period, 0))
}
