package metrics

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// KeyMetrics — 핵심 지표 (key-metrics). 비율 다수, 일부 절대값 int64.
type KeyMetrics struct {
	Symbol                                 string  `json:"symbol"`
	Date                                   string  `json:"date"`
	FiscalYear                             string  `json:"fiscalYear"`
	Period                                 string  `json:"period"`
	ReportedCurrency                       string  `json:"reportedCurrency"`
	MarketCap                              int64   `json:"marketCap"`
	EnterpriseValue                        int64   `json:"enterpriseValue"`
	EvToSales                              float64 `json:"evToSales"`
	EvToOperatingCashFlow                  float64 `json:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64 `json:"evToFreeCashFlow"`
	EvToEBITDA                             float64 `json:"evToEBITDA"`
	NetDebtToEBITDA                        float64 `json:"netDebtToEBITDA"`
	CurrentRatio                           float64 `json:"currentRatio"`
	IncomeQuality                          float64 `json:"incomeQuality"`
	GrahamNumber                           float64 `json:"grahamNumber"`
	GrahamNetNet                           float64 `json:"grahamNetNet"`
	TaxBurden                              float64 `json:"taxBurden"`
	InterestBurden                         float64 `json:"interestBurden"`
	WorkingCapital                         int64   `json:"workingCapital"`
	InvestedCapital                        int64   `json:"investedCapital"`
	ReturnOnAssets                         float64 `json:"returnOnAssets"`
	OperatingReturnOnAssets                float64 `json:"operatingReturnOnAssets"`
	ReturnOnTangibleAssets                 float64 `json:"returnOnTangibleAssets"`
	ReturnOnEquity                         float64 `json:"returnOnEquity"`
	ReturnOnInvestedCapital                float64 `json:"returnOnInvestedCapital"`
	ReturnOnCapitalEmployed                float64 `json:"returnOnCapitalEmployed"`
	EarningsYield                          float64 `json:"earningsYield"`
	FreeCashFlowYield                      float64 `json:"freeCashFlowYield"`
	CapexToOperatingCashFlow               float64 `json:"capexToOperatingCashFlow"`
	CapexToDepreciation                    float64 `json:"capexToDepreciation"`
	CapexToRevenue                         float64 `json:"capexToRevenue"`
	SalesGeneralAndAdministrativeToRevenue float64 `json:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDevelopementToRevenue       float64 `json:"researchAndDevelopementToRevenue"`
	StockBasedCompensationToRevenue        float64 `json:"stockBasedCompensationToRevenue"`
	IntangiblesToTotalAssets               float64 `json:"intangiblesToTotalAssets"`
	AverageReceivables                     int64   `json:"averageReceivables"`
	AveragePayables                        int64   `json:"averagePayables"`
	AverageInventory                       int64   `json:"averageInventory"`
	DaysOfSalesOutstanding                 float64 `json:"daysOfSalesOutstanding"`
	DaysOfPayablesOutstanding              float64 `json:"daysOfPayablesOutstanding"`
	DaysOfInventoryOutstanding             float64 `json:"daysOfInventoryOutstanding"`
	OperatingCycle                         float64 `json:"operatingCycle"`
	CashConversionCycle                    float64 `json:"cashConversionCycle"`
	FreeCashFlowToEquity                   int64   `json:"freeCashFlowToEquity"`
	FreeCashFlowToFirm                     float64 `json:"freeCashFlowToFirm"`
	TangibleAssetValue                     int64   `json:"tangibleAssetValue"`
	NetCurrentAssetValue                   int64   `json:"netCurrentAssetValue"`
}

