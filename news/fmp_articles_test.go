package news

import (
	"context"
	"os"
	"testing"
)

func TestFMPArticles_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/fmp-articles.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.FMPArticles(context.Background(), 0)
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Title == "" || rows[0].Tickers != "NYSE:MRK" || rows[0].Author == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/fmp-articles" || cap.query.Get("page") != "0" {
		t.Errorf("delegation: path=%q page=%q", cap.path, cap.query.Get("page"))
	}
}
