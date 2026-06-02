package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestStockPeers_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/peers-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.StockPeers(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol == "" || rows[0].MktCap <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
