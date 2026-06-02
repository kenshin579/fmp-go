// 실행: FMP_API_KEY=... go run examples/cot/main.go
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

	list, err := c.COT.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("COT 보고 대상 상품 %d개\n", len(list))

	analysis, err := c.COT.Analysis(ctx, "ES", "", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(analysis) > 0 {
		a := analysis[0]
		fmt.Printf("ES %s: %s (순포지션 %d, 심리 %s)\n", a.Date, a.MarketSituation, a.NetPostion, a.MarketSentiment)
	}
}
