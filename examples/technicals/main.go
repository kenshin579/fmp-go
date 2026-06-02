// 실행: FMP_API_KEY=... go run examples/technicals/main.go
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

	sma, err := c.TechnicalIndicators.SMA(ctx, "AAPL", 10, "1day", "", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(sma) > 0 {
		fmt.Printf("AAPL %s SMA(10): %.3f (종가 %.2f)\n", sma[0].Date, sma[0].SMA, sma[0].Close)
	}

	rsi, err := c.TechnicalIndicators.RSI(ctx, "AAPL", 14, "1day", "", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(rsi) > 0 {
		fmt.Printf("AAPL %s RSI(14): %.2f\n", rsi[0].Date, rsi[0].RSI)
	}
}
