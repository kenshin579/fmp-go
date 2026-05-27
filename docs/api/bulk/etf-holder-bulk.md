# ETF Holder Bulk

The ETF Holder Bulk API allows users to quickly retrieve detailed information about the assets and shares held by Exchange-Traded Funds (ETFs). This API provides insights into the weight each asset carries within the ETF, along with key financial information related to these holdings.

## Endpoint

`GET https://financialmodelingprep.com/stable/etf-holder-bulk?part=1`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| part* | string | 1 |

## Description

The ETF Holder Bulk API enables users to access:

Comprehensive Asset Lists: Retrieve a list of all assets held by an ETF, including individual stocks, bonds, and other securities.

- Share Information: View the number of shares an ETF holds for each asset, providing clarity on the distribution of holdings.

- Weight Percentage: Analyze the percentage weight of each asset within the ETF, helping investors understand its contribution to the ETF's overall value.

- Market Value: Get up-to-date market values for each asset held by the ETF, giving a complete picture of the ETF's composition.

- ISIN and CUSIP Identifiers: Identify assets by their ISIN or CUSIP for more precise tracking and research.

The ETF Holder Bulk API is an essential tool for financial analysts, institutional investors, and portfolio managers who need to analyze ETF composition, asset allocation, and potential risks or opportunities.

## Response (example)

```json
[
	{
		"symbol": "EXCH.AS",
		"name": "SAMSUNG ELECTRO MECHANICS LTD",
		"sharesNumber": "15514",
		"asset": "009150.KS",
		"weightPercentage": "0.09611",
		"cusip": "",
		"isin": "KR7009150004",
		"marketValue": "1553142.49",
		"lastUpdated\"": "2024-09-06\""
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/etf-holder-bulk · 카테고리: bulk
