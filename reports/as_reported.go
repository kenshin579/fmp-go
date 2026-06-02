package reports

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// AsReportedStatement — SEC 원문 재무제표 (income/balance/cashflow-as-reported 공용).
// data 는 XBRL/GAAP 태그(소문자) → 값(int/float 혼재) 동적 맵.
type AsReportedStatement struct {
	Symbol           string                 `json:"symbol"`           // 종목 심볼
	FiscalYear       int                    `json:"fiscalYear"`       // 회계연도(숫자)
	Period           string                 `json:"period"`           // 기간 (FY/Q1..)
	ReportedCurrency *string                `json:"reportedCurrency"` // 보고 통화(null 가능)
	Date             string                 `json:"date"`             // 기준일
	Data             map[string]json.Number `json:"data"`             // GAAP 태그 → 수치
}

// AsReportedFull — SEC 원문 전체 재무제표 (financial-statement-full-as-reported).
// data 값이 숫자/문자열/불리언 혼재 → any.
type AsReportedFull struct {
	Symbol           string         `json:"symbol"`           // 종목 심볼
	FiscalYear       int            `json:"fiscalYear"`       // 회계연도(숫자)
	Period           string         `json:"period"`           // 기간
	ReportedCurrency *string        `json:"reportedCurrency"` // 보고 통화(null 가능)
	Date             string         `json:"date"`             // 기준일
	Data             map[string]any `json:"data"`             // GAAP 태그 → 값(혼합 타입)
}

// IncomeStatementAsReported 는 SEC 원문 손익계산서를 조회한다.
func (c *Client) IncomeStatementAsReported(ctx context.Context, symbol, period string, limit int) ([]AsReportedStatement, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[AsReportedStatement](ctx, c.http, "/stable/income-statement-as-reported", asReportedParams(symbol, period, limit))
}

// BalanceSheetStatementAsReported 는 SEC 원문 대차대조표를 조회한다.
func (c *Client) BalanceSheetStatementAsReported(ctx context.Context, symbol, period string, limit int) ([]AsReportedStatement, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[AsReportedStatement](ctx, c.http, "/stable/balance-sheet-statement-as-reported", asReportedParams(symbol, period, limit))
}

// CashFlowStatementAsReported 는 SEC 원문 현금흐름표를 조회한다.
func (c *Client) CashFlowStatementAsReported(ctx context.Context, symbol, period string, limit int) ([]AsReportedStatement, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[AsReportedStatement](ctx, c.http, "/stable/cash-flow-statement-as-reported", asReportedParams(symbol, period, limit))
}

// FinancialStatementFullAsReported 는 SEC 원문 전체 재무제표를 조회한다.
func (c *Client) FinancialStatementFullAsReported(ctx context.Context, symbol, period string, limit int) ([]AsReportedFull, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[AsReportedFull](ctx, c.http, "/stable/financial-statement-full-as-reported", asReportedParams(symbol, period, limit))
}
