# 15 Min Interval Stock Chart

Access stock price and volume data with the FMP 15-Minute Interval Stock Chart API. Retrieve detailed stock data in 15-minute intervals, including open, high, low, close prices, and trading volume. This API is ideal for creating intraday charts and analyzing medium-term price trends during the trading day.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/15min?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |
| nonadjusted | boolean | false |

## Description

The FMP 15-Minute Interval Stock Chart API is designed to provide a more balanced view of stock price movements throughout the trading day. By delivering key data at 15-minute intervals, this API offers medium-term insights for traders and investors who need to monitor stock trends in a concise but effective format. Key features include:

- Medium-Term Price Analysis: Monitor price fluctuations over 15-minute intervals, ideal for traders who need to identify intraday trends without analyzing every minute.

- Comprehensive Data Points: Access key metrics such as open, high, low, close prices, and trading volume to create detailed intraday charts.

- Flexible Intraday Monitoring: This API is suitable for traders and investors who need to track stock performance throughout the trading day, making it easier to spot price movements and trends.

- Historical Data Access: Retrieve historical 15-minute interval data to conduct in-depth analysis of past trading sessions and identify recurring patterns.

- Efficient Data Retrieval: Ideal for those who want a balance between fast-moving data (such as 1-minute intervals) and longer-term intraday data for smarter decision-making.

Example Use Case
A swing trader can use the 15-Minute Interval Stock Chart API to monitor Apple stock throughout the trading day, analyzing medium-term price movements to make strategic trade entries and exits based on significant fluctuations.

## Response (example)

```json
[
	{
		"date": "2025-02-04 15:45:00",
		"open": 232.25,
		"low": 232.18,
		"high": 233.13,
		"close": 232.79,
		"volume": 2535629
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/intraday-15-min · 카테고리: chart
