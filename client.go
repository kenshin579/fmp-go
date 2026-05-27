// Package fmp 는 Financial Modeling Prep(FMP) API 의 Go 클라이언트다.
package fmp

import (
	"errors"
	"time"

	"github.com/kenshin579/fmp-go/company"
	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 fmp-go 라이브러리의 단일 진입점.
type Client struct {
	http *httpclient.Client

	Company *company.Client // 회사 정보(프로필 등)
}

// NewClient 는 API 키로 Client 를 만든다.
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("fmp: apiKey is required")
	}
	cfg := clientOptions{timeout: 30 * time.Second}
	for _, opt := range opts {
		opt(&cfg)
	}
	hc := httpclient.New(httpclient.Config{
		APIKey:     apiKey,
		BaseURL:    cfg.baseURL,
		Timeout:    cfg.timeout,
		HTTPClient: cfg.httpClient,
	})
	c := &Client{http: hc}
	c.Company = company.New(hc)
	return c, nil
}
