# 1-Minute Interval Index Price

Retrieve 1-minute interval intraday data for stock indexes using the Intraday 1-Minute Price Data API. This API provides granular price information, helping users track short-term price movements and trading volume within each minute.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/1min?symbol=^VIX`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | ^VIX |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |

## Description

The FMP Intraday 1-Minute Price Data API delivers high-frequency price data for stock indexes, offering insights into market fluctuations on a minute-by-minute basis. This level of detail is ideal for active traders and analysts who require real-time market insights for rapid decision-making. Key features include:

- Granular Price Data: Access open, high, low, and close prices for each minute of the trading day.

- Minute-by-Minute Tracking: Monitor short-term price movements and trends in real time.

- Volume Information: Analyze trading volume for each minute, offering insights into market liquidity and activity levels.

- Supports Intraday Trading: Perfect for day traders and high-frequency trading strategies that rely on detailed intraday data.

This API is particularly useful for day traders, quants, and financial analysts who need real-time data to track rapid price movements and make timely trading decisions.

Example Use Case
A day trader specializing in short-term stock index trades uses the Intraday 1-Minute Price Data API to track real-time price changes in the S&P 500 index (^GSPC). With access to minute-by-minute data, they can react to price movements and adjust their trading strategies in real time, optimizing their entry and exit points for maximum profitability.

## Response (example)

```json
[
	{
		"date": "2026-04-08 15:59:00",
		"open": 21.1,
		"low": 21.02,
		"high": 21.1,
		"close": 21.03,
		"volume": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/index-intraday-1-min · 카테고리: indexes
