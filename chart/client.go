// Package chart 는 FMP 과거 시세 API sub-client (EOD/intraday).
// fmp.Client.Chart 로 접근.
package chart

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 과거 시세 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// eodParams 는 symbol(필수) + from/to(비어있지 않으면) 쿼리 맵.
func eodParams(symbol, from, to string) map[string]string {
	q := map[string]string{"symbol": symbol}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return q
}

// intradayParams 는 eodParams + nonadjusted(true 일 때만).
func intradayParams(symbol, from, to string, nonadjusted bool) map[string]string {
	q := eodParams(symbol, from, to)
	if nonadjusted {
		q["nonadjusted"] = "true"
	}
	return q
}
