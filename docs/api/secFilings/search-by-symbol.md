# SEC Filings By Symbol

Search and retrieve SEC filings by company symbol using the FMP SEC Filings By Symbol API. Gain direct access to regulatory filings such as 8-K, 10-K, and 10-Q reports for publicly traded companies.

## Endpoint

`GET https://financialmodelingprep.com/stable/sec-filings-search/symbol?symbol=AAPL&from=2024-01-01&to=2024-03-01&page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| from* | string | 2024-01-01 |
| to* | string | 2024-03-01 |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP SEC Filings By Symbol API allows users to search for and retrieve SEC filings based on a specific company's stock symbol. This API provides crucial regulatory documents that are essential for compliance monitoring, financial analysis, and investment research:

- Company-Specific Filings: Access detailed SEC filings for any publicly traded company by simply entering its stock symbol.

- Direct Document Links: Receive direct links to the full SEC filings and related exhibits, ensuring full transparency for your research.

- Real-Time Data Updates: The API provides real-time updates, giving you access to the most recent filings as soon as they are made available by the SEC.

This API is invaluable for investors, analysts, and compliance officers who need to monitor and review regulatory filings tied to a specific company.

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

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-by-symbol · 카테고리: secFilings
