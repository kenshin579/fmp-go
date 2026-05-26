# Historical Cryptocurrency Light Chart

Access historical end-of-day prices for a variety of cryptocurrencies with the Historical Cryptocurrency Price Snapshot API. Track trends in price and trading volume over time to better understand market behavior.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/light?symbol=BTCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | BTCUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The Historical Cryptocurrency Price Snapshot API provides crucial insights into the performance of cryptocurrencies over time by offering:

- End-of-Day Prices: Retrieve historical end-of-day prices for cryptocurrencies, allowing you to analyze long-term market trends and patterns.

- Trading Volume Data: Access volume data to evaluate market activity during specific time frames.

- Price Trend Analysis: Use this data to review how a cryptocurrency's value has changed, assisting in making informed investment decisions.

This API is essential for traders, analysts, and investors looking to perform technical analysis or monitor how the market has evolved over time.

Example Use Case
An analyst can use the Historical Cryptocurrency Price Snapshot API to backtest trading strategies by reviewing past price movements and identifying patterns that could influence future price action.

## Response (example)

```json
[
	{
		"symbol": "BTCUSD",
		"date": "2025-07-24",
		"price": 118741.16,
		"volume": 75302985728
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cryptocurrency-historical-price-eod-light · 카테고리: crypto
