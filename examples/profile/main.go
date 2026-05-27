// examples/profile — FMP 회사 프로필 조회 예제.
//
// 실행: FMP_API_KEY=... go run ./examples/profile AAPL
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	fmp "github.com/kenshin579/fmp-go"
)

func main() {
	symbol := "AAPL"
	if len(os.Args) > 1 {
		symbol = os.Args[1]
	}
	client, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatalf("NewClientFromEnv: %v", err)
	}
	p, err := client.Company.Profile(context.Background(), symbol)
	if err != nil {
		log.Fatalf("Profile(%s): %v", symbol, err)
	}
	fmt.Printf("%s (%s)\nCEO: %s\nIPO: %s\nWeb: %s\nSector: %s / %s\n%s\n",
		p.CompanyName, p.Symbol, p.CEO, p.IPODate, p.Website, p.Sector, p.Industry, p.Description)
}
