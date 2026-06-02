package statements

import "context"

// IncomeStatementGrowth — 손익계산서 성장률 (income-statement-growth)
type IncomeStatementGrowth struct {
	Symbol                                    string  `json:"symbol"`
	Date                                      string  `json:"date"`
	FiscalYear                                string  `json:"fiscalYear"`
	Period                                    string  `json:"period"`
	ReportedCurrency                          string  `json:"reportedCurrency"`
	GrowthRevenue                             float64 `json:"growthRevenue"`
	GrowthCostOfRevenue                       float64 `json:"growthCostOfRevenue"`
	GrowthGrossProfit                         float64 `json:"growthGrossProfit"`
	GrowthGrossProfitRatio                    float64 `json:"growthGrossProfitRatio"`
	GrowthResearchAndDevelopmentExpenses      float64 `json:"growthResearchAndDevelopmentExpenses"`
	GrowthGeneralAndAdministrativeExpenses    float64 `json:"growthGeneralAndAdministrativeExpenses"`
	GrowthSellingAndMarketingExpenses         float64 `json:"growthSellingAndMarketingExpenses"`
	GrowthOtherExpenses                       float64 `json:"growthOtherExpenses"`
	GrowthOperatingExpenses                   float64 `json:"growthOperatingExpenses"`
	GrowthCostAndExpenses                     float64 `json:"growthCostAndExpenses"`
	GrowthInterestIncome                      float64 `json:"growthInterestIncome"`
	GrowthInterestExpense                     float64 `json:"growthInterestExpense"`
	GrowthDepreciationAndAmortization         float64 `json:"growthDepreciationAndAmortization"`
	GrowthEBITDA                              float64 `json:"growthEBITDA"`
	GrowthOperatingIncome                     float64 `json:"growthOperatingIncome"`
	GrowthIncomeBeforeTax                     float64 `json:"growthIncomeBeforeTax"`
	GrowthIncomeTaxExpense                    float64 `json:"growthIncomeTaxExpense"`
	GrowthNetIncome                           float64 `json:"growthNetIncome"`
	GrowthEPS                                 float64 `json:"growthEPS"`
	GrowthEPSDiluted                          float64 `json:"growthEPSDiluted"`
	GrowthWeightedAverageShsOut               float64 `json:"growthWeightedAverageShsOut"`
	GrowthWeightedAverageShsOutDil            float64 `json:"growthWeightedAverageShsOutDil"`
	GrowthEBIT                                float64 `json:"growthEBIT"`
	GrowthNonOperatingIncomeExcludingInterest float64 `json:"growthNonOperatingIncomeExcludingInterest"`
	GrowthNetInterestIncome                   float64 `json:"growthNetInterestIncome"`
	GrowthTotalOtherIncomeExpensesNet         float64 `json:"growthTotalOtherIncomeExpensesNet"`
	GrowthNetIncomeFromContinuingOperations   float64 `json:"growthNetIncomeFromContinuingOperations"`
	GrowthOtherAdjustmentsToNetIncome         float64 `json:"growthOtherAdjustmentsToNetIncome"`
	GrowthNetIncomeDeductions                 float64 `json:"growthNetIncomeDeductions"`
}

// IncomeStatementGrowth 는 종목의 손익계산서 성장률 시계열을 조회한다.
func (c *Client) IncomeStatementGrowth(ctx context.Context, p Params) ([]IncomeStatementGrowth, error) {
	return fetchList[IncomeStatementGrowth](ctx, c, "/stable/income-statement-growth", p, p.queryParams())
}

