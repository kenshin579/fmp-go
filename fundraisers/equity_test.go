package fundraisers

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestLatestEquityOffering_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/equity-latest.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.LatestEquityOffering(context.Background(), 0, 10, "")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}

	// int64 field parses correctly
	if rows[0].TotalOfferingAmount != 186842 {
		t.Errorf("TotalOfferingAmount=%d want 186842", rows[0].TotalOfferingAmount)
	}

	// bool field parses correctly
	if !rows[0].IsAmendment {
		t.Errorf("IsAmendment should be true")
	}

	// nullable *bool that is null in fixture should be nil
	if rows[0].SecuritiesOfferedAreOfEquityType != nil {
		t.Errorf("SecuritiesOfferedAreOfEquityType should be nil (null in fixture), got %v", rows[0].SecuritiesOfferedAreOfEquityType)
	}
}

func TestLatestEquityOffering_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/equity-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.LatestEquityOffering(context.Background(), 0, 10, "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/fundraising-latest" {
		t.Errorf("path=%q want /stable/fundraising-latest", cap.path)
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q want 0", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "10" {
		t.Errorf("limit=%q want 10", cap.query.Get("limit"))
	}
}

func TestEquityOfferingByCIK_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/equity-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.EquityOfferingByCIK(context.Background(), "0001234567")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/fundraising" {
		t.Errorf("path=%q want /stable/fundraising", cap.path)
	}
	if cap.query.Get("cik") != "0001234567" {
		t.Errorf("cik=%q want 0001234567", cap.query.Get("cik"))
	}
}

func TestEquityOfferingByCIK_EmptyCIK(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.EquityOfferingByCIK(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty cik")
	}
}

func TestEquityOfferingSearch_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/crowdfunding-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.EquityOfferingSearch(context.Background(), "Tesla")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/fundraising-search" {
		t.Errorf("path=%q want /stable/fundraising-search", cap.path)
	}
	if cap.query.Get("name") != "Tesla" {
		t.Errorf("name=%q want Tesla", cap.query.Get("name"))
	}
}

func TestEquityOfferingSearch_EmptyName(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()

	_, err := c.EquityOfferingSearch(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty name")
	}
}
