// 실행: FMP_API_KEY=... go run examples/dcf/main.go
package main

import (
	"context"
	"fmt"
	"log"

	fmp "github.com/kenshin579/fmp-go"
	"github.com/kenshin579/fmp-go/dcf"
)

func main() {
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	v, err := c.DCF.DiscountedCashFlow(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	if len(v) > 0 {
		fmt.Printf("AAPL %s 내재가치 %.2f vs 현재가 %.2f\n", v[0].Date, v[0].DCF, v[0].StockPrice)
	}

	rows, err := c.DCF.CustomDiscountedCashFlow(ctx, dcf.CustomDCFParams{Symbol: "AAPL"})
	if err != nil {
		log.Fatal(err)
	}
	if len(rows) > 0 {
		fmt.Printf("Custom DCF %s년: WACC %.4f, 주당가치 %.2f\n", rows[0].Year, rows[0].WACC, rows[0].EquityValuePerShare)
	}
}
