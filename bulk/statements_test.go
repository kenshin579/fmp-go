package bulk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func newTestClient(t *testing.T, status int, body string) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, srv.Close
}

type capturedReq struct {
	path  string
	query url.Values
}

func newCapturingClient(t *testing.T, body string) (*Client, *capturedReq, func()) {
	t.Helper()
	cap := &capturedReq{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cap.path = r.URL.Path
		cap.query = r.URL.Query()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, cap, srv.Close
}

const csvBody = "symbol,date,revenue\nAAPL,2024-09-28,391035000000\n"

func TestIncomeStatement_ParsesCSV(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, csvBody)
	defer cleanup()

	body, err := c.IncomeStatement(context.Background(), "2024", "FY")
	if err != nil {
		t.Fatalf("IncomeStatement: %v", err)
	}
	if len(body) == 0 {
		t.Fatal("expected non-empty body")
	}
	if !strings.Contains(string(body), "symbol") {
		t.Errorf("body does not contain 'symbol': %q", string(body))
	}
}

func TestIncomeStatement_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()

	_, err := c.IncomeStatement(context.Background(), "2024", "FY")
	if err != nil {
		t.Fatalf("IncomeStatement: %v", err)
	}
	if cap.path != "/stable/income-statement-bulk" {
		t.Errorf("path = %q, want /stable/income-statement-bulk", cap.path)
	}
	if cap.query.Get("year") != "2024" {
		t.Errorf("year = %q, want 2024", cap.query.Get("year"))
	}
	if cap.query.Get("period") != "FY" {
		t.Errorf("period = %q, want FY", cap.query.Get("period"))
	}
}

func TestIncomeStatement_EmptyYearGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, csvBody)
	defer cleanup()

	_, err := c.IncomeStatement(context.Background(), "", "FY")
	if err == nil {
		t.Fatal("expected error for empty year, got nil")
	}
}

func TestIncomeStatementGrowth_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.IncomeStatementGrowth(context.Background(), "2024", "FY")
	if err != nil {
		t.Fatalf("IncomeStatementGrowth: %v", err)
	}
	if cap.path != "/stable/income-statement-growth-bulk" {
		t.Errorf("path = %q", cap.path)
	}
}

func TestBalanceSheetStatement_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.BalanceSheetStatement(context.Background(), "2024", "FY")
	if err != nil {
		t.Fatalf("BalanceSheetStatement: %v", err)
	}
	if cap.path != "/stable/balance-sheet-statement-bulk" {
		t.Errorf("path = %q", cap.path)
	}
}

func TestBalanceSheetStatementGrowth_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.BalanceSheetStatementGrowth(context.Background(), "2024", "FY")
	if err != nil {
		t.Fatalf("BalanceSheetStatementGrowth: %v", err)
	}
	if cap.path != "/stable/balance-sheet-statement-growth-bulk" {
		t.Errorf("path = %q", cap.path)
	}
}

func TestCashFlowStatement_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.CashFlowStatement(context.Background(), "2024", "FY")
	if err != nil {
		t.Fatalf("CashFlowStatement: %v", err)
	}
	if cap.path != "/stable/cash-flow-statement-bulk" {
		t.Errorf("path = %q", cap.path)
	}
}

func TestCashFlowStatementGrowth_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.CashFlowStatementGrowth(context.Background(), "2024", "FY")
	if err != nil {
		t.Fatalf("CashFlowStatementGrowth: %v", err)
	}
	if cap.path != "/stable/cash-flow-statement-growth-bulk" {
		t.Errorf("path = %q", cap.path)
	}
}
