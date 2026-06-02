package ratios

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func TestRatiosTTM_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/ratios-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.RatiosTTM(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("RatiosTTM: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.GrossProfitMarginTTM == 0 {
		t.Error("GrossProfitMarginTTM not parsed")
	}
	if r.EnterpriseValueTTM == 0 {
		t.Error("EnterpriseValueTTM not parsed")
	}
	if r.PriceToEarningsRatioTTM == 0 {
		t.Error("PriceToEarningsRatioTTM not parsed")
	}
}

func TestRatiosTTM_PathAndQuery(t *testing.T) {
	raw, err := os.ReadFile("testdata/ratios-ttm.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	var gotPath, gotSymbol string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotSymbol = r.URL.Query().Get("symbol")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	}))
	defer srv.Close()

	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	if _, err := c.RatiosTTM(context.Background(), "AAPL"); err != nil {
		t.Fatalf("RatiosTTM: %v", err)
	}

	if gotPath != "/stable/ratios-ttm" {
		t.Errorf("path = %q, want /stable/ratios-ttm", gotPath)
	}
	if gotSymbol != "AAPL" {
		t.Errorf("symbol query = %q, want AAPL", gotSymbol)
	}
}

func TestRatiosTTM_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.RatiosTTM(context.Background(), "  "); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

func TestRatiosTTM_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.RatiosTTM(context.Background(), "NOPE")
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
