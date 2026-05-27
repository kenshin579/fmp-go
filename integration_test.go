//go:build integration

package fmp_test

import (
	"context"
	"os"
	"testing"

	fmp "github.com/kenshin579/fmp-go"
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
