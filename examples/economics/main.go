// 실행: FMP_API_KEY=... go run examples/economics/main.go
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

	rates, err := c.Economics.TreasuryRates(ctx, "", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(rates) > 0 {
		fmt.Printf("%s 국채: 2년 %.2f%%, 10년 %.2f%%, 30년 %.2f%%\n",
			rates[0].Date, rates[0].Year2, rates[0].Year10, rates[0].Year30)
	}

	rp, err := c.Economics.MarketRiskPremium(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("리스크 프리미엄 국가 %d개\n", len(rp))
}
