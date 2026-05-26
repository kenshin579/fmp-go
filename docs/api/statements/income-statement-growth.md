# Income Statement Growth

Track key financial growth metrics with the Income Statement Growth API. Analyze how revenue, profits, and expenses have evolved over time, offering insights into a company’s financial health and operational efficiency.

## Endpoint

`GET https://financialmodelingprep.com/stable/income-statement-growth?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | Q1,Q2,Q3,Q4,FY,annual,quarter |

## Description

The Income Statement Growth API provides critical growth data, allowing users to track year-over-year changes in key income statement items, such as:

- Revenue Growth: Monitor changes in a company's total revenue, helping gauge overall business performance.

- Profit Growth: Assess fluctuations in gross profit, operating income, and net income, offering insights into profitability trends.

- Expense Growth: Analyze growth in operating expenses, cost of revenue, and specific line items like research and development or interest expenses.

This API is a valuable tool for investors, analysts, and financial professionals who want to track a company's financial trends over time.

Example Use Case
A financial analyst can use the Income Statement Growth API to evaluate Apple's revenue and net income trends over the past few years, identifying whether the company is experiencing consistent growth or declines in profitability.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2024-09-28",
		"fiscalYear": "2024",
		"period": "FY",
		"reportedCurrency": "USD",
		"growthRevenue": 0.020219940775141214,
		"growthCostOfRevenue": -0.017675600199872046,
		"growthGrossProfit": 0.06819471705252206,
		"growthGrossProfitRatio": 0.04776303446712012,
		"growthResearchAndDevelopmentExpenses": 0.04863780712017383,
		"growthGeneralAndAdministrativeExpenses": 0,
		"growthSellingAndMarketingExpenses": 0,
		"growthOtherExpenses": -1,
		"growthOperatingExpenses": 0.04776924900176856,
		"growthCostAndExpenses": -0.004331112631234571,
		"growthInterestIncome": -1,
		"growthInterestExpense": -1,
		"growthDepreciationAndAmortization": -0.006424168764649709,
		"growthEBITDA": 0.07026704816404387,
		"growthOperatingIncome": 0.07799581805933456,
		"growthIncomeBeforeTax": 0.08571604417246959,
		"growthIncomeTaxExpense": 0.7770145152619318,
		"growthNetIncome": -0.033599670086086914,
		"growthEPS": -0.008116883116883088,
		"growthEPSDiluted": -0.008156606851549727,
		"growthWeightedAverageShsOut": -0.02543458616683152,
		"growthWeightedAverageShsOutDil": -0.02557791606880283,
		"growthEBIT": 0.0471407082579099,
		"growthNonOperatingIncomeExcludingInterest": 1,
		"growthNetInterestIncome": 1,
		"growthTotalOtherIncomeExpensesNet": 1.4761061946902654,
		"growthNetIncomeFromContinuingOperations": -0.033599670086086914,
		"growthOtherAdjustmentsToNetIncome": 0,
		"growthNetIncomeDeductions": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/income-statement-growth · 카테고리: statements
