// examples/company — FMP 회사 정보 조회 예제.
//
// 실행: FMP_API_KEY=... go run examples/company/main.go
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

	mc, err := c.Company.MarketCap(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL marketCap=%d (%s)\n", mc.MarketCap, mc.Date)

	peers, err := c.Company.StockPeers(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("peers: %d개\n", len(peers))

	execs, err := c.Company.KeyExecutives(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range execs {
		fmt.Printf("  %s — %s\n", e.Title, e.Name)
	}
}