// KeyMetricsTTM — TTM 핵심 지표 (key-metrics-ttm). symbol 외 메타 필드 없음.
type KeyMetricsTTM struct {
	Symbol                                    string  `json:"symbol"`
	MarketCap                                 int64   `json:"marketCap"`
	EnterpriseValueTTM                        int64   `json:"enterpriseValueTTM"`
	EvToSalesTTM                              float64 `json:"evToSalesTTM"`
	EvToOperatingCashFlowTTM                  float64 `json:"evToOperatingCashFlowTTM"`
	EvToFreeCashFlowTTM                       float64 `json:"evToFreeCashFlowTTM"`
	EvToEBITDATTM                             float64 `json:"evToEBITDATTM"`
	NetDebtToEBITDATTM                        float64 `json:"netDebtToEBITDATTM"`
	CurrentRatioTTM                           float64 `json:"currentRatioTTM"`
	IncomeQualityTTM                          float64 `json:"incomeQualityTTM"`
	GrahamNumberTTM                           float64 `json:"grahamNumberTTM"`
	GrahamNetNetTTM                           float64 `json:"grahamNetNetTTM"`
	TaxBurdenTTM                              float64 `json:"taxBurdenTTM"`
	InterestBurdenTTM                         float64 `json:"interestBurdenTTM"`
	WorkingCapitalTTM                         int64   `json:"workingCapitalTTM"`
	InvestedCapitalTTM                        int64   `json:"investedCapitalTTM"`
	ReturnOnAssetsTTM                         float64 `json:"returnOnAssetsTTM"`
	OperatingReturnOnAssetsTTM                float64 `json:"operatingReturnOnAssetsTTM"`
	ReturnOnTangibleAssetsTTM                 float64 `json:"returnOnTangibleAssetsTTM"`
	ReturnOnEquityTTM                         float64 `json:"returnOnEquityTTM"`
	ReturnOnInvestedCapitalTTM                float64 `json:"returnOnInvestedCapitalTTM"`
	ReturnOnCapitalEmployedTTM                float64 `json:"returnOnCapitalEmployedTTM"`
	EarningsYieldTTM                          float64 `json:"earningsYieldTTM"`
	FreeCashFlowYieldTTM                      float64 `json:"freeCashFlowYieldTTM"`
	CapexToOperatingCashFlowTTM               float64 `json:"capexToOperatingCashFlowTTM"`
	CapexToDepreciationTTM                    float64 `json:"capexToDepreciationTTM"`
	CapexToRevenueTTM                         float64 `json:"capexToRevenueTTM"`
	SalesGeneralAndAdministrativeToRevenueTTM float64 `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopementToRevenueTTM       float64 `json:"researchAndDevelopementToRevenueTTM"`
	StockBasedCompensationToRevenueTTM        float64 `json:"stockBasedCompensationToRevenueTTM"`
	IntangiblesToTotalAssetsTTM               float64 `json:"intangiblesToTotalAssetsTTM"`
	AverageReceivablesTTM                     int64   `json:"averageReceivablesTTM"`
	AveragePayablesTTM                        int64   `json:"averagePayablesTTM"`
	AverageInventoryTTM                       int64   `json:"averageInventoryTTM"`
	DaysOfSalesOutstandingTTM                 float64 `json:"daysOfSalesOutstandingTTM"`
	DaysOfPayablesOutstandingTTM              float64 `json:"daysOfPayablesOutstandingTTM"`
	DaysOfInventoryOutstandingTTM             float64 `json:"daysOfInventoryOutstandingTTM"`
	OperatingCycleTTM                         float64 `json:"operatingCycleTTM"`
	CashConversionCycleTTM                    float64 `json:"cashConversionCycleTTM"`
	FreeCashFlowToEquityTTM                   int64   `json:"freeCashFlowToEquityTTM"`
	FreeCashFlowToFirmTTM                     float64 `json:"freeCashFlowToFirmTTM"`
	TangibleAssetValueTTM                     int64   `json:"tangibleAssetValueTTM"`
	NetCurrentAssetValueTTM                   int64   `json:"netCurrentAssetValueTTM"`
}

// KeyMetrics 는 종목의 핵심 지표 시계열을 조회한다.
func (c *Client) KeyMetrics(ctx context.Context, symbol, period string, limit int) ([]KeyMetrics, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[KeyMetrics](ctx, c.http, "/stable/key-metrics", listParams(symbol, period, limit))
}

// KeyMetricsTTM 는 종목의 TTM 핵심 지표를 조회한다.
func (c *Client) KeyMetricsTTM(ctx context.Context, symbol string) ([]KeyMetricsTTM, error) {
	return fetch.ListBySymbol[KeyMetricsTTM](ctx, c.http, "/stable/key-metrics-ttm", symbol)
}
