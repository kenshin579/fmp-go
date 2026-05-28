# FMP Go SDK — Statements + Ratios (v0.2.0) 설계

- 작성일: 2026-05-28
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/sdk-statements-ratios`
- 토픽: SDK에 재무제표(손익·대차) + 비율 엔드포인트 3개를 추가해 moneyflow US 재무 통합(서브프로젝트 B)의 데이터 소스 확보

## 배경 / 목적

moneyflow가 국내 종목 재무 탭과 **동일한 4섹션(손익·재무상태·비율·성장률, 15필드)** 을 해외(US) 종목에도 제공하려고 한다. 현재 `fmp-go` SDK는 `Company.Profile`만 노출하므로, US 재무 데이터를 위해 **FMP의 3개 엔드포인트**를 SDK에 추가하고 v0.2.0으로 릴리스한다.

(서브프로젝트 B의 통합 작업은 moneyflow 레포의 별도 spec에서 다루며, 본 SDK의 v0.2.0 태그 릴리스 이후 시작.)

## 결정 사항 (브레인스토밍)

- **범위**: FMP의 income-statement / balance-sheet-statement / ratios 세 엔드포인트를 SDK에 추가. cash-flow-statement / key-metrics 등은 후속(범위 밖).
- **패키지 분리**: `Company`(기존)와 동격 위상으로 두 서비스 서브패키지 신규.
  - `statements` — 재무제표(income + balance) 묶음.
  - `ratios` — 재무비율(ratios) 단독.
- **API 기준선**: 기존 v0.1.0과 동일하게 FMP **stable** 엔드포인트.
- **응답 형태**: 각 엔드포인트는 객체 배열(여러 기간). SDK는 배열 그대로 반환(Profile은 첫 요소만 반환했지만, 재무는 시계열이라 배열 노출).
- **공통 파라미터**: `{ Symbol string; Period string; Limit int }`. `Period`는 `"annual"`/`"quarter"`, 기본 빈 값(FMP 기본 = annual).
- **필드 매핑**: FMP가 주는 모든 필드를 응답 struct에 노출(faithful). moneyflow는 필요한 부분 선택해 소비.
- **에러**: 빈 배열 → `httpclient.ErrNotFound`(Company.Profile과 동일 sentinel 재사용).

## 데이터 소스 (FMP 엔드포인트)

| 엔드포인트 | 경로 | 응답 핵심 필드(moneyflow가 소비) |
|---|---|---|
| Income Statement | `GET /stable/income-statement?symbol=X&period=...&limit=...` | date, period, reportedCurrency, revenue, costOfRevenue, grossProfit, operatingIncome, netIncome, eps, ebitda 등 |
| Balance Sheet Statement | `GET /stable/balance-sheet-statement?symbol=X&period=...&limit=...` | date, period, reportedCurrency, totalAssets, totalLiabilities, totalStockholdersEquity, totalEquity 등 |
| Ratios | `GET /stable/ratios?symbol=X&period=...&limit=...` | date, period, returnOnEquity, debtRatio, debtToEquityRatio, bookValuePerShare, priceToEarningsRatio 등 |

(전체 필드는 카탈로그 `docs/api/<...>/*.md`에 이미 캡처돼 있음 — Task 1에서 정확 필드 목록을 참조해 struct 정의 확정.)

## 아키텍처 (Company.Profile 패턴 차용)

```
fmp-go/
├── client.go               # Client 에 Statements, Ratios 필드 추가
├── statements/
│   ├── client.go           # statements.New(hc); IncomeStatement(...), BalanceSheetStatement(...)
│   ├── income.go           # IncomeStatement struct + IncomeStatement 메서드
│   ├── balance.go          # BalanceSheetStatement struct + BalanceSheetStatement 메서드
│   └── *_test.go           # fixture/스텁 단위테스트
├── ratios/
│   ├── client.go           # ratios.New(hc); Ratios(...)
│   ├── ratios.go           # Ratio struct + Ratios 메서드
│   └── *_test.go
└── docs/api/...            # (선택) 카탈로그 보강
```

루트 `Client`:
```go
type Client struct {
    http       *httpclient.Client
    Company    *company.Client
    Statements *statements.Client
    Ratios     *ratios.Client
}
```
`NewClient`에서 `c.Statements = statements.New(hc)`, `c.Ratios = ratios.New(hc)` 추가.

### 메서드 시그니처
```go
// statements 패키지
type Params struct {
    Symbol string
    Period string // "annual" | "quarter" (빈 값 → FMP 기본 annual)
    Limit  int    // 0 → 기본
}

func (c *Client) IncomeStatement(ctx context.Context, p Params) ([]IncomeStatement, error)
func (c *Client) BalanceSheetStatement(ctx context.Context, p Params) ([]BalanceSheetStatement, error)

// ratios 패키지
type Params struct { Symbol, Period string; Limit int }
func (c *Client) Ratios(ctx context.Context, p Params) ([]Ratio, error)
```

각 메서드: `httpclient.GetJSON(ctx, "/stable/...", params(map[string]string), &out)`. 빈 결과면 `httpclient.ErrNotFound`. 단일 종목/기간 시 `out[0]`만 쓰는 게 아니라 **시계열 슬라이스 그대로 반환**(`Company.Profile`이 첫 요소만 반환했던 것과 다름 — 재무는 본질적으로 다기간).

### 응답 구조체 (faithful)
각 응답 struct는 FMP가 주는 모든 필드를 노출(`Profile`과 동일 원칙). 정확한 필드 목록·타입은 카탈로그(`docs/api/statements/income-statement.md` 등) + 실응답 fixture로 확정 — 구현 시 자동 캐치(Task 1에서 fixture 캡처).

예시 minimum 필드:
```go
type IncomeStatement struct {
    Date              string  `json:"date"`               // "2024-09-28"
    Symbol            string  `json:"symbol"`
    ReportedCurrency  string  `json:"reportedCurrency"`   // "USD"
    Period            string  `json:"period"`             // "FY"/"Q1".."Q4"
    Revenue           int64   `json:"revenue"`
    CostOfRevenue     int64   `json:"costOfRevenue"`
    GrossProfit       int64   `json:"grossProfit"`
    OperatingIncome   int64   `json:"operatingIncome"`
    NetIncome         int64   `json:"netIncome"`
    EPS               float64 `json:"eps"`
    EBITDA            int64   `json:"ebitda"`
    // ... 그 외 FMP가 주는 모든 필드(구현 시 fixture로 확정 추가)
}
```

## 에러 / 결측
- 빈 배열 → `httpclient.ErrNotFound`.
- 비-200 / `Error Message` 바디 → `httpclient.APIError`(기존 매핑 재사용).
- 일부 필드 결측 → JSON 기본값(0/"")으로 디코드. 소비 측 책임.

## 테스트
- **순수 매퍼**: 각 응답을 `IncomeStatement`/`BalanceSheetStatement`/`Ratio` 슬라이스로 디코드 — fixture JSON으로 핵심 필드 매핑 검증(faithful 디코딩).
- **Delegation**: stub backend로 Client 메서드가 internal/httpclient에 위임하는지(Company.Profile 테스트 패턴).
- **에러 경로**: ErrNotFound(빈 배열), APIError(비-200) — httpclient 레이어가 이미 테스트하므로 SDK 메서드 1개만 sample 확인.
- **통합 테스트**(build tag): `FMP_API_KEY` 있으면 실 AAPL annual/quarter 호출로 계약 검증(키 없으면 Skip).

## 릴리스
- 작업 완료 후 `main`에 머지 → **`scripts/release.sh v0.2.0`** → 태그 push + GitHub Release.
- moneyflow `go.mod`에 `github.com/kenshin579/fmp-go v0.2.0` bump.

## 범위 밖 / 후속
- cash-flow-statement / key-metrics / income-statement-growth 엔드포인트.
- 영문 카탈로그 갱신(필요 시).

## 위험 / 주의
- FMP 응답 필드 타입(int64 vs float)이 ssue될 수 있음 — 매우 큰 금액(>2^53)이면 float64 정밀도 손실. 일반 상장사는 int64 안전. 카탈로그 응답을 fixture로 캡처해 매핑 시 확정.
- `Period` 파라미터: `"annual"`/`"quarter"`. FMP는 `quarter` 단수형 사용 — typo 주의.
- `Limit` 파라미터: 0이면 SDK는 쿼리에 포함 안 함(FMP 기본 적용).
