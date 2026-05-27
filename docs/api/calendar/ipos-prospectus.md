# IPOs Prospectus

Access comprehensive information on IPO prospectuses with the FMP IPO Prospectus API. Get key financial details, such as public offering prices, discounts, commissions, proceeds before expenses, and more. This API also provides links to official SEC prospectuses, helping investors stay informed on companies entering the public market.

## Endpoint

`GET https://financialmodelingprep.com/stable/ipos-prospectus`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |

## Description

The FMP IPO Prospectus API offers detailed insights into IPO filings, providing essential information to investors, analysts, and regulatory professionals. With this API, users can access:

- Public Offering Prices: View the price per share and total amount raised through the IPO.

- Discounts and Commissions: Understand the fees and commissions deducted from the gross proceeds of the IPO.

- Proceeds Before Expenses: See the net proceeds the company expects to raise after expenses.

- Filing and IPO Dates: Track when companies file their prospectuses and their scheduled IPO dates.

- CIK and Form Type: Get key regulatory details, including the CIK number and the form type (e.g., 424B4).

- Direct SEC Links: Access the full IPO prospectus filed with the SEC for complete details on the offering.

This API is an invaluable tool for anyone looking to analyze IPO financial details before making investment decisions.

Example Use Case
An investment advisor can use the IPO Prospectus API to review a company's IPO financials and prospectus filings, helping them evaluate whether to recommend the IPO to clients based on the offering's structure.

## Response (example)

```json
[
	{
		"symbol": "ATAK",
		"acceptedDate": "2025-02-03",
		"filingDate": "2025-02-03",
		"ipoDate": "2022-03-20",
		"cik": "0001883788",
		"pricePublicPerShare": 0.78,
		"pricePublicTotal": 4649936.72,
		"discountsAndCommissionsPerShare": 0.04,
		"discountsAndCommissionsTotal": 254909.67,
		"proceedsBeforeExpensesPerShare": 0.74,
		"proceedsBeforeExpensesTotal": 4395207.05,
		"form": "424B4",
		"url": "https://www.sec.gov/Archives/edgar/data/1883788/000149315225004604/form424b4.htm"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/ipos-prospectus · 카테고리: calendar
