package search

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSearchCIK_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-cik.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchCIK(context.Background(), "320193")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].CIK == "" || rows[0].Symbol != "AAPL" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/search-cik" || cap.query.Get("cik") != "320193" {
		t.Errorf("delegation: path=%q cik=%q", cap.path, cap.query.Get("cik"))
	}
	if _, err := c.SearchCIK(context.Background(), " "); err == nil {
		t.Fatal("want empty cik guard")
	}
}

func TestSearchCUSIP_ParsesFloatMarketCap(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-cusip.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchCUSIP(context.Background(), "037833100")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].CUSIP != "037833100" || rows[0].MarketCap <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestSearchISIN_ParsesIntMarketCap(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-isin.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchISIN(context.Background(), "US0378331005")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].ISIN != "US0378331005" || rows[0].MarketCap <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
