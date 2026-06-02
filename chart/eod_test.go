package chart

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

func TestHistoricalPriceEODLight_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/eod-light.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.HistoricalPriceEODLight(context.Background(), "AAPL", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Symbol != "AAPL" {
		t.Errorf("Symbol=%q, want AAPL", rows[0].Symbol)
	}
	if rows[0].Price == 0 {
		t.Error("Price must not be zero")
	}
	if rows[0].Volume == 0 {
		t.Error("Volume must not be zero")
	}
}

func TestHistoricalPriceEODFull_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/eod-full.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.HistoricalPriceEODFull(context.Background(), "AAPL", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Open == 0 {
		t.Error("Open must not be zero")
	}
	if rows[0].Close == 0 {
		t.Error("Close must not be zero")
	}
	if rows[0].VWAP == 0 {
		t.Error("VWAP must not be zero")
	}
}

func TestHistoricalPriceEODFull_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/eod-full.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.HistoricalPriceEODFull(context.Background(), "AAPL", "2024-01-01", "2024-09-27")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/historical-price-eod/full" {
		t.Errorf("path=%q, want /stable/historical-price-eod/full", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q, want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("from") != "2024-01-01" {
		t.Errorf("from=%q, want 2024-01-01", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2024-09-27" {
		t.Errorf("to=%q, want 2024-09-27", cap.query.Get("to"))
	}
}

func TestHistoricalPriceEODDividendAdjusted_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/eod-dividend-adjusted.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.HistoricalPriceEODDividendAdjusted(context.Background(), "AAPL", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].AdjClose == 0 {
		t.Error("AdjClose must not be zero")
	}
}

func TestHistoricalPriceEODNonSplitAdjusted_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/eod-non-split-adjusted.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.HistoricalPriceEODNonSplitAdjusted(context.Background(), "AAPL", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].AdjClose == 0 {
		t.Error("AdjClose must not be zero")
	}
	if cap.path != "/stable/historical-price-eod/non-split-adjusted" {
		t.Errorf("path=%q, want /stable/historical-price-eod/non-split-adjusted", cap.path)
	}
}

func TestHistoricalPriceEODLight_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.HistoricalPriceEODLight(context.Background(), "", "", "")
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}
