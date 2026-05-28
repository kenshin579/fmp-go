# FMP Go SDK — Statements + Ratios (v0.2.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** SDK에 FMP 재무제표(`/stable/income-statement`, `/stable/balance-sheet-statement`)와 재무비율(`/stable/ratios`) 엔드포인트 3개를 추가해 moneyflow US 재무 통합(서브프로젝트 B)의 데이터 소스 확보, `v0.2.0` 릴리스 준비.

**Architecture:** 기존 `Company.Profile`(v0.1.0)과 동일한 sub-client 패턴. 신규 `statements` 패키지(IncomeStatement + BalanceSheetStatement 메서드)와 `ratios` 패키지(Ratios 메서드)를 추가하고 `fmp.Client`에 필드로 노출. 각 메서드는 시계열 배열을 반환(`[]T`). 응답 매핑은 카탈로그 fixture로 fidelity 검증.

**Tech Stack:** Go 1.25+, 표준 라이브러리(`net/http/httptest`), `internal/httpclient`(이미 구현됨). 외부 의존성 추가 없음.

**Spec:** `docs/superpowers/specs/2026-05-28-statements-ratios-design.md`

**Repo:** `github.com/kenshin579/fmp-go`, branch `feature/sdk-statements-ratios`. Go: `unset GOROOT` if GOROOT error.

**확정된 사실(조사 완료):**
- 카탈로그 파일 존재: `docs/api/statements/income-statement.md`, `docs/api/statements/balance-sheet-statement.md`, `docs/api/statements/metrics-ratios.md`(엔드포인트 URL `/stable/ratios`). 각 md에 응답 예시 JSON 포함 — fixture 소스.
- 기존 패턴: `company/{client.go, profile.go, profile_test.go, testdata/profile_aapl.json}` — 신규 패키지가 그대로 모방.
- 기존 `internal/httpclient.Client`의 `GetJSON(ctx, path, params, out)` 사용. 빈 결과 → `httpclient.ErrNotFound`.
- 기존 `fmp.Client`는 `Company *company.Client`만 보유. 신규로 `Statements`, `Ratios` 추가.

---

## File Structure
- Create: `statements/client.go` — Statements sub-client + `Params` 타입.
- Create: `statements/income.go` — `IncomeStatement` 구조체 + `IncomeStatement(ctx, p)` 메서드.
- Create: `statements/balance.go` — `BalanceSheetStatement` 구조체 + `BalanceSheetStatement(ctx, p)` 메서드.
- Create: `statements/income_test.go`, `statements/balance_test.go` — fixture/스텁 단위테스트.
- Create: `statements/testdata/income-statement-aapl.json`, `statements/testdata/balance-sheet-statement-aapl.json`.
- Create: `ratios/client.go` — Ratios sub-client + `Params` 타입.
- Create: `ratios/ratios.go` — `Ratio` 구조체 + `Ratios(ctx, p)` 메서드.
- Create: `ratios/ratios_test.go`, `ratios/testdata/ratios-aapl.json`.
- Modify: `client.go` — `fmp.Client`에 `Statements *statements.Client` + `Ratios *ratios.Client` 추가.
- Modify: `client_test.go` — NewClient 테스트에 `c.Statements`/`c.Ratios` non-nil 어셔션 추가.
- Modify: `integration_test.go` — (선택) statements/ratios 통합테스트 추가.

---

## Task 1: `statements` 패키지 — IncomeStatement (TDD)

**Files:** `statements/client.go`, `statements/income.go`, `statements/income_test.go`, `statements/testdata/income-statement-aapl.json`

- [ ] **Step 1: fixture 확보**

`docs/api/statements/income-statement.md`의 `## Response (example)` json 블록 내용을 그대로 `statements/testdata/income-statement-aapl.json`에 저장(JSON 배열, AAPL FY2024). `FMP_API_KEY`가 설정돼 있으면 라이브 응답으로 대체 가능(더 정확):
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
mkdir -p statements/testdata
if [ -n "$FMP_API_KEY" ]; then
  curl -sS "https://financialmodelingprep.com/stable/income-statement?symbol=AAPL&period=annual&limit=2&apikey=$FMP_API_KEY" \
    | python3 -m json.tool > statements/testdata/income-statement-aapl.json
