# Statements-Core 확장 (v0.9.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development. Steps use checkbox (`- [ ]`) syntax.

**Goal:** 기존 `statements/` 패키지에 현금흐름표 + TTM 3종 + 성장률 4종(총 8 endpoint)을 추가한다.

**Architecture:** 기존 `Params{Symbol,Period,Limit}` 재사용 + `ttmQueryParams()` 추가. statements 패키지 private generic helper `fetchList[T]`(symbol 가드 + GetJSON + 빈 결과 ErrNotFound)로 income/balance/cashflow/ttm/growth 전부 통일. TTM 은 코어 구조체(IncomeStatement/BalanceSheetStatement/CashFlowStatement) 재사용. 구조체/JSON 태그는 `docs/superpowers/specs/2026-06-02-statements-core-expansion-design.md` 의 verbatim 정의를 따른다.

**Tech Stack:** Go 1.25 generics, `internal/httpclient.GetJSON`, 기존 statements 테스트 패턴.

참고: `unset GOROOT` 후 go 명령 실행. 커밋 메시지 한국어 `feat(statements): ...`.

---

### Task 1: helper + Params.ttmQueryParams + income/balance 리팩터 + CashFlowStatement

**Files:**
- Modify: `statements/client.go` (ttmQueryParams + fetchList helper 추가; import fmt/strings/context 정리)
- Modify: `statements/income.go` (fetchList 사용으로 단순화)
- Modify: `statements/balance.go` (fetchList 사용으로 단순화)
- Create: `statements/cashflow.go` (CashFlowStatement struct + 메서드)
- Create: `statements/cashflow_test.go`
- Create: `statements/testdata/cash-flow-statement.json`

- [ ] **Step 1:** `statements/client.go` 에 `ttmQueryParams()` 와 generic `fetchList[T]` 추가 (스펙 "내부 helper + Params 확장" 절 코드 그대로). 필요한 import: `context`, `fmt`, `strconv`, `strings`, `internal/httpclient`.
- [ ] **Step 2:** `income.go` / `balance.go` 의 메서드 본문을 `return fetchList[T](ctx, c, path, p, p.queryParams())` 형태로 교체. struct 정의는 그대로 유지. 미사용 import(`fmt`,`strings`,`httpclient`) 제거.
- [ ] **Step 3:** `cashflow.go` 작성 — 스펙의 `CashFlowStatement` struct 그대로 + 메서드:
```go
func (c *Client) CashFlowStatement(ctx context.Context, p Params) ([]CashFlowStatement, error) {
	return fetchList[CashFlowStatement](ctx, c, "/stable/cash-flow-statement", p, p.queryParams())
}
```
- [ ] **Step 4:** fixture `testdata/cash-flow-statement.json` — 스펙 cashflow 필드를 가진 1건 배열(현실적 AAPL 값, freeCashFlow 0 아님). 단위테스트: 파싱(freeCashFlow/netIncome 검증) + delegation(path `/stable/cash-flow-statement`, query symbol/period) + 빈 symbol 가드 + 빈 배열 ErrNotFound.
- [ ] **Step 5:** `unset GOROOT && go test ./statements/ && go vet ./statements/ && gofmt -l statements/` 통과(기존 income/balance 테스트 회귀 없음). 커밋 `feat(statements): CashFlowStatement + fetchList helper + ttmQueryParams`.

### Task 2: TTM 3종 (코어 구조체 재사용)

**Files:**
- Create: `statements/ttm.go`
- Create: `statements/ttm_test.go`
- Create: `statements/testdata/income-statement-ttm.json`, `balance-sheet-statement-ttm.json`, `cash-flow-statement-ttm.json`

- [ ] **Step 1:** `ttm.go` — 3 메서드, 코어 구조체 반환, `p.ttmQueryParams()` 사용:
```go
func (c *Client) IncomeStatementTTM(ctx context.Context, p Params) ([]IncomeStatement, error) {
	return fetchList[IncomeStatement](ctx, c, "/stable/income-statement-ttm", p, p.ttmQueryParams())
}
func (c *Client) BalanceSheetStatementTTM(ctx context.Context, p Params) ([]BalanceSheetStatement, error) {
	return fetchList[BalanceSheetStatement](ctx, c, "/stable/balance-sheet-statement-ttm", p, p.ttmQueryParams())
}
func (c *Client) CashFlowStatementTTM(ctx context.Context, p Params) ([]CashFlowStatement, error) {
	return fetchList[CashFlowStatement](ctx, c, "/stable/cash-flow-statement-ttm", p, p.ttmQueryParams())
}
```
- [ ] **Step 2:** fixtures — 각 endpoint 1건(코어 구조체와 동일 필드, period 본문엔 존재). 단위테스트: 각 TTM 파싱 1건 + capturing 으로 path `*-ttm` 확인 + 쿼리에 `period` 키 **없음** 확인(Period 지정해도 생략). 빈 symbol 가드 1건.
- [ ] **Step 3:** `go test ./statements/ && go vet && gofmt -l` 통과. 커밋 `feat(statements): IncomeStatementTTM + BalanceSheetStatementTTM + CashFlowStatementTTM`.

