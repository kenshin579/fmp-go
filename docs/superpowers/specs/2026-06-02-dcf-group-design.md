# FMP Go SDK — Discounted Cash Flow 그룹 (v0.19.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/dcf-group`
- 토픽: FMP `discountedCashFlow` 카테고리 4 endpoint. 캠페인 17번째 그룹. DCF 밸류에이션.

## 결정 사항
- 신규 `dcf/` 패키지, internal/fetch. 3 구조체.
- `DCFValue`(discounted-cash-flow + levered-discounted-cash-flow 공유, 4필드), `CustomDCFAdvanced`(상세 다년 투영), `CustomDCFLevered`(레버드 변형, 필드셋 다름 — 별도).
- 두 custom 은 동일 쿼리(symbol + 18 override) 이나 응답 shape 다름.
- custom 18 override 는 `*float64`(미설정 시 쿼리 제외) → `CustomDCFParams` 구조체.
- **JSON 키 주의**: 단순 DCF 의 `"Stock Price"`(공백+대문자), custom 의 `costofDebt`(소문자 of) vs `costOfEquity`(대문자 Of).
- 카탈로그 파일명-경로 불일치: dcf-advanced.md→`/stable/discounted-cash-flow`(단순), custom-dcf-advanced.md→`/stable/custom-discounted-cash-flow`(상세).
- 릴리스 `v0.19.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | 반환 |
|---|---|---|
| `DiscountedCashFlow(ctx, symbol)` | `/stable/discounted-cash-flow` | `[]DCFValue` |
| `LeveredDiscountedCashFlow(ctx, symbol)` | `/stable/levered-discounted-cash-flow` | `[]DCFValue` |
| `CustomDiscountedCashFlow(ctx, p CustomDCFParams)` | `/stable/custom-discounted-cash-flow` | `[]CustomDCFAdvanced` |
| `CustomLeveredDiscountedCashFlow(ctx, p CustomDCFParams)` | `/stable/custom-levered-discounted-cash-flow` | `[]CustomDCFLevered` |

파일: `dcf/client.go`(New), `dcf/dcf.go`(DCFValue + 2 단순 method), `dcf/custom.go`(CustomDCFParams + 2 custom struct + 2 method).
- 단순 2개: symbol 빈값 가드 + `fetch.List[DCFValue](..., {"symbol": symbol})`.
- custom 2개: p.Symbol 빈값 가드 + `p.queryParams()`.

## 루트 Client 와이어
```go
TechnicalIndicators *technicals.Client
DCF *dcf.Client // DCF 밸류에이션
```
`c.DCF = dcf.New(hc)`. `TestNewClient_HasDCF`.

## 응답 타입 (faithful)
```go
// DCFValue — 단순 DCF 내재가치 (discounted-cash-flow / levered-discounted-cash-flow 공유)
type DCFValue struct {
	Symbol     string  `json:"symbol"`      // 종목 심볼
	Date       string  `json:"date"`        // 기준일
	DCF        float64 `json:"dcf"`         // 주당 내재가치
	StockPrice float64 `json:"Stock Price"` // 현재가(FMP 키 "Stock Price")
}

// CustomDCFParams — custom DCF 입력. Symbol 필수, 나머지 override 는 미설정 시 제외.
type CustomDCFParams struct {
	Symbol                                     string
	RevenueGrowthPct                           *float64
	EbitdaPct                                  *float64
	DepreciationAndAmortizationPct             *float64
	CashAndShortTermInvestmentsPct             *float64
	ReceivablesPct                             *float64
	InventoriesPct                             *float64
	PayablePct                                 *float64
	EbitPct                                    *float64
	CapitalExpenditurePct                      *float64
	OperatingCashFlowPct                       *float64
	SellingGeneralAndAdministrativeExpensesPct *float64
	TaxRate                                    *float64
	LongTermGrowthRate                         *float64
	CostOfDebt                                 *float64
	CostOfEquity                               *float64
	MarketRiskPremium                          *float64
	Beta                                       *float64
	RiskFreeRate                               *float64
}

