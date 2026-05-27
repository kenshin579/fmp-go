// Package company 는 FMP 회사 정보 API sub-client 다. fmp.Client.Company 로 접근한다.
package company

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 회사 정보 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
