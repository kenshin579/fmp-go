# 30 Min Interval Stock Chart

Access stock price and volume data with the FMP 30-Minute Interval Stock Chart API. Retrieve essential stock data in 30-minute intervals, including open, high, low, close prices, and trading volume. This API is perfect for creating intraday charts and tracking medium-term price movements for more strategic trading decisions.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/30min?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |
| nonadjusted | boolean | false |

## Description

The FMP 30-Minute Interval Stock Chart API is designed for traders and investors seeking medium-term price insights without monitoring every minute of the trading day. By delivering key stock metrics in 30-minute intervals, it offers a well-balanced view of stock performance over time. Key features include:

- Efficient Medium-Term Analysis: Monitor stock price fluctuations at 30-minute intervals, providing a clear view of price movements without the noise of smaller time frames.

- Detailed Price Metrics: Access important data points such as open, high, low, close prices, and trading volume to build comprehensive intraday charts.

- Ideal for Intraday Strategies: This API supports trading strategies that rely on medium-term price movements and volume patterns, making it ideal for day traders and investors.

- Historical Data Availability: Retrieve historical data for 30-minute intervals, helping you analyze trends and patterns from past trading sessions.

- Optimized for Trend Tracking: With data available at 30-minute intervals, this API offers an efficient solution for those looking to identify key trends during the trading day.

Example Use Case
A day trader uses the 30-Minute Interval Stock Chart API to monitor the performance of Apple stock over the course of a trading day, identifying important price patterns and volume changes to make calculated buy and sell decisions.

## Response (example)

```json
[
	{
		"date": "2025-02-04 15:30:00",
		"open": 232.29,
		"low": 232.01,
		"high": 233.13,
		"close": 232.79,
		"volume": 3476320
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/intraday-30-min · 카테고리: chart
