package directory

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Exchange — 거래소 목록 (available-exchanges)
type Exchange struct {
	Exchange     string `json:"exchange"`     // 거래소 코드
	Name         string `json:"name"`         // 거래소명
	CountryName  string `json:"countryName"`  // 국가명
	CountryCode  string `json:"countryCode"`  // 국가 코드
	SymbolSuffix string `json:"symbolSuffix"` // 심볼 접미사
	Delay        string `json:"delay"`        // 시세 지연(Real-time 등)
}

// Sector — 섹터 목록 (available-sectors)
type Sector struct {
	Sector string `json:"sector"` // 섹터명
}

// Industry — 산업 목록 (available-industries)
type Industry struct {
	Industry string `json:"industry"` // 산업명
}

// Country — 국가 목록 (available-countries)
type Country struct {
	Country string `json:"country"` // 국가 코드
}

// AvailableExchanges 는 지원 거래소 목록을 조회한다.
func (c *Client) AvailableExchanges(ctx context.Context) ([]Exchange, error) {
	return fetch.List[Exchange](ctx, c.http, "/stable/available-exchanges", nil)
}

// AvailableSectors 는 지원 섹터 목록을 조회한다.
func (c *Client) AvailableSectors(ctx context.Context) ([]Sector, error) {
	return fetch.List[Sector](ctx, c.http, "/stable/available-sectors", nil)
}

// AvailableIndustries 는 지원 산업 목록을 조회한다.
func (c *Client) AvailableIndustries(ctx context.Context) ([]Industry, error) {
	return fetch.List[Industry](ctx, c.http, "/stable/available-industries", nil)
}

// AvailableCountries 는 지원 국가 목록을 조회한다.
func (c *Client) AvailableCountries(ctx context.Context) ([]Country, error) {
	return fetch.List[Country](ctx, c.http, "/stable/available-countries", nil)
}
