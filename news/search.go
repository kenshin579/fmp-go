package news

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SearchStockNews 는 종목별 주식 뉴스를 조회한다.
func (c *Client) SearchStockNews(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/stock", symbols)
}

// SearchCryptoNews 는 종목별 암호화폐 뉴스를 조회한다.
func (c *Client) SearchCryptoNews(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/crypto", symbols)
}

// SearchForexNews 는 통화쌍별 외환 뉴스를 조회한다.
func (c *Client) SearchForexNews(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/forex", symbols)
}

// SearchPressReleases 는 종목별 보도자료를 조회한다.
func (c *Client) SearchPressReleases(ctx context.Context, symbols ...string) ([]Article, error) {
	return fetch.ListBySymbols[Article](ctx, c.http, "/stable/news/press-releases", symbols)
}
