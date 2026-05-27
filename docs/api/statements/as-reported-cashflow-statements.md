# As Reported Cashflow Statements

View cash flow statements as reported by the company with the As Reported Cash Flow Statements API. Analyze a company's cash flows related to operations, investments, and financing directly from official reports.

## Endpoint

`GET https://financialmodelingprep.com/stable/cash-flow-statement-as-reported?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | annual,quarter |

## Description

The As Reported Cash Flow Statements API provides access to unadjusted cash flow data as reported by companies. This includes:

- Operational Cash Flows: Examine the cash generated or used in day-to-day business activities.

- Investment Cash Flows: Access cash movements related to investments in assets, acquisitions, and securities.

- Financing Cash Flows: View cash from equity, debt issuance, and dividend payments.

This API is ideal for users looking for a clear understanding of a company's cash flow management based on official filings.

Example Use Case
A financial analyst can use this API to track Apple's cash flow trends during Q1 2010, helping assess how effectively the company is managing its cash for operations and investments.

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
			"cashcashequivalentsrestrictedcashandrestrictedcashequivalents": 29943000000,
			"netincomeloss": 93736000000,
			"depreciationdepletionandamortization": 11445000000,
			"sharebasedcompensation": 11688000000,
			"othernoncashincomeexpense": 2266000000,
			"increasedecreaseinaccountsreceivable": 3788000000,
			"increasedecreaseinotherreceivables": 1356000000,
			"increasedecreaseininventories": 1046000000,
			"increasedecreaseinotheroperatingassets": 11731000000,
			"increasedecreaseinaccountspayable": 6020000000,
			"increasedecreaseinotheroperatingliabilities": 15552000000,
			"netcashprovidedbyusedinoperatingactivities": 118254000000,
			"paymentstoacquireavailableforsalesecuritiesdebt": 48656000000,
			"proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities": 51211000000,
			"proceedsfromsaleofavailableforsalesecuritiesdebt": 11135000000,
			"paymentstoacquirepropertyplantandequipment": 9447000000,
			"paymentsforproceedsfromotherinvestingactivities": 1308000000,
			"netcashprovidedbyusedininvestingactivities": 2935000000,
			"paymentsrelatedtotaxwithholdingforsharebasedcompensation": 5600000000,
			"paymentsofdividends": 15234000000,
			"paymentsforrepurchaseofcommonstock": 94949000000,
			"repaymentsoflongtermdebt": 9958000000,
			"proceedsfromrepaymentsofcommercialpaper": 3960000000,
			"proceedsfrompaymentsforotherfinancingactivities": -361000000,
			"netcashprovidedbyusedinfinancingactivities": -121983000000,
			"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect": -794000000,
			"incometaxespaidnet": 26102000000
		}
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/as-reported-cashflow-statements · 카테고리: statements
