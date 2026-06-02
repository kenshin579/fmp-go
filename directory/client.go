// Package directory 는 FMP 목록 API sub-client (심볼/거래소/섹터/산업/국가).
// fmp.Client.Directory 로 접근.
package directory

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 목록 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
