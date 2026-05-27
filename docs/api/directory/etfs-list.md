# ETF Symbol Search

Quickly find ticker symbols and company names for Exchange Traded Funds (ETFs) using the FMP ETF Symbol Search API. This tool simplifies identifying specific ETFs by their name or ticker.

## Endpoint

`GET https://financialmodelingprep.com/stable/etf-list`

## Description

The FMP ETF Symbol Search API allows users to efficiently locate the ticker symbols and names of various Exchange Traded Funds (ETFs). This API is essential for:

- Simple ETF Lookup: Access a database of ETF symbols and company names with minimal effort. By searching with a company name or part of it, users can quickly find relevant ETF symbols.

- Fast, Accurate Data: The API delivers up-to-date information, ensuring users are provided with the latest ETF symbols and names across multiple exchanges.

- Focus on ETFs: The API is designed specifically for ETF-related searches, making it an invaluable resource for investors, traders, and analysts focusing on this market segment.

## Response (example)

```json
[
	{
		"symbol": "GULF",
		"name": "WisdomTree Middle East Dividend Fund"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/etfs-list · 카테고리: directory
