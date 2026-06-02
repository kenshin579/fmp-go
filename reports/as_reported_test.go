package reports

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
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

func TestIncomeStatementAsReported_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/income-statement-as-reported.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.IncomeStatementAsReported(context.Background(), "AAPL", "annual", 1)
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
	if r.FiscalYear != 2024 {
		t.Errorf("FiscalYear=%d, want 2024", r.FiscalYear)
	}
	if r.ReportedCurrency != nil {
		t.Errorf("ReportedCurrency=%v, want nil", r.ReportedCurrency)
	}

	gp, err := r.Data["grossprofit"].Int64()
	if err != nil {
		t.Errorf("grossprofit Int64 error: %v", err)
	}
	if gp <= 0 {
		t.Errorf("grossprofit=%d, want >0", gp)
	}

	eps, err := r.Data["earningspersharediluted"].Float64()
	if err != nil {
		t.Errorf("earningspersharediluted Float64 error: %v", err)
	}
	if eps <= 0 {
		t.Errorf("earningspersharediluted=%f, want >0", eps)
	}
}

func TestFinancialStatementFullAsReported_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-statement-full-as-reported.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.FinancialStatementFullAsReported(context.Background(), "AAPL", "annual", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	r := rows[0]

	v, ok := r.Data["documenttype"].(string)
	if !ok {
		t.Fatalf("documenttype not a string, got %T", r.Data["documenttype"])
	}
	if v != "10-K" {
		t.Errorf("documenttype=%q, want 10-K", v)
	}

	if _, present := r.Data["netincomeloss"]; !present {
		t.Error("netincomeloss key missing from Data")
	}
}

func TestIncomeStatementAsReported_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/income-statement-as-reported.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.IncomeStatementAsReported(context.Background(), "AAPL", "annual", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/income-statement-as-reported" {
		t.Errorf("path=%q, want /stable/income-statement-as-reported", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q, want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("period") != "annual" {
		t.Errorf("period=%q, want annual", cap.query.Get("period"))
	}
	if cap.query.Get("limit") != "1" {
		t.Errorf("limit=%q, want 1", cap.query.Get("limit"))
	}
}

func TestIncomeStatementAsReported_EmptySymbol(t *testing.T) {
	raw, _ := os.ReadFile("testdata/income-statement-as-reported.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	_, err := c.IncomeStatementAsReported(context.Background(), "", "annual", 1)
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}
