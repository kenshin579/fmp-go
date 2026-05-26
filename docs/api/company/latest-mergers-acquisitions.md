# Latest Mergers & Acquisitions

Access real-time data on the latest mergers and acquisitions with the FMP Latest Mergers and Acquisitions API. This API provides key information such as the transaction date, company names, and links to detailed filing information for further analysis.

## Endpoint

`GET https://financialmodelingprep.com/stable/mergers-acquisitions-latest?page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP Latest Mergers and Acquisitions API delivers the most recent information on corporate mergers and acquisitions, giving users access to essential data about company takeovers and transactions. Key features include:

- Transaction Details: Get information on the companies involved, including acquiring and targeted firms.

- Filing Information: Access official filings and documents from the SEC for a deeper analysis of the deal.

- Timely Updates: Stay informed with the most recent mergers and acquisitions data, providing insights into market consolidation.

This API is ideal for analysts, investors, and corporate strategists looking to track corporate activity and make informed decisions based on the latest M&A trends.

Example Use Case
An investment analyst can use the Latest Mergers and Acquisitions API to track recent acquisitions and evaluate the impact of these deals on the companies involved. The data can be used to assess market consolidation, competitive dynamics, and potential investment opportunities.

## Response (example)

```json
[
	{
		"symbol": "ALGT",
		"companyName": "Allegiant Travel CO",
		"cik": "0001362468",
		"targetedCompanyName": "Sun Country Airlines Holdings, Inc.",
		"targetedCik": "0001743907",
		"targetedSymbol": "SNCY",
		"transactionDate": "2026-03-27",
		"acceptedDate": "2026-03-27 17:15:41",
		"link": "https://www.sec.gov/Archives/edgar/data/1362468/000114036126011799/ny20065073x3_s4.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/latest-mergers-acquisitions · 카테고리: company