fi
```
키 없으면 docs의 예시 JSON을 그대로 옮김. 결과는 유효한 JSON 배열이어야 한다:
```bash
python3 -m json.tool statements/testdata/income-statement-aapl.json > /dev/null && echo "valid JSON"
```

- [ ] **Step 2: 실패하는 테스트 작성**

Create `statements/client.go`:
```go
// Package statements 는 FMP 재무제표 API sub-client.
// fmp.Client.Statements 로 접근.
package statements

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 재무제표 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// Params 는 재무제표 조회 공통 파라미터.
type Params struct {
	Symbol string
	Period string // "annual" | "quarter" (빈 값 → FMP 기본 annual)
	Limit  int    // 0 → 쿼리에 미포함(FMP 기본)
}

// queryParams 는 Params 를 httpclient 쿼리 맵으로 변환한다.
func (p Params) queryParams() map[string]string {
	q := map[string]string{"symbol": p.Symbol}
	if p.Period != "" {
		q["period"] = p.Period
	}
	if p.Limit > 0 {
		q["limit"] = fmtInt(p.Limit)
	}
	return q
}

// fmtInt — strconv.Itoa 보다 명시적 의도.
func fmtInt(n int) string { return strconvItoa(n) }
```
`fmtInt`는 임시 — Step 3에서 표준 `strconv`로 정리한다(빌드 깨짐 의도적, 컴파일 강제 실패).

Create `statements/income_test.go`:
```go
package statements

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func newTestClient(t *testing.T, status int, body string) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, srv.Close
}

