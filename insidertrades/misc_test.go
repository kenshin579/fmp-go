package insidertrades

import (
	"context"
	"fmt"
	"os"
	"testing"
)

// --- TransactionTypes ---

func TestTransactionTypes_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/transaction-type.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.TransactionTypes(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].TransactionType == "" {
		t.Errorf("TransactionType empty")
	}
	if cap.path != "/stable/insider-trading-transaction-type" {
		t.Errorf("path=%q", cap.path)
	}
}

// --- Statistics ---

func TestStatistics_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/statistics.json")
	c, _, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.Statistics(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.AverageAcquired == 0 {
		t.Errorf("AverageAcquired=0")
	}
	if r.CIK == "" {
		t.Errorf("CIK empty")
	}
	if r.Year == 0 {
		t.Errorf("Year=0")
	}
}

func TestStatistics_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/statistics.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Statistics(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/insider-trading/statistics" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
}

func TestStatistics_EmptySymbol(t *testing.T) {
	c, _, cleanup := newCapturingClient(t, "[]")
	defer cleanup()

	_, err := c.Statistics(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// --- AcquisitionOwnership ---

func TestAcquisitionOwnership_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/acquisition-ownership.json")
	c, _, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.AcquisitionOwnership(context.Background(), "AAPL", 5)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.PercentOfClass == "" {
		t.Errorf("PercentOfClass empty")
	}
	if r.NameOfReportingPerson == "" {
		t.Errorf("NameOfReportingPerson empty")
	}
}

func TestAcquisitionOwnership_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/acquisition-ownership.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.AcquisitionOwnership(context.Background(), "AAPL", 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/acquisition-of-beneficial-ownership" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
	if cap.query.Get("limit") != fmt.Sprintf("%d", 5) {
		t.Errorf("limit=%q", cap.query.Get("limit"))
	}
}

func TestAcquisitionOwnership_EmptySymbol(t *testing.T) {
	c, _, cleanup := newCapturingClient(t, "[]")
	defer cleanup()

	_, err := c.AcquisitionOwnership(context.Background(), "", 5)
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// --- SearchReportingName ---

func TestSearchReportingName_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/reporting-name.json")
	c, _, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.SearchReportingName(context.Background(), "Cook")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.ReportingName == "" {
		t.Errorf("ReportingName empty")
	}
	if r.ReportingCik == "" {
		t.Errorf("ReportingCik empty")
	}
}

func TestSearchReportingName_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/reporting-name.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.SearchReportingName(context.Background(), "Cook")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/insider-trading/reporting-name" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("name") != "Cook" {
		t.Errorf("name=%q", cap.query.Get("name"))
	}
}

func TestSearchReportingName_EmptyName(t *testing.T) {
	c, _, cleanup := newCapturingClient(t, "[]")
	defer cleanup()

	_, err := c.SearchReportingName(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty name")
	}
}
