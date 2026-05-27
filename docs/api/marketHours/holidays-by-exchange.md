# Holidays By Exchange

Retrieve a list of market holidays and non-trading days for a specific stock exchange using the Holidays By Exchange API. Plan your trading schedule by knowing exactly when exchanges like NASDAQ, NYSE, and others are closed.

## Endpoint

`GET https://financialmodelingprep.com/stable/holidays-by-exchange?exchange=NASDAQ`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| exchange* | string | NASDAQ |
| from | date | 2025-04-27 |
| to | date | 2026-04-27 |

## Response (example)

```json
[
	{
		"exchange": "NASDAQ",
		"date": "2026-04-03",
		"name": "Good Friday",
		"isClosed": true,
		"adjOpenTime": null,
		"adjCloseTime": null
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/holidays-by-exchange · 카테고리: marketHours
