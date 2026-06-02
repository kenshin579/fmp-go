package dcf

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestCustomDiscountedCashFlow_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/custom-discounted-cash-flow.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.CustomDiscountedCashFlow(context.Background(), CustomDCFParams{Symbol: "AAPL"})
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Year == "" {
		t.Errorf("Year must not be empty")
	}
	if rows[0].WACC == 0 {
		t.Errorf("WACC must not be zero")
	}
	if rows[0].EquityValuePerShare == 0 {
		t.Errorf("EquityValuePerShare must not be zero")
	}
	if rows[0].CostOfDebt == 0 {
		t.Errorf("CostOfDebt must not be zero (costofDebt key mapping)")
	}
	if rows[0].UFCF == 0 {
		t.Errorf("UFCF must not be zero")
	}
}

func TestCustomLeveredDiscountedCashFlow_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/custom-levered-discounted-cash-flow.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.CustomLeveredDiscountedCashFlow(context.Background(), CustomDCFParams{Symbol: "AAPL"})
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].FreeCashFlow == 0 {
		t.Errorf("FreeCashFlow must not be zero")
	}
	if rows[0].OperatingCashFlow == 0 {
		t.Errorf("OperatingCashFlow must not be zero")
	}
	if rows[0].WACC == 0 {
		t.Errorf("WACC must not be zero")
	}
}

func TestCustomDiscountedCashFlow_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/custom-discounted-cash-flow.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	ptr := func(f float64) *float64 { return &f }
	_, err := c.CustomDiscountedCashFlow(context.Background(), CustomDCFParams{
		Symbol:  "AAPL",
		Beta:    ptr(1.2),
		TaxRate: ptr(0.21),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/custom-discounted-cash-flow" {
		t.Errorf("path=%q want /stable/custom-discounted-cash-flow", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("beta") != "1.2" {
		t.Errorf("beta=%q want 1.2", cap.query.Get("beta"))
	}
	if cap.query.Get("taxRate") != "0.21" {
		t.Errorf("taxRate=%q want 0.21", cap.query.Get("taxRate"))
	}
	if _, ok := cap.query["revenueGrowthPct"]; ok {
		t.Error("revenueGrowthPct key must be absent when not set")
	}
}

func TestCustomDiscountedCashFlow_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.CustomDiscountedCashFlow(context.Background(), CustomDCFParams{})
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}
