# Unadjusted Stock Price

Access stock price and volume data without adjustments for stock splits with the FMP Unadjusted Stock Price Chart API. Get accurate insights into stock performance, including open, high, low, and close prices, along with trading volume, without split-related changes.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/non-split-adjusted?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The FMP Unadjusted Stock Price Chart API provides unadjusted historical price data, allowing traders, analysts, and investors to view stock performance without split-related adjustments. This is useful for users who want a clear view of how stock prices moved before and after stock splits. Key features include:

- Unadjusted Price Data: Access historical stock prices&mdash;open, high, low, and close&mdash;without any adjustments for stock splits.

- Volume Data: Retrieve daily trading volume for further analysis of market activity.

- Pre-Split Analysis: See how stock prices performed in their original form, making it easier to analyze trends prior to a split event.

- Clear Historical View: For investors and analysts looking to avoid the distortions caused by stock splits, this API delivers clear and unmodified data.This API is ideal for anyone who needs accurate, split-free stock data for more precise historical analysis.

Example Use Case
A market researcher analyzing Apple stock performance before and after a split can use the Unadjusted Stock Price Chart API to get a clear view of stock prices without any split-related adjustments.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2025-02-04",
		"adjOpen": 227.2,
		"adjHigh": 233.13,
		"adjLow": 226.65,
		"adjClose": 232.8,
		"volume": 44489128
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/historical-price-eod-non-split-adjusted · 카테고리: chart
