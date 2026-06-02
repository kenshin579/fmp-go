// 실행: FMP_API_KEY=... go run examples/reports/main.go
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

	rows, err := c.Reports.IncomeStatementAsReported(ctx, "AAPL", "annual", 1)
	if err != nil {
		log.Fatal(err)
	}
	if len(rows) > 0 {
		fmt.Printf("AAPL %d %s SEC 원문 손익계산서: data 항목 %d개\n", rows[0].FiscalYear, rows[0].Period, len(rows[0].Data))
	}

	dates, err := c.Reports.FinancialReportDates(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	if len(dates) > 0 {
		fmt.Printf("최신 보고서 %d %s — JSON: %s\n", dates[0].FiscalYear, dates[0].Period, dates[0].LinkJson)
	}
}
