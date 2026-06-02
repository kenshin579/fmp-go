// 실행: FMP_API_KEY=... go run examples/chart/main.go
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

	eod, err := c.Chart.HistoricalPriceEODLight(ctx, "AAPL", "", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(eod) > 0 {
		fmt.Printf("AAPL 최신 EOD %s: 종가 %.2f, 거래량 %d\n", eod[0].Date, eod[0].Price, eod[0].Volume)
	}

	bars, err := c.Chart.Intraday1Hour(ctx, "AAPL", "", "", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 1시간봉 %d개\n", len(bars))
}
