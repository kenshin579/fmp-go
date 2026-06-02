// 실행: FMP_API_KEY=... go run examples/assets/main.go
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

	crypto, err := c.Assets.CryptoList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("암호화폐 %d개\n", len(crypto))

	forex, err := c.Assets.ForexList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("외환 페어 %d개\n", len(forex))

	// 시세는 기존 클라이언트로: c.Quote.Quote(ctx, "BTCUSD"), c.Chart.HistoricalPriceEODLight(ctx, "BTCUSD", "", "")
}
