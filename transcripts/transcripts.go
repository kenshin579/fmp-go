package transcripts

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// EarningCallTranscript — 실적 발표 트랜스크립트 본문 (earning-call-transcript)
type EarningCallTranscript struct {
	Symbol  string `json:"symbol"`  // 종목 심볼
	Period  string `json:"period"`  // 분기(예: Q3)
	Year    int    `json:"year"`    // 연도
	Date    string `json:"date"`    // 발표일
	Content string `json:"content"` // 전문(텍스트)
}

// LatestEarningCallTranscript — 최신 트랜스크립트 목록 (earning-call-transcript-latest)
type LatestEarningCallTranscript struct {
	Symbol     string `json:"symbol"`     // 종목 심볼
	Period     string `json:"period"`     // 분기(예: Q3)
	FiscalYear int    `json:"fiscalYear"` // 회계연도
	Date       string `json:"date"`       // 발표일
}

// EarningCallTranscriptDate — 종목별 트랜스크립트 가용 일자 (earning-call-transcript-dates)
type EarningCallTranscriptDate struct {
	Quarter    int    `json:"quarter"`    // 분기(숫자 1~4)
	FiscalYear int    `json:"fiscalYear"` // 회계연도
	Date       string `json:"date"`       // 발표일
}

// Transcript 는 종목의 특정 연도/분기 실적 발표 전문을 조회한다. symbol/year/quarter 필수.
func (c *Client) Transcript(ctx context.Context, symbol, year, quarter string, limit int) ([]EarningCallTranscript, error) {
	if strings.TrimSpace(symbol) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: symbol, year, quarter must not be empty")
	}
	q := map[string]string{"symbol": symbol, "year": year, "quarter": quarter}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[EarningCallTranscript](ctx, c.http, "/stable/earning-call-transcript", q)
}

// Latest 는 최신 등록된 트랜스크립트 목록을 조회한다.
func (c *Client) Latest(ctx context.Context, page, limit int) ([]LatestEarningCallTranscript, error) {
	return fetch.List[LatestEarningCallTranscript](ctx, c.http, "/stable/earning-call-transcript-latest", pageParams(page, limit))
}

// Dates 는 종목의 트랜스크립트 가용 일자 목록을 조회한다. symbol 필수.
func (c *Client) Dates(ctx context.Context, symbol string) ([]EarningCallTranscriptDate, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EarningCallTranscriptDate](ctx, c.http, "/stable/earning-call-transcript-dates", map[string]string{"symbol": symbol})
}
