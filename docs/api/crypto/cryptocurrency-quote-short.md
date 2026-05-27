# Cryptocurrency Quote Short

Access real-time cryptocurrency quotes with the FMP Cryptocurrency Quick Quote API. Get a concise overview of current crypto prices, changes, and trading volume for a wide range of digital assets.

## Endpoint

`GET https://financialmodelingprep.com/stable/quote-short?symbol=BTCUSD`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | BTCUSD |

## Description

The FMP Cryptocurrency Quick Quote API provides users with immediate access to essential cryptocurrency price data. It's designed for traders, investors, and analysts who need up-to-the-minute information on the crypto market, including:

- Real-Time Crypto Prices: Retrieve the latest prices for popular cryptocurrencies like Bitcoin, Ethereum, and more.

- Market Changes: View real-time price changes to stay informed of market fluctuations.

- Trading Volume: Access data on trading volumes to assess market activity and liquidity for specific cryptocurrencies.

This API offers a quick and effective way to monitor cryptocurrency prices and make informed decisions based on real-time market data.

Example Use Case
A day trader can use the Cryptocurrency Quick Quote API to track the price of Bitcoin and monitor real-time changes in price and volume, helping them make quick trading decisions in volatile markets.

## Response (example)

```json
[
	{
		"symbol": "BTCUSD",
		"price": 118741.16,
		"change": -37.93,
		"volume": 75302985728
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cryptocurrency-quote-short · 카테고리: crypto
