package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestEmployeeCount_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/employee-count-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.EmployeeCount(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].EmployeeCount <= 0 || rows[0].FormType == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestHistoricalEmployeeCount_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/historical-employee-count-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HistoricalEmployeeCount(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
}
