# Balance Sheet Statement Growth

Analyze the growth of key balance sheet items over time with the Balance Sheet Statement Growth API. Track changes in assets, liabilities, and equity to understand the financial evolution of a company.

## Endpoint

`GET https://financialmodelingprep.com/stable/balance-sheet-statement-growth?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | Q1,Q2,Q3,Q4,FY,annual,quarter |

## Description

The Balance Sheet Statement Growth API provides year-over-year growth metrics for key balance sheet components. This API is ideal for:

- Asset Growth Analysis: Track changes in assets, such as cash, inventory, and long-term investments, to assess how a company's resources are expanding or contracting.

- Liability Growth Monitoring: Understand how short-term and long-term liabilities are evolving, including payables and debt.

- Equity Growth Tracking: Monitor shifts in shareholder equity, retained earnings, and total equity, offering insights into a company's financial health.

This API helps financial analysts and investors evaluate a company's stability and growth by examining the evolution of its balance sheet items.

Example Use Case
An investor can use the Balance Sheet Statement Growth API to analyze how Apple's cash reserves and debt levels have changed over the past year, helping them assess the company's liquidity and financial health.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2024-09-28",
		"fiscalYear": "2024",
		"period": "FY",
		"reportedCurrency": "USD",
		"growthCashAndCashEquivalents": -0.0007341898882029034,
		"growthShortTermInvestments": 0.11516302627413738,
		"growthCashAndShortTermInvestments": 0.058744212492892536,
		"growthNetReceivables": 0.08621792243994425,
		"growthInventory": 0.15084504817564365,
		"growthOtherCurrentAssets": -0.02776454576386526,
		"growthTotalCurrentAssets": 0.06562138667929733,
		"growthPropertyPlantEquipmentNet": -0.15992349565984992,
		"growthGoodwill": 0,
		"growthIntangibleAssets": 0,
		"growthGoodwillAndIntangibleAssets": 0,
		"growthLongTermInvestments": -0.09015953214513049,
		"growthTaxAssets": 0.09225857046829487,
		"growthOtherNonCurrentAssets": 0.5266933370120016,
		"growthTotalNonCurrentAssets": 0.014238076328719674,
		"growthOtherAssets": 0,
		"growthTotalAssets": 0.035160515396374756,
		"growthAccountPayables": 0.1014039066617687,
		"growthShortTermDebt": 0.32087050041121024,
		"growthTaxPayables": 2.01632838190271,
		"growthDeferredRevenue": 0.023322168465450935,
		"growthOtherCurrentLiabilities": -0.1254584832500786,
		"growthTotalCurrentLiabilities": 0.21391802240757563,
		"growthLongTermDebt": -0.10003043628845205,
		"growthDeferredRevenueNonCurrent": 0,
		"growthDeferredTaxLiabilitiesNonCurrent": 0,
		"growthOtherNonCurrentLiabilities": -0.09048495373370312,
		"growthTotalNonCurrentLiabilities": -0.09295867814151548,
		"growthOtherLiabilities": 0,
		"growthTotalLiabilities": 0.060574238130816666,
		"growthPreferredStock": 0,
		"growthCommonStock": 0.12821763398905328,
		"growthRetainedEarnings": -88.50467289719626,
		"growthAccumulatedOtherComprehensiveIncomeLoss": 0.3737338456164862,
		"growthOthertotalStockholdersEquity": 0,
		"growthTotalStockholdersEquity": -0.0836095645737457,
		"growthMinorityInterest": 0,
		"growthTotalEquity": -0.0836095645737457,
		"growthTotalLiabilitiesAndStockholdersEquity": 0.035160515396374756,
		"growthTotalInvestments": -0.04107194211936368,
		"growthTotalDebt": -0.0401393489845888,
		"growthNetDebt": -0.05469472282829777,
		"growthAccountsReceivables": 0.13223532601328453,
		"growthOtherReceivables": 0.04307907360930203,
		"growthPrepaids": 0,
		"growthTotalPayables": 0.5262653527335452,
		"growthOtherPayables": 0,
		"growthAccruedExpenses": 0,
		"growthCapitalLeaseObligationsCurrent": 0.03619047619047619,
		"growthAdditionalPaidInCapital": 0,
		"growthTreasuryStock": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/balance-sheet-statement-growth · 카테고리: statements
