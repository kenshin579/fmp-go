package quote

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestAftermarketQuote_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/aftermarket-quote-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	q, err := c.AftermarketQuote(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("AftermarketQuote: %v", err)
	}
	if q.Symbol != "AAPL" || q.BidPrice <= 0 || q.AskPrice <= 0 || q.Timestamp == 0 {
		t.Errorf("not parsed: %+v", q)
	}
}

func TestAftermarketTrade_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/aftermarket-trade-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	tr, err := c.AftermarketTrade(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("AftermarketTrade: %v", err)
	}
	if tr.Symbol != "AAPL" || tr.Price <= 0 || tr.TradeSize <= 0 {
		t.Errorf("not parsed: %+v", tr)
	}
}

func TestBatchAftermarket_ParsesFixtures(t *testing.T) {
	rawQ, _ := os.ReadFile("testdata/batch-aftermarket-quote.json")
	cQ, cleanupQ := newTestClient(t, http.StatusOK, string(rawQ))
	defer cleanupQ()
	qs, err := cQ.BatchAftermarketQuote(context.Background(), "AAPL", "MSFT")
	if err != nil || len(qs) != 2 || qs[1].Symbol != "MSFT" {
		t.Fatalf("BatchAftermarketQuote: err=%v rows=%+v", err, qs)
	}

	rawT, _ := os.ReadFile("testdata/batch-aftermarket-trade.json")
	cT, cleanupT := newTestClient(t, http.StatusOK, string(rawT))
	defer cleanupT()
	ts, err := cT.BatchAftermarketTrade(context.Background(), "AAPL", "MSFT")
	if err != nil || len(ts) != 2 || ts[1].Symbol != "MSFT" {
		t.Fatalf("BatchAftermarketTrade: err=%v rows=%+v", err, ts)
	}
}
