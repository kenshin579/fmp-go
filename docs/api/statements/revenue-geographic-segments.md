# Revenue Geographic Segments

Access detailed revenue breakdowns by geographic region with the Revenue Geographic Segments API. Analyze how different regions contribute to a company’s total revenue and identify key markets for growth.

## Endpoint

`GET https://financialmodelingprep.com/stable/revenue-geographic-segmentation?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| period | string | annual,quarter |
| structure | string | flat |

## Description

The Revenue Geographic Segments API allows users to retrieve revenue data segmented by geographical regions, helping investors and analysts understand the performance of a company in different markets. This API is ideal for:

- Regional Revenue Analysis: Break down revenue contributions by geographical area to see which regions are driving growth.

- Market Performance Insights: Analyze how a company is performing in key regions like the Americas, Europe, and Greater China.

- Global Strategy Planning: For businesses, understanding geographic revenue distribution can help in developing regional strategies and identifying new opportunities for expansion.

This API offers a granular view of regional revenue, making it easier to track a company's global financial performance.

Example Use Case
An investor can use the Revenue Geographic Segments API to track Apple's performance across key regions like the Americas, Europe, and Greater China, helping to identify emerging markets or regions with declining sales.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"fiscalYear": 2024,
		"period": "FY",
		"reportedCurrency": null,
		"date": "2024-09-28",
		"data": {
			"Americas Segment": 167045000000,
			"Europe Segment": 101328000000,
			"Greater China Segment": 66952000000,
			"Japan Segment": 25052000000,
			"Rest of Asia Pacific": 30658000000
		}
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/revenue-geographic-segments · 카테고리: statements
