# Key Metrics TTM

Retrieve a comprehensive set of trailing twelve-month (TTM) key performance metrics with the TTM Key Metrics API. Access data related to a company's profitability, capital efficiency, and liquidity, allowing for detailed analysis of its financial health over the past year.

## Endpoint

`GET https://financialmodelingprep.com/stable/key-metrics-ttm?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The TTM Key Metrics API provides valuable insights into a company's recent performance, capturing data over the trailing twelve-month period. This API is ideal for:

- Profitability Assessment: Understand a company's ability to generate profit, with metrics such as return on assets (ROA) and earnings yield.

- Liquidity and Solvency Analysis: Evaluate how efficiently a company manages its short-term obligations with ratios like the current ratio and cash conversion cycle.

- Capital Efficiency: Assess how well a company is using its capital with metrics like return on invested capital (ROIC) and return on equity (ROE).

- Operational Performance: Get insights into the operational efficiency of a company through operating cycle and days of inventory outstanding (DIO).

This API helps investors, analysts, and portfolio managers track financial performance trends and assess companies' efficiency in generating returns.

Example Use Case
An analyst can use the TTM Key Metrics API to compare the free cash flow yield of several companies within the same industry, helping them make better-informed investment decisions.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"marketCap": 3149833928000,
		"enterpriseValueTTM": 3216333928000,
		"evToSalesTTM": 8.126980816656559,
		"evToOperatingCashFlowTTM": 29.70001965021146,
		"evToFreeCashFlowTTM": 32.71990486169747,
		"evToEBITDATTM": 23.41672438697653,
		"netDebtToEBITDATTM": 0.48415749315627005,
		"currentRatioTTM": 0.9229383853427077,
		"incomeQualityTTM": 1.1263026521060842,
		"grahamNumberTTM": 25.198029099282905,
		"grahamNetNetTTM": -11.64435843011051,
		"taxBurdenTTM": 0.7646366484818603,
		"interestBurdenTTM": 1.0005649492739208,
		"workingCapitalTTM": -11125000000,
		"investedCapitalTTM": 34944000000,
		"returnOnAssetsTTM": 0.27943676707790227,
		"operatingReturnOnAssetsTTM": 0.35448090090471257,
		"returnOnTangibleAssetsTTM": 0.27943676707790227,
		"returnOnEquityTTM": 1.4534598087751787,
		"returnOnInvestedCapitalTTM": 0.45208108089346594,
		"returnOnCapitalEmployedTTM": 0.6292559583416784,
		"earningsYieldTTM": 0.030404739849149914,
		"freeCashFlowYieldTTM": 0.03120767705439485,
		"capexToOperatingCashFlowTTM": 0.09229504866382256,
		"capexToDepreciationTTM": 0.855956153121521,
		"capexToRevenueTTM": 0.025255205174853447,
		"salesGeneralAndAdministrativeToRevenueTTM": 0,
		"researchAndDevelopementToRevenueTTM": 0.08071053163533455,
		"stockBasedCompensationToRevenueTTM": 0.030263290883363655,
		"intangiblesToTotalAssetsTTM": 0,
		"averageReceivablesTTM": 62774500000,
		"averagePayablesTTM": 65435000000,
		"averageInventoryTTM": 7098500000,
		"daysOfSalesOutstandingTTM": 54.69650798463715,
		"daysOfPayablesOutstandingTTM": 106.76306476988712,
		"daysOfInventoryOutstandingTTM": 11.917937984569374,
		"operatingCycleTTM": 66.61444596920653,
		"cashConversionCycleTTM": -40.148618800680595,
		"freeCashFlowToEquityTTM": 31799000000,
		"freeCashFlowToFirmTTM": 85497710797.9578,
		"tangibleAssetValueTTM": 66758000000,
		"netCurrentAssetValueTTM": -144087000000
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/key-metrics-ttm · 카테고리: statements
