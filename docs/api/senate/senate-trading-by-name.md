# senate-trading-by-name

## Endpoint

`GET https://financialmodelingprep.com/stable/senate-trades-by-name?name=Jerry`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| name* | string | Jerry |

## Response (example)

```json
[
	{
		"symbol": "GOOG",
		"disclosureDate": "2025-10-27",
		"transactionDate": "2025-09-23",
		"firstName": "Jerry",
		"lastName": "Moran",
		"office": "Jerry Moran",
		"district": "KS",
		"owner": "Spouse",
		"assetDescription": "Alphabet Cl C",
		"assetType": "Stock",
		"type": "Sale (Partial)",
		"amount": "$1,001 - $15,000",
		"capitalGainsOver200USD": "False",
		"comment": "--",
		"link": "https://efdsearch.senate.gov/search/view/ptr/b83b6502-520b-4403-9777-60f6c2d93bc1/"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/senate-trading-by-name · 카테고리: senate
