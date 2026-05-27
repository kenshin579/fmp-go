# Mutual Fund Disclosures

Access comprehensive disclosure data for mutual funds with the FMP Mutual Fund Disclosures API. Analyze recent filings, balance sheets, and financial reports to gain insights into mutual fund portfolios.

## Endpoint

`GET https://financialmodelingprep.com/stable/funds/disclosure?symbol=VWO&year=2023&quarter=4`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | VWO |
| year* | string | 2023 |
| quarter* | string | 4 |
| cik | string | 0000857489 |

## Description

The FMP Mutual Fund Disclosures API provides detailed information on mutual fund holdings and recent filings, allowing investors and financial professionals to:

- Track Fund Holdings: Review the most recent disclosures of mutual fund holdings, including asset categories, issuer information, and country of investment. This helps users understand the portfolio composition of various mutual funds.

- Analyze Recent Filings: Obtain critical financial reports and filings from mutual funds, including balance data, market value in USD, percentage of total portfolio value, and more. These insights can help with investment analysis and strategy development.

- Gain Transparency into Investments: The API provides essential details like CUSIP, ISIN, issuer category, and fair value levels, offering full transparency into mutual fund investments.

For example, an investor can use this API to review the holdings of a mutual fund, such as Realty Income Corp, analyzing the balance, value in USD, and percentage of portfolio allocation to help make informed investment decisions.

## Response (example)

```json
[
	{
		"cik": "0000857489",
		"date": "2023-10-31",
		"acceptedDate": "2023-12-28 09:26:13",
		"symbol": "000089.SZ",
		"name": "Shenzhen Airport Co Ltd",
		"lei": "3003009W045RIKRBZI44",
		"title": "SHENZ AIRPORT-A",
		"cusip": "N/A",
		"isin": "CNE000000VK1",
		"balance": 2438784,
		"units": "NS",
		"cur_cd": "CNY",
		"valUsd": 2255873.6,
		"pctVal": 0.0023838966190458215,
		"payoffProfile": "Long",
		"assetCat": "EC",
		"issuerCat": "CORP",
		"invCountry": "CN",
		"isRestrictedSec": "N",
		"fairValLevel": "2",
		"isCashCollateral": "N",
		"isNonCashCollateral": "N",
		"isLoanByFund": "N"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/mutual-fund-disclosures · 카테고리: etfAndMutualFunds
