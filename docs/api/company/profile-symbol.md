# Company Profile Data

Access detailed company profile data with the FMP Company Profile Data API. This API provides key financial and operational information for a specific stock symbol, including the company's market capitalization, stock price, industry, and much more.

## Endpoint

`GET https://financialmodelingprep.com/stable/profile?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Company Profile Data API offers comprehensive insights into a company's financial status and operational details. This API is ideal for analysts, traders, and investors who need an in-depth look at a company's core financial metrics and business information. Key features include:

- Stock Price and Market Cap: Get the latest stock price and market capitalization for the requested symbol.

- Company Details: Access information like company name, description, CEO, and industry classification

- Financial Metrics: Track important financial metrics like dividend yield, stock beta, and trading range to assess performance and volatility.

- Global Identifiers: Retrieve global financial identifiers such as CIK, ISIN, and CUSIP to ensure accurate tracking across platforms.

- Contact Information: Obtain contact details like the company's address, phone number, and website for direct reference.

- IPO Data: Learn about the company's IPO date, sector, and whether it's actively trading.

Example Use Case
An investor researching potential tech investments can use the Company Profile Data API to review the current financial health of Apple Inc., assess its performance, and explore key metrics like its stock range and market cap to inform buying or selling decisions.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"price": 262.82,
		"marketCap": 3900351299800,
		"beta": 1.109,
		"lastDividend": 1.04,
		"range": "169.21-265.29",
		"change": 3.24,
		"changePercentage": 1.24817,
		"volume": 36725325,
		"averageVolume": 47424558,
		"companyName": "Apple Inc.",
		"currency": "USD",
		"cik": "0000320193",
		"isin": "US0378331005",
		"cusip": "037833100",
		"exchangeFullName": "NASDAQ Global Select",
		"exchange": "NASDAQ",
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

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/profile-symbol · 카테고리: company
