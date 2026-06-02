package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// CompanyNote — 회사 채권/노트 발행 정보
type CompanyNote struct {
	CIK      string `json:"cik"`      // SEC CIK
	Symbol   string `json:"symbol"`   // 종목 심볼
	Title    string `json:"title"`    // 노트명 (예: 0.000% Notes due 2025)
	Exchange string `json:"exchange"` // 거래소
}

// CompanyNotes 는 회사가 발행한 채권/노트 목록을 조회한다.
func (c *Client) CompanyNotes(ctx context.Context, symbol string) ([]CompanyNote, error) {
	return fetch.ListBySymbol[CompanyNote](ctx, c.http, "/stable/company-notes", symbol)
}
