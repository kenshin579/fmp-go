# Stock Splits Calendar

Stay informed about upcoming stock splits with the FMP Stock Splits Calendar API. This API provides essential data on upcoming stock splits across multiple companies, including the split date and ratio, helping you track changes in share structures before they occur.

## Endpoint

`GET https://financialmodelingprep.com/stable/splits-calendar`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-27 |
| page | number | 0 |

## Description

The FMP Stock Splits Calendar API offers timely information for investors and analysts who want to stay ahead of stock split events. This API provides:

- Upcoming Split Dates: Know when future stock splits are scheduled, allowing you to plan your investments around these events.

- Split Ratios: Access detailed split ratios, which show how many new shares (numerator) are issued for each old share (denominator).

- Market Insight: Use this data to evaluate how upcoming splits might impact stock prices, liquidity, and shareholder value.

This API helps users monitor stock split announcements across the market, ensuring they have the information needed to make informed investment decisions.

Example Use Case
A portfolio manager can use the Stock Splits Calendar API to stay updated on upcoming stock splits, such as a 1-for-100 split scheduled for GBK.ST on February 29, 2024, to adjust their strategies accordingly.

## Response (example)

```json
[
	{
		"symbol": "EYEN",
		"date": "2025-02-03",
		"numerator": 1,
		"denominator": 80
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/splits-calendar · 카테고리: calendar
