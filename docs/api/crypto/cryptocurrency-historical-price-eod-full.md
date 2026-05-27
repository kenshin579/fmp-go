# Historical Cryptocurrency Full Chart

Access comprehensive end-of-day (EOD) price data for cryptocurrencies with the Full Historical Cryptocurrency Data API. Analyze long-term price trends, market movements, and trading volumes to inform strategic decisions.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/full?symbol=BTCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | BTCUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The Full Historical Cryptocurrency Data API provides extensive historical data, including:

- End-of-Day (EOD) Prices: Retrieve daily open, high, low, close (OHLC) price data for cryptocurrencies.

- Comprehensive Market Data: Access trading volumes, price changes, and VWAP (Volume Weighted Average Price) to gain insights into market behavior.

- Analyze Long-Term Trends: Review historical price data to track long-term trends, volatility, and market cycles, enabling better decision-making for investors and analysts.

This API is essential for long-term investors, analysts, and institutions seeking to evaluate market movements, identify trends, and support strategic planning.

Example Use Case
A long-term cryptocurrency investor could use the Full Historical Cryptocurrency Data API to analyze Bitcoin's market performance over the past year, identifying key resistance levels and potential buying opportunities based on historical price trends.

## Response (example)

```json
[
	{
		"symbol": "BTCUSD",
		"date": "2025-07-24",
		"open": 118779.09,
		"high": 119535.45,
		"low": 117435.22,
		"close": 118741.16,
		"volume": 75302985728,
		"change": -37.93,
		"changePercent": -0.03193323,
		"vwap": 118570.61
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cryptocurrency-historical-price-eod-full · 카테고리: crypto
