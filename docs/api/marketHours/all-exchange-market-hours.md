# All Exchange Market Hours

View the market hours for all exchanges. Check when different markets are active.

## Endpoint

`GET https://financialmodelingprep.com/stable/all-exchange-market-hours`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| timestamp | string | 1769527402 |

## Response (example)

```json
[
	{
		"exchange": "ASX",
		"name": "Australian Securities Exchange",
		"openingHour": "10:00 AM +10:00",
		"closingHour": "04:00 PM +10:00",
		"timezone": "Australia/Sydney",
		"isMarketOpen": true
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/all-exchange-market-hours · 카테고리: marketHours
