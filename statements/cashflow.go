package statements

import "context"

// CashFlowStatement 는 FMP 현금흐름표 한 기간(연간/분기/TTM). faithful 매핑.
type CashFlowStatement struct {
	Date                                   string `json:"date"`
	Symbol                                 string `json:"symbol"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	CIK                                    string `json:"cik"`
	FilingDate                             string `json:"filingDate"`
	AcceptedDate                           string `json:"acceptedDate"`
	FiscalYear                             string `json:"fiscalYear"`
	Period                                 string `json:"period"`
	NetIncome                              int64  `json:"netIncome"`
	DepreciationAndAmortization            int64  `json:"depreciationAndAmortization"`
	DeferredIncomeTax                      int64  `json:"deferredIncomeTax"`
	StockBasedCompensation                 int64  `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                 int64  `json:"changeInWorkingCapital"`
	AccountsReceivables                    int64  `json:"accountsReceivables"`
	Inventory                              int64  `json:"inventory"`
	AccountsPayables                       int64  `json:"accountsPayables"`
	OtherWorkingCapital                    int64  `json:"otherWorkingCapital"`
	OtherNonCashItems                      int64  `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities   int64  `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment int64  `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                        int64  `json:"acquisitionsNet"`
	PurchasesOfInvestments                 int64  `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments           int64  `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivities               int64  `json:"otherInvestingActivities"`
	NetCashProvidedByInvestingActivities   int64  `json:"netCashProvidedByInvestingActivities"`
	NetDebtIssuance                        int64  `json:"netDebtIssuance"`
	LongTermNetDebtIssuance                int64  `json:"longTermNetDebtIssuance"`
	ShortTermNetDebtIssuance               int64  `json:"shortTermNetDebtIssuance"`
	NetStockIssuance                       int64  `json:"netStockIssuance"`
	NetCommonStockIssuance                 int64  `json:"netCommonStockIssuance"`
	CommonStockIssuance                    int64  `json:"commonStockIssuance"`
	CommonStockRepurchased                 int64  `json:"commonStockRepurchased"`
	NetPreferredStockIssuance              int64  `json:"netPreferredStockIssuance"`
	NetDividendsPaid                       int64  `json:"netDividendsPaid"`
	CommonDividendsPaid                    int64  `json:"commonDividendsPaid"`
	PreferredDividendsPaid                 int64  `json:"preferredDividendsPaid"`
	OtherFinancingActivities               int64  `json:"otherFinancingActivities"`
	NetCashProvidedByFinancingActivities   int64  `json:"netCashProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash             int64  `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                        int64  `json:"netChangeInCash"`
	CashAtEndOfPeriod                      int64  `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                int64  `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                      int64  `json:"operatingCashFlow"`
	CapitalExpenditure                     int64  `json:"capitalExpenditure"`
	FreeCashFlow                           int64  `json:"freeCashFlow"`
	IncomeTaxesPaid                        int64  `json:"incomeTaxesPaid"`
	InterestPaid                           int64  `json:"interestPaid"`
}

// CashFlowStatement 는 종목의 현금흐름표 시계열을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) CashFlowStatement(ctx context.Context, p Params) ([]CashFlowStatement, error) {
	return fetchList[CashFlowStatement](ctx, c, "/stable/cash-flow-statement", p, p.queryParams())
}
