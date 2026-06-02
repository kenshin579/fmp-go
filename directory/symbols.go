package directory

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// SymbolName — 심볼+이름 (etf-list / actively-trading-list 공유)
type SymbolName struct {
	Symbol string `json:"symbol"` // 종목 심볼
	Name   string `json:"name"`   // 종목명
}

// CompanySymbol — 회사 심볼 목록 (stock-list)
type CompanySymbol struct {
	Symbol      string `json:"symbol"`      // 종목 심볼
	CompanyName string `json:"companyName"` // 회사명
}

// FinancialSymbol — 재무제표 제공 심볼 (financial-statement-symbol-list)
type FinancialSymbol struct {
	Symbol            string `json:"symbol"`            // 종목 심볼
	CompanyName       string `json:"companyName"`       // 회사명
	TradingCurrency   string `json:"tradingCurrency"`   // 거래 통화
	ReportingCurrency string `json:"reportingCurrency"` // 보고 통화
}

// CIKEntry — CIK 목록 (cik-list). cik 0-padded 문자열.
type CIKEntry struct {
	CIK         string `json:"cik"`         // SEC CIK(0-padded)
	CompanyName string `json:"companyName"` // 회사명
}

// SymbolChange — 심볼 변경 이력 (symbol-change)
type SymbolChange struct {
	Date        string `json:"date"`        // 변경일
	CompanyName string `json:"companyName"` // 회사명
	OldSymbol   string `json:"oldSymbol"`   // 이전 심볼
	NewSymbol   string `json:"newSymbol"`   // 신규 심볼
}

// TranscriptSymbol — 실적 트랜스크립트 보유 심볼 (earnings-transcript-list).
// noOfTranscripts 는 FMP 가 문자열로 반환.
type TranscriptSymbol struct {
	Symbol          string `json:"symbol"`          // 종목 심볼
	CompanyName     string `json:"companyName"`     // 회사명
	NoOfTranscripts string `json:"noOfTranscripts"` // 트랜스크립트 수(문자열)
}

// CompanySymbolsList 는 전체 회사 심볼 목록을 조회한다.
func (c *Client) CompanySymbolsList(ctx context.Context) ([]CompanySymbol, error) {
	return fetch.List[CompanySymbol](ctx, c.http, "/stable/stock-list", nil)
}

// FinancialSymbolsList 는 재무제표 제공 심볼 목록을 조회한다.
func (c *Client) FinancialSymbolsList(ctx context.Context) ([]FinancialSymbol, error) {
	return fetch.List[FinancialSymbol](ctx, c.http, "/stable/financial-statement-symbol-list", nil)
}

// CIKList 는 CIK 목록을 페이지 단위로 조회한다.
func (c *Client) CIKList(ctx context.Context, page, limit int) ([]CIKEntry, error) {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[CIKEntry](ctx, c.http, "/stable/cik-list", q)
}

// SymbolChangesList 는 심볼 변경 이력을 조회한다.
func (c *Client) SymbolChangesList(ctx context.Context, invalid bool, limit int) ([]SymbolChange, error) {
	q := map[string]string{"invalid": strconv.FormatBool(invalid)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[SymbolChange](ctx, c.http, "/stable/symbol-change", q)
}

// ETFsList 는 ETF 목록을 조회한다.
func (c *Client) ETFsList(ctx context.Context) ([]SymbolName, error) {
	return fetch.List[SymbolName](ctx, c.http, "/stable/etf-list", nil)
}

// ActivelyTradingList 는 활발히 거래되는 종목 목록을 조회한다.
func (c *Client) ActivelyTradingList(ctx context.Context) ([]SymbolName, error) {
	return fetch.List[SymbolName](ctx, c.http, "/stable/actively-trading-list", nil)
}

// EarningsTranscriptList 는 실적 트랜스크립트 보유 심볼 목록을 조회한다.
func (c *Client) EarningsTranscriptList(ctx context.Context) ([]TranscriptSymbol, error) {
	return fetch.List[TranscriptSymbol](ctx, c.http, "/stable/earnings-transcript-list", nil)
}
