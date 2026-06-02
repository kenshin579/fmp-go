package company

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Profile 은 FMP /stable/profile 응답의 단일 종목 프로필이다.
type Profile struct {
	Symbol            string  `json:"symbol"`
	CompanyName       string  `json:"companyName"`
	Price             float64 `json:"price"`
	MarketCap         int64   `json:"marketCap"`
	Beta              float64 `json:"beta"`
	LastDividend      float64 `json:"lastDividend"`
	Range             string  `json:"range"`
	Change            float64 `json:"change"`
	ChangePercentage  float64 `json:"changePercentage"`
	Volume            int64   `json:"volume"`
	AverageVolume     int64   `json:"averageVolume"`
	Currency          string  `json:"currency"`
	CIK               string  `json:"cik"`
	ISIN              string  `json:"isin"`
	CUSIP             string  `json:"cusip"`
	Exchange          string  `json:"exchange"`
	ExchangeFullName  string  `json:"exchangeFullName"`
	Industry          string  `json:"industry"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	CEO               string  `json:"ceo"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	Image             string  `json:"image"`
	IPODate           string  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsAdr             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
}

// Profile 은 종목의 회사 프로필을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) Profile(ctx context.Context, symbol string) (*Profile, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []Profile
	if err := c.http.GetJSON(ctx, "/stable/profile", map[string]string{"symbol": symbol}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return &out[0], nil
}

// ProfileByCIK 는 CIK 로 회사 프로필을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) ProfileByCIK(ctx context.Context, cik string) (*Profile, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.One[Profile](ctx, c.http, "/stable/profile-cik", map[string]string{"cik": cik})
}
