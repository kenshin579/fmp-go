package company

import (
	"context"
	"os"
	"testing"
)

func TestLatestMergersAcquisitions_DelegatesPage(t *testing.T) {
	raw, _ := os.ReadFile("testdata/mergers-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.LatestMergersAcquisitions(context.Background(), 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].TargetedSymbol == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/mergers-acquisitions-latest" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}

func TestSearchMergersAcquisitions_DelegatesNameAndGuard(t *testing.T) {
	raw, _ := os.ReadFile("testdata/mergers-latest.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	if _, err := c.SearchMergersAcquisitions(context.Background(), "Apple"); err != nil {
		t.Fatalf("Search: %v", err)
	}
	if cap.path != "/stable/mergers-acquisitions-search" || cap.query.Get("name") != "Apple" {
		t.Errorf("delegation: path=%q name=%q", cap.path, cap.query.Get("name"))
	}
	if _, err := c.SearchMergersAcquisitions(context.Background(), "  "); err == nil {
		t.Fatal("want empty name guard")
	}
}
