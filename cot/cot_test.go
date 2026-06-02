package cot

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

func TestAnalysis_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/analysis.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Analysis(context.Background(), "ES", "", "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Symbol != "ES" {
		t.Errorf("Symbol=%q want ES", r.Symbol)
	}
	if r.NetPostion == 0 {
		t.Errorf("NetPostion should be non-zero (typo key maps)")
	}
	if r.ReversalTrend != false {
		t.Errorf("ReversalTrend=%v want false", r.ReversalTrend)
	}
	if r.MarketSentiment == "" {
		t.Errorf("MarketSentiment should not be empty")
	}
}

func TestAnalysis_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/analysis.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Analysis(context.Background(), "ES", "2025-01-01", "2025-02-04")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/commitment-of-traders-analysis" {
		t.Errorf("path=%q want /stable/commitment-of-traders-analysis", cap.path)
	}
	if cap.query.Get("symbol") != "ES" {
		t.Errorf("symbol=%q want ES", cap.query.Get("symbol"))
	}
	if cap.query.Get("from") != "2025-01-01" {
		t.Errorf("from=%q want 2025-01-01", cap.query.Get("from"))
	}
	if cap.query.Get("to") != "2025-02-04" {
		t.Errorf("to=%q want 2025-02-04", cap.query.Get("to"))
	}
}

func TestList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/list.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.List(context.Background())
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Symbol != "ES" {
		t.Errorf("Symbol=%q want ES", r.Symbol)
	}
	if r.Name == "" {
		t.Errorf("Name should not be empty")
	}
}

func TestList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.List(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/commitment-of-traders-list" {
		t.Errorf("path=%q want /stable/commitment-of-traders-list", cap.path)
	}
}
