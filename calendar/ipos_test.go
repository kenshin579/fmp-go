package calendar

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestIPOsCalendar_NullableFields(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ipos-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.IPOsCalendar(context.Background(), "2025-02-01", "2025-02-28")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Company == "" || rows[0].Exchange != "NYSE" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if rows[0].Shares != nil || rows[0].PriceRange != nil || rows[0].MarketCap != nil {
		t.Errorf("nullables should be nil: %+v", rows[0])
	}
	if cap.path != "/stable/ipos-calendar" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestIPODisclosures_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ipos-disclosure.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.IPODisclosures(context.Background(), "", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Form != "CERT" || rows[0].URL == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestIPOProspectuses_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ipos-prospectus.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.IPOProspectuses(context.Background(), "", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].PricePublicPerShare <= 0 || rows[0].ProceedsBeforeExpensesTotal <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
