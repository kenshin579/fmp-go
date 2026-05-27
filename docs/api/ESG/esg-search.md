# ESG Investment Search

Align your investments with your values using the FMP ESG Investment Search API. Discover companies and funds based on Environmental, Social, and Governance (ESG) scores, performance, controversies, and business involvement criteria.

## Endpoint

`GET https://financialmodelingprep.com/stable/esg-disclosures?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP ESG Investment Search API is designed to help investors find companies and funds that align with their Environmental, Social, and Governance (ESG) values. This powerful tool allows you to:

- Search by ESG Scores: Identify companies and funds with strong ESG ratings that meet your investment criteria.

- Evaluate Performance: Filter investments based on their ESG performance to ensure they align with your values and financial goals.

- Assess Controversies: Avoid investments in companies involved in significant ESG controversies by filtering based on controversy scores.

- Apply Business Involvement Screens: Screen companies and funds based on specific business activities or sectors that align with your ESG principles.

Examples Use Cases

- An investor focused on sustainability might search for companies with an ESG scores of 80 or higher to ensure strong environmental and social practices.

- An investor concerned about environmental impact could search for companies with low ESG controversy scores to avoid potential risks.

## Response (example)

```json
[
	{
		"date": "2024-12-28",
		"acceptedDate": "2025-01-30",
		"symbol": "AAPL",
		"cik": "0000320193",
		"companyName": "Apple Inc.",
		"formType": "8-K",
		"environmentalScore": 52.52,
		"socialScore": 45.18,
		"governanceScore": 60.74,
		"ESGScore": 52.81,
		"url": "https://www.sec.gov/Archives/edgar/data/320193/000032019325000007/0000320193-25-000007-index.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/esg-search · 카테고리: ESG
