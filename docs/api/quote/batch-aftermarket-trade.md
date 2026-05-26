# Batch Aftermarket Trade

Retrieve real-time aftermarket trading data for multiple stocks with the FMP Batch Aftermarket Trade API. Track post-market trade prices, volumes, and timestamps across several companies simultaneously.

## Endpoint

`GET https://financialmodelingprep.com/stable/batch-aftermarket-trade?symbols=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbols* | string | AAPL |

## Description

The FMP Batch Aftermarket Trade API provides detailed aftermarket trading data for multiple stocks in a single request. This API is perfect for:

- Monitoring Multiple Stocks: Stay updated on post-market trades for various companies, allowing you to track price movements and trading activity after regular market hours.

- Efficient Data Access: Instead of retrieving data for each stock individually, this API lets you access aftermarket trading information for a batch of stocks all at once.

- Enhanced Investment Decisions: Use real-time data from the aftermarket session to analyze trends or patterns across multiple stocks, helping you prepare for the next trading day.

With this API, investors can efficiently track post-market activity for several stocks, enabling more comprehensive analysis and strategy adjustments.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"price": 232.53,
		"tradeSize": 132,
		"timestamp": 1738715334311
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/batch-aftermarket-trade · 카테고리: quote
