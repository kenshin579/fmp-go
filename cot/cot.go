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
