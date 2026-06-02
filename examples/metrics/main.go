// 실행: FMP_API_KEY=... go run examples/metrics/main.go
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

	km, err := c.Metrics.KeyMetrics(ctx, "AAPL", "annual", 1)
	if err != nil {
		log.Fatal(err)
	}
	if len(km) > 0 {
		fmt.Printf("AAPL %s 핵심지표: ROE %.4f, FCF수익률 %.4f, EV/EBITDA %.2f\n",
			km[0].FiscalYear, km[0].ReturnOnEquity, km[0].FreeCashFlowYield, km[0].EvToEBITDA)
	}

	s, err := c.Metrics.FinancialScores(ctx, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("재무점수: Altman Z %.2f, Piotroski %d/9\n", s.AltmanZScore, s.PiotroskiScore)
}
