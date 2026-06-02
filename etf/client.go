// Package etf 는 FMP ETF/뮤추얼펀드 API sub-client.
// fmp.Client.ETF 로 접근.
package etf

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 ETF/펀드 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
