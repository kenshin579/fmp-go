# Full Cryptocurrency Quotes

Access real-time cryptocurrency quotes with the FMP Full Cryptocurrency Quotes API. Track live prices, trading volumes, and price changes for a wide range of digital assets.

## Endpoint

`GET https://financialmodelingprep.com/stable/batch-crypto-quotes`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| short | boolean | true |

## Description

The FMP Full Cryptocurrency Quotes API offers comprehensive real-time data on cryptocurrency prices, including the latest trading prices, volumes, and price fluctuations. This API is essential for:

- Monitoring Market Prices: Keep track of live cryptocurrency prices to make informed trading decisions.

- Analyzing Market Movements: Stay updated with real-time changes and volume data to identify potential opportunities in the digital asset market.

- Portfolio Management: Use the API to follow the performance of specific cryptocurrencies in your portfolio and adjust your strategy accordingly.

This API is ideal for traders, investors, and analysts who want accurate and up-to-date information about cryptocurrency markets.

## Response (example)

```json
[
	{
		"symbol": "00USD",
		"price": 0.03071157,
		"change": -0.0026034,
		"volume": 169600
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/full-cryptocurrency-quotes · 카테고리: quote