func TestIncomeStatement_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/income-statement-aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil {
		t.Fatalf("fixture is not a JSON array: %v", err)
	}
	if len(arr) == 0 {
		t.Fatal("fixture array empty")
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.IncomeStatement(context.Background(), Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("IncomeStatement: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("rows empty")
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.Revenue <= 0 {
		t.Errorf("Revenue = %d, want > 0", r.Revenue)
	}
	if r.GrossProfit <= 0 {
		t.Errorf("GrossProfit = %d, want > 0", r.GrossProfit)
	}
	if r.OperatingIncome <= 0 {
		t.Errorf("OperatingIncome = %d, want > 0", r.OperatingIncome)
	}
	if r.NetIncome <= 0 {
		t.Errorf("NetIncome = %d, want > 0", r.NetIncome)
	}
	if r.EPS == 0 {
		t.Error("EPS not parsed")
	}
	if r.ReportedCurrency != "USD" {
		t.Errorf("ReportedCurrency = %q", r.ReportedCurrency)
	}
}

func TestIncomeStatement_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.IncomeStatement(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestIncomeStatement_EmptySymbol(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	if _, err := c.IncomeStatement(context.Background(), Params{Symbol: "  "}); err == nil {
		t.Fatal("expected error for empty symbol")
	}
}
```

- [ ] **Step 3: 실패 확인**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go test ./statements/
```
Expected: 컴파일 실패(`strconvItoa` 미정의, `IncomeStatement` 타입/메서드 미정의).

- [ ] **Step 4: 구현**

Fix `statements/client.go` — `fmtInt` 구현을 `strconv.Itoa`로 교체:
```go
package statements

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 재무제표 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// Params 는 재무제표 조회 공통 파라미터.
type Params struct {
	Symbol string
	Period string // "annual" | "quarter" (빈 값 → FMP 기본 annual)
	Limit  int    // 0 → 쿼리 미포함
}

// queryParams 는 Params 를 httpclient 쿼리 맵으로 변환한다.
func (p Params) queryParams() map[string]string {
	q := map[string]string{"symbol": p.Symbol}
	if p.Period != "" {
		q["period"] = p.Period
	}
	if p.Limit > 0 {
		q["limit"] = strconv.Itoa(p.Limit)
	}
	return q
}
```

Create `statements/income.go`:
```go
package statements

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// IncomeStatement 는 FMP /stable/income-statement 응답 한 기간(연간/분기).
// 필드는 FMP 응답을 충실히 매핑한다(faithful).
type IncomeStatement struct {
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	CIK                                     string  `json:"cik"`
	FilingDate                              string  `json:"filingDate"`
	AcceptedDate                            string  `json:"acceptedDate"`
	FiscalYear                              string  `json:"fiscalYear"`
	Period                                  string  `json:"period"` // "FY"/"Q1"..
	Revenue                                 int64   `json:"revenue"`
	CostOfRevenue                           int64   `json:"costOfRevenue"`
	GrossProfit                             int64   `json:"grossProfit"`
	ResearchAndDevelopmentExpenses          int64   `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        int64   `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             int64   `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses int64   `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           int64   `json:"otherExpenses"`
	OperatingExpenses                       int64   `json:"operatingExpenses"`
	CostAndExpenses                         int64   `json:"costAndExpenses"`
	NetInterestIncome                       int64   `json:"netInterestIncome"`
	InterestIncome                          int64   `json:"interestIncome"`
	InterestExpense                         int64   `json:"interestExpense"`
	DepreciationAndAmortization             int64   `json:"depreciationAndAmortization"`
	EBITDA                                  int64   `json:"ebitda"`
	EBIT                                    int64   `json:"ebit"`
	NonOperatingIncomeExcludingInterest     int64   `json:"nonOperatingIncomeExcludingInterest"`
	OperatingIncome                         int64   `json:"operatingIncome"`
	TotalOtherIncomeExpensesNet             int64   `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         int64   `json:"incomeBeforeTax"`
	IncomeTaxExpense                        int64   `json:"incomeTaxExpense"`
	NetIncomeFromContinuingOperations       int64   `json:"netIncomeFromContinuingOperations"`
	NetIncomeFromDiscontinuedOperations     int64   `json:"netIncomeFromDiscontinuedOperations"`
	OtherAdjustmentsToNetIncome             int64   `json:"otherAdjustmentsToNetIncome"`
	NetIncome                               int64   `json:"netIncome"`
	NetIncomeDeductions                     int64   `json:"netIncomeDeductions"`
	BottomLineNetIncome                     int64   `json:"bottomLineNetIncome"`
	EPS                                     float64 `json:"eps"`
	EPSDiluted                              float64 `json:"epsDiluted"`
	WeightedAverageShsOut                   int64   `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                int64   `json:"weightedAverageShsOutDil"`
}

// IncomeStatement 는 종목의 손익계산서 시계열을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) IncomeStatement(ctx context.Context, p Params) ([]IncomeStatement, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []IncomeStatement
	if err := c.http.GetJSON(ctx, "/stable/income-statement", p.queryParams(), &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
```

- [ ] **Step 5: 통과 확인**

Run: `go test ./statements/ -v && go vet ./statements/`
Expected: 3개 테스트 PASS.

- [ ] **Step 6: Commit**

```bash
git add statements/
git commit -m "feat(statements): IncomeStatement 엔드포인트 — /stable/income-statement"
```

---

## Task 2: `statements/balance.go` — BalanceSheetStatement (TDD)

**Files:** `statements/balance.go`, `statements/balance_test.go`, `statements/testdata/balance-sheet-statement-aapl.json`

- [ ] **Step 1: fixture 확보**

`docs/api/statements/balance-sheet-statement.md`의 `## Response (example)` json 블록을 `statements/testdata/balance-sheet-statement-aapl.json`에 저장(JSON 배열). 키 있으면 라이브:
```bash
[ -n "$FMP_API_KEY" ] && curl -sS "https://financialmodelingprep.com/stable/balance-sheet-statement?symbol=AAPL&period=annual&limit=2&apikey=$FMP_API_KEY" | python3 -m json.tool > statements/testdata/balance-sheet-statement-aapl.json
python3 -m json.tool statements/testdata/balance-sheet-statement-aapl.json > /dev/null && echo "valid JSON"
```

- [ ] **Step 2: 실패하는 테스트 작성**

Create `statements/balance_test.go`:
```go
package statements

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func TestBalanceSheetStatement_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/balance-sheet-statement-aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil || len(arr) == 0 {
		t.Fatalf("fixture invalid/empty: %v", err)
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.BalanceSheetStatement(context.Background(), Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("BalanceSheetStatement: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("rows empty")
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", r.Symbol)
	}
	if r.TotalAssets <= 0 {
		t.Errorf("TotalAssets = %d", r.TotalAssets)
	}
	if r.TotalLiabilities <= 0 {
		t.Errorf("TotalLiabilities = %d", r.TotalLiabilities)
	}
	if r.TotalStockholdersEquity <= 0 {
		t.Errorf("TotalStockholdersEquity = %d", r.TotalStockholdersEquity)
	}
}

func TestBalanceSheetStatement_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.BalanceSheetStatement(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
```

- [ ] **Step 3: 실패 확인** — `go test ./statements/ -run BalanceSheet` → 컴파일 실패(`BalanceSheetStatement` 타입/메서드 미정의).

- [ ] **Step 4: 구현 (`statements/balance.go`)**

핵심 필드(테스트가 검증) + faithful 추가 필드. 카탈로그 `balance-sheet-statement.md`의 응답 예시에서 모든 키를 struct로 정의:
```go
package statements

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// BalanceSheetStatement 는 FMP /stable/balance-sheet-statement 응답 한 기간.
type BalanceSheetStatement struct {
	Date                                    string `json:"date"`
	Symbol                                  string `json:"symbol"`
	ReportedCurrency                        string `json:"reportedCurrency"`
	CIK                                     string `json:"cik"`
	FilingDate                              string `json:"filingDate"`
	AcceptedDate                            string `json:"acceptedDate"`
	FiscalYear                              string `json:"fiscalYear"`
	Period                                  string `json:"period"`
	CashAndCashEquivalents                  int64  `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    int64  `json:"shortTermInvestments"`
	CashAndShortTermInvestments             int64  `json:"cashAndShortTermInvestments"`
	NetReceivables                          int64  `json:"netReceivables"`
	AccountsReceivables                     int64  `json:"accountsReceivables"`
	OtherReceivables                        int64  `json:"otherReceivables"`
	Inventory                               int64  `json:"inventory"`
	Prepaids                                int64  `json:"prepaids"`
	OtherCurrentAssets                      int64  `json:"otherCurrentAssets"`
	TotalCurrentAssets                      int64  `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               int64  `json:"propertyPlantEquipmentNet"`
	Goodwill                                int64  `json:"goodwill"`
	IntangibleAssets                        int64  `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             int64  `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     int64  `json:"longTermInvestments"`
	TaxAssets                               int64  `json:"taxAssets"`
	OtherNonCurrentAssets                   int64  `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   int64  `json:"totalNonCurrentAssets"`
	OtherAssets                             int64  `json:"otherAssets"`
	TotalAssets                             int64  `json:"totalAssets"`
	TotalPayables                           int64  `json:"totalPayables"`
	AccountPayables                         int64  `json:"accountPayables"`
	OtherPayables                           int64  `json:"otherPayables"`
	AccruedExpenses                         int64  `json:"accruedExpenses"`
	ShortTermDebt                           int64  `json:"shortTermDebt"`
	CapitalLeaseObligationsCurrent          int64  `json:"capitalLeaseObligationsCurrent"`
	TaxPayables                             int64  `json:"taxPayables"`
	DeferredRevenue                         int64  `json:"deferredRevenue"`
	OtherCurrentLiabilities                 int64  `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 int64  `json:"totalCurrentLiabilities"`
	LongTermDebt                            int64  `json:"longTermDebt"`
	CapitalLeaseObligationsNonCurrent       int64  `json:"capitalLeaseObligationsNonCurrent"`
	DeferredRevenueNonCurrent               int64  `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        int64  `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              int64  `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              int64  `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        int64  `json:"otherLiabilities"`
	CapitalLeaseObligations                 int64  `json:"capitalLeaseObligations"`
	TotalLiabilities                        int64  `json:"totalLiabilities"`
	TreasuryStock                           int64  `json:"treasuryStock"`
	PreferredStock                          int64  `json:"preferredStock"`
	CommonStock                             int64  `json:"commonStock"`
	RetainedEarnings                        int64  `json:"retainedEarnings"`
	AdditionalPaidInCapital                 int64  `json:"additionalPaidInCapital"`
	AccumulatedOtherComprehensiveIncomeLoss int64  `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OtherTotalStockholdersEquity            int64  `json:"otherTotalStockholdersEquity"`
	TotalStockholdersEquity                 int64  `json:"totalStockholdersEquity"`
	TotalEquity                             int64  `json:"totalEquity"`
	MinorityInterest                        int64  `json:"minorityInterest"`
	TotalLiabilitiesAndTotalEquity          int64  `json:"totalLiabilitiesAndTotalEquity"`
	TotalInvestments                        int64  `json:"totalInvestments"`
	TotalDebt                               int64  `json:"totalDebt"`
	NetDebt                                 int64  `json:"netDebt"`
}

// BalanceSheetStatement 는 종목의 대차대조표 시계열을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) BalanceSheetStatement(ctx context.Context, p Params) ([]BalanceSheetStatement, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []BalanceSheetStatement
	if err := c.http.GetJSON(ctx, "/stable/balance-sheet-statement", p.queryParams(), &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
```

> 카탈로그의 fixture 내 응답 키와 비교해 빠진 필드/타입 불일치가 있으면 추가/수정 후 테스트 재실행. 카탈로그가 라이브 응답과 약간 다를 수 있다.

- [ ] **Step 5: 통과 확인** — `go test ./statements/ -v && go vet ./statements/` → 새 테스트 + Task 1 테스트 전부 PASS.

- [ ] **Step 6: Commit**
```bash
git add statements/balance.go statements/balance_test.go statements/testdata/balance-sheet-statement-aapl.json
git commit -m "feat(statements): BalanceSheetStatement 엔드포인트 — /stable/balance-sheet-statement"
```

---

## Task 3: `ratios` 패키지 — Ratios (TDD)

**Files:** `ratios/client.go`, `ratios/ratios.go`, `ratios/ratios_test.go`, `ratios/testdata/ratios-aapl.json`

- [ ] **Step 1: fixture 확보**

`docs/api/statements/metrics-ratios.md`의 응답 예시(엔드포인트 `/stable/ratios`)를 `ratios/testdata/ratios-aapl.json`에 저장:
```bash
mkdir -p ratios/testdata
[ -n "$FMP_API_KEY" ] && curl -sS "https://financialmodelingprep.com/stable/ratios?symbol=AAPL&period=annual&limit=2&apikey=$FMP_API_KEY" | python3 -m json.tool > ratios/testdata/ratios-aapl.json
python3 -m json.tool ratios/testdata/ratios-aapl.json > /dev/null && echo "valid JSON"
```

- [ ] **Step 2: client.go + 실패 테스트**

Create `ratios/client.go`:
```go
// Package ratios 는 FMP 재무비율 API sub-client.
// fmp.Client.Ratios 로 접근.
package ratios

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 재무비율 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// Params 는 재무비율 조회 파라미터.
type Params struct {
	Symbol string
	Period string // "annual" | "quarter"
	Limit  int
}

func (p Params) queryParams() map[string]string {
	q := map[string]string{"symbol": p.Symbol}
	if p.Period != "" {
		q["period"] = p.Period
	}
	if p.Limit > 0 {
		q["limit"] = strconv.Itoa(p.Limit)
	}
	return q
}
```

Create `ratios/ratios_test.go`:
```go
package ratios

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func newTestClient(t *testing.T, status int, body string) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, srv.Close
}

func TestRatios_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/ratios-aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil || len(arr) == 0 {
		t.Fatalf("fixture invalid/empty: %v", err)
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	rows, err := c.Ratios(context.Background(), Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("Ratios: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("rows empty")
	}
	r := rows[0]
	if r.Symbol != "AAPL" {
		t.Errorf("Symbol = %q", r.Symbol)
	}
	// 핵심 매핑: BPS, 부채비율(D/E), 매출총이익률 정도 확인
	if r.BookValuePerShare == 0 {
		t.Error("BookValuePerShare not parsed")
	}
	if r.DebtToEquityRatio == 0 {
		t.Error("DebtToEquityRatio not parsed")
	}
}

func TestRatios_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()
	_, err := c.Ratios(context.Background(), Params{Symbol: "NOPE"})
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
```

- [ ] **Step 3: 실패 확인** — `cd /Users/frankoh/src/workspace_moneyflow/fmp-go && go test ./ratios/` → 컴파일 실패(`Ratio` 타입/메서드 미정의).

- [ ] **Step 4: 구현 (`ratios/ratios.go`)**

`metrics-ratios.md`의 응답 키 전체를 struct에 매핑. 핵심(moneyflow가 소비) 필드는 명시적으로 포함:
```go
package ratios

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Ratio 는 FMP /stable/ratios 응답 한 기간.
// FMP 응답 키를 충실히 매핑(faithful). 모든 비율은 소수(0.XX), 백분율 환산은 소비 측 책임.
type Ratio struct {
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	FiscalYear                              string  `json:"fiscalYear"`
	Period                                  string  `json:"period"`
	GrossProfitMargin                       float64 `json:"grossProfitMargin"`
	EBITMargin                              float64 `json:"ebitMargin"`
	EBITDAMargin                            float64 `json:"ebitdaMargin"`
	OperatingProfitMargin                   float64 `json:"operatingProfitMargin"`
	PretaxProfitMargin                      float64 `json:"pretaxProfitMargin"`
	ContinuousOperationsProfitMargin        float64 `json:"continuousOperationsProfitMargin"`
	NetProfitMargin                         float64 `json:"netProfitMargin"`
	BottomLineProfitMargin                  float64 `json:"bottomLineProfitMargin"`
	ReceivablesTurnover                     float64 `json:"receivablesTurnover"`
	PayablesTurnover                        float64 `json:"payablesTurnover"`
	InventoryTurnover                       float64 `json:"inventoryTurnover"`
	FixedAssetTurnover                      float64 `json:"fixedAssetTurnover"`
	AssetTurnover                           float64 `json:"assetTurnover"`
	CurrentRatio                            float64 `json:"currentRatio"`
	QuickRatio                              float64 `json:"quickRatio"`
	SolvencyRatio                           float64 `json:"solvencyRatio"`
	CashRatio                               float64 `json:"cashRatio"`
	PriceToEarningsRatio                    float64 `json:"priceToEarningsRatio"`
	PriceToEarningsGrowthRatio              float64 `json:"priceToEarningsGrowthRatio"`
	ForwardPriceToEarningsGrowthRatio       float64 `json:"forwardPriceToEarningsGrowthRatio"`
	PriceToBookRatio                        float64 `json:"priceToBookRatio"`
	PriceToSalesRatio                       float64 `json:"priceToSalesRatio"`
	PriceToFreeCashFlowRatio                float64 `json:"priceToFreeCashFlowRatio"`
	PriceToOperatingCashFlowRatio           float64 `json:"priceToOperatingCashFlowRatio"`
	DebtToAssetsRatio                       float64 `json:"debtToAssetsRatio"`
	DebtToEquityRatio                       float64 `json:"debtToEquityRatio"`
	DebtToCapitalRatio                      float64 `json:"debtToCapitalRatio"`
	LongTermDebtToCapitalRatio              float64 `json:"longTermDebtToCapitalRatio"`
	FinancialLeverageRatio                  float64 `json:"financialLeverageRatio"`
	WorkingCapitalTurnoverRatio             float64 `json:"workingCapitalTurnoverRatio"`
	OperatingCashFlowRatio                  float64 `json:"operatingCashFlowRatio"`
	OperatingCashFlowSalesRatio             float64 `json:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio      float64 `json:"freeCashFlowOperatingCashFlowRatio"`
	DebtServiceCoverageRatio                float64 `json:"debtServiceCoverageRatio"`
	InterestCoverageRatio                   float64 `json:"interestCoverageRatio"`
	ShortTermOperatingCashFlowCoverageRatio float64 `json:"shortTermOperatingCashFlowCoverageRatio"`
	OperatingCashFlowCoverageRatio          float64 `json:"operatingCashFlowCoverageRatio"`
	CapitalExpenditureCoverageRatio         float64 `json:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio       float64 `json:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio                     float64 `json:"dividendPayoutRatio"`
	DividendYield                           float64 `json:"dividendYield"`
	DividendYieldPercentage                 float64 `json:"dividendYieldPercentage"`
	RevenuePerShare                         float64 `json:"revenuePerShare"`
	NetIncomePerShare                       float64 `json:"netIncomePerShare"`
	InterestDebtPerShare                    float64 `json:"interestDebtPerShare"`
	CashPerShare                            float64 `json:"cashPerShare"`
	BookValuePerShare                       float64 `json:"bookValuePerShare"`
	TangibleBookValuePerShare               float64 `json:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare              float64 `json:"shareholdersEquityPerShare"`
	OperatingCashFlowPerShare               float64 `json:"operatingCashFlowPerShare"`
	CapexPerShare                           float64 `json:"capexPerShare"`
	FreeCashFlowPerShare                    float64 `json:"freeCashFlowPerShare"`
	NetIncomePerEBT                         float64 `json:"netIncomePerEBT"`
	EBTPerEBIT                              float64 `json:"ebtPerEbit"`
	PriceToFairValue                        float64 `json:"priceToFairValue"`
	DebtToMarketCap                         float64 `json:"debtToMarketCap"`
	EffectiveTaxRate                        float64 `json:"effectiveTaxRate"`
	EnterpriseValueMultiple                 float64 `json:"enterpriseValueMultiple"`
}

// Ratios 는 종목의 재무비율 시계열을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) Ratios(ctx context.Context, p Params) ([]Ratio, error) {
	if strings.TrimSpace(p.Symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	var out []Ratio
	if err := c.http.GetJSON(ctx, "/stable/ratios", p.queryParams(), &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return out, nil
}
```

> 카탈로그의 응답 예시 키 전체와 비교해 빠진 필드(예: `dividendYieldPercentage`가 카탈로그에 없을 수도) 발견 시 add/remove. 핵심은 fixture가 디코드 통과하는지.

- [ ] **Step 5: 통과 확인** — `cd /Users/frankoh/src/workspace_moneyflow/fmp-go && go test ./ratios/ -v && go vet ./ratios/` → PASS.

- [ ] **Step 6: Commit**
```bash
git add ratios/
git commit -m "feat(ratios): Ratios 엔드포인트 — /stable/ratios"
```

---

## Task 4: 루트 `fmp.Client` 와이어링

**Files:** `client.go`, `client_test.go`

- [ ] **Step 1: 실패하는 테스트 작성**

`client_test.go`에 NewClient/NewClientFromEnv가 신규 sub-client를 노출하는지 어셔션 추가. 기존 `TestNewClient_HasCompany`/`TestNewClientFromEnv_Reads` 다음에 추가:
```go
func TestNewClient_HasStatementsAndRatios(t *testing.T) {
	c, err := fmp.NewClient("k123")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Statements == nil {
		t.Fatal("Statements nil")
	}
	if c.Ratios == nil {
		t.Fatal("Ratios nil")
	}
}
```

- [ ] **Step 2: 실패 확인** — `cd /Users/frankoh/src/workspace_moneyflow/fmp-go && go test . -run TestNewClient_HasStatementsAndRatios` → 컴파일 실패(`c.Statements`/`c.Ratios` 미정의).

- [ ] **Step 3: 구현 — client.go 수정**

`client.go`의 imports에 추가:
```go
	"github.com/kenshin579/fmp-go/ratios"
	"github.com/kenshin579/fmp-go/statements"
```
`Client` 구조체에 필드 추가:
```go
type Client struct {
	http *httpclient.Client

	Company    *company.Client
	Statements *statements.Client
	Ratios     *ratios.Client
}
```
`NewClient` 본문에서 sub-client 생성부에 추가(`c.Company = company.New(hc)` 다음):
```go
	c.Statements = statements.New(hc)
	c.Ratios = ratios.New(hc)
```

- [ ] **Step 4: 통과 확인**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./...
```
Expected: 전체 PASS(신규 + 기존).

- [ ] **Step 5: Commit**
```bash
git add client.go client_test.go
git commit -m "feat: fmp.Client 에 Statements, Ratios sub-client 노출"
```

---

## Task 5: 통합 테스트 + 전체 검증

**Files:** Modify `integration_test.go`

- [ ] **Step 1: 통합 테스트(build tag) 보강**

`integration_test.go`에 두 테스트 추가(기존 `TestIntegration_CompanyProfile` 다음, `//go:build integration` 파일 안):
```go
func TestIntegration_IncomeStatement(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY not set")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	rows, err := c.Statements.IncomeStatement(context.Background(), statements.Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("IncomeStatement: %v", err)
	}
	if len(rows) == 0 || rows[0].Revenue <= 0 {
		t.Errorf("unexpected: %+v", rows)
	}
}

func TestIntegration_Ratios(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY not set")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	rows, err := c.Ratios.Ratios(context.Background(), ratios.Params{Symbol: "AAPL", Period: "annual", Limit: 2})
	if err != nil {
		t.Fatalf("Ratios: %v", err)
	}
	if len(rows) == 0 {
		t.Error("empty ratios rows")
	}
}
```
imports에 추가(`integration_test.go` 상단):
```go
import (
	// ... 기존 imports
	"github.com/kenshin579/fmp-go/ratios"
	"github.com/kenshin579/fmp-go/statements"
)
```

- [ ] **Step 2: 전체 빌드/테스트**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
unset GOROOT
go build ./... && go vet ./... && go test ./... && gofmt -l . | grep -v node_modules
```
Expected: build/vet/test 전부 PASS. gofmt 출력은 비어야 함(기존 미정렬 파일 제외).

- [ ] **Step 3: 통합 테스트 (선택, FMP_API_KEY 있을 때)**

Run: `go test -tags integration -run TestIntegration ./...`
Expected: PASS(키 있을 때) 또는 Skip(없을 때).

- [ ] **Step 4: Commit**
```bash
git add integration_test.go
git commit -m "test: statements/ratios 통합 테스트 (build tag integration)"
```

- [ ] **Step 5: 릴리스 안내 (실행은 사용자가)**

본 plan 범위 내에서 자동 릴리스하지 않는다. PR 머지 후, 사용자가 `./scripts/release.sh v0.2.0`을 실행해 태그 + GitHub Release를 생성한다. moneyflow 통합(서브프로젝트 B)은 이 태그 릴리스 후 시작.

---

## 자기 점검 메모 (작성자용)
- **패턴 일관**: 신규 `statements`, `ratios`는 기존 `company` 패키지와 동일 구조(client.go + 도메인 파일 + 테스트 + testdata).
- **공통 Params**: 각 패키지에 독립 `Params{Symbol, Period, Limit}` — 패키지 간 import 없이 cohesion 유지.
- **시계열 반환**: Profile은 첫 요소(`*Profile`)였으나 statements/ratios는 본질적으로 다기간 → `[]T` 그대로 반환.
- **빈 결과 → `httpclient.ErrNotFound`**: Profile과 동일 sentinel 재사용.
- **빈 symbol 가드**: 모든 메서드에 `strings.TrimSpace(symbol) == "" → error`(Profile과 동일).
- **fixture 출처**: 1순위 `FMP_API_KEY` 라이브 응답, 2순위 카탈로그 md의 응답 예시. 어느 쪽이든 유효한 JSON 배열이어야.
- **faithful 매핑**: 카탈로그 응답 키 전부를 struct에 반영. 누락 발견 시 add. moneyflow가 사용할 핵심 필드(IncomeStatement.{Revenue, GrossProfit, OperatingIncome, NetIncome, EPS, WeightedAverageShsOut}, BalanceSheet.{TotalAssets, TotalLiabilities, TotalStockholdersEquity}, Ratio.{BookValuePerShare, DebtToEquityRatio})는 테스트가 명시적으로 검증.
- **moneyflow ROE는 SDK 미제공**: `metrics-ratios` 엔드포인트에 ROE 없음 — moneyflow 어댑터에서 `netIncome / totalStockholdersEquity`로 직접 계산(spec 변경 사항).
- **릴리스 분리**: 본 plan은 v0.2.0 준비까지. 태그 push + Release 생성은 사용자 트리거.
