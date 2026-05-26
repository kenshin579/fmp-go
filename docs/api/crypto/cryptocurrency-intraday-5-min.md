# 5-Minute Interval Cryptocurrency Data

Analyze short-term price trends with the 5-Minute Interval Cryptocurrency Data API. Access real-time, intraday price data for cryptocurrencies to monitor rapid market movements and optimize trading strategies.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/5min?symbol=BTCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | BTCUSD |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |

## Description

The 5-Minute Interval Cryptocurrency Data API provides detailed intraday data for cryptocurrencies, including:

- Short-Term Price Movements: Track prices in 5-minute intervals, offering granular insights into cryptocurrency performance throughout the trading day.

- Real-Time Market Analysis: Access real-time updates on open, high, low, and close (OHLC) prices, as well as trading volumes, to capture intraday market shifts.

- Support for Technical Analysis: Use 5-minute interval data to perform advanced technical analysis, such as identifying support and resistance levels, spotting short-term trends, or implementing day trading strategies.

This API is essential for active traders, analysts, and investors who need to stay informed of fast-moving price changes and capitalize on short-term market fluctuations.

Example Use Case
A day trader uses the 5-Minute Interval Cryptocurrency Data API to track Bitcoin's price movements throughout the day. By analyzing the short-term price trends, the trader identifies optimal entry and exit points for their trades.

## Response (example)

```json
[
	{
		"date": "2025-07-24 12:25:00",
		"open": 118988.32,
		"low": 118797.03,
		"high": 118997.22,
		"close": 118797.03,
		"volume": 208601161.95555556
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cryptocurrency-intraday-5-min · 카테고리: crypto
