package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestCompanyNotes_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/company-notes-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.CompanyNotes(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Title == "" || rows[0].Exchange == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}
