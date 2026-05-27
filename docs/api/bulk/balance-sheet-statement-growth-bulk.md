# Balance Sheet Statement Growth Bulk

The Balance Sheet Growth Bulk API allows users to retrieve growth data across multiple companies’ balance sheets, enabling detailed analysis of how financial positions have changed over time.

## Endpoint

`GET https://financialmodelingprep.com/stable/balance-sheet-statement-growth-bulk?year=2026&period=Q1`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| year* | string | 2026 |
| period* | string | Q1,Q2,Q3,Q4,FY |

## Description

This API is designed for:

- Trend Analysis: Track the growth or decline of financial metrics such as cash and short-term investments, receivables, total liabilities, and equity.

- Comparative Insights: Analyze changes in financial positions across multiple companies over different periods to spot trends, risks, and opportunities.

- Long-Term Financial Health Assessment: Assess how a company's balance sheet has evolved, providing deeper insights into its long-term financial stability.

This API is essential for tracking the development of assets, liabilities, and equity, providing insights into a company's financial trajectory.

## Response (example)

```json
[
	{
		"symbol": "000001.SZ",
		"date": "2025-03-31",
		"fiscalYear": "2025",
		"period": "Q1",
		"reportedCurrency": "CNY",
		"growthCashAndCashEquivalents": "0.09574482145872953",
		"growthShortTermInvestments": "0",
		"growthCashAndShortTermInvestments": "0.09574482145872953",
		"growthNetReceivables": "0",
		"growthInventory": "0",
		"growthOtherCurrentAssets": "0",
		"growthTotalCurrentAssets": "0.09574482145872953",
		"growthPropertyPlantEquipmentNet": "-0.06373337231398918",
		"growthGoodwill": "0",
		"growthIntangibleAssets": "-0.03270278935556268",
		"growthGoodwillAndIntangibleAssets": "-0.01477618426770969",
		"growthLongTermInvestments": "-0.0774117797082201",
		"growthTaxAssets": "0",
		"growthOtherNonCurrentAssets": "0.07678934705504345",
		"growthTotalNonCurrentAssets": "-0.01112505367669385",
		"growthOtherAssets": "0.001488576544346165",
		"growthTotalAssets": "0.001488576544346165",
		"growthAccountPayables": "0",
		"growthShortTermDebt": "0",
		"growthTaxPayables": "-0.0279424216765453",
		"growthDeferredRevenue": "0",
		"growthOtherCurrentLiabilities": "0.12022416350749959",
		"growthTotalCurrentLiabilities": "0",
		"growthLongTermDebt": "0",
		"growthDeferredRevenueNonCurrent": "0",
		"growthDeferredTaxLiabilitiesNonCurrent": "0",
		"growthOtherNonCurrentLiabilities": "0",
		"growthTotalNonCurrentLiabilities": "0",
		"growthOtherLiabilities": "-0.0005084911577141635",
		"growthTotalLiabilities": "-0.0005084911577141635",
		"growthPreferredStock": "0",
		"growthCommonStock": "0",
		"growthRetainedEarnings": "0.049325752755485314",
		"growthAccumulatedOtherComprehensiveIncomeLoss": "0",
		"growthOthertotalStockholdersEquity": "-0.0035208940994345805",
		"growthTotalStockholdersEquity": "0.022774946346510602",
		"growthMinorityInterest": "0",
		"growthTotalEquity": "0.022774946346510602",
		"growthTotalLiabilitiesAndStockholdersEquity": "0.001488576544346165",
		"growthTotalInvestments": "-0.0774117797082201",
		"growthTotalDebt": "0",
		"growthNetDebt": "-0.09574482145872953",
		"growthAccountsReceivables": "0",
		"growthOtherReceivables": "0",
		"growthPrepaids": "0",
		"growthTotalPayables": "-0.12022416350749959",
		"growthOtherPayables": "-0.12022416350749959",
		"growthAccruedExpenses": "0",
		"growthCapitalLeaseObligationsCurrent": "0",
		"growthAdditionalPaidInCapital": "0",
		"growthTreasuryStock": "0"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/balance-sheet-statement-growth-bulk · 카테고리: bulk
