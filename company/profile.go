package company

import (
	"context"
	"encoding/json"
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

// profileFields — Profile 과 같은 필드를 갖되 UnmarshalJSON 을 물려받지 않는 별칭.
// 아래 UnmarshalJSON 이 자기 자신을 무한히 부르지 않게 하는 표준 수법이다.
type profileFields Profile

// UnmarshalJSON — volume·averageVolume 이 소수로 와도 프로필을 살려서 받는다.
//
// FMP 는 이 둘을 정수로 문서화했지만 실제로는 소수가 섞여 온다(운영 관측: 0.656,
// 104964205.74442, 45424.207). 기본 디코더는 그때 **구조체 전체**를 포기하므로,
// 거래량 한 필드 때문에 회사명·섹터·CEO 까지 전부 날아간다. 그래서 두 필드만
// json.Number 로 받아 int64 로 버림한다.
//
// 필드 타입을 float64 로 바꾸지 않는 이유는 거래량이 주식 수이기 때문이다 —
// 소비자 코드에 소수 거래량이 퍼지는 것보다, 들어올 때 한 번 정수로 맞추는 편이
// 낫다. 같은 위험은 quote·chart·search 의 Volume 필드에도 있지만, 실제로 깨진
// 것을 본 곳만 고친다.
func (p *Profile) UnmarshalJSON(data []byte) error {
	var raw struct {
		profileFields
		Volume        json.Number `json:"volume"`
		AverageVolume json.Number `json:"averageVolume"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*p = Profile(raw.profileFields)
	p.Volume = truncInt64(raw.Volume)
	p.AverageVolume = truncInt64(raw.AverageVolume)
	return nil
}

// truncInt64 — JSON 숫자를 int64 로 버림한다. 빈 값(필드 없음)과 int64 로 담을 수
// 없는 값은 0 이다. 후자는 거래량으로 성립하지 않는 수라 버리는 편이 낫다.
func truncInt64(n json.Number) int64 {
	if n == "" {
		return 0
	}
	if i, err := n.Int64(); err == nil {
		return i
	}
	f, err := n.Float64()
	if err != nil {
		return 0
	}
	return int64(f)
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
