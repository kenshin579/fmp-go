# COT Report

Access comprehensive Commitment of Traders (COT) reports with the FMP COT Report API. This API provides detailed information about long and short positions across various sectors, helping you assess market sentiment and track positions in commodities, indices, and financial instruments.

## Endpoint

`GET https://financialmodelingprep.com/stable/commitment-of-traders-report`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol | string | AAPL |
| from | date | 2024-01-01 |
| to | date | 2024-03-01 |

## Description

The FMP COT Report API is designed for traders, analysts, and market observers to evaluate the positions of market participants. This includes:

- Market Sentiment Tracking: Understand how commercial and non-commercial traders are positioned, giving you insights into the current sentiment of a specific market.

- Sector-Wide Analysis: Analyze trader positions across different sectors such as soft commodities, energy, and financials, offering a holistic view of market trends.

- Long and Short Positions: Get detailed data on long, short, and spread positions, helping you make informed decisions on market direction.

This API is perfect for anyone looking to gain a deeper understanding of market dynamics by observing how various market participants are positioned.

Example Use Case
A commodity trader can use the COT Report API to analyze the open interest and trader positions in the cocoa market, identifying trends in long and short positions to refine their trading strategy.

## Response (example)

```json
[
	{
		"symbol": "KC",
		"date": "2024-02-27 00:00:00",
		"name": "Coffee (KC)",
		"sector": "SOFTS",
		"marketAndExchangeNames": "COFFEE C - ICE FUTURES U.S.",
		"cftcContractMarketCode": "083731",
		"cftcMarketCode": "ICUS",
		"cftcRegionCode": "1",
		"cftcCommodityCode": "83",
		"openInterestAll": 209453,
		"noncommPositionsLongAll": 75330,
		"noncommPositionsShortAll": 23630,
		"noncommPositionsSpreadAll": 47072,
		"commPositionsLongAll": 79690,
		"commPositionsShortAll": 132114,
		"totReptPositionsLongAll": 202092,
		"totReptPositionsShortAll": 202816,
		"nonreptPositionsLongAll": 7361,
		"nonreptPositionsShortAll": 6637,
		"openInterestOld": 179986,
		"noncommPositionsLongOld": 75483,
		"noncommPositionsShortOld": 35395,
		"noncommPositionsSpreadOld": 27067,
		"commPositionsLongOld": 70693,
		"commPositionsShortOld": 111666,
		"totReptPositionsLongOld": 173243,
		"totReptPositionsShortOld": 174128,
		"nonreptPositionsLongOld": 6743,
		"nonreptPositionsShortOld": 5858,
		"openInterestOther": 29467,
		"noncommPositionsLongOther": 18754,
		"noncommPositionsShortOther": 7142,
		"noncommPositionsSpreadOther": 1098,
		"commPositionsLongOther": 8997,
		"commPositionsShortOther": 20448,
		"totReptPositionsLongOther": 28849,
		"totReptPositionsShortOther": 28688,
		"nonreptPositionsLongOther": 618,
		"nonreptPositionsShortOther": 779,
		"changeInOpenInterestAll": 2957,
		"changeInNoncommLongAll": -3545,
		"changeInNoncommShortAll": 618,
		"changeInNoncommSpeadAll": 1575,
		"changeInCommLongAll": 4978,
		"changeInCommShortAll": 802,
		"changeInTotReptLongAll": 3008,
		"changeInTotReptShortAll": 2995,
		"changeInNonreptLongAll": -51,
		"changeInNonreptShortAll": -38,
		"pctOfOpenInterestAll": 100,
		"pctOfOiNoncommLongAll": 36,
		"pctOfOiNoncommShortAll": 11.3,
		"pctOfOiNoncommSpreadAll": 22.5,
		"pctOfOiCommLongAll": 38,
		"pctOfOiCommShortAll": 63.1,
		"pctOfOiTotReptLongAll": 96.5,
		"pctOfOiTotReptShortAll": 96.8,
		"pctOfOiNonreptLongAll": 3.5,
		"pctOfOiNonreptShortAll": 3.2,
		"pctOfOpenInterestOl": 100,
		"pctOfOiNoncommLongOl": 41.9,
		"pctOfOiNoncommShortOl": 19.7,
		"pctOfOiNoncommSpreadOl": 15,
		"pctOfOiCommLongOl": 39.3,
		"pctOfOiCommShortOl": 62,
		"pctOfOiTotReptLongOl": 96.3,
		"pctOfOiTotReptShortOl": 96.7,
		"pctOfOiNonreptLongOl": 3.7,
		"pctOfOiNonreptShortOl": 3.3,
		"pctOfOpenInterestOther": 100,
		"pctOfOiNoncommLongOther": 63.6,
		"pctOfOiNoncommShortOther": 24.2,
		"pctOfOiNoncommSpreadOther": 3.7,
		"pctOfOiCommLongOther": 30.5,
		"pctOfOiCommShortOther": 69.4,
		"pctOfOiTotReptLongOther": 97.9,
		"pctOfOiTotReptShortOther": 97.4,
		"pctOfOiNonreptLongOther": 2.1,
		"pctOfOiNonreptShortOther": 2.6,
		"tradersTotAll": 357,
		"tradersNoncommLongAll": 132,
		"tradersNoncommShortAll": 77,
		"tradersNoncommSpreadAll": 94,
		"tradersCommLongAll": 106,
		"tradersCommShortAll": 119,
		"tradersTotReptLongAll": 286,
		"tradersTotReptShortAll": 250,
		"tradersTotOl": 351,
		"tradersNoncommLongOl": 136,
		"tradersNoncommShortOl": 72,
		"tradersNoncommSpeadOl": 88,
		"tradersCommLongOl": 94,
		"tradersCommShortOl": 114,
		"tradersTotReptLongOl": 269,
		"tradersTotReptShortOl": 239,
		"tradersTotOther": 164,
		"tradersNoncommLongOther": 31,
		"tradersNoncommShortOther": 34,
		"tradersNoncommSpreadOther": 16,
		"tradersCommLongOther": 59,
		"tradersCommShortOther": 68,
		"tradersTotReptLongOther": 102,
		"tradersTotReptShortOther": 106,
		"concGrossLe4TdrLongAll": 16,
		"concGrossLe4TdrShortAll": 23.7,
		"concGrossLe8TdrLongAll": 25.8,
		"concGrossLe8TdrShortAll": 38.9,
		"concNetLe4TdrLongAll": 9.8,
		"concNetLe4TdrShortAll": 16.2,
		"concNetLe8TdrLongAll": 17.7,
		"concNetLe8TdrShortAll": 25.4,
		"concGrossLe4TdrLongOl": 13.6,
		"concGrossLe4TdrShortOl": 24.7,
		"concGrossLe8TdrLongOl": 23.2,
		"concGrossLe8TdrShortOl": 40.3,
		"concNetLe4TdrLongOl": 11.3,
		"concNetLe4TdrShortOl": 18.2,
		"concNetLe8TdrLongOl": 20.3,
		"concNetLe8TdrShortOl": 31.9,
		"concGrossLe4TdrLongOther": 68.2,
		"concGrossLe4TdrShortOther": 29.1,
		"concGrossLe8TdrLongOther": 77.8,
		"concGrossLe8TdrShortOther": 47.3,
		"concNetLe4TdrLongOther": 64.7,
		"concNetLe4TdrShortOther": 26.7,
		"concNetLe8TdrLongOther": 73.9,
		"concNetLe8TdrShortOther": 44.2,
		"contractUnits": "(CONTRACTS OF 37,500 POUNDS)"
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cot-report · 카테고리: commitmentOfTraders
