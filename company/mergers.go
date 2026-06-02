package company

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// MergerAcquisition — M&A 공시 (latest / search 공용)
type MergerAcquisition struct {
	Symbol              string `json:"symbol"`              // 인수 회사 심볼
	CompanyName         string `json:"companyName"`         // 인수 회사명
	CIK                 string `json:"cik"`                 // 인수 회사 CIK
	TargetedCompanyName string `json:"targetedCompanyName"` // 피인수 회사명
	TargetedCik         string `json:"targetedCik"`         // 피인수 회사 CIK
	TargetedSymbol      string `json:"targetedSymbol"`      // 피인수 회사 심볼
	TransactionDate     string `json:"transactionDate"`     // 거래일
	AcceptedDate        string `json:"acceptedDate"`        // 수리일시
	Link                string `json:"link"`                // 원문 URL
}

// LatestMergersAcquisitions 는 최신 M&A 공시를 페이지 단위로 조회한다.
func (c *Client) LatestMergersAcquisitions(ctx context.Context, page int) ([]MergerAcquisition, error) {
	return fetch.List[MergerAcquisition](ctx, c.http, "/stable/mergers-acquisitions-latest", map[string]string{"page": strconv.Itoa(page)})
}

// SearchMergersAcquisitions 는 회사명으로 M&A 공시를 검색한다.
func (c *Client) SearchMergersAcquisitions(ctx context.Context, name string) ([]MergerAcquisition, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[MergerAcquisition](ctx, c.http, "/stable/mergers-acquisitions-search", map[string]string{"name": name})
}
