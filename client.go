// Package fmp 는 Financial Modeling Prep(FMP) API 의 Go 클라이언트다.
package fmp

import (
	"errors"
	"time"

	"github.com/kenshin579/fmp-go/company"
	"github.com/kenshin579/fmp-go/internal/httpclient"
	"github.com/kenshin579/fmp-go/quote"
	"github.com/kenshin579/fmp-go/ratios"
	"github.com/kenshin579/fmp-go/statements"
)

// Client 는 fmp-go 라이브러리의 단일 진입점.
type Client struct {
	http *httpclient.Client

	Company    *company.Client    // 회사 정보(프로필 등)
	Statements *statements.Client // 재무제표(소득, 대차대조표 등)
	Ratios     *ratios.Client     // 재무비율
	Quote      *quote.Client      // 시세(실시간/배치/자산군)
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
	c.Statements = statements.New(hc)
	c.Ratios = ratios.New(hc)
	c.Quote = quote.New(hc)
	return c, nil
}
