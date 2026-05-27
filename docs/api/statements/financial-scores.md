# Financial Scores

Assess a company's financial strength using the Financial Health Scores API. This API provides key metrics such as the Altman Z-Score and Piotroski Score, giving users insights into a company’s overall financial health and stability.

## Endpoint

`GET https://financialmodelingprep.com/stable/financial-scores?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The Financial Health Scores API offers a detailed evaluation of a company's financial stability by calculating various scores and metrics. This API is ideal for:

- Bankruptcy Risk Analysis: Use the Altman Z-Score to assess the likelihood of a company facing financial distress.

- Profitability and Efficiency Evaluation: The Piotroski Score helps determine a company's financial strength by measuring profitability and operational efficiency.

- Working Capital Management: Track changes in working capital to understand how a company manages its short-term assets and liabilities.

- Leverage and Capital Structure: Assess the relationship between a company's total liabilities and market capitalization to evaluate its financial leverage.

This API is a powerful tool for investors and analysts who need to evaluate the financial strength of a company to make informed decisions.

Example Use Case
A financial analyst uses the Financial Health Scores API to check Apple's Altman Z-Score and Piotroski Score before recommending it as a stable investment to clients.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"reportedCurrency": "USD",
		"altmanZScore": 9.322985825443649,
		"piotroskiScore": 8,
		"workingCapital": -11125000000,
		"totalAssets": 344085000000,
		"retainedEarnings": -11221000000,
		"ebit": 125675000000,
		"marketCap": 3259495258000,
		"totalLiabilities": 277327000000,
		"revenue": 395760000000
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/financial-scores · 카테고리: statements
