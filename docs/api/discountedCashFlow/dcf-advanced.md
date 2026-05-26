# DCF Valuation

Estimate the intrinsic value of a company with the FMP Discounted Cash Flow Valuation API. Calculate the DCF valuation based on expected future cash flows and discount rates.

## Endpoint

`GET https://financialmodelingprep.com/stable/discounted-cash-flow?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Discounted Cash Flow (DCF) Valuation API provides investors with a powerful tool to estimate the value of an investment. DCF is a widely used valuation method that calculates the present value of a company's expected future cash flows. This API allows you to:

- Calculate DCF Valuation: Easily compute the DCF valuation by providing the company's expected future cash flows and the appropriate discount rate.

- Assess Investment Opportunities: Use DCF to compare the intrinsic value of different investments, helping you identify undervalued or overvalued assets.

- Evaluate Investment Risk: Analyze the riskiness of an investment by understanding how sensitive the DCF valuation is to changes in cash flows or discount rates.

The FMP Discounted Cash Flow Valuation API simplifies the DCF calculation process, allowing users to input the necessary financial data and quickly obtain a valuation result.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2026-04-08",
		"dcf": 159.36622443786206,
		"Stock Price": 258.25
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/dcf-advanced · 카테고리: discountedCashFlow
