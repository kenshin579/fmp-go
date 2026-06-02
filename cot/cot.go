package cot

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// COTAnalysis — COT 분석/요약 (commitment-of-traders-analysis).
type COTAnalysis struct {
	Symbol                       string  `json:"symbol"`                       // 심볼
	Date                         string  `json:"date"`                         // 일자
	Name                         string  `json:"name"`                         // 상품명
	Sector                       string  `json:"sector"`                       // 섹터
	Exchange                     string  `json:"exchange"`                     // 거래소
	CurrentLongMarketSituation   float64 `json:"currentLongMarketSituation"`   // 현재 롱 비중
	CurrentShortMarketSituation  float64 `json:"currentShortMarketSituation"`  // 현재 숏 비중
	MarketSituation              string  `json:"marketSituation"`              // 현재 시장 상황
	PreviousLongMarketSituation  float64 `json:"previousLongMarketSituation"`  // 이전 롱 비중
	PreviousShortMarketSituation float64 `json:"previousShortMarketSituation"` // 이전 숏 비중
	PreviousMarketSituation      string  `json:"previousMarketSituation"`      // 이전 시장 상황
	NetPostion                   int64   `json:"netPostion"`                   // 순포지션(FMP 오타 netPostion)
	PreviousNetPosition          int64   `json:"previousNetPosition"`          // 이전 순포지션
	ChangeInNetPosition          float64 `json:"changeInNetPosition"`          // 순포지션 변화
	MarketSentiment              string  `json:"marketSentiment"`              // 시장 심리
	ReversalTrend                bool    `json:"reversalTrend"`                // 반전 추세 여부
}

// COTList — COT 보고 대상 상품 목록 (commitment-of-traders-list).
type COTList struct {
	Symbol string `json:"symbol"` // 심볼
	Name   string `json:"name"`   // 상품명
}

// Analysis 는 COT 분석/요약을 조회한다. symbol/from/to 는 선택 필터.
func (c *Client) Analysis(ctx context.Context, symbol, from, to string) ([]COTAnalysis, error) {
	return fetch.List[COTAnalysis](ctx, c.http, "/stable/commitment-of-traders-analysis", filterParams(symbol, from, to))
}

// List 는 COT 보고 대상 상품 목록을 조회한다.
func (c *Client) List(ctx context.Context) ([]COTList, error) {
	return fetch.List[COTList](ctx, c.http, "/stable/commitment-of-traders-list", nil)
}

