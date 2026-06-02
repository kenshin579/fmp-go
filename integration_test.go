//go:build integration

package fmp_test

import (
	"context"
	"os"
	"testing"
	"time"

	fmp "github.com/kenshin579/fmp-go"
	"github.com/kenshin579/fmp-go/dcf"
	"github.com/kenshin579/fmp-go/insidertrades"
	"github.com/kenshin579/fmp-go/ratios"
	"github.com/kenshin579/fmp-go/search"
	"github.com/kenshin579/fmp-go/statements"
)

// 실행: FMP_API_KEY=... go test -tags integration -run TestIntegration ./...
func TestIntegration_CompanyProfile(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY not set")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	p, err := c.Company.Profile(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("Profile: %v", err)
	}
	if p.Symbol != "AAPL" || p.CompanyName == "" {
		t.Errorf("unexpected profile: %+v", p)
	}
}

func TestIntegration_IncomeStatement(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY not set")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	rows, err := c.Statements.IncomeStatement(context.Background(), statements.Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("IncomeStatement: %v", err)
	}
	if len(rows) == 0 || rows[0].Revenue <= 0 {
		t.Errorf("unexpected: %+v", rows)
	}
}

func TestIntegration_Ratios(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY not set")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	rows, err := c.Ratios.Ratios(context.Background(), ratios.Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("Ratios: %v", err)
	}
	if len(rows) == 0 {
		t.Error("empty ratios rows")
	}
}

func TestIntegration_BalanceSheetStatement(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY not set")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	rows, err := c.Statements.BalanceSheetStatement(context.Background(), statements.Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("BalanceSheetStatement: %v", err)
	}
	if len(rows) == 0 || rows[0].TotalAssets <= 0 {
		t.Errorf("unexpected: %+v", rows)
	}
}

