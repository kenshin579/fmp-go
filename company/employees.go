package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// EmployeeCount — 직원 수 (employee-count / historical 공용, SEC 공시 기준)
type EmployeeCount struct {
	Symbol         string `json:"symbol"`         // 종목 심볼
	CIK            string `json:"cik"`            // SEC CIK
	AcceptanceTime string `json:"acceptanceTime"` // 공시 수리 시각
	PeriodOfReport string `json:"periodOfReport"` // 보고 기준일
	CompanyName    string `json:"companyName"`    // 회사명
	FormType       string `json:"formType"`       // 공시 양식 (예: 10-K)
	FilingDate     string `json:"filingDate"`     // 공시일
	EmployeeCount  int64  `json:"employeeCount"`  // 직원 수
	Source         string `json:"source"`         // 원문 URL
}

// EmployeeCount 는 종목의 최신 직원 수를 조회한다.
func (c *Client) EmployeeCount(ctx context.Context, symbol string) ([]EmployeeCount, error) {
	return fetch.ListBySymbol[EmployeeCount](ctx, c.http, "/stable/employee-count", symbol)
}

// HistoricalEmployeeCount 는 종목의 직원 수 시계열을 조회한다.
func (c *Client) HistoricalEmployeeCount(ctx context.Context, symbol string) ([]EmployeeCount, error) {
	return fetch.ListBySymbol[EmployeeCount](ctx, c.http, "/stable/historical-employee-count", symbol)
}
