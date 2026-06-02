package esg

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// ESGRating — ESG 등급 (esg-ratings)
type ESGRating struct {
	Symbol        string `json:"symbol"`        // 종목 심볼
	CIK           string `json:"cik"`           // SEC CIK(0-padded)
	CompanyName   string `json:"companyName"`   // 회사명
	Industry      string `json:"industry"`      // 산업
	FiscalYear    int    `json:"fiscalYear"`    // 회계연도
	ESGRiskRating string `json:"ESGRiskRating"` // ESG 리스크 등급(예: B)
	IndustryRank  string `json:"industryRank"`  // 산업 내 순위("4 out of 5")
}

// ESGDisclosure — ESG 공시 (esg-disclosures)
type ESGDisclosure struct {
	Date               string  `json:"date"`               // 공시일
	AcceptedDate       string  `json:"acceptedDate"`       // 수리일
	Symbol             string  `json:"symbol"`             // 종목 심볼
	CIK                string  `json:"cik"`                // SEC CIK(0-padded)
	CompanyName        string  `json:"companyName"`        // 회사명
	FormType           string  `json:"formType"`           // 공시 양식(8-K 등)
	EnvironmentalScore float64 `json:"environmentalScore"` // 환경 점수
	SocialScore        float64 `json:"socialScore"`        // 사회 점수
	GovernanceScore    float64 `json:"governanceScore"`    // 지배구조 점수
	ESGScore           float64 `json:"ESGScore"`           // 종합 ESG 점수
	URL                string  `json:"url"`                // 공시 원문 URL
}

// ESGBenchmark — 섹터별 ESG 벤치마크 (esg-benchmark)
type ESGBenchmark struct {
	FiscalYear         int     `json:"fiscalYear"`         // 회계연도
	Sector             string  `json:"sector"`             // 섹터
	EnvironmentalScore float64 `json:"environmentalScore"` // 환경 점수
	SocialScore        float64 `json:"socialScore"`        // 사회 점수
	GovernanceScore    float64 `json:"governanceScore"`    // 지배구조 점수
	ESGScore           float64 `json:"ESGScore"`           // 종합 ESG 점수
}

// Ratings 는 종목의 ESG 등급을 조회한다.
func (c *Client) Ratings(ctx context.Context, symbol string) ([]ESGRating, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[ESGRating](ctx, c.http, "/stable/esg-ratings", map[string]string{"symbol": symbol})
}

// Disclosures 는 종목의 ESG 공시를 조회한다.
func (c *Client) Disclosures(ctx context.Context, symbol string) ([]ESGDisclosure, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[ESGDisclosure](ctx, c.http, "/stable/esg-disclosures", map[string]string{"symbol": symbol})
}

// Benchmark 는 연도별 섹터 ESG 벤치마크를 조회한다. year 필수.
func (c *Client) Benchmark(ctx context.Context, year string) ([]ESGBenchmark, error) {
	if strings.TrimSpace(year) == "" {
		return nil, fmt.Errorf("fmp: year must not be empty")
	}
	return fetch.List[ESGBenchmark](ctx, c.http, "/stable/esg-benchmark", map[string]string{"year": year})
}
