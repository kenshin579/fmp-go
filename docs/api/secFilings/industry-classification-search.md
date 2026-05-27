# Industry Classification Search

Search and retrieve industry classification details for companies, including SIC codes, industry titles, and business information, with the FMP Industry Classification Search API.

## Endpoint

`GET https://financialmodelingprep.com/stable/industry-classification-search`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol | string | AAPL |
| cik | string | 320193 |
| sicCode | string | 7371 |

## Description

The FMP Industry Classification Search API allows users to search for company information based on their Standard Industrial Classification (SIC) codes. This API provides:

- Company Lookup by Industry: Search for companies by industry classifications, retrieving details such as SIC codes, industry titles, and company contact information.

- Business Information Access: Get comprehensive company information, including business addresses and phone numbers, making it easier to identify and classify businesses by their industry.

- SIC Code Matching: Use this API to match companies with their corresponding industry sectors, enhancing your ability to perform industry-specific research and classification.

This API is valuable for businesses, investors, and researchers who need detailed company information tied to specific industry sectors.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"name": "APPLE INC.",
		"cik": "0000320193",
		"sicCode": "3571",
		"industryTitle": "ELECTRONIC COMPUTERS",
		"businessAddress": "['ONE APPLE PARK WAY', 'CUPERTINO CA 95014']",
		"phoneNumber": "(408) 996-1010"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/industry-classification-search · 카테고리: secFilings
