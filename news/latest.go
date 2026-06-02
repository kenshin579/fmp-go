package news

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

func pageParams(page int) map[string]string {
	return map[string]string{"page": strconv.Itoa(page)}
}

// StockNewsLatest 는 최신 주식 뉴스를 페이지 단위로 조회한다.
func (c *Client) StockNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/stock-latest", pageParams(page))
}

// CryptoNewsLatest 는 최신 암호화폐 뉴스를 조회한다.
func (c *Client) CryptoNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/crypto-latest", pageParams(page))
}

// ForexNewsLatest 는 최신 외환 뉴스를 조회한다.
func (c *Client) ForexNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/forex-latest", pageParams(page))
}

// GeneralNewsLatest 는 최신 일반 경제 뉴스를 조회한다.
func (c *Client) GeneralNewsLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/general-latest", pageParams(page))
}

// PressReleasesLatest 는 최신 보도자료를 조회한다.
func (c *Client) PressReleasesLatest(ctx context.Context, page int) ([]Article, error) {
	return fetch.List[Article](ctx, c.http, "/stable/news/press-releases-latest", pageParams(page))
}
