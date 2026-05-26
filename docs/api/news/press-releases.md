# Press Releases

Access official company press releases with the FMP Press Releases API. Get real-time updates on corporate announcements, earnings reports, mergers, and more.

## Endpoint

`GET https://financialmodelingprep.com/stable/news/press-releases-latest?page=0&limit=20`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| from | date | 2026-01-27 |
| to | date | 2026-04-28 |
| page | number | 0 |
| limit | number | 20 |

## Description

The Press Releases API provides real-time access to official company announcements, allowing investors, analysts, and business professionals to stay informed on the latest developments. This API is crucial for:

- Company Announcements: Stay informed about earnings reports, product launches, mergers, and more directly from companies.

- Strategic Updates: Track leadership changes, business restructuring, and other significant corporate strategies that may affect a company's market standing.

- Market Impact Analysis: Analyze how company press releases influence stock prices, company valuations, and market sentiment.

This API ensures that you have access to the most current press releases, helping you make informed decisions based on the latest corporate disclosures.

Example Use Case
A financial analyst uses the Press Releases API to monitor corporate announcements from publicly traded companies, providing critical insights for investment decisions.

## Response (example)

```json
[
	{
		"symbol": "LNW",
		"publishedDate": "2025-02-03 23:32:00",
		"publisher": "PRNewsWire",
		"title": "Rosen Law Firm Encourages Light & Wonder, Inc. Investors to Inquire About Securities Class Action Investigation - LNW",
		"image": "https://images.financialmodelingprep.com/news/rosen-law-firm-encourages-light-wonder-inc-investors-to-20250203.jpg",
		"site": "prnewswire.com",
		"text": "NEW YORK , Feb. 3, 2025 /PRNewswire/ -- Why: Rosen Law Firm, a global investor rights law firm, continues to investigate potential securities claims on behalf of shareholders of Light & Wonder, Inc. (NASDAQ: LNW) resulting from allegations that Light & Wonder may have issued materially misleading business information to the investing public. So What: If you purchased Light & Wonder securities you may be entitled to compensation without payment of any out of pocket fees or costs through a contingency fee arrangement.",
		"url": "https://www.prnewswire.com/news-releases/rosen-law-firm-encourages-light--wonder-inc-investors-to-inquire-about-securities-class-action-investigation--lnw-302366877.html"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/press-releases · 카테고리: news
