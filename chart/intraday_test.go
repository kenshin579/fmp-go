package chart

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestIntraday1Min_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/intraday-1min.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Intraday1Min(context.Background(), "AAPL", "", "", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rows))
	}
	if rows[0].Close == 0 {
		t.Error("Close must not be zero")
	}
	if rows[0].Volume == 0 {
		t.Error("Volume must not be zero")
	}
	if rows[0].Open == 0 {
		t.Error("Open must not be zero")
	}
}

func TestIntraday1Min_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/intraday-1min.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Intraday1Min(context.Background(), "AAPL", "2024-09-27", "2024-09-27", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/historical-chart/1min" {
		t.Errorf("path=%q, want /stable/historical-chart/1min", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q, want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("from") != "2024-09-27" {
		t.Errorf("from=%q, want 2024-09-27", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2024-09-27" {
		t.Errorf("to=%q, want 2024-09-27", cap.query.Get("to"))
	}
	if cap.query.Get("nonadjusted") != "true" {
		t.Errorf("nonadjusted=%q, want true", cap.query.Get("nonadjusted"))
	}
}

func TestIntraday5Min_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/intraday-1min.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Intraday5Min(context.Background(), "AAPL", "", "", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/historical-chart/5min" {
		t.Errorf("path=%q, want /stable/historical-chart/5min", cap.path)
	}
	if _, ok := cap.query["nonadjusted"]; ok {
		t.Error("nonadjusted must not be present when nonadjusted=false")
	}
}

func TestIntraday1Min_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.Intraday1Min(context.Background(), "", "", "", false)
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}
