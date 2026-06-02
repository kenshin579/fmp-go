package news

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

func TestStockNewsLatest_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/stock-news-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.StockNewsLatest(context.Background(), 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "INSG" || rows[0].Title == "" || rows[0].URL == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/news/stock-latest" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}

func TestGeneralNewsLatest_NullSymbolDecodes(t *testing.T) {
	raw, _ := os.ReadFile("testdata/general-news-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.GeneralNewsLatest(context.Background(), 1)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "" {
		t.Errorf("null symbol should decode to empty string, got %q", rows[0].Symbol)
	}
	if rows[0].Publisher != "CNBC" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestCryptoForexPressLatest_Delegate(t *testing.T) {
	cases := []struct {
		name string
		call func(c *Client) ([]Article, error)
		path string
	}{
		{"crypto", func(c *Client) ([]Article, error) { return c.CryptoNewsLatest(context.Background(), 0) }, "/stable/news/crypto-latest"},
		{"forex", func(c *Client) ([]Article, error) { return c.ForexNewsLatest(context.Background(), 0) }, "/stable/news/forex-latest"},
		{"press", func(c *Client) ([]Article, error) { return c.PressReleasesLatest(context.Background(), 0) }, "/stable/news/press-releases-latest"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, cap, cleanup := newCapturingClient(t, `[]`)
			defer cleanup()
			if _, err := tc.call(c); err != nil {
				t.Fatalf("%s: %v", tc.name, err)
			}
			if cap.path != tc.path {
				t.Errorf("path=%q want %q", cap.path, tc.path)
			}
		})
	}
}
