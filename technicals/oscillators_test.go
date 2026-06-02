package technicals

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestRSI_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/rsi.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.RSI(context.Background(), "AAPL", 14, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Close == 0 {
		t.Errorf("embedded Bar.Close not promoted: got 0")
	}
	if rows[0].RSI == 0 {
		t.Errorf("RSI field not parsed: got 0")
	}
}

func TestRSI_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/rsi.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.RSI(context.Background(), "AAPL", 14, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/technical-indicators/rsi" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestStandardDeviation_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/standarddeviation.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.StandardDeviation(context.Background(), "AAPL", 20, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Close == 0 {
		t.Errorf("embedded Bar.Close not promoted: got 0")
	}
	if rows[0].StandardDeviation == 0 {
		t.Errorf("StandardDeviation field not parsed (camelCase JSON key): got 0")
	}
}

func TestStandardDeviation_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/standarddeviation.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.StandardDeviation(context.Background(), "AAPL", 20, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/technical-indicators/standarddeviation" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestADX_Parse(t *testing.T) {
	raw, _ := os.ReadFile("testdata/adx.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.ADX(context.Background(), "AAPL", 14, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Close == 0 {
		t.Errorf("embedded Bar.Close not promoted: got 0")
	}
	if rows[0].ADX == 0 {
		t.Errorf("ADX field not parsed: got 0")
	}
}

func TestADX_Delegation(t *testing.T) {
	raw, _ := os.ReadFile("testdata/adx.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.ADX(context.Background(), "AAPL", 14, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/technical-indicators/adx" {
		t.Errorf("path=%q", cap.path)
	}
}

func TestWilliams_Delegation(t *testing.T) {
	// Williams: path 확인만 (rsi.json 바디 재사용 — williams 필드 없어도 path 검증에 충분)
	raw, _ := os.ReadFile("testdata/rsi.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()

	_, err := c.Williams(context.Background(), "AAPL", 14, "1day", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cap.path != "/stable/technical-indicators/williams" {
		t.Errorf("path=%q", cap.path)
	}
}
