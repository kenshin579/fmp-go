package assets

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CryptoListItem — 암호화폐 목록 (cryptocurrency-list)
type CryptoListItem struct {
	Symbol            string   `json:"symbol"`
	Name              string   `json:"name"`
	Exchange          string   `json:"exchange"`
	IcoDate           string   `json:"icoDate"`
	CirculatingSupply float64  `json:"circulatingSupply"`
	TotalSupply       *float64 `json:"totalSupply"`
}

// ForexPair — 외환 페어 목록 (forex-list)
type ForexPair struct {
	Symbol       string `json:"symbol"`
	FromCurrency string `json:"fromCurrency"`
	ToCurrency   string `json:"toCurrency"`
	FromName     string `json:"fromName"`
	ToName       string `json:"toName"`
}

// CommodityListItem — 원자재 목록 (commodities-list)
type CommodityListItem struct {
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Exchange   *string `json:"exchange"`
	TradeMonth string  `json:"tradeMonth"`
	Currency   string  `json:"currency"`
}

// CryptoList 는 전체 암호화폐 목록을 조회한다.
func (c *Client) CryptoList(ctx context.Context) ([]CryptoListItem, error) {
	return fetch.List[CryptoListItem](ctx, c.http, "/stable/cryptocurrency-list", nil)
}

// ForexList 는 전체 외환 페어 목록을 조회한다.
func (c *Client) ForexList(ctx context.Context) ([]ForexPair, error) {
	return fetch.List[ForexPair](ctx, c.http, "/stable/forex-list", nil)
}

// CommodityList 는 전체 원자재 목록을 조회한다.
func (c *Client) CommodityList(ctx context.Context) ([]CommodityListItem, error) {
	return fetch.List[CommodityListItem](ctx, c.http, "/stable/commodities-list", nil)
}
