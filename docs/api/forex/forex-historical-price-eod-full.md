# Historical Forex Full Chart

Access comprehensive historical end-of-day forex price data with the Full Historical Forex Chart API. Gain detailed insights into currency pair movements, including open, high, low, close (OHLC) prices, volume, and percentage changes.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/full?symbol=EURUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | EURUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The Full Historical Forex Chart API provides extensive historical price data for a wide range of currency pairs, offering traders and analysts a deeper understanding of market trends. This data includes open, high, low, and close prices, as well as volume, VWAP (Volume Weighted Average Price), and percentage changes. This API is ideal for:

- Detailed Trend Analysis: Review comprehensive historical price data to analyze long-term trends and patterns in forex markets.

- Advanced Technical Analysis: Use OHLC data to apply technical indicators and identify potential trading signals.

- Strategy Backtesting: Access detailed historical data to validate and optimize trading strategies using real market conditions from past periods.

This API is an essential resource for traders, analysts, and portfolio managers seeking to understand forex market movements and refine their strategies with comprehensive data.

Example Use Case
A portfolio manager uses the Full Historical Forex Chart API to analyze the EUR/USD pair's daily open, high, low, and close prices over the last decade. By reviewing these trends, the manager develops a more informed strategy for managing currency exposure.

## Response (example)

```json
[
	{
		"symbol": "EURUSD",
		"date": "2025-07-24",
		"open": 1.17744,
		"high": 1.17911,
		"low": 1.17371,
		"close": 1.17639,
		"volume": 182290,
		"change": -0.00105,
		"changePercent": -0.08917652,
		"vwap": 1.18
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/forex-historical-price-eod-full · 카테고리: forex
