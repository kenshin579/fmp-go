package statements

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// --- IncomeStatementGrowth ---

func TestIncomeStatementGrowth_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/income-statement-growth.json")
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

	rows, err := c.IncomeStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("IncomeStatementGrowth: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.GrowthRevenue == 0 {
		t.Error("GrowthRevenue not parsed (got 0)")
	}
	if r.GrowthNetIncome == 0 {
		t.Error("GrowthNetIncome not parsed (got 0)")
	}
}

func TestIncomeStatementGrowth_Delegation(t *testing.T) {
	raw, err := os.ReadFile("testdata/income-statement-growth.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var capturedPath string
	var capturedQuery map[string]string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		q := r.URL.Query()
		capturedQuery = make(map[string]string, len(q))
		for k, v := range q {
			if len(v) > 0 {
				capturedQuery[k] = v[0]
			}
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	_, err = c.IncomeStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("IncomeStatementGrowth: %v", err)
	}

	if capturedPath != "/stable/income-statement-growth" {
		t.Errorf("path = %q, want /stable/income-statement-growth", capturedPath)
	}
	if capturedQuery["symbol"] != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", capturedQuery["symbol"])
	}
	if capturedQuery["period"] != "annual" {
		t.Errorf("query period = %q, want annual", capturedQuery["period"])
	}
}

