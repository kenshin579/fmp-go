package quote

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

func TestQuote_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/quote-aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	q, err := c.Quote(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("Quote: %v", err)
	}
	if q.Symbol != "AAPL" {
		t.Errorf("Symbol = %q", q.Symbol)
	}
	if q.Price <= 0 || q.MarketCap <= 0 || q.Volume <= 0 {
		t.Errorf("core numeric fields not parsed: %+v", q)
	}
	if q.Exchange != "NASDAQ" {
		t.Errorf("Exchange = %q", q.Exchange)
	}
	if q.Timestamp == 0 {
		t.Error("Timestamp not parsed")
	}
}

func TestQuote_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.Quote(context.Background(), "NOPE"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestQuote_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.Quote(context.Background(), "  "); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}

func TestBatchQuote_ParsesFixtureAndEmptyGuard(t *testing.T) {
	raw, err := os.ReadFile("testdata/batch-quote.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil || len(arr) < 2 {
		t.Fatalf("fixture must have >=2 items")
	}
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.BatchQuote(context.Background(), "AAPL", "MSFT")
	if err != nil {
		t.Fatalf("BatchQuote: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("len = %d, want 2", len(rows))
	}
	if rows[1].Symbol != "MSFT" {
		t.Errorf("rows[1].Symbol = %q", rows[1].Symbol)
	}

	if _, err := c.BatchQuote(context.Background()); err == nil {
		t.Fatal("expected error for empty symbols")
	}
}
