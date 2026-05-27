# Financial Statement Symbols List

Access a comprehensive list of companies with available financial statements through the FMP Financial Statement Symbols List API. Find companies listed on major global exchanges and obtain up-to-date financial data including income statements, balance sheets, and cash flow statements, are provided.

## Endpoint

`GET https://financialmodelingprep.com/stable/financial-statement-symbol-list`

## Description

The FMP Financial Statement Symbols List API provides a complete list of companies for which financial statements are available through our API. This endpoint is essential for:

- Comprehensive Company Coverage: Discover all companies with available financial statements, including those listed on major exchanges like the NYSE and NASDAQ, as well as international exchanges.

- Access to Global Financial Data: Gain insights into companies from around the world by accessing their financial statements through this extensive symbol list.

- Up-to-Date Information: Stay informed with regularly updated lists, ensuring you have access to the latest financial statements for public companies.

Example: An investor can use the Financial Statement Symbols List API to find the ticker symbol for a company they are interested in, access its financial statements, and make informed investment decisions based on the latest available data.

## Response (example)

```json
[
	{
		"symbol": "6898.HK",
		"companyName": "China Aluminum Cans Holdings Limited",
		"tradingCurrency": "HKD",
		"reportingCurrency": "HKD"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/financial-symbols-list · 카테고리: directory
