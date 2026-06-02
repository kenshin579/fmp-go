package search

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// ExchangeVariant — 거래소별 심볼 변형 (search-exchange-variants). profile 유사.
type ExchangeVariant struct {
	Symbol            string  `json:"symbol"`            // 종목 심볼
	Price             float64 `json:"price"`             // 현재가
	Beta              float64 `json:"beta"`              // 베타
	VolAvg            int64   `json:"volAvg"`            // 평균 거래량
	MktCap            int64   `json:"mktCap"`            // 시가총액
	LastDiv           float64 `json:"lastDiv"`           // 최근 배당
	Range             string  `json:"range"`             // 52주 범위
	Changes           float64 `json:"changes"`           // 등락액
	CompanyName       string  `json:"companyName"`       // 회사명
	Currency          string  `json:"currency"`          // 통화
	CIK               string  `json:"cik"`               // SEC CIK 번호
	ISIN              string  `json:"isin"`              // ISIN 코드
	CUSIP             string  `json:"cusip"`             // CUSIP 코드
	Exchange          string  `json:"exchange"`          // 거래소 전체명
	ExchangeShortName string  `json:"exchangeShortName"` // 거래소 약칭
	Industry          string  `json:"industry"`          // 산업 분류
	Website           string  `json:"website"`           // 웹사이트 URL
	Description       string  `json:"description"`       // 회사 설명
	CEO               string  `json:"ceo"`               // 대표이사
	Sector            string  `json:"sector"`            // 섹터 분류
	Country           string  `json:"country"`           // 국가 코드
	FullTimeEmployees string  `json:"fullTimeEmployees"` // 정규직 직원 수
	Phone             string  `json:"phone"`             // 전화번호
	Address           string  `json:"address"`           // 주소
	City              string  `json:"city"`              // 도시
	State             string  `json:"state"`             // 주/도
	Zip               string  `json:"zip"`               // 우편번호
	DCFDiff           float64 `json:"dcfDiff"`           // DCF 차이
	DCF               float64 `json:"dcf"`               // DCF 내재가치
	Image             string  `json:"image"`             // 로고 이미지 URL
	IPODate           string  `json:"ipoDate"`           // IPO 날짜
	DefaultImage      bool    `json:"defaultImage"`      // 기본 이미지 여부
	IsETF             bool    `json:"isEtf"`             // ETF 여부
	IsActivelyTrading bool    `json:"isActivelyTrading"` // 활성 거래 여부
	IsADR             bool    `json:"isAdr"`             // ADR 여부
	IsFund            bool    `json:"isFund"`            // 펀드 여부
}

// SearchExchangeVariants 는 한 종목의 거래소별 변형(다른 거래소 상장)을 조회한다.
func (c *Client) SearchExchangeVariants(ctx context.Context, symbol string) ([]ExchangeVariant, error) {
	return fetch.ListBySymbol[ExchangeVariant](ctx, c.http, "/stable/search-exchange-variants", symbol)
}
