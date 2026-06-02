// Package search 는 FMP 검색(search) API sub-client.
// fmp.Client.Search 로 접근.
package search

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 검색 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
