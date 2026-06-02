package form13f

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

func TestLatestFilings_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.LatestFilings(context.Background(), 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	if rows[0].Cik == "" {
		t.Errorf("Cik is empty: %+v", rows[0])
	}
	if rows[0].Name == "" {
		t.Errorf("Name is empty: %+v", rows[0])
	}
}

func TestLatestFilings_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.LatestFilings(context.Background(), 0, 5)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/latest" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "5" {
		t.Errorf("limit=%q", cap.query.Get("limit"))
	}
}

func TestExtract_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/extract.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Extract(context.Background(), "0001388838", "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	if rows[0].Shares == 0 {
		t.Errorf("Shares is 0: %+v", rows[0])
	}
	if rows[0].Value == 0 {
		t.Errorf("Value is 0: %+v", rows[0])
	}
}

func TestExtract_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/extract.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Extract(context.Background(), "0001388838", "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/extract" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("cik") != "0001388838" {
		t.Errorf("cik=%q", cap.query.Get("cik"))
	}
	if cap.query.Get("year") != "2023" {
		t.Errorf("year=%q", cap.query.Get("year"))
	}
	if cap.query.Get("quarter") != "3" {
		t.Errorf("quarter=%q", cap.query.Get("quarter"))
	}
}

func TestExtract_EmptyCikGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/extract.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.Extract(context.Background(), "", "2023", "3")
	if err == nil {
		t.Error("expected error for empty cik")
	}
}

func TestFilingDates_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dates.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.FilingDates(context.Background(), "0001067983")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	if rows[0].Year == 0 {
		t.Errorf("Year is 0: %+v", rows[0])
	}
	if rows[0].Quarter == 0 {
		t.Errorf("Quarter is 0: %+v", rows[0])
	}
}

func TestFilingDates_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/dates.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.FilingDates(context.Background(), "0001067983")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/dates" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("cik") != "0001067983" {
		t.Errorf("cik=%q", cap.query.Get("cik"))
	}
}

func TestIndustrySummary_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-summary.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.IndustrySummary(context.Background(), "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) < 1 {
		t.Fatalf("expected >=1 rows, got %d", len(rows))
	}
	if rows[0].IndustryValue == 0 {
		t.Errorf("IndustryValue is 0: %+v", rows[0])
	}
	if rows[0].IndustryTitle == "" {
		t.Errorf("IndustryTitle is empty: %+v", rows[0])
	}
}

func TestIndustrySummary_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-summary.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.IndustrySummary(context.Background(), "2023", "3")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/institutional-ownership/industry-summary" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("year") != "2023" {
		t.Errorf("year=%q", cap.query.Get("year"))
	}
	if cap.query.Get("quarter") != "3" {
		t.Errorf("quarter=%q", cap.query.Get("quarter"))
	}
}

func TestIndustrySummary_EmptyYearGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-summary.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.IndustrySummary(context.Background(), "", "3")
	if err == nil {
		t.Error("expected error for empty year")
	}
}
