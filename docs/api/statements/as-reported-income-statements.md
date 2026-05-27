# As Reported Income Statements

Retrieve income statements as they were reported by the company with the As Reported Income Statements API. Access raw financial data directly from official company filings, including revenue, expenses, and net income.

## Endpoint

`GET https://financialmodelingprep.com/stable/income-statement-as-reported?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | annual,quarter |

## Description

The As Reported Income Statements API provides a clear and direct view of a company's financial performance as reported in their official financial statements. This API is useful for:

- Direct Financial Insights: Access income statement data as reported by the company, without adjustments.

- Comprehensive Expense Tracking: See detailed breakdowns of revenue, cost of goods sold, and operating expenses.

- In-Depth Analysis: Use the raw data to perform your own calculations and build models based on official figures.

This API allows investors and analysts to rely on the most accurate, company-provided financial information for evaluating profitability and operational efficiency.

Example Use Case
A financial analyst can use the As Reported Income Statements API to access Apple's quarterly income statements, allowing them to compare operating income and net profit for different fiscal periods without any adjustments.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"fiscalYear": 2024,
		"period": "FY",
		"reportedCurrency": null,
		"date": "2024-09-27",
		"data": {
			"revenuefromcontractwithcustomerexcludingassessedtax": 391035000000,
			"costofgoodsandservicessold": 210352000000,
			"grossprofit": 180683000000,
			"researchanddevelopmentexpense": 31370000000,
			"sellinggeneralandadministrativeexpense": 26097000000,
			"operatingexpenses": 57467000000,
			"operatingincomeloss": 123216000000,
			"nonoperatingincomeexpense": 269000000,
			"incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterest": 123485000000,
			"incometaxexpensebenefit": 29749000000,
			"netincomeloss": 93736000000,
			"earningspersharebasic": 6.11,
			"earningspersharediluted": 6.08,
			"weightedaveragenumberofsharesoutstandingbasic": 15343783000,
			"weightedaveragenumberofdilutedsharesoutstanding": 15408095000,
			"othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftax": 395000000,
			"othercomprehensiveincomelossderivativeinstrumentgainlossbeforereclassificationaftertax": -832000000,
			"othercomprehensiveincomelossderivativeinstrumentgainlossreclassificationaftertax": 1337000000,
			"othercomprehensiveincomelossderivativeinstrumentgainlossafterreclassificationandtax": -2169000000,
			"othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftax": 5850000000,
			"othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftax": -204000000,
			"othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftax": 6054000000,
			"othercomprehensiveincomelossnetoftaxportionattributabletoparent": 4280000000,
			"comprehensiveincomenetoftax": 98016000000
		}
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/as-reported-income-statements · 카테고리: statements
