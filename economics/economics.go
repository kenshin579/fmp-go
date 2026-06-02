package economics

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// TreasuryRate — 미국 국채 수익률 곡선 (treasury-rates)
type TreasuryRate struct {
	Date   string  `json:"date"`   // 일자
	Month1 float64 `json:"month1"` // 1개월
	Month2 float64 `json:"month2"` // 2개월
	Month3 float64 `json:"month3"` // 3개월
	Month6 float64 `json:"month6"` // 6개월
	Year1  float64 `json:"year1"`  // 1년
	Year2  float64 `json:"year2"`  // 2년
	Year3  float64 `json:"year3"`  // 3년
	Year5  float64 `json:"year5"`  // 5년
	Year7  float64 `json:"year7"`  // 7년
	Year10 float64 `json:"year10"` // 10년
	Year20 float64 `json:"year20"` // 20년
	Year30 float64 `json:"year30"` // 30년
}

// EconomicIndicator — 경제 지표 (economic-indicators). name 으로 지표 선택.
type EconomicIndicator struct {
	Name  string  `json:"name"`  // 지표명
	Date  string  `json:"date"`  // 일자
	Value float64 `json:"value"` // 값
}

// EconomicCalendarEvent — 경제 캘린더 이벤트 (economic-calendar)
type EconomicCalendarEvent struct {
	Date             string   `json:"date"`             // 일시(YYYY-MM-DD HH:MM:SS)
	Country          string   `json:"country"`          // 국가
	Event            string   `json:"event"`            // 이벤트명
	Currency         string   `json:"currency"`         // 통화
	Previous         float64  `json:"previous"`         // 이전값
	Estimate         *float64 `json:"estimate"`         // 예상치(null 가능)
	Actual           float64  `json:"actual"`           // 실제값
	Change           float64  `json:"change"`           // 변동
	Impact           string   `json:"impact"`           // 영향도(Low/Medium/High)
	ChangePercentage float64  `json:"changePercentage"` // 변동률
	Unit             *string  `json:"unit"`             // 단위(null 가능)
}

// RiskPremium — 국가별 리스크 프리미엄 (market-risk-premium)
type RiskPremium struct {
	Country                string  `json:"country"`                // 국가
	Continent              string  `json:"continent"`              // 대륙
	CountryRiskPremium     float64 `json:"countryRiskPremium"`     // 국가 리스크 프리미엄
	TotalEquityRiskPremium float64 `json:"totalEquityRiskPremium"` // 총 주식 리스크 프리미엄
}

// TreasuryRates 는 미국 국채 수익률 곡선을 조회한다.
func (c *Client) TreasuryRates(ctx context.Context, from, to string) ([]TreasuryRate, error) {
	return fetch.List[TreasuryRate](ctx, c.http, "/stable/treasury-rates", fromToParams(from, to))
}

// EconomicIndicators 는 지정 경제 지표 시계열을 조회한다. name 필수(예: GDP, CPI).
func (c *Client) EconomicIndicators(ctx context.Context, name, from, to string) ([]EconomicIndicator, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	q := fromToParams(from, to)
	q["name"] = name
	return fetch.List[EconomicIndicator](ctx, c.http, "/stable/economic-indicators", q)
}

// EconomicCalendar 는 경제 캘린더 이벤트를 조회한다.
func (c *Client) EconomicCalendar(ctx context.Context, country, from, to string) ([]EconomicCalendarEvent, error) {
	q := fromToParams(from, to)
	if country != "" {
		q["country"] = country
	}
	return fetch.List[EconomicCalendarEvent](ctx, c.http, "/stable/economic-calendar", q)
}

// MarketRiskPremium 은 국가별 리스크 프리미엄을 조회한다.
func (c *Client) MarketRiskPremium(ctx context.Context) ([]RiskPremium, error) {
	return fetch.List[RiskPremium](ctx, c.http, "/stable/market-risk-premium", nil)
}
