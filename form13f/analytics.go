package form13f

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// HolderAnalytics 는 기관 투자자별 13F 보유 분석 데이터 (filings-extract-with-analytics-by-holder).
type HolderAnalytics struct {
	// 보고 기준일
	Date string `json:"date"`
	// 기관 투자자 CIK
	Cik string `json:"cik"`
	// 제출일
	FilingDate string `json:"filingDate"`
	// 기관 투자자명
	InvestorName string `json:"investorName"`
	// 종목 티커
	Symbol string `json:"symbol"`
	// 증권명
	SecurityName string `json:"securityName"`
	// 증권 유형
	TypeOfSecurity string `json:"typeOfSecurity"`
	// 증권 CUSIP
	SecurityCusip string `json:"securityCusip"`
	// 주식 유형
	SharesType string `json:"sharesType"`
	// 풋/콜/주식 구분
	PutCallShare string `json:"putCallShare"`
	// 투자 재량
	InvestmentDiscretion string `json:"investmentDiscretion"`
	// 산업 분류
	IndustryTitle string `json:"industryTitle"`
	// 포트폴리오 내 비중
	Weight float64 `json:"weight"`
	// 이전 분기 비중
	LastWeight float64 `json:"lastWeight"`
	// 비중 변화
	ChangeInWeight float64 `json:"changeInWeight"`
	// 비중 변화율 (%)
	ChangeInWeightPercentage float64 `json:"changeInWeightPercentage"`
	// 시가총액 (보유 기준)
	MarketValue int64 `json:"marketValue"`
	// 이전 분기 시가총액
	LastMarketValue int64 `json:"lastMarketValue"`
	// 시가총액 변화
	ChangeInMarketValue int64 `json:"changeInMarketValue"`
	// 시가총액 변화율 (%)
	ChangeInMarketValuePercentage float64 `json:"changeInMarketValuePercentage"`
	// 보유 주식 수
	SharesNumber int64 `json:"sharesNumber"`
	// 이전 분기 보유 주식 수
	LastSharesNumber int64 `json:"lastSharesNumber"`
	// 주식 수 변화
	ChangeInSharesNumber int64 `json:"changeInSharesNumber"`
	// 주식 수 변화율 (%)
	ChangeInSharesNumberPercentage float64 `json:"changeInSharesNumberPercentage"`
	// 분기말 주가
	QuarterEndPrice float64 `json:"quarterEndPrice"`
	// 평균 매입가
	AvgPricePaid float64 `json:"avgPricePaid"`
	// 신규 편입 여부
	IsNew bool `json:"isNew"`
	// 전량 매도 여부
	IsSoldOut bool `json:"isSoldOut"`
	// 지분율
	Ownership float64 `json:"ownership"`
	// 이전 분기 지분율
	LastOwnership float64 `json:"lastOwnership"`
	// 지분율 변화
	ChangeInOwnership float64 `json:"changeInOwnership"`
	// 지분율 변화율 (%)
	ChangeInOwnershipPercentage float64 `json:"changeInOwnershipPercentage"`
	// 보유 기간 (분기 수)
	HoldingPeriod int64 `json:"holdingPeriod"`
	// 최초 편입일
	FirstAdded string `json:"firstAdded"`
	// 성과 (손익)
	Performance int64 `json:"performance"`
	// 성과율 (%)
	PerformancePercentage float64 `json:"performancePercentage"`
	// 이전 분기 성과
	LastPerformance int64 `json:"lastPerformance"`
	// 성과 변화
	ChangeInPerformance int64 `json:"changeInPerformance"`
	// 성과 계산 포함 여부
	IsCountedForPerformance bool `json:"isCountedForPerformance"`
}

// IndustryBreakdown 는 기관 투자자 산업별 보유 비중 분석 데이터 (holders-industry-breakdown).
type IndustryBreakdown struct {
	// 보고 기준일
	Date string `json:"date"`
	// 기관 투자자 CIK
	Cik string `json:"cik"`
	// 기관 투자자명
	InvestorName string `json:"investorName"`
	// 산업 분류
	IndustryTitle string `json:"industryTitle"`
	// 포트폴리오 내 비중
	Weight float64 `json:"weight"`
	// 이전 분기 비중
	LastWeight float64 `json:"lastWeight"`
	// 비중 변화
	ChangeInWeight float64 `json:"changeInWeight"`
	// 비중 변화율 (%)
	ChangeInWeightPercentage float64 `json:"changeInWeightPercentage"`
	// 성과 (손익)
	Performance int64 `json:"performance"`
	// 성과율 (%)
	PerformancePercentage float64 `json:"performancePercentage"`
	// 이전 분기 성과
	LastPerformance int64 `json:"lastPerformance"`
	// 성과 변화
	ChangeInPerformance int64 `json:"changeInPerformance"`
}

// ExtractAnalyticsByHolder 는 종목별 기관 투자자 분석 데이터를 조회한다.
func (c *Client) ExtractAnalyticsByHolder(ctx context.Context, symbol, year, quarter string, page, limit int) ([]HolderAnalytics, error) {
	if strings.TrimSpace(symbol) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: symbol, year, quarter must not be empty")
	}
	q := pageParams(page, limit)
	q["symbol"] = symbol
	q["year"] = year
	q["quarter"] = quarter
	return fetch.List[HolderAnalytics](ctx, c.http, "/stable/institutional-ownership/extract-analytics/holder", q)
}

// HoldersIndustryBreakdown 는 기관 투자자의 산업별 보유 비중을 조회한다.
func (c *Client) HoldersIndustryBreakdown(ctx context.Context, cik, year, quarter string) ([]IndustryBreakdown, error) {
	if strings.TrimSpace(cik) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: cik, year, quarter must not be empty")
	}
	return fetch.List[IndustryBreakdown](ctx, c.http, "/stable/institutional-ownership/holder-industry-breakdown", map[string]string{"cik": cik, "year": year, "quarter": quarter})
}
