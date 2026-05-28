// Package statements 는 FMP 재무제표 API sub-client.
// fmp.Client.Statements 로 접근.
package statements

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 재무제표 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// Params 는 재무제표 조회 공통 파라미터.
type Params struct {
	Symbol string
	Period string // "annual" | "quarter" (빈 값 → FMP 기본 annual)
	Limit  int    // 0 → 쿼리 미포함(FMP 기본)
}

// queryParams 는 Params 를 httpclient 쿼리 맵으로 변환한다.
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
