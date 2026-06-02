// 실행: FMP_API_KEY=... go run examples/secfilings/main.go
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

	prof, err := c.SECFilings.Profile(ctx, "AAPL", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(prof) > 0 {
		fmt.Printf("%s — CEO %s, 거래소 %s, 직원 %s\n", prof[0].RegistrantName, prof[0].CEO, prof[0].Exchange, prof[0].Employees)
	}

	to := time.Now().Format("2006-01-02")
	from := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	filings, err := c.SECFilings.Latest8K(ctx, from, to, 0, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("최근 7일 8-K 공시 %d건\n", len(filings))
}
