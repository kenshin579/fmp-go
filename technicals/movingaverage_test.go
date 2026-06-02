package technicals

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

func TestSMA_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sma.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.SMA(context.Background(), "AAPL", 10, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Close == 0 {
		t.Errorf("embedded Bar.Close not promoted: got 0")
	}
	if rows[0].SMA == 0 {
		t.Errorf("SMA field not parsed: got 0")
	}
}

func TestSMA_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sma.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.SMA(context.Background(), "AAPL", 10, "1day", "2026-01-01", "2026-04-08")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/technical-indicators/sma" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
	if cap.query.Get("periodLength") != "10" {
		t.Errorf("periodLength=%q", cap.query.Get("periodLength"))
	}
	if cap.query.Get("timeframe") != "1day" {
		t.Errorf("timeframe=%q", cap.query.Get("timeframe"))
	}
	if cap.query.Get("from") != "2026-01-01" {
		t.Errorf("from=%q", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2026-04-08" {
		t.Errorf("to=%q", cap.query.Get("to"))
	}
}

func TestEMA_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ema.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.EMA(context.Background(), "AAPL", 10, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Close == 0 {
		t.Errorf("embedded Bar.Close not promoted: got 0")
	}
	if rows[0].EMA == 0 {
		t.Errorf("EMA field not parsed: got 0")
	}
}

func TestEMA_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/ema.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.EMA(context.Background(), "AAPL", 10, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/technical-indicators/ema" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestSMA_Guards(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sma.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	ctx := context.Background()

	if _, err := c.SMA(ctx, "", 10, "1day", "", ""); err == nil {
		t.Error("expected error for empty symbol")
	}
	if _, err := c.SMA(ctx, "AAPL", 0, "1day", "", ""); err == nil {
		t.Error("expected error for periodLength=0")
	}
	if _, err := c.SMA(ctx, "AAPL", 10, "", "", ""); err == nil {
		t.Error("expected error for empty timeframe")
	}
}
