# Search Isin

Easily search and retrieve the International Securities Identification Number (ISIN) for financial securities using the FMP ISIN API. Find key details such as company name, stock symbol, and market capitalization associated with the ISIN.

## Endpoint

`GET https://financialmodelingprep.com/stable/search-isin?isin=US0378331005`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| isin* | string | US0378331005 |

## Description

The FMP ISIN API allows users to quickly retrieve comprehensive financial information linked to a specific ISIN (International Securities Identification Number). This twelve-character alphanumeric code uniquely identifies financial securities globally, making it an essential tool for investors, traders, and financial analysts.

Key features of the ISIN API include:

- Accurate Identification: Quickly find stock symbols and company names associated with a specific ISIN, ensuring precise identification of global securities.

- Comprehensive Data: Retrieve relevant financial details such as the company name, stock symbol, ISIN, and market capitalization.

- Global Coverage: The ISIN API supports a wide range of international securities, including stocks, bonds, and mutual funds, offering a broad range of search capabilities across global markets.

This API is a valuable resource for financial professionals needing to identify and analyze securities efficiently by their ISIN for global investments or research.

Example: An investor can use the ISIN API to locate the ISIN and market capitalization for Apple Inc. by searching for the stock symbol "AAPL," streamlining global investment research.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"name": "Apple Inc.",
		"isin": "US0378331005",
		"marketCap": 3900351299800
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-isin · 카테고리: search
