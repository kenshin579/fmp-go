package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON_InjectsAPIKeyAndDecodes(t *testing.T) {
	var gotKey, gotSymbol, gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKey = r.URL.Query().Get("apikey")
		gotSymbol = r.URL.Query().Get("symbol")
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"symbol":"AAPL","companyName":"Apple Inc."}]`))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "k123", BaseURL: srv.URL})
	var out []map[string]any
	err := c.GetJSON(context.Background(), "/stable/profile", map[string]string{"symbol": "AAPL"}, &out)
	if err != nil {
		t.Fatalf("GetJSON: %v", err)
	}
	if gotKey != "k123" {
		t.Errorf("apikey = %q, want k123", gotKey)
	}
	if gotSymbol != "AAPL" {
		t.Errorf("symbol = %q, want AAPL", gotSymbol)
	}
	if gotPath != "/stable/profile" {
		t.Errorf("path = %q", gotPath)
	}
	if len(out) != 1 || out[0]["symbol"] != "AAPL" {
		t.Errorf("decoded = %v", out)
	}
}

func TestGetJSON_HTTPErrorStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"Error Message":"Invalid API KEY."}`))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "bad", BaseURL: srv.URL})
	var out []map[string]any
	err := c.GetJSON(context.Background(), "/stable/profile", nil, &out)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("want *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 401 || apiErr.Message != "Invalid API KEY." {
		t.Errorf("APIError = %+v", apiErr)
	}
}

func TestGetJSON_ErrorBodyWith200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Error Message":"Limit Reach."}`))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "k", BaseURL: srv.URL})
	var out []map[string]any
	err := c.GetJSON(context.Background(), "/stable/profile", nil, &out)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("want *APIError, got %T: %v", err, err)
	}
	if apiErr.Message != "Limit Reach." {
		t.Errorf("message = %q", apiErr.Message)
	}
}
