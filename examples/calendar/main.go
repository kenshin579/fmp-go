// 실행: FMP_API_KEY=... go run examples/calendar/main.go
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

	earnings, err := c.Calendar.EarningsCalendar(ctx, "2025-02-01", "2025-02-07")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("실적 일정 %d건\n", len(earnings))
	for _, e := range earnings[:min(3, len(earnings))] {
		fmt.Printf("  %s %s\n", e.Date, e.Symbol)
	}

	divs, err := c.Calendar.CompanyDividends(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 배당 이력 %d건\n", len(divs))
	if len(divs) > 0 {
		fmt.Printf("  최근 배당락 %s, 배당금 %.4f\n", divs[0].Date, divs[0].Dividend)
	}
}
