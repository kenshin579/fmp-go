package senate

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CongressTrade — 의회 의원 거래 공시 (senate/house latest·trades·by-name 공유).
// 모든 값이 문자열(amount 는 범위, capitalGainsOver200USD 는 "True"/"False").
type CongressTrade struct {
	Symbol                 string `json:"symbol"`                 // 종목 심볼
	DisclosureDate         string `json:"disclosureDate"`         // 공시일
	TransactionDate        string `json:"transactionDate"`        // 거래일
	FirstName              string `json:"firstName"`              // 이름
	LastName               string `json:"lastName"`               // 성
	Office                 string `json:"office"`                 // 의원명/직위
	District               string `json:"district"`               // 주(상원)/주+선거구(하원)
	Owner                  string `json:"owner"`                  // 소유자(Spouse/Joint 등)
	AssetDescription       string `json:"assetDescription"`       // 자산 설명
	AssetType              string `json:"assetType"`              // 자산 유형(Stock 등)
	Type                   string `json:"type"`                   // 거래 유형(Purchase/Sale)
	Amount                 string `json:"amount"`                 // 금액 범위("$1,001 - $15,000")
	CapitalGainsOver200USD string `json:"capitalGainsOver200USD"` // 200달러 초과 자본이득 여부(문자열)
	Comment                string `json:"comment"`                // 비고
	Link                   string `json:"link"`                   // 공시 원문 URL
}

// SenateLatest 는 최신 상원 의원 거래 공시를 조회한다.
func (c *Client) SenateLatest(ctx context.Context, page, limit int) ([]CongressTrade, error) {
	return fetch.List[CongressTrade](ctx, c.http, "/stable/senate-latest", pageParams(page, limit))
}

// SenateTrades 는 특정 종목의 상원 거래 공시를 조회한다. symbol 필수.
func (c *Client) SenateTrades(ctx context.Context, symbol string, page, limit int) ([]CongressTrade, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := pageParams(page, limit)
	q["symbol"] = symbol
	return fetch.List[CongressTrade](ctx, c.http, "/stable/senate-trades", q)
}

// SenateTradesByName 은 의원명으로 상원 거래 공시를 조회한다. name 필수.
func (c *Client) SenateTradesByName(ctx context.Context, name string) ([]CongressTrade, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[CongressTrade](ctx, c.http, "/stable/senate-trades-by-name", map[string]string{"name": name})
}

// HouseLatest 는 최신 하원 의원 거래 공시를 조회한다.
func (c *Client) HouseLatest(ctx context.Context, page, limit int) ([]CongressTrade, error) {
	return fetch.List[CongressTrade](ctx, c.http, "/stable/house-latest", pageParams(page, limit))
}

// HouseTrades 는 특정 종목의 하원 거래 공시를 조회한다. symbol 필수.
func (c *Client) HouseTrades(ctx context.Context, symbol string, page, limit int) ([]CongressTrade, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := pageParams(page, limit)
	q["symbol"] = symbol
	return fetch.List[CongressTrade](ctx, c.http, "/stable/house-trades", q)
}

// HouseTradesByName 은 의원명으로 하원 거래 공시를 조회한다. name 필수.
func (c *Client) HouseTradesByName(ctx context.Context, name string) ([]CongressTrade, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[CongressTrade](ctx, c.http, "/stable/house-trades-by-name", map[string]string{"name": name})
}
