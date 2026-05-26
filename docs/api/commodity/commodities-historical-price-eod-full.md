# Full Chart

Access full historical end-of-day price data for commodities with the FMP Comprehensive Commodities Price API. This API enables users to analyze long-term price trends, patterns, and market movements in great detail.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/full?symbol=GCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | GCUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The FMP Comprehensive Commodities Price API provides detailed historical data for various commodities, including opening, high, low, and closing prices, as well as trading volume and price changes. This API is designed for investors, analysts, and traders who need in-depth market insights to evaluate the performance of commodities over time and make data-driven decisions.

- Detailed Historical Data: Access historical end-of-day data, including opening, closing, high, and low prices, trading volume, and price changes.

- Trend Analysis: Analyze long-term price trends and market patterns to better understand the volatility and movement of commodities.

- Comprehensive View: Evaluate not only price movements but also volume and volatility to get a full picture of market conditions.

This API is a powerful tool for professionals looking to assess long-term trends and patterns in commodity markets, helping to predict future price movements or develop investment strategies based on historical data.

## Response (example)

```json
[
	{
		"symbol": "GCUSD",
		"date": "2025-07-24",
		"open": 3398.6,
		"high": 3401.1,
		"low": 3355.2,
		"close": 3373.8,
		"volume": 174758,
		"change": -24.8,
		"changePercent": -0.72971223,
		"vwap": 3376.7
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/commodities-historical-price-eod-full · 카테고리: commodity
