# Stock Quote

Access real-time stock quotes with the FMP Stock Quote API. Get up-to-the-minute prices, changes, and volume data for individual stocks.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Stock Quote API provides detailed, real-time stock data for individual stocks, making it a valuable tool for investors, traders, and financial analysts. This API helps you:

- Monitor Real-Time Prices: Stay updated with the latest stock prices to make informed trading decisions.

- Analyze Stock Movements: Track key data points such as price changes, volume, day highs and lows, and yearly highs and lows.

- Portfolio Tracking: Use real-time data to keep track of stock performance in your portfolio.

Whether you are monitoring individual stocks or building trading strategies, this API ensures that you have the most up-to-date information.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"name": "Apple Inc.",
		"price": 232.8,
		"changePercentage": 2.1008,
		"change": 4.79,
		"volume": 44489128,
		"dayLow": 226.65,
		"dayHigh": 233.13,
		"yearHigh": 260.1,
		"yearLow": 164.08,
		"marketCap": 3500823120000,
		"priceAvg50": 240.2278,
		"priceAvg200": 219.98755,
		"exchange": "NASDAQ",
		"open": 227.2,
		"previousClose": 228.01,
		"timestamp": 1738702801
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/quote · 카테고리: quote
