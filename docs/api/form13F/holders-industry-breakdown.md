# Holders Industry Breakdown

The Holders Industry Breakdown API provides an overview of the sectors and industries that institutional holders are investing in. This API helps analyze how institutional investors distribute their holdings across different industries and track changes in their investment strategies over time.

## Endpoint

`GET https://financialmodelingprep.com/stable/institutional-ownership/holder-industry-breakdown?cik=0001067983&year=2023&quarter=3`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| cik* | string | 0001067983 |
| year* | string | 2023 |
| quarter* | string | 3 |

## Description

The Holders Industry Breakdown API allows users to retrieve data on the industries institutional investors are focusing on, including the weight of their holdings in each sector and how that weight changes over time. This API provides detailed insights into the industry allocation of institutional investors, making it easier to understand their sector focus and strategy.

- Industry Focus Analysis: Understand which industries are receiving the most investment from major institutional holders.

- Portfolio Diversification: Analyze how diversified institutional investors' portfolios are across different sectors.

- Investment Trend Insights: Track changes in the weight of industry holdings to identify shifts in institutional investment strategies.

This API is ideal for investors, analysts, and portfolio managers looking to gain insights into institutional investment behavior across various industries.

Example Use Case
A financial analyst can use the Holders Industry Breakdown API to analyze Berkshire Hathaway's sector focus, identifying whether they are increasing or decreasing their exposure to industries like technology or healthcare over time.

## Response (example)

```json
[
	{
		"date": "2023-09-30",
		"cik": "0001067983",
		"investorName": "BERKSHIRE HATHAWAY INC",
		"industryTitle": "ELECTRONIC COMPUTERS",
		"weight": 49.7704,
		"lastWeight": 51.0035,
		"changeInWeight": -1.2332,
		"changeInWeightPercentage": -2.4178,
		"performance": -20838154294,
		"performancePercentage": -178.2938,
		"lastPerformance": 26615340304,
		"changeInPerformance": -47453494598
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/holders-industry-breakdown · 카테고리: form13F
