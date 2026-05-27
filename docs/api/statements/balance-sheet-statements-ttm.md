# Balance Sheet Statements TTM

## Endpoint

`GET https://financialmodelingprep.com/stable/balance-sheet-statement-ttm?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 5 |

## Response (example)

```json
[
	{
		"date": "2024-12-28",
		"symbol": "AAPL",
		"reportedCurrency": "USD",
		"cik": "0000320193",
		"filingDate": "2025-01-31",
		"acceptedDate": "2025-01-31 06:01:27",
		"fiscalYear": "2025",
		"period": "Q1",
		"cashAndCashEquivalents": 30299000000,
		"shortTermInvestments": 23476000000,
		"cashAndShortTermInvestments": 53775000000,
		"netReceivables": 59306000000,
		"accountsReceivables": 29639000000,
		"otherReceivables": 29667000000,
		"inventory": 6911000000,
		"prepaids": 0,
		"otherCurrentAssets": 13248000000,
		"totalCurrentAssets": 133240000000,
		"propertyPlantEquipmentNet": 46069000000,
		"goodwill": 0,
		"intangibleAssets": 0,
		"goodwillAndIntangibleAssets": 0,
		"longTermInvestments": 87593000000,
		"taxAssets": 0,
		"otherNonCurrentAssets": 77183000000,
		"totalNonCurrentAssets": 210845000000,
		"otherAssets": 0,
		"totalAssets": 344085000000,
		"totalPayables": 61910000000,
		"accountPayables": 61910000000,
		"otherPayables": 0,
		"accruedExpenses": 0,
		"shortTermDebt": 12843000000,
		"capitalLeaseObligationsCurrent": 0,
		"taxPayables": 0,
		"deferredRevenue": 8461000000,
		"otherCurrentLiabilities": 61151000000,
		"totalCurrentLiabilities": 144365000000,
		"longTermDebt": 83956000000,
		"deferredRevenueNonCurrent": 0,
		"deferredTaxLiabilitiesNonCurrent": 0,
		"otherNonCurrentLiabilities": 49006000000,
		"totalNonCurrentLiabilities": 132962000000,
		"otherLiabilities": 0,
		"capitalLeaseObligations": 0,
		"totalLiabilities": 277327000000,
		"treasuryStock": 0,
		"preferredStock": 0,
		"commonStock": 84768000000,
		"retainedEarnings": -11221000000,
		"additionalPaidInCapital": 0,
		"accumulatedOtherComprehensiveIncomeLoss": -6789000000,
		"otherTotalStockholdersEquity": 0,
		"totalStockholdersEquity": 66758000000,
		"totalEquity": 66758000000,
		"minorityInterest": 0,
		"totalLiabilitiesAndTotalEquity": 344085000000,
		"totalInvestments": 111069000000,
		"totalDebt": 96799000000,
		"netDebt": 66500000000
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/balance-sheet-statements-ttm · 카테고리: statements
