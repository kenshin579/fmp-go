package marketperf

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

func TestBiggestGainers_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/biggest-gainers.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.BiggestGainers(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("len=%d, want 2", len(rows))
	}
	if rows[0].Symbol == "" {
		t.Error("rows[0].Symbol is empty")
	}
	if rows[0].ChangesPercentage == 0 {
		t.Error("rows[0].ChangesPercentage is 0")
	}
	if rows[0].Price == 0 {
		t.Error("rows[0].Price is 0")
	}
}

func TestBiggestGainers_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/biggest-gainers.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.BiggestGainers(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/biggest-gainers" {
		t.Errorf("path=%q, want /stable/biggest-gainers", cap.path)
	}
}

func TestBiggestLosers_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/biggest-gainers.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.BiggestLosers(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/biggest-losers" {
		t.Errorf("path=%q, want /stable/biggest-losers", cap.path)
	}
}

func TestMostActives_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/biggest-gainers.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.MostActives(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/most-actives" {
		t.Errorf("path=%q, want /stable/most-actives", cap.path)
	}
}
