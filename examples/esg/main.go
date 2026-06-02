// 실행: FMP_API_KEY=... go run examples/esg/main.go
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

	ratings, err := c.ESG.Ratings(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	if len(ratings) > 0 {
		fmt.Printf("AAPL ESG 등급: %s (산업 %s, 순위 %s)\n", ratings[0].ESGRiskRating, ratings[0].Industry, ratings[0].IndustryRank)
	}

	bench, err := c.ESG.Benchmark(ctx, "2023")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("2023 섹터 벤치마크 %d건\n", len(bench))
}
