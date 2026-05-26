# Stock Split Details

Access detailed information on stock splits for a specific company using the FMP Stock Split Details API. This API provides essential data, including the split date and the split ratio, helping users understand changes in a company's share structure after a stock split.

## Endpoint

`GET https://financialmodelingprep.com/stable/splits?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| limit | number | 100 |

## Description

The FMP Stock Split Details API is designed to offer critical insights into a company's stock split history. With this API, users can:

- Split Date Information: Access the exact date of a company's stock split to understand when the changes occurred.

- Split Ratio Details: Retrieve the split ratio, represented by the numerator and denominator, to see how many new shares are issued for every old share.

- Historical Reference: Track and analyze the impact of stock splits on a company's share price and market performance.

This API is ideal for investors and analysts who need to monitor stock split events and assess their effects on stock ownership and market trends.

Example Use Case
An investor looking to track Apple's stock split history can use the Stock Split Details API to retrieve detailed data on the company's past splits, including the date and ratio, allowing them to assess how splits have impacted stock value over time.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"date": "2020-08-31",
		"numerator": 4,
		"denominator": 1
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/splits-company · 카테고리: calendar
