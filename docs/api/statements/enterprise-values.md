# Enterprise Values

Access a company's enterprise value using the Enterprise Values API. This metric offers a comprehensive view of a company's total market value by combining both its equity (market capitalization) and debt, providing a better understanding of its worth.

## Endpoint

`GET https://financialmodelingprep.com/stable/enterprise-values?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | Q1,Q2,Q3,Q4,FY,annual,quarter |

## Description

The Enterprise Values API provides key financial data to help assess a company's value by including:

- Market Capitalization: The total value of all outstanding shares based on the current stock price.

- Debt & Cash: Includes total debt and subtracts cash and cash equivalents to get a full picture of a company's financial standing.

- Comprehensive Valuation: Enterprise value includes both equity and debt, making it a preferred measure for evaluating potential buyouts, mergers, or acquisitions.

This API is ideal for analysts, investors, and finance professionals who need a complete understanding of a company's valuation, especially when considering its overall market position.

Example Use Case
A financial analyst uses the Enterprise Values API to assess Apple's total market value, factoring in debt and subtracting cash reserves, to determine whether it's a good acquisition target.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2024-09-28",
		"stockPrice": 227.79,
		"numberOfShares": 15343783000,
		"marketCapitalization": 3495160329570,
		"minusCashAndCashEquivalents": 29943000000,
		"addTotalDebt": 106629000000,
		"enterpriseValue": 3571846329570
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/enterprise-values · 카테고리: statements
