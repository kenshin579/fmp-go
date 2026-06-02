// Package form13f 는 FMP 기관 13F 보유/분석 API sub-client.
// fmp.Client.Form13F 로 접근.
package form13f

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 13F sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

func pageParams(page, limit int) map[string]string {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
