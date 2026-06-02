package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestSharesFloat_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/shares-float-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	s, err := c.SharesFloat(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("SharesFloat: %v", err)
	}
	if s.Symbol != "AAPL" || s.OutstandingShares <= 0 {
		t.Errorf("not parsed: %+v", s)
	}
}

func TestAllSharesFloat_DelegatesPage(t *testing.T) {
	raw, _ := os.ReadFile("testdata/all-shares-float.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.AllSharesFloat(context.Background(), 2)
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.path != "/stable/shares-float-all" || cap.query.Get("page") != "2" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}
