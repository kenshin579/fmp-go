# Nasdaq Index

Access comprehensive data for the Nasdaq index with the Nasdaq Index API. Monitor real-time movements and track the historical performance of companies listed on this prominent stock exchange.

## Endpoint

`GET https://financialmodelingprep.com/stable/nasdaq-constituent`

## Description

The FMP Nasdaq Index API provides up-to-date information on companies listed on the Nasdaq stock exchange. This API offers key details about each constituent, such as company name, symbol, sector, sub-sector, headquarters, and founding date. Whether you're tracking real-time movements or conducting historical analysis, this API is essential for those who need data on one of the world's largest stock exchanges. Key features include:

- Company Information: Access detailed data for Nasdaq-listed companies, including industry classification and headquarters location.

- Real-Time Monitoring: Track current and up-to-date information on Nasdaq constituents.

- Historical Insights: Analyze data about companies' founding dates and industry segments to understand long-term trends.

- Sector and Sub-Sector Breakdown: Evaluate the distribution of companies across various industries and sectors.

This API is a valuable resource for traders, portfolio managers, and analysts who need real-time insights and historical data on Nasdaq-listed companies.

Example Use Case
A financial analyst monitoring the technology sector uses the Nasdaq Index API to track the real-time performance of Nasdaq-listed companies, such as Apple Inc. (AAPL). By retrieving sector-specific data, the analyst can make informed decisions on market trends and identify potential investment opportunities in the tech industry.

## Response (example)

```json
[
	{
		"symbol": "ADBE",
		"name": "Adobe Inc.",
		"sector": "Technology",
		"subSector": "Software - Infrastructure",
		"headQuarter": "San Jose, CA",
		"dateFirstAdded": null,
		"cik": "0000796343",
		"founded": "1982-12-01"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/nasdaq · 카테고리: indexes
