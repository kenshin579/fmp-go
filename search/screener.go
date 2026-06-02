package search

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// ScreenerParams — company-screener 필터. 빈 값/0/nil 은 쿼리에서 생략.
// 숫자 0 = 미지정(생략). boolean 은 *bool(false 와 미지정 구분).
type ScreenerParams struct {
	MarketCapMoreThan      int64
	MarketCapLowerThan     int64
	Sector                 string
	Industry               string
	BetaMoreThan           float64
	BetaLowerThan          float64
	PriceMoreThan          float64
	PriceLowerThan         float64
	DividendMoreThan       float64
	DividendLowerThan      float64
	VolumeMoreThan         int64
	VolumeLowerThan        int64
	Exchange               string
	Country                string
	IsEtf                  *bool
	IsFund                 *bool
	IsActivelyTrading      *bool
	Limit                  int
	IncludeAllShareClasses *bool
}

// toMap 은 비제로/non-nil 필터만 쿼리 맵으로 만든다.
func (p ScreenerParams) toMap() map[string]string {
	m := map[string]string{}
	putI := func(k string, v int64) {
		if v != 0 {
			m[k] = strconv.FormatInt(v, 10)
		}
	}
	putF := func(k string, v float64) {
		if v != 0 {
			m[k] = strconv.FormatFloat(v, 'f', -1, 64)
		}
	}
	putS := func(k, v string) {
		if v != "" {
			m[k] = v
		}
	}
	putB := func(k string, v *bool) {
		if v != nil {
			m[k] = strconv.FormatBool(*v)
		}
	}
	putI("marketCapMoreThan", p.MarketCapMoreThan)
	putI("marketCapLowerThan", p.MarketCapLowerThan)
	putS("sector", p.Sector)
	putS("industry", p.Industry)
	putF("betaMoreThan", p.BetaMoreThan)
	putF("betaLowerThan", p.BetaLowerThan)
	putF("priceMoreThan", p.PriceMoreThan)
	putF("priceLowerThan", p.PriceLowerThan)
	putF("dividendMoreThan", p.DividendMoreThan)
	putF("dividendLowerThan", p.DividendLowerThan)
	putI("volumeMoreThan", p.VolumeMoreThan)
	putI("volumeLowerThan", p.VolumeLowerThan)
	putS("exchange", p.Exchange)
	putS("country", p.Country)
	putB("isEtf", p.IsEtf)
	putB("isFund", p.IsFund)
	putB("isActivelyTrading", p.IsActivelyTrading)
	if p.Limit > 0 {
		m["limit"] = strconv.Itoa(p.Limit)
	}
	putB("includeAllShareClasses", p.IncludeAllShareClasses)
	return m
}

// ScreenerResult — 스크리너 결과. nullable(MarketCap/Beta/LastAnnualDividend) → 포인터.
type ScreenerResult struct {
	Symbol             string   `json:"symbol"`             // 종목 심볼
	CompanyName        string   `json:"companyName"`        // 회사명
	MarketCap          *int64   `json:"marketCap"`          // 시가총액(결측 가능)
	Sector             string   `json:"sector"`             // 섹터
	Industry           string   `json:"industry"`           // 산업
	Beta               *float64 `json:"beta"`               // 베타(결측 가능)
	Price              float64  `json:"price"`              // 현재가
	LastAnnualDividend *float64 `json:"lastAnnualDividend"` // 최근 연간 배당(결측 가능)
	Volume             int64    `json:"volume"`             // 거래량
	Exchange           string   `json:"exchange"`           // 거래소
	ExchangeShortName  string   `json:"exchangeShortName"`  // 거래소 약칭
	Country            string   `json:"country"`            // 국가
	IsEtf              bool     `json:"isEtf"`              // ETF 여부
	IsFund             bool     `json:"isFund"`             // 펀드 여부
	IsActivelyTrading  bool     `json:"isActivelyTrading"`  // 거래 활성 여부
}

// CompanyScreener 는 다중 필터로 종목을 스크리닝한다. 빈 params 는 전체 스크리닝.
func (c *Client) CompanyScreener(ctx context.Context, params ScreenerParams) ([]ScreenerResult, error) {
	return fetch.List[ScreenerResult](ctx, c.http, "/stable/company-screener", params.toMap())
}
