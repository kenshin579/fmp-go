package analyst

import (
	"context"
	"os"
	"testing"
)

func TestFinancialEstimates_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/analyst-estimates.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.FinancialEstimates(context.Background(), "AAPL", "annual", 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].RevenueAvg <= 0 || rows[0].EpsAvg <= 0 || rows[0].NumAnalystsEps <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/analyst-estimates" || cap.query.Get("symbol") != "AAPL" || cap.query.Get("period") != "annual" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q symbol=%q period=%q page=%q", cap.path, cap.query.Get("symbol"), cap.query.Get("period"), cap.query.Get("page"))
	}
}

func TestFinancialEstimates_Guards(t *testing.T) {
	c, cleanup := newTestClient(t, 200, `[]`)
	defer cleanup()
	if _, err := c.FinancialEstimates(context.Background(), "  ", "annual", 0); err == nil {
		t.Error("want empty symbol guard")
	}
	if _, err := c.FinancialEstimates(context.Background(), "AAPL", "", 0); err == nil {
		t.Error("want empty period guard")
	}
}
