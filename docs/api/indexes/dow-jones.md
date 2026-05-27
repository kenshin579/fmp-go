# Dow Jones

Access data on the Dow Jones Industrial Average using the Dow Jones API. Track current values, analyze trends, and get detailed information about the companies that make up this important stock index.

## Endpoint

`GET https://financialmodelingprep.com/stable/dowjones-constituent`

## Description

The FMP Dow Jones Industrial Average API provides comprehensive information on the companies that are part of this iconic index. This API offers key details such as company name, symbol, sector, sub-sector, headquarters, and founding date, helping investors and analysts monitor the performance of one of the most widely followed stock market indexes. Key features include:

- Detailed Company Information: Access key details about Dow Jones constituents, including sector, sub-sector, and geographic location.

- Track Real-Time Trends: Follow current movements and trends in the Dow Jones Industrial Average.

- Sector Breakdown: Analyze how the index is divided across different sectors and sub-sectors for deeper insights.

- Historical Additions: See when companies were first added to the Dow Jones, providing context on index changes.

This API is ideal for financial professionals, portfolio managers, and analysts who need accurate and up-to-date information on the Dow Jones Industrial Average.

Example Use Case
A portfolio manager tracking the Dow Jones Industrial Average uses the Dow Jones API to access detailed data on newly added companies, like Amazon (AMZN). By analyzing the sector and sub-sector breakdown, the manager can evaluate the impact of changes in the index on their investment strategy.

## Response (example)

```json
[
	{
		"symbol": "NVDA",
		"name": "Nvidia",
		"sector": "Technology",
		"subSector": "Semiconductors",
		"headQuarter": "Santa Clara, CA",
		"dateFirstAdded": "2024-11-08",
		"cik": "0001045810",
		"founded": "1993-04-05"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/dow-jones · 카테고리: indexes
