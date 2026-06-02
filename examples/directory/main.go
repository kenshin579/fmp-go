// 실행: FMP_API_KEY=... go run examples/directory/main.go
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

	exchanges, err := c.Directory.AvailableExchanges(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("지원 거래소 %d개\n", len(exchanges))

	sectors, err := c.Directory.AvailableSectors(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("섹터: ")
	for _, s := range sectors {
		fmt.Printf("%s ", s.Sector)
	}
	fmt.Println()
}
