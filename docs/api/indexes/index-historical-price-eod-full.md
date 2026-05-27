# Historical Index Full Chart

Access full historical end-of-day prices for stock indexes using the Detailed Historical Price Data API. This API provides comprehensive information, including open, high, low, close prices, volume, and additional metrics for detailed financial analysis.

## Endpoint

`GET https://financialmodelingprep.com/stable/historical-price-eod/full?symbol=^GSPC`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | ^GSPC |
| from | date | 2025-09-09 |
| to | date | 2025-12-09 |

## Description

The FMP Detailed Historical Price Data API offers full end-of-day price data for stock indexes, making it a powerful tool for in-depth financial analysis. It includes a range of price points&mdash;open, high, low, close&mdash;along with volume, price changes, and volume-weighted average price (VWAP). Key features include:

- Complete Price Data: Access open, high, low, and close prices for stock indexes on specific dates.

- Volume Information: Track trading volume to assess market activity and liquidity.

- Price Movement Insights: Analyze daily price changes and percentage changes to understand market trends.

- Volume-Weighted Average Price (VWAP): Get VWAP data for each trading day, helping in performance benchmarking and trading decisions.

This API is ideal for financial analysts, quants, and traders who need comprehensive historical price data to build models, conduct backtesting, or analyze market trends.

Example Use Case
A quantitative analyst developing an algorithmic trading model requires complete historical price data for the S&P 500 index (^GSPC). Using the Detailed Historical Price Data API, they can retrieve open, high, low, and close prices, along with VWAP and volume data for each trading day. This detailed information helps refine the model's predictions and backtesting performance.

## Response (example)

(문서에서 응답 예시를 찾지 못함)

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/index-historical-price-eod-full · 카테고리: indexes
