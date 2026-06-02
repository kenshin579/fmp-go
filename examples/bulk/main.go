// 실행: FMP_API_KEY=... go run examples/bulk/main.go
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

	scores, err := c.Bulk.Scores(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("scores-bulk CSV: %d bytes\n", len(scores))

	date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	eod, err := c.Bulk.EOD(ctx, date)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("eod-bulk(%s) CSV: %d bytes\n", date, len(eod))
	// 반환은 원시 CSV — encoding/csv 로 파싱하세요.
}
