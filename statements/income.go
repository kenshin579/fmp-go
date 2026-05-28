package statements

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// IncomeStatement 는 FMP /stable/income-statement 응답 한 기간(연간/분기).
// 필드는 FMP 응답을 충실히 매핑한다(faithful).
type IncomeStatement struct {
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	CIK                                     string  `json:"cik"`
	FilingDate                              string  `json:"filingDate"`
	AcceptedDate                            string  `json:"acceptedDate"`
	FiscalYear                              string  `json:"fiscalYear"`
	Period                                  string  `json:"period"` // "FY"/"Q1"..
	Revenue                                 int64   `json:"revenue"`
	CostOfRevenue                           int64   `json:"costOfRevenue"`
	GrossProfit                             int64   `json:"grossProfit"`
	ResearchAndDevelopmentExpenses          int64   `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        int64   `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             int64   `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses int64   `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           int64   `json:"otherExpenses"`
	OperatingExpenses                       int64   `json:"operatingExpenses"`
	CostAndExpenses                         int64   `json:"costAndExpenses"`
	NetInterestIncome                       int64   `json:"netInterestIncome"`
	InterestIncome                          int64   `json:"interestIncome"`
	InterestExpense                         int64   `json:"interestExpense"`
	DepreciationAndAmortization             int64   `json:"depreciationAndAmortization"`
	EBITDA                                  int64   `json:"ebitda"`
	EBIT                                    int64   `json:"ebit"`
	NonOperatingIncomeExcludingInterest     int64   `json:"nonOperatingIncomeExcludingInterest"`
	OperatingIncome                         int64   `json:"operatingIncome"`
	TotalOtherIncomeExpensesNet             int64   `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         int64   `json:"incomeBeforeTax"`
	IncomeTaxExpense                        int64   `json:"incomeTaxExpense"`
	NetIncomeFromContinuingOperations       int64   `json:"netIncomeFromContinuingOperations"`
	NetIncomeFromDiscontinuedOperations     int64   `json:"netIncomeFromDiscontinuedOperations"`
	OtherAdjustmentsToNetIncome             int64   `json:"otherAdjustmentsToNetIncome"`
	NetIncome                               int64   `json:"netIncome"`
	NetIncomeDeductions                     int64   `json:"netIncomeDeductions"`
	BottomLineNetIncome                     int64   `json:"bottomLineNetIncome"`
	EPS                                     float64 `json:"eps"`
	EPSDiluted                              float64 `json:"epsDiluted"`
	WeightedAverageShsOut                   int64   `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                int64   `json:"weightedAverageShsOutDil"`
}

// IncomeStatement 는 종목의 손익계산서 시계열을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) IncomeStatement(ctx context.Context, p Params) ([]IncomeStatement, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []IncomeStatement
	if err := c.http.GetJSON(ctx, "/stable/income-statement", p.queryParams(), &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
