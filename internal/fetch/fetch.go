// Package fetch 는 FMP sub-client 들이 공유하는 generic 조회 helper.
// 모든 그룹 패키지(quote/company/...)가 이 helper 로 endpoint 를 구현한다.
package fetch

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// OneBySymbol — {symbol} 단일 레코드. 빈 symbol 가드, 빈 배열 → httpclient.ErrNotFound.
func OneBySymbol[T any](ctx context.Context, hc *httpclient.Client, path, symbol string) (*T, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return One[T](ctx, hc, path, map[string]string{"symbol": symbol})
}

// ListBySymbol — {symbol} 리스트(시계열/다건). 빈 symbol 가드.
func ListBySymbol[T any](ctx context.Context, hc *httpclient.Client, path, symbol string) ([]T, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return List[T](ctx, hc, path, map[string]string{"symbol": symbol})
}

// ListBySymbols — {symbols:쉼표 join} 배치 리스트. 빈 symbols 가드.
func ListBySymbols[T any](ctx context.Context, hc *httpclient.Client, path string, symbols []string) ([]T, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("fmp: symbols must not be empty")
	}
	return List[T](ctx, hc, path, map[string]string{"symbols": strings.Join(symbols, ",")})
}

// One — 임의 params 단일 레코드. 빈 배열 → httpclient.ErrNotFound.
func One[T any](ctx context.Context, hc *httpclient.Client, path string, params map[string]string) (*T, error) {
	var out []T
	if err := hc.GetJSON(ctx, path, params, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return &out[0], nil
}

// List — 임의 params 리스트.
func List[T any](ctx context.Context, hc *httpclient.Client, path string, params map[string]string) ([]T, error) {
	var out []T
	err := hc.GetJSON(ctx, path, params, &out)
	return out, err
}
