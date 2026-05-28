package ratios

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

func TestRatios_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/ratios-aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil || len(arr) == 0 {
		t.Fatalf("fixture invalid/empty: %v", err)
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Ratios(context.Background(), Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("Ratios: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("rows empty")
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q", r.Symbol)
	}
	if r.BookValuePerShare == 0 {
		t.Error("BookValuePerShare not parsed")
	}
	if r.DebtToEquityRatio == 0 {
		t.Error("DebtToEquityRatio not parsed")
	}
}

func TestRatios_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.Ratios(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestRatios_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.Ratios(context.Background(), Params{Symbol: "  "}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}
