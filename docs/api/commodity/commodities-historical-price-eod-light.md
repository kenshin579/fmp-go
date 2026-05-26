# Light Chart

Access historical end-of-day prices for various commodities with the FMP Historical Commodities Price API. Analyze past price movements, trading volume, and trends to support informed decision-making.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/light?symbol=GCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | GCUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The FMP Historical Commodities Price API offers users access to end-of-day pricing data for a wide range of commodities. This API is designed for investors, traders, and analysts who need to perform historical analysis on commodities markets, track price trends, and make informed predictions based on past data.

- End-of-Day Pricing: Retrieve accurate historical prices for commodities, including key metrics like trading volume, to analyze market performance over time.

- Comprehensive Historical Data: Access a detailed record of price changes for commodities over any chosen period.

- Trading Volume Insights: Evaluate the trading activity for each commodity with volume data included alongside price information.

This API is ideal for financial professionals looking to analyze historical commodity data for research, risk management, or strategic trading purposes.

## Response (example)

```json
[
	{
		"symbol": "GCUSD",
		"date": "2025-07-24",
		"price": 3373.8,
		"volume": 174758
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/commodities-historical-price-eod-light · 카테고리: commodity
