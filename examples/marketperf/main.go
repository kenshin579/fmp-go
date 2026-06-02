// 실행: FMP_API_KEY=... go run examples/marketperf/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	fmp "github.com/kenshin579/fmp-go"
)

func main() {
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	gainers, err := c.MarketPerformance.BiggestGainers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("상승률 상위:")
	for _, g := range gainers[:min(3, len(gainers))] {
		fmt.Printf("  %s %s: %.2f%%\n", g.Symbol, g.Name, g.ChangesPercentage)
	}

	date := time.Now().Format("2006-01-02")
	sectors, err := c.MarketPerformance.SectorPerformanceSnapshot(ctx, date, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s 섹터 성과 %d건\n", date, len(sectors))
}
