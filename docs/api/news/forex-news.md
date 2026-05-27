# Forex News

Stay updated with the latest forex news articles from various sources using the FMP Forex News API. Access headlines, snippets, and publication URLs for comprehensive market insights.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/forex-latest?page=0&limit=20`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The Forex News API provides up-to-date reports on currency markets, ensuring you stay informed about:

- Currency Market Movements: Get real-time updates on the forex market, including major events and macro-economic trends that influence currency pairs.

- Currency Pair Analysis: Stay informed on specific currency pair movements, such as EUR/USD, GBP/USD, or JPY/CHF, to better understand market conditions.

- Market Sentiment Updates: Follow forex-related news to gauge investor sentiment and market dynamics in the foreign exchange sector.

This API is essential for traders, analysts, and financial professionals who need to stay on top of the ever-changing forex markets.

Example Use Case
A forex trader uses the Forex News API to track the latest news on currency pairs, helping them make quick and informed trading decisions.

## Response (example)

```json
[
	{
		"symbol": "XAUUSD",
		"publishedDate": "2025-02-03 23:55:44",
		"publisher": "FX Street",
		"title": "United Arab Emirates Gold price today: Gold steadies, according to FXStreet data",
		"image": "https://images.financialmodelingprep.com/news/united-arab-emirates-gold-price-today-gold-steadies-according-20250203.jpg",
		"site": "fxstreet.com",
		"text": "Gold prices remained broadly unchanged in United Arab Emirates on Tuesday, according to data compiled by FXStreet.",
		"url": "https://www.fxstreet.com/news/united-arab-emirates-gold-price-today-gold-steadies-according-to-fxstreet-data-202502040455"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/forex-news · 카테고리: news
