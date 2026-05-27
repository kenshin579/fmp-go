# Aftermarket Quote

Access real-time aftermarket stock quotes with the FMP Aftermarket Quote API. Track bid and ask prices, volume, and other relevant data outside of regular trading hours.

## Endpoint

`GET https://financialmodelingprep.com/stable/aftermarket-quote?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Aftermarket Stock Quote API provides comprehensive quotes for stocks traded outside of normal market hours. This API is essential for:

- Tracking Aftermarket Stock Movers: See real-time bid and ask prices, volumes, and other key metrics after the stock market closes.

- Strategic Analysis: Use aftermarket stock quotes to gain insights into market sentiment and stock performance beyond regular trading hours, helping you make better decisions for the next trading session.

- Efficient Market Monitoring: Stay updated on price movements and trends that can affect next-day trading strategies.

With the Aftermarket Stock Price API, investors can efficiently monitor post-market movements, bid-ask spreads, and trading volumes to stay ahead of potential shifts in the market.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"bidSize": 1,
		"bidPrice": 232.45,
		"askSize": 3,
		"askPrice": 232.64,
		"volume": 41647042,
		"timestamp": 1738715334311
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/aftermarket-quote · 카테고리: quote
