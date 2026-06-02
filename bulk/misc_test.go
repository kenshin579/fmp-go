package bulk

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

const profileCSV = "symbol,companyName\nAAPL,Apple Inc\n"

// Profile

func TestProfile_ParsesCSV(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, profileCSV)
	defer cleanup()

	body, err := c.Profile(context.Background(), "0")
	if err != nil {
		t.Fatalf("Profile: %v", err)
	}
	if len(body) == 0 {
		t.Fatal("expected non-empty body")
	}
	if !strings.Contains(string(body), "symbol") {
		t.Errorf("body does not contain 'symbol': %q", string(body))
	}
}

func TestProfile_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, profileCSV)
	defer cleanup()

	_, err := c.Profile(context.Background(), "0")
	if err != nil {
		t.Fatalf("Profile: %v", err)
	}
	if cap.path != "/stable/profile-bulk" {
		t.Errorf("path = %q, want /stable/profile-bulk", cap.path)
	}
	if cap.query.Get("part") != "0" {
		t.Errorf("part = %q, want 0", cap.query.Get("part"))
	}
}

func TestProfile_EmptyPartGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, profileCSV)
	defer cleanup()

	_, err := c.Profile(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty part, got nil")
	}
}

// ETFHolder

func TestETFHolder_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()

	_, err := c.ETFHolder(context.Background(), "1")
	if err != nil {
		t.Fatalf("ETFHolder: %v", err)
	}
	if cap.path != "/stable/etf-holder-bulk" {
		t.Errorf("path = %q, want /stable/etf-holder-bulk", cap.path)
	}
	if cap.query.Get("part") != "1" {
		t.Errorf("part = %q, want 1", cap.query.Get("part"))
	}
}

func TestETFHolder_EmptyPartGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, csvBody)
	defer cleanup()

	_, err := c.ETFHolder(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty part, got nil")
	}
}

// EOD

func TestEOD_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()

	_, err := c.EOD(context.Background(), "2024-10-22")
	if err != nil {
		t.Fatalf("EOD: %v", err)
	}
	if cap.path != "/stable/eod-bulk" {
		t.Errorf("path = %q, want /stable/eod-bulk", cap.path)
	}
	if cap.query.Get("date") != "2024-10-22" {
		t.Errorf("date = %q, want 2024-10-22", cap.query.Get("date"))
	}
}

func TestEOD_EmptyDateGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, csvBody)
	defer cleanup()

	_, err := c.EOD(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty date, got nil")
	}
}

// EarningsSurprises

func TestEarningsSurprises_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()

	_, err := c.EarningsSurprises(context.Background(), "2024")
	if err != nil {
		t.Fatalf("EarningsSurprises: %v", err)
	}
	if cap.path != "/stable/earnings-surprises-bulk" {
		t.Errorf("path = %q, want /stable/earnings-surprises-bulk", cap.path)
	}
	if cap.query.Get("year") != "2024" {
		t.Errorf("year = %q, want 2024", cap.query.Get("year"))
	}
}

func TestEarningsSurprises_EmptyYearGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, csvBody)
	defer cleanup()

	_, err := c.EarningsSurprises(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty year, got nil")
	}
}

// Scores

func TestScores_ParsesCSV(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, csvBody)
	defer cleanup()

	body, err := c.Scores(context.Background())
	if err != nil {
		t.Fatalf("Scores: %v", err)
	}
	if len(body) == 0 {
		t.Fatal("expected non-empty body")
	}
}

func TestScores_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()

	_, err := c.Scores(context.Background())
	if err != nil {
		t.Fatalf("Scores: %v", err)
	}
	if cap.path != "/stable/scores-bulk" {
		t.Errorf("path = %q, want /stable/scores-bulk", cap.path)
	}
}

// No-param methods delegation

func TestPeers_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.Peers(context.Background())
	if err != nil {
		t.Fatalf("Peers: %v", err)
	}
	if cap.path != "/stable/peers-bulk" {
		t.Errorf("path = %q, want /stable/peers-bulk", cap.path)
	}
}

func TestDCF_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.DCF(context.Background())
	if err != nil {
		t.Fatalf("DCF: %v", err)
	}
	if cap.path != "/stable/dcf-bulk" {
		t.Errorf("path = %q, want /stable/dcf-bulk", cap.path)
	}
}

func TestRating_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.Rating(context.Background())
	if err != nil {
		t.Fatalf("Rating: %v", err)
	}
	if cap.path != "/stable/rating-bulk" {
		t.Errorf("path = %q, want /stable/rating-bulk", cap.path)
	}
}

func TestRatiosTTM_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.RatiosTTM(context.Background())
	if err != nil {
		t.Fatalf("RatiosTTM: %v", err)
	}
	if cap.path != "/stable/ratios-ttm-bulk" {
		t.Errorf("path = %q, want /stable/ratios-ttm-bulk", cap.path)
	}
}

func TestKeyMetricsTTM_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.KeyMetricsTTM(context.Background())
	if err != nil {
		t.Fatalf("KeyMetricsTTM: %v", err)
	}
	if cap.path != "/stable/key-metrics-ttm-bulk" {
		t.Errorf("path = %q, want /stable/key-metrics-ttm-bulk", cap.path)
	}
}

func TestPriceTargetSummary_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.PriceTargetSummary(context.Background())
	if err != nil {
		t.Fatalf("PriceTargetSummary: %v", err)
	}
	if cap.path != "/stable/price-target-summary-bulk" {
		t.Errorf("path = %q, want /stable/price-target-summary-bulk", cap.path)
	}
}

func TestUpgradesDowngradesConsensus_DelegatesPath(t *testing.T) {
	c, cap, cleanup := newCapturingClient(t, csvBody)
	defer cleanup()
	_, err := c.UpgradesDowngradesConsensus(context.Background())
	if err != nil {
		t.Fatalf("UpgradesDowngradesConsensus: %v", err)
	}
	if cap.path != "/stable/upgrades-downgrades-consensus-bulk" {
		t.Errorf("path = %q, want /stable/upgrades-downgrades-consensus-bulk", cap.path)
	}
}
