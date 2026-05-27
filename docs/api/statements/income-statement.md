# Income Statement

Access detailed income statement data for publicly traded companies with the Income Statements API. Track profitability, compare competitors, and identify business trends with up-to-date financial data.

## Endpoint

`GET https://financialmodelingprep.com/stable/income-statement?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | Q1,Q2,Q3,Q4,FY,annual,quarter |

## Description

The FMP Income Statements API provides comprehensive access to income statement data for a wide range of companies. This API is essential for:

- Profitability Tracking: Monitor a company's revenue, expenses, and net income over time. The income statement, also known as the profit and loss statement, provides a detailed view of a company's financial performance during a specific period.

- Competitive Analysis: Use the API to compare a company's financial performance to its competitors. By analyzing income statements across companies, investors can identify which businesses are leading in profitability and efficiency.

- Trend Identification: Detect trends in a company's business by examining changes in revenue, expenses, and net income over multiple periods. This data is crucial for understanding a company's financial health and growth prospects.

Example
Financial Ratio Calculation: An investor can use the Income Statements API to calculate key financial ratios, such as the price-to-earnings ratio (P/E ratio) and gross margin. These ratios help investors assess a company's valuation and profitability, enabling more informed investment decisions.

## Response (example)

```json
[
	{
		"date": "2024-09-28",
		"symbol": "AAPL",
		"reportedCurrency": "USD",
		"cik": "0000320193",
		"filingDate": "2024-11-01",
		"acceptedDate": "2024-11-01 06:01:36",
		"fiscalYear": "2024",
		"period": "FY",
		"revenue": 391035000000,
		"costOfRevenue": 210352000000,
		"grossProfit": 180683000000,
		"researchAndDevelopmentExpenses": 31370000000,
		"generalAndAdministrativeExpenses": 0,
		"sellingAndMarketingExpenses": 0,
		"sellingGeneralAndAdministrativeExpenses": 26097000000,
		"otherExpenses": 0,
		"operatingExpenses": 57467000000,
		"costAndExpenses": 267819000000,
		"netInterestIncome": 0,
		"interestIncome": 0,
		"interestExpense": 0,
		"depreciationAndAmortization": 11445000000,
		"ebitda": 134661000000,
		"ebit": 123216000000,
		"nonOperatingIncomeExcludingInterest": 0,
		"operatingIncome": 123216000000,
		"totalOtherIncomeExpensesNet": 269000000,
		"incomeBeforeTax": 123485000000,
		"incomeTaxExpense": 29749000000,
		"netIncomeFromContinuingOperations": 93736000000,
		"netIncomeFromDiscontinuedOperations": 0,
		"otherAdjustmentsToNetIncome": 0,
		"netIncome": 93736000000,
		"netIncomeDeductions": 0,
		"bottomLineNetIncome": 93736000000,
		"eps": 6.11,
		"epsDiluted": 6.08,
		"weightedAverageShsOut": 15343783000,
		"weightedAverageShsOutDil": 15408095000
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/income-statement · 카테고리: statements
