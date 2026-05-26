# Cash Flow Statement Bulk

The Cash Flow Statement Bulk API provides access to detailed cash flow reports for a wide range of companies. This API enables users to retrieve bulk cash flow statement data, helping to analyze companies’ operating, investing, and financing activities over time.

## Endpoint

`GET https://financialmodelingprep.com/stable/cash-flow-statement-bulk?year=2026&period=Q1`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| year* | string | 2026 |
| period* | string | Q1,Q2,Q3,Q4,FY |

## Description

This API is essential for:

- Tracking Cash Movements: Understand how a company generates and uses cash in its operations, investments, and financing activities.

- Free Cash Flow Analysis: Analyze free cash flow to assess a company's ability to generate cash after accounting for capital expenditures.

- Comparative Analysis: Access data for multiple companies at once to compare their cash flow trends, helping to identify companies with strong or weak cash flow management.

## Response (example)

```json
[
	{
		"date": "2025-03-31",
		"symbol": "000001.SZ",
		"reportedCurrency": "CNY",
		"cik": "0000000000",
		"filingDate": "2025-03-31",
		"acceptedDate": "2025-03-31 00:00:00",
		"fiscalYear": "2025",
		"period": "Q1",
		"netIncome": "0",
		"depreciationAndAmortization": "0",
		"deferredIncomeTax": "0",
		"stockBasedCompensation": "0",
		"changeInWorkingCapital": "0",
		"accountsReceivables": "0",
		"inventory": "0",
		"accountsPayables": "0",
		"otherWorkingCapital": "0",
		"otherNonCashItems": "162946000000",
		"netCashProvidedByOperatingActivities": "162946000000",
		"investmentsInPropertyPlantAndEquipment": "-338000000",
		"acquisitionsNet": "0",
		"purchasesOfInvestments": "-227916000000",
		"salesMaturitiesOfInvestments": "253172000000",
		"otherInvestingActivities": "25000000",
		"netCashProvidedByInvestingActivities": "24943000000",
		"netDebtIssuance": "0",
		"longTermNetDebtIssuance": "0",
		"shortTermNetDebtIssuance": "0",
		"netStockIssuance": "0",
		"netCommonStockIssuance": "0",
		"commonStockIssuance": "0",
		"commonStockRepurchased": "0",
		"netPreferredStockIssuance": "0",
		"netDividendsPaid": "-2538000000",
		"commonDividendsPaid": "-2538000000",
		"preferredDividendsPaid": "0",
		"otherFinancingActivities": "-155860000000",
		"netCashProvidedByFinancingActivities": "-158398000000",
		"effectOfForexChangesOnCash": "-130000000",
		"netChangeInCash": "29361000000",
		"cashAtEndOfPeriod": "286307000000",
		"cashAtBeginningOfPeriod": "256946000000",
		"operatingCashFlow": "162946000000",
		"capitalExpenditure": "-338000000",
		"freeCashFlow": "162608000000",
		"incomeTaxesPaid": "0",
		"interestPaid": "0"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cash-flow-statement-bulk · 카테고리: bulk
