# Technical Indicators 그룹 (v0.18.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `technicals/` 패키지 9 endpoint, 임베디드 Bar + 9 thin 구조체.

**Architecture:** internal/fetch(List) + validate/indicatorParams helper. Bar 임베디드(JSON 필드 승격). 9 메서드 동일 시그니처. 구조체는 스펙 verbatim.

참고: `unset GOROOT`. 커밋 한국어 `feat(technicals): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의. 메서드명==타입명 합법.

---

### Task 1: 패키지 scaffold + 이동평균 5종 (SMA/EMA/WMA/DEMA/TEMA)

**Files:** Create `technicals/client.go`, `technicals/movingaverage.go`, `technicals/movingaverage_test.go`, testdata `sma.json`, `ema.json`.

- [ ] **Step 1:** `technicals/client.go` — Client + New + validate + indicatorParams (스펙 helpers 절 그대로). import: fmt, strconv, strings, internal/httpclient.
- [ ] **Step 2:** `technicals/movingaverage.go` — `Bar`(임베디드) + `SMA`/`EMA`/`WMA`/`DEMA`/`TEMA` struct(각 Bar 임베드 + 지표 필드) + 5 메서드. import: context, internal/fetch.
```go
func (c *Client) SMA(ctx context.Context, symbol string, periodLength int, timeframe, from, to string) ([]SMA, error) {
	if err := validate(symbol, periodLength, timeframe); err != nil {
		return nil, err
	}
	return fetch.List[SMA](ctx, c.http, "/stable/technical-indicators/sma", indicatorParams(symbol, periodLength, timeframe, from, to))
}
// EMA → .../ema ([]EMA), WMA → .../wma, DEMA → .../dema, TEMA → .../tema (동일 패턴)
```
struct 예:
```go
type Bar struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}
type SMA struct {
	Bar
	SMA float64 `json:"sma"`
}
type EMA struct {
	Bar
	EMA float64 `json:"ema"`
}
type WMA struct {
	Bar
	WMA float64 `json:"wma"`
}
type DEMA struct {
	Bar
	DEMA float64 `json:"dema"`
}
type TEMA struct {
	Bar
	TEMA float64 `json:"tema"`
}
```
- [ ] **Step 3:** fixtures: sma.json `[{date:"2026-04-08 00:00:00",open,high,low,close:258.9,volume:39655304,sma:253.754}]`, ema.json(같은 OHLCV + ema).
- [ ] **Step 4:** `movingaverage_test.go`(헬퍼 calendar 패턴 정의): SMA 파싱(Close!=0(임베디드 Bar 승격 확인), SMA!=0) + delegation(SMA("AAPL",10,"1day","2026-01-01","2026-04-08") path `/stable/technical-indicators/sma` + symbol/periodLength=="10"/timeframe=="1day"/from/to) / EMA 파싱(EMA!=0) + path / 가드 3종(빈 symbol, periodLength 0, 빈 timeframe — SMA 대표).
- [ ] **Step 5:** `unset GOROOT && go test ./technicals/ && go vet ./technicals/ && gofmt -l technicals/`. 커밋 `feat(technicals): 이동평균 5종(SMA/EMA/WMA/DEMA/TEMA)`.

### Task 2: 오실레이터/기타 4종 (RSI/StandardDeviation/Williams/ADX)

**Files:** Create `technicals/oscillators.go`, `technicals/oscillators_test.go`, testdata `rsi.json`, `standarddeviation.json`, `adx.json`.

- [ ] **Step 1:** `technicals/oscillators.go` — `RSI`/`StandardDeviation`/`Williams`/`ADX` struct(Bar 임베드) + 4 메서드:
```go
type RSI struct {
	Bar
	RSI float64 `json:"rsi"`
}
type StandardDeviation struct {
	Bar
	StandardDeviation float64 `json:"standardDeviation"`
}
type Williams struct {
	Bar
	Williams float64 `json:"williams"`
}
type ADX struct {
	Bar
	ADX float64 `json:"adx"`
}
// RSI → /stable/technical-indicators/rsi
// StandardDeviation → /stable/technical-indicators/standarddeviation (path 소문자, JSON 키 standardDeviation)
// Williams → /stable/technical-indicators/williams
// ADX → /stable/technical-indicators/adx
```
각 메서드 validate + fetch.List + indicatorParams.
- [ ] **Step 2:** fixtures: rsi.json(OHLCV + rsi), standarddeviation.json(OHLCV + standardDeviation 키), adx.json(OHLCV + adx).
- [ ] **Step 3:** `oscillators_test.go`(헬퍼 재사용): RSI 파싱(RSI!=0) + path / StandardDeviation 파싱(StandardDeviation!=0 — JSON 키 standardDeviation 매핑 확인) + delegation(path `/stable/technical-indicators/standarddeviation`) / ADX 파싱(ADX!=0) + path / Williams path.
- [ ] **Step 4:** `go test ./technicals/ && go vet && gofmt -l`. 커밋 `feat(technicals): RSI/StandardDeviation/Williams/ADX 4종`.

### Task 3: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/technicals/main.go`.

- [ ] **Step 1:** `client.go` — import `technicals`, struct 에 `TechnicalIndicators *technicals.Client`, NewClient 에 `c.TechnicalIndicators = technicals.New(hc)`. (gofmt -w client.go.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasTechnicalIndicators`.
- [ ] **Step 3:** README 표 행 신규: `| Technical Indicators | \`client.TechnicalIndicators\` | SMA, EMA, WMA, DEMA, TEMA, RSI, StandardDeviation, Williams, ADX — 9 endpoint |`.
- [ ] **Step 4:** `examples/technicals/main.go` — NewClientFromEnv → SMA("AAPL",10,"1day","","") 첫 행 + RSI("AAPL",14,"1day","","") 첫 행 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_TechnicalIndicators`: SMA("AAPL",10,"1day","","") len>0 & rows[0].Close>0 & rows[0].SMA>0 / RSI("AAPL",14,"1day","","") len>0 / ADX("AAPL",14,"1day","","") len>0.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(technicals): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 9 endpoint: 이동평균 5=T1, 오실레이터 4=T2, 와이어/문서=T3.
- Bar 임베디드 — JSON 필드 승격으로 OHLCV+지표값 한 객체 파싱.
- standard-deviation path 소문자 / JSON 키 camelCase 주의.
- 메서드명==타입명(SMA 등) 합법.
