package directory

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestAvailableExchanges_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/available-exchanges.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.AvailableExchanges(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Exchange != "AMEX" {
		t.Errorf("Exchange=%q (want AMEX)", rows[0].Exchange)
	}
	if rows[0].CountryCode != "US" {
		t.Errorf("CountryCode=%q (want US)", rows[0].CountryCode)
	}
	if rows[0].Delay == "" {
		t.Errorf("Delay is empty")
	}
}

func TestAvailableExchanges_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/available-exchanges.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.AvailableExchanges(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/available-exchanges" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestAvailableSectors_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/available-sectors.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.AvailableSectors(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Sector != "Basic Materials" {
		t.Errorf("Sector=%q (want Basic Materials)", rows[0].Sector)
	}
	if cap.path != "/stable/available-sectors" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestAvailableIndustries_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/available-industries.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.AvailableIndustries(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Industry != "Steel" {
		t.Errorf("Industry=%q (want Steel)", rows[0].Industry)
	}
	if cap.path != "/stable/available-industries" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestAvailableCountries_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/available-countries.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.AvailableCountries(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Country != "US" {
		t.Errorf("Country=%q (want US)", rows[0].Country)
	}
	if cap.path != "/stable/available-countries" {
		t.Errorf("path=%q", cap.path)
	}
}
