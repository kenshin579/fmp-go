# Search Forex News

Search for foreign exchange news using the FMP Search Forex News API. Find targeted news on specific currency pairs by entering their symbols for focused updates.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/forex?symbols=EURUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbols* | string | EURUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The Search Forex News API allows users to look up forex news by entering a currency pair, such as EUR/USD or GBP/USD. This API is perfect for:

- Targeted News Search: Easily find news about specific currency pairs to track the latest developments in the forex market.

- Historical News Access: Look up both current and historical forex news to analyze long-term trends and market movements.

- Symbol-Based Retrieval: Enter specific currency pair symbols to retrieve relevant news for informed decision-making.

This API is ideal for forex traders who need quick access to news related to specific currency pairs.

Example Use Case
A currency trader uses the Search Forex News API to search for the latest news on EUR/USD, helping them understand recent price fluctuations before entering a trade.

## Response (example)

```json
[
	{
		"symbol": "EURUSD",
		"publishedDate": "2025-02-03 18:43:01",
		"publisher": "FX Street",
		"title": "EUR/USD trims losses but still sheds weight",
		"image": "https://images.financialmodelingprep.com/news/eurusd-trims-losses-but-still-sheds-weight-20250203.jpg",
		"site": "fxstreet.com",
		"text": "EUR/USD dropped sharply following fresh tariff threats from US President Donald Trump, impacting the markets. However, significant declines in global risk markets eased as the Trump administration offered 30-day concessions on impending tariffs for Canada and Mexico.",
		"url": "https://www.fxstreet.com/news/eur-usd-trims-losses-but-still-sheds-weight-202502032343"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-forex-news · 카테고리: news
