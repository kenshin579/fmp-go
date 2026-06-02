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

func TestCashFlowStatement_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/cash-flow-statement.json")
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

	rows, err := c.CashFlowStatement(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("CashFlowStatement: %v", err)
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
	if r.NetIncome == 0 {
		t.Error("NetIncome not parsed (got 0)")
	}
}

func TestCashFlowStatement_Delegation(t *testing.T) {
	raw, err := os.ReadFile("testdata/cash-flow-statement.json")
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
	_, err = c.CashFlowStatement(context.Background(), Params{Symbol: "AAPL", Period: "annual"})
	if err != nil {
		t.Fatalf("CashFlowStatement: %v", err)
	}

	if capturedPath != "/stable/cash-flow-statement" {
		t.Errorf("path = %q, want /stable/cash-flow-statement", capturedPath)
	}
	if capturedQuery["symbol"] != "AAPL" {
		t.Errorf("query symbol = %q, want AAPL", capturedQuery["symbol"])
	}
	if capturedQuery["period"] != "annual" {
		t.Errorf("query period = %q, want annual", capturedQuery["period"])
	}
}

func TestCashFlowStatement_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.CashFlowStatement(context.Background(), Params{Symbol: ""}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

func TestCashFlowStatement_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.CashFlowStatement(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
