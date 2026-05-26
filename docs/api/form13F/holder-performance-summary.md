# Holder Performance Summary

The Holder Performance Summary API provides insights into the performance of institutional investors based on their stock holdings. This data helps track how well institutional holders are performing, their portfolio changes, and how their performance compares to benchmarks like the S&P 500.

## Endpoint

`GET https://financialmodelingprep.com/stable/institutional-ownership/holder-performance-summary?cik=0001067983&page=0`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| cik* | string | 0001067983 |
| page | number | 0 |

## Description

The Holder Performance Summary API allows users to view performance metrics for institutional holders, such as market value changes, portfolio turnover, and relative performance against benchmarks. This API is ideal for:

- Institutional Investor Analysis: Track how well institutional investors are performing based on stock picks, changes in holdings, and market value.

- Portfolio Turnover Analysis: See how frequently an institution buys or sells securities, providing insights into their trading strategy.

- Performance Benchmarking: Compare an institution's performance against the S&P 500 and other benchmarks over different timeframes (1 year, 3 years, 5 years).

This API offers a comprehensive view of an institutional holder's performance over time, helping investors and analysts track key players in the market.

Example Use Case
An investment manager can use the Holder Performance Summary API to analyze Berkshire Hathaway's performance over the last five years and compare it to the S&P 500, assessing how well their investment strategy has fared.

## Response (example)

```json
[
	{
		"date": "2024-09-30",
		"cik": "0001067983",
		"investorName": "BERKSHIRE HATHAWAY INC",
		"portfolioSize": 40,
		"securitiesAdded": 3,
		"securitiesRemoved": 4,
		"marketValue": 266378900503,
		"previousMarketValue": 279969062343,
		"changeInMarketValue": -13590161840,
		"changeInMarketValuePercentage": -4.8542,
		"averageHoldingPeriod": 18,
		"averageHoldingPeriodTop10": 31,
		"averageHoldingPeriodTop20": 27,
		"turnover": 0.175,
		"turnoverAlternateSell": 13.9726,
		"turnoverAlternateBuy": 1.1974,
		"performance": 17707926874,
		"performancePercentage": 6.325,
		"lastPerformance": 38318168662,
		"changeInPerformance": -20610241788,
		"performance1year": 89877376224,
		"performancePercentage1year": 28.5368,
		"performance3year": 91730847239,
		"performancePercentage3year": 31.2597,
		"performance5year": 157058602844,
		"performancePercentage5year": 73.1617,
		"performanceSinceInception": 182067479115,
		"performanceSinceInceptionPercentage": 198.2138,
		"performanceRelativeToSP500Percentage": 6.325,
		"performance1yearRelativeToSP500Percentage": 28.5368,
		"performance3yearRelativeToSP500Percentage": 36.5632,
		"performance5yearRelativeToSP500Percentage": 36.1296,
		"performanceSinceInceptionRelativeToSP500Percentage": 37.0968
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/holder-performance-summary · 카테고리: form13F
