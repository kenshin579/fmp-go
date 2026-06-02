package etf

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// ETFHolding — ETF 보유 종목 (etf/holdings)
type ETFHolding struct {
	Symbol           string  `json:"symbol"`
	Asset            string  `json:"asset"`
	Name             string  `json:"name"`
	ISIN             string  `json:"isin"`
	SecurityCusip    string  `json:"securityCusip"`
	SharesNumber     int64   `json:"sharesNumber"`
	WeightPercentage float64 `json:"weightPercentage"`
	MarketValue      float64 `json:"marketValue"`
	UpdatedAt        string  `json:"updatedAt"`
	Updated          string  `json:"updated"`
}

// ETFSectorExposure — ETF info 의 sectorsList 중첩 원소.
type ETFSectorExposure struct {
	Industry string  `json:"industry"`
	Exposure float64 `json:"exposure"`
}

// ETFInformation — ETF/펀드 프로필 (etf/info)
type ETFInformation struct {
	Symbol                string              `json:"symbol"`
	Name                  string              `json:"name"`
	Description           string              `json:"description"`
	ISIN                  string              `json:"isin"`
	AssetClass            string              `json:"assetClass"`
	SecurityCusip         string              `json:"securityCusip"`
	Domicile              string              `json:"domicile"`
	Website               string              `json:"website"`
	ETFCompany            string              `json:"etfCompany"`
	ExpenseRatio          float64             `json:"expenseRatio"`
	AssetsUnderManagement int64               `json:"assetsUnderManagement"`
	AvgVolume             int64               `json:"avgVolume"`
	InceptionDate         string              `json:"inceptionDate"`
	NAV                   float64             `json:"nav"`
	NAVCurrency           string              `json:"navCurrency"`
	HoldingsCount         int                 `json:"holdingsCount"`
	UpdatedAt             string              `json:"updatedAt"`
	SectorsList           []ETFSectorExposure `json:"sectorsList"`
}

// ETFCountryWeighting — ETF 국가 비중 (etf/country-weightings). weightPercentage 는 "97.29%" 문자열.
type ETFCountryWeighting struct {
	Country          string `json:"country"`
	WeightPercentage string `json:"weightPercentage"`
}

// ETFSectorWeighting — ETF 섹터 비중 (etf/sector-weightings). weightPercentage 는 숫자.
type ETFSectorWeighting struct {
	Symbol           string  `json:"symbol"`
	Sector           string  `json:"sector"`
	WeightPercentage float64 `json:"weightPercentage"`
}

// ETFAssetExposure — 특정 종목을 보유한 ETF 노출 (etf/asset-exposure)
type ETFAssetExposure struct {
	Symbol           string  `json:"symbol"`
	Asset            string  `json:"asset"`
	SharesNumber     int64   `json:"sharesNumber"`
	WeightPercentage float64 `json:"weightPercentage"`
	MarketValue      float64 `json:"marketValue"`
}

func (c *Client) Holdings(ctx context.Context, symbol string) ([]ETFHolding, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[ETFHolding](ctx, c.http, "/stable/etf/holdings", map[string]string{"symbol": symbol})
}
func (c *Client) Information(ctx context.Context, symbol string) ([]ETFInformation, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[ETFInformation](ctx, c.http, "/stable/etf/info", map[string]string{"symbol": symbol})
}
func (c *Client) CountryWeightings(ctx context.Context, symbol string) ([]ETFCountryWeighting, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[ETFCountryWeighting](ctx, c.http, "/stable/etf/country-weightings", map[string]string{"symbol": symbol})
}
func (c *Client) SectorWeightings(ctx context.Context, symbol string) ([]ETFSectorWeighting, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[ETFSectorWeighting](ctx, c.http, "/stable/etf/sector-weightings", map[string]string{"symbol": symbol})
}
func (c *Client) AssetExposure(ctx context.Context, symbol string) ([]ETFAssetExposure, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[ETFAssetExposure](ctx, c.http, "/stable/etf/asset-exposure", map[string]string{"symbol": symbol})
}
