package company

import (
	"context"
	"os"
	"testing"
)

func TestDelistedCompanies_DelegatesPage(t *testing.T) {
	raw, _ := os.ReadFile("testdata/delisted-companies.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.DelistedCompanies(context.Background(), 1)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].DelistedDate == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/delisted-companies" || cap.query.Get("page") != "1" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}
