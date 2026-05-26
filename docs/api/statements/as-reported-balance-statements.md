# As Reported Balance Statements

Access balance sheets as reported by the company with the As Reported Balance Statements API. View detailed financial data on assets, liabilities, and equity directly from official filings.

## Endpoint

`GET https://financialmodelingprep.com/stable/balance-sheet-statement-as-reported?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |
| period | string | annual,quarter |

## Description

The As Reported Balance Statements API offers unadjusted balance sheet data as reported by companies. It provides insight into a company's financial position, including:

- Asset Overview: View cash, receivables, inventory, and long-term assets as reported.

- Liability Breakdown: Access current and non-current liabilities, deferred revenues, and more.

- Equity Insights: Examine stockholders' equity, including retained earnings and stock details.

This API is ideal for analysts and investors who want raw, as-reported balance sheet data to perform accurate financial assessments.

Example Use Case
An investment analyst can use the As Reported Balance Statements API to evaluate Apple's asset-liability structure for Q1 2010, helping to understand the company's financial position during that period without any adjustments.

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
			"cashandcashequivalentsatcarryingvalue": 29943000000,
			"marketablesecuritiescurrent": 35228000000,
			"accountsreceivablenetcurrent": 33410000000,
			"nontradereceivablescurrent": 32833000000,
			"inventorynet": 7286000000,
			"otherassetscurrent": 14287000000,
			"assetscurrent": 152987000000,
			"marketablesecuritiesnoncurrent": 91479000000,
			"propertyplantandequipmentnet": 45680000000,
			"otherassetsnoncurrent": 74834000000,
			"assetsnoncurrent": 211993000000,
			"assets": 364980000000,
			"accountspayablecurrent": 68960000000,
			"otherliabilitiescurrent": 78304000000,
			"contractwithcustomerliabilitycurrent": 8249000000,
			"commercialpaper": 10000000000,
			"longtermdebtcurrent": 10912000000,
			"liabilitiescurrent": 176392000000,
			"longtermdebtnoncurrent": 85750000000,
			"otherliabilitiesnoncurrent": 45888000000,
			"liabilitiesnoncurrent": 131638000000,
			"liabilities": 308030000000,
			"commonstocksharesoutstanding": 15116786000,
			"commonstocksharesissued": 15116786000,
			"commonstocksincludingadditionalpaidincapital": 83276000000,
			"retainedearningsaccumulateddeficit": -19154000000,
			"accumulatedothercomprehensiveincomelossnetoftax": -7172000000,
			"stockholdersequity": 56950000000,
			"liabilitiesandstockholdersequity": 364980000000,
			"commonstockparorstatedvaluepershare": 0.00001,
			"commonstocksharesauthorized": 50400000000
		}
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/as-reported-balance-statements · 카테고리: statements
