# Index Short Quote

Access concise stock index quotes with the Stock Index Short Quote API. This API provides a snapshot of the current price, change, and volume for stock indexes, making it ideal for users who need a quick overview of market movements.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote-short?symbol=^VIX`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | ^VIX |

## Description

The Stock Index Short Quote API delivers simplified, real-time index data, offering essential metrics such as price, change, and volume. This API is a valuable tool for traders, investors, and analysts who need a quick overview of an index's current standing without unnecessary details. Key features include:

- Real-Time Index Data: Get current price, change, and volume for stock indexes.

- Simplified Data: Designed for users who need only the essential figures, providing a clear and efficient market snapshot.

- Wide Market Coverage: Retrieve short quotes for a wide range of global indexes.

This API is perfect for traders and analysts who want to stay updated on index performance at a glance, enabling them to react quickly to market shifts.

Example Use Case
A trader monitoring the S&P 500 throughout the trading day can use the Stock Index Short Quote API to quickly access real-time price changes, helping them make decisions on whether to buy or sell without delving into more complex data.

## Response (example)

```json
[
	{
		"symbol": "^VIX",
		"price": 16.37,
		"change": -0.93,
		"volume": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/index-quote-short · 카테고리: indexes
