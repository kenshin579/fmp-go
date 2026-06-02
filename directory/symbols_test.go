package directory

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

func TestCompanySymbolsList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/stock-list.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CompanySymbolsList(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].CompanyName == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestFinancialSymbolsList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-statement-symbol-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.FinancialSymbolsList(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].TradingCurrency == "" || rows[0].ReportingCurrency == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/financial-statement-symbol-list" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestCompanySymbolsList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/stock-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CompanySymbolsList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/stock-list" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestCIKList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/cik-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.CIKList(context.Background(), 0, 10)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].CIK != "0002036063" {
		t.Errorf("CIK 0-padding not preserved: %q", rows[0].CIK)
	}
	if cap.path != "/stable/cik-list" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("page") != "0" || cap.query.Get("limit") != "10" {
		t.Errorf("query page=%q limit=%q", cap.query.Get("page"), cap.query.Get("limit"))
	}
}

func TestSymbolChangesList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/symbol-change.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SymbolChangesList(context.Background(), true, 10)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].NewSymbol != "X" {
		t.Errorf("NewSymbol=%q", rows[0].NewSymbol)
	}
	if cap.path != "/stable/symbol-change" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("invalid") != "true" || cap.query.Get("limit") != "10" {
		t.Errorf("query invalid=%q limit=%q", cap.query.Get("invalid"), cap.query.Get("limit"))
	}
}

func TestETFsList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/etf-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.ETFsList(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "SPY" || rows[0].Name == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/etf-list" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestActivelyTradingList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/actively-trading-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.ActivelyTradingList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/actively-trading-list" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestEarningsTranscriptList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/earnings-transcript-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.EarningsTranscriptList(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].NoOfTranscripts != "16" {
		t.Errorf("NoOfTranscripts=%q (want string \"16\")", rows[0].NoOfTranscripts)
	}
	if cap.path != "/stable/earnings-transcript-list" {
		t.Errorf("path=%q", cap.path)
	}
}
