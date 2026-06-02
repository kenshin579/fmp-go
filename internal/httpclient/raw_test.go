package httpclient

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRaw_ReturnsCsvBodyVerbatim(t *testing.T) {
	want := "symbol,price\nAAPL,225.5\n"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(want))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "k", BaseURL: srv.URL})
	got, err := c.GetRaw(context.Background(), "/stable/bulk-csv", nil)
	if err != nil {
		t.Fatalf("GetRaw: %v", err)
	}
	if string(got) != want {
		t.Errorf("body = %q, want %q", string(got), want)
	}
}

func TestGetRaw_MapsNon200ToAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"Error Message":"Invalid API KEY"}`))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "bad", BaseURL: srv.URL})
	_, err := c.GetRaw(context.Background(), "/stable/bulk-csv", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("want *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 401 {
		t.Errorf("StatusCode = %d, want 401", apiErr.StatusCode)
	}
	if apiErr.Message != "Invalid API KEY" {
		t.Errorf("Message = %q, want %q", apiErr.Message, "Invalid API KEY")
	}
}
