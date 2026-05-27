# Historical Market Sector Performance

Access historical sector performance data using the Historical Market Sector Performance API. Review how different sectors have performed over time across various stock exchanges.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-sector-performance?sector=Energy`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | string | 2024-02-01 |
| exchange | string | NASDAQ |
| sector* | string | Energy |
| to | string | 2024-03-01 |

## Description

The FMP Historical Market Sector Performance API provides detailed historical data on the performance of market sectors, such as Energy, Technology, Healthcare, and others. This API allows users to track and analyze sector-specific trends over time, helping identify long-term patterns and market movements. Key features include:

- Historical Sector Performance: Access historical data on average percentage changes in various sectors over time.

- Exchange-Specific Data: Track how sectors have performed on different stock exchanges, including NASDAQ, NYSE, and others.

- Long-Term Market Trends: Analyze trends and sector performance data over extended periods, offering insights for long-term investment strategies.

- Cross-Sector Analysis: Compare the performance of multiple sectors to see how different areas of the market have evolved.

This API is ideal for financial researchers, portfolio managers, and investors who need to review historical sector performance for trend analysis, sector rotation strategies, and long-term planning.

Example Use Case
An investor uses the Historical Market Sector Performance API to review the Energy sector's historical performance on NASDAQ. By analyzing data from a specific date, showing an average change of 0.64%, the investor can track the sector's performance over time and make more informed decisions about future investments in the Energy sector.

## Response (example)

```json
[
	{
		"date": "2024-02-01",
		"sector": "Energy",
		"exchange": "NASDAQ",
		"averageChange": 0.6397534025664513
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/historical-sector-performance · 카테고리: marketPerformance
