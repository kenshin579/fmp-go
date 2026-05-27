# Company Profile by CIK

Retrieve detailed company profile data by CIK (Central Index Key) with the FMP Company Profile by CIK API. This API allows users to search for companies using their unique CIK identifier and access a full range of company data, including stock price, market capitalization, industry, and much more.

## Endpoint

`GET https://financialmodelingprep.com/stable/profile-cik?cik=320193`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| cik* | string | 320193 |

## Description

The FMP Company Profile by CIK API provides comprehensive company information for users who want to look up firms using the CIK code. Ideal for compliance officers, analysts, and investors, this API allows access to vital company details based on their CIK number. Key features include:

- Company Lookup by CIK: Easily find companies using their Central Index Key for fast and accurate identification.

- Stock Price & Market Cap: Get the most up-to-date stock price and market capitalization data for the requested company.

- Comprehensive Financial Data: Access essential financial metrics like beta, dividend yield, and trading range to evaluate a company's performance.

- Global Identifiers: Retrieve key identifiers such as CIK, ISIN, and CUSIP to streamline cross-platform tracking of companies.

- Company Information: Get in-depth details on the company's business operations, CEO, sector, and contact information.

- IPO & Industry Data: View company industry, sector, and IPO details to better understand its market position.

Example Use Case
A compliance officer conducting a regulatory review can use the Company Profile by CIK API to quickly retrieve comprehensive data on Apple Inc. using its unique CIK number, ensuring accuracy in cross-referencing the company across different databases.

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

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/profile-cik · 카테고리: company
