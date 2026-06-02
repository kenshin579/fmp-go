package metrics

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

func TestKeyMetrics_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/key-metrics.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.KeyMetrics(context.Background(), "AAPL", "annual", 1)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol=%q want AAPL", r.Symbol)
	}
	if r.MarketCap == 0 {
		t.Error("MarketCap must not be 0")
	}
	if r.FreeCashFlowToFirm == 0 {
		t.Error("FreeCashFlowToFirm must not be 0")
	}
	if r.ResearchAndDevelopementToRevenue == 0 {
		t.Error("ResearchAndDevelopementToRevenue must not be 0")
	}
	if r.ReturnOnEquity == 0 {
		t.Error("ReturnOnEquity must not be 0")
	}
}

func TestKeyMetrics_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/key-metrics.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.KeyMetrics(context.Background(), "AAPL", "annual", 2)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/key-metrics" {
		t.Errorf("path=%q want /stable/key-metrics", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("period") != "annual" {
		t.Errorf("period=%q want annual", cap.query.Get("period"))
	}
	if cap.query.Get("limit") != "2" {
		t.Errorf("limit=%q want 2", cap.query.Get("limit"))
	}
}

func TestKeyMetrics_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.KeyMetrics(context.Background(), "", "annual", 0)
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}

func TestKeyMetricsTTM_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/key-metrics-ttm.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.KeyMetricsTTM(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].EnterpriseValueTTM == 0 {
		t.Error("EnterpriseValueTTM must not be 0")
	}
	if cap.path != "/stable/key-metrics-ttm" {
		t.Errorf("path=%q want /stable/key-metrics-ttm", cap.path)
	}
}
