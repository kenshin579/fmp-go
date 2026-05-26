# Commodities Quote

Access price quotes for all commodities traded worldwide with the FMP Global Commodities API. Track market movements and identify investment opportunities with comprehensive price data.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote?symbol=GCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | GCUSD |

## Description

The FMP Global Commodities API provides a complete list of price quotes for all commodities traded on exchanges around the world. This API is an essential tool for investors and traders who want to:

- Commodity Prices: Access commodity price quotes, including current prices, highs, lows, and opening prices.

- Track Market Movements: Follow the fluctuations in commodity prices over time to spot trends and make informed decisions.

- Identify Investment Opportunities: Use detailed commodity price data to uncover potential investment opportunities in global markets.

This Commodities API provides a global view of prices, enabling users to stay informed about market conditions and make data-driven investment decisions.

## Response (example)

```json
[
	{
		"symbol": "GCUSD",
		"name": "Gold Futures",
		"price": 3375.3,
		"changePercentage": -0.65635,
		"change": -22.3,
		"volume": 170936,
		"dayLow": 3355.2,
		"dayHigh": 3401.1,
		"yearHigh": 3509.9,
		"yearLow": 2354.6,
		"marketCap": null,
		"priceAvg50": 3358.706,
		"priceAvg200": 3054.501,
		"exchange": "COMMODITY",
		"open": 3398.6,
		"previousClose": 3397.6,
		"timestamp": 1753372205
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/commodities-quote · 카테고리: commodity
