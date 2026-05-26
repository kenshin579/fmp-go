# Crowdfunding Campaign Search

Search for crowdfunding campaigns by company name, campaign name, or platform with the FMP Crowdfunding Campaign Search API. Access detailed information to track and analyze crowdfunding activities.

## Endpoint

`GET https://financialmodelingprep.com/stable/crowdfunding-offerings-search?name=enotap`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| name* | string | enotap |

## Description

The FMP Crowdfunding Campaign Search API allows users to search for crowdfunding campaigns based on company name, campaign name, or platform. This API is a valuable tool for investors and analysts who need to:

- Find Specific Campaigns: Quickly access information on specific crowdfunding campaigns, including the amount raised, number of backers, and investment deadlines.

- Track Company Activity: Monitor the crowdfunding activity of particular companies to identify trends or patterns over time.

- Identify Investment Opportunities: Use crowdfunding data to discover potential investment opportunities based on recent and ongoing campaigns.

This API provides comprehensive details about crowdfunding campaigns, enabling users to make informed decisions based on up-to-date information.

## Response (example)

```json
[
	{
		"cik": "0001912939",
		"name": "Enotap LLC",
		"date": null
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/crowdfunding-search · 카테고리: Fundraisers
