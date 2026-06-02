// 실행: FMP_API_KEY=... go run examples/insidertrades/main.go
package main

import (
	"context"
	"fmt"
	"log"

	fmp "github.com/kenshin579/fmp-go"
	"github.com/kenshin579/fmp-go/insidertrades"
)

func main() {
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	trades, err := c.InsiderTrades.SearchInsiderTrades(ctx, insidertrades.SearchParams{Symbol: "AAPL", Limit: 5})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 내부자 거래 %d건\n", len(trades))
	for _, tr := range trades {
		fmt.Printf("  %s %s %s %d주 @ %.2f\n", tr.TransactionDate, tr.ReportingName, tr.TransactionType, tr.SecuritiesTransacted, tr.Price)
	}

	stats, err := c.InsiderTrades.Statistics(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	if len(stats) > 0 {
		fmt.Printf("최근 통계 %dQ%d: 취득 %d / 처분 %d\n", stats[0].Year, stats[0].Quarter, stats[0].TotalAcquired, stats[0].TotalDisposed)
	}
}
