package company

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Executive — 주요 임원 (key-executives). pay/yearBorn/titleSince 는 결측 가능(null) → 포인터.
type Executive struct {
	Title       string `json:"title"`       // 직책
	Name        string `json:"name"`        // 이름
	Pay         *int64 `json:"pay"`         // 보수(결측 가능)
	CurrencyPay string `json:"currencyPay"` // 보수 통화
	Gender      string `json:"gender"`      // 성별
	YearBorn    *int   `json:"yearBorn"`    // 출생연도(결측 가능)
	TitleSince  *int64 `json:"titleSince"`  // 현 직책 부임(결측 가능)
	Active      bool   `json:"active"`      // 재직 여부
}

// ExecutiveCompensation — 임원 보수 공시 (governance-executive-compensation)
type ExecutiveCompensation struct {
	CIK                       string `json:"cik"`                       // SEC CIK
	Symbol                    string `json:"symbol"`                    // 종목 심볼
	CompanyName               string `json:"companyName"`               // 회사명
	FilingDate                string `json:"filingDate"`                // 공시일
	AcceptedDate              string `json:"acceptedDate"`              // 수리일시
	NameAndPosition           string `json:"nameAndPosition"`           // 이름·직책
	Year                      int    `json:"year"`                      // 회계연도
	Salary                    int64  `json:"salary"`                    // 급여
	Bonus                     int64  `json:"bonus"`                     // 상여
	StockAward                int64  `json:"stockAward"`                // 주식 보상
	OptionAward               int64  `json:"optionAward"`               // 옵션 보상
	IncentivePlanCompensation int64  `json:"incentivePlanCompensation"` // 성과급
	AllOtherCompensation      int64  `json:"allOtherCompensation"`      // 기타 보상
	Total                     int64  `json:"total"`                     // 총 보수
}

// ExecutiveCompensationBenchmark — 산업별 평균 임원 보수
type ExecutiveCompensationBenchmark struct {
	IndustryTitle       string  `json:"industryTitle"`       // 산업 분류
	Year                int     `json:"year"`                // 연도
	AverageCompensation float64 `json:"averageCompensation"` // 평균 보수
}

// KeyExecutives 는 종목의 주요 임원 목록을 조회한다.
func (c *Client) KeyExecutives(ctx context.Context, symbol string) ([]Executive, error) {
	return fetch.ListBySymbol[Executive](ctx, c.http, "/stable/key-executives", symbol)
}

// ExecutiveCompensation 은 종목의 임원 보수 공시를 조회한다.
func (c *Client) ExecutiveCompensation(ctx context.Context, symbol string) ([]ExecutiveCompensation, error) {
	return fetch.ListBySymbol[ExecutiveCompensation](ctx, c.http, "/stable/governance-executive-compensation", symbol)
}

// ExecutiveCompensationBenchmark 은 연도별 산업 평균 임원 보수를 조회한다.
func (c *Client) ExecutiveCompensationBenchmark(ctx context.Context, year int) ([]ExecutiveCompensationBenchmark, error) {
	params := map[string]string{}
	if year > 0 {
		params["year"] = strconv.Itoa(year)
	}
	return fetch.List[ExecutiveCompensationBenchmark](ctx, c.http, "/stable/executive-compensation-benchmark", params)
}
