// Package quote 는 FMP 시세(quote) API sub-client.
// fmp.Client.Quote 로 접근.
package quote

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 시세 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
