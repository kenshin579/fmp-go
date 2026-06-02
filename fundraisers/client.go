// Package fundraisers 는 FMP 크라우드펀딩/지분공모 API sub-client.
// fmp.Client.Fundraisers 로 접근.
package fundraisers

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 펀드레이저 sub-client.
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
