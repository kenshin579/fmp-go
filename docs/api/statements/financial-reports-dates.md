# Financial Reports Dates

## Endpoint

`GET https://financialmodelingprep.com/stable/financial-reports-dates?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"fiscalYear": 2025,
		"period": "Q1",
		"linkXlsx": "https://financialmodelingprep.com/stable/financial-reports-json?symbol=AAPL&year=2025&period=Q1&apikey=YOUR_API_KEY",
		"linkJson": "https://financialmodelingprep.com/stable/financial-reports-xlsx?symbol=AAPL&year=2025&period=Q1&apikey=YOUR_API_KEY"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/financial-reports-dates · 카테고리: statements
