# 1-Minute Interval Commodities Chart

Track short-term price movements for commodities with the FMP 1-Minute Interval Commodities Chart API. This API provides detailed 1-minute interval data, enabling precise monitoring of intraday market changes.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-chart/1min?symbol=GCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | GCUSD |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |

## Description

The FMP 1-Minute Interval Commodities Chart API delivers minute-by-minute price data for commodities, including open, high, low, and close prices, as well as trading volume. This API is ideal for day traders, analysts, and market participants who require highly granular data to monitor price fluctuations and respond to market trends with speed and accuracy.

- Intraday Data: Access up-to-the-minute price data for commodities, making it easier to track short-term price movements.

- Detailed Price Information: View open, high, low, and close prices, along with trading volume, for precise analysis of market trends.

- Fast Decision-Making: The 1-minute interval data supports fast decision-making for intraday trading, allowing users to act on market opportunities as they arise.

This API is a valuable resource for active traders and investors who need to stay on top of price changes in the fast-moving commodities market.

## Response (example)

```json
[
	{
		"date": "2025-07-24 12:18:00",
		"open": 3374.5,
		"low": 3373.7,
		"high": 3374.5,
		"close": 3374,
		"volume": 123
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/commodities-intraday-1-min · 카테고리: commodity
