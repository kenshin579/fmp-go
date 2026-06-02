package ratios

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// RatioTTM — TTM 재무비율 (ratios-ttm). symbol 외 메타 없음.
type RatioTTM struct {
	Symbol                                     string  `json:"symbol"`
	GrossProfitMarginTTM                       float64 `json:"grossProfitMarginTTM"`
	EbitMarginTTM                              float64 `json:"ebitMarginTTM"`
	EbitdaMarginTTM                            float64 `json:"ebitdaMarginTTM"`
	OperatingProfitMarginTTM                   float64 `json:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM                      float64 `json:"pretaxProfitMarginTTM"`
	ContinuousOperationsProfitMarginTTM        float64 `json:"continuousOperationsProfitMarginTTM"`
	NetProfitMarginTTM                         float64 `json:"netProfitMarginTTM"`
	BottomLineProfitMarginTTM                  float64 `json:"bottomLineProfitMarginTTM"`
	ReceivablesTurnoverTTM                     float64 `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                        float64 `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                       float64 `json:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM                      float64 `json:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM                           float64 `json:"assetTurnoverTTM"`
	CurrentRatioTTM                            float64 `json:"currentRatioTTM"`
	QuickRatioTTM                              float64 `json:"quickRatioTTM"`
	SolvencyRatioTTM                           float64 `json:"solvencyRatioTTM"`
	CashRatioTTM                               float64 `json:"cashRatioTTM"`
	PriceToEarningsRatioTTM                    float64 `json:"priceToEarningsRatioTTM"`
	PriceToEarningsGrowthRatioTTM              float64 `json:"priceToEarningsGrowthRatioTTM"`
	ForwardPriceToEarningsGrowthRatioTTM       float64 `json:"forwardPriceToEarningsGrowthRatioTTM"`
	PriceToBookRatioTTM                        float64 `json:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                       float64 `json:"priceToSalesRatioTTM"`
	PriceToFreeCashFlowRatioTTM                float64 `json:"priceToFreeCashFlowRatioTTM"`
	PriceToOperatingCashFlowRatioTTM           float64 `json:"priceToOperatingCashFlowRatioTTM"`
	DebtToAssetsRatioTTM                       float64 `json:"debtToAssetsRatioTTM"`
	DebtToEquityRatioTTM                       float64 `json:"debtToEquityRatioTTM"`
	DebtToCapitalRatioTTM                      float64 `json:"debtToCapitalRatioTTM"`
	LongTermDebtToCapitalRatioTTM              float64 `json:"longTermDebtToCapitalRatioTTM"`
	FinancialLeverageRatioTTM                  float64 `json:"financialLeverageRatioTTM"`
	WorkingCapitalTurnoverRatioTTM             float64 `json:"workingCapitalTurnoverRatioTTM"`
	OperatingCashFlowRatioTTM                  float64 `json:"operatingCashFlowRatioTTM"`
	OperatingCashFlowSalesRatioTTM             float64 `json:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM      float64 `json:"freeCashFlowOperatingCashFlowRatioTTM"`
	DebtServiceCoverageRatioTTM                float64 `json:"debtServiceCoverageRatioTTM"`
	InterestCoverageRatioTTM                   float64 `json:"interestCoverageRatioTTM"`
	ShortTermOperatingCashFlowCoverageRatioTTM float64 `json:"shortTermOperatingCashFlowCoverageRatioTTM"`
	OperatingCashFlowCoverageRatioTTM          float64 `json:"operatingCashFlowCoverageRatioTTM"`
	CapitalExpenditureCoverageRatioTTM         float64 `json:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTM       float64 `json:"dividendPaidAndCapexCoverageRatioTTM"`
	DividendPayoutRatioTTM                     float64 `json:"dividendPayoutRatioTTM"`
	DividendYieldTTM                           float64 `json:"dividendYieldTTM"`
	EnterpriseValueTTM                         int64   `json:"enterpriseValueTTM"`
	RevenuePerShareTTM                         float64 `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM                       float64 `json:"netIncomePerShareTTM"`
	InterestDebtPerShareTTM                    float64 `json:"interestDebtPerShareTTM"`
	CashPerShareTTM                            float64 `json:"cashPerShareTTM"`
	BookValuePerShareTTM                       float64 `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM               float64 `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM              float64 `json:"shareholdersEquityPerShareTTM"`
	OperatingCashFlowPerShareTTM               float64 `json:"operatingCashFlowPerShareTTM"`
	CapexPerShareTTM                           float64 `json:"capexPerShareTTM"`
	FreeCashFlowPerShareTTM                    float64 `json:"freeCashFlowPerShareTTM"`
	NetIncomePerEBTTTM                         float64 `json:"netIncomePerEBTTTM"`
	EbtPerEbitTTM                              float64 `json:"ebtPerEbitTTM"`
	PriceToFairValueTTM                        float64 `json:"priceToFairValueTTM"`
	DebtToMarketCapTTM                         float64 `json:"debtToMarketCapTTM"`
	EffectiveTaxRateTTM                        float64 `json:"effectiveTaxRateTTM"`
	EnterpriseValueMultipleTTM                 float64 `json:"enterpriseValueMultipleTTM"`
}

// RatiosTTM 는 종목의 TTM 재무비율을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) RatiosTTM(ctx context.Context, symbol string) ([]RatioTTM, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []RatioTTM
	if err := c.http.GetJSON(ctx, "/stable/ratios-ttm", map[string]string{"symbol": symbol}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
