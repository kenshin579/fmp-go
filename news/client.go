// Package news 는 FMP 뉴스(news) API sub-client.
// fmp.Client.News 로 접근.
package news

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 뉴스 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
