// Package secfilings 는 FMP SEC 공시 검색/분류/프로필 API sub-client.
// fmp.Client.SECFilings 로 접근.
package secfilings

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 SEC 공시 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// filingParams: from/to/page 항상, 주param(key!="" 일 때)/limit(>0) 조건부.
func filingParams(key, val, from, to string, page, limit int) map[string]string {
	q := map[string]string{"from": from, "to": to, "page": strconv.Itoa(page)}
	if key != "" {
		q[key] = val
	}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
