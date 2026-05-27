# Index Quote

Access real-time stock index quotes with the Stock Index Quote API. Stay updated with the latest price changes, daily highs and lows, volume, and other key metrics for major stock indices around the world.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote?symbol=^VIX`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | ^VIX |

## Description

The Stock Index Quote API provides real-time data on the performance of stock indices, offering a comprehensive view of current market conditions. This API is essential for:

- Tracking Market Performance: Monitor the real-time movements of key stock indices, like the S&P 500 or NASDAQ, to stay informed about overall market trends.

- Portfolio Management: Use index data to evaluate the health of your investments relative to the broader market.

- Global Market Insights: Access index data across various markets and exchanges, allowing for a global market view.

- Day Trading: Keep track of daily price movements, highs, lows, and volumes for real-time decision-making.

Example Use Case

A trader could use the Stock Index Quote API to track the S&P 500's daily performance in real-time, enabling them to make informed trading decisions based on market trends and volume.

## Response (example)

```json
[
	{
		"symbol": "^VIX",
		"name": "CBOE Volatility Index",
		"price": 16.37,
		"changePercentage": -5.37572,
		"change": -0.93,
		"volume": 0,
		"dayLow": 16.02,
		"dayHigh": 17.22,
		"yearHigh": 60.13,
		"yearLow": 12.7,
		"marketCap": 0,
		"priceAvg50": 16.5992,
		"priceAvg200": 19.3432,
		"exchange": "INDEX",
		"open": 17.02,
		"previousClose": 17.3,
		"timestamp": 1761336901
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/index-quote · 카테고리: indexes
