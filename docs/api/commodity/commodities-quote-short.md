# Commodities Quote Short

Get fast and accurate quotes for commodities with the FMP Commodities Quick Quote API. Instantly access the current price, recent changes, and trading volume for various commodities.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote-short?symbol=GCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | GCUSD |

## Description

The FMP Commodities Quick Quote API provides a concise and efficient way to retrieve key information on commodities. Whether you're looking for the latest price, recent market changes, or trading volume, this API delivers the essential data you need for quick analysis and decision-making.

- Instant Price Updates: Receive price data for various commodities, ensuring you're always up to date with the market.

- Market Change Tracking: Stay informed about price changes, allowing for fast reactions to market movements.

- Volume Insights: Access the latest trading volume data to gauge market activity and liquidity.

This API is ideal for investors, traders, and financial analysts who need quick access to essential market data without the complexity of in-depth reports.
Example: For instance, you can use this API to instantly retrieve the current price of gold (symbol: GCUSD), see the price change (-7.2), and track the trading volume (69,930), providing a snapshot of the market's performance at a glance.

## Response (example)

```json
[
	{
		"symbol": "GCUSD",
		"price": 3375.3,
		"change": -22.3,
		"volume": 170936
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/commodities-quote-short · 카테고리: commodity
