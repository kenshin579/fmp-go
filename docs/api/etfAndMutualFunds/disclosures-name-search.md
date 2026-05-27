# Mutual Fund & ETF Disclosure Name Search

Easily search for mutual fund and ETF disclosures by name using the Mutual Fund & ETF Disclosure Name Search API. This API allows you to find specific reports and filings based on the fund or ETF name, providing essential details like CIK number, entity information, and reporting file number.

## Endpoint

`GET https://financialmodelingprep.com/stable/funds/disclosure-holders-search?name=Federated Hermes Government Income Securities, Inc.`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| name* | string | Federated Hermes Government Income Securities, Inc. |

## Description

The Mutual Fund & ETF Disclosure Name Search API helps users quickly locate disclosure documents for mutual funds and ETFs by searching with a specific fund name. It returns critical data such as the fund's symbol, CIK, class information, and the address of the reporting entity. Ideal for investors, analysts, and researchers looking for detailed disclosure information for compliance, research, or investment decision-making.

- Fund Name Search: Look up disclosures for mutual funds and ETFs using the fund or entity name.

- Key Filing Details: Get important information like CIK number, series and class IDs, entity name, and reporting file number.

- Comprehensive Results: The API returns address details and filing information for the searched fund or ETF entity, making it easy to locate relevant documents.

This API is perfect for anyone conducting due diligence or research on mutual funds and ETFs, allowing for precise and efficient disclosure searches.

Example Use Case
A financial analyst can use the Mutual Fund & ETF Disclosure Name Search API to retrieve specific disclosures for a mutual fund by entering its name, helping the analyst review relevant regulatory filings and reports for the fund.

## Response (example)

```json
[
	{
		"symbol": "FGOAX",
		"cik": "0000355691",
		"classId": "C000024574",
		"seriesId": "S000009042",
		"entityName": "Federated Hermes Government Income Securities, Inc.",
		"entityOrgType": "30",
		"seriesName": "Federated Hermes Government Income Securities, Inc.",
		"className": "Class A Shares",
		"reportingFileNumber": "811-03266",
		"address": "4000 ERICSSON DRIVE",
		"city": "WARRENDALE",
		"zipCode": "15086-7561",
		"state": "PA"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/disclosures-name-search · 카테고리: etfAndMutualFunds
