package fmp_test

import (
	"testing"

	fmp "github.com/kenshin579/fmp-go"
)

func TestNewClient_EmptyKey(t *testing.T) {
	_, err := fmp.NewClient("")
	if err == nil {
		t.Fatal("expected error for empty apiKey")
	}
}

func TestNewClient_HasCompany(t *testing.T) {
	c, err := fmp.NewClient("k123")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Company == nil {
		t.Fatal("Company sub-client is nil")
	}
}

func TestNewClientFromEnv_MissingEnv(t *testing.T) {
	t.Setenv("FMP_API_KEY", "")
	if _, err := fmp.NewClientFromEnv(); err == nil {
		t.Fatal("expected error when FMP_API_KEY unset")
	}
}

func TestNewClientFromEnv_Reads(t *testing.T) {
	t.Setenv("FMP_API_KEY", "envkey")
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	if c.Company == nil {
		t.Fatal("Company nil")
	}
}

func TestNewClient_HasStatementsAndRatios(t *testing.T) {
	c, err := fmp.NewClient("k123")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Statements == nil {
		t.Fatal("Statements nil")
	}
	if c.Ratios == nil {
		t.Fatal("Ratios nil")
	}
}

func TestNewClient_HasQuote(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Quote == nil {
		t.Fatal("Quote sub-client is nil")
	}
}

func TestNewClient_HasSearch(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Search == nil {
		t.Fatal("Search sub-client is nil")
	}
}

func TestNewClient_HasNews(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.News == nil {
		t.Fatal("News sub-client is nil")
	}
}

func TestNewClient_HasAnalyst(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Analyst == nil {
		t.Fatal("Analyst sub-client is nil")
	}
}

func TestNewClient_HasCalendar(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Calendar == nil {
		t.Fatal("Calendar sub-client is nil")
	}
}

func TestNewClient_HasMetrics(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Metrics == nil {
		t.Fatal("Metrics sub-client is nil")
	}
}

func TestNewClient_HasReports(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Reports == nil {
		t.Fatal("Reports sub-client is nil")
	}
}

func TestNewClient_HasChart(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Chart == nil {
		t.Fatal("Chart sub-client is nil")
	}
}

func TestNewClient_HasMarketPerformance(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.MarketPerformance == nil {
		t.Fatal("MarketPerformance sub-client is nil")
	}
}

func TestNewClient_HasDirectory(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Directory == nil {
		t.Fatal("Directory sub-client is nil")
	}
}

func TestNewClient_HasEconomics(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Economics == nil {
		t.Fatal("Economics sub-client is nil")
	}
}

func TestNewClient_HasMarketHours(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.MarketHours == nil {
		t.Fatal("MarketHours sub-client is nil")
	}
}

func TestNewClient_HasInsiderTrades(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.InsiderTrades == nil {
		t.Fatal("InsiderTrades sub-client is nil")
	}
}

func TestNewClient_HasTechnicalIndicators(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.TechnicalIndicators == nil {
		t.Fatal("TechnicalIndicators sub-client is nil")
	}
}

func TestNewClient_HasDCF(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.DCF == nil {
		t.Fatal("DCF sub-client is nil")
	}
}

func TestNewClient_HasSenate(t *testing.T) {
	c, err := fmp.NewClient("k")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Senate == nil {
		t.Fatal("Senate sub-client is nil")
	}
}
