// 실행: FMP_API_KEY=... go run examples/search/main.go
package main

import (
	"context"
	"fmt"
	"log"

	fmp "github.com/kenshin579/fmp-go"
	"github.com/kenshin579/fmp-go/search"
)

func main() {
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	hits, err := c.Search.SearchSymbol(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	for _, h := range hits {
		fmt.Printf("%s — %s (%s)\n", h.Symbol, h.Name, h.Exchange)
	}

	tech, err := c.Search.CompanyScreener(ctx, search.ScreenerParams{Sector: "Technology", Limit: 5})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tech screener: %d개\n", len(tech))
}
