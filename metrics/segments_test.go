package metrics

import (
	"context"
	"net/http"
	"os"
	"testing"
)

// --- RevenueGeographicSegmentation ---

func TestRevenueGeographicSegmentation_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/revenue-geographic-segmentation.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.RevenueGeographicSegmentation(context.Background(), "AAPL", "annual")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d want 1", len(rows))
	}
	r := rows[0]
	if r.FiscalYear != 2024 {
		t.Errorf("FiscalYear=%d want 2024", r.FiscalYear)
	}
	if r.ReportedCurrency != nil {
		t.Errorf("ReportedCurrency=%v want nil", r.ReportedCurrency)
	}
	if len(r.Data) == 0 {
		t.Error("Data must not be empty")
	}
	if r.Data["Americas Segment"] != 167045000000 {
		t.Errorf("Americas Segment=%d want 167045000000", r.Data["Americas Segment"])
	}
}

func TestRevenueGeographicSegmentation_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/revenue-geographic-segmentation.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.RevenueGeographicSegmentation(context.Background(), "AAPL", "annual")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/revenue-geographic-segmentation" {
		t.Errorf("path=%q want /stable/revenue-geographic-segmentation", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("period") != "annual" {
		t.Errorf("period=%q want annual", cap.query.Get("period"))
	}
}

// --- RevenueProductSegmentation ---

func TestRevenueProductSegmentation_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/revenue-product-segmentation.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.RevenueProductSegmentation(context.Background(), "AAPL", "annual")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d want 1", len(rows))
	}
	r := rows[0]
	if r.Data["iPhone"] != 201183000000 {
		t.Errorf("iPhone=%d want 201183000000", r.Data["iPhone"])
	}
	if r.ReportedCurrency != nil {
		t.Errorf("ReportedCurrency=%v want nil", r.ReportedCurrency)
	}
}

func TestRevenueProductSegmentation_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/revenue-product-segmentation.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.RevenueProductSegmentation(context.Background(), "AAPL", "annual")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/revenue-product-segmentation" {
		t.Errorf("path=%q want /stable/revenue-product-segmentation", cap.path)
	}
}

// --- Empty symbol guard ---

func TestRevenueSegmentation_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.RevenueGeographicSegmentation(context.Background(), "", "annual")
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}
