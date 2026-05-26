# Search Mergers & Acquisitions

Search for specific mergers and acquisitions data with the FMP Search Mergers and Acquisitions API. Retrieve detailed information on M&A activity, including acquiring and targeted companies, transaction dates, and links to official SEC filings.

## Endpoint

`GET https://financialmodelingprep.com/stable/mergers-acquisitions-search?name=Apple`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| name* | string | Apple |

## Description

The FMP Search Mergers and Acquisitions API allows users to find mergers and acquisitions by company name, enabling a deeper understanding of corporate activity. This API is useful for those needing detailed data on past and ongoing deals, including:

- Company-Specific M&A Data: Search for M&A transactions involving specific companies, either as the acquirer or target.

- Transaction Dates: Access the exact dates of the transactions for precise tracking.

- Filing Links: Obtain links to official SEC documents for detailed information on the terms and conditions of the deal.

This API is perfect for financial analysts, researchers, and corporate strategists who need comprehensive M&A data to inform business or investment decisions.

Example Use Case
A corporate strategist can use the Search Mergers and Acquisitions API to identify past acquisition targets of a competitor. This information can help shape competitive strategies or identify industry trends that may affect future business opportunities.

## Response (example)

```json
[
	{
		"symbol": "PEGY",
		"companyName": "Pineapple Energy Inc.",
		"cik": "0000022701",
		"targetedCompanyName": "Communications Systems, Inc.",
		"targetedCik": "0000022701",
		"targetedSymbol": "JCS",
		"transactionDate": "2021-11-12",
		"acceptedDate": "2021-11-12 09:54:22",
		"link": "https://www.sec.gov/Archives/edgar/data/22701/000089710121000932/a211292_s-4.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-mergers-acquisitions · 카테고리: company
