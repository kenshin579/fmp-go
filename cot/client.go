// Package cot 는 FMP 상품선물 COT(Commitment of Traders) 리포트 API sub-client.
// fmp.Client.COT 로 접근.
package cot

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 COT sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

func filterParams(symbol, from, to string) map[string]string {
	q := map[string]string{}
	if symbol != "" {
		q["symbol"] = symbol
	}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return q
}
