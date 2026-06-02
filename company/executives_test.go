package company

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestKeyExecutives_ParsesFixtureWithNullable(t *testing.T) {
	raw, _ := os.ReadFile("testdata/key-executives-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.KeyExecutives(context.Background(), "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Pay == nil || *rows[0].Pay <= 0 {
		t.Errorf("row0 Pay should be set: %+v", rows[0])
	}
	if rows[0].YearBorn == nil || *rows[0].YearBorn != 1960 {
		t.Errorf("row0 YearBorn: %v", rows[0].YearBorn)
	}
	if rows[1].Pay != nil {
		t.Errorf("row1 Pay should be nil")
	}
	if rows[1].YearBorn != nil {
		t.Errorf("row1 YearBorn should be nil")
	}
}

func TestExecutiveCompensation_ParsesFixture(t *testing.T) {
	raw, _ := os.ReadFile("testdata/executive-compensation-aapl.json")
	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()
	rows, err := c.ExecutiveCompensation(context.Background(), "AAPL")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Total <= 0 || rows[0].Year == 0 || rows[0].NameAndPosition == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
}

func TestExecutiveCompensationBenchmark_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/executive-compensation-benchmark.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.ExecutiveCompensationBenchmark(context.Background(), 2024)
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].AverageCompensation <= 0 || rows[0].IndustryTitle == "" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/executive-compensation-benchmark" || cap.query.Get("year") != "2024" {
		t.Errorf("delegation: path=%q year=%q", cap.path, cap.query.Get("year"))
	}
}
