package secfilings

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestProfile_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/profile.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Profile(context.Background(), "AAPL", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d want 1", len(rows))
	}
	p := rows[0]
	if p.RegistrantName == "" {
		t.Errorf("RegistrantName must not be empty")
	}
	if !p.IsActive {
		t.Errorf("IsActive should be true, got %v", p.IsActive)
	}
	if p.SecurityType != nil {
		t.Errorf("SecurityType should be nil (null), got %v", p.SecurityType)
	}
	if p.Employees != "164000" {
		t.Errorf("Employees=%q want 164000", p.Employees)
	}
	if p.SICCode != "3571" {
		t.Errorf("SICCode=%q want 3571", p.SICCode)
	}
	if p.IsEtf {
		t.Errorf("IsEtf should be false")
	}
	if p.IsAdr {
		t.Errorf("IsAdr should be false")
	}
	if p.IsFund {
		t.Errorf("IsFund should be false")
	}
}

func TestProfile_DelegationBySymbol(t *testing.T) {
	raw, _ := os.ReadFile("testdata/profile.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Profile(context.Background(), "AAPL", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-profile" {
		t.Errorf("path=%q want /stable/sec-profile", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q want AAPL", cap.query.Get("symbol"))
	}
}

func TestProfile_DelegationByCIK(t *testing.T) {
	raw, _ := os.ReadFile("testdata/profile.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Profile(context.Background(), "", "320193")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-profile" {
		t.Errorf("path=%q want /stable/sec-profile", cap.path)
	}
	if cap.query.Get("cik") != "320193" {
		t.Errorf("cik=%q want 320193", cap.query.Get("cik"))
	}
}

func TestProfile_EmptyGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.Profile(context.Background(), "", "")
	if err == nil {
		t.Error("expected error when both symbol and cik are empty")
	}
}
