# FMP Go SDK — Reports 그룹 (v0.11.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/reports-group`
- 토픽: FMP statements 카테고리의 as-reported(SEC 원문) + 보고서 endpoint 7개. 전체 API 커버리지 캠페인 9번째 그룹. statements 분해의 3/3(마지막).

## 배경 / 목적

statements 27 endpoint 분해의 3번째이자 마지막 하위 그룹. SEC 원문(as-reported) 재무제표와 보고서 메타/링크/10-K JSON. 이 그룹 완료 시 statements 카테고리 전체 커버.

## 결정 사항 (브레인스토밍)

- **범위**: 7 endpoint. 신규 `reports/` 패키지. as-reported 3 + full 1 + latest 1 + dates 1 + 10-K JSON 1.
- **xlsx 제외**: `financial-reports-xlsx` 는 바이너리 파일 다운로드(JSON 아님)로 타입드 SDK 패턴에 부적합 → 범위 제외. 필요 시 `FinancialReportDate.LinkXlsx` URL 로 직접 다운로드 가능(범위 밖 명시).
- **as-reported 공유 구조체**: income/balance/cashflow-as-reported 응답 wrapper 동일(`symbol/fiscalYear/period/reportedCurrency/date/data`) → 단일 `AsReportedStatement` 공유, `Data map[string]json.Number`(int/float 혼재, 문자열 없음). path 만 다른 3 메서드.
- **full 별도 구조체**: `financial-statement-full-as-reported` 의 data 는 숫자+문자열+불리언 혼재 → `AsReportedFull` 의 `Data map[string]any`(json.Number 불가).
- **10-K JSON 원시 반환**: `financial-reports-json` 은 symbol/period/year 외 안정 스키마 없음(수십 동적 섹션) → `[]map[string]any` 원시 반환(타입화 안 함). 정직하게 동적 처리.
- **고정 구조체**: latest-financial-statements(`calendarYear` 주의, symbol 없음 page/limit), financial-reports-dates(linkXlsx/linkJson) 는 작은 고정 필드 struct.
- **internal/fetch 사용**: List/ListBySymbol(calendar/metrics 컨벤션). list 는 빈 결과 ErrNotFound 안 던짐.
- **릴리스**: `v0.11.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | helper | 반환 |
|---|---|---|---|---|
| `reports/as_reported.go` | `IncomeStatementAsReported(ctx, symbol, period, limit)` | `/stable/income-statement-as-reported` | List+guard | `[]AsReportedStatement` |
| | `BalanceSheetStatementAsReported(ctx, symbol, period, limit)` | `/stable/balance-sheet-statement-as-reported` | List+guard | `[]AsReportedStatement` |
| | `CashFlowStatementAsReported(ctx, symbol, period, limit)` | `/stable/cash-flow-statement-as-reported` | List+guard | `[]AsReportedStatement` |
| | `FinancialStatementFullAsReported(ctx, symbol, period, limit)` | `/stable/financial-statement-full-as-reported` | List+guard | `[]AsReportedFull` |
| `reports/reports.go` | `LatestFinancialStatements(ctx, page, limit)` | `/stable/latest-financial-statements` | List | `[]LatestFinancialStatement` |
| | `FinancialReportDates(ctx, symbol)` | `/stable/financial-reports-dates` | ListBySymbol | `[]FinancialReportDate` |
| | `FinancialReportJSON(ctx, symbol string, year int, period string)` | `/stable/financial-reports-json` | List+guard | `[]map[string]any` |
| `reports/client.go` | `New(http)` + helpers | — | — | `*Client` |

helpers (client.go):
- `asReportedParams(symbol, period string, limit int) map[string]string` — symbol 항상, period/limit 조건부.
- as-reported/10-K 메서드는 호출 전 빈 symbol 가드.

## 루트 Client 와이어
```go
type Client struct {
	...
	Metrics *metrics.Client
	Reports *reports.Client // 보고서(as-reported/latest/dates/10-K JSON)
}
```
`NewClient` 에 `c.Reports = reports.New(hc)`. `client_test.go` 에 `TestNewClient_HasReports`.

## 응답 타입 (faithful)

```go
import "encoding/json"

// AsReportedStatement — SEC 원문 재무제표 (income/balance/cashflow-as-reported 공용).
// data 는 XBRL/GAAP 태그(소문자) → 값(int/float 혼재) 동적 맵.
type AsReportedStatement struct {
	Symbol           string                 `json:"symbol"`           // 종목 심볼
	FiscalYear       int                    `json:"fiscalYear"`       // 회계연도(숫자)
	Period           string                 `json:"period"`           // 기간 (FY/Q1..)
	ReportedCurrency *string                `json:"reportedCurrency"` // 보고 통화(null 가능)
	Date             string                 `json:"date"`             // 기준일
	Data             map[string]json.Number `json:"data"`             // GAAP 태그 → 수치
}

