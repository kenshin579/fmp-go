package ratios

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Ratio 는 FMP /stable/ratios 응답 한 기간.
// FMP 응답 키를 충실히 매핑(faithful). 모든 비율은 소수(0.XX), 백분율 환산은 소비 측 책임.
type Ratio struct {
	// String fields
	Date             string `json:"date"`
	Symbol           string `json:"symbol"`
	ReportedCurrency string `json:"reportedCurrency"`
	FiscalYear       string `json:"fiscalYear"`
	Period           string `json:"period"`

	// Margin ratios
	GrossProfitMargin                float64 `json:"grossProfitMargin"`
	EbitMargin                       float64 `json:"ebitMargin"`
	EbitdaMargin                     float64 `json:"ebitdaMargin"`
	OperatingProfitMargin            float64 `json:"operatingProfitMargin"`
	PretaxProfitMargin               float64 `json:"pretaxProfitMargin"`
	ContinuousOperationsProfitMargin float64 `json:"continuousOperationsProfitMargin"`
	NetProfitMargin                  float64 `json:"netProfitMargin"`
	BottomLineProfitMargin           float64 `json:"bottomLineProfitMargin"`

	// Turnover ratios
	ReceivablesTurnover float64 `json:"receivablesTurnover"`
	PayablesTurnover    float64 `json:"payablesTurnover"`
	InventoryTurnover   float64 `json:"inventoryTurnover"`
	FixedAssetTurnover  float64 `json:"fixedAssetTurnover"`
	AssetTurnover       float64 `json:"assetTurnover"`

	// Liquidity ratios
	CurrentRatio  float64 `json:"currentRatio"`
	QuickRatio    float64 `json:"quickRatio"`
	SolvencyRatio float64 `json:"solvencyRatio"`
	CashRatio     float64 `json:"cashRatio"`

	// Price/valuation ratios
	PriceToEarningsRatio              float64 `json:"priceToEarningsRatio"`
	PriceToEarningsGrowthRatio        float64 `json:"priceToEarningsGrowthRatio"`
	ForwardPriceToEarningsGrowthRatio float64 `json:"forwardPriceToEarningsGrowthRatio"`
	PriceToBookRatio                  float64 `json:"priceToBookRatio"`
	PriceToSalesRatio                 float64 `json:"priceToSalesRatio"`
	PriceToFreeCashFlowRatio          float64 `json:"priceToFreeCashFlowRatio"`
	PriceToOperatingCashFlowRatio     float64 `json:"priceToOperatingCashFlowRatio"`
	PriceToFairValue                  float64 `json:"priceToFairValue"`

	// Debt ratios
	DebtToAssetsRatio          float64 `json:"debtToAssetsRatio"`
	DebtToEquityRatio          float64 `json:"debtToEquityRatio"`
	DebtToCapitalRatio         float64 `json:"debtToCapitalRatio"`
	LongTermDebtToCapitalRatio float64 `json:"longTermDebtToCapitalRatio"`
	FinancialLeverageRatio     float64 `json:"financialLeverageRatio"`
	DebtToMarketCap            float64 `json:"debtToMarketCap"`

	// Cash flow ratios
	WorkingCapitalTurnoverRatio             float64 `json:"workingCapitalTurnoverRatio"`
	OperatingCashFlowRatio                  float64 `json:"operatingCashFlowRatio"`
	OperatingCashFlowSalesRatio             float64 `json:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio      float64 `json:"freeCashFlowOperatingCashFlowRatio"`
	DebtServiceCoverageRatio                float64 `json:"debtServiceCoverageRatio"`
	InterestCoverageRatio                   float64 `json:"interestCoverageRatio"`
	ShortTermOperatingCashFlowCoverageRatio float64 `json:"shortTermOperatingCashFlowCoverageRatio"`
	OperatingCashFlowCoverageRatio          float64 `json:"operatingCashFlowCoverageRatio"`
	CapitalExpenditureCoverageRatio         float64 `json:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio       float64 `json:"dividendPaidAndCapexCoverageRatio"`

	// Dividend ratios
	DividendPayoutRatio     float64 `json:"dividendPayoutRatio"`
	DividendYield           float64 `json:"dividendYield"`
	DividendYieldPercentage float64 `json:"dividendYieldPercentage"`

	// Per-share metrics
	RevenuePerShare              float64 `json:"revenuePerShare"`
	NetIncomePerShare            float64 `json:"netIncomePerShare"`
	InterestDebtPerShare         float64 `json:"interestDebtPerShare"`
	CashPerShare                 float64 `json:"cashPerShare"`
	BookValuePerShare            float64 `json:"bookValuePerShare"`
	TangibleBookValuePerShare    float64 `json:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare   float64 `json:"shareholdersEquityPerShare"`
	OperatingCashFlowPerShare    float64 `json:"operatingCashFlowPerShare"`
	CapexPerShare                float64 `json:"capexPerShare"`
	FreeCashFlowPerShare         float64 `json:"freeCashFlowPerShare"`

	// Other metrics
	NetIncomePerEBT        float64 `json:"netIncomePerEBT"`
	EbtPerEbit             float64 `json:"ebtPerEbit"`
	EffectiveTaxRate       float64 `json:"effectiveTaxRate"`
	EnterpriseValueMultiple float64 `json:"enterpriseValueMultiple"`
}

// Ratios 는 종목의 재무비율 시계열을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) Ratios(ctx context.Context, p Params) ([]Ratio, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []Ratio
	if err := c.http.GetJSON(ctx, "/stable/ratios", p.queryParams(), &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
