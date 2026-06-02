package marketperf

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SectorPE — 섹터 PER (sector-pe-snapshot / historical-sector-pe 공유)
type SectorPE struct {
	Date     string  `json:"date"`     // 일자
	Sector   string  `json:"sector"`   // 섹터명
	Exchange string  `json:"exchange"` // 거래소
	PE       float64 `json:"pe"`       // 섹터 평균 PER
}

// IndustryPE — 산업 PER (industry-pe-snapshot / historical-industry-pe 공유)
type IndustryPE struct {
	Date     string  `json:"date"`     // 일자
	Industry string  `json:"industry"` // 산업명
	Exchange string  `json:"exchange"` // 거래소
	PE       float64 `json:"pe"`       // 산업 평균 PER
}

// SectorPESnapshot 은 특정 일자의 섹터별 PER 을 조회한다. date 필수.
func (c *Client) SectorPESnapshot(ctx context.Context, date, exchange, sector string) ([]SectorPE, error) {
	if strings.TrimSpace(date) == "" {
		return nil, fmt.Errorf("fmp: date must not be empty")
	}
	return fetch.List[SectorPE](ctx, c.http, "/stable/sector-pe-snapshot", snapshotParams(date, exchange, "sector", sector))
}

// IndustryPESnapshot 은 특정 일자의 산업별 PER 을 조회한다. date 필수.
func (c *Client) IndustryPESnapshot(ctx context.Context, date, exchange, industry string) ([]IndustryPE, error) {
	if strings.TrimSpace(date) == "" {
		return nil, fmt.Errorf("fmp: date must not be empty")
	}
	return fetch.List[IndustryPE](ctx, c.http, "/stable/industry-pe-snapshot", snapshotParams(date, exchange, "industry", industry))
}

// HistoricalSectorPE 는 섹터 PER 시계열을 조회한다. sector 필수.
func (c *Client) HistoricalSectorPE(ctx context.Context, sector, from, to, exchange string) ([]SectorPE, error) {
	if strings.TrimSpace(sector) == "" {
		return nil, fmt.Errorf("fmp: sector must not be empty")
	}
	return fetch.List[SectorPE](ctx, c.http, "/stable/historical-sector-pe", historicalParams("sector", sector, from, to, exchange))
}

// HistoricalIndustryPE 는 산업 PER 시계열을 조회한다. industry 필수.
func (c *Client) HistoricalIndustryPE(ctx context.Context, industry, from, to, exchange string) ([]IndustryPE, error) {
	if strings.TrimSpace(industry) == "" {
		return nil, fmt.Errorf("fmp: industry must not be empty")
	}
	return fetch.List[IndustryPE](ctx, c.http, "/stable/historical-industry-pe", historicalParams("industry", industry, from, to, exchange))
}
