# Latest Senate Financial Disclosures

Access the latest financial disclosures from U.S. Senate members with the FMP Latest Senate Financial Disclosures API. Track recent trades, asset ownership, and transaction details for enhanced transparency in government financial activities.

## Endpoint

`GET https://financialmodelingprep.com/stable/senate-latest?page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP Latest Senate Financial Disclosures API provides up-to-date information on trades and asset ownership by U.S. Senate members. With this API, users can:

- Monitor Senate Member Transactions: Access real-time disclosures detailing trades, sales, and purchases made by U.S. Senate members and their families.

- Detailed Transaction Data: Retrieve transaction details, including asset types (stocks, bonds, real estate), transaction dates, amounts, and ownership types.

- Stay Informed: Follow recent disclosures to stay informed about financial activity by key political figures.

This API is essential for those who want to track political figures' financial activities and understand their investment behaviors.

## Response (example)

```json
[
	{
		"symbol": "PEP",
		"disclosureDate": "2026-04-08",
		"transactionDate": "2026-03-30",
		"firstName": "Sheldon",
		"lastName": "Whitehouse",
		"office": "Sheldon Whitehouse",
		"district": "RI",
		"owner": "Spouse",
		"assetDescription": "PepsiCo Inc",
		"assetType": "Stock",
		"type": "Sale",
		"amount": "$1,001 - $15,000",
		"comment": "",
		"link": "https://efdsearch.senate.gov/search/view/ptr/853d0789-28db-4789-9654-a73cff7740d7/"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/senate-latest · 카테고리: senate