// BalanceSheetStatementGrowth — 대차대조표 성장률 (balance-sheet-statement-growth)
type BalanceSheetStatementGrowth struct {
	Symbol                                        string  `json:"symbol"`
	Date                                          string  `json:"date"`
	FiscalYear                                    string  `json:"fiscalYear"`
	Period                                        string  `json:"period"`
	ReportedCurrency                              string  `json:"reportedCurrency"`
	GrowthCashAndCashEquivalents                  float64 `json:"growthCashAndCashEquivalents"`
	GrowthShortTermInvestments                    float64 `json:"growthShortTermInvestments"`
	GrowthCashAndShortTermInvestments             float64 `json:"growthCashAndShortTermInvestments"`
	GrowthNetReceivables                          float64 `json:"growthNetReceivables"`
	GrowthInventory                               float64 `json:"growthInventory"`
	GrowthOtherCurrentAssets                      float64 `json:"growthOtherCurrentAssets"`
	GrowthTotalCurrentAssets                      float64 `json:"growthTotalCurrentAssets"`
	GrowthPropertyPlantEquipmentNet               float64 `json:"growthPropertyPlantEquipmentNet"`
	GrowthGoodwill                                float64 `json:"growthGoodwill"`
	GrowthIntangibleAssets                        float64 `json:"growthIntangibleAssets"`
	GrowthGoodwillAndIntangibleAssets             float64 `json:"growthGoodwillAndIntangibleAssets"`
	GrowthLongTermInvestments                     float64 `json:"growthLongTermInvestments"`
	GrowthTaxAssets                               float64 `json:"growthTaxAssets"`
	GrowthOtherNonCurrentAssets                   float64 `json:"growthOtherNonCurrentAssets"`
	GrowthTotalNonCurrentAssets                   float64 `json:"growthTotalNonCurrentAssets"`
	GrowthOtherAssets                             float64 `json:"growthOtherAssets"`
	GrowthTotalAssets                             float64 `json:"growthTotalAssets"`
	GrowthAccountPayables                         float64 `json:"growthAccountPayables"`
	GrowthShortTermDebt                           float64 `json:"growthShortTermDebt"`
	GrowthTaxPayables                             float64 `json:"growthTaxPayables"`
	GrowthDeferredRevenue                         float64 `json:"growthDeferredRevenue"`
	GrowthOtherCurrentLiabilities                 float64 `json:"growthOtherCurrentLiabilities"`
	GrowthTotalCurrentLiabilities                 float64 `json:"growthTotalCurrentLiabilities"`
	GrowthLongTermDebt                            float64 `json:"growthLongTermDebt"`
	GrowthDeferredRevenueNonCurrent               float64 `json:"growthDeferredRevenueNonCurrent"`
	GrowthDeferredTaxLiabilitiesNonCurrent        float64 `json:"growthDeferredTaxLiabilitiesNonCurrent"`
	GrowthOtherNonCurrentLiabilities              float64 `json:"growthOtherNonCurrentLiabilities"`
	GrowthTotalNonCurrentLiabilities              float64 `json:"growthTotalNonCurrentLiabilities"`
	GrowthOtherLiabilities                        float64 `json:"growthOtherLiabilities"`
	GrowthTotalLiabilities                        float64 `json:"growthTotalLiabilities"`
	GrowthPreferredStock                          float64 `json:"growthPreferredStock"`
	GrowthCommonStock                             float64 `json:"growthCommonStock"`
	GrowthRetainedEarnings                        float64 `json:"growthRetainedEarnings"`
	GrowthAccumulatedOtherComprehensiveIncomeLoss float64 `json:"growthAccumulatedOtherComprehensiveIncomeLoss"`
	GrowthOthertotalStockholdersEquity            float64 `json:"growthOthertotalStockholdersEquity"`
	GrowthTotalStockholdersEquity                 float64 `json:"growthTotalStockholdersEquity"`
	GrowthMinorityInterest                        float64 `json:"growthMinorityInterest"`
	GrowthTotalEquity                             float64 `json:"growthTotalEquity"`
	GrowthTotalLiabilitiesAndStockholdersEquity   float64 `json:"growthTotalLiabilitiesAndStockholdersEquity"`
	GrowthTotalInvestments                        float64 `json:"growthTotalInvestments"`
	GrowthTotalDebt                               float64 `json:"growthTotalDebt"`
	GrowthNetDebt                                 float64 `json:"growthNetDebt"`
	GrowthAccountsReceivables                     float64 `json:"growthAccountsReceivables"`
	GrowthOtherReceivables                        float64 `json:"growthOtherReceivables"`
	GrowthPrepaids                                float64 `json:"growthPrepaids"`
	GrowthTotalPayables                           float64 `json:"growthTotalPayables"`
	GrowthOtherPayables                           float64 `json:"growthOtherPayables"`
	GrowthAccruedExpenses                         float64 `json:"growthAccruedExpenses"`
	GrowthCapitalLeaseObligationsCurrent          float64 `json:"growthCapitalLeaseObligationsCurrent"`
	GrowthAdditionalPaidInCapital                 float64 `json:"growthAdditionalPaidInCapital"`
	GrowthTreasuryStock                           float64 `json:"growthTreasuryStock"`
}

// BalanceSheetStatementGrowth 는 종목의 대차대조표 성장률 시계열을 조회한다.
func (c *Client) BalanceSheetStatementGrowth(ctx context.Context, p Params) ([]BalanceSheetStatementGrowth, error) {
	return fetchList[BalanceSheetStatementGrowth](ctx, c, "/stable/balance-sheet-statement-growth", p, p.queryParams())
}

