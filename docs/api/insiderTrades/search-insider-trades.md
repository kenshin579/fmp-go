# Search Insider Trades

Search insider trading activity by company or symbol using the Search Insider Trades API. Find specific trades made by corporate insiders, including executives and directors.

## Endpoint

`GET https://financialmodelingprep.com/stable/insider-trading/search?page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol | string | AAPL |
| page | number | 0 |
| limit | number | 100 |
| reportingCik | string | 0001496686 |
| companyCik | string | 0000320193 |
| transactionType | string | S-Sale |

## Description

The FMP Search Insider Trades API allows users to search for specific insider trading activities based on a company or stock symbol. This API provides detailed information on stock transactions by corporate insiders, including transaction dates, types, amounts, and roles within the company. Key features include:

- Company-Specific Searches: Search insider trading activity by entering the stock symbol or company name to retrieve relevant transactions.

- Detailed Transaction Information: Access detailed data such as transaction type (purchase or sale), number of securities transacted, and price.

- Insider Roles: Understand the roles of the insiders involved in the transactions, such as directors or executives.

- Direct Links to Filings: Each transaction includes a link to the official SEC filing for deeper analysis and verification.

This API is perfect for investors, financial researchers, and analysts who need to investigate insider trading activities of specific companies or individuals.

Example Use Case
An investment analyst uses the Search Insider Trades API to investigate recent sales of Apple (AAPL) stock by Chris Kondo, the Principal Accounting Officer. By retrieving detailed information about the transaction, including the sale of 8,706 shares at $225, the analyst can better assess the implications for the company's financial performance and strategy.

## Response (example)

```json
[
	{
		"symbol": "LAB",
		"filingDate": "2026-04-08",
		"transactionDate": "2026-04-06",
		"reportingCik": "0001559779",
		"companyCik": "0001162194",
		"transactionType": "M-Exempt",
		"securitiesOwned": 6790596,
		"reportingName": "Egholm Michael",
		"typeOfOwner": "director, officer: President & CEO",
		"acquisitionOrDisposition": "A",
		"directOrIndirect": "D",
		"formType": "4",
		"securitiesTransacted": 196513,
		"price": 0,
		"securityName": "Common Stock",
		"url": "https://www.sec.gov/Archives/edgar/data/1162194/000119312526148615/0001193125-26-148615-index.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-insider-trades · 카테고리: insiderTrades
