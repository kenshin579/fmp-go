// Package economics 는 FMP 경제 API sub-client (국채/지표/캘린더/리스크프리미엄).
// fmp.Client.Economics 로 접근.
package economics

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 경제 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// fromToParams 는 from/to(비어있지 않으면) 쿼리 맵.
func fromToParams(from, to string) map[string]string {
	q := map[string]string{}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return q
}
