package statements

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// --- IncomeStatementTTM ---

func TestIncomeStatementTTM_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/income-statement-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil {
		t.Fatalf("fixture is not a JSON array: %v", err)
	}
	if len(arr) == 0 {
		t.Fatal("fixture array empty")
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.IncomeStatementTTM(context.Background(), Params{Symbol: "AAPL"})
	if err != nil {
		t.Fatalf("IncomeStatementTTM: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.Revenue == 0 {
		t.Error("Revenue not parsed (got 0)")
	}
}

func TestIncomeStatementTTM_PeriodOmitted(t *testing.T) {
	raw, err := os.ReadFile("testdata/income-statement-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var capturedPath string
	var capturedQuery url.Values

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		capturedQuery = r.URL.Query()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	// Period is intentionally set to verify ttmQueryParams omits it.
	_, err = c.IncomeStatementTTM(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("IncomeStatementTTM: %v", err)
	}

	if capturedPath != "/stable/income-statement-ttm" {
		t.Errorf("path = %q, want /stable/income-statement-ttm", capturedPath)
	}
	if got := capturedQuery.Get("symbol"); got != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", got)
	}
	if _, ok := capturedQuery["period"]; ok {
		t.Error("query must NOT contain 'period' for TTM endpoint, but it was present")
	}
}

func TestIncomeStatementTTM_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.IncomeStatementTTM(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

// --- BalanceSheetStatementTTM ---

func TestBalanceSheetStatementTTM_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/balance-sheet-statement-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil {
		t.Fatalf("fixture is not a JSON array: %v", err)
	}
	if len(arr) == 0 {
		t.Fatal("fixture array empty")
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.BalanceSheetStatementTTM(context.Background(), Params{Symbol: "AAPL"})
	if err != nil {
		t.Fatalf("BalanceSheetStatementTTM: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.TotalAssets == 0 {
		t.Error("TotalAssets not parsed (got 0)")
	}
}

func TestBalanceSheetStatementTTM_PeriodOmitted(t *testing.T) {
	raw, err := os.ReadFile("testdata/balance-sheet-statement-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var capturedPath string
	var capturedQuery url.Values

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		capturedQuery = r.URL.Query()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	_, err = c.BalanceSheetStatementTTM(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("BalanceSheetStatementTTM: %v", err)
	}

	if capturedPath != "/stable/balance-sheet-statement-ttm" {
		t.Errorf("path = %q, want /stable/balance-sheet-statement-ttm", capturedPath)
	}
	if got := capturedQuery.Get("symbol"); got != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", got)
	}
	if _, ok := capturedQuery["period"]; ok {
		t.Error("query must NOT contain 'period' for TTM endpoint, but it was present")
	}
}

func TestBalanceSheetStatementTTM_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.BalanceSheetStatementTTM(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

// --- CashFlowStatementTTM ---

func TestCashFlowStatementTTM_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/cash-flow-statement-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil {
		t.Fatalf("fixture is not a JSON array: %v", err)
	}
	if len(arr) == 0 {
		t.Fatal("fixture array empty")
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.CashFlowStatementTTM(context.Background(), Params{Symbol: "AAPL"})
	if err != nil {
		t.Fatalf("CashFlowStatementTTM: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.FreeCashFlow == 0 {
		t.Error("FreeCashFlow not parsed (got 0)")
	}
}

func TestCashFlowStatementTTM_PeriodOmitted(t *testing.T) {
	raw, err := os.ReadFile("testdata/cash-flow-statement-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var capturedPath string
	var capturedQuery url.Values

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		capturedQuery = r.URL.Query()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	_, err = c.CashFlowStatementTTM(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("CashFlowStatementTTM: %v", err)
	}

	if capturedPath != "/stable/cash-flow-statement-ttm" {
		t.Errorf("path = %q, want /stable/cash-flow-statement-ttm", capturedPath)
	}
	if got := capturedQuery.Get("symbol"); got != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", got)
	}
	if _, ok := capturedQuery["period"]; ok {
		t.Error("query must NOT contain 'period' for TTM endpoint, but it was present")
	}
}

func TestCashFlowStatementTTM_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.CashFlowStatementTTM(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}
