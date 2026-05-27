# 5 Min Interval Stock Chart

Access stock price and volume data with the FMP 5-Minute Interval Stock Chart API. Retrieve detailed stock data in 5-minute intervals, including open, high, low, and close prices, along with trading volume for each 5-minute period. This API is perfect for short-term trading analysis and building intraday charts.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/5min?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |
| nonadjusted | boolean | false |

## Description

The FMP 5-Minute Interval Stock Chart API provides users with valuable stock data over 5-minute intervals, allowing for better insight into intraday market activity. It's designed for investors and traders who need quick, accurate data to track short-term price movements. Key features include:

- Short-Term Price Analysis: Track stock price movements over short periods with 5-minute interval data, providing an ideal solution for intraday traders.

- Precise Trading Data: Get open, high, low, and close prices, along with trading volume, for each 5-minute period to identify patterns and trends.

- Intraday Charting: Build detailed intraday charts for any stock symbol, allowing for enhanced visualization of short-term price trends.

- Historical Data Access: Use the API to retrieve historical 5-minute interval data, providing a broader scope for price analysis and trend identification.

Efficient for Active Traders: This API is perfect for day traders and active investors who need fast, reliable data to make informed trading decisions.

Example Use Case
A day trader can use the 5-Minute Interval Stock Chart API to monitor Apple's stock throughout the trading day, identifying short-term trends and making timely trading decisions based on price fluctuations.

## Response (example)

```json
[
	{
		"date": "2025-02-04 15:55:00",
		"open": 232.87,
		"low": 232.72,
		"high": 233.13,
		"close": 232.79,
		"volume": 1555040
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/intraday-5-min · 카테고리: chart
