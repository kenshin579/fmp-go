// Package marketperf 는 FMP 시장 성과 API sub-client (등락/섹터/산업/PE).
// fmp.Client.MarketPerformance 로 접근.
package marketperf

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 시장 성과 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// snapshotParams: date(필수) + exchange + dimension(sector|industry, 비어있지 않으면).
func snapshotParams(date, exchange, dimKey, dimVal string) map[string]string {
	q := map[string]string{"date": date}
	if exchange != "" {
		q["exchange"] = exchange
	}
	if dimVal != "" {
		q[dimKey] = dimVal
	}
	return q
}

// historicalParams: dimension(필수) + from + to + exchange(비어있지 않으면).
func historicalParams(dimKey, dimVal, from, to, exchange string) map[string]string {
	q := map[string]string{dimKey: dimVal}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	if exchange != "" {
		q["exchange"] = exchange
	}
	return q
}
