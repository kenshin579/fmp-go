// 실행: FMP_API_KEY=... go run examples/form13f/main.go
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

	filings, err := c.Form13F.LatestFilings(ctx, 0, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("최신 13F 공시 %d건\n", len(filings))

	pos, err := c.Form13F.PositionsSummary(ctx, "AAPL", "2023", "3")
	if err != nil {
		log.Fatal(err)
	}
	if len(pos) > 0 {
		fmt.Printf("AAPL 2023 Q3 보유 기관 %d개 (총 투자 %d)\n", pos[0].InvestorsHolding, pos[0].TotalInvested)
	}
}
