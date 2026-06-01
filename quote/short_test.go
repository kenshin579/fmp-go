package quote

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func TestQuoteShort_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/quote-short-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	q, err := c.QuoteShort(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("QuoteShort: %v", err)
	}
	if q.Symbol != "AAPL" || q.Price <= 0 || q.Volume <= 0 {
		t.Errorf("not parsed: %+v", q)
	}
}

func TestQuoteShort_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.QuoteShort(context.Background(), "NOPE"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestBatchQuoteShort_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/batch-quote-short.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.BatchQuoteShort(context.Background(), "AAPL", "MSFT")
	if err != nil {
		t.Fatalf("BatchQuoteShort: %v", err)
	}
	if len(rows) != 2 || rows[1].Symbol != "MSFT" {
		t.Errorf("rows = %+v", rows)
	}
}
