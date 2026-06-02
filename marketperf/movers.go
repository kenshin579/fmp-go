package marketperf

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// MarketMover — 등락 상위 종목 (biggest-gainers/losers/most-actives 공유)
type MarketMover struct {
	Symbol            string  `json:"symbol"`            // 종목 심볼
	Price             float64 `json:"price"`             // 현재가
	Name              string  `json:"name"`              // 종목명
	Change            float64 `json:"change"`            // 변동액
	ChangesPercentage float64 `json:"changesPercentage"` // 변동률(%) — FMP 키 's' 포함
	Exchange          string  `json:"exchange"`          // 거래소
}

// BiggestGainers 는 상승률 상위 종목을 조회한다.
func (c *Client) BiggestGainers(ctx context.Context) ([]MarketMover, error) {
	return fetch.List[MarketMover](ctx, c.http, "/stable/biggest-gainers", nil)
}

// BiggestLosers 는 하락률 상위 종목을 조회한다.
func (c *Client) BiggestLosers(ctx context.Context) ([]MarketMover, error) {
	return fetch.List[MarketMover](ctx, c.http, "/stable/biggest-losers", nil)
}

// MostActives 는 거래 활발 종목을 조회한다.
func (c *Client) MostActives(ctx context.Context) ([]MarketMover, error) {
	return fetch.List[MarketMover](ctx, c.http, "/stable/most-actives", nil)
}