// CashFlowStatementGrowth — 현금흐름표 성장률 (cash-flow-statement-growth). FMP 오타 Activites 보존.
type CashFlowStatementGrowth struct {
	Symbol                                         string  `json:"symbol"`
	Date                                           string  `json:"date"`
	FiscalYear                                     string  `json:"fiscalYear"`
	Period                                         string  `json:"period"`
	ReportedCurrency                               string  `json:"reportedCurrency"`
	GrowthNetIncome                                float64 `json:"growthNetIncome"`
	GrowthDepreciationAndAmortization              float64 `json:"growthDepreciationAndAmortization"`
	GrowthDeferredIncomeTax                        float64 `json:"growthDeferredIncomeTax"`
	GrowthStockBasedCompensation                   float64 `json:"growthStockBasedCompensation"`
	GrowthChangeInWorkingCapital                   float64 `json:"growthChangeInWorkingCapital"`
	GrowthAccountsReceivables                      float64 `json:"growthAccountsReceivables"`
	GrowthInventory                                float64 `json:"growthInventory"`
	GrowthAccountsPayables                         float64 `json:"growthAccountsPayables"`
	GrowthOtherWorkingCapital                      float64 `json:"growthOtherWorkingCapital"`
	GrowthOtherNonCashItems                        float64 `json:"growthOtherNonCashItems"`
	GrowthNetCashProvidedByOperatingActivites      float64 `json:"growthNetCashProvidedByOperatingActivites"`
	GrowthInvestmentsInPropertyPlantAndEquipment   float64 `json:"growthInvestmentsInPropertyPlantAndEquipment"`
	GrowthAcquisitionsNet                          float64 `json:"growthAcquisitionsNet"`
	GrowthPurchasesOfInvestments                   float64 `json:"growthPurchasesOfInvestments"`
	GrowthSalesMaturitiesOfInvestments             float64 `json:"growthSalesMaturitiesOfInvestments"`
	GrowthOtherInvestingActivites                  float64 `json:"growthOtherInvestingActivites"`
	GrowthNetCashUsedForInvestingActivites         float64 `json:"growthNetCashUsedForInvestingActivites"`
	GrowthDebtRepayment                            float64 `json:"growthDebtRepayment"`
	GrowthCommonStockIssued                        float64 `json:"growthCommonStockIssued"`
	GrowthCommonStockRepurchased                   float64 `json:"growthCommonStockRepurchased"`
	GrowthDividendsPaid                            float64 `json:"growthDividendsPaid"`
	GrowthOtherFinancingActivites                  float64 `json:"growthOtherFinancingActivites"`
	GrowthNetCashUsedProvidedByFinancingActivities float64 `json:"growthNetCashUsedProvidedByFinancingActivities"`
	GrowthEffectOfForexChangesOnCash               float64 `json:"growthEffectOfForexChangesOnCash"`
	GrowthNetChangeInCash                          float64 `json:"growthNetChangeInCash"`
	GrowthCashAtEndOfPeriod                        float64 `json:"growthCashAtEndOfPeriod"`
	GrowthCashAtBeginningOfPeriod                  float64 `json:"growthCashAtBeginningOfPeriod"`
	GrowthOperatingCashFlow                        float64 `json:"growthOperatingCashFlow"`
	GrowthCapitalExpenditure                       float64 `json:"growthCapitalExpenditure"`
	GrowthFreeCashFlow                             float64 `json:"growthFreeCashFlow"`
	GrowthNetDebtIssuance                          float64 `json:"growthNetDebtIssuance"`
	GrowthLongTermNetDebtIssuance                  float64 `json:"growthLongTermNetDebtIssuance"`
	GrowthShortTermNetDebtIssuance                 float64 `json:"growthShortTermNetDebtIssuance"`
	GrowthNetStockIssuance                         float64 `json:"growthNetStockIssuance"`
	GrowthPreferredDividendsPaid                   float64 `json:"growthPreferredDividendsPaid"`
	GrowthIncomeTaxesPaid                          float64 `json:"growthIncomeTaxesPaid"`
	GrowthInterestPaid                             float64 `json:"growthInterestPaid"`
}

// CashFlowStatementGrowth 는 종목의 현금흐름표 성장률 시계열을 조회한다.
func (c *Client) CashFlowStatementGrowth(ctx context.Context, p Params) ([]CashFlowStatementGrowth, error) {
	return fetchList[CashFlowStatementGrowth](ctx, c, "/stable/cash-flow-statement-growth", p, p.queryParams())
}

