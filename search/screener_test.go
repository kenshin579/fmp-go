package search

import (
	"context"
	"os"
	"testing"
)

func TestScreenerParams_ToMap(t *testing.T) {
	f := false
	tr := true
	p := ScreenerParams{
		MarketCapMoreThan: 1000000,
		Sector:            "Technology",
		PriceMoreThan:     10.5,
		IsEtf:             &f,
		IsActivelyTrading: &tr,
		Limit:             50,
	}
	m := p.toMap()
	if m["marketCapMoreThan"] != "1000000" {
		t.Errorf("marketCapMoreThan=%q", m["marketCapMoreThan"])
	}
	if m["sector"] != "Technology" {
		t.Errorf("sector=%q", m["sector"])
	}
	if m["priceMoreThan"] != "10.5" {
		t.Errorf("priceMoreThan=%q", m["priceMoreThan"])
	}
	if m["isEtf"] != "false" {
		t.Errorf("isEtf=%q want false", m["isEtf"])
	}
	if m["isActivelyTrading"] != "true" {
		t.Errorf("isActivelyTrading=%q", m["isActivelyTrading"])
	}
	if m["limit"] != "50" {
		t.Errorf("limit=%q", m["limit"])
	}
	if _, ok := m["industry"]; ok {
		t.Error("industry should be omitted")
	}
	if _, ok := m["isFund"]; ok {
		t.Error("isFund(nil) should be omitted")
	}
}

func TestScreenerParams_ToMap_Empty(t *testing.T) {
	if len(ScreenerParams{}.toMap()) != 0 {
		t.Error("empty params should yield empty map")
	}
}

func TestCompanyScreener_ParsesNullableAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/screener.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	tr := true
	rows, err := c.CompanyScreener(context.Background(), ScreenerParams{Sector: "Technology", IsActivelyTrading: &tr, Limit: 5})
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].MarketCap != nil || rows[0].Beta != nil || rows[0].LastAnnualDividend != nil {
		t.Errorf("row0 nullables should be nil: %+v", rows[0])
	}
	if rows[1].MarketCap == nil || *rows[1].MarketCap <= 0 {
		t.Errorf("row1 MarketCap should be set")
	}
	if rows[1].Beta == nil || rows[1].LastAnnualDividend == nil {
		t.Errorf("row1 Beta/Dividend should be set")
	}
	if cap.path != "/stable/company-screener" || cap.query.Get("sector") != "Technology" || cap.query.Get("limit") != "5" {
		t.Errorf("delegation: path=%q sector=%q limit=%q", cap.path, cap.query.Get("sector"), cap.query.Get("limit"))
	}
}
