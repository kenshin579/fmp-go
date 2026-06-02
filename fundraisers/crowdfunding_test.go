package fundraisers

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

func TestLatestCrowdfunding_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crowdfunding-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.LatestCrowdfunding(context.Background(), 0, 10)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].CompanyName == "" {
		t.Errorf("CompanyName empty: %+v", rows[0])
	}
	if rows[0].TotalAssetMostRecentFiscalYear <= 0 {
		t.Errorf("TotalAssetMostRecentFiscalYear not parsed: %v", rows[0].TotalAssetMostRecentFiscalYear)
	}
	// compensationAmount is a string field (free text)
	_ = rows[0].CompensationAmount
}

func TestLatestCrowdfunding_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crowdfunding-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.LatestCrowdfunding(context.Background(), 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/crowdfunding-offerings-latest" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "10" {
		t.Errorf("limit=%q", cap.query.Get("limit"))
	}
}

func TestCrowdfundingByCIK_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crowdfunding-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CrowdfundingByCIK(context.Background(), "0001234567")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/crowdfunding-offerings" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("cik") != "0001234567" {
		t.Errorf("cik=%q", cap.query.Get("cik"))
	}
}

func TestCrowdfundingByCIK_EmptyCIK(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.CrowdfundingByCIK(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty cik")
	}
}

func TestCrowdfundingSearch_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crowdfunding-search.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CrowdfundingSearch(context.Background(), "Acme")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Name != "Acme Inc" {
		t.Errorf("Name=%q", rows[0].Name)
	}
	// null date should deserialize to empty string
	if rows[0].Date != "" {
		t.Errorf("Date=%q (expected empty string from null)", rows[0].Date)
	}
}

func TestCrowdfundingSearch_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crowdfunding-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CrowdfundingSearch(context.Background(), "Acme")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/crowdfunding-offerings-search" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("name") != "Acme" {
		t.Errorf("name=%q", cap.query.Get("name"))
	}
}

func TestCrowdfundingSearch_EmptyName(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.CrowdfundingSearch(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty name")
	}
}
