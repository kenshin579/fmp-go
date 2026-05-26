# All Industry Classification

Access comprehensive industry classification data for companies across all sectors with the FMP All Industry Classification API. Retrieve key details such as SIC codes, industry titles, and business contact information.

## Endpoint

`GET https://financialmodelingprep.com/stable/all-industry-classification`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP All Industry Classification API provides a complete overview of companies classified by industry sector. Users can retrieve:

- Full Industry Classification Data: Access detailed information on companies, including SIC codes, industry titles, and business addresses, for all available industries.

- Comprehensive Company Information: Get relevant details such as company names, CIK numbers, SIC codes, phone numbers, and addresses, helping you identify and analyze businesses across various industries.

- Cross-Industry Analysis: Use this API to study companies within specific industries or across multiple sectors for a complete industry overview.

This API is ideal for investors, analysts, and market researchers looking for extensive industry classification and business data.

## Response (example)

```json
[
	{
		"symbol": "0Q16.L",
		"name": "BANK OF AMERICA CORP /DE/",
		"cik": "0000070858",
		"sicCode": "6021",
		"industryTitle": "NATIONAL COMMERCIAL BANKS",
		"businessAddress": "['BANK OF AMERICA CORPORATE CENTER', 'CHARLOTTE NC 28255']",
		"phoneNumber": "7043868486"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/all-industry-classification · 카테고리: secFilings
