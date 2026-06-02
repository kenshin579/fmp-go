package marketperf

import (
	"context"
	"net/http"
	"os"
	"testing"
)

// ---------- SectorPerformanceSnapshot ----------

func TestSectorPerformanceSnapshot_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-performance-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.SectorPerformanceSnapshot(context.Background(), "2024-09-27", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d, want 1", len(rows))
	}
	if rows[0].Sector != "Technology" {
		t.Errorf("Sector=%q, want Technology", rows[0].Sector)
	}
	if rows[0].AverageChange == 0 {
		t.Error("AverageChange is 0")
	}
}

func TestSectorPerformanceSnapshot_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-performance-snapshot.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.SectorPerformanceSnapshot(context.Background(), "2024-09-27", "NASDAQ", "Technology")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sector-performance-snapshot" {
		t.Errorf("path=%q, want /stable/sector-performance-snapshot", cap.path)
	}
	if cap.query.Get("date") != "2024-09-27" {
		t.Errorf("date=%q, want 2024-09-27", cap.query.Get("date"))
	}
	if cap.query.Get("sector") != "Technology" {
		t.Errorf("sector=%q, want Technology", cap.query.Get("sector"))
	}
	if cap.query.Get("exchange") != "NASDAQ" {
		t.Errorf("exchange=%q, want NASDAQ", cap.query.Get("exchange"))
	}
}

// ---------- IndustryPerformanceSnapshot ----------

func TestIndustryPerformanceSnapshot_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-performance-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.IndustryPerformanceSnapshot(context.Background(), "2024-09-27", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d, want 1", len(rows))
	}
	if rows[0].Industry != "Semiconductors" {
		t.Errorf("Industry=%q, want Semiconductors", rows[0].Industry)
	}
	if rows[0].AverageChange == 0 {
		t.Error("AverageChange is 0")
	}
}

// ---------- HistoricalSectorPerformance ----------

func TestHistoricalSectorPerformance_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-performance-snapshot.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.HistoricalSectorPerformance(context.Background(), "Technology", "2024-01-01", "2024-09-27", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/historical-sector-performance" {
		t.Errorf("path=%q, want /stable/historical-sector-performance", cap.path)
	}
	if cap.query.Get("sector") != "Technology" {
		t.Errorf("sector=%q, want Technology", cap.query.Get("sector"))
	}
	if cap.query.Get("from") != "2024-01-01" {
		t.Errorf("from=%q, want 2024-01-01", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2024-09-27" {
		t.Errorf("to=%q, want 2024-09-27", cap.query.Get("to"))
	}
}

// ---------- Guard tests ----------

func TestSectorPerformanceSnapshot_EmptyDate(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-performance-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	_, err := c.SectorPerformanceSnapshot(context.Background(), "", "", "")
	if err == nil {
		t.Fatal("expected error for empty date, got nil")
	}
}

func TestHistoricalSectorPerformance_EmptySector(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-performance-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	_, err := c.HistoricalSectorPerformance(context.Background(), "", "", "", "")
	if err == nil {
		t.Fatal("expected error for empty sector, got nil")
	}
}
