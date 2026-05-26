# Key Metrics TTM Bulk

The Key Metrics TTM Bulk API allows users to retrieve trailing twelve months (TTM) data for all companies available in the database. The API provides critical financial ratios and metrics based on each company’s latest financial report, offering insights into company performance and financial health.

## Endpoint

`GET https://financialmodelingprep.com/stable/key-metrics-ttm-bulk`

## Description

This API gives access to:

- Market and Enterprise Value Metrics: Get TTM market capitalization, enterprise value, and other valuation multiples such as EV to sales, operating cash flow, and free cash flow.

- Profitability and Return Ratios: Track key ratios including return on assets (ROA), return on equity (ROE), return on invested capital (ROIC), and more.

- Operational Efficiency Metrics: Access metrics like the cash conversion cycle, days of payables, receivables, and inventory outstanding, providing insight into a company's operational efficiency.

- Liquidity and Leverage Ratios: Monitor liquidity with the current ratio and assess financial leverage through net debt to EBITDA and other relevant ratios.

- Cash Flow and Yield Metrics: Evaluate cash flow-related metrics, such as free cash flow yield, earnings yield, and capex to revenue ratios, helping investors understand how well a company generates and uses cash.

This API is especially useful for analysts, portfolio managers, and institutional investors seeking to monitor key financial metrics across a large universe of companies. The API is non-filterable, providing the most recent TTM data based on the latest financial filings.

## Response (example)

```json
[
	{
		"symbol": "000001.SZ",
		"marketCap": "249171756000",
		"enterpriseValueTTM": "-496959244000",
		"evToSalesTTM": "-2.95816117050406",
		"evToOperatingCashFlowTTM": "-2.9831814247210167",
		"evToFreeCashFlowTTM": "-3.028355803098073",
		"evToEBITDATTM": "-14.656106051669223",
		"netDebtToEBITDATTM": "-22.004571192638906",
		"currentRatioTTM": "0",
		"incomeQualityTTM": "15.217593861331872",
		"grahamNumberTTM": "31.017865999534138",
		"grahamNetNetTTM": "-199.05514330278228",
		"taxBurdenTTM": "0.8225101702576465",
		"interestBurdenTTM": "1.4030970878917606",
		"workingCapitalTTM": "746131000000",
		"investedCapitalTTM": "772543000000",
		"returnOnAssetsTTM": "0.007558510437605078",
		"operatingReturnOnAssetsTTM": "0.013555578495362656",
		"returnOnTangibleAssetsTTM": "0.007576346366296015",
		"returnOnEquityTTM": "0.09082717681735725",
		"returnOnInvestedCapitalTTM": "0.011141314993384131",
		"returnOnCapitalEmployedTTM": "0.013545504233575834",
		"earningsYieldTTM": "0.14960077934639543",
		"freeCashFlowYieldTTM": "0.6585898925077207",
		"capexToOperatingCashFlowTTM": "0.014917130388325619",
		"capexToDepreciationTTM": "1.855862584017924",
		"capexToRevenueTTM": "0.014792018857591847",
		"salesGeneralAndAdministrativeToRevenueTTM": "0.10163337222314817",
		"researchAndDevelopementToRevenueTTM": "0",
		"stockBasedCompensationToRevenueTTM": "0",
		"intangiblesToTotalAssetsTTM": "0.002354159621091415",
		"averageReceivablesTTM": "0",
		"averagePayablesTTM": "0",
		"averageInventoryTTM": "0",
		"daysOfSalesOutstandingTTM": "0",
		"daysOfPayablesOutstandingTTM": "0",
		"daysOfInventoryOutstandingTTM": "0",
		"operatingCycleTTM": "0",
		"cashConversionCycleTTM": "0",
		"freeCashFlowToEquityTTM": "910233000000",
		"freeCashFlowToFirmTTM": "-35237570137.11014",
		"tangibleAssetValueTTM": "492510000000",
		"netCurrentAssetValueTTM": "-4525615000000"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/key-metrics-ttm-bulk · 카테고리: bulk
