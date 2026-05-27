# ESG Ratings

Access comprehensive ESG ratings for companies and funds with the FMP ESG Ratings API. Make informed investment decisions based on environmental, social, and governance (ESG) performance data.

## Endpoint

`GET https://financialmodelingprep.com/stable/esg-ratings?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP ESG Ratings API provides detailed ESG ratings for companies and funds, helping investors and analysts assess the sustainability and ethical impact of their investments. This API is essential for:

- Evaluating ESG Performance: Access ESG ratings that reflect a company's or fund's performance across environmental, social, and governance criteria, sourced from corporate sustainability reports, ESG research firms, and government agencies.

- Informed Investment Decisions: Use ESG ratings to identify companies and funds that align with your ethical and sustainability goals, ensuring that your investments support positive social and environmental outcomes.

- Filtering Based on ESG Scores: Customize your search to filter for companies with high ESG ratings or low ESG controversy scores, helping you focus on organizations that meet your specific ESG criteria.

This API is a valuable tool for socially conscious investors, financial analysts, and asset managers who prioritize ESG factors in their investment strategies.

Examples Use Cases

- High ESG Performance: An investor interested in companies with strong ESG practices can filter for those with an ESG rating of 80 or higher, ensuring that their investments align with their values.

- Low ESG Controversy: An analyst focused on minimizing environmental risks in their portfolio may filter for companies with low ESG controversy scores, indicating fewer issues related to environmental or social impacts.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"cik": "0000320193",
		"companyName": "Apple Inc.",
		"industry": "CONSUMER ELECTRONICS",
		"fiscalYear": 2024,
		"ESGRiskRating": "B",
		"industryRank": "4 out of 5"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/esg-ratings · 카테고리: ESG
