# Equity Offering Search

Easily search for equity offerings by company name or stock symbol with the FMP Equity Offering Search API. Access detailed information about recent share issuances to stay informed on company fundraising activities.

## Endpoint

`GET https://financialmodelingprep.com/stable/fundraising-search?name=NJOY`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| name* | string | NJOY |

## Description

The FMP Equity Offering Search API allows users to quickly find relevant equity offering data, including details on recent share issuances and filing dates. This API is essential for investors, analysts, and compliance officers who want to:

- Track Company Equity Offerings: Search by company name or ticker symbol to find recent equity offerings.

- Analyze Issuance Data: Access key information such as offering dates, company names, and CIK (Central Index Key) numbers to get a comprehensive view of recent share issuances.

- Stay Informed About Market Activity: Use the API to monitor fundraising activities, assess the impact of equity offerings, and make informed investment decisions.

This API provides an efficient way to stay on top of market events by offering a quick search for new equity issuances from companies across various sectors.

Example Use Case
An investor can use the Equity Offering Search API to identify which companies are issuing new shares, allowing them to assess the impact of equity offerings on their portfolio or potential investments.

## Response (example)

```json
[
	{
		"cik": "0001547416",
		"name": "NJOY INC",
		"date": "2014-02-28 16:00:25"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/equity-offering-search · 카테고리: Fundraisers
