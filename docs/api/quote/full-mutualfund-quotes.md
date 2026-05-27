# Mutual Fund Price Quotes

Access real-time quotes for mutual funds with the FMP Mutual Fund Price Quotes API. Track current prices, performance changes, and key data for various mutual funds.

## Endpoint

`GET https://financialmodelingprep.com/stable/batch-mutualfund-quotes`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| short | boolean | true |

## Description

The FMP Mutual Fund Price Quotes API provides real-time price information and performance updates for mutual funds. Investors and analysts can use this API to:

- Monitor Mutual Fund Performance: Stay updated on the latest price movements and performance changes of mutual funds.

- Track Investment Value: Use price data to assess the value of your mutual fund investments in real-time.

- Analyze Trends: Compare performance across multiple mutual funds to make informed investment decisions and portfolio adjustments.

This API is an essential tool for investors seeking to stay informed on mutual fund prices and performance data.

## Response (example)

```json
[
	{
		"symbol": "ARCFX",
		"price": 9.84,
		"change": 0.01,
		"volume": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/full-mutualfund-quotes · 카테고리: quote
