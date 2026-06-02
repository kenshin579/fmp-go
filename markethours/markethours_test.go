package markethours

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

// --- ExchangeMarketHours ---

func TestExchangeMarketHours_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/exchange-market-hours.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.ExchangeMarketHours(context.Background(), "NASDAQ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected len=1, got %d", len(rows))
	}
	if !rows[0].IsMarketOpen {
		t.Errorf("IsMarketOpen should be true")
	}
	if rows[0].Timezone == "" {
		t.Errorf("Timezone must not be empty")
	}
	if rows[0].Exchange != "NASDAQ" {
		t.Errorf("Exchange=%q, want NASDAQ", rows[0].Exchange)
	}
}

func TestExchangeMarketHours_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/exchange-market-hours.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.ExchangeMarketHours(context.Background(), "NASDAQ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/exchange-market-hours" {
		t.Errorf("path=%q, want /stable/exchange-market-hours", cap.path)
	}
	if cap.query.Get("exchange") != "NASDAQ" {
		t.Errorf("exchange query=%q, want NASDAQ", cap.query.Get("exchange"))
	}
}

func TestExchangeMarketHours_EmptyExchangeGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.ExchangeMarketHours(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty exchange, got nil")
	}
}

// --- AllExchangeMarketHours ---

func TestAllExchangeMarketHours_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/all-exchange-market-hours.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.AllExchangeMarketHours(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("expected len=2, got %d", len(rows))
	}
}

func TestAllExchangeMarketHours_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/all-exchange-market-hours.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.AllExchangeMarketHours(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/all-exchange-market-hours" {
		t.Errorf("path=%q, want /stable/all-exchange-market-hours", cap.path)
	}
}

// --- HolidaysByExchange ---

func TestHolidaysByExchange_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/holidays-by-exchange.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.HolidaysByExchange(context.Background(), "NASDAQ", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected len=1, got %d", len(rows))
	}
	if !rows[0].IsClosed {
		t.Errorf("IsClosed should be true")
	}
	if rows[0].AdjOpenTime != nil {
		t.Errorf("AdjOpenTime should be nil")
	}
	if rows[0].AdjCloseTime != nil {
		t.Errorf("AdjCloseTime should be nil")
	}
	if rows[0].Name == "" {
		t.Errorf("Name must not be empty")
	}
}

func TestHolidaysByExchange_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/holidays-by-exchange.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.HolidaysByExchange(context.Background(), "NASDAQ", "2025-01-01", "2026-01-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/holidays-by-exchange" {
		t.Errorf("path=%q, want /stable/holidays-by-exchange", cap.path)
	}
	if cap.query.Get("exchange") != "NASDAQ" {
		t.Errorf("exchange query=%q, want NASDAQ", cap.query.Get("exchange"))
	}
	if cap.query.Get("from") != "2025-01-01" {
		t.Errorf("from query=%q, want 2025-01-01", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2026-01-01" {
		t.Errorf("to query=%q, want 2026-01-01", cap.query.Get("to"))
	}
}

func TestHolidaysByExchange_EmptyExchangeGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.HolidaysByExchange(context.Background(), "", "", "")
	if err == nil {
		t.Error("expected error for empty exchange, got nil")
	}
}
