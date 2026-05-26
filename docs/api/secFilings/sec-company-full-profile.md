# SEC Company Full Profile

Retrieve detailed company profiles, including business descriptions, executive details, contact information, and financial data with the FMP SEC Company Full Profile API.

## Endpoint

`GET https://financialmodelingprep.com/stable/sec-profile?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| cik-A | string | 320193 |

## Description

The FMP SEC Company Full Profile API offers comprehensive data on companies registered with the SEC. This API is ideal for:

- Detailed Company Profiles: Access in-depth information on a company's operations, SIC code, CEO, fiscal year, and employee count.

- Executive and Contact Information: Retrieve key executive details and contact information, including business and mailing addresses, phone numbers, and website links.

- Company Description and Operations: Get a detailed company description, including its products, services, markets, and business sectors, allowing for a full understanding of its operations.

- Financial and Regulatory Data: This API provides essential financial data like fiscal year end, IPO date, and links to SEC filings.

This API is crucial for investors, analysts, and researchers who need detailed corporate profiles for financial analysis, competitive research, and investment decision-making.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"cik": "0000320193",
		"registrantName": "Apple Inc.",
		"sicCode": "3571",
		"sicDescription": "Electronic Computers",
		"sicGroup": "Consumer Electronics",
		"isin": "US0378331005",
		"businessAddress": "ONE APPLE PARK WAY,CUPERTINO CA 95014,(408) 996-1010",
		"mailingAddress": "ONE APPLE PARK WAY,CUPERTINO CA 95014",
		"phoneNumber": "(408) 996-1010",
		"postalCode": "95014",
		"city": "Cupertino",
		"state": "CA",
		"country": "US",
		"description": "Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories worldwide. The company offers iPhone, a line of smartphones; Mac, a line of personal computers; iPad, a line of multi-purpose tablets; and wearables, home, and accessories comprising AirPods, Apple TV, Apple Watch, Beats products, and HomePod. It also provides AppleCare support and cloud services; and operates various platforms, including the App Store that allow customers to discov...",
		"ceo": "Mr. Timothy D. Cook",
		"website": "https://www.apple.com",
		"exchange": "NASDAQ",
		"stateLocation": "CA",
		"stateOfIncorporation": "CA",
		"fiscalYearEnd": "09-28",
		"ipoDate": "1980-12-12",
		"employees": "164000",
		"secFilingsUrl": "https://www.sec.gov/cgi-bin/browse-edgar?CIK=0000320193",
		"taxIdentificationNumber": "94-2404110",
		"fiftyTwoWeekRange": "164.08 - 260.1",
		"isActive": true,
		"assetType": "stock",
		"openFigiComposite": "BBG000B9XRY4",
		"priceCurrency": "USD",
		"marketSector": "Technology",
		"securityType": null,
		"isEtf": false,
		"isAdr": false,
		"isFund": false
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/sec-company-full-profile · 카테고리: secFilings
