package etf

import (
	"context"
	"net/http"
	"os"
	"testing"
)

// --- DisclosureHoldersSearch ---

func TestDisclosureHoldersSearch_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure-holders-search.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.DisclosureHoldersSearch(context.Background(), "Vanguard")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.CIK != "0000355691" {
		t.Errorf("CIK=%q want 0000355691", r.CIK)
	}
	if r.EntityOrgType != "30" {
		t.Errorf("EntityOrgType=%q want 30", r.EntityOrgType)
	}
}

func TestDisclosureHoldersSearch_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure-holders-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.DisclosureHoldersSearch(context.Background(), "Vanguard")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/funds/disclosure-holders-search" {
		t.Errorf("path=%q want /stable/funds/disclosure-holders-search", cap.path)
	}
	if cap.query.Get("name") != "Vanguard" {
		t.Errorf("name=%q want Vanguard", cap.query.Get("name"))
	}
}

func TestDisclosureHoldersSearch_EmptyName(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.DisclosureHoldersSearch(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty name")
	}
}

// --- DisclosureDates ---

func TestDisclosureDates_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure-dates.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.DisclosureDates(context.Background(), "VFIAX", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Year != 2024 {
		t.Errorf("Year=%d want 2024", r.Year)
	}
	if r.Quarter != 4 {
		t.Errorf("Quarter=%d want 4", r.Quarter)
	}
}

func TestDisclosureDates_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure-dates.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.DisclosureDates(context.Background(), "VFIAX", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/funds/disclosure-dates" {
		t.Errorf("path=%q want /stable/funds/disclosure-dates", cap.path)
	}
	if cap.query.Get("symbol") != "VFIAX" {
		t.Errorf("symbol=%q want VFIAX", cap.query.Get("symbol"))
	}
}

func TestDisclosureDates_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.DisclosureDates(context.Background(), "", "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// --- LatestDisclosureHolders ---

func TestLatestDisclosureHolders_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure-holders-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.LatestDisclosureHolders(context.Background(), "VFIAX")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.WeightPercent == 0 {
		t.Errorf("WeightPercent should not be zero")
	}
	if r.Holder == "" {
		t.Errorf("Holder should not be empty")
	}
}

func TestLatestDisclosureHolders_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure-holders-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.LatestDisclosureHolders(context.Background(), "VFIAX")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/funds/disclosure-holders-latest" {
		t.Errorf("path=%q want /stable/funds/disclosure-holders-latest", cap.path)
	}
}

func TestLatestDisclosureHolders_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.LatestDisclosureHolders(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// --- Disclosure ---

func TestDisclosure_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Disclosure(context.Background(), "AAPL", "2023", "4", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.IsRestrictedSec != "N" {
		t.Errorf("IsRestrictedSec=%q want N", r.IsRestrictedSec)
	}
	if r.CurCd != "USD" {
		t.Errorf("CurCd=%q want USD", r.CurCd)
	}
	if r.PctVal == 0 {
		t.Errorf("PctVal should not be zero")
	}
}

func TestDisclosure_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/disclosure.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Disclosure(context.Background(), "AAPL", "2023", "4", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/funds/disclosure" {
		t.Errorf("path=%q want /stable/funds/disclosure", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("year") != "2023" {
		t.Errorf("year=%q want 2023", cap.query.Get("year"))
	}
	if cap.query.Get("quarter") != "4" {
		t.Errorf("quarter=%q want 4", cap.query.Get("quarter"))
	}
}

func TestDisclosure_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.Disclosure(context.Background(), "", "2023", "4", "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}
