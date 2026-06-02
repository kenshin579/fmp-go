package analyst

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestRatingsSnapshot_ParsesNoDateField(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ratings-snapshot.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	r, err := c.RatingsSnapshot(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("RatingsSnapshot: %v", err)
	}
	if r.Rating != "A-" || r.OverallScore != 4 {
		t.Errorf("not parsed: %+v", r)
	}
	if r.Date != "" {
		t.Errorf("snapshot Date should be empty, got %q", r.Date)
	}
}

func TestHistoricalRatings_ParsesWithDate(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ratings-historical.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.HistoricalRatings(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Date == "" || rows[0].PriceToBookScore != 1 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/ratings-historical" {
		t.Errorf("path=%q", cap.path)
	}
}
