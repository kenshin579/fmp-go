package etf

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

// --- Holdings ---

func TestHoldings_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/holdings.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Holdings(context.Background(), "SPY")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.MarketValue == 0 {
		t.Errorf("MarketValue should not be zero: %v", r.MarketValue)
	}
	if r.SharesNumber == 0 {
		t.Errorf("SharesNumber should not be zero: %v", r.SharesNumber)
	}
	if r.Symbol != "SPY" {
		t.Errorf("Symbol=%q want SPY", r.Symbol)
	}
}

func TestHoldings_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/holdings.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Holdings(context.Background(), "SPY")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/etf/holdings" {
		t.Errorf("path=%q want /stable/etf/holdings", cap.path)
	}
	if cap.query.Get("symbol") != "SPY" {
		t.Errorf("symbol=%q want SPY", cap.query.Get("symbol"))
	}
}

func TestHoldings_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.Holdings(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// --- Information ---

func TestInformation_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/info.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.Information(context.Background(), "SPY")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if len(r.SectorsList) != 2 {
		t.Errorf("SectorsList len=%d want 2", len(r.SectorsList))
	}
	if r.SectorsList[0].Industry == "" {
		t.Errorf("SectorsList[0].Industry is empty")
	}
	if r.NAV == 0 {
		t.Errorf("NAV should not be zero")
	}
	if r.HoldingsCount == 0 {
		t.Errorf("HoldingsCount should not be zero")
	}
	if r.AssetsUnderManagement == 0 {
		t.Errorf("AssetsUnderManagement should not be zero")
	}
}

func TestInformation_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/info.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.Information(context.Background(), "SPY")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/etf/info" {
		t.Errorf("path=%q want /stable/etf/info", cap.path)
	}
}

func TestInformation_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.Information(context.Background(), "  ")
	if err == nil {
		t.Error("expected error for blank symbol")
	}
}

// --- CountryWeightings ---

func TestCountryWeightings_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/country-weightings.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CountryWeightings(context.Background(), "SPY")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].WeightPercentage != "97.29%" {
		t.Errorf("WeightPercentage=%q want 97.29%%", rows[0].WeightPercentage)
	}
}

func TestCountryWeightings_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/country-weightings.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CountryWeightings(context.Background(), "SPY")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/etf/country-weightings" {
		t.Errorf("path=%q want /stable/etf/country-weightings", cap.path)
	}
}

func TestCountryWeightings_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.CountryWeightings(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// --- SectorWeightings ---

func TestSectorWeightings_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-weightings.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SectorWeightings(context.Background(), "SPY")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.WeightPercentage == 0 {
		t.Errorf("WeightPercentage should not be zero")
	}
	if r.Symbol != "SPY" {
		t.Errorf("Symbol=%q want SPY", r.Symbol)
	}
}

func TestSectorWeightings_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/sector-weightings.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SectorWeightings(context.Background(), "SPY")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/etf/sector-weightings" {
		t.Errorf("path=%q want /stable/etf/sector-weightings", cap.path)
	}
}

func TestSectorWeightings_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.SectorWeightings(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// --- AssetExposure ---

func TestAssetExposure_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/asset-exposure.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.AssetExposure(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.Asset != "AAPL" {
		t.Errorf("Asset=%q want AAPL", r.Asset)
	}
	if r.WeightPercentage == 0 {
		t.Errorf("WeightPercentage should not be zero")
	}
}

func TestAssetExposure_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/asset-exposure.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.AssetExposure(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/etf/asset-exposure" {
		t.Errorf("path=%q want /stable/etf/asset-exposure", cap.path)
	}
}

func TestAssetExposure_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.AssetExposure(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}
