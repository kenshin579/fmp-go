package dcf

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CustomDCFParams — custom DCF 입력. Symbol 필수, 나머지 override 는 미설정 시 제외.
type CustomDCFParams struct {
	Symbol                                     string
	RevenueGrowthPct                           *float64
	EbitdaPct                                  *float64
	DepreciationAndAmortizationPct             *float64
	CashAndShortTermInvestmentsPct             *float64
	ReceivablesPct                             *float64
	InventoriesPct                             *float64
	PayablePct                                 *float64
	EbitPct                                    *float64
	CapitalExpenditurePct                      *float64
	OperatingCashFlowPct                       *float64
	SellingGeneralAndAdministrativeExpensesPct *float64
	TaxRate                                    *float64
	LongTermGrowthRate                         *float64
	CostOfDebt                                 *float64
	CostOfEquity                               *float64
	MarketRiskPremium                          *float64
	Beta                                       *float64
	RiskFreeRate                               *float64
}

func (p CustomDCFParams) queryParams() map[string]string {
	q := map[string]string{"symbol": p.Symbol}
	add := func(key string, v *float64) {
		if v != nil {
			q[key] = strconv.FormatFloat(*v, 'f', -1, 64)
		}
	}
	add("revenueGrowthPct", p.RevenueGrowthPct)
	add("ebitdaPct", p.EbitdaPct)
	add("depreciationAndAmortizationPct", p.DepreciationAndAmortizationPct)
	add("cashAndShortTermInvestmentsPct", p.CashAndShortTermInvestmentsPct)
	add("receivablesPct", p.ReceivablesPct)
	add("inventoriesPct", p.InventoriesPct)
	add("payablePct", p.PayablePct)
	add("ebitPct", p.EbitPct)
	add("capitalExpenditurePct", p.CapitalExpenditurePct)
	add("operatingCashFlowPct", p.OperatingCashFlowPct)
	add("sellingGeneralAndAdministrativeExpensesPct", p.SellingGeneralAndAdministrativeExpensesPct)
	add("taxRate", p.TaxRate)
	add("longTermGrowthRate", p.LongTermGrowthRate)
	add("costOfDebt", p.CostOfDebt)
	add("costOfEquity", p.CostOfEquity)
	add("marketRiskPremium", p.MarketRiskPremium)
	add("beta", p.Beta)
	add("riskFreeRate", p.RiskFreeRate)
	return q
}

// CustomDCFAdvanced — 상세 다년 DCF 투영 (custom-discounted-cash-flow)
type CustomDCFAdvanced struct {
	Year                         string  `json:"year"`
	Symbol                       string  `json:"symbol"`
	Revenue                      int64   `json:"revenue"`
	RevenuePercentage            float64 `json:"revenuePercentage"`
	EBITDA                       int64   `json:"ebitda"`
	EBITDAPercentage             float64 `json:"ebitdaPercentage"`
	EBIT                         int64   `json:"ebit"`
	EBITPercentage               float64 `json:"ebitPercentage"`
	Depreciation                 int64   `json:"depreciation"`
	DepreciationPercentage       float64 `json:"depreciationPercentage"`
	TotalCash                    int64   `json:"totalCash"`
	TotalCashPercentage          float64 `json:"totalCashPercentage"`
	Receivables                  int64   `json:"receivables"`
	ReceivablesPercentage        float64 `json:"receivablesPercentage"`
	Inventories                  int64   `json:"inventories"`
	InventoriesPercentage        float64 `json:"inventoriesPercentage"`
	Payable                      int64   `json:"payable"`
	PayablePercentage            float64 `json:"payablePercentage"`
	CapitalExpenditure           int64   `json:"capitalExpenditure"`
	CapitalExpenditurePercentage float64 `json:"capitalExpenditurePercentage"`
	Price                        float64 `json:"price"`
	Beta                         float64 `json:"beta"`
	DilutedSharesOutstanding     int64   `json:"dilutedSharesOutstanding"`
	CostOfDebt                   float64 `json:"costofDebt"`
	TaxRate                      float64 `json:"taxRate"`
	AfterTaxCostOfDebt           float64 `json:"afterTaxCostOfDebt"`
	RiskFreeRate                 float64 `json:"riskFreeRate"`
	MarketRiskPremium            float64 `json:"marketRiskPremium"`
	CostOfEquity                 float64 `json:"costOfEquity"`
	TotalDebt                    int64   `json:"totalDebt"`
	TotalEquity                  int64   `json:"totalEquity"`
	TotalCapital                 int64   `json:"totalCapital"`
	DebtWeighting                float64 `json:"debtWeighting"`
	EquityWeighting              float64 `json:"equityWeighting"`
	WACC                         float64 `json:"wacc"`
	TaxRateCash                  int64   `json:"taxRateCash"`
	EBIAT                        int64   `json:"ebiat"`
	UFCF                         int64   `json:"ufcf"`
	SumPvUfcf                    int64   `json:"sumPvUfcf"`
	LongTermGrowthRate           float64 `json:"longTermGrowthRate"`
	TerminalValue                int64   `json:"terminalValue"`
	PresentTerminalValue         int64   `json:"presentTerminalValue"`
	EnterpriseValue              int64   `json:"enterpriseValue"`
	NetDebt                      int64   `json:"netDebt"`
	EquityValue                  int64   `json:"equityValue"`
	EquityValuePerShare          float64 `json:"equityValuePerShare"`
	FreeCashFlowT1               int64   `json:"freeCashFlowT1"`
}

