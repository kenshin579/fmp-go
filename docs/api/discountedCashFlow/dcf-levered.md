# Levered DCF

Analyze a company’s value with the FMP Levered Discounted Cash Flow (DCF) API, which incorporates the impact of debt. This API provides post-debt company valuation, offering investors a more accurate measure of a company's true worth by accounting for its debt obligations.

## Endpoint

`GET https://financialmodelingprep.com/stable/levered-discounted-cash-flow?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The Levered DCF API is designed for investors and analysts looking to assess a company's valuation with more precision. By factoring in debt, it delivers a realistic estimate of the company's value. Key features include:

- Post-Debt Valuation: Provides a clear picture of the company's value after considering its debt load, which is crucial for assessing the risk and return profile of an investment.

- DCF Value vs. Market Price: Compare the discounted cash flow valuation to the current stock price to assess whether a stock is overvalued or undervalued.

- Informed Investment Decisions: With a levered DCF approach, investors can make better decisions by understanding the impact of financial obligations on a company's value.

This API is essential for performing deeper financial analysis and gaining a holistic view of a company's valuation.

Example Use Case
An investor evaluating whether to buy Apple shares can use the Levered DCF API to compare the company's DCF value to its current stock price. If the DCF is significantly lower than the market price, the investor might reconsider the purchase, factoring in the company's debt obligations.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2026-04-08",
		"dcf": 152.32738976131944,
		"Stock Price": 258.25
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/dcf-levered · 카테고리: discountedCashFlow
