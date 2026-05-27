# U.S. House Trades

Track the financial trades made by U.S. House members and their families with the FMP U.S. House Trades API. Access real-time information on stock sales, purchases, and other investment activities to gain insight into their financial decisions.

## Endpoint

`GET https://financialmodelingprep.com/stable/house-trades?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP U.S. House Trades API provides a comprehensive view of the trading activities of U.S. House members and their spouses. This API offers detailed data on trades, including stock sales and purchases, ownership details, and transaction amounts. Users can:

- Monitor Trading Activity: Stay informed about the latest stock trades made by U.S. House members and their families.

- Understand Financial Moves: Gain insights into the financial decisions of government officials through detailed trade data.

- Transparency and Accountability: Use the data to follow the financial actions of U.S. House members, ensuring greater transparency in government.

This API is ideal for political analysts, journalists, and the general public interested in understanding the financial moves of U.S. House representatives.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"disclosureDate": "2026-04-08",
		"transactionDate": "2025-11-13",
		"firstName": "Ed",
		"lastName": "Case",
		"office": "Ed Case",
		"district": "HI01",
		"owner": "",
		"assetDescription": "Apple Inc",
		"assetType": "Stock",
		"type": "Purchase",
		"amount": "$1,001 - $15,000",
		"capitalGainsOver200USD": "False",
		"comment": "",
		"link": "https://disclosures-clerk.house.gov/public_disc/ptr-pdfs/2026/20034221.pdf"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/house-trading · 카테고리: senate
