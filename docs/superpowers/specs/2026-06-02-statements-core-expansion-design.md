# FMP Go SDK — Statements-Core 확장 (v0.9.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/statements-core-expansion`
- 토픽: FMP `statements` 카테고리 중 3대 재무제표의 누락분/파생 뷰 8 endpoint 추가. 전체 API 커버리지 캠페인 7번째 그룹. statements 분해의 1/3.

## 배경 / 목적

statements 카테고리는 27 endpoint 로 한 PR 에 과대 → 3 하위 그룹 분해 확정:
1. **statements-core 확장 (v0.9.0, 본 스펙)** — cashflow + TTM 3종 + growth 4종 = 8.
2. metrics 그룹 (v0.10.0) — key-metrics(+ttm), metrics-ratios-ttm, financial-scores, owner-earnings, enterprise-values, revenue segments.
3. as-reported/reports 그룹 (v0.11.0) — as-reported 4종 + latest + financial-reports.

본 그룹은 현재 income/balance 만 있는 `statements/` 패키지에 현금흐름표를 채우고, 세 제표의 TTM·성장률 뷰를 더해 "완결된 3대 재무제표 세트"를 만든다. moneyflow 재무 분석에 직접 쓰인다.

## 결정 사항 (브레인스토밍)

- **범위**: 8 endpoint. 기존 `statements/` 패키지 확장(신규 패키지 아님).
- **구조체 재사용**: TTM 응답은 비-TTM 코어 제표와 필드 동일 → `IncomeStatement`/`BalanceSheetStatement`/`CashFlowStatement` 를 TTM 메서드가 그대로 반환. 신규 구조체는 `CashFlowStatement` + growth 4종뿐.
- **Params 재사용**: 기존 `Params{Symbol, Period, Limit}`. TTM 은 `period` 미지원 → `ttmQueryParams()`(symbol+limit) 추가.
- **내부 helper 통일**: `fetch.List` 는 빈 결과에 ErrNotFound 를 안 던져 기존 statements 시맨틱(len==0 → ErrNotFound)과 불일치. statements 패키지에 private generic helper `fetchList[T]` 추가하고, 기존 income/balance 도 여기에 맞춰 통일(같은 파일 내 동일 보일러플레이트 정리).
- **FMP 필드명 그대로**: JSON 태그는 카탈로그 그대로 보존(오타 포함: `...Activites`, `growthOthertotalStockholdersEquity`, `ebitgrowth`/`epsgrowth` 등). Go 필드명만 관례대로 정리.
- **nullable**: `FinancialStatementGrowth` 의 5개 필드(`ebitdaGrowth`, `growthCapitalExpenditure`, `tenY/fiveY/threeYBottomLineNetIncomeGrowthPerShare`)는 카탈로그 예시에서 null → `*float64`. 나머지 growth 는 값 존재 → `float64`.
- **필드 한국어 주석**: 코어 income/balance 가 주석을 안 단 관례를 따라, 대량 숫자 필드 구조체(cashflow/growth)는 구조체 단위 주석만 유지(필드별 주석 생략, 기존 income.go/balance.go 와 일관).
- **릴리스**: `v0.9.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | query | 반환 |
|---|---|---|---|---|
| `cashflow.go` | `CashFlowStatement(ctx, p)` | `/stable/cash-flow-statement` | symbol,period,limit | `[]CashFlowStatement` |
| `ttm.go` | `IncomeStatementTTM(ctx, p)` | `/stable/income-statement-ttm` | symbol,limit | `[]IncomeStatement` |
| | `BalanceSheetStatementTTM(ctx, p)` | `/stable/balance-sheet-statement-ttm` | symbol,limit | `[]BalanceSheetStatement` |
| | `CashFlowStatementTTM(ctx, p)` | `/stable/cash-flow-statement-ttm` | symbol,limit | `[]CashFlowStatement` |
| `growth.go` | `IncomeStatementGrowth(ctx, p)` | `/stable/income-statement-growth` | symbol,period,limit | `[]IncomeStatementGrowth` |
| | `BalanceSheetStatementGrowth(ctx, p)` | `/stable/balance-sheet-statement-growth` | symbol,period,limit | `[]BalanceSheetStatementGrowth` |
| | `CashFlowStatementGrowth(ctx, p)` | `/stable/cash-flow-statement-growth` | symbol,period,limit | `[]CashFlowStatementGrowth` |
| | `FinancialStatementGrowth(ctx, p)` | `/stable/financial-growth` | symbol,period,limit | `[]FinancialStatementGrowth` |

> 경로 주의: FMP 는 cashflow 를 하이픈 형태 `cash-flow-statement` 로, 재무성장률을 `financial-growth` 로 노출(파일명 슬러그와 다름). 카탈로그 `Endpoint:` 줄 기준.

## 내부 helper + Params 확장

```go
// statements/client.go 에 추가
// ttmQueryParams 는 TTM endpoint 용 — period 미지원이므로 symbol(+limit)만.
func (p Params) ttmQueryParams() map[string]string {
	q := map[string]string{"symbol": p.Symbol}
	if p.Limit > 0 {
		q["limit"] = strconv.Itoa(p.Limit)
	}
	return q
}

