# Price Target Summary Bulk

The Price Target Summary Bulk API provides a comprehensive overview of price targets for all listed symbols over multiple timeframes. With this API, users can quickly retrieve price target data, helping investors and analysts compare current prices to projected targets across different periods.

## Endpoint

`GET https://financialmodelingprep.com/stable/price-target-summary-bulk`

## Description

This API enables users to access price targets for all companies, offering insights into:

- Price Targets Over Timeframes: Retrieve price target data for symbols, including insights for the last month, quarter, year, and all-time periods.

- Average Price Target: View the average price target set by analysts and market experts for each symbol.

- Price Target Differences: Analyze the percentage difference between current prices and price targets across various timeframes.

- Publisher Data: Identify the sources and publishers providing these price targets, offering an understanding of the context and reliability of the data.

The Price Target Summary Bulk API is ideal for institutional investors, analysts, and traders seeking a holistic view of stock price forecasts and analyst expectations.

## Response (example)

```json
[
	{
		"symbol": "A",
		"lastMonthCount": "0",
		"lastMonthAvgPriceTarget": "0",
		"lastQuarterCount": "1",
		"lastQuarterAvgPriceTarget": "116",
		"lastYearCount": "6",
		"lastYearAvgPriceTarget": "142.17",
		"allTimeCount": "18",
		"allTimeAvgPriceTarget": "146.61",
		"publishers": "[\"\"TheFly\""
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/price-target-summary-bulk · 카테고리: bulk
