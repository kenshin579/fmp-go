package form13f

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestHolderPerformanceSummary_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/holder-performance.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HolderPerformanceSummary(context.Background(), "0001067983", 0)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	if rows[0].InvestorName == "" {
		t.Errorf("InvestorName is empty: %+v", rows[0])
	}
	if rows[0].PortfolioSize == 0 {
		t.Errorf("PortfolioSize is 0: %+v", rows[0])
	}
	if rows[0].Turnover == 0 {
		t.Errorf("Turnover is 0: %+v", rows[0])
	}
	if rows[0].PerformancePercentage1year == 0 {
		t.Errorf("PerformancePercentage1year is 0: %+v", rows[0])
	}
}

func TestHolderPerformanceSummary_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/holder-performance.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.HolderPerformanceSummary(context.Background(), "0001067983", 0)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/holder-performance-summary" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("cik") != "0001067983" {
		t.Errorf("cik=%q", cap.query.Get("cik"))
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
}

func TestHolderPerformanceSummary_EmptyCikGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/holder-performance.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.HolderPerformanceSummary(context.Background(), "", 0)
	if err == nil {
		t.Error("expected error for empty cik")
	}
}

func TestPositionsSummary_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/positions-summary.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.PositionsSummary(context.Background(), "AAPL", "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	if rows[0].Symbol == "" {
		t.Errorf("Symbol is empty: %+v", rows[0])
	}
	if rows[0].InvestorsHolding == 0 {
		t.Errorf("InvestorsHolding is 0: %+v", rows[0])
	}
	if rows[0].PutCallRatio == 0 {
		t.Errorf("PutCallRatio is 0: %+v", rows[0])
	}
}

func TestPositionsSummary_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/positions-summary.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.PositionsSummary(context.Background(), "AAPL", "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/symbol-positions-summary" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
	if cap.query.Get("year") != "2023" {
		t.Errorf("year=%q", cap.query.Get("year"))
	}
	if cap.query.Get("quarter") != "3" {
		t.Errorf("quarter=%q", cap.query.Get("quarter"))
	}
}

func TestPositionsSummary_EmptySymbolGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/positions-summary.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.PositionsSummary(context.Background(), "", "2023", "3")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}
