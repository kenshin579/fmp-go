package calendar

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

func TestDateRange_OmitsEmpty(t *testing.T) {
	m := dateRange("2025-01-01", "")
	if m["from"] != "2025-01-01" {
		t.Errorf("from=%q", m["from"])
	}
	if _, ok := m["to"]; ok {
		t.Error("empty to should be omitted")
	}
	if len(dateRange("", "")) != 0 {
		t.Error("both empty should yield empty map")
	}
}

func TestDividendsCalendar_DelegatesDateRange(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dividends-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.DividendsCalendar(context.Background(), "2025-02-01", "2025-02-28")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Frequency != "Semi-Annual" || rows[0].Yield <= 0 {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/dividends-calendar" || cap.query.Get("from") != "2025-02-01" || cap.query.Get("to") != "2025-02-28" {
		t.Errorf("delegation: path=%q from=%q to=%q", cap.path, cap.query.Get("from"), cap.query.Get("to"))
	}
}

func TestCompanyDividends_DelegatesSymbol(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dividends-company.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.CompanyDividends(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].PaymentDate == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/dividends" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}
