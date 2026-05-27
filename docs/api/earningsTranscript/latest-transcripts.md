# Latest Earning Transcripts

Access available earnings transcripts for companies with the FMP Latest Earning Transcripts API. Retrieve a list of companies with earnings transcripts, along with the total number of transcripts available for each company.

## Endpoint

`GET https://financialmodelingprep.com/stable/earning-call-transcript-latest`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| limit | number | 100 |
| page | number | 0 |

## Description

The FMP Latest Earning Transcripts API provides users with essential data on the availability of earnings transcripts for various companies. This API is ideal for financial analysts, investors, and researchers looking to track earnings performance over time.

- Identify Available Transcripts: Quickly access a list of companies with earnings transcripts, complete with the number of available transcripts for each.

- Support Earnings Analysis: Use the transcript count to further analyze earnings call data and gain insights into company performance.

- Track Historical Data: Discover companies with multiple transcripts to track earnings calls over different quarters or years.

Example Use Case
An investor looking to analyze a company's earnings performance over several quarters can use the Earnings Transcript List API to identify companies with multiple earnings call transcripts and retrieve the necessary documents for deeper financial analysis.

## Response (example)

```json
[
	{
		"symbol": "CSWC",
		"period": "Q3",
		"fiscalYear": 2025,
		"date": "2025-02-04"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/latest-transcripts · 카테고리: earningsTranscript
