package senate

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

func TestSenateTrades_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/senate-trades.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SenateTrades(context.Background(), "AAPL", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d want 1", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol=%q want AAPL", r.Symbol)
	}
	if r.Amount != "$1,001 - $15,000" {
		t.Errorf("Amount=%q want $1,001 - $15,000", r.Amount)
	}
	if r.CapitalGainsOver200USD != "False" {
		t.Errorf("CapitalGainsOver200USD=%q want False", r.CapitalGainsOver200USD)
	}
	if r.Type != "Purchase" {
		t.Errorf("Type=%q want Purchase", r.Type)
	}
}

func TestSenateLatest_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/senate-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SenateLatest(context.Background(), 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d want 1", len(rows))
	}
	r := rows[0]
	if r.Symbol == "" {
		t.Errorf("Symbol should not be empty")
	}
	if r.CapitalGainsOver200USD != "" {
		t.Errorf("CapitalGainsOver200USD=%q want empty (absent key)", r.CapitalGainsOver200USD)
	}
}

func TestSenateLatest_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/senate-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SenateLatest(context.Background(), 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/senate-latest" {
		t.Errorf("path=%q want /stable/senate-latest", cap.path)
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q want 0", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "10" {
		t.Errorf("limit=%q want 10", cap.query.Get("limit"))
	}
}

func TestSenateTrades_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/senate-trades.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SenateTrades(context.Background(), "AAPL", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/senate-trades" {
		t.Errorf("path=%q want /stable/senate-trades", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
}

func TestSenateTradesByName_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/senate-trades.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SenateTradesByName(context.Background(), "Moran")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/senate-trades-by-name" {
		t.Errorf("path=%q want /stable/senate-trades-by-name", cap.path)
	}
	if cap.query.Get("name") != "Moran" {
		t.Errorf("name=%q want Moran", cap.query.Get("name"))
	}
}

func TestHouseLatest_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/house-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.HouseLatest(context.Background(), 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/house-latest" {
		t.Errorf("path=%q want /stable/house-latest", cap.path)
	}
}

func TestHouseTrades_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/house-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.HouseTrades(context.Background(), "AAPL", 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/house-trades" {
		t.Errorf("path=%q want /stable/house-trades", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
}

func TestHouseTradesByName_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/house-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.HouseTradesByName(context.Background(), "Pelosi")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/house-trades-by-name" {
		t.Errorf("path=%q want /stable/house-trades-by-name", cap.path)
	}
	if cap.query.Get("name") != "Pelosi" {
		t.Errorf("name=%q want Pelosi", cap.query.Get("name"))
	}
}

func TestSenateTrades_EmptySymbolGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/senate-trades.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.SenateTrades(context.Background(), "", 0, 5)
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}

func TestSenateTradesByName_EmptyNameGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/senate-trades.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.SenateTradesByName(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty name, got nil")
	}
}
