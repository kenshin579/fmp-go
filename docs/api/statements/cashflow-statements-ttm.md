# Cashflow Statements TTM

## Endpoint

`GET https://financialmodelingprep.com/stable/cash-flow-statement-ttm?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |

## Response (example)

```json
[
	{
		"date": "2024-12-28",
		"symbol": "AAPL",
		"reportedCurrency": "USD",
		"cik": "0000320193",
		"filingDate": "2025-01-31",
		"acceptedDate": "2025-01-31 06:01:27",
		"fiscalYear": "2025",
		"period": "Q1",
		"netIncome": 96150000000,
		"depreciationAndAmortization": 11677000000,
		"deferredIncomeTax": 0,
		"stockBasedCompensation": 11977000000,
		"changeInWorkingCapital": -8224000000,
		"accountsReceivables": -9505000000,
		"inventory": -694000000,
		"accountsPayables": 3891000000,
		"otherWorkingCapital": -1916000000,
		"otherNonCashItems": -3286000000,
		"netCashProvidedByOperatingActivities": 108294000000,
		"investmentsInPropertyPlantAndEquipment": -9995000000,
		"acquisitionsNet": 0,
		"purchasesOfInvestments": -45000000000,
		"salesMaturitiesOfInvestments": 67422000000,
		"otherInvestingActivities": -1627000000,
		"netCashProvidedByInvestingActivities": 10800000000,
		"netDebtIssuance": -10967000000,
		"longTermNetDebtIssuance": -10967000000,
		"shortTermNetDebtIssuance": 0,
		"netStockIssuance": -98416000000,
		"netCommonStockIssuance": -98416000000,
		"commonStockIssuance": 0,
		"commonStockRepurchased": -98416000000,
		"netPreferredStockIssuance": 0,
		"netDividendsPaid": -15265000000,
		"commonDividendsPaid": -15265000000,
		"preferredDividendsPaid": 0,
		"otherFinancingActivities": -6121000000,
		"netCashProvidedByFinancingActivities": -130769000000,
		"effectOfForexChangesOnCash": 0,
		"netChangeInCash": -11675000000,
		"cashAtEndOfPeriod": 30299000000,
		"cashAtBeginningOfPeriod": 41974000000,
		"operatingCashFlow": 108294000000,
		"capitalExpenditure": -9995000000,
		"freeCashFlow": 98299000000,
		"incomeTaxesPaid": 37498000000,
		"interestPaid": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cashflow-statements-ttm · 카테고리: statements
