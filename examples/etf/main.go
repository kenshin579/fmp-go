// 실행: FMP_API_KEY=... go run examples/etf/main.go
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

	info, err := c.ETF.Information(ctx, "SPY")
	if err != nil {
		log.Fatal(err)
	}
	if len(info) > 0 {
		fmt.Printf("SPY: %s (운용사 %s, AUM %d, 보수율 %.4f)\n", info[0].Name, info[0].ETFCompany, info[0].AssetsUnderManagement, info[0].ExpenseRatio)
	}

	holdings, err := c.ETF.Holdings(ctx, "SPY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SPY 보유 종목 %d개\n", len(holdings))
}
