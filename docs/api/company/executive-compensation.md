# Executive Compensation

Retrieve comprehensive compensation data for company executives with the FMP Executive Compensation API. This API provides detailed information on salaries, stock awards, total compensation, and other relevant financial data, including filing details and links to official documents.

## Endpoint

`GET https://financialmodelingprep.com/stable/governance-executive-compensation?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Executive Compensation API is designed to give investors, analysts, and researchers a complete overview of executive compensation for publicly traded companies. This API is beneficial for:

- Executive Salary & Benefits: Retrieve data on annual salaries, stock awards, bonuses, and incentive plans.

- Comprehensive Compensation Breakdown: Access detailed reports on total compensation, including base pay and additional awards or incentives.

- Filing Information: Includes key filing dates and direct links to SEC filings for deeper analysis of compensation packages.

This API provides valuable insights into how company executives are compensated, helping users understand leadership incentives and assess company governance.

Example Use Case
A compensation analyst can use the Executive Compensation API to compare CEO pay across different companies, analyzing how various forms of compensation&mdash;such as salary, stock awards, and performance incentives&mdash;impact executive behavior and company performance.

## Response (example)

```json
[
	{
		"cik": "0000320193",
		"symbol": "AAPL",
		"companyName": "Apple Inc.",
		"filingDate": "2026-01-08",
		"acceptedDate": "2026-01-08 16:31:36",
		"nameAndPosition": "Tim Cook Chief Executive Officer",
		"year": 2025,
		"salary": 3000000,
		"bonus": 0,
		"stockAward": 57535293,
		"optionAward": 0,
		"incentivePlanCompensation": 12000000,
		"allOtherCompensation": 1759518,
		"total": 74294811,
		"link": "https://www.sec.gov/Archives/edgar/data/320193/000130817926000008/0001308179-26-000008-index.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/executive-compensation · 카테고리: company
