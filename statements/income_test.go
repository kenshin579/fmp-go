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

func newTestClient(t *testing.T, status int, body string) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, srv.Close
}

func TestIncomeStatement_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/income-statement-aapl.json")
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

	rows, err := c.IncomeStatement(context.Background(), Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("IncomeStatement: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("rows empty")
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.Revenue <= 0 {
		t.Errorf("Revenue = %d, want > 0", r.Revenue)
	}
	if r.GrossProfit <= 0 {
		t.Errorf("GrossProfit = %d, want > 0", r.GrossProfit)
	}
	if r.OperatingIncome <= 0 {
		t.Errorf("OperatingIncome = %d, want > 0", r.OperatingIncome)
	}
	if r.NetIncome <= 0 {
		t.Errorf("NetIncome = %d, want > 0", r.NetIncome)
	}
	if r.EPS == 0 {
		t.Error("EPS not parsed")
	}
	if r.ReportedCurrency != "USD" {
		t.Errorf("ReportedCurrency = %q", r.ReportedCurrency)
	}
}

func TestIncomeStatement_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.IncomeStatement(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestIncomeStatement_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.IncomeStatement(context.Background(), Params{Symbol: "  "}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}
