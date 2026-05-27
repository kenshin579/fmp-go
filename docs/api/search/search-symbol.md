# Stock Symbol Search

Easily find the ticker symbol of any stock with the FMP Stock Symbol Search API. Search by symbol across multiple global markets.

## Endpoint

`GET https://financialmodelingprep.com/stable/search-symbol?query=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| query* | string | AAPL |
| limit | number | 50 |
| exchange | string | NASDAQ |

## Description

The FMP Stock Symbol Search API allows users to quickly and efficiently locate stock ticker symbols. Whether you're searching for U.S. stocks, international equities, or ETFs, this API provides fast, reliable results. Key features include:

- Simple Search: Enter a company name or ticker symbol to retrieve essential details like the symbol, company name, exchange, and currency.

- Global Market Access: Search across major stock exchanges, including NASDAQ, NYSE, and more.

- Accurate and Up-to-Date: The API delivers real-time results, ensuring you're always working with the latest ticker information.

The Stock Symbol Search API is perfect for traders, investors, or anyone needing quick access to stock symbols across different markets.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"name": "Apple Inc.",
		"currency": "USD",
		"exchangeFullName": "NASDAQ Global Select",
		"exchange": "NASDAQ"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-symbol · 카테고리: search
