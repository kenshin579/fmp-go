package analyst

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestPriceTargetConsensus_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/price-target-consensus.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	pt, err := c.PriceTargetConsensus(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("PriceTargetConsensus: %v", err)
	}
	if pt.TargetHigh <= 0 || pt.TargetConsensus <= 0 || pt.TargetMedian <= 0 {
		t.Errorf("not parsed: %+v", pt)
	}
}

func TestPriceTargetSummary_ParsesPublishersString(t *testing.T) {
	raw, _ := os.ReadFile("testdata/price-target-summary.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	s, err := c.PriceTargetSummary(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("PriceTargetSummary: %v", err)
	}
	if s.AllTimeCount != 167 || s.LastYearAvgPriceTarget <= 0 {
		t.Errorf("not parsed: %+v", s)
	}
	if !strings.HasPrefix(s.Publishers, "[") || !strings.Contains(s.Publishers, "Benzinga") {
		t.Errorf("publishers should be JSON-array string: %q", s.Publishers)
	}
}
