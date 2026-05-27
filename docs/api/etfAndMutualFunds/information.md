# ETF & Mutual Fund Information

Access comprehensive data on ETFs and mutual funds with the FMP ETF & Mutual Fund Information API. Retrieve essential details such as ticker symbol, fund name, expense ratio, assets under management, and more.

## Endpoint

`GET https://financialmodelingprep.com/stable/etf/info?symbol=SPY`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | SPY |

## Description

The FMP ETF & Mutual Fund Information API offers a detailed look into the financial and structural information of ETFs and mutual funds. This API enables investors to:

- Compare Funds: Evaluate different ETFs and mutual funds by reviewing key metrics like ticker symbol, name, expense ratio, and assets under management to choose the most cost-effective and suitable investment options.

- Identify Investment Opportunities: Use the detailed data to discover ETFs and mutual funds that align with your specific investment strategy, risk tolerance, and financial goals.

- Understand Investment Objectives: Learn more about the objectives and strategies of various ETFs and mutual funds, helping you assess their suitability for inclusion in your portfolio based on asset class, sector exposure, and expense ratios.

For example, an investor can use this API to compare the expense ratios of various ETFs and mutual funds, find funds with large assets under management, or analyze sector weightings to ensure their investments align with their market outlook.

## Response (example)

```json
[
	{
		"symbol": "SPY",
		"name": "SPDR S&P 500 ETF Trust",
		"description": "The Trust seeks to achieve its investment objective by holding a portfolio of the common stocks that are included in the index (the “Portfolio”), with the weight of each stock in the Portfolio substantially corresponding to the weight of such stock in the index.",
		"isin": "US78462F1030",
		"assetClass": "Equity",
		"securityCusip": "78462F103",
		"domicile": "US",
		"website": "https://www.ssga.com/us/en/institutional/etfs/spdr-sp-500-etf-trust-spy",
		"etfCompany": "SPDR",
		"expenseRatio": 0.0945,
		"assetsUnderManagement": 633120180000,
		"avgVolume": 46396400,
		"inceptionDate": "1993-01-22",
		"nav": 603.64,
		"navCurrency": "USD",
		"holdingsCount": 503,
		"updatedAt": "2024-12-03T20:32:48.873Z",
		"sectorsList": [
			{
				"industry": "Basic Materials",
				"exposure": 1.97
			},
			{
				"industry": "Communication Services",
				"exposure": 8.87
			},
			{
				"industry": "Consumer Cyclical",
				"exposure": 9.84
			}
		]
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/information · 카테고리: etfAndMutualFunds
