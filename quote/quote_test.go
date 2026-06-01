package quote

import (
	"context"
	"encoding/json"
	"errors"
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

type capturedReq struct {
	path  string
	query url.Values
}

// newCapturingClient 는 요청 path + query 를 캡처해 path 오타·파라미터 누락을 검증한다.
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

func TestQuote_DelegatesPathAndParams(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, `[{"symbol":"AAPL"}]`)
	defer cleanup()
	if _, err := c.Quote(context.Background(), "AAPL"); err != nil {
		t.Fatalf("Quote: %v", err)
	}
	if cap.path != "/stable/quote" {
		t.Errorf("path = %q, want /stable/quote", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol param = %q", cap.query.Get("symbol"))
	}
}

func TestBatchQuote_DelegatesSymbolsJoin(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, `[]`)
	defer cleanup()
	_, _ = c.BatchQuote(context.Background(), "AAPL", "MSFT", "GOOGL")
	if cap.path != "/stable/batch-quote" {
		t.Errorf("path = %q", cap.path)
	}
	if cap.query.Get("symbols") != "AAPL,MSFT,GOOGL" {
		t.Errorf("symbols param = %q, want comma-joined", cap.query.Get("symbols"))
	}
}

func TestExchangeQuotes_DelegatesPathAndExchange(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, `[]`)
	defer cleanup()
	_, _ = c.ExchangeQuotes(context.Background(), "NASDAQ")
	if cap.path != "/stable/batch-exchange-quote" {
		t.Errorf("path = %q", cap.path)
	}
	if cap.query.Get("exchange") != "NASDAQ" {
		t.Errorf("exchange param = %q", cap.query.Get("exchange"))
	}
}

func TestCryptoQuotes_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, `[]`)
	defer cleanup()
	_, _ = c.CryptoQuotes(context.Background())
	if cap.path != "/stable/batch-crypto-quotes" {
		t.Errorf("path = %q", cap.path)
	}
}

func TestCryptoQuotes_EmptyIsNotError(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	rows, err := c.CryptoQuotes(context.Background())
	if err != nil {
		t.Fatalf("empty list should not error: %v", err)
	}
	if len(rows) != 0 {
		t.Errorf("rows = %+v, want empty", rows)
	}
}
