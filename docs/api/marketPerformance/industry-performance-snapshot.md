# Industry Performance Snapshot

Access detailed performance data by industry using the Industry Performance Snapshot API. Analyze trends, movements, and daily performance metrics for specific industries across various stock exchanges.

## Endpoint

`GET https://financialmodelingprep.com/stable/industry-performance-snapshot?date=2024-02-01`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| date* | string | 2024-02-01 |
| exchange | string | NASDAQ |
| industry | string | Biotechnology |

## Description

The FMP Industry Performance Snapshot API provides a daily overview of how specific industries are performing across major stock exchanges. This API delivers key data, such as average percentage changes for industries like Advertising Agencies, Healthcare Equipment, or Technology Services, allowing users to track and compare performance trends within specific sectors. Key features include:

- Industry-Level Performance Data: View average percentage changes for specific industries across major exchanges.

- Real-Time Market Insights: Analyze industry performance trends and movements in real time with daily updates.

- Exchange-Specific Data: Compare how different industries are performing on various stock exchanges like NASDAQ, NYSE, and others.

- In-Depth Industry Comparisons: Track and analyze the performance of specific industries to understand market trends and identify growth opportunities.

This API is ideal for market analysts, portfolio managers, and investors who need to understand the performance dynamics of individual industries to guide investment strategies.

Example Use Case
A market analyst uses the Industry Performance Snapshot API to analyze the performance of the Advertising Agencies industry on a specific date, and finds that it posted an average gain of 3.87% on NASDAQ. This data helps the analyst recommend sector-specific investments and identify growth trends in the advertising sector.

## Response (example)

```json
[
	{
		"date": "2024-02-01",
		"industry": "Advertising Agencies",
		"exchange": "NASDAQ",
		"averageChange": 3.8660194344955996
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/industry-performance-snapshot · 카테고리: marketPerformance
