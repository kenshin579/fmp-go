package calendar

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// IPO — IPO 일정 (ipos-calendar). shares/priceRange/marketCap 결측 가능.
type IPO struct {
	Symbol     string  `json:"symbol"`     // 종목 심볼
	Date       string  `json:"date"`       // IPO 일자
	Daa        string  `json:"daa"`        // 공시 일시 (ISO8601)
	Company    string  `json:"company"`    // 회사명
	Exchange   string  `json:"exchange"`   // 거래소
	Actions    string  `json:"actions"`    // 상태 (예: Expected)
	Shares     *int64  `json:"shares"`     // 공모 주식 수(결측 가능)
	PriceRange *string `json:"priceRange"` // 공모가 범위(결측 가능)
	MarketCap  *int64  `json:"marketCap"`  // 시가총액(결측 가능)
}

// IPODisclosure — IPO 공시 서류 (ipos-disclosure)
type IPODisclosure struct {
	Symbol            string `json:"symbol"`            // 종목 심볼
	FilingDate        string `json:"filingDate"`        // 제출일
	AcceptedDate      string `json:"acceptedDate"`      // 수리일
	EffectivenessDate string `json:"effectivenessDate"` // 효력 발생일
	CIK               string `json:"cik"`               // SEC CIK
	Form              string `json:"form"`              // 공시 양식 (예: CERT)
	URL               string `json:"url"`               // 원문 URL
}

// IPOProspectus — IPO 투자설명서 (ipos-prospectus)
type IPOProspectus struct {
	Symbol                          string  `json:"symbol"`                          // 종목 심볼
	AcceptedDate                    string  `json:"acceptedDate"`                    // 수리일
	FilingDate                      string  `json:"filingDate"`                      // 제출일
	IPODate                         string  `json:"ipoDate"`                         // IPO 일자
	CIK                             string  `json:"cik"`                             // SEC CIK
	PricePublicPerShare             float64 `json:"pricePublicPerShare"`             // 주당 공모가
	PricePublicTotal                float64 `json:"pricePublicTotal"`                // 총 공모금액
	DiscountsAndCommissionsPerShare float64 `json:"discountsAndCommissionsPerShare"` // 주당 인수수수료
	DiscountsAndCommissionsTotal    float64 `json:"discountsAndCommissionsTotal"`    // 총 인수수수료
	ProceedsBeforeExpensesPerShare  float64 `json:"proceedsBeforeExpensesPerShare"`  // 주당 순수취금(비용 전)
	ProceedsBeforeExpensesTotal     float64 `json:"proceedsBeforeExpensesTotal"`     // 총 순수취금(비용 전)
	Form                            string  `json:"form"`                            // 공시 양식 (예: 424B4)
	URL                             string  `json:"url"`                             // 원문 URL
}

// IPOsCalendar 는 기간 내 IPO 일정을 조회한다. from/to 는 YYYY-MM-DD(선택).
func (c *Client) IPOsCalendar(ctx context.Context, from, to string) ([]IPO, error) {
	return fetch.List[IPO](ctx, c.http, "/stable/ipos-calendar", dateRange(from, to))
}

// IPODisclosures 는 기간 내 IPO 공시 서류를 조회한다.
func (c *Client) IPODisclosures(ctx context.Context, from, to string) ([]IPODisclosure, error) {
	return fetch.List[IPODisclosure](ctx, c.http, "/stable/ipos-disclosure", dateRange(from, to))
}

// IPOProspectuses 는 기간 내 IPO 투자설명서를 조회한다.
func (c *Client) IPOProspectuses(ctx context.Context, from, to string) ([]IPOProspectus, error) {
	return fetch.List[IPOProspectus](ctx, c.http, "/stable/ipos-prospectus", dateRange(from, to))
}
