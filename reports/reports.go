package reports

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// LatestFinancialStatement — 최신 재무제표 등록 목록 (latest-financial-statements).
// symbol 미입력, page/limit 페이징. calendarYear 사용(fiscalYear 아님).
type LatestFinancialStatement struct {
	Symbol       string `json:"symbol"`       // 종목 심볼
	CalendarYear int    `json:"calendarYear"` // 달력연도
	Period       string `json:"period"`       // 기간 (Q4 등)
	Date         string `json:"date"`         // 보고 기준일
	DateAdded    string `json:"dateAdded"`    // 등록 일시
}

// FinancialReportDate — 보고서 다운로드 링크 (financial-reports-dates).
type FinancialReportDate struct {
	Symbol     string `json:"symbol"`     // 종목 심볼
	FiscalYear int    `json:"fiscalYear"` // 회계연도
	Period     string `json:"period"`     // 기간 (FY/Q1..)
	LinkXlsx   string `json:"linkXlsx"`   // XLSX 다운로드 URL
	LinkJson   string `json:"linkJson"`   // JSON 다운로드 URL
}

// LatestFinancialStatements 는 최신 등록된 재무제표 목록을 페이지 단위로 조회한다.
func (c *Client) LatestFinancialStatements(ctx context.Context, page, limit int) ([]LatestFinancialStatement, error) {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[LatestFinancialStatement](ctx, c.http, "/stable/latest-financial-statements", q)
}

// FinancialReportDates 는 종목의 보고서 다운로드 링크 목록을 조회한다.
func (c *Client) FinancialReportDates(ctx context.Context, symbol string) ([]FinancialReportDate, error) {
	return fetch.ListBySymbol[FinancialReportDate](ctx, c.http, "/stable/financial-reports-dates", symbol)
}

// FinancialReportJSON 은 종목의 특정 연도/기간 10-K 보고서를 원시 JSON(동적 섹션)으로 조회한다.
func (c *Client) FinancialReportJSON(ctx context.Context, symbol string, year int, period string) ([]map[string]any, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := map[string]string{"symbol": symbol, "year": strconv.Itoa(year), "period": period}
	return fetch.List[map[string]any](ctx, c.http, "/stable/financial-reports-json", q)
}