// CustomDCFAdvanced — 상세 다년 DCF 투영 (custom-discounted-cash-flow)
type CustomDCFAdvanced struct {
	Year                         string  `json:"year"`
	Symbol                       string  `json:"symbol"`
	Revenue                      int64   `json:"revenue"`
	RevenuePercentage            float64 `json:"revenuePercentage"`
	EBITDA                       int64   `json:"ebitda"`
	EBITDAPercentage             float64 `json:"ebitdaPercentage"`
	EBIT                         int64   `json:"ebit"`
	EBITPercentage               float64 `json:"ebitPercentage"`
	Depreciation                 int64   `json:"depreciation"`
	DepreciationPercentage       float64 `json:"depreciationPercentage"`
	TotalCash                    int64   `json:"totalCash"`
	TotalCashPercentage          float64 `json:"totalCashPercentage"`
	Receivables                  int64   `json:"receivables"`
	ReceivablesPercentage        float64 `json:"receivablesPercentage"`
	Inventories                  int64   `json:"inventories"`
	InventoriesPercentage        float64 `json:"inventoriesPercentage"`
	Payable                      int64   `json:"payable"`
	PayablePercentage            float64 `json:"payablePercentage"`
	CapitalExpenditure           int64   `json:"capitalExpenditure"`
	CapitalExpenditurePercentage float64 `json:"capitalExpenditurePercentage"`
	Price                        float64 `json:"price"`
	Beta                         float64 `json:"beta"`
	DilutedSharesOutstanding     int64   `json:"dilutedSharesOutstanding"`
	CostOfDebt                   float64 `json:"costofDebt"` // FMP 키 소문자 of
	TaxRate                      float64 `json:"taxRate"`
	AfterTaxCostOfDebt           float64 `json:"afterTaxCostOfDebt"`
	RiskFreeRate                 float64 `json:"riskFreeRate"`
	MarketRiskPremium            float64 `json:"marketRiskPremium"`
	CostOfEquity                 float64 `json:"costOfEquity"`
	TotalDebt                    int64   `json:"totalDebt"`
	TotalEquity                  int64   `json:"totalEquity"`
	TotalCapital                 int64   `json:"totalCapital"`
	DebtWeighting                float64 `json:"debtWeighting"`
	EquityWeighting              float64 `json:"equityWeighting"`
	WACC                         float64 `json:"wacc"`
	TaxRateCash                  int64   `json:"taxRateCash"`
	EBIAT                        int64   `json:"ebiat"`
	UFCF                         int64   `json:"ufcf"`
	SumPvUfcf                    int64   `json:"sumPvUfcf"`
	LongTermGrowthRate           float64 `json:"longTermGrowthRate"`
	TerminalValue                int64   `json:"terminalValue"`
	PresentTerminalValue         int64   `json:"presentTerminalValue"`
	EnterpriseValue              int64   `json:"enterpriseValue"`
	NetDebt                      int64   `json:"netDebt"`
	EquityValue                  int64   `json:"equityValue"`
	EquityValuePerShare          float64 `json:"equityValuePerShare"`
	FreeCashFlowT1               int64   `json:"freeCashFlowT1"`
}

