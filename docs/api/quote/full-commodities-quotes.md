# Full Commodities Quotes

Get up-to-the-minute quotes for commodities with the FMP Commodities Quotes API. Track the latest prices, changes, and volumes for a wide range of commodities, including oil, gold, and agricultural products.

## Endpoint

`GET https://financialmodelingprep.com/stable/batch-commodity-quotes`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| short | boolean | true |

## Description

The FMP Commodities Quotes API provides access to latest pricing information for various commodities. This API is an essential tool for:

- Tracking Key Commodities: Monitor real-time prices for commodities such as oil, gold, natural gas, and agricultural products.

- Making Timely Investment Decisions: Stay informed about price changes and volume to make well-timed trades or investments.

- Market Analysis: Use live data to analyze trends and fluctuations in commodity markets, helping you stay ahead of market movements.

Whether you are a trader, investor, or analyst, this API delivers crucial data to keep you informed on the commodities markets.

## Response (example)

```json
[
	{
		"symbol": "DCUSD",
		"price": 19.89,
		"change": 0.23,
		"volume": 442
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/full-commodities-quotes · 카테고리: quote