func TestIntegration_Company(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if mc, err := c.Company.MarketCap(ctx, "AAPL"); err != nil || mc.MarketCap <= 0 {
		t.Errorf("MarketCap: err=%v mc=%+v", err, mc)
	}
	if rows, err := c.Company.StockPeers(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("StockPeers: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Company.KeyExecutives(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("KeyExecutives: err=%v len=%d", err, len(rows))
	}
	if sf, err := c.Company.SharesFloat(ctx, "AAPL"); err != nil {
		t.Errorf("SharesFloat: %v", err)
	} else {
		t.Logf("SharesFloat AAPL: %+v", sf) // 합성 struct 실 shape 확인용 로그
	}
	if rows, err := c.Company.DelistedCompanies(ctx, 0); err != nil || len(rows) == 0 {
		t.Errorf("DelistedCompanies: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_Search(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Search.SearchSymbol(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("SearchSymbol: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Search.SearchName(ctx, "Apple"); err != nil || len(rows) == 0 {
		t.Errorf("SearchName: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Search.CompanyScreener(ctx, search.ScreenerParams{Sector: "Technology", Limit: 5}); err != nil || len(rows) == 0 {
		t.Errorf("CompanyScreener: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_News(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.News.StockNewsLatest(ctx, 0); err != nil || len(rows) == 0 {
		t.Errorf("StockNewsLatest: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.News.SearchStockNews(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("SearchStockNews: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.News.FMPArticles(ctx, 0); err != nil || len(rows) == 0 {
		t.Errorf("FMPArticles: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_Analyst(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if g, err := c.Analyst.GradesConsensus(ctx, "AAPL"); err != nil || g.Symbol != "AAPL" {
		t.Errorf("GradesConsensus: err=%v g=%+v", err, g)
	}
	if pt, err := c.Analyst.PriceTargetConsensus(ctx, "AAPL"); err != nil || pt.TargetConsensus <= 0 {
		t.Errorf("PriceTargetConsensus: err=%v pt=%+v", err, pt)
	}
	if rows, err := c.Analyst.FinancialEstimates(ctx, "AAPL", "annual", 0); err != nil || len(rows) == 0 {
		t.Errorf("FinancialEstimates: err=%v len=%d", err, len(rows))
	} else {
		t.Logf("FinancialEstimate[0]: %+v", rows[0])
	}
}

func TestIntegration_Calendar(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Calendar.EarningsCalendar(ctx, "2025-02-01", "2025-02-28"); err != nil || len(rows) == 0 {
		t.Errorf("EarningsCalendar: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Calendar.CompanyDividends(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("CompanyDividends: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Calendar.SplitsCalendar(ctx, "2020-08-01", "2020-09-01"); err != nil {
		t.Errorf("SplitsCalendar: %v", err)
	} else {
		t.Logf("SplitsCalendar 건수: %d", len(rows))
	}
	if rows, err := c.Calendar.IPOsCalendar(ctx, "2025-02-01", "2025-02-28"); err != nil {
		t.Errorf("IPOsCalendar: %v", err)
	} else if len(rows) > 0 {
		t.Logf("IPO[0]: %+v", rows[0]) // PriceRange 실 타입 확인용 로그
	}
}

func TestIntegration_COT(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.COT.List(ctx); err != nil || len(rows) == 0 {
		t.Errorf("List: err=%v len=%d", err, len(rows))
	}
	if _, err := c.COT.Analysis(ctx, "ES", "", ""); err != nil {
		t.Errorf("Analysis: %v", err)
	}
	if _, err := c.COT.Report(ctx, "ES", "", ""); err != nil {
		t.Errorf("Report: %v", err)
	}
}

func TestIntegration_EarningsTranscripts(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.EarningsTranscripts.Latest(ctx, 0, 10); err != nil || len(rows) == 0 {
		t.Errorf("Latest: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.EarningsTranscripts.Dates(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("Dates: err=%v len=%d", err, len(rows))
	}
	if _, err := c.EarningsTranscripts.Transcript(ctx, "AAPL", "2023", "1", 0); err != nil {
		t.Errorf("Transcript: %v", err)
	}
}

func TestIntegration_ESG(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.ESG.Ratings(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("Ratings: err=%v len=%d", err, len(rows))
	}
	if _, err := c.ESG.Disclosures(ctx, "AAPL"); err != nil {
		t.Errorf("Disclosures: %v", err)
	}
	if rows, err := c.ESG.Benchmark(ctx, "2023"); err != nil || len(rows) == 0 {
		t.Errorf("Benchmark: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_Senate(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Senate.SenateLatest(ctx, 0, 10); err != nil || len(rows) == 0 {
		t.Errorf("SenateLatest: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Senate.HouseLatest(ctx, 0, 10); err != nil || len(rows) == 0 {
		t.Errorf("HouseLatest: err=%v len=%d", err, len(rows))
	}
	if _, err := c.Senate.SenateTrades(ctx, "AAPL", 0, 5); err != nil {
		t.Errorf("SenateTrades: %v", err)
	}
}

func TestIntegration_DCF(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.DCF.DiscountedCashFlow(ctx, "AAPL"); err != nil || len(rows) == 0 || rows[0].DCF <= 0 {
		t.Errorf("DiscountedCashFlow: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.DCF.CustomDiscountedCashFlow(ctx, dcf.CustomDCFParams{Symbol: "AAPL"}); err != nil || len(rows) == 0 {
		t.Errorf("CustomDiscountedCashFlow: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.DCF.CustomLeveredDiscountedCashFlow(ctx, dcf.CustomDCFParams{Symbol: "AAPL"}); err != nil || len(rows) == 0 {
		t.Errorf("CustomLeveredDiscountedCashFlow: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_TechnicalIndicators(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.TechnicalIndicators.SMA(ctx, "AAPL", 10, "1day", "", ""); err != nil || len(rows) == 0 || rows[0].Close <= 0 || rows[0].SMA <= 0 {
		t.Errorf("SMA: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.TechnicalIndicators.RSI(ctx, "AAPL", 14, "1day", "", ""); err != nil || len(rows) == 0 {
		t.Errorf("RSI: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.TechnicalIndicators.ADX(ctx, "AAPL", 14, "1day", "", ""); err != nil || len(rows) == 0 {
		t.Errorf("ADX: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_InsiderTrades(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.InsiderTrades.LatestInsiderTrades(ctx, "", 0, 5); err != nil || len(rows) == 0 {
		t.Errorf("LatestInsiderTrades: err=%v len=%d", err, len(rows))
	}
	if _, err := c.InsiderTrades.SearchInsiderTrades(ctx, insidertrades.SearchParams{Symbol: "AAPL", Limit: 5}); err != nil {
		t.Errorf("SearchInsiderTrades: %v", err)
	}
	if _, err := c.InsiderTrades.Statistics(ctx, "AAPL"); err != nil {
		t.Errorf("Statistics: %v", err)
	}
	if rows, err := c.InsiderTrades.TransactionTypes(ctx); err != nil || len(rows) == 0 {
		t.Errorf("TransactionTypes: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_MarketHours(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.MarketHours.AllExchangeMarketHours(ctx); err != nil || len(rows) == 0 {
		t.Errorf("AllExchangeMarketHours: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.MarketHours.ExchangeMarketHours(ctx, "NASDAQ"); err != nil || len(rows) == 0 || rows[0].Exchange == "" {
		t.Errorf("ExchangeMarketHours: err=%v len=%d", err, len(rows))
	}
	if _, err := c.MarketHours.HolidaysByExchange(ctx, "NASDAQ", "", ""); err != nil {
		t.Errorf("HolidaysByExchange: %v", err)
	}
}

func TestIntegration_Economics(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Economics.TreasuryRates(ctx, "", ""); err != nil || len(rows) == 0 {
		t.Errorf("TreasuryRates: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Economics.EconomicIndicators(ctx, "GDP", "", ""); err != nil || len(rows) == 0 {
		t.Errorf("EconomicIndicators: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Economics.MarketRiskPremium(ctx); err != nil || len(rows) == 0 {
		t.Errorf("MarketRiskPremium: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_Directory(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Directory.CompanySymbolsList(ctx); err != nil || len(rows) == 0 {
		t.Errorf("CompanySymbolsList: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Directory.AvailableExchanges(ctx); err != nil || len(rows) == 0 || rows[0].Exchange == "" {
		t.Errorf("AvailableExchanges: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Directory.AvailableSectors(ctx); err != nil || len(rows) == 0 {
		t.Errorf("AvailableSectors: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_MarketPerformance(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	date := time.Now().Format("2006-01-02")

	if rows, err := c.MarketPerformance.BiggestGainers(ctx); err != nil || len(rows) == 0 {
		t.Errorf("BiggestGainers: err=%v len=%d", err, len(rows))
	}
	if _, err := c.MarketPerformance.SectorPerformanceSnapshot(ctx, date, "", ""); err != nil {
		t.Errorf("SectorPerformanceSnapshot: %v", err) // 주말/휴일 빈 결과 허용
	}
	if _, err := c.MarketPerformance.SectorPESnapshot(ctx, date, "", ""); err != nil {
		t.Errorf("SectorPESnapshot: %v", err)
	}
}

func TestIntegration_Chart(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Chart.HistoricalPriceEODLight(ctx, "AAPL", "", ""); err != nil || len(rows) == 0 {
		t.Errorf("HistoricalPriceEODLight: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Chart.HistoricalPriceEODFull(ctx, "AAPL", "", ""); err != nil || len(rows) == 0 || rows[0].Close <= 0 {
		t.Errorf("HistoricalPriceEODFull: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Chart.Intraday1Hour(ctx, "AAPL", "", "", false); err != nil || len(rows) == 0 {
		t.Errorf("Intraday1Hour: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_Reports(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Reports.IncomeStatementAsReported(ctx, "AAPL", "annual", 1); err != nil || len(rows) == 0 || len(rows[0].Data) == 0 {
		t.Errorf("IncomeStatementAsReported: err=%v len=%d", err, len(rows))
	} else if _, ok := rows[0].Data["grossprofit"]; !ok {
		t.Errorf("IncomeStatementAsReported: grossprofit 키 없음: %v", rows[0].Data)
	}
	if rows, err := c.Reports.LatestFinancialStatements(ctx, 0, 5); err != nil || len(rows) == 0 {
		t.Errorf("LatestFinancialStatements: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Reports.FinancialReportDates(ctx, "AAPL"); err != nil || len(rows) == 0 || rows[0].LinkJson == "" {
		t.Errorf("FinancialReportDates: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Reports.FinancialReportJSON(ctx, "AAPL", 2022, "FY"); err != nil || len(rows) == 0 {
		t.Errorf("FinancialReportJSON: err=%v len=%d", err, len(rows))
	} else {
		keys := make([]string, 0, len(rows[0]))
		for k := range rows[0] {
			keys = append(keys, k)
		}
		t.Logf("FinancialReportJSON 섹션 키 일부: %v", keys)
	}
}

func TestIntegration_Metrics(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Metrics.KeyMetrics(ctx, "AAPL", "annual", 2); err != nil || len(rows) == 0 || rows[0].MarketCap == 0 {
		t.Errorf("KeyMetrics: err=%v len=%d", err, len(rows))
	} else {
		t.Logf("KeyMetrics[0]: %+v", rows[0]) // 절대값 타입 확인
	}
	if s, err := c.Metrics.FinancialScores(ctx, "AAPL"); err != nil || s.PiotroskiScore < 0 || s.PiotroskiScore > 9 {
		t.Errorf("FinancialScores: err=%v s=%+v", err, s)
	}
	if rows, err := c.Metrics.RevenueProductSegmentation(ctx, "AAPL", "annual"); err != nil || len(rows) == 0 || len(rows[0].Data) == 0 {
		t.Errorf("RevenueProductSegmentation: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Ratios.RatiosTTM(ctx, "AAPL"); err != nil || len(rows) == 0 {
		t.Errorf("RatiosTTM: err=%v len=%d", err, len(rows))
	}
}

func TestIntegration_StatementsCore(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if rows, err := c.Statements.CashFlowStatement(ctx, statements.Params{Symbol: "AAPL", Period: "annual", Limit: 2}); err != nil || len(rows) == 0 || rows[0].FreeCashFlow == 0 {
		t.Errorf("CashFlowStatement: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Statements.IncomeStatementTTM(ctx, statements.Params{Symbol: "AAPL"}); err != nil || len(rows) == 0 || rows[0].Revenue == 0 {
		t.Errorf("IncomeStatementTTM: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Statements.CashFlowStatementGrowth(ctx, statements.Params{Symbol: "AAPL", Period: "annual"}); err != nil || len(rows) == 0 {
		t.Errorf("CashFlowStatementGrowth: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Statements.FinancialStatementGrowth(ctx, statements.Params{Symbol: "AAPL", Period: "annual"}); err != nil || len(rows) == 0 {
		t.Errorf("FinancialStatementGrowth: err=%v len=%d", err, len(rows))
	} else {
		t.Logf("FinancialStatementGrowth[0]: %+v", rows[0]) // nullable 5필드 실제값 확인
	}
}

func TestIntegration_Quote(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY 미설정 — 통합 테스트 skip")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	ctx := context.Background()

	q, err := c.Quote.Quote(ctx, "AAPL")
	if err != nil {
		t.Fatalf("Quote: %v", err)
	}
	if q.Symbol != "AAPL" || q.Price <= 0 {
		t.Errorf("quote = %+v", q)
	}
	if _, err := c.Quote.QuoteShort(ctx, "AAPL"); err != nil {
		t.Errorf("QuoteShort: %v", err)
	}
	if _, err := c.Quote.PriceChange(ctx, "AAPL"); err != nil {
		t.Errorf("PriceChange: %v", err)
	}
	if rows, err := c.Quote.BatchQuote(ctx, "AAPL", "MSFT"); err != nil || len(rows) == 0 {
		t.Errorf("BatchQuote: err=%v len=%d", err, len(rows))
	}
	if rows, err := c.Quote.CryptoQuotes(ctx); err != nil || len(rows) == 0 {
		t.Errorf("CryptoQuotes: err=%v len=%d", err, len(rows))
	}
}
