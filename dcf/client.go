// Package dcf 는 FMP DCF 밸류에이션 API sub-client.
// fmp.Client.DCF 로 접근.
package dcf

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 DCF sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
