package fmp_test

import (
	"testing"

	fmp "github.com/kenshin579/fmp-go"
)

func TestNewClient_EmptyKey(t *testing.T) {
	_, err := fmp.NewClient("")
	if err == nil {
		t.Fatal("expected error for empty apiKey")
	}
}

func TestNewClient_HasCompany(t *testing.T) {
	c, err := fmp.NewClient("k123")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Company == nil {
		t.Fatal("Company sub-client is nil")
	}
}

func TestNewClientFromEnv_MissingEnv(t *testing.T) {
	t.Setenv("FMP_API_KEY", "")
	if _, err := fmp.NewClientFromEnv(); err == nil {
		t.Fatal("expected error when FMP_API_KEY unset")
	}
}

func TestNewClientFromEnv_Reads(t *testing.T) {
	t.Setenv("FMP_API_KEY", "envkey")
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	if c.Company == nil {
		t.Fatal("Company nil")
	}
}
