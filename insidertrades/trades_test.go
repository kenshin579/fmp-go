package insidertrades

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

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

func TestLatestInsiderTrades_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.LatestInsiderTrades(context.Background(), "", 0, 5)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol=%q", r.Symbol)
	}
	if r.SecuritiesTransacted == 0 {
		t.Errorf("SecuritiesTransacted=0")
	}
	if r.Price == 0 {
		t.Errorf("Price=0")
	}
	if r.DirectOrIndirect == "" {
		t.Errorf("DirectOrIndirect empty")
	}
	if r.FormType == "" {
		t.Errorf("FormType empty")
	}
	if cap.path != "/stable/insider-trading/latest" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "5" {
		t.Errorf("limit=%q", cap.query.Get("limit"))
	}
}

func TestSearchInsiderTrades_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.SearchInsiderTrades(context.Background(), SearchParams{
		Symbol:          "AAPL",
		TransactionType: "P-Purchase",
		Limit:           10,
	})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/insider-trading/search" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
	if cap.query.Get("transactionType") != "P-Purchase" {
		t.Errorf("transactionType=%q", cap.query.Get("transactionType"))
	}
	if cap.query.Get("limit") != "10" {
		t.Errorf("limit=%q", cap.query.Get("limit"))
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
}

func TestSearchInsiderTrades_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search.json")
	c, _, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	rows, err := c.SearchInsiderTrades(context.Background(), SearchParams{Symbol: "AAPL"})
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].TransactionType == "" {
		t.Errorf("TransactionType empty")
	}
}
