# SEC Filings By CIK

Search for SEC filings using the FMP SEC Filings By CIK API. Access detailed regulatory filings by Central Index Key (CIK) number, enabling you to track all filings related to a specific company or entity.

## Endpoint

`GET https://financialmodelingprep.com/stable/sec-filings-search/cik?cik=0000320193&from=2024-01-01&to=2024-03-01&page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| cik* | string | 0000320193 |
| from* | string | 2024-01-01 |
| to* | string | 2024-03-01 |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP SEC Filings By CIK API allows users to retrieve SEC filings by the Central Index Key (CIK) number, providing comprehensive access to a company or entity's official filings. This API is designed for:

- Entity-Specific Filings: Search for SEC filings linked to a specific CIK number, which uniquely identifies publicly traded companies, mutual funds, and other registrants.

- Real-Time Filings: Receive updates on the latest SEC submissions for the entity, including 8-K, 10-K, and 10-Q forms, among others.

- Direct Links to Filings: Access direct links to the official SEC filings and any associated documents or exhibits.

This API is ideal for financial analysts, investors, and compliance officers who require precise and up-to-date filings based on CIK identifiers.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"cik": "0000320193",
		"filingDate": "2024-02-28 00:00:00",
		"acceptedDate": "2024-02-28 17:09:05",
		"formType": "8-K",
		"link": "https://www.sec.gov/Archives/edgar/data/320193/000114036124010155/0001140361-24-010155-index.htm",
		"finalLink": "https://www.sec.gov/Archives/edgar/data/320193/000114036124010155/ny20022580x1_image01.jpg"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-by-cik · 카테고리: secFilings
