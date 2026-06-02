// 실행: FMP_API_KEY=... go run examples/analyst/main.go
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

	pt, err := c.Analyst.PriceTargetConsensus(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AAPL 목표가 컨센서스: %.2f (고 %.2f / 저 %.2f)\n", pt.TargetConsensus, pt.TargetHigh, pt.TargetLow)

	g, err := c.Analyst.GradesConsensus(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("등급 컨센서스: %s (매수 %d / 보유 %d / 매도 %d)\n", g.Consensus, g.Buy, g.Hold, g.Sell)
}
