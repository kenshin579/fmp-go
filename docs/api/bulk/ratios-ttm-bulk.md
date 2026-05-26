# Ratios TTM Bulk

The Ratios TTM Bulk API offers an efficient way to retrieve trailing twelve months (TTM) financial ratios for stocks. It provides users with detailed insights into a company’s profitability, liquidity, efficiency, leverage, and valuation ratios, all based on the most recent financial report.

## Endpoint

`GET https://financialmodelingprep.com/stable/ratios-ttm-bulk`

## Description

With this API, you can access a wide array of financial ratios including:

- Profitability Ratios: Gross profit margin, operating profit margin, net profit margin, and more, helping investors assess how well a company is generating profit from its operations.

- Liquidity Ratios: Key liquidity measures such as current ratio, quick ratio, and cash ratio to understand how well a company can meet its short-term liabilities.

- Efficiency Ratios: Metrics such as receivables turnover, inventory turnover, and asset turnover to evaluate how efficiently a company utilizes its assets.

- Leverage Ratios: Debt-to-assets, debt-to-equity, and debt-to-capital ratios, which provide insight into a company's capital structure and financial leverage.

- Valuation Ratios: Ratios such as price-to-earnings (P/E), price-to-book (P/B), and price-to-sales (P/S) to help investors determine whether a stock is overvalued or undervalued.

- Cash Flow Ratios: Free cash flow yield, operating cash flow coverage, and capital expenditure coverage ratios to assess how well a company manages its cash flow relative to its operations and investments.

This API is invaluable for financial analysts, institutional investors, and portfolio managers who need to track and compare TTM ratios across multiple companies for investment decision-making.

## Response (example)

```json
[
	{
		"symbol": "000001.SZ",
		"grossProfitMarginTTM": "1.1622776732779352",
		"ebitMarginTTM": "0.22525536322293388",
		"ebitdaMarginTTM": "0.2018381390033096",
		"operatingProfitMarginTTM": "0.4658682349579752",
		"pretaxProfitMarginTTM": "0.3160551441700993",
		"continuousOperationsProfitMarginTTM": "0.25995857044215337",
		"netProfitMarginTTM": "0.25995857044215337",
		"bottomLineProfitMarginTTM": "0.25995857044215337",
		"receivablesTurnoverTTM": "0",
		"payablesTurnoverTTM": "0",
		"inventoryTurnoverTTM": "0",
		"fixedAssetTurnoverTTM": "13.114441842310695",
		"assetTurnoverTTM": "0.029075827062555015",
		"currentRatioTTM": "0",
		"quickRatioTTM": "0",
		"solvencyRatioTTM": "0.008534174446189174",
		"cashRatioTTM": "0",
		"priceToEarningsRatioTTM": "6.68445715569793",
		"priceToEarningsGrowthRatioTTM": "-3.6096068640768793",
		"forwardPriceToEarningsGrowthRatioTTM": "2.4481492401413427",
		"priceToBookRatioTTM": "0.576796465809228",
		"priceToSalesRatioTTM": "1.483200528584014",
		"priceToFreeCashFlowRatioTTM": "1.518395607609901",
		"priceToOperatingCashFlowRatioTTM": "1.7523793147342828",
		"debtToAssetsRatioTTM": "0",
		"debtToEquityRatioTTM": "0",
		"debtToCapitalRatioTTM": "0",
		"longTermDebtToCapitalRatioTTM": "0",
		"financialLeverageRatioTTM": "11.416164801466868",
		"workingCapitalTurnoverRatioTTM": "0.23544250931631752",
		"operatingCashFlowRatioTTM": "0",
		"operatingCashFlowSalesRatioTTM": "0.991612895545132",
		"freeCashFlowOperatingCashFlowRatioTTM": "0.9850828696116743",
		"debtServiceCoverageRatioTTM": "0.24758322210087771",
		"interestCoverageRatioTTM": "0.7914088096104842",
		"shortTermOperatingCashFlowCoverageRatioTTM": "0",
		"operatingCashFlowCoverageRatioTTM": "0",
		"capitalExpenditureCoverageRatioTTM": "67.03702213279678",
		"dividendPaidAndCapexCoverageRatioTTM": "6.192364879934577",
		"dividendPayoutRatioTTM": "0.5590996519509067",
		"dividendYieldTTM": "0.10335",
		"enterpriseValueTTM": "-496959244000",
		"revenuePerShareTTM": "7.389154370023568",
		"netIncomePerShareTTM": "1.9208740068077172",
		"interestDebtPerShareTTM": "4.349676503966586",
		"cashPerShareTTM": "32.81790720767194",
		"bookValuePerShareTTM": "22.260885357516656",
		"tangibleBookValuePerShareTTM": "21.662613507347245",
		"shareholdersEquityPerShareTTM": "22.260885357516656",
		"operatingCashFlowPerShareTTM": "7.327180760489036",
		"capexPerShareTTM": "0.10930051078304583",
		"freeCashFlowPerShareTTM": "7.21788024970599",
		"netIncomePerEBTTTM": "0.8225101702576465",
		"ebtPerEbitTTM": "0.6784217520188082",
		"priceToFairValueTTM": "0.576796465809228",
		"debtToMarketCapTTM": "0",
		"effectiveTaxRateTTM": "0.17748982974235347",
		"enterpriseValueMultipleTTM": "-14.656106051669223",
		"dividendPerShareTTM": "1.327"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/ratios-ttm-bulk · 카테고리: bulk
