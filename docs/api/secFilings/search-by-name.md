# SEC Filings By Name

Search for SEC filings by company or entity name using the FMP SEC Filings By Name API. Quickly retrieve official filings for any organization based on its name.

## Endpoint

`GET https://financialmodelingprep.com/stable/sec-filings-company-search/name?company=Berkshire`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| company* | string | Berkshire |

## Description

The FMP SEC Filings By Name API enables users to search for SEC filings using a company or entity name, providing access to detailed regulatory filings. This API is essential for:

- Entity-Specific Search: Find SEC filings for companies, mutual funds, and other entities by searching their name.

- Comprehensive Filing Access: Get access to key filings such as 8-K, 10-K, 10-Q forms, and more, with the ability to view specific company filings.

- Company Information: Along with SEC filings, receive additional details such as CIK number, business address, and contact information.

This API is ideal for investors, financial analysts, and regulatory compliance officers who need to locate filings based on company or entity names.

## Response (example)

```json
[
	{
		"symbol": "None",
		"name": "BERKSHIRE MULTIFAMILY VALUE FUND II LP",
		"cik": "0001418405",
		"sicCode": "",
		"industryTitle": "",
		"businessAddress": "c/o Berkshire Property Advisors LLC, Boston MA 02108",
		"phoneNumber": "(617) 646-2300"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-by-name · 카테고리: secFilings
