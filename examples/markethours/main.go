// 실행: FMP_API_KEY=... go run examples/markethours/main.go
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

	hours, err := c.MarketHours.ExchangeMarketHours(ctx, "NASDAQ")
	if err != nil {
		log.Fatal(err)
	}
	if len(hours) > 0 {
		fmt.Printf("NASDAQ 개장 여부: %v (%s~%s %s)\n",
			hours[0].IsMarketOpen, hours[0].OpeningHour, hours[0].ClosingHour, hours[0].Timezone)
	}

	all, err := c.MarketHours.AllExchangeMarketHours(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("전체 거래소 %d개\n", len(all))
}
