package form13f

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestExtractAnalyticsByHolder_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/analytics-holder.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.ExtractAnalyticsByHolder(context.Background(), "AAPL", "2023", "3", 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	r := rows[0]
	if r.InvestorName == "" {
		t.Errorf("InvestorName is empty: %+v", r)
	}
	if r.Weight == 0 {
		t.Errorf("Weight is 0: %+v", r)
	}
	if r.MarketValue == 0 {
		t.Errorf("MarketValue is 0: %+v", r)
	}
	// bool フィールド検証: fixture値と一致するか
	if r.IsNew != false {
		t.Errorf("IsNew expected false, got %v", r.IsNew)
	}
	if r.IsSoldOut != false {
		t.Errorf("IsSoldOut expected false, got %v", r.IsSoldOut)
	}
	if r.IsCountedForPerformance != true {
		t.Errorf("IsCountedForPerformance expected true, got %v", r.IsCountedForPerformance)
	}
}

func TestExtractAnalyticsByHolder_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/analytics-holder.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.ExtractAnalyticsByHolder(context.Background(), "AAPL", "2023", "3", 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/extract-analytics/holder" {
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
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "10" {
		t.Errorf("limit=%q", cap.query.Get("limit"))
	}
}

func TestExtractAnalyticsByHolder_EmptySymbolGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/analytics-holder.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.ExtractAnalyticsByHolder(context.Background(), "", "2023", "3", 0, 10)
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

func TestHoldersIndustryBreakdown_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-breakdown.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HoldersIndustryBreakdown(context.Background(), "0001067983", "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	r := rows[0]
	if r.IndustryTitle == "" {
		t.Errorf("IndustryTitle is empty: %+v", r)
	}
	if r.Weight == 0 {
		t.Errorf("Weight is 0: %+v", r)
	}
}

func TestHoldersIndustryBreakdown_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-breakdown.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.HoldersIndustryBreakdown(context.Background(), "0001067983", "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/holder-industry-breakdown" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("cik") != "0001067983" {
		t.Errorf("cik=%q", cap.query.Get("cik"))
	}
	if cap.query.Get("year") != "2023" {
		t.Errorf("year=%q", cap.query.Get("year"))
	}
	if cap.query.Get("quarter") != "3" {
		t.Errorf("quarter=%q", cap.query.Get("quarter"))
	}
}

func TestHoldersIndustryBreakdown_EmptyCikGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-breakdown.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.HoldersIndustryBreakdown(context.Background(), "", "2023", "3")
	if err == nil {
		t.Error("expected error for empty cik")
	}
}
