package quote

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestExchangeQuotes_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/exchange-quotes.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.ExchangeQuotes(context.Background(), "NASDAQ")
	if err != nil {
		t.Fatalf("ExchangeQuotes: %v", err)
	}
	if len(rows) != 2 || rows[0].Symbol != "AAPL" || rows[0].Price <= 0 {
		t.Errorf("rows = %+v", rows)
	}
}

func TestCryptoQuotes_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crypto-quotes.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CryptoQuotes(context.Background())
	if err != nil {
		t.Fatalf("CryptoQuotes: %v", err)
	}
	if len(rows) != 2 || rows[0].Symbol != "BTCUSD" || rows[0].Price <= 0 {
		t.Errorf("rows = %+v", rows)
	}
}
