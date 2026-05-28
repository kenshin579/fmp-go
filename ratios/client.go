// Package ratios 는 FMP 재무비율 API sub-client.
// fmp.Client.Ratios 로 접근.
package ratios

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 재무비율 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// Params 는 재무비율 조회 파라미터.
type Params struct {
	Symbol string
	Period string // "annual" | "quarter"
	Limit  int
}

func (p Params) queryParams() map[string]string {
	q := map[string]string{"symbol": p.Symbol}
	if p.Period != "" {
		q["period"] = p.Period
	}
	if p.Limit > 0 {
		q["limit"] = strconv.Itoa(p.Limit)
	}
	return q
}
