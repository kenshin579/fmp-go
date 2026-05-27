# Latest 8-K SEC Filings

Stay up-to-date with the most recent 8-K filings from publicly traded companies using the FMP Latest 8-K SEC Filings API. Get real-time access to significant company events such as mergers, acquisitions, leadership changes, and other material events that may impact the market.

## Endpoint

`GET https://financialmodelingprep.com/stable/sec-filings-8k?from=2024-01-01&to=2024-03-01&page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from* | string | 2024-01-01 |
| to* | string | 2024-03-01 |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP Latest 8-K SEC Filings API provides timely updates on essential corporate events that are required to be disclosed to the public. These filings offer critical insights for investors and analysts, including:

- Real-Time Filings: Access the latest 8-K filings as they are submitted to the SEC, ensuring you stay informed of key corporate developments.

- Material Events: Track significant corporate events such as mergers, acquisitions, bankruptcies, changes in leadership, and more.

- Direct Filing Links: Get direct access to SEC filing documents, providing you with complete details and disclosures from the companies.

This API is an invaluable tool for investors, analysts, and professionals who need to stay informed of market-moving corporate events.

## Response (example)

```json
[
	{
		"symbol": "BROS",
		"cik": "0001866581",
		"filingDate": "2024-03-01 00:00:00",
		"acceptedDate": "2024-02-29 21:43:41",
		"formType": "8-K",
		"hasFinancials": false,
		"link": "https://www.sec.gov/Archives/edgar/data/1866581/000162828024008098/0001628280-24-008098-index.htm",
		"finalLink": "https://www.sec.gov/Archives/edgar/data/1866581/000162828024008098/exhibit11-8xkfeb2024.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/8k-latest · 카테고리: secFilings
