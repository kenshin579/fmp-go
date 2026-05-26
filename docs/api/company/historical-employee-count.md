# Company Historical Employee Count

Access historical employee count data for a company based on specific reporting periods. The FMP Company Historical Employee Count API provides insights into how a company’s workforce has evolved over time, allowing users to analyze growth trends and operational changes.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-employee-count?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 100 |

## Description

The FMP Company Historical Employee Count API is designed for users who need to track workforce trends for a company across various reporting periods. This data is especially useful for analyzing long-term growth, staffing changes, and the relationship between workforce size and financial performance. Key features include:

- Historical Employee Count: Retrieve workforce size over different periods to analyze growth or decline trends.

- Report Periods: Gain insights into specific timeframes of employee data, tied to annual or quarterly financial reports.

- Filing Date and Form Type: Understand when the employee data was reported, along with the corresponding SEC form type (e.g., 10-K).

- Direct SEC Links: Access the original SEC filings for in-depth research and transparency.

This API is ideal for HR analysts, investors, and business strategists who want to track workforce changes and assess their impact on company operations.

Example Use Case
A financial analyst can use the Company Historical Employee Count API to compare the employee count of Apple Inc. over a five-year period to evaluate how workforce changes correlate with revenue growth and market expansion.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"cik": "0000320193",
		"acceptanceTime": "2025-10-31 06:01:26",
		"periodOfReport": "2025-09-27",
		"companyName": "Apple Inc.",
		"formType": "10-K",
		"filingDate": "2025-10-31",
		"employeeCount": 166000,
		"source": "https://www.sec.gov/Archives/edgar/data/320193/000032019325000079/0000320193-25-000079-index.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/historical-employee-count · 카테고리: company
