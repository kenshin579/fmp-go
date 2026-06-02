// 실행: FMP_API_KEY=... go run examples/transcripts/main.go
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

	latest, err := c.EarningsTranscripts.Latest(ctx, 0, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("최신 트랜스크립트 %d건\n", len(latest))

	dates, err := c.EarningsTranscripts.Dates(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 가용 트랜스크립트 %d건\n", len(dates))
	for _, d := range dates[:min(3, len(dates))] {
		fmt.Printf("  %dQ%d %s\n", d.FiscalYear, d.Quarter, d.Date)
	}
}
