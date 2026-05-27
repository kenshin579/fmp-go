# Price Target Consensus

Access analysts' consensus price targets with the FMP Price Target Consensus API. This API provides high, low, median, and consensus price targets for stocks, offering investors a comprehensive view of market expectations for future stock prices.

## Endpoint

`GET https://financialmodelingprep.com/stable/price-target-consensus?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Price Target Consensus API delivers key insights into stock price forecasts by aggregating price targets from analysts. This allows investors to make more informed decisions based on the following metrics:

- High Price Target: See the highest price target forecasted by analysts.

- Low Price Target: Access the lowest expected price for a stock, providing insight into downside risk.

- Median Price Target: Get the median price target to understand the central tendency of analysts' predictions.

- Consensus Price Target: Retrieve the overall consensus target, which reflects the average of analysts' forecasts.

This API offers a broad perspective on price expectations, helping users to evaluate the potential range of stock movements based on expert predictions.

Example Use Case
A portfolio manager can use the Price Target Consensus API to assess the potential upside and downside for a stock, using the high, low, median, and consensus price targets to create risk-reward scenarios for investment decisions.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"targetHigh": 300,
		"targetLow": 200,
		"targetConsensus": 251.7,
		"targetMedian": 258
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/price-target-consensus · 카테고리: analyst
