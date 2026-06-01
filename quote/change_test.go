package quote

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestPriceChange_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/price-change-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	pc, err := c.PriceChange(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("PriceChange: %v", err)
	}
	if pc.Symbol != "AAPL" {
		t.Errorf("Symbol = %q", pc.Symbol)
	}
	if pc.D1 == 0 || pc.Y1 == 0 || pc.Max == 0 {
		t.Errorf("period fields not parsed: %+v", pc)
	}
	if pc.YTD == 0 {
		t.Error("YTD not parsed")
	}
}