// CustomDCFLevered — 레버드 다년 DCF 투영 (custom-levered-discounted-cash-flow).
type CustomDCFLevered struct {
	Year                         string  `json:"year"`
	Symbol                       string  `json:"symbol"`
	Revenue                      int64   `json:"revenue"`
	RevenuePercentage            float64 `json:"revenuePercentage"`
	CapitalExpenditure           int64   `json:"capitalExpenditure"`
	CapitalExpenditurePercentage float64 `json:"capitalExpenditurePercentage"`
	Price                        float64 `json:"price"`
	Beta                         float64 `json:"beta"`
	DilutedSharesOutstanding     int64   `json:"dilutedSharesOutstanding"`
	CostOfDebt                   float64 `json:"costofDebt"`
	TaxRate                      float64 `json:"taxRate"`
	AfterTaxCostOfDebt           float64 `json:"afterTaxCostOfDebt"`
	RiskFreeRate                 float64 `json:"riskFreeRate"`
	MarketRiskPremium            float64 `json:"marketRiskPremium"`
	CostOfEquity                 float64 `json:"costOfEquity"`
	TotalDebt                    int64   `json:"totalDebt"`
	TotalEquity                  int64   `json:"totalEquity"`
	TotalCapital                 int64   `json:"totalCapital"`
	DebtWeighting                float64 `json:"debtWeighting"`
	EquityWeighting              float64 `json:"equityWeighting"`
	WACC                         float64 `json:"wacc"`
	OperatingCashFlow            int64   `json:"operatingCashFlow"`
	PvLfcf                       int64   `json:"pvLfcf"`
	SumPvLfcf                    int64   `json:"sumPvLfcf"`
	FreeCashFlow                 int64   `json:"freeCashFlow"`
	OperatingCashFlowPercentage  float64 `json:"operatingCashFlowPercentage"`
	LongTermGrowthRate           float64 `json:"longTermGrowthRate"`
	TerminalValue                int64   `json:"terminalValue"`
	PresentTerminalValue         int64   `json:"presentTerminalValue"`
	EnterpriseValue              int64   `json:"enterpriseValue"`
	NetDebt                      int64   `json:"netDebt"`
	EquityValue                  int64   `json:"equityValue"`
	EquityValuePerShare          float64 `json:"equityValuePerShare"`
	FreeCashFlowT1               int64   `json:"freeCashFlowT1"`
}

// CustomDiscountedCashFlow 는 override 를 적용한 상세 DCF 투영을 조회한다.
func (c *Client) CustomDiscountedCashFlow(ctx context.Context, p CustomDCFParams) ([]CustomDCFAdvanced, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[CustomDCFAdvanced](ctx, c.http, "/stable/custom-discounted-cash-flow", p.queryParams())
}

// CustomLeveredDiscountedCashFlow 는 override 를 적용한 레버드 DCF 투영을 조회한다.
func (c *Client) CustomLeveredDiscountedCashFlow(ctx context.Context, p CustomDCFParams) ([]CustomDCFLevered, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[CustomDCFLevered](ctx, c.http, "/stable/custom-levered-discounted-cash-flow", p.queryParams())
}
