# Mutual Fund & ETF Disclosure

Access the latest disclosures from mutual funds and ETFs with the FMP Mutual Fund & ETF Disclosure API. This API provides updates on filings, changes in holdings, and other critical disclosure data for mutual funds and ETFs.

## Endpoint

`GET https://financialmodelingprep.com/stable/funds/disclosure-holders-latest?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |

## Description

The FMP Mutual Fund & ETF Disclosure API delivers up-to-date information on the holdings and strategy changes of mutual funds and ETFs. This API is designed for investors, analysts, and financial professionals who need to:

- Track Fund Holdings: Stay informed on the latest holdings disclosed by mutual funds and ETFs, including the number of shares held and the percentage of the portfolio they represent.

- Monitor Strategy Changes: Detect changes in fund strategy by reviewing updated disclosures, which may reveal shifts in investment focus or portfolio rebalancing.

- Gain Insight into Major Funds: Understand the investment decisions of significant institutional players, such as Vanguard or BlackRock, by accessing their most recent filings.

For example, an investor might use this API to track the latest disclosure from Vanguard's mutual fund, analyzing whether the fund increased or decreased its position in a particular stock, and use that information to support their own investment strategy.

## Response (example)

```json
[
	{
		"cik": "0000106444",
		"holder": "VANGUARD FIXED INCOME SECURITIES FUNDS",
		"shares": 67030000,
		"dateReported": "2024-07-31",
		"change": 0,
		"weightPercent": 0.03840197
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/latest-disclosures · 카테고리: etfAndMutualFunds
