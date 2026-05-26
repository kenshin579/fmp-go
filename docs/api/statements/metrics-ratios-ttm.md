# Financial Ratios TTM

Gain access to trailing twelve-month (TTM) financial ratios with the TTM Ratios API. This API provides key performance metrics over the past year, including profitability, liquidity, and efficiency ratios.

## Endpoint

`GET https://financialmodelingprep.com/stable/ratios-ttm?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The TTM Ratios API offers a comprehensive view of a company's financial performance, making it an essential tool for investors, analysts, and decision-makers. This API is ideal for:

- Profitability Analysis: Understand how efficiently a company generates profit using metrics like gross profit margin, net profit margin, and EBITDA margin.

- Liquidity Assessment: Evaluate a company's ability to meet short-term obligations with ratios such as the current ratio and quick ratio.

- Efficiency Insight: Examine how well a company manages its assets and liabilities with key efficiency ratios like asset turnover and inventory turnover.

- Leverage Evaluation: Assess a company's debt levels and leverage through metrics like the debt-to-equity ratio and financial leverage ratio.

This API provides insights into a company's performance across key areas, helping users make more informed decisions by analyzing trends over the past twelve months.

Example Use Case
An investor uses the TTM Ratios API to analyze Apple's liquidity and profitability ratios, helping them decide whether to invest in the company based on its trailing twelve-month financial performance.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"grossProfitMarginTTM": 0.46518849807964424,
		"ebitMarginTTM": 0.3175535678188801,
		"ebitdaMarginTTM": 0.34705882352941175,
		"operatingProfitMarginTTM": 0.3175535678188801,
		"pretaxProfitMarginTTM": 0.31773296947645036,
		"continuousOperationsProfitMarginTTM": 0.24295027289266222,
		"netProfitMarginTTM": 0.24295027289266222,
		"bottomLineProfitMarginTTM": 0.24295027289266222,
		"receivablesTurnoverTTM": 6.673186524129093,
		"payablesTurnoverTTM": 3.4187853335486995,
		"inventoryTurnoverTTM": 30.626103313558097,
		"fixedAssetTurnoverTTM": 8.590592372311098,
		"assetTurnoverTTM": 1.1501809145995903,
		"currentRatioTTM": 0.9229383853427077,
		"quickRatioTTM": 0.8750666712845911,
		"solvencyRatioTTM": 0.3888081578786054,
		"cashRatioTTM": 0.20987774044955496,
		"priceToEarningsRatioTTM": 32.889608822880916,
		"priceToEarningsGrowthRatioTTM": 9.104441715061135,
		"forwardPriceToEarningsGrowthRatioTTM": 9.104441715061135,
		"priceToBookRatioTTM": 47.370141231313106,
		"priceToSalesRatioTTM": 7.958949686678795,
		"priceToFreeCashFlowRatioTTM": 32.04339747098139,
		"priceToOperatingCashFlowRatioTTM": 29.201395167968677,
		"debtToAssetsRatioTTM": 0.28132292892744526,
		"debtToEquityRatioTTM": 1.4499985020521886,
		"debtToCapitalRatioTTM": 0.5918364851397372,
		"longTermDebtToCapitalRatioTTM": 0.557055084464615,
		"financialLeverageRatioTTM": 5.154213727193745,
		"workingCapitalTurnoverRatioTTM": -22.92267593397046,
		"operatingCashFlowRatioTTM": 0.7501402694558931,
		"operatingCashFlowSalesRatioTTM": 0.2736355366889024,
		"freeCashFlowOperatingCashFlowRatioTTM": 0.9077049513361775,
		"debtServiceCoverageRatioTTM": 8.390251498870981,
		"interestCoverageRatioTTM": 0,
		"shortTermOperatingCashFlowCoverageRatioTTM": 8.432142022891847,
		"operatingCashFlowCoverageRatioTTM": 1.1187512267688715,
		"capitalExpenditureCoverageRatioTTM": 10.834817408704351,
		"dividendPaidAndCapexCoverageRatioTTM": 4.287173396674584,
		"dividendPayoutRatioTTM": 0.15876235049401977,
		"dividendYieldTTM": 0.0047691720717283476,
		"enterpriseValueTTM": 3216333928000,
		"revenuePerShareTTM": 26.24103186081379,
		"netIncomePerShareTTM": 6.375265851569754,
		"interestDebtPerShareTTM": 6.418298067250137,
		"cashPerShareTTM": 3.565573803101025,
		"bookValuePerShareTTM": 4.426417032959892,
		"tangibleBookValuePerShareTTM": 4.426417032959892,
		"shareholdersEquityPerShareTTM": 4.426417032959892,
		"operatingCashFlowPerShareTTM": 7.180478836504368,
		"capexPerShareTTM": 0.6627226436447186,
		"freeCashFlowPerShareTTM": 6.5177561928596495,
		"netIncomePerEBTTTM": 0.7646366484818603,
		"ebtPerEbitTTM": 1.0005649492739208,
		"priceToFairValueTTM": 47.370141231313106,
		"debtToMarketCapTTM": 0.030731461471514124,
		"effectiveTaxRateTTM": 0.23536335151813975,
		"enterpriseValueMultipleTTM": 23.41672438697653
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/metrics-ratios-ttm · 카테고리: statements
