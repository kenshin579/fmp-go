package esg

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

func TestRatings_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/esg-ratings.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Ratings(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol=%q", r.Symbol)
	}
	if r.FiscalYear != 2024 {
		t.Errorf("FiscalYear=%d", r.FiscalYear)
	}
	if r.ESGRiskRating == "" {
		t.Error("ESGRiskRating must not be empty")
	}
	if r.IndustryRank == "" {
		t.Error("IndustryRank must not be empty")
	}
}

func TestRatings_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/esg-ratings.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Ratings(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/esg-ratings" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
}

func TestRatings_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.Ratings(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

func TestDisclosures_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/esg-disclosures.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Disclosures(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.ESGScore == 0 {
		t.Error("ESGScore must not be zero")
	}
	if r.EnvironmentalScore == 0 {
		t.Error("EnvironmentalScore must not be zero")
	}
}

func TestDisclosures_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/esg-disclosures.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Disclosures(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/esg-disclosures" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
}

func TestBenchmark_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/esg-benchmark.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Benchmark(context.Background(), "2023")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Sector == "" {
		t.Error("Sector must not be empty")
	}
	if r.ESGScore == 0 {
		t.Error("ESGScore must not be zero")
	}
	if r.FiscalYear != 2023 {
		t.Errorf("FiscalYear=%d", r.FiscalYear)
	}
}

func TestBenchmark_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/esg-benchmark.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Benchmark(context.Background(), "2023")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/esg-benchmark" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("year") != "2023" {
		t.Errorf("year=%q", cap.query.Get("year"))
	}
}

func TestBenchmark_EmptyYear(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.Benchmark(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty year")
	}
}
