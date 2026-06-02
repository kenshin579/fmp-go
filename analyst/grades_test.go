package analyst

import (
	"context"
	"errors"
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

func TestGrades_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/grades.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.Grades(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].GradingCompany == "" || rows[0].Action != "maintain" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/grades" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}

func TestGradesConsensus_ParsesSingle(t *testing.T) {
	raw, _ := os.ReadFile("testdata/grades-consensus.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	g, err := c.GradesConsensus(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("GradesConsensus: %v", err)
	}
	if g.Buy != 29 || g.Consensus != "Buy" {
		t.Errorf("not parsed: %+v", g)
	}
}

func TestGradesConsensus_EmptyNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.GradesConsensus(context.Background(), "NOPE"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestHistoricalGrades_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/grades-historical.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.HistoricalGrades(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].AnalystRatingsBuy != 8 || rows[0].Date == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
