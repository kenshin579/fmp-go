# FMP Go SDK — Analyst 그룹 (v0.7.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/analyst-group`
- 토픽: FMP `analyst` 카테고리 8 endpoint 추가. 전체 API 커버리지 캠페인 5번째 그룹.

## 배경 / 목적

추천 순서의 다음 그룹. 애널리스트 등급/컨센서스/목표주가/평가점수/재무추정 — 종목 분석에 강력. 외부 개발자 가치 + moneyflow 종목 상세 보강.

## 결정 사항 (브레인스토밍)

- **범위**: analyst 8 endpoint 전부. 신규 `analyst/` 패키지, `internal/fetch` 사용.
- **단일 *T 4개**(consensus/snapshot/price-target × 2) + **list []T 3개**(grades/historical-grades/historical-ratings) + **financial-estimates**(params).
- **Rating 공유**: ratings-snapshot/ratings-historical → `Rating`(snapshot 은 Date "").
- **FinancialEstimate 합성**: analyst-estimates 카탈로그 예시 없음 → FMP 공개 shape 합성 + 통합 검증.
- **PriceTargetSummary.Publishers**: JSON 배열 문자열 그대로 string.
- **템플릿 계승**: 필드 한국어 주석, fixture + delegation 테스트.
- **릴리스**: `v0.7.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | helper | 반환 |
|---|---|---|---|---|
| `estimates.go` | `FinancialEstimates(ctx, symbol, period, page)` | `/stable/analyst-estimates` | List{symbol,period,page}+가드 | `[]FinancialEstimate` |
| `grades.go` | `Grades(ctx, symbol)` | `/stable/grades` | ListBySymbol | `[]Grade` |
| | `GradesConsensus(ctx, symbol)` | `/stable/grades-consensus` | OneBySymbol | `*GradesConsensus` |
| | `HistoricalGrades(ctx, symbol)` | `/stable/grades-historical` | ListBySymbol | `[]HistoricalGrade` |
| `ratings.go` | `RatingsSnapshot(ctx, symbol)` | `/stable/ratings-snapshot` | OneBySymbol | `*Rating` |
| | `HistoricalRatings(ctx, symbol)` | `/stable/ratings-historical` | ListBySymbol | `[]Rating` |
| `price_target.go` | `PriceTargetConsensus(ctx, symbol)` | `/stable/price-target-consensus` | OneBySymbol | `*PriceTargetConsensus` |
| | `PriceTargetSummary(ctx, symbol)` | `/stable/price-target-summary` | OneBySymbol | `*PriceTargetSummary` |
| `client.go` | `New(http)` | — | — | `*Client` |

## 응답 타입 (faithful, 필드 한국어 주석)

```go
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

// Rating — 종합 평가 점수 (ratings-snapshot / ratings-historical 공용). snapshot 은 Date "".
type Rating struct {
	Symbol                  string `json:"symbol"`                  // 종목 심볼
	Date                    string `json:"date"`                    // 기준일(snapshot 은 빈 문자열)
	Rating                  string `json:"rating"`                  // 등급 (예: A-)
	OverallScore            int    `json:"overallScore"`            // 종합 점수
	DiscountedCashFlowScore int    `json:"discountedCashFlowScore"` // DCF 점수
	ReturnOnEquityScore     int    `json:"returnOnEquityScore"`     // ROE 점수
	ReturnOnAssetsScore     int    `json:"returnOnAssetsScore"`     // ROA 점수
	DebtToEquityScore       int    `json:"debtToEquityScore"`       // 부채비율 점수
	PriceToEarningsScore    int    `json:"priceToEarningsScore"`    // PER 점수
	PriceToBookScore        int    `json:"priceToBookScore"`        // PBR 점수
}

// PriceTargetConsensus — 목표주가 컨센서스 (price-target-consensus)
type PriceTargetConsensus struct {
	Symbol          string  `json:"symbol"`          // 종목 심볼
	TargetHigh      float64 `json:"targetHigh"`      // 최고 목표가
	TargetLow       float64 `json:"targetLow"`       // 최저 목표가
	TargetConsensus float64 `json:"targetConsensus"` // 평균 목표가
	TargetMedian    float64 `json:"targetMedian"`    // 중앙값 목표가
}

// PriceTargetSummary — 목표주가 요약 (price-target-summary)
type PriceTargetSummary struct {
	Symbol                    string  `json:"symbol"`                    // 종목 심볼
	LastMonthCount            int     `json:"lastMonthCount"`            // 최근 1개월 리포트 수
	LastMonthAvgPriceTarget   float64 `json:"lastMonthAvgPriceTarget"`   // 최근 1개월 평균 목표가
	LastQuarterCount          int     `json:"lastQuarterCount"`          // 최근 분기 리포트 수
	LastQuarterAvgPriceTarget float64 `json:"lastQuarterAvgPriceTarget"` // 최근 분기 평균 목표가
	LastYearCount             int     `json:"lastYearCount"`             // 최근 1년 리포트 수
	LastYearAvgPriceTarget    float64 `json:"lastYearAvgPriceTarget"`    // 최근 1년 평균 목표가
	AllTimeCount              int     `json:"allTimeCount"`              // 전체 리포트 수
	AllTimeAvgPriceTarget     float64 `json:"allTimeAvgPriceTarget"`     // 전체 평균 목표가
	Publishers                string  `json:"publishers"`                // 발행처 목록(JSON 배열 문자열)
}

