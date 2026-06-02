package marketperf

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSectorPESnapshot_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-pe-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.SectorPESnapshot(context.Background(), "2024-09-27", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d, want 1", len(rows))
	}
	if rows[0].Sector != "Technology" {
		t.Errorf("Sector=%q, want Technology", rows[0].Sector)
	}
	if rows[0].PE == 0 {
		t.Error("PE is 0")
	}
}

func TestSectorPESnapshot_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-pe-snapshot.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.SectorPESnapshot(context.Background(), "2024-09-27", "NASDAQ", "Technology")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sector-pe-snapshot" {
		t.Errorf("path=%q, want /stable/sector-pe-snapshot", cap.path)
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

func TestIndustryPESnapshot_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-pe-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.IndustryPESnapshot(context.Background(), "2024-09-27", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d, want 1", len(rows))
	}
	if rows[0].Industry != "Semiconductors" {
		t.Errorf("Industry=%q, want Semiconductors", rows[0].Industry)
	}
	if rows[0].PE == 0 {
		t.Error("PE is 0")
	}
}

func TestHistoricalSectorPE_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-pe-snapshot.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.HistoricalSectorPE(context.Background(), "Technology", "2024-01-01", "2024-09-27", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/historical-sector-pe" {
		t.Errorf("path=%q, want /stable/historical-sector-pe", cap.path)
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

func TestSectorPESnapshot_EmptyDateGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.SectorPESnapshot(context.Background(), "", "", "")
	if err == nil {
		t.Fatal("expected error for empty date, got nil")
	}
}
