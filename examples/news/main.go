// 실행: FMP_API_KEY=... go run examples/news/main.go
package main

import (
	"context"
	"fmt"
	"log"

	fmp "github.com/kenshin579/fmp-go"
)

func main() {
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	hits, err := c.News.SearchStockNews(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	for i, a := range hits {
		if i >= 3 {
			break
		}
		fmt.Printf("[%s] %s — %s\n", a.PublishedDate, a.Title, a.Site)
	}

	latest, err := c.News.StockNewsLatest(ctx, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("latest stock news: %d개\n", len(latest))
}