// COTReport — COT 전체 포지션 리포트 (commitment-of-traders-report). FMP 키 그대로(오타/접미사 보존).
type COTReport struct {
	// 식별 필드
	Symbol                 string `json:"symbol"`
	Date                   string `json:"date"`
	Name                   string `json:"name"`
	Sector                 string `json:"sector"`
	MarketAndExchangeNames string `json:"marketAndExchangeNames"`
	CftcContractMarketCode string `json:"cftcContractMarketCode"`
	CftcMarketCode         string `json:"cftcMarketCode"`
	CftcRegionCode         string `json:"cftcRegionCode"`
	CftcCommodityCode      string `json:"cftcCommodityCode"`
	ContractUnits          string `json:"contractUnits"`

	// All — 오픈 인터레스트 (포지션 수)
	OpenInterestAll           int64 `json:"openInterestAll"`
	NoncommPositionsLongAll   int64 `json:"noncommPositionsLongAll"`
	NoncommPositionsShortAll  int64 `json:"noncommPositionsShortAll"`
	NoncommPositionsSpreadAll int64 `json:"noncommPositionsSpreadAll"`
	CommPositionsLongAll      int64 `json:"commPositionsLongAll"`
	CommPositionsShortAll     int64 `json:"commPositionsShortAll"`
	TotReptPositionsLongAll   int64 `json:"totReptPositionsLongAll"`
	TotReptPositionsShortAll  int64 `json:"totReptPositionsShortAll"`
	NonreptPositionsLongAll   int64 `json:"nonreptPositionsLongAll"`
	NonreptPositionsShortAll  int64 `json:"nonreptPositionsShortAll"`

	// Old — 기존(Old) 오픈 인터레스트
	OpenInterestOld           int64 `json:"openInterestOld"`
	NoncommPositionsLongOld   int64 `json:"noncommPositionsLongOld"`
	NoncommPositionsShortOld  int64 `json:"noncommPositionsShortOld"`
	NoncommPositionsSpreadOld int64 `json:"noncommPositionsSpreadOld"`
	CommPositionsLongOld      int64 `json:"commPositionsLongOld"`
	CommPositionsShortOld     int64 `json:"commPositionsShortOld"`
	TotReptPositionsLongOld   int64 `json:"totReptPositionsLongOld"`
	TotReptPositionsShortOld  int64 `json:"totReptPositionsShortOld"`
	NonreptPositionsLongOld   int64 `json:"nonreptPositionsLongOld"`
	NonreptPositionsShortOld  int64 `json:"nonreptPositionsShortOld"`

	// Other — 기타(Other) 오픈 인터레스트
	OpenInterestOther           int64 `json:"openInterestOther"`
	NoncommPositionsLongOther   int64 `json:"noncommPositionsLongOther"`
	NoncommPositionsShortOther  int64 `json:"noncommPositionsShortOther"`
	NoncommPositionsSpreadOther int64 `json:"noncommPositionsSpreadOther"`
	CommPositionsLongOther      int64 `json:"commPositionsLongOther"`
	CommPositionsShortOther     int64 `json:"commPositionsShortOther"`
	TotReptPositionsLongOther   int64 `json:"totReptPositionsLongOther"`
	TotReptPositionsShortOther  int64 `json:"totReptPositionsShortOther"`
	NonreptPositionsLongOther   int64 `json:"nonreptPositionsLongOther"`
	NonreptPositionsShortOther  int64 `json:"nonreptPositionsShortOther"`

	// Change — 포지션 변화량 (FMP 오타 "Spead" 보존)
	ChangeInOpenInterestAll int64 `json:"changeInOpenInterestAll"`
	ChangeInNoncommLongAll  int64 `json:"changeInNoncommLongAll"`
	ChangeInNoncommShortAll int64 `json:"changeInNoncommShortAll"`
	ChangeInNoncommSpeadAll int64 `json:"changeInNoncommSpeadAll"`
	ChangeInCommLongAll     int64 `json:"changeInCommLongAll"`
	ChangeInCommShortAll    int64 `json:"changeInCommShortAll"`
	ChangeInTotReptLongAll  int64 `json:"changeInTotReptLongAll"`
	ChangeInTotReptShortAll int64 `json:"changeInTotReptShortAll"`
	ChangeInNonreptLongAll  int64 `json:"changeInNonreptLongAll"`
	ChangeInNonreptShortAll int64 `json:"changeInNonreptShortAll"`

	// PctOf — All 오픈 인터레스트 비중
	PctOfOpenInterestAll    float64 `json:"pctOfOpenInterestAll"`
	PctOfOiNoncommLongAll   float64 `json:"pctOfOiNoncommLongAll"`
	PctOfOiNoncommShortAll  float64 `json:"pctOfOiNoncommShortAll"`
	PctOfOiNoncommSpreadAll float64 `json:"pctOfOiNoncommSpreadAll"`
	PctOfOiCommLongAll      float64 `json:"pctOfOiCommLongAll"`
	PctOfOiCommShortAll     float64 `json:"pctOfOiCommShortAll"`
	PctOfOiTotReptLongAll   float64 `json:"pctOfOiTotReptLongAll"`
	PctOfOiTotReptShortAll  float64 `json:"pctOfOiTotReptShortAll"`
	PctOfOiNonreptLongAll   float64 `json:"pctOfOiNonreptLongAll"`
	PctOfOiNonreptShortAll  float64 `json:"pctOfOiNonreptShortAll"`

	// PctOf — Old (접미사 Ol) 오픈 인터레스트 비중
	PctOfOpenInterestOl    float64 `json:"pctOfOpenInterestOl"`
	PctOfOiNoncommLongOl   float64 `json:"pctOfOiNoncommLongOl"`
	PctOfOiNoncommShortOl  float64 `json:"pctOfOiNoncommShortOl"`
	PctOfOiNoncommSpreadOl float64 `json:"pctOfOiNoncommSpreadOl"`
	PctOfOiCommLongOl      float64 `json:"pctOfOiCommLongOl"`
	PctOfOiCommShortOl     float64 `json:"pctOfOiCommShortOl"`
	PctOfOiTotReptLongOl   float64 `json:"pctOfOiTotReptLongOl"`
	PctOfOiTotReptShortOl  float64 `json:"pctOfOiTotReptShortOl"`
	PctOfOiNonreptLongOl   float64 `json:"pctOfOiNonreptLongOl"`
	PctOfOiNonreptShortOl  float64 `json:"pctOfOiNonreptShortOl"`

	// PctOf — Other 오픈 인터레스트 비중
	PctOfOpenInterestOther    float64 `json:"pctOfOpenInterestOther"`
	PctOfOiNoncommLongOther   float64 `json:"pctOfOiNoncommLongOther"`
	PctOfOiNoncommShortOther  float64 `json:"pctOfOiNoncommShortOther"`
	PctOfOiNoncommSpreadOther float64 `json:"pctOfOiNoncommSpreadOther"`
	PctOfOiCommLongOther      float64 `json:"pctOfOiCommLongOther"`
	PctOfOiCommShortOther     float64 `json:"pctOfOiCommShortOther"`
	PctOfOiTotReptLongOther   float64 `json:"pctOfOiTotReptLongOther"`
	PctOfOiTotReptShortOther  float64 `json:"pctOfOiTotReptShortOther"`
	PctOfOiNonreptLongOther   float64 `json:"pctOfOiNonreptLongOther"`
	PctOfOiNonreptShortOther  float64 `json:"pctOfOiNonreptShortOther"`

	// Traders — All 트레이더 수
	TradersTotAll           int64 `json:"tradersTotAll"`
	TradersNoncommLongAll   int64 `json:"tradersNoncommLongAll"`
	TradersNoncommShortAll  int64 `json:"tradersNoncommShortAll"`
	TradersNoncommSpreadAll int64 `json:"tradersNoncommSpreadAll"`
	TradersCommLongAll      int64 `json:"tradersCommLongAll"`
	TradersCommShortAll     int64 `json:"tradersCommShortAll"`
	TradersTotReptLongAll   int64 `json:"tradersTotReptLongAll"`
	TradersTotReptShortAll  int64 `json:"tradersTotReptShortAll"`

	// Traders — Old (접미사 Ol, FMP 오타 "Spead" 보존)
	TradersTotOl          int64 `json:"tradersTotOl"`
	TradersNoncommLongOl  int64 `json:"tradersNoncommLongOl"`
	TradersNoncommShortOl int64 `json:"tradersNoncommShortOl"`
	TradersNoncommSpeadOl int64 `json:"tradersNoncommSpeadOl"`
	TradersCommLongOl     int64 `json:"tradersCommLongOl"`
	TradersCommShortOl    int64 `json:"tradersCommShortOl"`
	TradersTotReptLongOl  int64 `json:"tradersTotReptLongOl"`
	TradersTotReptShortOl int64 `json:"tradersTotReptShortOl"`

	// Traders — Other 트레이더 수
	TradersTotOther           int64 `json:"tradersTotOther"`
	TradersNoncommLongOther   int64 `json:"tradersNoncommLongOther"`
	TradersNoncommShortOther  int64 `json:"tradersNoncommShortOther"`
	TradersNoncommSpreadOther int64 `json:"tradersNoncommSpreadOther"`
	TradersCommLongOther      int64 `json:"tradersCommLongOther"`
	TradersCommShortOther     int64 `json:"tradersCommShortOther"`
	TradersTotReptLongOther   int64 `json:"tradersTotReptLongOther"`
	TradersTotReptShortOther  int64 `json:"tradersTotReptShortOther"`

	// Conc — All 집중도
	ConcGrossLe4TdrLongAll  float64 `json:"concGrossLe4TdrLongAll"`
	ConcGrossLe4TdrShortAll float64 `json:"concGrossLe4TdrShortAll"`
	ConcGrossLe8TdrLongAll  float64 `json:"concGrossLe8TdrLongAll"`
	ConcGrossLe8TdrShortAll float64 `json:"concGrossLe8TdrShortAll"`
	ConcNetLe4TdrLongAll    float64 `json:"concNetLe4TdrLongAll"`
	ConcNetLe4TdrShortAll   float64 `json:"concNetLe4TdrShortAll"`
	ConcNetLe8TdrLongAll    float64 `json:"concNetLe8TdrLongAll"`
	ConcNetLe8TdrShortAll   float64 `json:"concNetLe8TdrShortAll"`

	// Conc — Old (접미사 Ol) 집중도
	ConcGrossLe4TdrLongOl  float64 `json:"concGrossLe4TdrLongOl"`
	ConcGrossLe4TdrShortOl float64 `json:"concGrossLe4TdrShortOl"`
	ConcGrossLe8TdrLongOl  float64 `json:"concGrossLe8TdrLongOl"`
	ConcGrossLe8TdrShortOl float64 `json:"concGrossLe8TdrShortOl"`
	ConcNetLe4TdrLongOl    float64 `json:"concNetLe4TdrLongOl"`
	ConcNetLe4TdrShortOl   float64 `json:"concNetLe4TdrShortOl"`
	ConcNetLe8TdrLongOl    float64 `json:"concNetLe8TdrLongOl"`
	ConcNetLe8TdrShortOl   float64 `json:"concNetLe8TdrShortOl"`

	// Conc — Other 집중도
	ConcGrossLe4TdrLongOther  float64 `json:"concGrossLe4TdrLongOther"`
	ConcGrossLe4TdrShortOther float64 `json:"concGrossLe4TdrShortOther"`
	ConcGrossLe8TdrLongOther  float64 `json:"concGrossLe8TdrLongOther"`
	ConcGrossLe8TdrShortOther float64 `json:"concGrossLe8TdrShortOther"`
	ConcNetLe4TdrLongOther    float64 `json:"concNetLe4TdrLongOther"`
	ConcNetLe4TdrShortOther   float64 `json:"concNetLe4TdrShortOther"`
	ConcNetLe8TdrLongOther    float64 `json:"concNetLe8TdrLongOther"`
	ConcNetLe8TdrShortOther   float64 `json:"concNetLe8TdrShortOther"`
}

// Report 는 COT 전체 포지션 리포트를 조회한다. symbol/from/to 는 선택 필터.
func (c *Client) Report(ctx context.Context, symbol, from, to string) ([]COTReport, error) {
	return fetch.List[COTReport](ctx, c.http, "/stable/commitment-of-traders-report", filterParams(symbol, from, to))
}
