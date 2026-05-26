# Full Cryptocurrency Quote

Access real-time quotes for all cryptocurrencies with the FMP Full Cryptocurrency Quote API. Obtain comprehensive price data including current, high, low, and open prices.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote?symbol=BTCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | BTCUSD |

## Description

The Full Cryptocurrency Quote API provides real-time quotes for all cryptocurrencies traded on exchanges worldwide. This endpoint offers detailed information such as:

Current Price: Get the latest price of any cryptocurrency.

High, Low, and Open Prices: Access the highest, lowest, and opening prices for the day.

Investors can use the Full Cryptocurrency Quote API to:

- Monitor Real-Time Prices: Stay updated with real-time prices of all cryptocurrencies traded globally.

- Track Price Movements: Follow the movement of cryptocurrency prices over time to identify trends and patterns.

- Identify Investment Opportunities: Use comprehensive price data to spot potential investment opportunities.

- Make Informed Trading Decisions: Base your trading decisions on up-to-date and accurate cryptocurrency price data.

## Response (example)

```json
[
	{
		"symbol": "BTCUSD",
		"name": "Bitcoin USD",
		"price": 118741.16,
		"changePercentage": -0.03193323,
		"change": -37.93,
		"volume": 75302985728,
		"dayLow": 117435.22,
		"dayHigh": 119535.45,
		"yearHigh": 123091.61,
		"yearLow": 49121.24,
		"marketCap": 2344693699320,
		"priceAvg50": 109824.32,
		"priceAvg200": 98161.086,
		"exchange": "CRYPTO",
		"open": 118779.09,
		"previousClose": 118779.09,
		"timestamp": 1753374602
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cryptocurrency-quote · 카테고리: crypto
