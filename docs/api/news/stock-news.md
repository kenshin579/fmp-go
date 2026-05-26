# Stock News

Stay informed with the latest stock market news using the FMP Stock News Feed API. Access headlines, snippets, publication URLs, and ticker symbols for the most recent articles from a variety of sources.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/stock-latest?page=0&limit=20`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The Stock News API offers up-to-date information on stock market events, keeping traders, investors, and financial professionals informed about:

- Breaking Market News: Access the latest headlines that may impact stock prices and market movements.

- Company-Specific News: Stay updated on news related to individual stocks, including earnings reports, product announcements, and mergers.

- Market Trends and Analysis: Follow broader market trends and sentiment to make better investment decisions.

This API is designed to provide timely news that helps professionals track stock market developments and make informed decisions.

Example Use Case
A portfolio manager uses the Stock News API to track real-time updates on the stock markets, ensuring they are aware of any news that may affect the performance of the equities in their portfolio.

## Response (example)

```json
[
	{
		"symbol": "INSG",
		"publishedDate": "2025-02-03 23:53:40",
		"publisher": "Seeking Alpha",
		"title": "Q4 Earnings Release Looms For Inseego, But Don't Expect Miracles",
		"image": "https://images.financialmodelingprep.com/news/q4-earnings-release-looms-for-inseego-but-dont-expect-20250203.jpg",
		"site": "seekingalpha.com",
		"text": "Inseego's Q3 beat was largely due to a one-time debt restructuring gain, not sustainable earnings growth, raising concerns about future performance. The sale of its telematics business for $52 million allows INSG to focus on North America, but it remains to be seen if this was wise. Despite improved margins and reduced debt, Inseego's revenue growth is insufficient, and its high stock price remains unjustifiable for new investors.",
		"url": "https://seekingalpha.com/article/4754485-inseego-stock-q4-earnings-preview-monitor-growth-margins-closely"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/stock-news · 카테고리: news
