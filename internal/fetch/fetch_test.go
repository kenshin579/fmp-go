package fetch

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

type rec struct {
	Symbol string `json:"symbol"`
	V      int    `json:"v"`
}

type captured struct {
	path  string
	query url.Values
}

func newHC(t *testing.T, body string, capture *captured) *httpclient.Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if capture != nil {
			capture.path = r.URL.Path
			capture.query = r.URL.Query()
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL})
}

func TestOneBySymbol_ParsesAndDelegates(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[{"symbol":"AAPL","v":1}]`, cap)
	got, err := OneBySymbol[rec](context.Background(), hc, "/stable/x", "AAPL")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if got.Symbol != "AAPL" || got.V != 1 {
		t.Errorf("got %+v", got)
	}
	if cap.path != "/stable/x" || cap.query.Get("symbol") != "AAPL" {
		t.Errorf("delegation: path=%q symbol=%q", cap.path, cap.query.Get("symbol"))
	}
}

func TestOneBySymbol_EmptyArrayNotFound(t *testing.T) {
	hc := newHC(t, `[]`, nil)
	if _, err := OneBySymbol[rec](context.Background(), hc, "/x", "AAPL"); !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestOneBySymbol_EmptySymbolGuard(t *testing.T) {
	hc := newHC(t, `[]`, nil)
	if _, err := OneBySymbol[rec](context.Background(), hc, "/x", "  "); err == nil {
		t.Fatal("want guard error")
	}
}

func TestListBySymbols_JoinsAndGuards(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[{"symbol":"A"},{"symbol":"B"}]`, cap)
	rows, err := ListBySymbols[rec](context.Background(), hc, "/x", []string{"A", "B"})
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.query.Get("symbols") != "A,B" {
		t.Errorf("symbols=%q want A,B", cap.query.Get("symbols"))
	}
	if _, err := ListBySymbols[rec](context.Background(), hc, "/x", nil); err == nil {
		t.Fatal("want empty symbols guard")
	}
}

func TestListBySymbol_ParsesAndDelegates(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[{"symbol":"AAPL","v":1},{"symbol":"AAPL","v":2}]`, cap)
	rows, err := ListBySymbol[rec](context.Background(), hc, "/x", "AAPL")
	if err != nil || len(rows) != 2 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if cap.query.Get("symbol") != "AAPL" {
		t.Errorf("symbol=%q", cap.query.Get("symbol"))
	}
}

func TestList_ArbitraryParamsAndEmptyOK(t *testing.T) {
	cap := &captured{}
	hc := newHC(t, `[]`, cap)
	rows, err := List[rec](context.Background(), hc, "/x", map[string]string{"page": "2"})
	if err != nil {
		t.Fatalf("empty list should not error: %v", err)
	}
	if len(rows) != 0 {
		t.Errorf("rows=%+v", rows)
	}
	if cap.query.Get("page") != "2" {
		t.Errorf("page=%q", cap.query.Get("page"))
	}
}
