package calendar

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSplitsCalendar_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/splits-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SplitsCalendar(context.Background(), "2020-08-01", "2020-09-01")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Numerator != 4 || rows[0].Denominator != 1 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/splits-calendar" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("from") != "2020-08-01" || cap.query.Get("to") != "2020-09-01" {
		t.Errorf("query=%v", cap.query)
	}
}

func TestCompanySplits_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/splits.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.CompanySplits(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/splits" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("query=%v", cap.query)
	}
}

func TestCompanySplits_EmptySymbolGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.CompanySplits(context.Background(), "")
	if err == nil {
		t.Error("expected guard error for empty symbol")
	}
}
