package form13f

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Filing 은 기관 투자자 SEC 최신 제출 파일 정보.
type Filing struct {
	Cik          string `json:"cik"`
	Name         string `json:"name"`
	Date         string `json:"date"`
	FilingDate   string `json:"filingDate"`
	AcceptedDate string `json:"acceptedDate"`
	FormType     string `json:"formType"`
	Link         string `json:"link"`
	FinalLink    string `json:"finalLink"`
}

// Holding 은 13F 파일링에서 추출한 개별 보유 종목 정보.
type Holding struct {
	Date          string `json:"date"`
	FilingDate    string `json:"filingDate"`
	AcceptedDate  string `json:"acceptedDate"`
	Cik           string `json:"cik"`
	SecurityCusip string `json:"securityCusip"`
	Symbol        string `json:"symbol"`
	NameOfIssuer  string `json:"nameOfIssuer"`
	Shares        int64  `json:"shares"`
	TitleOfClass  string `json:"titleOfClass"`
	SharesType    string `json:"sharesType"`
	PutCallShare  string `json:"putCallShare"`
	Value         int64  `json:"value"`
	Link          string `json:"link"`
	FinalLink     string `json:"finalLink"`
}

// FilingDate 는 특정 기관의 13F 파일링 날짜/연도/분기 정보.
type FilingDate struct {
	Date    string `json:"date"`
	Year    int64  `json:"year"`
	Quarter int64  `json:"quarter"`
}

// IndustrySummary 는 특정 기간의 산업별 보유 가치 요약.
type IndustrySummary struct {
	IndustryTitle string `json:"industryTitle"`
	IndustryValue int64  `json:"industryValue"`
	Date          string `json:"date"`
}

// LatestFilings 는 기관 투자자의 최신 SEC 제출 파일 목록을 반환한다.
func (c *Client) LatestFilings(ctx context.Context, page, limit int) ([]Filing, error) {
	return fetch.List[Filing](ctx, c.http, "/stable/institutional-ownership/latest", pageParams(page, limit))
}

// Extract 는 특정 기관(cik)의 연도/분기 13F 파일링에서 보유 종목 상세 정보를 반환한다.
func (c *Client) Extract(ctx context.Context, cik, year, quarter string) ([]Holding, error) {
	if strings.TrimSpace(cik) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: cik, year, quarter must not be empty")
	}
	return fetch.List[Holding](ctx, c.http, "/stable/institutional-ownership/extract", map[string]string{"cik": cik, "year": year, "quarter": quarter})
}

// FilingDates 는 특정 기관(cik)의 13F 파일링 날짜 목록을 반환한다.
func (c *Client) FilingDates(ctx context.Context, cik string) ([]FilingDate, error) {
	if strings.TrimSpace(cik) == "" {
		return nil, fmt.Errorf("fmp: cik must not be empty")
	}
	return fetch.List[FilingDate](ctx, c.http, "/stable/institutional-ownership/dates", map[string]string{"cik": cik})
}

// IndustrySummary 는 특정 연도/분기의 산업별 기관 보유 가치 요약을 반환한다.
func (c *Client) IndustrySummary(ctx context.Context, year, quarter string) ([]IndustrySummary, error) {
	if strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: year, quarter must not be empty")
	}
	return fetch.List[IndustrySummary](ctx, c.http, "/stable/institutional-ownership/industry-summary", map[string]string{"year": year, "quarter": quarter})
}
