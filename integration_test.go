//go:build integration

package fmp_test

import (
	"context"
	"os"
	"testing"

	fmp "github.com/kenshin579/fmp-go"
	"github.com/kenshin579/fmp-go/ratios"
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
