package insidertrades

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// InsiderTransactionType — 거래 유형 코드 목록 (insider-trading-transaction-type)
type InsiderTransactionType struct {
	TransactionType string `json:"transactionType"` // 거래 유형 코드
}

// TradeStatistics — 종목 내부자 거래 통계 (insider-trading/statistics)
type TradeStatistics struct {
	Symbol                string  `json:"symbol"`                // 종목 심볼
	CIK                   string  `json:"cik"`                   // 회사 CIK
	Year                  int64   `json:"year"`                  // 연도
	Quarter               int64   `json:"quarter"`               // 분기
	AcquiredTransactions  int64   `json:"acquiredTransactions"`  // 취득 거래 수
	DisposedTransactions  int64   `json:"disposedTransactions"`  // 처분 거래 수
	AcquiredDisposedRatio float64 `json:"acquiredDisposedRatio"` // 취득/처분 비율
	TotalAcquired         int64   `json:"totalAcquired"`         // 총 취득
	TotalDisposed         int64   `json:"totalDisposed"`         // 총 처분
	AverageAcquired       float64 `json:"averageAcquired"`       // 평균 취득
	AverageDisposed       float64 `json:"averageDisposed"`       // 평균 처분
	TotalPurchases        int64   `json:"totalPurchases"`        // 총 매수
	TotalSales            int64   `json:"totalSales"`            // 총 매도
}

// AcquisitionOwnership — 수익적 소유 취득 (acquisition-of-beneficial-ownership).
// FMP 가 의결권/지분율 수치를 전부 문자열로 반환 → 전 필드 string.
type AcquisitionOwnership struct {
	CIK                              string `json:"cik"`                              // CIK
	Symbol                           string `json:"symbol"`                           // 종목 심볼
	FilingDate                       string `json:"filingDate"`                       // 공시일
	AcceptedDate                     string `json:"acceptedDate"`                     // 수리일
	CUSIP                            string `json:"cusip"`                            // CUSIP
	NameOfReportingPerson            string `json:"nameOfReportingPerson"`            // 보고자명
	CitizenshipOrPlaceOfOrganization string `json:"citizenshipOrPlaceOfOrganization"` // 시민권/설립지
	SoleVotingPower                  string `json:"soleVotingPower"`                  // 단독 의결권
	SharedVotingPower                string `json:"sharedVotingPower"`                // 공동 의결권
	SoleDispositivePower             string `json:"soleDispositivePower"`             // 단독 처분권
	SharedDispositivePower           string `json:"sharedDispositivePower"`           // 공동 처분권
	AmountBeneficiallyOwned          string `json:"amountBeneficiallyOwned"`          // 수익적 소유량
	PercentOfClass                   string `json:"percentOfClass"`                   // 클래스 비율(%)
	TypeOfReportingPerson            string `json:"typeOfReportingPerson"`            // 보고자 유형
	URL                              string `json:"url"`                              // 공시 URL
}

// ReportingName — 보고자명 검색 결과 (insider-trading/reporting-name)
type ReportingName struct {
	ReportingCik  string `json:"reportingCik"`  // 보고자 CIK
	ReportingName string `json:"reportingName"` // 보고자명
}

// TransactionTypes 는 내부자 거래 유형 코드 목록을 조회한다.
func (c *Client) TransactionTypes(ctx context.Context) ([]InsiderTransactionType, error) {
	return fetch.List[InsiderTransactionType](ctx, c.http, "/stable/insider-trading-transaction-type", nil)
}

// Statistics 는 종목의 내부자 거래 통계를 조회한다. symbol 필수.
func (c *Client) Statistics(ctx context.Context, symbol string) ([]TradeStatistics, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[TradeStatistics](ctx, c.http, "/stable/insider-trading/statistics", map[string]string{"symbol": symbol})
}

// AcquisitionOwnership 는 종목의 수익적 소유 취득 공시를 조회한다. symbol 필수.
func (c *Client) AcquisitionOwnership(ctx context.Context, symbol string, limit int) ([]AcquisitionOwnership, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	q := map[string]string{"symbol": symbol}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[AcquisitionOwnership](ctx, c.http, "/stable/acquisition-of-beneficial-ownership", q)
}

// SearchReportingName 은 보고자명을 검색한다. name 필수.
func (c *Client) SearchReportingName(ctx context.Context, name string) ([]ReportingName, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("fmp: name must not be empty")
	}
	return fetch.List[ReportingName](ctx, c.http, "/stable/insider-trading/reporting-name", map[string]string{"name": name})
}
