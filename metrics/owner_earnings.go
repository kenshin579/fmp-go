package metrics

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// OwnerEarning — 오너 어닝 (owner-earnings). 버핏식 소유주 이익.
type OwnerEarning struct {
	Symbol                 string  `json:"symbol"`                 // 종목 심볼
	ReportedCurrency       string  `json:"reportedCurrency"`       // 보고 통화
	FiscalYear             string  `json:"fiscalYear"`             // 회계연도(문자열)
	Period                 string  `json:"period"`                 // 기간 (FY/Q1..)
	Date                   string  `json:"date"`                   // 기준일
	AveragePPE             float64 `json:"averagePPE"`             // 평균 유형자산 비율
	MaintenanceCapex       int64   `json:"maintenanceCapex"`       // 유지보수 capex
	OwnersEarnings         int64   `json:"ownersEarnings"`         // 소유주 이익
	GrowthCapex            int64   `json:"growthCapex"`            // 성장 capex
	OwnersEarningsPerShare float64 `json:"ownersEarningsPerShare"` // 주당 소유주 이익
}

// OwnerEarnings 는 종목의 오너 어닝 시계열을 조회한다.
func (c *Client) OwnerEarnings(ctx context.Context, symbol string, limit int) ([]OwnerEarning, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[OwnerEarning](ctx, c.http, "/stable/owner-earnings", listParams(symbol, "", limit))
}