func TestIncomeStatementGrowth_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.IncomeStatementGrowth(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

func TestIncomeStatementGrowth_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.IncomeStatementGrowth(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

// --- BalanceSheetStatementGrowth ---

func TestBalanceSheetStatementGrowth_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/balance-sheet-statement-growth.json")
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

	rows, err := c.BalanceSheetStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("BalanceSheetStatementGrowth: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.GrowthTotalAssets == 0 {
		t.Error("GrowthTotalAssets not parsed (got 0)")
	}
}

func TestBalanceSheetStatementGrowth_Delegation(t *testing.T) {
	raw, err := os.ReadFile("testdata/balance-sheet-statement-growth.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var capturedPath string
	var capturedQuery map[string]string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		q := r.URL.Query()
		capturedQuery = make(map[string]string, len(q))
		for k, v := range q {
			if len(v) > 0 {
				capturedQuery[k] = v[0]
			}
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	_, err = c.BalanceSheetStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("BalanceSheetStatementGrowth: %v", err)
	}

	if capturedPath != "/stable/balance-sheet-statement-growth" {
		t.Errorf("path = %q, want /stable/balance-sheet-statement-growth", capturedPath)
	}
	if capturedQuery["symbol"] != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", capturedQuery["symbol"])
	}
}

func TestBalanceSheetStatementGrowth_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.BalanceSheetStatementGrowth(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

// --- CashFlowStatementGrowth ---

func TestCashFlowStatementGrowth_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/cash-flow-statement-growth.json")
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

	rows, err := c.CashFlowStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("CashFlowStatementGrowth: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.GrowthFreeCashFlow == 0 {
		t.Error("GrowthFreeCashFlow not parsed (got 0)")
	}
	// Verify the verbatim typo-tagged field parses correctly.
	if r.GrowthOtherInvestingActivites == 0 {
		t.Error("GrowthOtherInvestingActivites (typo tag) not parsed (got 0)")
	}
	if r.GrowthNetCashProvidedByOperatingActivites == 0 {
		t.Error("GrowthNetCashProvidedByOperatingActivites (typo tag) not parsed (got 0)")
	}
}

func TestCashFlowStatementGrowth_Delegation(t *testing.T) {
	raw, err := os.ReadFile("testdata/cash-flow-statement-growth.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var capturedPath string
	var capturedQuery map[string]string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		q := r.URL.Query()
		capturedQuery = make(map[string]string, len(q))
		for k, v := range q {
			if len(v) > 0 {
				capturedQuery[k] = v[0]
			}
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	_, err = c.CashFlowStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("CashFlowStatementGrowth: %v", err)
	}

	if capturedPath != "/stable/cash-flow-statement-growth" {
		t.Errorf("path = %q, want /stable/cash-flow-statement-growth", capturedPath)
	}
	if capturedQuery["symbol"] != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", capturedQuery["symbol"])
	}
}

func TestCashFlowStatementGrowth_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.CashFlowStatementGrowth(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

// --- FinancialStatementGrowth ---

func TestFinancialStatementGrowth_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/financial-growth.json")
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

	rows, err := c.FinancialStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("FinancialStatementGrowth: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("want 2 rows, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.RevenueGrowth == 0 {
		t.Error("RevenueGrowth not parsed (got 0)")
	}
}

func TestFinancialStatementGrowth_NullableFields(t *testing.T) {
	raw, err := os.ReadFile("testdata/financial-growth.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.FinancialStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("FinancialStatementGrowth: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("want 2 rows, got %d", len(rows))
	}

	// row[0]: all 5 nullable fields should be non-nil with expected values.
	r0 := rows[0]
	if r0.EBITDAGrowth == nil {
		t.Fatal("row[0].EBITDAGrowth: want non-nil, got nil")
	}
	if *r0.EBITDAGrowth != 0.05 {
		t.Errorf("row[0].EBITDAGrowth = %v, want 0.05", *r0.EBITDAGrowth)
	}
	if r0.GrowthCapitalExpenditure == nil {
		t.Error("row[0].GrowthCapitalExpenditure: want non-nil, got nil")
	}
	if r0.TenYBottomLineNetIncomeGrowthPerShare == nil {
		t.Error("row[0].TenYBottomLineNetIncomeGrowthPerShare: want non-nil, got nil")
	}
	if r0.FiveYBottomLineNetIncomeGrowthPerShare == nil {
		t.Error("row[0].FiveYBottomLineNetIncomeGrowthPerShare: want non-nil, got nil")
	}
	if r0.ThreeYBottomLineNetIncomeGrowthPerShare == nil {
		t.Error("row[0].ThreeYBottomLineNetIncomeGrowthPerShare: want non-nil, got nil")
	}

	// row[1]: all 5 nullable fields should be nil (JSON null).
	r1 := rows[1]
	if r1.EBITDAGrowth != nil {
		t.Errorf("row[1].EBITDAGrowth: want nil, got %v", *r1.EBITDAGrowth)
	}
	if r1.GrowthCapitalExpenditure != nil {
		t.Errorf("row[1].GrowthCapitalExpenditure: want nil, got %v", *r1.GrowthCapitalExpenditure)
	}
	if r1.TenYBottomLineNetIncomeGrowthPerShare != nil {
		t.Errorf("row[1].TenYBottomLineNetIncomeGrowthPerShare: want nil, got %v", *r1.TenYBottomLineNetIncomeGrowthPerShare)
	}
	if r1.FiveYBottomLineNetIncomeGrowthPerShare != nil {
		t.Errorf("row[1].FiveYBottomLineNetIncomeGrowthPerShare: want nil, got %v", *r1.FiveYBottomLineNetIncomeGrowthPerShare)
	}
	if r1.ThreeYBottomLineNetIncomeGrowthPerShare != nil {
		t.Errorf("row[1].ThreeYBottomLineNetIncomeGrowthPerShare: want nil, got %v", *r1.ThreeYBottomLineNetIncomeGrowthPerShare)
	}
}

func TestFinancialStatementGrowth_Delegation(t *testing.T) {
	raw, err := os.ReadFile("testdata/financial-growth.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var capturedPath string
	var capturedQuery map[string]string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		q := r.URL.Query()
		capturedQuery = make(map[string]string, len(q))
		for k, v := range q {
			if len(v) > 0 {
				capturedQuery[k] = v[0]
			}
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	_, err = c.FinancialStatementGrowth(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("FinancialStatementGrowth: %v", err)
	}

	if capturedPath != "/stable/financial-growth" {
		t.Errorf("path = %q, want /stable/financial-growth", capturedPath)
	}
	if capturedQuery["symbol"] != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", capturedQuery["symbol"])
	}
	if capturedQuery["period"] != "annual" {
		t.Errorf("query period = %q, want annual", capturedQuery["period"])
	}
}

func TestFinancialStatementGrowth_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.FinancialStatementGrowth(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

func TestFinancialStatementGrowth_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.FinancialStatementGrowth(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
