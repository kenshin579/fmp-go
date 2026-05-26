# SEC Filings Company Search By Symbol

Find company information and regulatory filings using a stock symbol with the FMP SEC Filings Company Search By Symbol API. Quickly access essential company details based on stock ticker symbols.

## Endpoint

`GET https://financialmodelingprep.com/stable/sec-filings-company-search/symbol?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP SEC Filings Company Search By Symbol API allows users to search for a company's SEC filings by simply entering its stock symbol. This API provides valuable information such as:

- Stock Symbol-Based Search: Enter a company's ticker symbol to find official SEC filings and corporate details.

- Detailed Company Information: Retrieve the company's name, CIK number, industry classification (SIC code), and business address.

- Filing Access: Access crucial SEC filings, enabling comprehensive regulatory research and corporate event tracking.

This API is perfect for investors, financial analysts, and compliance professionals who need to quickly pull company-specific SEC filings and information using a stock symbol.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"name": "APPLE INC.",
		"cik": "0000320193",
		"sicCode": "3571",
		"industryTitle": "ELECTRONIC COMPUTERS",
		"businessAddress": "ONE APPLE PARK WAY, CUPERTINO CA 95014",
		"phoneNumber": "(408) 996-1010"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/company-search-by-symbol · 카테고리: secFilings
