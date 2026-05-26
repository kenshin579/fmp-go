# Cash Flow Statement Growth Bulk

The Cash Flow Statement Growth Bulk API allows you to retrieve bulk growth data for cash flow statements, enabling you to track changes in cash flows over time. This API is ideal for analyzing the cash flow growth trends of multiple companies simultaneously.

## Endpoint

`GET https://financialmodelingprep.com/stable/cash-flow-statement-growth-bulk?year=2026&period=Q1`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| year* | string | 2026 |
| period* | string | Q1,Q2,Q3,Q4,FY |

## Description

This API helps you:

- Track Growth Trends: Monitor changes in key cash flow metrics such as operating cash flow, capital expenditures, and free cash flow over time.

- Compare Company Performance: Quickly analyze the growth in cash flow activities for several companies, making it easier to identify high-growth firms or companies with declining cash flow.

- Understand Financial Health: Evaluate how companies are managing their cash flow, whether it's through improved operations or changes in investment and financing activities.

## Response (example)

```json
[
	{
		"symbol": "000001.SZ",
		"date": "2025-03-31",
		"fiscalYear": "2025",
		"period": "Q1",
		"reportedCurrency": "CNY",
		"growthNetIncome": "0",
		"growthDepreciationAndAmortization": "0",
		"growthDeferredIncomeTax": "0",
		"growthStockBasedCompensation": "0",
		"growthChangeInWorkingCapital": "0",
		"growthAccountsReceivables": "0",
		"growthInventory": "0",
		"growthAccountsPayables": "0",
		"growthOtherWorkingCapital": "0",
		"growthOtherNonCashItems": "3.2072823819457614",
		"growthNetCashProvidedByOperatingActivites": "3.2072823819457614",
		"growthInvestmentsInPropertyPlantAndEquipment": "0.7332280978689818",
		"growthAcquisitionsNet": "0",
		"growthPurchasesOfInvestments": "-0.12254537395030414",
		"growthSalesMaturitiesOfInvestments": "0.3847853673478318",
		"growthOtherInvestingActivites": "-0.8417721518987342",
		"growthNetCashUsedForInvestingActivites": "2.1699343339587243",
		"growthDebtRepayment": "1",
		"growthCommonStockIssued": "0",
		"growthCommonStockRepurchased": "0",
		"growthDividendsPaid": "0.6798284344644885",
		"growthOtherFinancingActivites": "-1.7077146619443309",
		"growthNetCashUsedProvidedByFinancingActivities": "-3.2122934677858628",
		"growthEffectOfForexChangesOnCash": "-1.0731570061902083",
		"growthNetChangeInCash": "2.348938711752274",
		"growthCashAtEndOfPeriod": "0.11426914604625096",
		"growthCashAtBeginningOfPeriod": "-0.07809495106059301",
		"growthOperatingCashFlow": "3.2072823819457614",
		"growthCapitalExpenditure": "0.7332280978689818",
		"growthFreeCashFlow": "3.16553689621649",
		"growthNetDebtIssuance": "1",
		"growthLongTermNetDebtIssuance": "1",
		"growthShortTermNetDebtIssuance": "0",
		"growthNetStockIssuance": "0",
		"growthPreferredDividendsPaid": "0.6798284344644885",
		"growthIncomeTaxesPaid": "0",
		"growthInterestPaid": "0"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cash-flow-statement-growth-bulk · 카테고리: bulk
