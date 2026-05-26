# Search Crypto News

Search for cryptocurrency news using the FMP Search Crypto News API. Retrieve news related to specific coins or tokens by entering their name or symbol.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/crypto?symbols=BTCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbols* | string | BTCUSD |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The Search Crypto News API allows users to look up cryptocurrency news by entering a coin name or symbol. This API is helpful for:

- Targeted Searches: Quickly find news on specific cryptocurrencies by entering their name or ticker symbol.

- Real-Time & Historical News: Retrieve both current and past news on digital assets to track market trends and price drivers.

- Symbol-Based Lookups: Find news related to your preferred coins, such as Bitcoin (BTC) or Ethereum (ETH).

This API is ideal for cryptocurrency investors who need fast access to news that could affect the value of their digital assets.

Example Use Case
A crypto investor uses the Search Crypto News API to search for news on Ethereum to understand the recent market movements before making a trade.

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

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-crypto-news · 카테고리: news
