# SEC Filings By Form Type

Search for specific SEC filings by form type with the FMP SEC Filings By Form Type API. Retrieve filings such as 10-K, 10-Q, 8-K, and others, filtered by the exact type of document you're looking for.

## Endpoint

`GET https://financialmodelingprep.com/stable/sec-filings-search/form-type?formType=8-K&from=2024-01-01&to=2024-03-01&page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| formType* | string | 8-K |
| from* | string | 2024-01-01 |
| to* | string | 2024-03-01 |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP SEC Filings By Form Type API allows users to filter and retrieve SEC filings based on the document's form type. Whether you're looking for annual reports (10-K), quarterly earnings (10-Q), or event-related filings (8-K), this API provides a streamlined way to access the exact forms needed for analysis or compliance:

- Targeted Filings Search: Search for SEC filings by form type to retrieve specific reports such as 8-K, 10-K, 10-Q, and more.

- Direct Links to Documents: Access the full filing and any associated exhibits directly from the SEC, ensuring complete visibility into company disclosures.

- Regulatory Compliance Monitoring: Use this API to monitor filings related to compliance events, mergers, acquisitions, financial disclosures, and governance updates.

This API is an essential tool for investors, analysts, and regulatory professionals who need quick access to specific types of filings for compliance, analysis, or investment decisions.

## Response (example)

```json
[
	{
		"symbol": "BROS",
		"cik": "0001866581",
		"filingDate": "2024-03-01 00:00:00",
		"acceptedDate": "2024-02-29 21:43:41",
		"formType": "8-K",
		"link": "https://www.sec.gov/Archives/edgar/data/1866581/000162828024008098/0001628280-24-008098-index.htm",
		"finalLink": "https://www.sec.gov/Archives/edgar/data/1866581/000162828024008098/exhibit11-8xkfeb2024.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-by-form-type · 카테고리: secFilings
