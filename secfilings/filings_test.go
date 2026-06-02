package secfilings

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

// ---- LatestFinancials ----

func TestLatestFinancials_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financials-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.LatestFinancials(context.Background(), "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if !rows[0].HasFinancials {
		t.Errorf("HasFinancials should be true: %+v", rows[0])
	}
	if rows[0].FormType == "" {
		t.Errorf("FormType must not be empty: %+v", rows[0])
	}
	if rows[0].Symbol != "AAPL" {
		t.Errorf("Symbol=%q want AAPL", rows[0].Symbol)
	}
}

func TestLatestFinancials_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financials-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.LatestFinancials(context.Background(), "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-financials" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("from") != "2024-01-01" {
		t.Errorf("from=%q", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2024-03-01" {
		t.Errorf("to=%q", cap.query.Get("to"))
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
}

func TestLatestFinancials_EmptyFromGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.LatestFinancials(context.Background(), "", "2024-03-01", 0, 5)
	if err == nil {
		t.Error("expected error for empty from")
	}
}

// ---- Latest8K ----

func TestLatest8K_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/8k-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Latest8K(context.Background(), "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].HasFinancials {
		t.Errorf("HasFinancials should be false: %+v", rows[0])
	}
}

func TestLatest8K_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/8k-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Latest8K(context.Background(), "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-8k" {
		t.Errorf("path=%q", cap.path)
	}
}

// ---- SearchBySymbol ----

func TestSearchBySymbol_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-symbol.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchBySymbol(context.Background(), "AAPL", "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].FormType == "" {
		t.Errorf("FormType must not be empty: %+v", rows[0])
	}
}

func TestSearchBySymbol_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-symbol.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SearchBySymbol(context.Background(), "AAPL", "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-search/symbol" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
	if cap.query.Get("from") != "2024-01-01" {
		t.Errorf("from=%q", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2024-03-01" {
		t.Errorf("to=%q", cap.query.Get("to"))
	}
}

func TestSearchBySymbol_EmptySymbolGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.SearchBySymbol(context.Background(), "", "2024-01-01", "2024-03-01", 0, 5)
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// ---- SearchByCIK ----

func TestSearchByCIK_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-symbol.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SearchByCIK(context.Background(), "0000320193", "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-search/cik" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("cik") != "0000320193" {
		t.Errorf("cik=%q", cap.query.Get("cik"))
	}
}

// ---- SearchByFormType ----

func TestSearchByFormType_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-symbol.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SearchByFormType(context.Background(), "10-K", "2024-01-01", "2024-03-01", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-search/form-type" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("formType") != "10-K" {
		t.Errorf("formType=%q", cap.query.Get("formType"))
	}
}
