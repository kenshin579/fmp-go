# Exponential Moving Average

## Endpoint

`GET https://financialmodelingprep.com/stable/technical-indicators/ema?symbol=AAPL&periodLength=10&timeframe=1day`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| periodLength* | number | 10 |
| timeframe* | string | 1min,5min,15min,30min,1hour,4hour,1day |
| from | date | 2026-01-08 |
| to | date | 2026-04-08 |

## Response (example)

```json
[
	{
		"date": "2026-04-08 00:00:00",
		"open": 258.45,
		"high": 259.75,
		"low": 256.53,
		"close": 258.9,
		"volume": 39655304,
		"ema": 254.84409682340092
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/exponential-moving-average · 카테고리: technicalIndicators
