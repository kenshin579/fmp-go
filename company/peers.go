package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Peer — 동종업계 비교 종목
type Peer struct {
	Symbol      string  `json:"symbol"`      // 종목 심볼
	CompanyName string  `json:"companyName"` // 회사명
	Price       float64 `json:"price"`       // 현재가
	MktCap      int64   `json:"mktCap"`      // 시가총액
}

// StockPeers 는 종목의 동종업계 비교 종목을 조회한다.
func (c *Client) StockPeers(ctx context.Context, symbol string) ([]Peer, error) {
	return fetch.ListBySymbol[Peer](ctx, c.http, "/stable/stock-peers", symbol)
}
