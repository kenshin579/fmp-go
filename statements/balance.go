package statements

import "context"

// BalanceSheetStatement 는 FMP /stable/balance-sheet-statement 응답 한 기간.
// 필드는 FMP 응답을 충실히 매핑한다(faithful).
type BalanceSheetStatement struct {
	Date                                    string `json:"date"`
	Symbol                                  string `json:"symbol"`
	ReportedCurrency                        string `json:"reportedCurrency"`
	CIK                                     string `json:"cik"`
	FilingDate                              string `json:"filingDate"`
	AcceptedDate                            string `json:"acceptedDate"`
	FiscalYear                              string `json:"fiscalYear"`
	Period                                  string `json:"period"`
	CashAndCashEquivalents                  int64  `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    int64  `json:"shortTermInvestments"`
	CashAndShortTermInvestments             int64  `json:"cashAndShortTermInvestments"`
	NetReceivables                          int64  `json:"netReceivables"`
	AccountsReceivables                     int64  `json:"accountsReceivables"`
	OtherReceivables                        int64  `json:"otherReceivables"`
	Inventory                               int64  `json:"inventory"`
	Prepaids                                int64  `json:"prepaids"`
	OtherCurrentAssets                      int64  `json:"otherCurrentAssets"`
	TotalCurrentAssets                      int64  `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               int64  `json:"propertyPlantEquipmentNet"`
	Goodwill                                int64  `json:"goodwill"`
	IntangibleAssets                        int64  `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             int64  `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     int64  `json:"longTermInvestments"`
	TaxAssets                               int64  `json:"taxAssets"`
	OtherNonCurrentAssets                   int64  `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   int64  `json:"totalNonCurrentAssets"`
	OtherAssets                             int64  `json:"otherAssets"`
	TotalAssets                             int64  `json:"totalAssets"`
	TotalPayables                           int64  `json:"totalPayables"`
	AccountPayables                         int64  `json:"accountPayables"`
	OtherPayables                           int64  `json:"otherPayables"`
	AccruedExpenses                         int64  `json:"accruedExpenses"`
	ShortTermDebt                           int64  `json:"shortTermDebt"`
	CapitalLeaseObligationsCurrent          int64  `json:"capitalLeaseObligationsCurrent"`
	TaxPayables                             int64  `json:"taxPayables"`
	DeferredRevenue                         int64  `json:"deferredRevenue"`
	OtherCurrentLiabilities                 int64  `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 int64  `json:"totalCurrentLiabilities"`
	LongTermDebt                            int64  `json:"longTermDebt"`
	DeferredRevenueNonCurrent               int64  `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        int64  `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              int64  `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              int64  `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        int64  `json:"otherLiabilities"`
	CapitalLeaseObligations                 int64  `json:"capitalLeaseObligations"`
	TotalLiabilities                        int64  `json:"totalLiabilities"`
	TreasuryStock                           int64  `json:"treasuryStock"`
	PreferredStock                          int64  `json:"preferredStock"`
	CommonStock                             int64  `json:"commonStock"`
	RetainedEarnings                        int64  `json:"retainedEarnings"`
	AdditionalPaidInCapital                 int64  `json:"additionalPaidInCapital"`
	AccumulatedOtherComprehensiveIncomeLoss int64  `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OtherTotalStockholdersEquity            int64  `json:"otherTotalStockholdersEquity"`
	TotalStockholdersEquity                 int64  `json:"totalStockholdersEquity"`
	TotalEquity                             int64  `json:"totalEquity"`
	MinorityInterest                        int64  `json:"minorityInterest"`
	TotalLiabilitiesAndTotalEquity          int64  `json:"totalLiabilitiesAndTotalEquity"`
	TotalInvestments                        int64  `json:"totalInvestments"`
	TotalDebt                               int64  `json:"totalDebt"`
	NetDebt                                 int64  `json:"netDebt"`
}

// BalanceSheetStatement 는 종목의 대차대조표 시계열을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) BalanceSheetStatement(ctx context.Context, p Params) ([]BalanceSheetStatement, error) {
	return fetchList[BalanceSheetStatement](ctx, c, "/stable/balance-sheet-statement", p, p.queryParams())
}
