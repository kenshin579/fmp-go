package analyst

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// Grade — 개별 애널리스트 등급 변경 (grades)
type Grade struct {
	Symbol         string `json:"symbol"`         // 종목 심볼
	Date           string `json:"date"`           // 등급 변경일
	GradingCompany string `json:"gradingCompany"` // 평가 기관
	PreviousGrade  string `json:"previousGrade"`  // 이전 등급
	NewGrade       string `json:"newGrade"`       // 신규 등급
	Action         string `json:"action"`         // 조치 (maintain/upgrade/downgrade)
}

// GradesConsensus — 등급 컨센서스 집계 (grades-consensus)
type GradesConsensus struct {
	Symbol     string `json:"symbol"`     // 종목 심볼
	StrongBuy  int    `json:"strongBuy"`  // 적극 매수 수
	Buy        int    `json:"buy"`        // 매수
	Hold       int    `json:"hold"`       // 보유
	Sell       int    `json:"sell"`       // 매도
	StrongSell int    `json:"strongSell"` // 적극 매도
	Consensus  string `json:"consensus"`  // 종합 컨센서스
}

// HistoricalGrade — 일자별 등급 분포 (grades-historical). FMP 응답에 StrongBuy 없음.
type HistoricalGrade struct {
	Symbol                   string `json:"symbol"`                   // 종목 심볼
	Date                     string `json:"date"`                     // 기준일
	AnalystRatingsBuy        int    `json:"analystRatingsBuy"`        // 매수 의견 수
	AnalystRatingsHold       int    `json:"analystRatingsHold"`       // 보유
	AnalystRatingsSell       int    `json:"analystRatingsSell"`       // 매도
	AnalystRatingsStrongSell int    `json:"analystRatingsStrongSell"` // 적극 매도
}

// Grades 는 종목의 개별 애널리스트 등급 변경 이력을 조회한다.
func (c *Client) Grades(ctx context.Context, symbol string) ([]Grade, error) {
	return fetch.ListBySymbol[Grade](ctx, c.http, "/stable/grades", symbol)
}

// GradesConsensus 는 종목의 등급 컨센서스 집계를 조회한다. 결과 없으면 httpclient.ErrNotFound.
func (c *Client) GradesConsensus(ctx context.Context, symbol string) (*GradesConsensus, error) {
	return fetch.OneBySymbol[GradesConsensus](ctx, c.http, "/stable/grades-consensus", symbol)
}

// HistoricalGrades 는 종목의 일자별 등급 분포를 조회한다.
func (c *Client) HistoricalGrades(ctx context.Context, symbol string) ([]HistoricalGrade, error) {
	return fetch.ListBySymbol[HistoricalGrade](ctx, c.http, "/stable/grades-historical", symbol)
}
