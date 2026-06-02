// Package markethours 는 FMP 거래소 운영시간/휴장일 API sub-client.
// fmp.Client.MarketHours 로 접근.
package markethours

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 거래소 운영시간 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
