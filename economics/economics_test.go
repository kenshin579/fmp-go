package economics

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

// TreasuryRates parse test
func TestTreasuryRates_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/treasury-rates.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.TreasuryRates(context.Background(), "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].Year10 == 0 {
		t.Errorf("Year10 should not be zero: %+v", rows[0])
	}
	if rows[0].Month1 == 0 {
		t.Errorf("Month1 should not be zero: %+v", rows[0])
	}
}

// TreasuryRates delegation test
func TestTreasuryRates_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/treasury-rates.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.TreasuryRates(context.Background(), "2025-01-01", "2025-02-03")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/treasury-rates" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("from") != "2025-01-01" || cap.query.Get("to") != "2025-02-03" {
		t.Errorf("from=%q to=%q", cap.query.Get("from"), cap.query.Get("to"))
	}
}

// EconomicIndicators delegation test
func TestEconomicIndicators_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/economic-indicators.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.EconomicIndicators(context.Background(), "CPI", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/economic-indicators" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("name") != "CPI" {
		t.Errorf("name=%q", cap.query.Get("name"))
	}
}

// EconomicIndicators empty name guard
func TestEconomicIndicators_EmptyNameGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/economic-indicators.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	_, err := c.EconomicIndicators(context.Background(), "", "", "")
	if err == nil {
		t.Error("expected error for empty name, got nil")
	}
}

// EconomicCalendar parse test (Estimate==nil, Unit==nil)
func TestEconomicCalendar_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/economic-calendar.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.EconomicCalendar(context.Background(), "US", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].Estimate != nil {
		t.Errorf("Estimate should be nil, got %v", rows[0].Estimate)
	}
	if rows[0].Unit != nil {
		t.Errorf("Unit should be nil, got %v", rows[0].Unit)
	}
	if rows[0].Previous == 0 {
		t.Errorf("Previous should not be zero: %+v", rows[0])
	}
	if rows[0].Actual == 0 {
		t.Errorf("Actual should not be zero: %+v", rows[0])
	}
}

// EconomicCalendar delegation test
func TestEconomicCalendar_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/economic-calendar.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.EconomicCalendar(context.Background(), "US", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/economic-calendar" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("country") != "US" {
		t.Errorf("country=%q", cap.query.Get("country"))
	}
}

// MarketRiskPremium parse test
func TestMarketRiskPremium_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/market-risk-premium.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.MarketRiskPremium(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].TotalEquityRiskPremium == 0 {
		t.Errorf("TotalEquityRiskPremium should not be zero: %+v", rows[0])
	}
	if rows[0].Country == "" {
		t.Errorf("Country should not be empty: %+v", rows[0])
	}
}

// MarketRiskPremium delegation test
func TestMarketRiskPremium_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/market-risk-premium.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.MarketRiskPremium(context.Background())
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/market-risk-premium" {
		t.Errorf("path=%q", cap.path)
	}
}
