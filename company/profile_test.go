package company

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func newTestClient(t *testing.T, status int, body string) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, srv.Close
}

func TestProfile_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/profile_aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil {
		t.Fatalf("fixture is not a JSON array: %v", err)
	}
	if len(arr) == 0 {
		t.Fatal("fixture array empty")
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	p, err := c.Profile(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("Profile: %v", err)
	}
	if p.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", p.Symbol)
	}
	if p.CompanyName == "" {
		t.Error("CompanyName empty")
	}
	if p.CEO == "" {
		t.Error("CEO empty")
	}
	if p.Website == "" {
		t.Error("Website empty")
	}
}

func TestProfile_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()

	_, err := c.Profile(context.Background(), "NOPE")
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
