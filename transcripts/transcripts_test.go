package transcripts

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

func TestTranscript_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/transcript.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Transcript(context.Background(), "AAPL", "2020", "Q3", 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" || rows[0].Year != 2020 || rows[0].Period != "Q3" || rows[0].Content == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestTranscript_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/transcript.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Transcript(context.Background(), "AAPL", "2020", "Q3", 0)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/earning-call-transcript" || cap.query.Get("symbol") != "AAPL" || cap.query.Get("year") != "2020" || cap.query.Get("quarter") != "Q3" {
		t.Errorf("delegation: path=%q symbol=%q year=%q quarter=%q", cap.path, cap.query.Get("symbol"), cap.query.Get("year"), cap.query.Get("quarter"))
	}
}

func TestLatest_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Latest(context.Background(), 0, 10)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].FiscalYear != 2025 || rows[0].Period == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestLatest_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Latest(context.Background(), 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/earning-call-transcript-latest" || cap.query.Get("page") != "0" || cap.query.Get("limit") != "10" {
		t.Errorf("delegation: path=%q page=%q limit=%q", cap.path, cap.query.Get("page"), cap.query.Get("limit"))
	}
}

func TestDates_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dates.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Dates(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Quarter != 1 || rows[0].FiscalYear != 2025 || rows[0].Date == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestDates_Delegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dates.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Dates(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/earning-call-transcript-dates" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}

func TestTranscript_GuardEmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.Transcript(context.Background(), "", "2020", "Q3", 0)
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

func TestTranscript_GuardEmptyYear(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.Transcript(context.Background(), "AAPL", "", "Q3", 0)
	if err == nil {
		t.Error("expected error for empty year")
	}
}

func TestTranscript_GuardEmptyQuarter(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.Transcript(context.Background(), "AAPL", "2020", "", 0)
	if err == nil {
		t.Error("expected error for empty quarter")
	}
}

func TestDates_GuardEmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.Dates(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}
