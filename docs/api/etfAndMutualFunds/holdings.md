# ETF & Fund Holdings

Get a detailed breakdown of the assets held within ETFs and mutual funds using the FMP ETF & Fund Holdings API. Access real-time data on the specific securities and their weights in the portfolio, providing insights into asset composition and fund strategies.

## Endpoint

`GET https://financialmodelingprep.com/stable/etf/holdings?symbol=SPY`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | SPY |

## Description

The FMP ETF & Fund Holdings API offers comprehensive information about the underlying assets that make up ETFs and mutual funds. This API is crucial for investors and analysts who need:

- Detailed Portfolio Insights: Gain visibility into the specific assets held within an ETF or mutual fund, including information such as asset names, symbols, ISINs, market values, and weight percentages. This helps investors understand a fund's exposure to particular stocks, sectors, or markets.

- Real-Time Updates: Stay informed with up-to-date information on fund holdings. The API provides real-time updates, ensuring that you always have access to the most current data on fund compositions.

- Investment Strategy Analysis: Use the holdings data to evaluate the investment strategy of different ETFs and mutual funds. By analyzing the securities and their respective weightings, you can make informed decisions about potential risks and opportunities.

For example, an investor interested in the SPY ETF can use this API to view Apple Inc.'s (AAPL) share count, market value, and its percentage weight in the fund, helping to assess the exposure to the tech sector.

## Response (example)

```json
[
	{
		"symbol": "SPY",
		"asset": "AAPL",
		"name": "APPLE INC",
		"isin": "US0378331005",
		"securityCusip": "037833100",
		"sharesNumber": 188106081,
		"weightPercentage": 7.137,
		"marketValue": 44744793487.47,
		"updatedAt": "2025-01-16 05:01:09",
		"updated": "2025-02-04 19:02:31"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/holdings · 카테고리: etfAndMutualFunds
