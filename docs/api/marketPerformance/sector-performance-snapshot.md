# Market Sector Performance Snapshot

Get a snapshot of sector performance using the Market Sector Performance Snapshot API. Analyze how different industries are performing in the market based on average changes across sectors.

## Endpoint

`GET https://financialmodelingprep.com/stable/sector-performance-snapshot?date=2024-02-01`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| date* | string | 2024-02-01 |
| exchange | string | NASDAQ |
| sector | string | Energy |

## Description

The FMP Market Sector Performance Snapshot API provides real-time insights into the performance of different sectors across various stock exchanges. This API allows users to track the average performance of industries like Basic Materials, Technology, Healthcare, and more, helping analysts and investors understand how different parts of the market are doing at any given moment. Key features include:

- Sector-Specific Performance Data: Access performance data for various sectors, including the average percentage change for each sector.

- Exchange-Specific Analysis: Analyze sector performance across specific exchanges such as NASDAQ, NYSE, and others.

- Daily Snapshots: Get daily updates on sector performance to track trends and market dynamics in real time.

- Cross-Industry Comparisons: Compare the performance of different sectors to identify growth or decline in key areas of the market.

This API is ideal for financial analysts, portfolio managers, and traders who need to track sector-level performance to make informed investment decisions.

Example Use Case
A portfolio manager uses the Market Sector Performance Snapshot API to review how different sectors performed on NASDAQ on a specific date. By identifying that the Basic Materials sector experienced an average decline of -0.31%, the manager can adjust their sector allocations and shift their focus to outperforming industries.

## Response (example)

```json
[
	{
		"date": "2024-02-01",
		"sector": "Basic Materials",
		"exchange": "NASDAQ",
		"averageChange": -0.31481377464310634
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/sector-performance-snapshot · 카테고리: marketPerformance
