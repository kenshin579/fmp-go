// Package fmp 는 Financial Modeling Prep(FMP) API 의 Go 클라이언트다.
package fmp

import (
	"errors"
	"time"

	"github.com/kenshin579/fmp-go/analyst"
	"github.com/kenshin579/fmp-go/calendar"
	"github.com/kenshin579/fmp-go/chart"
	"github.com/kenshin579/fmp-go/company"
	"github.com/kenshin579/fmp-go/internal/httpclient"
	"github.com/kenshin579/fmp-go/marketperf"
	"github.com/kenshin579/fmp-go/metrics"
	"github.com/kenshin579/fmp-go/news"
	"github.com/kenshin579/fmp-go/quote"
	"github.com/kenshin579/fmp-go/ratios"
	"github.com/kenshin579/fmp-go/reports"
	"github.com/kenshin579/fmp-go/search"
	"github.com/kenshin579/fmp-go/statements"
)

// Client 는 fmp-go 라이브러리의 단일 진입점.
type Client struct {
	http *httpclient.Client

	Analyst    *analyst.Client    // 애널리스트(등급/목표주가/추정)
	Company    *company.Client    // 회사 정보(프로필 등)
	Statements *statements.Client // 재무제표(소득, 대차대조표 등)
	Ratios     *ratios.Client     // 재무비율
	Quote      *quote.Client      // 시세(실시간/배치/자산군)
	Search     *search.Client     // 검색(심볼/식별자/스크리너)
	News       *news.Client       // 뉴스(주식/암호화폐/외환/보도자료/일반)
	Calendar   *calendar.Client   // 캘린더(배당/실적/IPO/분할)
	Metrics    *metrics.Client    // 지표(key-metrics/scores/owner-earnings/EV/segments)
	Reports    *reports.Client    // 보고서(as-reported/latest/dates/10-K JSON)
	Chart      *chart.Client      // 과거 시세(EOD/intraday)

	MarketPerformance *marketperf.Client // 시장 성과(등락/섹터/산업/PE)
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
	c.Analyst = analyst.New(hc)
	c.Company = company.New(hc)
	c.Statements = statements.New(hc)
	c.Ratios = ratios.New(hc)
	c.Quote = quote.New(hc)
	c.Search = search.New(hc)
	c.News = news.New(hc)
	c.Calendar = calendar.New(hc)
	c.Metrics = metrics.New(hc)
	c.Reports = reports.New(hc)
	c.Chart = chart.New(hc)
	c.MarketPerformance = marketperf.New(hc)
	return c, nil
}
