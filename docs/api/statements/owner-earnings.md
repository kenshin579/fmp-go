# Owner Earnings

Retrieve a company's owner earnings with the Owner Earnings API, which provides a more accurate representation of cash available to shareholders by adjusting net income. This metric is crucial for evaluating a company’s profitability from the perspective of investors.

## Endpoint

`GET https://financialmodelingprep.com/stable/owner-earnings?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |

## Description

The Owner Earnings API offers a detailed breakdown of a company's cash flow adjusted for key factors, such as capital expenditures and depreciation. It is designed for:

- Investor Evaluation: Calculate cash truly available to shareholders, giving a clearer picture of profitability beyond net income.

- Valuation Analysis: Use owner earnings to make informed decisions when valuing a company for long-term investments.

- Capex Insight: Get insights into both maintenance and growth capital expenditures (Capex) to assess how much of the company's income is being reinvested.

- Owner Earnings Per Share: Track the value available to each share, helping determine if a stock is a good investment.

This API provides a robust view of a company's profitability and cash flow potential, especially for value investors looking for long-term returns.

Example Use Case
An investor uses the Owner Earnings API to evaluate Apple's true cash earnings before purchasing additional shares, ensuring that the company's income aligns with their long-term investment strategy.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"reportedCurrency": "USD",
		"fiscalYear": "2025",
		"period": "Q1",
		"date": "2024-12-28",
		"averagePPE": 0.13969,
		"maintenanceCapex": -2279964750,
		"ownersEarnings": 27655035250,
		"growthCapex": -660035250,
		"ownersEarningsPerShare": 1.83
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/owner-earnings · 카테고리: statements
