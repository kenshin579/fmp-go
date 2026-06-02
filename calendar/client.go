// Package calendar 는 FMP 캘린더(calendar) API sub-client.
// fmp.Client.Calendar 로 접근.
package calendar

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 캘린더 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// dateRange 는 from/to(YYYY-MM-DD) 를 쿼리 맵으로 만든다. 빈 값은 제외.
func dateRange(from, to string) map[string]string {
	m := map[string]string{}
	if from != "" {
		m["from"] = from
	}
	if to != "" {
		m["to"] = to
	}
	return m
}
