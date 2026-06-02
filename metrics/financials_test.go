package metrics

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// --- FinancialScores ---

func TestFinancialScores_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/financial-scores.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	r, err := c.FinancialScores(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if r.PiotroskiScore != 8 {
		t.Errorf("PiotroskiScore=%d want 8", r.PiotroskiScore)
	}
	if r.AltmanZScore == 0 {
		t.Error("AltmanZScore must not be 0")
	}
	if r.MarketCap == 0 {
		t.Error("MarketCap must not be 0")
	}
}

func TestFinancialScores_EmptyArray_ErrNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.FinancialScores(context.Background(), "AAPL")
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Errorf("err=%v want httpclient.ErrNotFound", err)
	}
}

func TestFinancialScores_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.FinancialScores(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}

// --- OwnerEarnings ---

func TestOwnerEarnings_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/owner-earnings.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.OwnerEarnings(context.Background(), "AAPL", 1)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.OwnersEarnings == 0 {
		t.Error("OwnersEarnings must not be 0")
	}
	if r.OwnersEarningsPerShare == 0 {
		t.Error("OwnersEarningsPerShare must not be 0")
	}
	if r.FiscalYear != "2025" {
		t.Errorf("FiscalYear=%q want 2025", r.FiscalYear)
	}
}

func TestOwnerEarnings_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.OwnerEarnings(context.Background(), "", 0)
	if err == nil {
		t.Error("expected error for empty symbol, got nil")
	}
}

// --- EnterpriseValues ---

func TestEnterpriseValues_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/enterprise-values.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.EnterpriseValues(context.Background(), "AAPL", "annual", 1)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	r := rows[0]
	if r.EnterpriseValue == 0 {
		t.Error("EnterpriseValue must not be 0")
	}
	if r.StockPrice == 0 {
		t.Error("StockPrice must not be 0")
	}
}

func TestEnterpriseValues_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/enterprise-values.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.EnterpriseValues(context.Background(), "AAPL", "annual", 2)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/enterprise-values" {
		t.Errorf("path=%q want /stable/enterprise-values", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
	if cap.query.Get("period") != "annual" {
		t.Errorf("period=%q want annual", cap.query.Get("period"))
	}
	if cap.query.Get("limit") != "2" {
		t.Errorf("limit=%q want 2", cap.query.Get("limit"))
	}
}
