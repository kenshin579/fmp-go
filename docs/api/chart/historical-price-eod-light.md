# Stock Chart Light

Access simplified stock chart data using the FMP Basic Stock Chart API. This API provides essential charting information, including date, price, and trading volume, making it ideal for tracking stock performance with minimal data and creating basic price and volume charts.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/light?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The FMP Basic Stock Chart API delivers streamlined access to stock charting data for users who need to track price movements without overwhelming complexity. This API offers:

- Date & Price Information: Easily track daily price movements for a specific stock symbol.

- Volume Data: Stay informed about trading activity with volume data included for each date.

- Basic Charting Needs: Ideal for generating simple stock price and volume charts for historical performance analysis.

This API is perfect for users and developers who want a quick, straightforward way to visualize stock data without the need for detailed technical indicators.

Example Use Case
A financial app can use the Basic Stock Chart API to display a minimal chart showing a stock's daily closing price and volume, allowing users to quickly assess its performance over time.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2025-02-04",
		"price": 232.8,
		"volume": 44489128
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/historical-price-eod-light · 카테고리: chart
