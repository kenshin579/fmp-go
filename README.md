# fmp-go

[![Go Reference](https://pkg.go.dev/badge/github.com/kenshin579/fmp-go.svg)](https://pkg.go.dev/github.com/kenshin579/fmp-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A Go client library for the [Financial Modeling Prep (FMP)](https://financialmodelingprep.com/) API.

## Installation

```bash
go get github.com/kenshin579/fmp-go@latest
```

Requires Go 1.25+.

## Usage

```go
client, _ := fmp.NewClientFromEnv() // reads FMP_API_KEY
ctx := context.Background()

profile, _ := client.Company.Profile(ctx, "AAPL") // company profile
fmt.Println(profile.CompanyName, profile.CEO, profile.Website)
```

Every service method takes a `context.Context` as its first argument, so
requests honor cancellation and deadlines (e.g. `context.WithTimeout`).

See the [`examples/`](examples/) directory for runnable examples of each category.

## Authentication

Provide your API key either through the `FMP_API_KEY` environment variable
(`fmp.NewClientFromEnv()`) or by passing it directly (`fmp.NewClient(apiKey)`).
The key is automatically injected as the `apikey` query parameter on every
request. The library targets the FMP **stable** endpoints
(`https://financialmodelingprep.com/stable/...`).

## Configuration

`NewClient` and `NewClientFromEnv` accept functional options:

```go
client, _ := fmp.NewClient(apiKey,
    fmp.WithTimeout(10*time.Second),  // HTTP timeout (default: 30s)
    fmp.WithBaseURL("https://..."),   // override base URL (testing/proxy)
    fmp.WithHTTPClient(customClient), // inject a custom *http.Client
)
```

## Error Handling

All service methods return standard Go errors:

- When a lookup yields no result (e.g. an empty array), the method returns
  `fmp.ErrNotFound` — check it with `errors.Is`.
- A non-200 HTTP response or an FMP `"Error Message"` body is returned as
  `*fmp.APIError`, exposing `StatusCode` and `Message` via `errors.As`.

```go
profile, err := client.Company.Profile(ctx, "AAPL")
if errors.Is(err, fmp.ErrNotFound) {
    // no data for this symbol
}

var apiErr *fmp.APIError
if errors.As(err, &apiErr) {
    log.Printf("FMP error: status=%d message=%s", apiErr.StatusCode, apiErr.Message)
}
```

## Coverage

**222 endpoints across 28 categories** (~84% of the 263 documented FMP stable
endpoints), expanding incrementally by category.

| Category | Service | Endpoints |
|----------|---------|-----------|
| Analyst | `client.Analyst` | Grades, GradesConsensus, HistoricalGrades, RatingsSnapshot, HistoricalRatings, PriceTargetConsensus, PriceTargetSummary, FinancialEstimates — 8 endpoints |
| Company | `client.Company` | Profile, ProfileByCIK, MarketCap(+historical/batch), SharesFloat(+all), EmployeeCount(+historical), KeyExecutives, ExecutiveCompensation(+benchmark), StockPeers, CompanyNotes, Mergers(latest/search), DelistedCompanies — 17 endpoints |
| Statements | `client.Statements` | IncomeStatement, BalanceSheetStatement, CashFlowStatement, IncomeStatementTTM, BalanceSheetStatementTTM, CashFlowStatementTTM, IncomeStatementGrowth, BalanceSheetStatementGrowth, CashFlowStatementGrowth, FinancialStatementGrowth — 10 endpoints |
| Ratios | `client.Ratios` | Ratios, RatiosTTM — 2 endpoints |
| Metrics | `client.Metrics` | KeyMetrics, KeyMetricsTTM, FinancialScores, OwnerEarnings, EnterpriseValues, RevenueGeographicSegmentation, RevenueProductSegmentation — 7 endpoints |
| Reports | `client.Reports` | IncomeStatementAsReported, BalanceSheetStatementAsReported, CashFlowStatementAsReported, FinancialStatementFullAsReported, LatestFinancialStatements, FinancialReportDates, FinancialReportJSON — 7 endpoints |
| Chart | `client.Chart` | HistoricalPriceEODLight, HistoricalPriceEODFull, HistoricalPriceEODDividendAdjusted, HistoricalPriceEODNonSplitAdjusted, Intraday1Min, Intraday5Min, Intraday15Min, Intraday30Min, Intraday1Hour, Intraday4Hour — 10 endpoints |
| Market Performance | `client.MarketPerformance` | BiggestGainers, BiggestLosers, MostActives, SectorPerformanceSnapshot, IndustryPerformanceSnapshot, HistoricalSectorPerformance, HistoricalIndustryPerformance, SectorPESnapshot, IndustryPESnapshot, HistoricalSectorPE, HistoricalIndustryPE — 11 endpoints |
| Directory | `client.Directory` | CompanySymbolsList, FinancialSymbolsList, CIKList, SymbolChangesList, ETFsList, ActivelyTradingList, EarningsTranscriptList, AvailableExchanges, AvailableSectors, AvailableIndustries, AvailableCountries — 11 endpoints |
| Economics | `client.Economics` | TreasuryRates, EconomicIndicators, EconomicCalendar, MarketRiskPremium — 4 endpoints |
| Market Hours | `client.MarketHours` | ExchangeMarketHours, AllExchangeMarketHours, HolidaysByExchange — 3 endpoints |
| Insider Trades | `client.InsiderTrades` | LatestInsiderTrades, SearchInsiderTrades, TransactionTypes, Statistics, AcquisitionOwnership, SearchReportingName — 6 endpoints |
| Technical Indicators | `client.TechnicalIndicators` | SMA, EMA, WMA, DEMA, TEMA, RSI, StandardDeviation, Williams, ADX — 9 endpoints |
| DCF | `client.DCF` | DiscountedCashFlow, LeveredDiscountedCashFlow, CustomDiscountedCashFlow, CustomLeveredDiscountedCashFlow — 4 endpoints |
| Senate/House | `client.Senate` | SenateLatest, SenateTrades, SenateTradesByName, HouseLatest, HouseTrades, HouseTradesByName — 6 endpoints |
| ESG | `client.ESG` | Ratings, Disclosures, Benchmark — 3 endpoints |
| Earnings Transcripts | `client.EarningsTranscripts` | Transcript, Latest, Dates — 3 endpoints |
| Commitment of Traders | `client.COT` | Report, Analysis, List — 3 endpoints |
| Fundraisers | `client.Fundraisers` | LatestCrowdfunding, CrowdfundingByCIK, CrowdfundingSearch, LatestEquityOffering, EquityOfferingByCIK, EquityOfferingSearch — 6 endpoints |
| Form 13F | `client.Form13F` | LatestFilings, Extract, FilingDates, ExtractAnalyticsByHolder, HolderPerformanceSummary, HoldersIndustryBreakdown, PositionsSummary, IndustrySummary — 8 endpoints |
| ETF & Mutual Funds | `client.ETF` | Holdings, Information, CountryWeightings, SectorWeightings, AssetExposure, DisclosureHoldersSearch, DisclosureDates, LatestDisclosureHolders, Disclosure — 9 endpoints |
| SEC Filings | `client.SECFilings` | LatestFinancials, Latest8K, SearchBySymbol, SearchByCIK, SearchByFormType, SearchByName, CompanySearchBySymbol, CompanySearchByCIK, Profile, IndustryClassificationList, IndustryClassificationSearch, AllIndustryClassification — 12 endpoints |
| Assets | `client.Assets` | CryptoList, ForexList, CommodityList — 3 endpoints (for crypto/forex/commodity quotes and time series use `client.Quote`/`client.Chart`) |
| Bulk | `client.Bulk` | Profile, ETFHolder, EOD, IncomeStatement(+Growth), BalanceSheetStatement(+Growth), CashFlowStatement(+Growth), EarningsSurprises, RatiosTTM, KeyMetricsTTM, Scores, DCF, Peers, PriceTargetSummary, Rating, UpgradesDowngradesConsensus — 18 endpoints (returns raw CSV `[]byte`) |
| Quote | `client.Quote` | Quote, QuoteShort, PriceChange, AftermarketQuote/Trade, Batch(Quote/Short/Aftermarket), asset classes (Exchange/Index/Commodity/Crypto/ETF/Forex/MutualFund) — 16 endpoints |
| Search | `client.Search` | SearchSymbol, SearchName, SearchCIK, SearchCUSIP, SearchISIN, SearchExchangeVariants, CompanyScreener — 7 endpoints |
| News | `client.News` | StockNewsLatest, CryptoNewsLatest, ForexNewsLatest, GeneralNewsLatest, PressReleasesLatest, SearchStockNews, SearchCryptoNews, SearchForexNews, SearchPressReleases, FMPArticles — 10 endpoints |
| Calendar | `client.Calendar` | DividendsCalendar, CompanyDividends, EarningsCalendar, CompanyEarnings, IPOsCalendar, IPODisclosures, IPOProspectuses, SplitsCalendar, CompanySplits — 9 endpoints |

> The complete API documentation catalog lives in [`docs/api/`](docs/api/).

## Development

```bash
go build ./...
go vet ./...
go test ./...                          # unit tests
go test -tags integration ./...        # integration tests (requires FMP_API_KEY)
```

## License

Released under the [MIT License](LICENSE).
