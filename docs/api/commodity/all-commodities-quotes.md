# All Commodities Quotes

Access quotes for multiple commodities at once with the FMP Batch Commodities Quotes API. Instantly track price changes, volume, and other key metrics for a broad range of commodities.

## Endpoint

`GET https://financialmodelingprep.com/stable/batch-commodity-quotes`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| short | boolean | true |

## Description

The FMP Batch Commodities Quotes API allows users to retrieve live price data for a wide selection of commodities in one request. This API is designed for investors, traders, and analysts who need to monitor several commodities simultaneously and make quick, informed decisions based on market information.

- Batch Quotes: Retrieve quotes for multiple commodities in a single API call, simplifying the process of tracking a wide range of assets.

- Updates: Get up-to-the-minute pricing, ensuring you're always working with the most current market data.

- Market Metrics: Access additional metrics such as price changes and trading volume to provide context to market movements.

This API is essential for professionals who need efficient access to commodity prices without having to query each asset individually.

You can use this API to simultaneously retrieve the latest price for commodities such as DCUSD (current price: $22.29, change: -0.2, volume: 284), allowing for fast analysis and comparison of market data.

## Response (example)

```json
[
	{
		"symbol": "DCUSD",
		"price": 17.18,
		"change": -0.21,
		"volume": 284
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/all-commodities-quotes · 카테고리: commodity