// FinancialEstimate — 애널리스트 재무 추정 (analyst-estimates). 카탈로그 예시 없음 → FMP 공개 shape 합성, 통합 확정.
type FinancialEstimate struct {
	Symbol             string  `json:"symbol"`             // 종목 심볼
	Date               string  `json:"date"`               // 추정 기준일
	RevenueLow         int64   `json:"revenueLow"`         // 매출 추정 최저
	RevenueHigh        int64   `json:"revenueHigh"`        // 매출 추정 최고
	RevenueAvg         int64   `json:"revenueAvg"`         // 매출 추정 평균
	EbitdaLow          int64   `json:"ebitdaLow"`          // EBITDA 최저
	EbitdaHigh         int64   `json:"ebitdaHigh"`         // EBITDA 최고
	EbitdaAvg          int64   `json:"ebitdaAvg"`          // EBITDA 평균
	EbitLow            int64   `json:"ebitLow"`            // EBIT 최저
	EbitHigh           int64   `json:"ebitHigh"`           // EBIT 최고
	EbitAvg            int64   `json:"ebitAvg"`            // EBIT 평균
	NetIncomeLow       int64   `json:"netIncomeLow"`       // 순이익 최저
	NetIncomeHigh      int64   `json:"netIncomeHigh"`      // 순이익 최고
	NetIncomeAvg       int64   `json:"netIncomeAvg"`       // 순이익 평균
	SgaExpenseLow      int64   `json:"sgaExpenseLow"`      // 판관비 최저
	SgaExpenseHigh     int64   `json:"sgaExpenseHigh"`     // 판관비 최고
	SgaExpenseAvg      int64   `json:"sgaExpenseAvg"`      // 판관비 평균
	EpsLow             float64 `json:"epsLow"`             // EPS 최저
	EpsHigh            float64 `json:"epsHigh"`            // EPS 최고
	EpsAvg             float64 `json:"epsAvg"`             // EPS 평균
	NumAnalystsRevenue int     `json:"numAnalystsRevenue"` // 매출 추정 애널리스트 수
	NumAnalystsEps     int     `json:"numAnalystsEps"`     // EPS 추정 애널리스트 수
}
```

## 시그니처 규칙
- 단일 *T(symbol): GradesConsensus/RatingsSnapshot/PriceTargetConsensus/PriceTargetSummary → `fetch.OneBySymbol`.
- list []T(symbol): Grades/HistoricalGrades/HistoricalRatings → `fetch.ListBySymbol`.
- financial-estimates: 빈 symbol/period 가드 + `fetch.List[FinancialEstimate](ctx, c.http, path, {"symbol","period","page"})`.

## 루트 Client 와이어
```go
type Client struct {
	...
	News    *news.Client
	Analyst *analyst.Client // 애널리스트 (신규)
}
```
`NewClient` 에 `c.Analyst = analyst.New(hc)`. `client_test.go` 에 `TestNewClient_HasAnalyst`.

## 테스트
- fixture 단위: 7 endpoint + Rating(snapshot date "" / historical date 값) 양쪽 + FinancialEstimate(합성).
- delegation: Grades(symbol) / RatingsSnapshot(symbol) / FinancialEstimates(symbol,period,page) path+param.
- 가드: 단일 빈 symbol(OneBySymbol 내장 검증 대표), financial-estimates 빈 symbol/period.
- 단일 *T 빈 배열 → ErrNotFound 1건.
- 통합(`//go:build integration`): GradesConsensus("AAPL") / PriceTargetConsensus("AAPL") / FinancialEstimates("AAPL","annual",0) — estimates 실 shape 로그.

## 문서 / 릴리스
- README 커버리지 표 Analyst 행(8 endpoint).
- `examples/analyst/main.go` — PriceTargetConsensus + GradesConsensus.
- 릴리스 `v0.7.0`.

## 범위 밖 / 위험
- 나머지 23 그룹 별도 PR(다음: calendar → statements 확장).
- **FinancialEstimate 합성 shape** — 통합테스트 확정, 다르면 조정(shares-float 선례).
- financial-estimates `period` 필수("annual"/"quarter") 빈 가드. limit FMP 기본값(page 만 노출, 후속 추가 가능).
- Rating 공유 — snapshot 빈 Date fixture 검증.
- HistoricalGrade 에 StrongBuy 필드 없음(FMP 응답 그대로).