// fetchList 는 symbol 가드 후 GetJSON + 빈 결과 ErrNotFound 를 묶는다.
// 기존 income/balance 및 신규 cashflow/ttm/growth 가 공용.
func fetchList[T any](ctx context.Context, c *Client, path string, p Params, q map[string]string) ([]T, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []T
	if err := c.http.GetJSON(ctx, path, q, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
```

기존 income/balance 는 다음과 같이 단순화:
```go
func (c *Client) IncomeStatement(ctx context.Context, p Params) ([]IncomeStatement, error) {
	return fetchList[IncomeStatement](ctx, c, "/stable/income-statement", p, p.queryParams())
}
func (c *Client) BalanceSheetStatement(ctx context.Context, p Params) ([]BalanceSheetStatement, error) {
	return fetchList[BalanceSheetStatement](ctx, c, "/stable/balance-sheet-statement", p, p.queryParams())
}
```

## 응답 타입 (faithful, JSON 태그 카탈로그 그대로)

### CashFlowStatement (cashflow-statement / cash-flow-statement-ttm 공용)
```go
// CashFlowStatement 는 FMP 현금흐름표 한 기간(연간/분기/TTM). faithful 매핑.
type CashFlowStatement struct {
	Date                                  string `json:"date"`
	Symbol                                string `json:"symbol"`
	ReportedCurrency                      string `json:"reportedCurrency"`
	CIK                                   string `json:"cik"`
	FilingDate                            string `json:"filingDate"`
	AcceptedDate                          string `json:"acceptedDate"`
	FiscalYear                            string `json:"fiscalYear"`
	Period                                string `json:"period"`
	NetIncome                             int64  `json:"netIncome"`
	DepreciationAndAmortization           int64  `json:"depreciationAndAmortization"`
	DeferredIncomeTax                     int64  `json:"deferredIncomeTax"`
	StockBasedCompensation                int64  `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                int64  `json:"changeInWorkingCapital"`
	AccountsReceivables                   int64  `json:"accountsReceivables"`
	Inventory                             int64  `json:"inventory"`
	AccountsPayables                      int64  `json:"accountsPayables"`
	OtherWorkingCapital                   int64  `json:"otherWorkingCapital"`
	OtherNonCashItems                     int64  `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities  int64  `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment int64 `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                       int64  `json:"acquisitionsNet"`
	PurchasesOfInvestments                int64  `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments          int64  `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivities              int64  `json:"otherInvestingActivities"`
	NetCashProvidedByInvestingActivities  int64  `json:"netCashProvidedByInvestingActivities"`
	NetDebtIssuance                       int64  `json:"netDebtIssuance"`
	LongTermNetDebtIssuance               int64  `json:"longTermNetDebtIssuance"`
	ShortTermNetDebtIssuance              int64  `json:"shortTermNetDebtIssuance"`
	NetStockIssuance                      int64  `json:"netStockIssuance"`
	NetCommonStockIssuance                int64  `json:"netCommonStockIssuance"`
	CommonStockIssuance                   int64  `json:"commonStockIssuance"`
	CommonStockRepurchased                int64  `json:"commonStockRepurchased"`
	NetPreferredStockIssuance             int64  `json:"netPreferredStockIssuance"`
	NetDividendsPaid                      int64  `json:"netDividendsPaid"`
	CommonDividendsPaid                   int64  `json:"commonDividendsPaid"`
	PreferredDividendsPaid                int64  `json:"preferredDividendsPaid"`
	OtherFinancingActivities              int64  `json:"otherFinancingActivities"`
	NetCashProvidedByFinancingActivities  int64  `json:"netCashProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash            int64  `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                       int64  `json:"netChangeInCash"`
	CashAtEndOfPeriod                     int64  `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod               int64  `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                     int64  `json:"operatingCashFlow"`
	CapitalExpenditure                    int64  `json:"capitalExpenditure"`
	FreeCashFlow                          int64  `json:"freeCashFlow"`
	IncomeTaxesPaid                       int64  `json:"incomeTaxesPaid"`
	InterestPaid                          int64  `json:"interestPaid"`
}
```

### IncomeStatementGrowth (income-statement-growth)
```go
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
```

### BalanceSheetStatementGrowth (balance-sheet-statement-growth)
```go
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
```

### CashFlowStatementGrowth (cash-flow-statement-growth, FMP 오타 `Activites` 보존)
```go
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
```

### FinancialStatementGrowth (financial-growth, 마지막 5필드 nullable → *float64)
```go
type FinancialStatementGrowth struct {
	Symbol                                string   `json:"symbol"`
	Date                                  string   `json:"date"`
	FiscalYear                            string   `json:"fiscalYear"`
	Period                                string   `json:"period"`
	ReportedCurrency                      string   `json:"reportedCurrency"`
	RevenueGrowth                         float64  `json:"revenueGrowth"`
	GrossProfitGrowth                     float64  `json:"grossProfitGrowth"`
	EBITGrowth                            float64  `json:"ebitgrowth"`
	OperatingIncomeGrowth                 float64  `json:"operatingIncomeGrowth"`
	NetIncomeGrowth                       float64  `json:"netIncomeGrowth"`
	EPSGrowth                             float64  `json:"epsgrowth"`
	EPSDilutedGrowth                      float64  `json:"epsdilutedGrowth"`
	WeightedAverageSharesGrowth           float64  `json:"weightedAverageSharesGrowth"`
	WeightedAverageSharesDilutedGrowth    float64  `json:"weightedAverageSharesDilutedGrowth"`
	DividendsPerShareGrowth               float64  `json:"dividendsPerShareGrowth"`
	OperatingCashFlowGrowth               float64  `json:"operatingCashFlowGrowth"`
	ReceivablesGrowth                     float64  `json:"receivablesGrowth"`
	InventoryGrowth                       float64  `json:"inventoryGrowth"`
	AssetGrowth                           float64  `json:"assetGrowth"`
	BookValuePerShareGrowth               float64  `json:"bookValueperShareGrowth"`
	DebtGrowth                            float64  `json:"debtGrowth"`
	RDExpenseGrowth                       float64  `json:"rdexpenseGrowth"`
	SGAExpensesGrowth                     float64  `json:"sgaexpensesGrowth"`
	FreeCashFlowGrowth                    float64  `json:"freeCashFlowGrowth"`
	TenYRevenueGrowthPerShare             float64  `json:"tenYRevenueGrowthPerShare"`
	FiveYRevenueGrowthPerShare            float64  `json:"fiveYRevenueGrowthPerShare"`
	ThreeYRevenueGrowthPerShare           float64  `json:"threeYRevenueGrowthPerShare"`
	TenYOperatingCFGrowthPerShare         float64  `json:"tenYOperatingCFGrowthPerShare"`
	FiveYOperatingCFGrowthPerShare        float64  `json:"fiveYOperatingCFGrowthPerShare"`
	ThreeYOperatingCFGrowthPerShare       float64  `json:"threeYOperatingCFGrowthPerShare"`
	TenYNetIncomeGrowthPerShare           float64  `json:"tenYNetIncomeGrowthPerShare"`
	FiveYNetIncomeGrowthPerShare          float64  `json:"fiveYNetIncomeGrowthPerShare"`
	ThreeYNetIncomeGrowthPerShare         float64  `json:"threeYNetIncomeGrowthPerShare"`
	TenYShareholdersEquityGrowthPerShare  float64  `json:"tenYShareholdersEquityGrowthPerShare"`
	FiveYShareholdersEquityGrowthPerShare float64  `json:"fiveYShareholdersEquityGrowthPerShare"`
	ThreeYShareholdersEquityGrowthPerShare float64 `json:"threeYShareholdersEquityGrowthPerShare"`
	TenYDividendPerShareGrowthPerShare    float64  `json:"tenYDividendperShareGrowthPerShare"`
	FiveYDividendPerShareGrowthPerShare   float64  `json:"fiveYDividendperShareGrowthPerShare"`
	ThreeYDividendPerShareGrowthPerShare  float64  `json:"threeYDividendperShareGrowthPerShare"`
	EBITDAGrowth                          *float64 `json:"ebitdaGrowth"`                          // null 가능
	GrowthCapitalExpenditure              *float64 `json:"growthCapitalExpenditure"`              // null 가능
	TenYBottomLineNetIncomeGrowthPerShare *float64 `json:"tenYBottomLineNetIncomeGrowthPerShare"` // null 가능
	FiveYBottomLineNetIncomeGrowthPerShare *float64 `json:"fiveYBottomLineNetIncomeGrowthPerShare"` // null 가능
	ThreeYBottomLineNetIncomeGrowthPerShare *float64 `json:"threeYBottomLineNetIncomeGrowthPerShare"` // null 가능
}
```

## 테스트
- fixture 단위: cashflow / 4 growth — 핵심 필드 파싱 검증. FinancialStatementGrowth 는 null row + 값 row 양쪽 → 5개 포인터 null→nil 및 값 검증.
- TTM: income/balance/cashflow TTM 각 1건 — 코어 구조체 재사용 파싱 + path 가 `*-ttm` 인지, 쿼리에 `period` 없는지 검증(capturing).
- delegation: CashFlowStatement(p) path+symbol+period / 한 growth path / TTM 쿼리 period 생략.
- 가드: 빈 symbol 대표 1건(예: CashFlowStatement). 빈 결과 → ErrNotFound 1건.
- 기존 income/balance 테스트는 helper 리팩터 후에도 그대로 통과해야 함(회귀 없음).
- 통합(`//go:build integration`): CashFlowStatement(AAPL,annual) freeCashFlow!=0 / IncomeStatementTTM(AAPL) / FinancialStatementGrowth(AAPL,annual) — 마지막 5필드 실 null 여부 t.Logf.

## 문서 / 릴리스
- README 커버리지 표 Statements 행 갱신: IncomeStatement, BalanceSheetStatement, CashFlowStatement, (Income/Balance/CashFlow)TTM, (Income/Balance/CashFlow/Financial)Growth — 10 endpoint.
- `examples/` 는 별도 추가 안 함(statements 는 기존 income 예시 흐름; cashflow 한 줄 통합테스트로 충분). → examples 갱신 생략(YAGNI).
- 릴리스 `v0.9.0`.

## 범위 밖 / 위험
- metrics 그룹(v0.10.0), as-reported 그룹(v0.11.0) 은 별도 PR.
- TTM 응답의 `period` 필드는 본문엔 있으나 쿼리 파라미터로는 미지원 — `ttmQueryParams()` 가 생략.
- 일부 growth/cashflow JSON 키에 FMP 오타(`Activites`, `Othertotal`, `ebitgrowth`) — JSON 태그 그대로 보존(통합테스트로 매핑 확인).
- 거대 정수 매출/현금흐름이 드물게 소수로 올 가능성 — 기존 income/balance 가 int64 를 쓰므로 동일 정책 유지(문제 시 후속 조정).
- FinancialStatementGrowth 의 nullable 5필드는 카탈로그 예시 기준 — 통합테스트로 실제 null 여부 재확인.
```
