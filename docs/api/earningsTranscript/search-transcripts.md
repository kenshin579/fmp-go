# Earnings Transcript

Access the full transcript of a company’s earnings call with the FMP Earnings Transcript API. Stay informed about a company’s financial performance, future plans, and overall strategy by analyzing management's communication.

## Endpoint

`GET https://financialmodelingprep.com/stable/earning-call-transcript?symbol=AAPL&year=2020&quarter=3`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| year* | string | 2020 |
| quarter* | string | 3 |
| limit | number | 1 |

## Description

The FMP Earnings Transcript API provides complete access to the text transcript of a company's earnings call. This API is essential for:

- In-Depth Financial Analysis: Gain valuable insights into a company's financial performance by reviewing what executives say during earnings calls. The transcript can provide context and details beyond what's available in standard financial reports.

- Strategic Planning: Learn about a company's future plans and strategic direction straight from management. Understanding the company's priorities and challenges can help investors make informed decisions.

- Risk Identification: Use the transcript to identify any potential red flags or areas of concern that might not be immediately apparent in the earnings report. This can include management's tone, response to analysts' questions, or any mention of operational or financial difficulties.

Example Use Case
Investor Insight: An investor might use the Earnings Transcript API to review the most recent earnings call for a retail company. By analyzing the transcript, the investor can assess the company's response to market trends, management's outlook on upcoming quarters, and any potential risks that were discussed.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"period": "Q3",
		"year": 2020,
		"date": "2020-07-30",
		"content": "Operator: Good day, everyone. Welcome to the Apple Incorporated Third Quarter Fiscal Year 2020 Earnings Conference Call. Today's call is being recorded. At this time, for opening remarks and introductions, I would like to turn things over to Mr. Tejas Gala, Senior Manager, Corporate Finance and Investor Relations. Please go ahead, sir.\nTejas Gala: Thank you. Good afternoon and thank you for joining us. Speaking first today is Apple's CEO, Tim Cook; and he'll be followed by CFO, Luca Maestri. Aft..."
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-transcripts · 카테고리: earningsTranscript
