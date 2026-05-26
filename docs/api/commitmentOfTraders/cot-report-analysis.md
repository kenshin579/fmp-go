# COT Analysis By Dates

Gain in-depth insights into market sentiment with the FMP COT Report Analysis API. Analyze the Commitment of Traders (COT) reports for a specific date range to evaluate market dynamics, sentiment, and potential reversals across various sectors.

## Endpoint

`GET https://financialmodelingprep.com/stable/commitment-of-traders-analysis`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol | string | AAPL |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |

## Description

The FMP COT Report Analysis API is designed for traders, analysts, and market strategists to interpret the long and short positions of traders over time, helping to track sentiment trends and potential market shifts. This API includes:

- Market Sentiment Evaluation: Analyze the bullish or bearish sentiment based on long and short positions, helping you gauge the current market situation.

- Net Position Changes: Track changes in net positions to understand whether sentiment is becoming more bullish or bearish.

- Historical Sentiment Comparison: Compare current market sentiment with previous periods to detect trends or potential reversal points in the market.

This API enables market participants to make informed decisions by providing detailed insights into how traders are positioned in various markets and how sentiment evolves over time.

Example Use Case
A commodity trader can use the COT Report Analysis API to assess the bullish sentiment in the energy market by tracking changes in the net position of Brent crude oil traders, allowing them to refine their trading strategy accordingly.

## Response (example)

```json
[
	{
		"symbol": "B6",
		"date": "2024-02-27 00:00:00",
		"name": "British Pound (B6)",
		"sector": "CURRENCIES",
		"exchange": "BRITISH POUND - CHICAGO MERCANTILE EXCHANGE",
		"currentLongMarketSituation": 66.85,
		"currentShortMarketSituation": 33.15,
		"marketSituation": "Bullish",
		"previousLongMarketSituation": 67.97,
		"previousShortMarketSituation": 32.03,
		"previousMarketSituation": "Bullish",
		"netPostion": 46358,
		"previousNetPosition": 46312,
		"changeInNetPosition": 0.1,
		"marketSentiment": "Increasing Bullish",
		"reversalTrend": false
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cot-report-analysis · 카테고리: commitmentOfTraders
