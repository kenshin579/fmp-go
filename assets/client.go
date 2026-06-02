// Package assets 는 FMP 암호화폐/외환/원자재 목록 API sub-client.
// fmp.Client.Assets 로 접근. (시세/시계열은 client.Quote, client.Chart 사용)
package assets

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 자산 목록 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
