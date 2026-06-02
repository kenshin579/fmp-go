package search

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

func TestSearchSymbol_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-symbol-aapl.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchSymbol(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].Exchange != "NASDAQ" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/search-symbol" || cap.query.Get("query") != "AAPL" {
		t.Errorf("delegation: path=%q query=%q", cap.path, cap.query.Get("query"))
	}
	if _, err := c.SearchSymbol(context.Background(), "  "); err == nil {
		t.Fatal("want empty query guard")
	}
}

func TestSearchName_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-name-apple.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchName(context.Background(), "Apple")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/search-name" || cap.query.Get("query") != "Apple" {
		t.Errorf("delegation: path=%q query=%q", cap.path, cap.query.Get("query"))
	}
}
