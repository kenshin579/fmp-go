# FMP Articles

Access the latest articles from Financial Modeling Prep with the FMP Articles API. Get comprehensive updates including headlines, snippets, and publication URLs.

## Endpoint

`GET https://financialmodelingprep.com/stable/fmp-articles?page=0&limit=20`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 20 |

## Description

The FMP Articles API provides access to a curated list of the most recent articles published by Financial Modeling Prep. This endpoint offers:

- Headlines: Stay informed with the latest headlines covering a wide range of financial topics.

- Snippets: Quickly grasp the key points of each article with concise snippets.

- Publication URLs: Access the full articles through provided URLs for in-depth reading.

This API is updated regularly to ensure you have access to the most current content, helping you stay informed about the latest trends, insights, and analyses from Financial Modeling Prep.

## Response (example)

```json
[
	{
		"title": "Merck Shares Plunge 8% as Weak Guidance Overshadows Strong Revenue Growth",
		"date": "2025-02-04 09:33:00",
		"content": "<p><a href='https://financialmodelingprep.com/financial-summary/MRK'>Merck & Co (NYSE:MRK)</a> saw its stock sink over 8% in pre-market today after delivering mixed fourth-quarter results, with earnings missing expectations, revenue exceeding forecasts, and full-year guidance coming in below analyst estimates.</p>\n<p>For Q4, the pharmaceutical giant reported adjusted earnings per share (EPS) of $1.72, falling short of the $1.81 consensus estimate. However, revenue climbed 7% year-over-year to $1...",
		"tickers": "NYSE:MRK",
		"image": "https://cdn.financialmodelingprep.com/images/fmp-1738679603793.jpg",
		"link": "https://financialmodelingprep.com/market-news/fmp-merck-shares-plunge-8-as-weak-guidance-overshadows-strong-revenue-growth",
		"author": "Davit Kirakosyan",
		"site": "Financial Modeling Prep"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/fmp-articles · 카테고리: news