// AsReportedFull — SEC 원문 전체 재무제표 (financial-statement-full-as-reported).
// data 값이 숫자/문자열/불리언 혼재 → any.
type AsReportedFull struct {
	Symbol           string         `json:"symbol"`           // 종목 심볼
	FiscalYear       int            `json:"fiscalYear"`       // 회계연도(숫자)
	Period           string         `json:"period"`           // 기간
	ReportedCurrency *string        `json:"reportedCurrency"` // 보고 통화(null 가능)
	Date             string         `json:"date"`             // 기준일
	Data             map[string]any `json:"data"`             // GAAP 태그 → 값(혼합 타입)
}

// LatestFinancialStatement — 최신 재무제표 등록 목록 (latest-financial-statements).
// symbol 미입력, page/limit 페이징. calendarYear 사용(fiscalYear 아님).
type LatestFinancialStatement struct {
	Symbol       string `json:"symbol"`       // 종목 심볼
	CalendarYear int    `json:"calendarYear"` // 달력연도
	Period       string `json:"period"`       // 기간 (Q4 등)
	Date         string `json:"date"`         // 보고 기준일
	DateAdded    string `json:"dateAdded"`    // 등록 일시
}

// FinancialReportDate — 보고서 다운로드 링크 (financial-reports-dates).
type FinancialReportDate struct {
	Symbol     string `json:"symbol"`     // 종목 심볼
	FiscalYear int    `json:"fiscalYear"` // 회계연도
	Period     string `json:"period"`     // 기간 (FY/Q1..)
	LinkXlsx   string `json:"linkXlsx"`   // XLSX 다운로드 URL
	LinkJson   string `json:"linkJson"`   // JSON 다운로드 URL
}
```

`FinancialReportJSON` 은 구조체 없이 `[]map[string]any` 반환(동적 10-K 섹션). 첫 원소에 `symbol/period/year`(year 는 문자열) + 다수 섹션 키.

## 시그니처 규칙
- as-reported 4개: `(ctx, symbol, period string, limit int)` → 빈 symbol 가드 + `fetch.List[T](ctx, c.http, path, asReportedParams(symbol, period, limit))`.
- LatestFinancialStatements: `(ctx, page, limit int)` → `fetch.List[LatestFinancialStatement]` with {page, limit}(둘 다 0 이상이면 포함, page 는 0 도 포함).
- FinancialReportDates: `(ctx, symbol string)` → `fetch.ListBySymbol`.
- FinancialReportJSON: `(ctx, symbol string, year int, period string)` → 빈 symbol 가드 + `fetch.List[map[string]any]` with {symbol, year, period}.

## 테스트
- fixture 단위: AsReportedStatement(data map json.Number — Int64/Float64 파싱 검증, reportedCurrency null→nil, fiscalYear int) / AsReportedFull(data map[string]any 혼합값 — 숫자/문자열 키 둘 다) / LatestFinancialStatement(calendarYear, dateAdded) / FinancialReportDate(linkXlsx/linkJson) / FinancialReportJSON([]map[string]any — symbol/year/섹션 키 존재).
- delegation: IncomeStatementAsReported(symbol,period,limit) path+쿼리 / LatestFinancialStatements(page,limit) path+page/limit / FinancialReportJSON(symbol,year,period) path+쿼리.
- 가드: as-reported/10-K 빈 symbol 대표 1건.
- 통합(`//go:build integration`): IncomeStatementAsReported(AAPL,annual,1) data 비어있지 않음 + grossprofit 키 Int64 파싱 / LatestFinancialStatements(0,5) / FinancialReportDates(AAPL) link 비어있지 않음 / FinancialReportJSON(AAPL,2022,"FY") len>0 — t.Logf 로 첫 원소 키 일부 로그.

## 문서 / 릴리스
- README 커버리지 표 Reports 행 신규(7 endpoint).
- `examples/reports/main.go` — IncomeStatementAsReported + FinancialReportDates.
- 릴리스 `v0.11.0`. 이 그룹으로 statements 카테고리 전체 완료(누적 ~93/263).

## 범위 밖 / 위험
- **financial-reports-xlsx 제외**: 바이너리 다운로드. LinkXlsx 로 대체. 후속 별도 다운로드 메서드 가능.
- `json.Number` 디코딩: 대상 타입이 json.Number 면 표준 디코더가 리터럴 보존(UseNumber 불필요). full 은 문자열 혼재라 any 필수.
- 10-K JSON 동적 스키마 — 타입화 안 함(map). 안정 스키마 없음.
- as-reported data 키는 회사/연도별 가변 — fixture 는 대표 키 일부만.
- 다음 그룹: statements 완료 후 나머지 ~19 카테고리(chart/insiderTrades/secFilings/marketPerformance 등).
