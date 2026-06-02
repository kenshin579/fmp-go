// Package insidertrades 는 FMP 내부자 거래 API sub-client.
// fmp.Client.InsiderTrades 로 접근.
package insidertrades

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 내부자 거래 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
