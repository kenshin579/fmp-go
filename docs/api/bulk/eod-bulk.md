# Eod Bulk

The EOD Bulk API allows users to retrieve end-of-day stock price data for multiple symbols in bulk. This API is ideal for financial analysts, traders, and investors who need to assess valuations for a large number of companies.

## Endpoint

`GET https://financialmodelingprep.com/stable/eod-bulk?date=2024-10-22`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| date* | string | 2024-10-22 |

## Description

The EOD Bulk API provides:

- Historical Stock Prices: Access end-of-day stock prices for multiple symbols on a specific date.

- Open, High, Low, Close Prices: Retrieve detailed price data, including opening, high, low, and closing prices for each symbol.

- Volume and Adjusted Close: Get trading volume and adjusted closing prices to analyze stock performance and trading activity.

- Historical Data Analysis: Use historical stock prices to conduct technical analysis, backtesting, and trend forecasting.

This API is designed for users who need to analyze stock prices across a wide range of companies, making it an efficient solution for bulk data retrieval.

## Response (example)

```json
[
	{
		"symbol": "EGS745W1C011.CA",
		"date": "2024-10-22",
		"open": "2.67",
		"low": "2.7",
		"high": "2.9",
		"close": "2.93",
		"adjClose": "2.93",
		"volume": "920904"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/eod-bulk · 카테고리: bulk
