package assets

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

func TestCryptoList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crypto-list.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CryptoList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].CirculatingSupply == 0 {
		t.Errorf("CirculatingSupply should not be zero: %+v", rows[0])
	}
	if rows[0].TotalSupply == nil {
		t.Errorf("rows[0].TotalSupply should not be nil (value present)")
	}
	if rows[1].TotalSupply != nil {
		t.Errorf("rows[1].TotalSupply should be nil (null in JSON): %v", rows[1].TotalSupply)
	}
}

func TestCryptoList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crypto-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CryptoList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/cryptocurrency-list" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestForexList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/forex-list.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.ForexList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].FromCurrency != "EUR" {
		t.Errorf("FromCurrency=%q", rows[0].FromCurrency)
	}
	if rows[0].ToCurrency != "USD" {
		t.Errorf("ToCurrency=%q", rows[0].ToCurrency)
	}
}

func TestForexList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/forex-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.ForexList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/forex-list" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestCommodityList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/commodities-list.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CommodityList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].Exchange != nil {
		t.Errorf("Exchange should be nil (null in JSON): %v", rows[0].Exchange)
	}
	if rows[0].TradeMonth != "Dec" {
		t.Errorf("TradeMonth=%q", rows[0].TradeMonth)
	}
	if rows[0].Currency != "USD" {
		t.Errorf("Currency=%q", rows[0].Currency)
	}
}

func TestCommodityList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/commodities-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CommodityList(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/commodities-list" {
		t.Errorf("path=%q", cap.path)
	}
}
