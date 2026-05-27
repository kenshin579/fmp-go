# Historical Forex Light Chart

Access historical end-of-day forex prices with the Historical Forex Light Chart API. Track long-term price trends across different currency pairs to enhance your trading and analysis strategies.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/light?symbol=EURUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | EURUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The Historical Forex Light Chart API provides end-of-day forex prices for a wide range of currency pairs. This data is invaluable for traders and analysts looking to:

- Analyze Long-Term Trends: Review historical price data to identify patterns and trends that could influence future market movements.

- Backtest Trading Strategies: Use past data to validate trading strategies by simulating market conditions over extended timeframes.

- Compare Forex Pair Performance: Analyze the performance of different forex pairs over time, helping you make more informed trading decisions.

This API is essential for forex traders, analysts, and investors who need access to accurate historical data for market analysis and strategy development.

Example Use Case
A forex trader uses the Historical Forex Light Chart API to review end-of-day prices for the EUR/USD currency pair over the past five years. By analyzing this data, the trader identifies key support and resistance levels, helping refine their trading strategy.

## Response (example)

```json
[
	{
		"symbol": "EURUSD",
		"date": "2025-07-24",
		"price": 1.17639,
		"volume": 182290
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/forex-historical-price-eod-light · 카테고리: forex
