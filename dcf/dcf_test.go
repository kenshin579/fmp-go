package dcf

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

func TestDiscountedCashFlow_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/discounted-cash-flow.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.DiscountedCashFlow(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" {
		t.Errorf("Symbol=%q want AAPL", rows[0].Symbol)
	}
	if rows[0].DCF == 0 {
		t.Errorf("DCF must not be zero")
	}
	if rows[0].StockPrice == 0 {
		t.Errorf("StockPrice must not be zero (\"Stock Price\" key mapping)")
	}
	if cap.path != "/stable/discounted-cash-flow" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}

func TestLeveredDiscountedCashFlow_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/levered-discounted-cash-flow.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.LeveredDiscountedCashFlow(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" {
		t.Errorf("Symbol=%q want AAPL", rows[0].Symbol)
	}
	if rows[0].DCF == 0 {
		t.Errorf("DCF must not be zero")
	}
	if rows[0].StockPrice == 0 {
		t.Errorf("StockPrice must not be zero (\"Stock Price\" key mapping)")
	}
	if cap.path != "/stable/levered-discounted-cash-flow" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}

func TestDiscountedCashFlow_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.DiscountedCashFlow(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

func TestLeveredDiscountedCashFlow_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.LeveredDiscountedCashFlow(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}
