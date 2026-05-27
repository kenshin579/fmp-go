# 5-Minute Interval Index Price

Retrieve 5-minute interval intraday price data for stock indexes using the Intraday 5-Minute Price Data API. This API provides crucial insights into price movements and trading volume within 5-minute windows, ideal for traders who require short-term data.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/5min?symbol=^VIX`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | ^VIX |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |

## Description

The FMP Intraday 5-Minute Price Data API offers real-time price and volume data for stock indexes, updated every 5 minutes during active market hours. This API is designed for traders and analysts who need detailed, short-term data to track price fluctuations and make timely decisions. Key features include:

- 5-Minute Interval Data: Access open, high, low, and close prices for each 5-minute interval throughout the trading day.

- Real-Time Tracking: Stay up-to-date with price changes and market trends in near real-time.

- Volume Data: Analyze trading volume in 5-minute intervals to gauge market activity and liquidity.

- Supports Short-Term Trading: Ideal for short-term and swing traders looking for frequent updates to inform their strategies.

This API is perfect for day traders, quants, and financial professionals who need to monitor price movements closely and execute trades based on short-term fluctuations.

Example Use Case
A swing trader monitoring the S&P 500 index (^GSPC) uses the Intraday 5-Minute Price Data API to track price movements over the course of the trading day. By analyzing the 5-minute intervals, they can time their trades more accurately, reacting quickly to short-term market changes and optimizing their strategy for maximum return.

## Response (example)

```json
[
	{
		"date": "2026-04-08 15:55:00",
		"open": 21.26,
		"low": 21.02,
		"high": 21.3,
		"close": 21.03,
		"volume": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/index-intraday-5-min · 카테고리: indexes
