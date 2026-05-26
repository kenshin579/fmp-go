# Company Name Search

Search for ticker symbols, company names, and exchange details for equity securities and ETFs listed on various exchanges with the FMP Name Search API. This endpoint is useful for retrieving ticker symbols when you know the full or partial company or asset name but not the symbol identifier.

## Endpoint

`GET https://financialmodelingprep.com/stable/search-name?query=AA`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| query* | string | AA |
| limit | number | 50 |
| exchange | string | NASDAQ |

## Description

The FMP Name Search API provides an easy way to find the ticker symbols and exchange information for companies and ETFs. This endpoint is useful for retrieving ticker symbols when you know the company or asset name but not the symbol identifier.

Key Features of the Name Search API

- Simple Company Name Lookup: Enter a company or asset name, and retrieve the corresponding ticker symbol, company name, and exchange details.

- Equity Securities and ETFs: Supports searches for a variety of listed equity securities and exchange-traded funds (ETFs) across major exchanges.

- Accurate and Up-to-Date Data: Receive real-time, accurate search results, ensuring you're always working with the latest available information.

How Investors and Analysts Can Benefit

- Quick Symbol Lookup: Easily locate ticker symbols when you know the company name but not the corresponding symbol.

- Broad Market Coverage: Search across multiple exchanges for both domestic and international companies, helping you stay informed about different markets.

- Streamlined Workflow: Enhance your research and investment decisions by quickly identifying the correct symbols for analysis or trade execution.

## Response (example)

```json
[
	{
		"symbol": "AAGUSD",
		"name": "AAG USD",
		"currency": "USD",
		"exchangeFullName": "CCC",
		"exchange": "CRYPTO"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-name · 카테고리: search
