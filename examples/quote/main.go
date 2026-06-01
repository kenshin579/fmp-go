// examples/quote — FMP 주식 시세 조회 예제.
//
// 실행: FMP_API_KEY=... go run ./examples/quote
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

	q, err := c.Quote.Quote(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL %.2f (%.2f%%) vol=%d\n", q.Price, q.ChangePercentage, q.Volume)

	batch, err := c.Quote.BatchQuote(ctx, "AAPL", "MSFT", "GOOGL")
	if err != nil {
		log.Fatal(err)
	}
	for _, b := range batch {
		fmt.Printf("  %s %.2f\n", b.Symbol, b.Price)
	}

	pc, err := c.Quote.PriceChange(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 1Y=%.2f%% YTD=%.2f%%\n", pc.Y1, pc.YTD)

	crypto, err := c.Quote.CryptoQuotes(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("crypto quotes: %d개\n", len(crypto))
}
