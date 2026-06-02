// Package transcripts 는 FMP 실적 발표 트랜스크립트 API sub-client.
// fmp.Client.EarningsTranscripts 로 접근.
package transcripts

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 트랜스크립트 sub-client.
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
