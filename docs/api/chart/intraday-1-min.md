# 1 Min Interval Stock Chart

Access precise intraday stock price and volume data with the FMP 1-Minute Interval Stock Chart API. Retrieve real-time or historical stock data in 1-minute intervals, including key information such as open, high, low, and close prices, and trading volume for each minute.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/1min?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |
| nonadjusted | boolean | false |

## Description

The FMP 1-Minute Interval Stock Chart API is designed for traders, analysts, and investors who need detailed intraday stock data for technical analysis, high-frequency trading, or algorithmic strategies. With this API, you can:

- Detailed Intraday Data: Get stock prices at 1-minute intervals, including open, high, low, and close prices, as well as trading volume for each minute.

- Real-Time and Historical Data: Access real-time minute-by-minute data or retrieve historical data using specific date ranges, allowing for long-term analysis.

- Customization with Date Parameters: Easily pull data for any desired time frame, including historical data going back over 30 years, by setting the "from" and "to" parameters.

- Intraday Charting: Perfect for building detailed intraday charts that provide deeper insights into short-term stock movements.

- Perfect for Day Traders: For day traders or algorithmic traders, this API offers the precision needed to identify short-term trends, fluctuations, and trading opportunities.

Example Use Case
A day trader can use the 1-Minute Interval Stock Chart API to track Apple's stock price movements throughout the trading day, enabling them to make timely buy and sell decisions based on real-time price changes and volume spikes.

## Response (example)

```json
[
	{
		"date": "2025-02-04 15:59:00",
		"open": 233.01,
		"low": 232.72,
		"high": 233.13,
		"close": 232.79,
		"volume": 720121
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/intraday-1-min · 카테고리: chart
