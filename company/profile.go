package company

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Profile 은 FMP /stable/profile 응답의 단일 종목 프로필이다.
type Profile struct {
	Symbol            string  `json:"symbol"`            // 종목 심볼 (예: AAPL)
	CompanyName       string  `json:"companyName"`       // 회사명
	Price             float64 `json:"price"`             // 현재가
	MarketCap         int64   `json:"marketCap"`         // 시가총액
	Beta              float64 `json:"beta"`              // 베타 (시장 대비 변동성)
	LastDividend      float64 `json:"lastDividend"`      // 최근 배당금
	Range             string  `json:"range"`             // 52주 가격 범위
	Change            float64 `json:"change"`            // 전일 대비 등락액
	ChangePercentage  float64 `json:"changePercentage"`  // 등락률 (%)
	Volume            int64   `json:"volume"`            // 거래량
	AverageVolume     int64   `json:"averageVolume"`     // 평균 거래량
	Currency          string  `json:"currency"`          // 통화
	CIK               string  `json:"cik"`               // SEC CIK
	ISIN              string  `json:"isin"`              // ISIN 코드
	CUSIP             string  `json:"cusip"`             // CUSIP 코드
	Exchange          string  `json:"exchange"`          // 거래소 코드
	ExchangeFullName  string  `json:"exchangeFullName"`  // 거래소 전체명
	Industry          string  `json:"industry"`          // 산업
	Sector            string  `json:"sector"`            // 섹터
	Country           string  `json:"country"`           // 국가
	Website           string  `json:"website"`           // 웹사이트
	Description       string  `json:"description"`       // 회사 설명
	CEO               string  `json:"ceo"`               // CEO
	FullTimeEmployees string  `json:"fullTimeEmployees"` // 정규직 직원 수(문자열)
	Phone             string  `json:"phone"`             // 전화번호
	Address           string  `json:"address"`           // 주소
	City              string  `json:"city"`              // 도시
	State             string  `json:"state"`             // 주/도
	Zip               string  `json:"zip"`               // 우편번호
	Image             string  `json:"image"`             // 로고 이미지 URL
	IPODate           string  `json:"ipoDate"`           // 상장일
	DefaultImage      bool    `json:"defaultImage"`      // 기본 이미지 여부
	IsEtf             bool    `json:"isEtf"`             // ETF 여부
	IsActivelyTrading bool    `json:"isActivelyTrading"` // 거래 활성 여부
	IsAdr             bool    `json:"isAdr"`             // ADR 여부
	IsFund            bool    `json:"isFund"`            // 펀드 여부
}

// Profile 은 종목의 회사 프로필을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) Profile(ctx context.Context, symbol string) (*Profile, error) {
	return fetch.OneBySymbol[Profile](ctx, c.http, "/stable/profile", symbol)
}

// ProfileByCIK 는 CIK 로 회사 프로필을 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) ProfileByCIK(ctx context.Context, cik string) (*Profile, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.One[Profile](ctx, c.http, "/stable/profile-cik", map[string]string{"cik": cik})
}
