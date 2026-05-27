# Forex Quote

Access real-time forex quotes for currency pairs with the Forex Quote API. Retrieve up-to-date information on exchange rates and price changes to help monitor market movements.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote?symbol=EURUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | EURUSD |

## Description

The Fx Quotes API provides live exchange rate data for various currency pairs, delivering essential insights for traders and financial analysts. Here's how it can help you:

- Live Forex Quotes: Get up-to-the-minute exchange rates and price updates for different forex pairs, such as EUR/USD.

- Detailed Price Information: Access key data, including the current price, day's high and low, year's high and low, and percentage changes.

- Monitor Market Movements: Track the opening and closing prices, as well as 50-day and 200-day moving averages, to gain a comprehensive view of market trends.

This API is essential for forex traders and financial professionals who need accurate and timely currency exchange data to make informed decisions.

Example Use Case
A forex trader uses the Fx Quotes API to monitor the EUR/USD exchange rate throughout the day. By tracking live price changes and percentage movements, the trader can time their trades and react quickly to market fluctuations.

## Response (example)

```json
[
	{
		"symbol": "EURUSD",
		"name": "EUR/USD",
		"price": 1.17598,
		"changePercentage": -0.14754,
		"change": -0.0017376,
		"volume": 184065,
		"dayLow": 1.17371,
		"dayHigh": 1.17911,
		"yearHigh": 1.18303,
		"yearLow": 1.01838,
		"marketCap": null,
		"priceAvg50": 1.15244,
		"priceAvg200": 1.08866,
		"exchange": "FOREX",
		"open": 1.17744,
		"previousClose": 1.17772,
		"timestamp": 1753374603
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/forex-quote · 카테고리: forex
