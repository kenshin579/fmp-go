package form13f

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// HolderPerformance 는 기관 투자자의 보유 포트폴리오 성과 요약 정보.
type HolderPerformance struct {
	Date                                               string  `json:"date"`
	Cik                                                string  `json:"cik"`
	InvestorName                                       string  `json:"investorName"`
	PortfolioSize                                      int64   `json:"portfolioSize"`
	SecuritiesAdded                                    int64   `json:"securitiesAdded"`
	SecuritiesRemoved                                  int64   `json:"securitiesRemoved"`
	MarketValue                                        int64   `json:"marketValue"`
	PreviousMarketValue                                int64   `json:"previousMarketValue"`
	ChangeInMarketValue                                int64   `json:"changeInMarketValue"`
	ChangeInMarketValuePercentage                      float64 `json:"changeInMarketValuePercentage"`
	AverageHoldingPeriod                               int64   `json:"averageHoldingPeriod"`
	AverageHoldingPeriodTop10                          int64   `json:"averageHoldingPeriodTop10"`
	AverageHoldingPeriodTop20                          int64   `json:"averageHoldingPeriodTop20"`
	Turnover                                           float64 `json:"turnover"`
	TurnoverAlternateSell                              float64 `json:"turnoverAlternateSell"`
	TurnoverAlternateBuy                               float64 `json:"turnoverAlternateBuy"`
	Performance                                        int64   `json:"performance"`
	PerformancePercentage                              float64 `json:"performancePercentage"`
	LastPerformance                                    int64   `json:"lastPerformance"`
	ChangeInPerformance                                int64   `json:"changeInPerformance"`
	Performance1year                                   int64   `json:"performance1year"`
	PerformancePercentage1year                         float64 `json:"performancePercentage1year"`
	Performance3year                                   int64   `json:"performance3year"`
	PerformancePercentage3year                         float64 `json:"performancePercentage3year"`
	Performance5year                                   int64   `json:"performance5year"`
	PerformancePercentage5year                         float64 `json:"performancePercentage5year"`
	PerformanceSinceInception                          int64   `json:"performanceSinceInception"`
	PerformanceSinceInceptionPercentage                float64 `json:"performanceSinceInceptionPercentage"`
	PerformanceRelativeToSP500Percentage               float64 `json:"performanceRelativeToSP500Percentage"`
	Performance1yearRelativeToSP500Percentage          float64 `json:"performance1yearRelativeToSP500Percentage"`
	Performance3yearRelativeToSP500Percentage          float64 `json:"performance3yearRelativeToSP500Percentage"`
	Performance5yearRelativeToSP500Percentage          float64 `json:"performance5yearRelativeToSP500Percentage"`
	PerformanceSinceInceptionRelativeToSP500Percentage float64 `json:"performanceSinceInceptionRelativeToSP500Percentage"`
}

// PositionSummary 는 특정 종목에 대한 기관 보유 포지션 요약 정보.
type PositionSummary struct {
	Symbol                   string  `json:"symbol"`
	Cik                      string  `json:"cik"`
	Date                     string  `json:"date"`
	InvestorsHolding         int64   `json:"investorsHolding"`
	LastInvestorsHolding     int64   `json:"lastInvestorsHolding"`
	InvestorsHoldingChange   int64   `json:"investorsHoldingChange"`
	NumberOf13Fshares        int64   `json:"numberOf13Fshares"`
	LastNumberOf13Fshares    int64   `json:"lastNumberOf13Fshares"`
	NumberOf13FsharesChange  int64   `json:"numberOf13FsharesChange"`
	TotalInvested            int64   `json:"totalInvested"`
	LastTotalInvested        int64   `json:"lastTotalInvested"`
	TotalInvestedChange      int64   `json:"totalInvestedChange"`
	OwnershipPercent         float64 `json:"ownershipPercent"`
	LastOwnershipPercent     float64 `json:"lastOwnershipPercent"`
	OwnershipPercentChange   float64 `json:"ownershipPercentChange"`
	NewPositions             int64   `json:"newPositions"`
	LastNewPositions         int64   `json:"lastNewPositions"`
	NewPositionsChange       int64   `json:"newPositionsChange"`
	IncreasedPositions       int64   `json:"increasedPositions"`
	LastIncreasedPositions   int64   `json:"lastIncreasedPositions"`
	IncreasedPositionsChange int64   `json:"increasedPositionsChange"`
	ClosedPositions          int64   `json:"closedPositions"`
	LastClosedPositions      int64   `json:"lastClosedPositions"`
	ClosedPositionsChange    int64   `json:"closedPositionsChange"`
	ReducedPositions         int64   `json:"reducedPositions"`
	LastReducedPositions     int64   `json:"lastReducedPositions"`
	ReducedPositionsChange   int64   `json:"reducedPositionsChange"`
	TotalCalls               int64   `json:"totalCalls"`
	LastTotalCalls           int64   `json:"lastTotalCalls"`
	TotalCallsChange         int64   `json:"totalCallsChange"`
	TotalPuts                int64   `json:"totalPuts"`
	LastTotalPuts            int64   `json:"lastTotalPuts"`
	TotalPutsChange          int64   `json:"totalPutsChange"`
	PutCallRatio             float64 `json:"putCallRatio"`
	LastPutCallRatio         float64 `json:"lastPutCallRatio"`
	PutCallRatioChange       float64 `json:"putCallRatioChange"`
}

// HolderPerformanceSummary 는 CIK 기관의 포트폴리오 성과 요약 목록을 반환한다.
func (c *Client) HolderPerformanceSummary(ctx context.Context, cik string, page int) ([]HolderPerformance, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[HolderPerformance](ctx, c.http, "/stable/institutional-ownership/holder-performance-summary", map[string]string{"cik": cik, "page": strconv.Itoa(page)})
}

// PositionsSummary 는 특정 종목의 기관 포지션 요약 목록을 반환한다.
func (c *Client) PositionsSummary(ctx context.Context, symbol, year, quarter string) ([]PositionSummary, error) {
	if strings.TrimSpace(symbol) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: symbol, year, quarter must not be empty")
	}
	return fetch.List[PositionSummary](ctx, c.http, "/stable/institutional-ownership/symbol-positions-summary", map[string]string{"symbol": symbol, "year": year, "quarter": quarter})
}
