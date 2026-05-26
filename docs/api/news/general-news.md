# General News

Access the latest general news articles from a variety of sources with the FMP General News API. Obtain headlines, snippets, and publication URLs for comprehensive news coverage.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/general-latest?page=0&limit=20`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The FMP General News API provides access to the latest general news articles from a wide range of sources. This endpoint includes:

- Headlines: Stay informed with the latest headlines on current events.

- Snippets: Get brief summaries of the articles to quickly understand the key points.

- Publication URLs: Access full articles through provided URLs for detailed information.

This API is updated daily to ensure you have the most current news. Simply provide the date range you are interested in, and the endpoint will return a list of all general news articles published during that period.

## Response (example)

```json
[
	{
		"symbol": null,
		"publishedDate": "2025-02-03 23:51:37",
		"publisher": "CNBC",
		"title": "Asia tech stocks rise after Trump pauses tariffs on China and Mexico",
		"image": "https://images.financialmodelingprep.com/news/asia-tech-stocks-rise-after-trump-pauses-tariffs-on-20250203.jpg",
		"site": "cnbc.com",
		"text": "Gains in Asian tech companies were broad-based, with stocks in Japan, South Korea and Hong Kong advancing. Semiconductor players Advantest and Lasertec led gains among Japanese tech stocks.",
		"url": "https://www.cnbc.com/2025/02/04/asia-tech-stocks-rise-after-trump-pauses-tariffs-on-china-and-mexico.html"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/general-news · 카테고리: news
