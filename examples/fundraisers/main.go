// 실행: FMP_API_KEY=... go run examples/fundraisers/main.go
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

	cf, err := c.Fundraisers.LatestCrowdfunding(ctx, 0, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("최신 크라우드펀딩 %d건\n", len(cf))

	eo, err := c.Fundraisers.LatestEquityOffering(ctx, 0, 5, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("최신 지분공모 %d건\n", len(eo))
}
