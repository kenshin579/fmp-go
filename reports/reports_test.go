package reports

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestLatestFinancialStatements_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/latest-financial-statements.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.LatestFinancialStatements(context.Background(), 0, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol=%q, want AAPL", r.Symbol)
	}
	if r.CalendarYear != 2024 {
		t.Errorf("CalendarYear=%d, want 2024", r.CalendarYear)
	}
	if r.DateAdded == "" {
		t.Error("DateAdded must not be empty")
	}
}

func TestLatestFinancialStatements_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/latest-financial-statements.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.LatestFinancialStatements(context.Background(), 0, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/latest-financial-statements" {
		t.Errorf("path=%q, want /stable/latest-financial-statements", cap.path)
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q, want 0", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "5" {
		t.Errorf("limit=%q, want 5", cap.query.Get("limit"))
	}
}

func TestFinancialReportDates_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-reports-dates.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.FinancialReportDates(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.LinkXlsx == "" {
		t.Error("LinkXlsx must not be empty")
	}
	if r.LinkJson == "" {
		t.Error("LinkJson must not be empty")
	}
	if r.FiscalYear != 2022 {
		t.Errorf("FiscalYear=%d, want 2022", r.FiscalYear)
	}
}

func TestFinancialReportDates_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-reports-dates.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.FinancialReportDates(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/financial-reports-dates" {
		t.Errorf("path=%q, want /stable/financial-reports-dates", cap.path)
	}
}

func TestFinancialReportJSON_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-reports-json.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.FinancialReportJSON(context.Background(), "AAPL", 2022, "FY")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	m := rows[0]
	if sym, ok := m["symbol"].(string); !ok || sym != "AAPL" {
		t.Errorf("symbol=%v, want AAPL", m["symbol"])
	}
	if yr, ok := m["year"].(string); !ok || yr != "2022" {
		t.Errorf("year=%v (type %T), want string \"2022\"", m["year"], m["year"])
	}
	if _, present := m["Cover Page"]; !present {
		t.Error("Cover Page key missing from result")
	}
}

func TestFinancialReportJSON_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-reports-json.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.FinancialReportJSON(context.Background(), "AAPL", 2022, "FY")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/financial-reports-json" {
		t.Errorf("path=%q, want /stable/financial-reports-json", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q, want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("year") != "2022" {
		t.Errorf("year=%q, want 2022", cap.query.Get("year"))
	}
	if cap.query.Get("period") != "FY" {
		t.Errorf("period=%q, want FY", cap.query.Get("period"))
	}
}

func TestFinancialReportJSON_EmptySymbol(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-reports-json.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	_, err := c.FinancialReportJSON(context.Background(), "", 2022, "FY")
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}
