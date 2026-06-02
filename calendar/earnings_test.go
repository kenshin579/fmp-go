package calendar

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestEarningsCalendar_ParsesValues(t *testing.T) {
	raw, _ := os.ReadFile("testdata/earnings-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.EarningsCalendar(context.Background(), "2024-11-01", "2024-11-30")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].EpsActual == nil || *rows[0].EpsActual != 3.32 {
		t.Errorf("EpsActual: %v", rows[0].EpsActual)
	}
	if rows[0].RevenueActual == nil || *rows[0].RevenueActual <= 0 {
		t.Errorf("RevenueActual: %v", rows[0].RevenueActual)
	}
	if cap.path != "/stable/earnings-calendar" || cap.query.Get("from") != "2024-11-01" {
		t.Errorf("delegation: path=%q from=%q", cap.path, cap.query.Get("from"))
	}
}

func TestCompanyEarnings_NullableNilVsValue(t *testing.T) {
	raw, _ := os.ReadFile("testdata/earnings-company.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CompanyEarnings(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].EpsActual != nil || rows[0].RevenueActual != nil {
		t.Errorf("row0 nullables should be nil: %+v", rows[0])
	}
	if rows[1].EpsActual == nil || *rows[1].EpsActual != 2.4 {
		t.Errorf("row1 EpsActual should be set")
	}
	if rows[1].RevenueActual == nil || *rows[1].RevenueActual <= 0 {
		t.Errorf("row1 RevenueActual should be set")
	}
}
