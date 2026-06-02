//go:build integration

package fmp_test

import (
	"context"
	"os"
	"testing"

	fmp "github.com/kenshin579/fmp-go"
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
