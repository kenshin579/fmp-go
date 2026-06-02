# fmp-go

Financial Modeling Prep(FMP) API 의 Go 클라이언트 라이브러리.

## 설치

```bash
go get github.com/kenshin579/fmp-go@v0.1.0
```

## 사용

```go
client, _ := fmp.NewClientFromEnv() // FMP_API_KEY
ctx := context.Background()

profile, _ := client.Company.Profile(ctx, "AAPL") // 회사 프로필
fmt.Println(profile.CompanyName, profile.CEO, profile.Website)
```

## 인증

발급받은 API 키를 `FMP_API_KEY` 환경변수로 두거나 `fmp.NewClient(apiKey)` 로 전달한다.
모든 요청에 `apikey` 쿼리로 자동 주입된다. FMP stable 엔드포인트
(`https://financialmodelingprep.com/stable/...`)를 사용한다.

## 커버리지

| 카테고리 | 서비스 | 엔드포인트 |
|----------|--------|-----------|
| Analyst | `client.Analyst` | Grades, GradesConsensus, HistoricalGrades, RatingsSnapshot, HistoricalRatings, PriceTargetConsensus, PriceTargetSummary, FinancialEstimates — 8 endpoint |
| Company | `client.Company` | Profile, ProfileByCIK, MarketCap(+historical/batch), SharesFloat(+all), EmployeeCount(+historical), KeyExecutives, ExecutiveCompensation(+benchmark), StockPeers, CompanyNotes, Mergers(latest/search), DelistedCompanies — 17 endpoint |
| Statements | `client.Statements` | IncomeStatement, BalanceSheetStatement, CashFlowStatement, IncomeStatementTTM, BalanceSheetStatementTTM, CashFlowStatementTTM, IncomeStatementGrowth, BalanceSheetStatementGrowth, CashFlowStatementGrowth, FinancialStatementGrowth — 10 endpoint |
| Ratios | `client.Ratios` | Ratios, RatiosTTM — 2 endpoint |
| Metrics | `client.Metrics` | KeyMetrics, KeyMetricsTTM, FinancialScores, OwnerEarnings, EnterpriseValues, RevenueGeographicSegmentation, RevenueProductSegmentation — 7 endpoint |
| Reports | `client.Reports` | IncomeStatementAsReported, BalanceSheetStatementAsReported, CashFlowStatementAsReported, FinancialStatementFullAsReported, LatestFinancialStatements, FinancialReportDates, FinancialReportJSON — 7 endpoint |
| Chart | `client.Chart` | HistoricalPriceEODLight, HistoricalPriceEODFull, HistoricalPriceEODDividendAdjusted, HistoricalPriceEODNonSplitAdjusted, Intraday1Min, Intraday5Min, Intraday15Min, Intraday30Min, Intraday1Hour, Intraday4Hour — 10 endpoint |
| Market Performance | `client.MarketPerformance` | BiggestGainers, BiggestLosers, MostActives, SectorPerformanceSnapshot, IndustryPerformanceSnapshot, HistoricalSectorPerformance, HistoricalIndustryPerformance, SectorPESnapshot, IndustryPESnapshot, HistoricalSectorPE, HistoricalIndustryPE — 11 endpoint |
| Directory | `client.Directory` | CompanySymbolsList, FinancialSymbolsList, CIKList, SymbolChangesList, ETFsList, ActivelyTradingList, EarningsTranscriptList, AvailableExchanges, AvailableSectors, AvailableIndustries, AvailableCountries — 11 endpoint |
| Economics | `client.Economics` | TreasuryRates, EconomicIndicators, EconomicCalendar, MarketRiskPremium — 4 endpoint |
| Market Hours | `client.MarketHours` | ExchangeMarketHours, AllExchangeMarketHours, HolidaysByExchange — 3 endpoint |
| Insider Trades | `client.InsiderTrades` | LatestInsiderTrades, SearchInsiderTrades, TransactionTypes, Statistics, AcquisitionOwnership, SearchReportingName — 6 endpoint |
| Technical Indicators | `client.TechnicalIndicators` | SMA, EMA, WMA, DEMA, TEMA, RSI, StandardDeviation, Williams, ADX — 9 endpoint |
| DCF | `client.DCF` | DiscountedCashFlow, LeveredDiscountedCashFlow, CustomDiscountedCashFlow, CustomLeveredDiscountedCashFlow — 4 endpoint |
| Senate/House | `client.Senate` | SenateLatest, SenateTrades, SenateTradesByName, HouseLatest, HouseTrades, HouseTradesByName — 6 endpoint |
| ESG | `client.ESG` | Ratings, Disclosures, Benchmark — 3 endpoint |
| Earnings Transcripts | `client.EarningsTranscripts` | Transcript, Latest, Dates — 3 endpoint |
| Commitment of Traders | `client.COT` | Report, Analysis, List — 3 endpoint |
| Fundraisers | `client.Fundraisers` | LatestCrowdfunding, CrowdfundingByCIK, CrowdfundingSearch, LatestEquityOffering, EquityOfferingByCIK, EquityOfferingSearch — 6 endpoint |
| Quote | `client.Quote` | Quote, QuoteShort, PriceChange, AftermarketQuote/Trade, Batch(Quote/Short/Aftermarket), 자산군(Exchange/Index/Commodity/Crypto/ETF/Forex/MutualFund) — 16 endpoint |
| Search | `client.Search` | SearchSymbol, SearchName, SearchCIK, SearchCUSIP, SearchISIN, SearchExchangeVariants, CompanyScreener — 7 endpoint |
| News | `client.News` | StockNewsLatest, CryptoNewsLatest, ForexNewsLatest, GeneralNewsLatest, PressReleasesLatest, SearchStockNews, SearchCryptoNews, SearchForexNews, SearchPressReleases, FMPArticles — 10 endpoint |
| Calendar | `client.Calendar` | DividendsCalendar, CompanyDividends, EarningsCalendar, CompanyEarnings, IPOsCalendar, IPODisclosures, IPOProspectuses, SplitsCalendar, CompanySplits — 9 endpoint |

> 전체 FMP API 커버리지를 목표로 카테고리 단위로 점진 확장한다.
> 전체 API 문서 카탈로그: `docs/api/`.

## 개발

```bash
go build ./...
go vet ./...
go test ./...                          # 단위 테스트
go test -tags integration ./...        # 통합(FMP_API_KEY 필요)
```
