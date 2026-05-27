# house-trading-by-name

## Endpoint

`GET https://financialmodelingprep.com/stable/house-trades-by-name?name=James`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| name* | string | James |

## Response (example)

```json
[
	{
		"symbol": "KD",
		"disclosureDate": "2026-01-26",
		"transactionDate": "2025-12-31",
		"firstName": "James French",
		"lastName": "Hill",
		"office": "James French Hill",
		"district": "AR02",
		"owner": "",
		"assetDescription": "Kyndryl Holdings Inc",
		"assetType": "Stock",
		"type": "Sale",
		"amount": "$1,001 - $15,000",
		"capitalGainsOver200USD": "False",
		"comment": "",
		"link": "https://disclosures-clerk.house.gov/public_disc/ptr-pdfs/2026/20033661.pdf"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/house-trading-by-name · 카테고리: senate
