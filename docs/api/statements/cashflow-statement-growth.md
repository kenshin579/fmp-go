# Cashflow Statement Growth

Measure the growth rate of a company’s cash flow with the FMP Cashflow Statement Growth API. Determine how quickly a company’s cash flow is increasing or decreasing over time.

## Endpoint

`GET https://financialmodelingprep.com/stable/cash-flow-statement-growth?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | Q1,Q2,Q3,Q4,FY,annual,quarter |

## Description

The FMP Cashflow Statement Growth API provides key insights into the cash flow growth rate of a company, an essential metric for assessing a company's financial health. This API is crucial for:

- Financial Performance Evaluation: Analyze the rate at which a company's cash flow is growing. A positive growth rate indicates that the company is generating more cash than it is using, which can signal strong financial health and operational efficiency.

- Investment Decision-Making: Use cash flow growth data to identify companies with strong cash flow generation capabilities. Companies with consistent positive cash flow growth are often more stable and may represent good investment opportunities.

- Risk Assessment: A negative cash flow growth rate can be a red flag, indicating that a company is using more cash than it is generating. This information can be used to evaluate the risk associated with investing in or continuing to hold a company's stock.

Example
Investor Analysis: An investor might use the Cashflow Growth API to assess a manufacturing company's financial health by examining its cash flow growth over the past five years. If the company shows consistent positive growth, the investor may decide to increase their investment in the company.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2024-09-28",
		"fiscalYear": "2024",
		"period": "FY",
		"reportedCurrency": "USD",
		"growthNetIncome": -0.033599670086086914,
		"growthDepreciationAndAmortization": -0.006424168764649709,
		"growthDeferredIncomeTax": 0,
		"growthStockBasedCompensation": 0.07892550540016616,
		"growthChangeInWorkingCapital": 1.555116314429071,
		"growthAccountsReceivables": -2.0473933649289098,
		"growthInventory": 0.3535228677379481,
		"growthAccountsPayables": 4.1868713605082055,
		"growthOtherWorkingCapital": 2.4402563136072373,
		"growthOtherNonCashItems": -0.017512348450830714,
		"growthNetCashProvidedByOperatingActivites": 0.06975566069312394,
		"growthInvestmentsInPropertyPlantAndEquipment": 0.13796879277306323,
		"growthAcquisitionsNet": 0,
		"growthPurchasesOfInvestments": -0.6486294175448107,
		"growthSalesMaturitiesOfInvestments": 0.3698202750801951,
		"growthOtherInvestingActivites": 0.02169035153328347,
		"growthNetCashUsedForInvestingActivites": -0.2078272604588394,
		"growthDebtRepayment": -0.012662502110417018,
		"growthCommonStockIssued": 0,
		"growthCommonStockRepurchased": -0.2243584784010316,
		"growthDividendsPaid": -0.013910149750415973,
		"growthOtherFinancingActivites": 0.03493013972055888,
		"growthNetCashUsedProvidedByFinancingActivities": -0.12439163778482412,
		"growthEffectOfForexChangesOnCash": 0,
		"growthNetChangeInCash": -1.1378472222222222,
		"growthCashAtEndOfPeriod": -0.02583205908188828,
		"growthCashAtBeginningOfPeriod": 0.23061216319013492,
		"growthOperatingCashFlow": 0.06975566069312394,
		"growthCapitalExpenditure": 0.13796879277306323,
		"growthFreeCashFlow": 0.092615279562982,
		"growthNetDebtIssuance": 0.3942026057973942,
		"growthLongTermNetDebtIssuance": -0.6812426135404356,
		"growthShortTermNetDebtIssuance": 1.995475113122172,
		"growthNetStockIssuance": -0.2243584784010316,
		"growthPreferredDividendsPaid": -0.013910149750415973,
		"growthIncomeTaxesPaid": 0.3973981476524439,
		"growthInterestPaid": -1
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cashflow-statement-growth · 카테고리: statements
