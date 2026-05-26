# Crypto News

Stay informed with the latest cryptocurrency news using the FMP Crypto News API. Access a curated list of articles from various sources, including headlines, snippets, and publication URLs.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/crypto-latest?page=0&limit=20`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The Crypto News API provides up-to-date news on cryptocurrencies, including key market events and trends. This API is critical for:

- Real-Time Updates: Receive the latest news on major cryptocurrencies like Bitcoin, Ethereum, and more.

- Market Sentiment Analysis: Follow news and reports that could influence crypto market sentiment and price movements.

- Cryptocurrency Trends: Stay informed about industry developments, new technologies, and regulatory updates.

This API is a must-have for anyone involved in the fast-moving world of cryptocurrency investing and trading.

Example Use Case
A crypto trader uses the Crypto News API to track daily news on Bitcoin and Ethereum, enabling them to stay ahead of market trends.

## Response (example)

```json
[
	{
		"symbol": "BTCUSD",
		"publishedDate": "2025-02-03 23:32:19",
		"publisher": "Coingape",
		"title": "Crypto Prices Today Feb 4: BTC & Altcoins Recover Amid Pause On Trump's Tariffs",
		"image": "https://images.financialmodelingprep.com/news/crypto-prices-today-feb-4-btc-altcoins-recover-amid-20250203.webp",
		"site": "coingape.com",
		"text": "Crypto prices today have shown signs of recovery as U.S. President Donald Trump's newly announced import tariffs on Canada and Mexico were paused for 30 days. Bitcoin (BTC) price regained its value, hitting a $102K high amid broader market recovery.",
		"url": "https://coingape.com/crypto-prices-today-feb-4-btc-altcoins-recover-amid-pause-on-trumps-tariffs/"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/crypto-news · 카테고리: news
