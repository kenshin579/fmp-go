// Package reports 는 FMP 보고서 API sub-client (as-reported/latest/dates/10-K).
// fmp.Client.Reports 로 접근.
package reports

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 보고서 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// asReportedParams 는 symbol(필수) + period(비어있지 않으면) + limit(>0) 쿼리 맵.
func asReportedParams(symbol, period string, limit int) map[string]string {
	q := map[string]string{"symbol": symbol}
	if period != "" {
		q["period"] = period
	}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
