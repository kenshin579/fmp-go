# IPOs Disclosure

Access a comprehensive list of disclosure filings for upcoming initial public offerings (IPOs) with the FMP IPO Disclosures API. Stay updated on regulatory filings, including filing dates, effectiveness dates, CIK numbers, and form types, with direct links to official SEC documents.

## Endpoint

`GET https://financialmodelingprep.com/stable/ipos-disclosure`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The FMP IPO Disclosures API provides users with timely and detailed information about regulatory filings for companies planning to go public. This API is essential for analysts, investors, and regulatory professionals who need insights into IPO filing activity. Key features include:

- Filing and Accepted Dates: Track when companies file IPO documents and when those filings are accepted by the SEC.

- Effectiveness Dates: Stay informed on the effectiveness dates, signaling when IPO filings become official.

- Form Types and CIK Numbers: Access key details such as the CIK number and form type (e.g., S-1, CERT) to understand the nature of the filing.

- Direct SEC Links: Get direct access to official SEC documents to review the details of each filing.

This API is a critical tool for those monitoring the regulatory process behind IPOs and understanding the disclosures that accompany companies entering the public market.

Example Use Case
An institutional investor can use the IPO Disclosures API to track regulatory filings for upcoming IPOs and analyze SEC documents before making investment decisions in new market entrants.

## Response (example)

```json
[
	{
		"symbol": "SCHM",
		"filingDate": "2025-02-03",
		"acceptedDate": "2025-02-03",
		"effectivenessDate": "2025-02-03",
		"cik": "0001454889",
		"form": "CERT",
		"url": "https://www.sec.gov/Archives/edgar/data/1454889/000114336225000044/SCCR020325.pdf"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/ipos-disclosure · 카테고리: calendar