// FinancialStatementGrowth — 종합 재무 성장률 (financial-growth). 마지막 5필드 null 가능.
type FinancialStatementGrowth struct {
	Symbol                                  string   `json:"symbol"`
	Date                                    string   `json:"date"`
	FiscalYear                              string   `json:"fiscalYear"`
	Period                                  string   `json:"period"`
	ReportedCurrency                        string   `json:"reportedCurrency"`
	RevenueGrowth                           float64  `json:"revenueGrowth"`
	GrossProfitGrowth                       float64  `json:"grossProfitGrowth"`
	EBITGrowth                              float64  `json:"ebitgrowth"`
	OperatingIncomeGrowth                   float64  `json:"operatingIncomeGrowth"`
	NetIncomeGrowth                         float64  `json:"netIncomeGrowth"`
	EPSGrowth                               float64  `json:"epsgrowth"`
	EPSDilutedGrowth                        float64  `json:"epsdilutedGrowth"`
	WeightedAverageSharesGrowth             float64  `json:"weightedAverageSharesGrowth"`
	WeightedAverageSharesDilutedGrowth      float64  `json:"weightedAverageSharesDilutedGrowth"`
	DividendsPerShareGrowth                 float64  `json:"dividendsPerShareGrowth"`
	OperatingCashFlowGrowth                 float64  `json:"operatingCashFlowGrowth"`
	ReceivablesGrowth                       float64  `json:"receivablesGrowth"`
	InventoryGrowth                         float64  `json:"inventoryGrowth"`
	AssetGrowth                             float64  `json:"assetGrowth"`
	BookValuePerShareGrowth                 float64  `json:"bookValueperShareGrowth"`
	DebtGrowth                              float64  `json:"debtGrowth"`
	RDExpenseGrowth                         float64  `json:"rdexpenseGrowth"`
	SGAExpensesGrowth                       float64  `json:"sgaexpensesGrowth"`
	FreeCashFlowGrowth                      float64  `json:"freeCashFlowGrowth"`
	TenYRevenueGrowthPerShare               float64  `json:"tenYRevenueGrowthPerShare"`
	FiveYRevenueGrowthPerShare              float64  `json:"fiveYRevenueGrowthPerShare"`
	ThreeYRevenueGrowthPerShare             float64  `json:"threeYRevenueGrowthPerShare"`
	TenYOperatingCFGrowthPerShare           float64  `json:"tenYOperatingCFGrowthPerShare"`
	FiveYOperatingCFGrowthPerShare          float64  `json:"fiveYOperatingCFGrowthPerShare"`
	ThreeYOperatingCFGrowthPerShare         float64  `json:"threeYOperatingCFGrowthPerShare"`
	TenYNetIncomeGrowthPerShare             float64  `json:"tenYNetIncomeGrowthPerShare"`
	FiveYNetIncomeGrowthPerShare            float64  `json:"fiveYNetIncomeGrowthPerShare"`
	ThreeYNetIncomeGrowthPerShare           float64  `json:"threeYNetIncomeGrowthPerShare"`
	TenYShareholdersEquityGrowthPerShare    float64  `json:"tenYShareholdersEquityGrowthPerShare"`
	FiveYShareholdersEquityGrowthPerShare   float64  `json:"fiveYShareholdersEquityGrowthPerShare"`
	ThreeYShareholdersEquityGrowthPerShare  float64  `json:"threeYShareholdersEquityGrowthPerShare"`
	TenYDividendPerShareGrowthPerShare      float64  `json:"tenYDividendperShareGrowthPerShare"`
	FiveYDividendPerShareGrowthPerShare     float64  `json:"fiveYDividendperShareGrowthPerShare"`
	ThreeYDividendPerShareGrowthPerShare    float64  `json:"threeYDividendperShareGrowthPerShare"`
	EBITDAGrowth                            *float64 `json:"ebitdaGrowth"`
	GrowthCapitalExpenditure                *float64 `json:"growthCapitalExpenditure"`
	TenYBottomLineNetIncomeGrowthPerShare   *float64 `json:"tenYBottomLineNetIncomeGrowthPerShare"`
	FiveYBottomLineNetIncomeGrowthPerShare  *float64 `json:"fiveYBottomLineNetIncomeGrowthPerShare"`
	ThreeYBottomLineNetIncomeGrowthPerShare *float64 `json:"threeYBottomLineNetIncomeGrowthPerShare"`
}

// FinancialStatementGrowth 는 종목의 종합 재무 성장률 시계열을 조회한다.
func (c *Client) FinancialStatementGrowth(ctx context.Context, p Params) ([]FinancialStatementGrowth, error) {
	return fetchList[FinancialStatementGrowth](ctx, c, "/stable/financial-growth", p, p.queryParams())
}
