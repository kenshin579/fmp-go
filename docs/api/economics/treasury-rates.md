# Treasury Rates

Access latest and historical Treasury rates for all maturities with the FMP Treasury Rates API. Track key benchmarks for interest rates across the economy.

## Endpoint

`GET https://financialmodelingprep.com/stable/treasury-rates`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The Treasury Rates API provides latest and historical data on Treasury rates for all maturities. These rates represent the interest rates that the US government pays on its debt obligations and serve as a critical benchmark for interest rates across the economy. Investors can use this API to:

- Track Treasury Rates Over Time: Monitor the movement of Treasury rates and understand how they change over different periods.

- Identify Interest Rate Trends: Analyze trends in interest rates to gain insights into the broader economic landscape.

- Make Informed Investment Decisions: Use the data to inform investment strategies based on current and historical interest rate information.

This API is an invaluable tool for investors, analysts, and economists who need accurate and timely information on Treasury rates.

## Response (example)

```json
[
	{
		"date": "2026-04-08",
		"month1": 3.67,
		"month2": 3.71,
		"month3": 3.69,
		"month6": 3.73,
		"year1": 3.69,
		"year2": 3.79,
		"year3": 3.78,
		"year5": 3.92,
		"year7": 4.1,
		"year10": 4.29,
		"year20": 4.87,
		"year30": 4.89
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/treasury-rates · 카테고리: economics
