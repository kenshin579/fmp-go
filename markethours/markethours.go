package markethours

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// ExchangeHours — 거래소 운영시간 (exchange-market-hours / all-exchange-market-hours 공유)
type ExchangeHours struct {
	Exchange     string `json:"exchange"`     // 거래소 코드
	Name         string `json:"name"`         // 거래소명
	OpeningHour  string `json:"openingHour"`  // 개장 시각(UTC offset 포함, 예 "09:30 AM -04:00")
	ClosingHour  string `json:"closingHour"`  // 폐장 시각
	Timezone     string `json:"timezone"`     // 타임존(예 America/New_York)
	IsMarketOpen bool   `json:"isMarketOpen"` // 현재 개장 여부
}

// ExchangeHoliday — 거래소 휴장일 (holidays-by-exchange). adj 시각은 null 가능.
type ExchangeHoliday struct {
	Exchange     string  `json:"exchange"`     // 거래소 코드
	Date         string  `json:"date"`         // 일자
	Name         string  `json:"name"`         // 휴일명
	IsClosed     bool    `json:"isClosed"`     // 휴장 여부
	AdjOpenTime  *string `json:"adjOpenTime"`  // 조정 개장 시각(null 가능)
	AdjCloseTime *string `json:"adjCloseTime"` // 조정 폐장 시각(null 가능)
}

// ExchangeMarketHours 는 특정 거래소의 운영시간을 조회한다. exchange 필수.
func (c *Client) ExchangeMarketHours(ctx context.Context, exchange string) ([]ExchangeHours, error) {
	if strings.TrimSpace(exchange) == "" {
		return nil, fmt.Errorf("fmp: exchange must not be empty")
	}
	return fetch.List[ExchangeHours](ctx, c.http, "/stable/exchange-market-hours", map[string]string{"exchange": exchange})
}

// AllExchangeMarketHours 는 전체 거래소 운영시간을 조회한다.
func (c *Client) AllExchangeMarketHours(ctx context.Context) ([]ExchangeHours, error) {
	return fetch.List[ExchangeHours](ctx, c.http, "/stable/all-exchange-market-hours", nil)
}

// HolidaysByExchange 는 특정 거래소의 휴장일을 조회한다. exchange 필수.
func (c *Client) HolidaysByExchange(ctx context.Context, exchange, from, to string) ([]ExchangeHoliday, error) {
	if strings.TrimSpace(exchange) == "" {
		return nil, fmt.Errorf("fmp: exchange must not be empty")
	}
	q := map[string]string{"exchange": exchange}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return fetch.List[ExchangeHoliday](ctx, c.http, "/stable/holidays-by-exchange", q)
}
