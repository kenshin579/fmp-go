# Search Press Releases

Search for company press releases with the FMP Search Press Releases API. Find specific corporate announcements and updates by entering a stock symbol or company name.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/press-releases?symbols=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbols* | string | AAPL |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The Search Press Releases API allows users to find specific press releases based on a company name or stock symbol, offering quick access to relevant announcements. This API is essential for:

- Targeted Searches: Narrow down your search to find exact press releases from a particular company.

- Symbol-Based Retrieval: Use stock symbols to pinpoint corporate disclosures, making it ideal for investors and analysts looking for precise data.

- Historical and Real-Time Access: Retrieve both current and past press releases, helping with long-term trend analysis.

This API is designed for professionals who need quick, reliable access to specific press releases, saving time and providing accurate data.

Example Use Case
An investor uses the Search Press Releases API to find the most recent earnings report of a specific company before making an investment decision.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"publishedDate": "2025-01-30 16:30:00",
		"publisher": "Business Wire",
		"title": "Apple reports first quarter results",
		"image": "https://images.financialmodelingprep.com/news/apple-reports-first-quarter-results-20250130.jpg",
		"site": "businesswire.com",
		"text": "CUPERTINO, Calif.--(BUSINESS WIRE)--Apple® today announced financial results for its fiscal 2025 first quarter ended December 28, 2024. The Company posted quarterly revenue of $124.3 billion, up 4 percent year over year, and quarterly diluted earnings per share of $2.40, up 10 percent year over year. “Today Apple is reporting our best quarter ever, with revenue of $124.3 billion, up 4 percent from a year ago,” said Tim Cook, Apple's CEO. “We were thrilled to bring customers our best-ever lineup.",
		"url": "https://www.businesswire.com/news/home/20250130261281/en/Apple-reports-first-quarter-results/"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/search-press-releases · 카테고리: news
