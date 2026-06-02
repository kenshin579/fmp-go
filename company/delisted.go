package company

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// DelistedCompany — 상장폐지 종목
type DelistedCompany struct {
	Symbol       string `json:"symbol"`       // 종목 심볼
	CompanyName  string `json:"companyName"`  // 회사명
	Exchange     string `json:"exchange"`     // 거래소
	IPODate      string `json:"ipoDate"`      // 상장일
	DelistedDate string `json:"delistedDate"` // 상장폐지일
}

// DelistedCompanies 는 상장폐지 종목을 페이지 단위로 조회한다.
func (c *Client) DelistedCompanies(ctx context.Context, page int) ([]DelistedCompany, error) {
	return fetch.List[DelistedCompany](ctx, c.http, "/stable/delisted-companies", map[string]string{"page": strconv.Itoa(page)})
}
