package search

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSearchExchangeVariants_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/exchange-variants-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchExchangeVariants(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].MktCap <= 0 || rows[0].CompanyName == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if _, err := c.SearchExchangeVariants(context.Background(), "  "); err == nil {
		t.Fatal("want empty symbol guard")
	}
}
