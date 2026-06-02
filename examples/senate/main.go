// 실행: FMP_API_KEY=... go run examples/senate/main.go
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

	senate, err := c.Senate.SenateLatest(ctx, 0, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("최신 상원 거래:")
	for _, t := range senate {
		fmt.Printf("  %s %s %s %s %s (%s)\n", t.TransactionDate, t.Office, t.Type, t.Symbol, t.Amount, t.AssetType)
	}

	house, err := c.Senate.HouseLatest(ctx, 0, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("최신 하원 거래 %d건\n", len(house))
}