### Task 3: growth 4종

**Files:**
- Create: `statements/growth.go` (4 struct + 4 메서드, 스펙 정의 그대로)
- Create: `statements/growth_test.go`
- Create: `statements/testdata/income-statement-growth.json`, `balance-sheet-statement-growth.json`, `cash-flow-statement-growth.json`, `financial-growth.json`

- [ ] **Step 1:** `growth.go` — 스펙의 `IncomeStatementGrowth`/`BalanceSheetStatementGrowth`/`CashFlowStatementGrowth`/`FinancialStatementGrowth` struct 그대로(JSON 태그 오타 보존, FinancialStatementGrowth 마지막 5필드 `*float64`) + 4 메서드:
```go
func (c *Client) IncomeStatementGrowth(ctx context.Context, p Params) ([]IncomeStatementGrowth, error) {
	return fetchList[IncomeStatementGrowth](ctx, c, "/stable/income-statement-growth", p, p.queryParams())
}
func (c *Client) BalanceSheetStatementGrowth(ctx context.Context, p Params) ([]BalanceSheetStatementGrowth, error) {
	return fetchList[BalanceSheetStatementGrowth](ctx, c, "/stable/balance-sheet-statement-growth", p, p.queryParams())
}
func (c *Client) CashFlowStatementGrowth(ctx context.Context, p Params) ([]CashFlowStatementGrowth, error) {
	return fetchList[CashFlowStatementGrowth](ctx, c, "/stable/cash-flow-statement-growth", p, p.queryParams())
}
func (c *Client) FinancialStatementGrowth(ctx context.Context, p Params) ([]FinancialStatementGrowth, error) {
	return fetchList[FinancialStatementGrowth](ctx, c, "/stable/financial-growth", p, p.queryParams())
}
```
- [ ] **Step 2:** fixtures — 각 1건. `financial-growth.json` 은 2건: (a) 마지막 5필드 값 존재, (b) 마지막 5필드 `null`. 단위테스트: 각 growth 핵심 필드 파싱 + FinancialStatementGrowth null→nil 및 값→non-nil 검증 + delegation(`financial-growth` path, period 쿼리).
- [ ] **Step 3:** `go test ./statements/ && go vet && gofmt -l` 통과. 커밋 `feat(statements): Income/Balance/CashFlow/Financial 성장률 4종`.

### Task 4: README + 통합테스트

**Files:**
- Modify: `README.md` (Statements 행)
- Modify: `integration_test.go` (TestIntegration_StatementsCore 추가)

- [ ] **Step 1:** README 커버리지 표 Statements 행을 다음으로 교체:
`| Statements | \`client.Statements\` | IncomeStatement, BalanceSheetStatement, CashFlowStatement, IncomeStatementTTM, BalanceSheetStatementTTM, CashFlowStatementTTM, IncomeStatementGrowth, BalanceSheetStatementGrowth, CashFlowStatementGrowth, FinancialStatementGrowth — 10 endpoint |`
- [ ] **Step 2:** `integration_test.go` 에 `TestIntegration_StatementsCore` 추가 — FMP_API_KEY 가드, AAPL 로: CashFlowStatement(annual, Limit 2) freeCashFlow!=0, IncomeStatementTTM, CashFlowStatementGrowth(annual), FinancialStatementGrowth(annual) — 마지막에 `t.Logf("FinancialStatementGrowth[0]: %+v")` 로 nullable 실제값 로그. import `statements` 패키지의 `Params` 사용.
- [ ] **Step 3:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 전부 통과. 커밋 `feat(statements): README 갱신 + statements-core 통합테스트`.

## Self-Review 메모
- 스펙 8 endpoint 모두 Task 1~3 에 매핑됨(cashflow=T1, ttm3=T2, growth4=T3). README/통합=T4.
- helper 리팩터로 income/balance 회귀 위험 → 각 Task 에서 기존 테스트 통과 확인.
- FinancialStatementGrowth nullable 5필드만 `*float64`, 나머지 growth 전부 `float64`.
- TTM 쿼리 period 생략은 `ttmQueryParams()` + 테스트로 보장.
