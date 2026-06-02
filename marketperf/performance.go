package marketperf

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SectorPerformance — 섹터 성과 (sector-performance-snapshot / historical-sector-performance 공유)
type SectorPerformance struct {
	Date          string  `json:"date"`          // 일자
	Sector        string  `json:"sector"`        // 섹터명
	Exchange      string  `json:"exchange"`      // 거래소
	AverageChange float64 `json:"averageChange"` // 평균 변동률
}

// IndustryPerformance — 산업 성과 (industry-performance-snapshot / historical-industry-performance 공유)
type IndustryPerformance struct {
	Date          string  `json:"date"`          // 일자
	Industry      string  `json:"industry"`      // 산업명
	Exchange      string  `json:"exchange"`      // 거래소
	AverageChange float64 `json:"averageChange"` // 평균 변동률
}

// SectorPerformanceSnapshot 은 특정 일자의 섹터별 성과를 조회한다. date 필수.
func (c *Client) SectorPerformanceSnapshot(ctx context.Context, date, exchange, sector string) ([]SectorPerformance, error) {
	if strings.TrimSpace(date) == "" {
		return nil, fmt.Errorf("fmp: date must not be empty")
	}
	return fetch.List[SectorPerformance](ctx, c.http, "/stable/sector-performance-snapshot", snapshotParams(date, exchange, "sector", sector))
}

// IndustryPerformanceSnapshot 은 특정 일자의 산업별 성과를 조회한다. date 필수.
func (c *Client) IndustryPerformanceSnapshot(ctx context.Context, date, exchange, industry string) ([]IndustryPerformance, error) {
	if strings.TrimSpace(date) == "" {
		return nil, fmt.Errorf("fmp: date must not be empty")
	}
	return fetch.List[IndustryPerformance](ctx, c.http, "/stable/industry-performance-snapshot", snapshotParams(date, exchange, "industry", industry))
}

// HistoricalSectorPerformance 는 섹터 성과 시계열을 조회한다. sector 필수.
func (c *Client) HistoricalSectorPerformance(ctx context.Context, sector, from, to, exchange string) ([]SectorPerformance, error) {
	if strings.TrimSpace(sector) == "" {
		return nil, fmt.Errorf("fmp: sector must not be empty")
	}
	return fetch.List[SectorPerformance](ctx, c.http, "/stable/historical-sector-performance", historicalParams("sector", sector, from, to, exchange))
}

// HistoricalIndustryPerformance 는 산업 성과 시계열을 조회한다. industry 필수.
func (c *Client) HistoricalIndustryPerformance(ctx context.Context, industry, from, to, exchange string) ([]IndustryPerformance, error) {
	if strings.TrimSpace(industry) == "" {
		return nil, fmt.Errorf("fmp: industry must not be empty")
	}
	return fetch.List[IndustryPerformance](ctx, c.http, "/stable/historical-industry-performance", historicalParams("industry", industry, from, to, exchange))
}
