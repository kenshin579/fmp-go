package insidertrades

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// InsiderTrade — 내부자 거래 (insider-trading/latest, /search 공유)
type InsiderTrade struct {
	Symbol                   string  `json:"symbol"`                   // 종목 심볼
	FilingDate               string  `json:"filingDate"`               // 공시일
	TransactionDate          string  `json:"transactionDate"`          // 거래일
	ReportingCik             string  `json:"reportingCik"`             // 보고자 CIK
	CompanyCik               string  `json:"companyCik"`               // 회사 CIK
	TransactionType          string  `json:"transactionType"`          // 거래 유형(예: P-Purchase)
	SecuritiesOwned          int64   `json:"securitiesOwned"`          // 보유 증권 수
	ReportingName            string  `json:"reportingName"`            // 보고자명
	TypeOfOwner              string  `json:"typeOfOwner"`              // 소유자 유형
	AcquisitionOrDisposition string  `json:"acquisitionOrDisposition"` // 취득/처분(A/D)
	DirectOrIndirect         string  `json:"directOrIndirect"`         // 직접/간접
	FormType                 string  `json:"formType"`                 // 양식 유형(예: 4)
	SecuritiesTransacted     int64   `json:"securitiesTransacted"`     // 거래 증권 수
	Price                    float64 `json:"price"`                    // 단가
	SecurityName             string  `json:"securityName"`             // 증권명
	URL                      string  `json:"url"`                      // 공시 URL
}

// SearchParams — SearchInsiderTrades 옵션. 빈 값은 쿼리에서 제외.
type SearchParams struct {
	Symbol          string
	Page            int
	Limit           int
	ReportingCik    string
	CompanyCik      string
	TransactionType string
}

func (p SearchParams) queryParams() map[string]string {
	q := map[string]string{"page": strconv.Itoa(p.Page)}
	if p.Limit > 0 {
		q["limit"] = strconv.Itoa(p.Limit)
	}
	if p.Symbol != "" {
		q["symbol"] = p.Symbol
	}
	if p.ReportingCik != "" {
		q["reportingCik"] = p.ReportingCik
	}
	if p.CompanyCik != "" {
		q["companyCik"] = p.CompanyCik
	}
	if p.TransactionType != "" {
		q["transactionType"] = p.TransactionType
	}
	return q
}

// LatestInsiderTrades 는 최신 내부자 거래를 조회한다.
func (c *Client) LatestInsiderTrades(ctx context.Context, date string, page, limit int) ([]InsiderTrade, error) {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	if date != "" {
		q["date"] = date
	}
	return fetch.List[InsiderTrade](ctx, c.http, "/stable/insider-trading/latest", q)
}

// SearchInsiderTrades 는 조건으로 내부자 거래를 검색한다.
func (c *Client) SearchInsiderTrades(ctx context.Context, p SearchParams) ([]InsiderTrade, error) {
	return fetch.List[InsiderTrade](ctx, c.http, "/stable/insider-trading/search", p.queryParams())
}
