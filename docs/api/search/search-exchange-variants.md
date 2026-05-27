# Exchange Variants

Search across multiple public exchanges to find where a given stock symbol is listed using the FMP Exchange Variants API. This allows users to quickly identify all the exchanges where a security is actively traded.

## Endpoint

`GET https://financialmodelingprep.com/stable/search-exchange-variants?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Exchange Variants API is a powerful tool that provides essential data on where a particular stock is listed across different global exchanges. This API is critical for:

- Multi-Exchange Search: Easily find all public exchanges where a specific stock is listed, ensuring you have a complete understanding of a company's trading activity worldwide.

- Detailed Stock Information: The API returns not only the exchanges where a stock is listed but also includes key financial data such as price, market cap, volume, and beta, allowing for a thorough analysis of the stock.

- Broad Market Coverage: With support for major international exchanges, users can access data from global markets, making it easier to track securities listed in different regions.

This API is a valuable resource for investors, traders, and analysts who need a global view of where securities are traded.

Example: A trader looking for Apple Inc. (AAPL) can use the Exchange Variants API to retrieve a list of exchanges where Apple's stock is traded, along with crucial financial data like market cap, price range, and average trading volume.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"price": 262.82,
		"beta": 1.109,
		"volAvg": 47424558,
		"mktCap": 3900351299800,
		"lastDiv": 1.04,
		"range": "169.21-288.62",
		"changes": 3.24,
		"companyName": "Apple Inc.",
		"currency": "USD",
		"cik": "0000320193",
		"isin": "US0378331005",
		"cusip": "037833100",
		"exchange": "NASDAQ Global Select",
		"exchangeShortName": "NASDAQ",
		"industry": "Consumer Electronics",
		"website": "https://www.apple.com",
		"description": "Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories worldwide. The company offers iPhone, a line of smartphones; Mac, a line of personal computers; iPad, a line of multi-purpose tablets; and wearables, home, and accessories comprising AirPods, Apple TV, Apple Watch, Beats products, and HomePod. It also provides AppleCare support and cloud services; and operates various platforms, including the App Store that allow customers to discov...",
		"ceo": "Timothy D. Cook",
		"sector": "Technology",
		"country": "US",
		"fullTimeEmployees": "164000",
		"phone": "(408) 996-1010",
		"address": "One Apple Park Way",
		"city": "Cupertino",
		"state": "CA",
		"zip": "95014",
		"dcfDiff": 105.92261,
		"dcf": 152.32738976131944,
		"image": "https://images.financialmodelingprep.com/symbol/AAPL.png",
		"ipoDate": "1980-12-12",
		"defaultImage": false,
		"isEtf": false,
		"isActivelyTrading": true,
		"isAdr": false,
		"isFund": false
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-exchange-variants · 카테고리: search
