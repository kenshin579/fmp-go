package company

import (
	"context"
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

type capturedReq struct {
	path  string
	query url.Values
}

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

func TestMarketCap_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/market-cap-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	m, err := c.MarketCap(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("MarketCap: %v", err)
	}
	if m.Symbol != "AAPL" || m.MarketCap <= 0 || m.Date == "" {
		t.Errorf("not parsed: %+v", m)
	}
}

func TestHistoricalMarketCap_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/historical-market-cap-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HistoricalMarketCap(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
}

func TestBatchMarketCap_DelegatesSymbols(t *testing.T) {
	raw, _ := os.ReadFile("testdata/batch-market-cap.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.BatchMarketCap(context.Background(), "AAPL", "MSFT")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/market-capitalization-batch" || cap.query.Get("symbols") != "AAPL,MSFT" {
		t.Errorf("delegation: path=%q symbols=%q", cap.path, cap.query.Get("symbols"))
	}
}

func TestProfileByCIK_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/profile-cik-aapl.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	p, err := c.ProfileByCIK(context.Background(), "320193")
	if err != nil {
		t.Fatalf("ProfileByCIK: %v", err)
	}
	if p.Symbol != "AAPL" {
		t.Errorf("Symbol=%q", p.Symbol)
	}
	if cap.path != "/stable/profile-cik" || cap.query.Get("cik") != "320193" {
		t.Errorf("delegation: path=%q cik=%q", cap.path, cap.query.Get("cik"))
	}
	if _, err := c.ProfileByCIK(context.Background(), "  "); err == nil {
		t.Fatal("want empty cik guard")
	}
}
