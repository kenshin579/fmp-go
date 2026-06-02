// Package esg 는 FMP ESG 평가/공시/벤치마크 API sub-client.
// fmp.Client.ESG 로 접근.
package esg

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 ESG sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
