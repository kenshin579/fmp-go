# FMP Go SDK — Economics 그룹 (v0.15.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/economics-group`
- 토픽: FMP `economics` 카테고리 4 endpoint. 캠페인 13번째 그룹. 국채금리/경제지표/경제캘린더/리스크프리미엄.

## 결정 사항

- **범위**: economics 4 endpoint. 신규 `economics/` 패키지, internal/fetch.
- **4 구조체**(공유 없음): `TreasuryRate`(만기 컬럼), `EconomicIndicator`, `EconomicCalendarEvent`(estimate/unit nullable), `RiskPremium`.
- **path 주의**: 실제 path 는 `economic-indicators`/`economic-calendar`(단수 economic).
- **EconomicIndicators name 필수** 가드. MarketRiskPremium 무파라미터.
- **nullable**: EconomicCalendarEvent.Estimate `*float64`, Unit `*string`.
- **릴리스**: `v0.15.0`.

## 패키지 구조 + endpoint 매핑

| 메서드 | path | query | 반환 |
|---|---|---|---|
| `TreasuryRates(ctx, from, to)` | `/stable/treasury-rates` | from,to | `[]TreasuryRate` |
| `EconomicIndicators(ctx, name, from, to)` | `/stable/economic-indicators` | name(필수),from,to | `[]EconomicIndicator` |
| `EconomicCalendar(ctx, country, from, to)` | `/stable/economic-calendar` | country,from,to | `[]EconomicCalendarEvent` |
| `MarketRiskPremium(ctx)` | `/stable/market-risk-premium` | — | `[]RiskPremium` |

파일: `economics/client.go`(New + fromToParams helper), `economics/economics.go`(4 struct + 4 method).
- `fromToParams(from, to string) map[string]string` — from/to 비어있지 않으면 포함(빈 맵 시작).
- EconomicIndicators: name 빈값 가드 + params{name, from?, to?}. EconomicCalendar: params{country?, from?, to?}. MarketRiskPremium: nil.

## 루트 Client 와이어
```go
Directory *directory.Client
Economics *economics.Client // 경제(국채/지표/캘린더/리스크프리미엄)
```
`c.Economics = economics.New(hc)`. `TestNewClient_HasEconomics`.

## 응답 타입 (faithful, 필드 한국어 주석)
```go
// TreasuryRate — 미국 국채 수익률 곡선 (treasury-rates)
type TreasuryRate struct {
	Date    string  `json:"date"`    // 일자
	Month1  float64 `json:"month1"`  // 1개월
	Month2  float64 `json:"month2"`  // 2개월
	Month3  float64 `json:"month3"`  // 3개월
	Month6  float64 `json:"month6"`  // 6개월
	Year1   float64 `json:"year1"`   // 1년
	Year2   float64 `json:"year2"`   // 2년
	Year3   float64 `json:"year3"`   // 3년
	Year5   float64 `json:"year5"`   // 5년
	Year7   float64 `json:"year7"`   // 7년
	Year10  float64 `json:"year10"`  // 10년
	Year20  float64 `json:"year20"`  // 20년
	Year30  float64 `json:"year30"`  // 30년
}

// EconomicIndicator — 경제 지표 (economic-indicators). name 으로 지표 선택.
type EconomicIndicator struct {
	Name  string  `json:"name"`  // 지표명
	Date  string  `json:"date"`  // 일자
	Value float64 `json:"value"` // 값
}

// EconomicCalendarEvent — 경제 캘린더 이벤트 (economic-calendar)
type EconomicCalendarEvent struct {
	Date             string   `json:"date"`             // 일시(YYYY-MM-DD HH:MM:SS)
	Country          string   `json:"country"`          // 국가
	Event            string   `json:"event"`            // 이벤트명
	Currency         string   `json:"currency"`         // 통화
	Previous         float64  `json:"previous"`         // 이전값
	Estimate         *float64 `json:"estimate"`         // 예상치(null 가능)
	Actual           float64  `json:"actual"`           // 실제값
	Change           float64  `json:"change"`           // 변동
	Impact           string   `json:"impact"`           // 영향도(Low/Medium/High)
	ChangePercentage float64  `json:"changePercentage"` // 변동률
	Unit             *string  `json:"unit"`             // 단위(null 가능)
}

// RiskPremium — 국가별 리스크 프리미엄 (market-risk-premium)
type RiskPremium struct {
	Country                string  `json:"country"`                // 국가
	Continent              string  `json:"continent"`              // 대륙
	CountryRiskPremium     float64 `json:"countryRiskPremium"`     // 국가 리스크 프리미엄
	TotalEquityRiskPremium float64 `json:"totalEquityRiskPremium"` // 총 주식 리스크 프리미엄
}
```

## 시그니처 규칙
- TreasuryRates: `(ctx, from, to)` → `fetch.List[TreasuryRate](..., fromToParams(from, to))`.
- EconomicIndicators: `(ctx, name, from, to)` → name 빈값 가드 + params{name}+from/to.
- EconomicCalendar: `(ctx, country, from, to)` → params{country?}+from/to.
- MarketRiskPremium: `(ctx)` → `fetch.List[RiskPremium](..., nil)`.

## 테스트
- fixture 단위: TreasuryRate(Year10!=0), EconomicIndicator(Value!=0), EconomicCalendarEvent(Estimate null→nil + Unit null→nil + Previous 값), RiskPremium(CountryRiskPremium!=0).
- delegation: TreasuryRates(from,to) path+from/to / EconomicIndicators("CPI",...) path+name / EconomicCalendar("US",...) path+country / MarketRiskPremium path.
- 가드: EconomicIndicators 빈 name 1건.
- 통합: TreasuryRates("","") len>0 / EconomicIndicators("GDP","","") len>0 / MarketRiskPremium len>0.

## 문서 / 릴리스
- README Economics 행(4 endpoint).
- `examples/economics/main.go` — TreasuryRates + MarketRiskPremium.
- 릴리스 `v0.15.0`.

## 범위 밖 / 위험
- EconomicIndicators name 유효값 24종(GDP/CPI/...) — 검증 안 함(FMP 위임).
- continent 가 드물게 null 가능성 — 카탈로그 예시 비-null 이라 string 유지(문제 시 후속).
- 다음 그룹: marketHours 또는 crypto/forex.
