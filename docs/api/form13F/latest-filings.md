# Institutional Ownership Filings

Stay up to date with the most recent SEC filings related to institutional ownership using the Institutional Ownership Filings API. This tool allows you to track the latest reports and disclosures from institutional investors, giving you a real-time view of major holdings and regulatory submissions.

## Endpoint

`GET https://financialmodelingprep.com/stable/institutional-ownership/latest?page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 100 |

## Description

The Institutional Ownership Filings API gives access to the latest SEC filings from institutional investors, providing insights into reports like Form 13F filings. It's perfect for staying on top of which institutions hold shares in specific companies and monitoring significant ownership changes.
This API is ideal for:

- Tracking Institutional Ownership: Stay updated on which institutions hold shares in specific companies.

- Monitoring Investor Activity: Access filings that show when large investors are buying or selling shares.

- Research & Analysis: Use this data for investment research and trend analysis to see which institutions are bullish or bearish on a company.

- Compliance & Governance: Utilize filings to ensure corporate actions comply with regulatory requirements.

This API ensures real-time access to the most recent institutional filings, keeping you informed about significant investor movements.

Example Use Case
An investment researcher can use the Institutional Ownership Filings API to monitor changes in institutional ownership for companies like Apple, identifying when major hedge funds or pension funds increase or decrease their stakes.

## Response (example)

```json
[
	{
		"cik": "0001963967",
		"name": "CPA ASSET MANAGEMENT LLC",
		"date": "2024-12-31",
		"filingDate": "2025-02-04 00:00:00",
		"acceptedDate": "2025-02-04 17:28:36",
		"formType": "13F-HR",
		"link": "https://www.sec.gov/Archives/edgar/data/1963967/000196396725000001/0001963967-25-000001-index.htm",
		"finalLink": "https://www.sec.gov/Archives/edgar/data/1963967/000196396725000001/boc2024q413f.xml"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/latest-filings · 카테고리: form13F
