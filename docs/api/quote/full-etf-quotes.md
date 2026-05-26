# ETF Price Quotes

Get real-time price quotes for exchange-traded funds (ETFs) with the FMP ETF Price Quotes API. Track current prices, performance changes, and key data for a wide variety of ETFs.

## Endpoint

`GET https://financialmodelingprep.com/stable/batch-etf-quotes`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| short | boolean | true |

## Description

The FMP ETF Price Quotes API allows investors to access real-time pricing information and performance updates for ETFs. This API is essential for those looking to:

- Monitor ETF Performance: Stay updated on the latest prices and performance metrics of different ETFs.

- Evaluate Investment Opportunities: Use real-time price data to assess the value of ETFs and make informed investment decisions.

- Compare ETFs: Easily track and compare the performance of multiple ETFs to optimize your portfolio strategy.

This API provides comprehensive information for investors and analysts looking to make data-driven decisions regarding their ETF investments.

## Response (example)

```json
[
	{
		"symbol": "GULF",
		"price": 16.335,
		"change": 0.13,
		"volume": 3032
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/full-etf-quotes · 카테고리: quote
