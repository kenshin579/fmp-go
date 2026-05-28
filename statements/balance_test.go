package statements

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func TestBalanceSheetStatement_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/balance-sheet-statement-aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil || len(arr) == 0 {
		t.Fatalf("fixture invalid/empty: %v", err)
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.BalanceSheetStatement(context.Background(), Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("BalanceSheetStatement: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("rows empty")
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.TotalAssets <= 0 {
		t.Errorf("TotalAssets = %d", r.TotalAssets)
	}
	if r.TotalLiabilities <= 0 {
		t.Errorf("TotalLiabilities = %d", r.TotalLiabilities)
	}
	if r.TotalStockholdersEquity <= 0 {
		t.Errorf("TotalStockholdersEquity = %d", r.TotalStockholdersEquity)
	}
}

func TestBalanceSheetStatement_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.BalanceSheetStatement(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
