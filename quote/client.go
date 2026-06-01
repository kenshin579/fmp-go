// Package quote 는 FMP 시세(quote) API sub-client.
// fmp.Client.Quote 로 접근.
package quote

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 시세 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// fetchOne 은 단일 심볼 조회 공통 — FMP 배열 응답의 첫 요소를 *T 로 반환.
// 빈 symbol → 에러, 빈 배열 → httpclient.ErrNotFound.
func fetchOne[T any](ctx context.Context, c *Client, path, symbol string) (*T, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []T
	if err := c.http.GetJSON(ctx, path, map[string]string{"symbol": symbol}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return &out[0], nil
}

// fetchBatch 은 symbols 배치 조회 공통 — 쉼표 join.
func fetchBatch[T any](ctx context.Context, c *Client, path string, symbols []string) ([]T, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("fmp: symbols must not be empty")
	}
	var out []T
	err := c.http.GetJSON(ctx, path, map[string]string{"symbols": strings.Join(symbols, ",")}, &out)
	return out, err
}

// fetchList 은 단일 키 또는 무파라미터 리스트 조회 공통.
func fetchList[T any](ctx context.Context, c *Client, path string, params map[string]string) ([]T, error) {
	var out []T
	err := c.http.GetJSON(ctx, path, params, &out)
	return out, err
}
