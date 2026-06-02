package secfilings

import (
	"context"
	"net/http"
	"os"
	"testing"
)

// ---- SearchByName ----

func TestSearchByName_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/company-search.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.SearchByName(context.Background(), "Apple")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].Name == "" {
		t.Errorf("Name must not be empty: %+v", rows[0])
	}
	if rows[0].SICCode != "3571" {
		t.Errorf("SICCode=%q want 3571", rows[0].SICCode)
	}
}

func TestSearchByName_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/company-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.SearchByName(context.Background(), "Apple")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-company-search/name" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("company") != "Apple" {
		t.Errorf("company=%q", cap.query.Get("company"))
	}
}

func TestSearchByName_EmptyCompanyGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.SearchByName(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty company")
	}
}

// ---- CompanySearchBySymbol ----

func TestCompanySearchBySymbol_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/company-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CompanySearchBySymbol(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-company-search/symbol" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
}

func TestCompanySearchBySymbol_EmptySymbolGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.CompanySearchBySymbol(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty symbol")
	}
}

// ---- CompanySearchByCIK ----

func TestCompanySearchByCIK_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/company-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.CompanySearchByCIK(context.Background(), "0000320193")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/sec-filings-company-search/cik" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("cik") != "0000320193" {
		t.Errorf("cik=%q", cap.query.Get("cik"))
	}
}

func TestCompanySearchByCIK_EmptyCIKGuard(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, "[]")
	defer cleanup()
	_, err := c.CompanySearchByCIK(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty cik")
	}
}

// ---- IndustryClassificationList ----

func TestIndustryClassificationList_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-classification-list.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.IndustryClassificationList(context.Background(), "Agriculture", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].Office == "" {
		t.Errorf("Office must not be empty: %+v", rows[0])
	}
	if rows[0].SICCode != "100" {
		t.Errorf("SICCode=%q want 100", rows[0].SICCode)
	}
}

func TestIndustryClassificationList_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-classification-list.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.IndustryClassificationList(context.Background(), "Agriculture", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/standard-industrial-classification-list" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("industryTitle") != "Agriculture" {
		t.Errorf("industryTitle=%q", cap.query.Get("industryTitle"))
	}
}

// ---- IndustryClassificationSearch ----

func TestIndustryClassificationSearch_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-classification-search.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.IndustryClassificationSearch(context.Background(), "AAPL", "", "")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0].Name == "" {
		t.Errorf("Name must not be empty: %+v", rows[0])
	}
	if rows[0].BusinessAddress == "" {
		t.Errorf("BusinessAddress must not be empty: %+v", rows[0])
	}
}

func TestIndustryClassificationSearch_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/industry-classification-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.IndustryClassificationSearch(context.Background(), "AAPL", "", "3571")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/industry-classification-search" {
		t.Errorf("path=%q", cap.path)
	}
}

// ---- AllIndustryClassification ----

func TestAllIndustryClassification_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/company-search.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	_, err := c.AllIndustryClassification(context.Background(), 0, 10)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if cap.path != "/stable/all-industry-classification" {
		t.Errorf("path=%q", cap.path)
	}
	if cap.query.Get("page") != "0" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
	if cap.query.Get("limit") != "10" {
		t.Errorf("limit=%q", cap.query.Get("limit"))
	}
}