// CustomDCFLevered — 레버드 다년 DCF 투영 (custom-levered-discounted-cash-flow).
// advanced 와 필드셋 다름(operating cash flow / levered FCF 중심).
type CustomDCFLevered struct {
	Year                         string  `json:"year"`
	Symbol                       string  `json:"symbol"`
	Revenue                      int64   `json:"revenue"`
	RevenuePercentage            float64 `json:"revenuePercentage"`
	CapitalExpenditure           int64   `json:"capitalExpenditure"`
	CapitalExpenditurePercentage float64 `json:"capitalExpenditurePercentage"`
	Price                        float64 `json:"price"`
	Beta                         float64 `json:"beta"`
	DilutedSharesOutstanding     int64   `json:"dilutedSharesOutstanding"`
	CostOfDebt                   float64 `json:"costofDebt"`
	TaxRate                      float64 `json:"taxRate"`
	AfterTaxCostOfDebt           float64 `json:"afterTaxCostOfDebt"`
	RiskFreeRate                 float64 `json:"riskFreeRate"`
	MarketRiskPremium            float64 `json:"marketRiskPremium"`
	CostOfEquity                 float64 `json:"costOfEquity"`
	TotalDebt                    int64   `json:"totalDebt"`
	TotalEquity                  int64   `json:"totalEquity"`
	TotalCapital                 int64   `json:"totalCapital"`
	DebtWeighting                float64 `json:"debtWeighting"`
	EquityWeighting              float64 `json:"equityWeighting"`
	WACC                         float64 `json:"wacc"`
	OperatingCashFlow            int64   `json:"operatingCashFlow"`
	PvLfcf                       int64   `json:"pvLfcf"`
	SumPvLfcf                    int64   `json:"sumPvLfcf"`
	FreeCashFlow                 int64   `json:"freeCashFlow"`
	OperatingCashFlowPercentage  float64 `json:"operatingCashFlowPercentage"`
	LongTermGrowthRate           float64 `json:"longTermGrowthRate"`
	TerminalValue                int64   `json:"terminalValue"`
	PresentTerminalValue         int64   `json:"presentTerminalValue"`
	EnterpriseValue              int64   `json:"enterpriseValue"`
	NetDebt                      int64   `json:"netDebt"`
	EquityValue                  int64   `json:"equityValue"`
	EquityValuePerShare          float64 `json:"equityValuePerShare"`
	FreeCashFlowT1               int64   `json:"freeCashFlowT1"`
}
```

## CustomDCFParams.queryParams
symbol 항상 포함. 각 `*float64` 필드가 non-nil 이면 `strconv.FormatFloat(*v, 'f', -1, 64)` 로 해당 쿼리 키 추가. 쿼리 키는 JSON override 명과 동일(revenueGrowthPct, ebitdaPct, ..., riskFreeRate).

## 시그니처 규칙
- DiscountedCashFlow/LeveredDiscountedCashFlow: `(ctx, symbol)` → symbol 가드 + `fetch.List[DCFValue](..., {"symbol": symbol})`.
- CustomDiscountedCashFlow/CustomLeveredDiscountedCashFlow: `(ctx, p CustomDCFParams)` → p.Symbol 가드 + `fetch.List[T](..., p.queryParams())`.

## 테스트
- fixture 단위: DCFValue(StockPrice 키 "Stock Price" 매핑, DCF!=0), CustomDCFAdvanced(WACC/UFCF/EquityValuePerShare 파싱, costofDebt 키), CustomDCFLevered(FreeCashFlow/OperatingCashFlow 파싱).
- delegation: DiscountedCashFlow("AAPL") path+symbol / CustomDiscountedCashFlow(CustomDCFParams{Symbol:"AAPL", Beta: ptr(1.2), TaxRate: ptr(0.21)}) path+symbol/beta/taxRate(설정한 override 만 쿼리 포함, 미설정 제외 확인).
- 가드: DiscountedCashFlow 빈 symbol, CustomDiscountedCashFlow 빈 Symbol.
- 통합: DiscountedCashFlow("AAPL") DCF>0 / CustomDiscountedCashFlow({Symbol:"AAPL"}) len>0 & WACC>0 / CustomLeveredDiscountedCashFlow({Symbol:"AAPL"}) len>0.

## 문서 / 릴리스
- README DCF 행(4 endpoint).
- `examples/dcf/main.go` — DiscountedCashFlow + CustomDiscountedCashFlow.
- 릴리스 `v0.19.0`.

## 범위 밖 / 위험
- bulk/dcf-bulk(문자열 수치)은 별도 bulk 그룹에서.
- custom monetary 필드 int64 가정(예시 정수) — 소수 반환 시 후속 조정.
- JSON 키 불일치(costofDebt/Stock Price) 테스트로 확인.
- 다음 그룹: crypto/forex/commodity 또는 secFilings.
