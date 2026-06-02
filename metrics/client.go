// Package metrics 는 FMP 지표 API sub-client (key-metrics/scores/owner-earnings/EV/segments).
// fmp.Client.Metrics 로 접근.
package metrics

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 지표 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// listParams 는 symbol(필수) + period(비어있지 않으면) + limit(>0) 쿼리 맵.
func listParams(symbol, period string, limit int) map[string]string {
	q := map[string]string{"symbol": symbol}
	if period != "" {
		q["period"] = period
	}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
