package news

import (
	"context"
	"os"
	"testing"
)

func TestSearchStockNews_ParsesAndDelegates(t *testing.T) {
	raw, _ := os.ReadFile("testdata/search-stock-news.json")
	c, cap, cleanup := newCapturingClient(t, string(raw))
	defer cleanup()
	rows, err := c.SearchStockNews(context.Background(), "AAPL", "MSFT")
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v len=%d", err, len(rows))
	}
	if rows[0].Symbol != "AAPL" {
		t.Errorf("not parsed: %+v", rows[0])
	}
	if cap.path != "/stable/news/stock" || cap.query.Get("symbols") != "AAPL,MSFT" {
		t.Errorf("delegation: path=%q symbols=%q", cap.path, cap.query.Get("symbols"))
	}
}

func TestSearchNews_EmptySymbolsGuard(t *testing.T) {
	c, cleanup := newTestClient(t, 200, `[]`)
	defer cleanup()
	if _, err := c.SearchStockNews(context.Background()); err == nil {
		t.Fatal("want empty symbols guard")
	}
}

func TestSearchCryptoForexPress_Delegate(t *testing.T) {
	cases := []struct {
		name string
		call func(c *Client) ([]Article, error)
		path string
	}{
		{"crypto", func(c *Client) ([]Article, error) { return c.SearchCryptoNews(context.Background(), "BTCUSD") }, "/stable/news/crypto"},
		{"forex", func(c *Client) ([]Article, error) { return c.SearchForexNews(context.Background(), "EURUSD") }, "/stable/news/forex"},
		{"press", func(c *Client) ([]Article, error) { return c.SearchPressReleases(context.Background(), "AAPL") }, "/stable/news/press-releases"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, cap, cleanup := newCapturingClient(t, `[]`)
			defer cleanup()
			if _, err := tc.call(c); err != nil {
				t.Fatalf("%s: %v", tc.name, err)
			}
			if cap.path != tc.path {
				t.Errorf("path=%q want %q", cap.path, tc.path)
			}
		})
	}
}
