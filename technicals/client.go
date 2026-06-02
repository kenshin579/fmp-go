// Package technicals 는 FMP 기술 지표 API sub-client (SMA/EMA/RSI/ADX 등).
// fmp.Client.TechnicalIndicators 로 접근.
package technicals

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 기술 지표 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

func validate(symbol string, periodLength int, timeframe string) error {
	if strings.TrimSpace(symbol) == "" {
		return fmt.Errorf("fmp: symbol must not be empty")
	}
	if periodLength <= 0 {
		return fmt.Errorf("fmp: periodLength must be > 0")
	}
	if strings.TrimSpace(timeframe) == "" {
		return fmt.Errorf("fmp: timeframe must not be empty")
	}
	return nil
}

func indicatorParams(symbol string, periodLength int, timeframe, from, to string) map[string]string {
	q := map[string]string{
		"symbol":       symbol,
		"periodLength": strconv.Itoa(periodLength),
		"timeframe":    timeframe,
	}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return q
}
