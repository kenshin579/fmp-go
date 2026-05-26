# Latest Financial Statements

## Endpoint

`GET https://financialmodelingprep.com/stable/latest-financial-statements?page=0&limit=250`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 250 |

## Response (example)

```json
[
	{
		"symbol": "FGFI",
		"calendarYear": 2024,
		"period": "Q4",
		"date": "2024-12-31",
		"dateAdded": "2025-03-13 17:03:59"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/latest-financial-statements · 카테고리: statements
