# Full Forex Quote

Retrieve real-time quotes for multiple forex currency pairs with the FMP Batch Forex Quote API. Get real-time price changes and updates for a variety of forex pairs in a single request.

## Endpoint

`GET https://financialmodelingprep.com/stable/batch-forex-quotes`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| short | boolean | true |

## Description

The FMP Batch Forex Quote API allows users to track real-time exchange rates for multiple currency pairs at once. This API is ideal for those who need to monitor numerous forex pairs simultaneously. Key features include:

- Multiple Currency Pair Tracking: Retrieve real-time quotes for several forex pairs in one request, streamlining market analysis.

- Comprehensive Forex Data: Access up-to-date prices, price changes, and trading volumes across a wide range of global currencies.

- Efficient Market Monitoring: Ideal for traders or analysts monitoring multiple currency pairs in fast-moving forex markets.

The Batch Forex Quote API is a powerful tool for tracking global forex market trends and staying informed on price fluctuations for multiple pairs.

## Response (example)

```json
[
	{
		"symbol": "AEDAUD",
		"price": 0.43575,
		"change": 0.0009547891,
		"volume": 344
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/full-forex-quotes · 카테고리: quote
