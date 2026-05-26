# Financial Statement Growth

Analyze the growth of key financial statement items across income, balance sheet, and cash flow statements with the Financial Statement Growth API. Track changes over time to understand trends in financial performance.

## Endpoint

`GET https://financialmodelingprep.com/stable/financial-growth?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | Q1,Q2,Q3,Q4,FY,annual,quarter |

## Description

The Financial Statement Growth API provides an overview of year-over-year growth in key financial metrics from income statements, balance sheets, and cash flow statements. It's designed for analysts and investors who want to:

- Assess Revenue Trends: See how a company's revenue has grown or contracted over time, highlighting overall business health.

- Evaluate Profitability Growth: Track growth in net income, operating income, and EBIT to gauge profitability.

- Monitor Asset & Debt Changes: Understand the growth or reduction in assets and liabilities, providing insights into financial management.

- Examine Cash Flow Changes: View growth in operating cash flow and free cash flow to analyze liquidity and capital efficiency.

This API helps in identifying long-term trends across financial statements, providing a comprehensive picture of a company's financial growth.

Example Use Case
An investor can use the Financial Statement Growth API to analyze Apple's revenue, net income, and free cash flow growth over the past few years, helping them assess the company's performance trends.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2024-09-28",
		"fiscalYear": "2024",
		"period": "FY",
		"reportedCurrency": "USD",
		"revenueGrowth": 0.020219940775141214,
		"grossProfitGrowth": 0.06819471705252206,
		"ebitgrowth": 0.07799581805933456,
		"operatingIncomeGrowth": 0.07799581805933456,
		"netIncomeGrowth": -0.033599670086086914,
		"epsgrowth": -0.008116883116883088,
		"epsdilutedGrowth": -0.008156606851549727,
		"weightedAverageSharesGrowth": -0.02543458616683152,
		"weightedAverageSharesDilutedGrowth": -0.02557791606880283,
		"dividendsPerShareGrowth": 0.040371570095532654,
		"operatingCashFlowGrowth": 0.06975566069312394,
		"receivablesGrowth": 0.08621792243994425,
		"inventoryGrowth": 0.15084504817564365,
		"assetGrowth": 0.035160515396374756,
		"bookValueperShareGrowth": -0.059693251557224776,
		"debtGrowth": -0.0401393489845888,
		"rdexpenseGrowth": 0.04863780712017383,
		"sgaexpensesGrowth": 0.04672709770575967,
		"freeCashFlowGrowth": 0.092615279562982,
		"tenYRevenueGrowthPerShare": 2.3937532854122625,
		"fiveYRevenueGrowthPerShare": 0.8093292228858464,
		"threeYRevenueGrowthPerShare": 0.163506592883552,
		"tenYOperatingCFGrowthPerShare": 2.1417809176982403,
		"fiveYOperatingCFGrowthPerShare": 1.051533221923415,
		"threeYOperatingCFGrowthPerShare": 0.23720294833900227,
		"tenYNetIncomeGrowthPerShare": 2.76381558093543,
		"fiveYNetIncomeGrowthPerShare": 1.0421744314966246,
		"threeYNetIncomeGrowthPerShare": 0.07761907162786884,
		"tenYShareholdersEquityGrowthPerShare": -0.19003774225234785,
		"fiveYShareholdersEquityGrowthPerShare": -0.24235004889283715,
		"threeYShareholdersEquityGrowthPerShare": -0.017459858915902907,
		"tenYDividendperShareGrowthPerShare": 1.1722201809466772,
		"fiveYDividendperShareGrowthPerShare": 0.29890046876764864,
		"threeYDividendperShareGrowthPerShare": 0.14617932692103452,
		"ebitdaGrowth": null,
		"growthCapitalExpenditure": null,
		"tenYBottomLineNetIncomeGrowthPerShare": null,
		"fiveYBottomLineNetIncomeGrowthPerShare": null,
		"threeYBottomLineNetIncomeGrowthPerShare": null
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/financial-statement-growth · 카테고리: statements
