# Filings Extract With Analytics By Holder

The Filings Extract With Analytics By Holder API provides an analytical breakdown of institutional filings. This API offers insight into stock movements, strategies, and portfolio changes by major institutional holders, helping you understand their investment behavior and track significant changes in stock ownership.

## Endpoint

`GET https://financialmodelingprep.com/stable/institutional-ownership/extract-analytics/holder?symbol=AAPL&year=2023&quarter=3&page=0&limit=10`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| year* | string | 2023 |
| quarter* | string | 3 |
| page | number | 0 |
| limit | number | 10 |

## Description

The Filings Extract With Analytics By Holder API allows users to extract detailed analytics from filings by institutional investors. It offers information such as shares held, changes in stock weight and market value, ownership percentages, and other important metrics that provide an analytical view of institutional investment strategies.

- Institutional Investor Analysis: Track the behavior of large institutional holders such as Vanguard, including their changes in stock positions and market value.

- Portfolio Movement Monitoring: Analyze stock movements and holding period data to see how long institutions have held a stock and when they increased or reduced their positions.

- Investment Strategy Insights: Understand investment strategies by looking at changes in weight, market value, and ownership over time.

This API offers granular insights into how institutions manage their portfolios, providing data to investors and analysts for deeper investment analysis.

Example Use Case
An investment analyst can use the Filings Extract With Analytics By Holder API to monitor Vanguard Group's activity in Apple Inc. stocks, seeing how much stock Vanguard holds, any changes in weight or market value, and when the stock was first added to their portfolio.

## Response (example)

```json
[
	{
		"date": "2023-09-30",
		"cik": "0000102909",
		"filingDate": "2023-12-18",
		"investorName": "VANGUARD GROUP INC",
		"symbol": "AAPL",
		"securityName": "APPLE INC",
		"typeOfSecurity": "COM",
		"securityCusip": "037833100",
		"sharesType": "SH",
		"putCallShare": "Share",
		"investmentDiscretion": "SOLE",
		"industryTitle": "ELECTRONIC COMPUTERS",
		"weight": 5.4673,
		"lastWeight": 5.996,
		"changeInWeight": -0.5287,
		"changeInWeightPercentage": -8.8175,
		"marketValue": 222572509140,
		"lastMarketValue": 252876459509,
		"changeInMarketValue": -30303950369,
		"changeInMarketValuePercentage": -11.9837,
		"sharesNumber": 1299997133,
		"lastSharesNumber": 1303688506,
		"changeInSharesNumber": -3691373,
		"changeInSharesNumberPercentage": -0.2831,
		"quarterEndPrice": 171.21,
		"avgPricePaid": 95.86,
		"isNew": false,
		"isSoldOut": false,
		"ownership": 8.3336,
		"lastOwnership": 8.305,
		"changeInOwnership": 0.0286,
		"changeInOwnershipPercentage": 0.3445,
		"holdingPeriod": 42,
		"firstAdded": "2013-06-30",
		"performance": -29671950396,
		"performancePercentage": -11.7338,
		"lastPerformance": 38078179274,
		"changeInPerformance": -67750129670,
		"isCountedForPerformance": true
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/filings-extract-with-analytics-by-holder · 카테고리: form13F
