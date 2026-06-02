# FMP Go SDK — Metrics 그룹 (v0.10.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/metrics-group`
- 토픽: FMP statements 카테고리의 파생 지표 8 endpoint. 전체 API 커버리지 캠페인 8번째 그룹. statements 분해의 2/3.

## 배경 / 목적

statements 27 endpoint 분해의 2번째 하위 그룹. 핵심 재무비율/지표/세그먼트 — moneyflow 종목 분석 화면에 직접 활용. statements-core(v0.9.0) 다음 단계.

## 결정 사항 (브레인스토밍)

- **범위**: 8 endpoint. 신규 `metrics/` 패키지 7개 + 기존 `ratios/` 패키지에 RatiosTTM 1개 추가.
- **metrics-ratios 제외**: `/stable/ratios` 는 이미 `ratios.Ratios` 로 구현됨(중복). 본 그룹에서 제외.
- **metrics-ratios-ttm 위치**: `/stable/ratios-ttm` 는 ratios 의 TTM 변형 → 신규 패키지 아닌 기존 `ratios/` 패키지에 `RatiosTTM` 으로 추가(응집).
- **internal/fetch 사용**: calendar/analyst 와 동일. list endpoint 는 빈 결과에 ErrNotFound 안 던짐(fetch.List 컨벤션), 단일(FinancialScores)만 OneBySymbol → ErrNotFound.
- **RevenueSegment 공유**: geographic/product segmentation 응답 shape 동일 → 단일 `RevenueSegment` 구조체 공유. `data map[string]int64`(세그먼트명 동적 키), `fiscalYear int`(숫자), `reportedCurrency *string`(null 가능).
- **fiscalYear 타입 불일치 주의**: key-metrics/owner-earnings 는 `string`("2024"), segmentation 은 `int`(2024). 공유 금지.
- **JSON 태그 오타 보존**: `researchAndDevelopementToRevenue[TTM]`(FMP 철자) 그대로.
- **필드 주석 정책**: 대량 비율 구조체(KeyMetrics/KeyMetricsTTM/RatioTTM)는 struct 단위 주석만(statements 관례). 작은/비자명 구조체(FinancialScores/OwnerEarning/EnterpriseValue/RevenueSegment)는 필드 주석 선택적.
- **릴리스**: `v0.10.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | helper | 반환 |
|---|---|---|---|---|
| `metrics/key_metrics.go` | `KeyMetrics(ctx, symbol, period, limit)` | `/stable/key-metrics` | List+guard | `[]KeyMetrics` |
| | `KeyMetricsTTM(ctx, symbol)` | `/stable/key-metrics-ttm` | ListBySymbol | `[]KeyMetricsTTM` |
| `metrics/scores.go` | `FinancialScores(ctx, symbol)` | `/stable/financial-scores` | OneBySymbol | `*FinancialScores` |
| `metrics/owner_earnings.go` | `OwnerEarnings(ctx, symbol, limit)` | `/stable/owner-earnings` | List+guard | `[]OwnerEarning` |
| `metrics/enterprise_values.go` | `EnterpriseValues(ctx, symbol, period, limit)` | `/stable/enterprise-values` | List+guard | `[]EnterpriseValue` |
| `metrics/segments.go` | `RevenueGeographicSegmentation(ctx, symbol, period)` | `/stable/revenue-geographic-segmentation` | List+guard | `[]RevenueSegment` |
| | `RevenueProductSegmentation(ctx, symbol, period)` | `/stable/revenue-product-segmentation` | List+guard | `[]RevenueSegment` |
| `metrics/client.go` | `New(http)` + `listParams` helper | — | — | `*Client` |
| `ratios/ratios_ttm.go` | `RatiosTTM(ctx, symbol)` | `/stable/ratios-ttm` | ListBySymbol | `[]RatioTTM` |

`listParams(symbol, period string, limit int) map[string]string` — symbol 항상, period 비어있지 않으면, limit>0 이면 포함.
List 계열 메서드는 호출 전 빈 symbol 가드(`strings.TrimSpace`).

## 루트 Client 와이어
```go
type Client struct {
	...
	Calendar *calendar.Client
	Metrics  *metrics.Client // 지표(key-metrics/scores/owner-earnings/EV/segments)
}
```
`NewClient` 에 `c.Metrics = metrics.New(hc)`. `client_test.go` 에 `TestNewClient_HasMetrics`. (ratios.RatiosTTM 은 기존 `c.Ratios` 로 노출.)

## 응답 타입 (faithful)

### KeyMetrics (key-metrics)
```go
// KeyMetrics — 핵심 지표 (key-metrics). 비율 다수, 일부 절대값 int64.
type KeyMetrics struct {
	Symbol                                 string  `json:"symbol"`
	Date                                   string  `json:"date"`
	FiscalYear                             string  `json:"fiscalYear"`
	Period                                 string  `json:"period"`
	ReportedCurrency                       string  `json:"reportedCurrency"`
	MarketCap                              int64   `json:"marketCap"`
	EnterpriseValue                        int64   `json:"enterpriseValue"`
	EvToSales                              float64 `json:"evToSales"`
	EvToOperatingCashFlow                  float64 `json:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64 `json:"evToFreeCashFlow"`
	EvToEBITDA                             float64 `json:"evToEBITDA"`
	NetDebtToEBITDA                        float64 `json:"netDebtToEBITDA"`
	CurrentRatio                           float64 `json:"currentRatio"`
	IncomeQuality                          float64 `json:"incomeQuality"`
	GrahamNumber                           float64 `json:"grahamNumber"`
	GrahamNetNet                           float64 `json:"grahamNetNet"`
	TaxBurden                              float64 `json:"taxBurden"`
	InterestBurden                         float64 `json:"interestBurden"`
	WorkingCapital                         int64   `json:"workingCapital"`
	InvestedCapital                        int64   `json:"investedCapital"`
	ReturnOnAssets                         float64 `json:"returnOnAssets"`
	OperatingReturnOnAssets                float64 `json:"operatingReturnOnAssets"`
	ReturnOnTangibleAssets                 float64 `json:"returnOnTangibleAssets"`
	ReturnOnEquity                         float64 `json:"returnOnEquity"`
	ReturnOnInvestedCapital                float64 `json:"returnOnInvestedCapital"`
	ReturnOnCapitalEmployed                float64 `json:"returnOnCapitalEmployed"`
	EarningsYield                          float64 `json:"earningsYield"`
	FreeCashFlowYield                      float64 `json:"freeCashFlowYield"`
	CapexToOperatingCashFlow               float64 `json:"capexToOperatingCashFlow"`
	CapexToDepreciation                    float64 `json:"capexToDepreciation"`
	CapexToRevenue                         float64 `json:"capexToRevenue"`
	SalesGeneralAndAdministrativeToRevenue float64 `json:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDevelopementToRevenue       float64 `json:"researchAndDevelopementToRevenue"`
	StockBasedCompensationToRevenue        float64 `json:"stockBasedCompensationToRevenue"`
	IntangiblesToTotalAssets               float64 `json:"intangiblesToTotalAssets"`
	AverageReceivables                     int64   `json:"averageReceivables"`
	AveragePayables                        int64   `json:"averagePayables"`
	AverageInventory                       int64   `json:"averageInventory"`
	DaysOfSalesOutstanding                 float64 `json:"daysOfSalesOutstanding"`
	DaysOfPayablesOutstanding              float64 `json:"daysOfPayablesOutstanding"`
	DaysOfInventoryOutstanding             float64 `json:"daysOfInventoryOutstanding"`
	OperatingCycle                         float64 `json:"operatingCycle"`
	CashConversionCycle                    float64 `json:"cashConversionCycle"`
	FreeCashFlowToEquity                   int64   `json:"freeCashFlowToEquity"`
	FreeCashFlowToFirm                     float64 `json:"freeCashFlowToFirm"`
	TangibleAssetValue                     int64   `json:"tangibleAssetValue"`
	NetCurrentAssetValue                   int64   `json:"netCurrentAssetValue"`
}
```

### KeyMetricsTTM (key-metrics-ttm) — date/fiscalYear/period/currency 없음
```go
// KeyMetricsTTM — TTM 핵심 지표 (key-metrics-ttm). symbol 외 메타 필드 없음.
type KeyMetricsTTM struct {
	Symbol                                    string  `json:"symbol"`
	MarketCap                                 int64   `json:"marketCap"`
	EnterpriseValueTTM                        int64   `json:"enterpriseValueTTM"`
	EvToSalesTTM                              float64 `json:"evToSalesTTM"`
	EvToOperatingCashFlowTTM                  float64 `json:"evToOperatingCashFlowTTM"`
	EvToFreeCashFlowTTM                       float64 `json:"evToFreeCashFlowTTM"`
	EvToEBITDATTM                             float64 `json:"evToEBITDATTM"`
	NetDebtToEBITDATTM                        float64 `json:"netDebtToEBITDATTM"`
	CurrentRatioTTM                           float64 `json:"currentRatioTTM"`
	IncomeQualityTTM                          float64 `json:"incomeQualityTTM"`
	GrahamNumberTTM                           float64 `json:"grahamNumberTTM"`
	GrahamNetNetTTM                           float64 `json:"grahamNetNetTTM"`
	TaxBurdenTTM                              float64 `json:"taxBurdenTTM"`
	InterestBurdenTTM                         float64 `json:"interestBurdenTTM"`
	WorkingCapitalTTM                         int64   `json:"workingCapitalTTM"`
	InvestedCapitalTTM                        int64   `json:"investedCapitalTTM"`
	ReturnOnAssetsTTM                         float64 `json:"returnOnAssetsTTM"`
	OperatingReturnOnAssetsTTM                float64 `json:"operatingReturnOnAssetsTTM"`
	ReturnOnTangibleAssetsTTM                 float64 `json:"returnOnTangibleAssetsTTM"`
	ReturnOnEquityTTM                         float64 `json:"returnOnEquityTTM"`
	ReturnOnInvestedCapitalTTM                float64 `json:"returnOnInvestedCapitalTTM"`
	ReturnOnCapitalEmployedTTM                float64 `json:"returnOnCapitalEmployedTTM"`
	EarningsYieldTTM                          float64 `json:"earningsYieldTTM"`
	FreeCashFlowYieldTTM                      float64 `json:"freeCashFlowYieldTTM"`
	CapexToOperatingCashFlowTTM               float64 `json:"capexToOperatingCashFlowTTM"`
	CapexToDepreciationTTM                    float64 `json:"capexToDepreciationTTM"`
	CapexToRevenueTTM                         float64 `json:"capexToRevenueTTM"`
	SalesGeneralAndAdministrativeToRevenueTTM float64 `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopementToRevenueTTM       float64 `json:"researchAndDevelopementToRevenueTTM"`
	StockBasedCompensationToRevenueTTM        float64 `json:"stockBasedCompensationToRevenueTTM"`
	IntangiblesToTotalAssetsTTM               float64 `json:"intangiblesToTotalAssetsTTM"`
	AverageReceivablesTTM                     int64   `json:"averageReceivablesTTM"`
	AveragePayablesTTM                        int64   `json:"averagePayablesTTM"`
	AverageInventoryTTM                       int64   `json:"averageInventoryTTM"`
	DaysOfSalesOutstandingTTM                 float64 `json:"daysOfSalesOutstandingTTM"`
	DaysOfPayablesOutstandingTTM              float64 `json:"daysOfPayablesOutstandingTTM"`
	DaysOfInventoryOutstandingTTM             float64 `json:"daysOfInventoryOutstandingTTM"`
	OperatingCycleTTM                         float64 `json:"operatingCycleTTM"`
	CashConversionCycleTTM                    float64 `json:"cashConversionCycleTTM"`
	FreeCashFlowToEquityTTM                   int64   `json:"freeCashFlowToEquityTTM"`
	FreeCashFlowToFirmTTM                     float64 `json:"freeCashFlowToFirmTTM"`
	TangibleAssetValueTTM                     int64   `json:"tangibleAssetValueTTM"`
	NetCurrentAssetValueTTM                   int64   `json:"netCurrentAssetValueTTM"`
}
```

### FinancialScores (financial-scores) — 단일 *T
```go
// FinancialScores — 재무 건전성 점수 (financial-scores)
type FinancialScores struct {
	Symbol           string  `json:"symbol"`           // 종목 심볼
	ReportedCurrency string  `json:"reportedCurrency"` // 보고 통화
	AltmanZScore     float64 `json:"altmanZScore"`     // Altman Z-Score
	PiotroskiScore   int     `json:"piotroskiScore"`   // Piotroski 점수(0~9)
	WorkingCapital   int64   `json:"workingCapital"`   // 운전자본
	TotalAssets      int64   `json:"totalAssets"`      // 총자산
	RetainedEarnings int64   `json:"retainedEarnings"` // 이익잉여금
	EBIT             int64   `json:"ebit"`             // EBIT
	MarketCap        int64   `json:"marketCap"`        // 시가총액
	TotalLiabilities int64   `json:"totalLiabilities"` // 총부채
	Revenue          int64   `json:"revenue"`          // 매출
}
```

### OwnerEarning (owner-earnings)
```go
// OwnerEarning — 오너 어닝 (owner-earnings). 버핏식 소유주 이익.
type OwnerEarning struct {
	Symbol                 string  `json:"symbol"`                 // 종목 심볼
	ReportedCurrency       string  `json:"reportedCurrency"`       // 보고 통화
	FiscalYear             string  `json:"fiscalYear"`             // 회계연도(문자열)
	Period                 string  `json:"period"`                 // 기간 (FY/Q1..)
	Date                   string  `json:"date"`                   // 기준일
	AveragePPE             float64 `json:"averagePPE"`             // 평균 유형자산 비율
	MaintenanceCapex       int64   `json:"maintenanceCapex"`       // 유지보수 capex
	OwnersEarnings         int64   `json:"ownersEarnings"`         // 소유주 이익
	GrowthCapex            int64   `json:"growthCapex"`            // 성장 capex
	OwnersEarningsPerShare float64 `json:"ownersEarningsPerShare"` // 주당 소유주 이익
}
```

### EnterpriseValue (enterprise-values)
```go
// EnterpriseValue — 기업가치 (enterprise-values)
type EnterpriseValue struct {
	Symbol                   string  `json:"symbol"`                   // 종목 심볼
	Date                     string  `json:"date"`                     // 기준일
	StockPrice               float64 `json:"stockPrice"`               // 주가
	NumberOfShares           int64   `json:"numberOfShares"`           // 발행주식수
	MarketCapitalization     int64   `json:"marketCapitalization"`     // 시가총액
	MinusCashAndCashEquivalents int64 `json:"minusCashAndCashEquivalents"` // (-)현금성자산
	AddTotalDebt             int64   `json:"addTotalDebt"`             // (+)총부채
	EnterpriseValue          int64   `json:"enterpriseValue"`          // 기업가치(EV)
}
```

### RevenueSegment (revenue-geographic-segmentation / revenue-product-segmentation 공용)
```go
// RevenueSegment — 매출 세그먼트 (지역/제품 공용). data 는 세그먼트명→매출 동적 맵.
type RevenueSegment struct {
	Symbol           string           `json:"symbol"`           // 종목 심볼
	FiscalYear       int              `json:"fiscalYear"`       // 회계연도(숫자)
	Period           string           `json:"period"`           // 기간 (FY/Q1..)
	ReportedCurrency *string          `json:"reportedCurrency"` // 보고 통화(null 가능)
	Date             string           `json:"date"`             // 기준일
	Data             map[string]int64 `json:"data"`             // 세그먼트명 → 매출액
}
```

### RatioTTM (ratios/ratios_ttm.go, ratios 패키지)
```go
// RatioTTM — TTM 재무비율 (ratios-ttm). symbol 외 메타 없음.
type RatioTTM struct {
	Symbol                                     string  `json:"symbol"`
	GrossProfitMarginTTM                       float64 `json:"grossProfitMarginTTM"`
	EbitMarginTTM                              float64 `json:"ebitMarginTTM"`
	EbitdaMarginTTM                            float64 `json:"ebitdaMarginTTM"`
	OperatingProfitMarginTTM                   float64 `json:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM                      float64 `json:"pretaxProfitMarginTTM"`
	ContinuousOperationsProfitMarginTTM        float64 `json:"continuousOperationsProfitMarginTTM"`
	NetProfitMarginTTM                         float64 `json:"netProfitMarginTTM"`
	BottomLineProfitMarginTTM                  float64 `json:"bottomLineProfitMarginTTM"`
	ReceivablesTurnoverTTM                     float64 `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                        float64 `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                       float64 `json:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM                      float64 `json:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM                           float64 `json:"assetTurnoverTTM"`
	CurrentRatioTTM                            float64 `json:"currentRatioTTM"`
	QuickRatioTTM                              float64 `json:"quickRatioTTM"`
	SolvencyRatioTTM                           float64 `json:"solvencyRatioTTM"`
	CashRatioTTM                               float64 `json:"cashRatioTTM"`
	PriceToEarningsRatioTTM                    float64 `json:"priceToEarningsRatioTTM"`
	PriceToEarningsGrowthRatioTTM              float64 `json:"priceToEarningsGrowthRatioTTM"`
	ForwardPriceToEarningsGrowthRatioTTM       float64 `json:"forwardPriceToEarningsGrowthRatioTTM"`
	PriceToBookRatioTTM                        float64 `json:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                       float64 `json:"priceToSalesRatioTTM"`
	PriceToFreeCashFlowRatioTTM                float64 `json:"priceToFreeCashFlowRatioTTM"`
	PriceToOperatingCashFlowRatioTTM           float64 `json:"priceToOperatingCashFlowRatioTTM"`
	DebtToAssetsRatioTTM                       float64 `json:"debtToAssetsRatioTTM"`
	DebtToEquityRatioTTM                       float64 `json:"debtToEquityRatioTTM"`
	DebtToCapitalRatioTTM                      float64 `json:"debtToCapitalRatioTTM"`
	LongTermDebtToCapitalRatioTTM              float64 `json:"longTermDebtToCapitalRatioTTM"`
	FinancialLeverageRatioTTM                  float64 `json:"financialLeverageRatioTTM"`
	WorkingCapitalTurnoverRatioTTM             float64 `json:"workingCapitalTurnoverRatioTTM"`
	OperatingCashFlowRatioTTM                  float64 `json:"operatingCashFlowRatioTTM"`
	OperatingCashFlowSalesRatioTTM             float64 `json:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM      float64 `json:"freeCashFlowOperatingCashFlowRatioTTM"`
	DebtServiceCoverageRatioTTM                float64 `json:"debtServiceCoverageRatioTTM"`
	InterestCoverageRatioTTM                   float64 `json:"interestCoverageRatioTTM"`
	ShortTermOperatingCashFlowCoverageRatioTTM float64 `json:"shortTermOperatingCashFlowCoverageRatioTTM"`
	OperatingCashFlowCoverageRatioTTM          float64 `json:"operatingCashFlowCoverageRatioTTM"`
	CapitalExpenditureCoverageRatioTTM         float64 `json:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTM       float64 `json:"dividendPaidAndCapexCoverageRatioTTM"`
	DividendPayoutRatioTTM                     float64 `json:"dividendPayoutRatioTTM"`
	DividendYieldTTM                           float64 `json:"dividendYieldTTM"`
	EnterpriseValueTTM                         int64   `json:"enterpriseValueTTM"`
	RevenuePerShareTTM                         float64 `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM                       float64 `json:"netIncomePerShareTTM"`
	InterestDebtPerShareTTM                    float64 `json:"interestDebtPerShareTTM"`
	CashPerShareTTM                            float64 `json:"cashPerShareTTM"`
	BookValuePerShareTTM                       float64 `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM               float64 `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM              float64 `json:"shareholdersEquityPerShareTTM"`
	OperatingCashFlowPerShareTTM               float64 `json:"operatingCashFlowPerShareTTM"`
	CapexPerShareTTM                           float64 `json:"capexPerShareTTM"`
	FreeCashFlowPerShareTTM                    float64 `json:"freeCashFlowPerShareTTM"`
	NetIncomePerEBTTTM                         float64 `json:"netIncomePerEBTTTM"`
	EbtPerEbitTTM                              float64 `json:"ebtPerEbitTTM"`
	PriceToFairValueTTM                        float64 `json:"priceToFairValueTTM"`
	DebtToMarketCapTTM                         float64 `json:"debtToMarketCapTTM"`
	EffectiveTaxRateTTM                        float64 `json:"effectiveTaxRateTTM"`
	EnterpriseValueMultipleTTM                 float64 `json:"enterpriseValueMultipleTTM"`
}
```

## 시그니처 규칙
- KeyMetrics/EnterpriseValues: `(ctx, symbol, period string, limit int)` → 빈 symbol 가드 + `fetch.List[T](ctx, c.http, path, listParams(symbol, period, limit))`.
- OwnerEarnings: `(ctx, symbol string, limit int)` → guard + List+listParams(symbol,"",limit).
- RevenueGeographic/ProductSegmentation: `(ctx, symbol, period string)` → guard + List+listParams(symbol,period,0).
- KeyMetricsTTM / ratios.RatiosTTM: `(ctx, symbol string)` → `fetch.ListBySymbol`(가드 내장).
- FinancialScores: `(ctx, symbol string)` → `fetch.OneBySymbol` → `*FinancialScores`.

## 테스트
- fixture 단위: 8 endpoint. KeyMetrics(주요 float/int 필드), KeyMetricsTTM, RatioTTM, FinancialScores(piotroskiScore int / altmanZScore float), OwnerEarning, EnterpriseValue, RevenueSegment(data 맵 키/값 + reportedCurrency null→nil + fiscalYear int).
- delegation: KeyMetrics(symbol,period,limit) path+쿼리 / RevenueGeographic path / FinancialScores path+symbol.
- 가드: List 계열 빈 symbol 대표 1건. FinancialScores 빈 배열 → ErrNotFound 1건.
- 통합(`//go:build integration`): KeyMetrics(AAPL,annual,2) / FinancialScores(AAPL) piotroskiScore 범위 / RevenueProductSegmentation(AAPL,annual) data 비어있지 않음 / ratios.RatiosTTM(AAPL).

## 문서 / 릴리스
- README 커버리지 표: Metrics 행 신규(7 endpoint) + Ratios 행에 RatiosTTM 추가(2 endpoint).
- `examples/metrics/main.go` — KeyMetrics + FinancialScores.
- 릴리스 `v0.10.0`.

## 범위 밖 / 위험
- as-reported/reports 그룹(v0.11.0) 별도 PR.
- segmentation `structure` 쿼리(flat 등) 미노출 — 기본값 사용(후속 가능).
- fiscalYear 타입 불일치(string vs int) — endpoint 별 정확히 구분.
- RevenueSegment.Data 값이 드물게 소수일 가능성 — 카탈로그 정수 기준 int64, 통합테스트로 확인.
- key-metrics 절대값(freeCashFlowToFirm 등)이 소수로 와 일부 int64 가정 깨질 수 있음 — freeCashFlowToFirm 은 float64 로 이미 처리. 통합테스트 t.Logf 로 재확인.
```
