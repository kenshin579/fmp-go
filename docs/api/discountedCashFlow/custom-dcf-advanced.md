# Custom DCF Advanced

Run a tailored Discounted Cash Flow (DCF) analysis using the FMP Custom DCF Advanced API. With detailed inputs, this API allows users to fine-tune their assumptions and variables, offering a more personalized and precise valuation for a company.

## Endpoint

`GET https://financialmodelingprep.com/stable/custom-discounted-cash-flow?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| revenueGrowthPct | number | 0.1094119804597946 |
| ebitdaPct | number | 0.31273548388 |
| depreciationAndAmortizationPct | number | 0.0345531631720999 |
| cashAndShortTermInvestmentsPct | number | 0.2344222126801843 |
| receivablesPct | number | 0.1533770531229388 |
| inventoriesPct | number | 0.0155245674227653 |
| payablePct | number | 0.1614868903169657 |
| ebitPct | number | 0.2781823207138459 |
| capitalExpenditurePct | number | 0.0306025847141713 |
| operatingCashFlowPct | number | 0.2886333485760204 |
| sellingGeneralAndAdministrativeExpensesPct | number | 0.0662854095187211 |
| taxRate | number | 0.14919579658453103 |
| longTermGrowthRate | number | 4 |
| costOfDebt | number | 3.64 |
| costOfEquity | number | 9.51168 |
| marketRiskPremium | number | 4.72 |
| beta | number | 1.244 |
| riskFreeRate | number | 3.64 |

## Description

The Custom DCF Advanced API is designed for financial analysts and investors who want to customize their DCF analysis based on their specific forecasts and assumptions. This API gives users the flexibility to modify key variables such as revenue growth, EBITDA, capital expenditures, and risk factors to achieve a tailored company valuation. Key features include:

- Customizable Inputs: Adjust core financial metrics such as revenue, EBITDA, and capital expenditures to fit your projections and forecasts.

- Advanced Financial Assumptions: Modify factors like the risk-free rate, market risk premium, tax rate, and WACC to create a more accurate valuation.

- Comprehensive Output: Get detailed results including equity value, free cash flow, terminal value, and equity value per share, all based on your custom inputs.

This API is ideal for professional analysts or advanced users looking to customize DCF models to reflect their investment strategy or valuation assumptions.

Example Use Case
An equity analyst might use the Custom DCF Advanced API to adjust Apple's financial forecasts, input a different market risk premium, or modify the long-term growth rate. These tailored inputs allow the analyst to create a unique valuation model for the company and make more informed investment decisions.

## Response (example)

```json
[
	{
		"year": "2030",
		"symbol": "AAPL",
		"revenue": 529528728806,
		"revenuePercentage": 4.09,
		"ebitda": 191125428209,
		"ebitdaPercentage": 36.09,
		"ebit": 177353356628,
		"ebitPercentage": 33.49,
		"depreciation": 15508463644,
		"depreciationPercentage": 2.93,
		"totalCash": 79685715467,
		"totalCashPercentage": 15.05,
		"receivables": 114078294622,
		"receivablesPercentage": 21.54,
		"inventories": 8411056160,
		"inventoriesPercentage": 1.59,
		"payable": 101862682518,
		"payablePercentage": 19.24,
		"capitalExpenditure": -14907445037,
		"capitalExpenditurePercentage": -2.82,
		"price": 262.82,
		"beta": 1.109,
		"dilutedSharesOutstanding": 15004697000,
		"costofDebt": 3.92,
		"taxRate": 15.61,
		"afterTaxCostOfDebt": 3.31,
		"riskFreeRate": 3.92,
		"marketRiskPremium": 4.72,
		"costOfEquity": 9.15,
		"totalDebt": 112377000000,
		"totalEquity": 3943534465540,
		"totalCapital": 4055911465540,
		"debtWeighting": 2.77,
		"equityWeighting": 97.23,
		"wacc": 8.99,
		"taxRateCash": 16785417,
		"ebiat": 147583856418,
		"ufcf": 145836268225,
		"sumPvUfcf": 505377678906,
		"longTermGrowthRate": 4,
		"terminalValue": 3038731862013,
		"presentTerminalValue": 1975763045693,
		"enterpriseValue": 2481140724600,
		"netDebt": 76443000000,
		"equityValue": 2404697724600,
		"equityValuePerShare": 160.26,
		"freeCashFlowT1": 151669718954
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/custom-dcf-advanced · 카테고리: discountedCashFlow
