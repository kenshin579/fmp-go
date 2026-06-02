# Chart 그룹 (v0.12.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `chart/` 패키지 10 endpoint(EOD 4 + intraday 6), 4 구조체.

**Architecture:** internal/fetch(List) + 패키지 helper eodParams/intradayParams. EODLight/EODFull/EODAdjusted/IntradayBar. 구조체는 스펙 verbatim.

**Tech Stack:** Go 1.25 generics, internal/fetch.

참고: `unset GOROOT`. 커밋 한국어 `feat(chart): ...`. 테스트 헬퍼는 calendar 패턴을 패키지 내부 정의.

---

### Task 1: chart 패키지 scaffold + EOD 4종

**Files:** Create `chart/client.go`, `chart/eod.go`, `chart/eod_test.go`, testdata `eod-light.json`, `eod-full.json`, `eod-dividend-adjusted.json`, `eod-non-split-adjusted.json`.

- [ ] **Step 1:** `chart/client.go`:
```go
// Package chart 는 FMP 과거 시세 API sub-client (EOD/intraday).
// fmp.Client.Chart 로 접근.
package chart

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 과거 시세 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// eodParams 는 symbol(필수) + from/to(비어있지 않으면) 쿼리 맵.
func eodParams(symbol, from, to string) map[string]string {
	q := map[string]string{"symbol": symbol}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return q
}

// intradayParams 는 eodParams + nonadjusted(true 일 때만).
func intradayParams(symbol, from, to string, nonadjusted bool) map[string]string {
	q := eodParams(symbol, from, to)
	if nonadjusted {
		q["nonadjusted"] = "true"
	}
	return q
}
```
- [ ] **Step 2:** `chart/eod.go` — 스펙 `EODLight`/`EODFull`/`EODAdjusted` struct + 4 메서드(빈 symbol 가드 + fetch.List + eodParams). import: context, fmt, strings, internal/fetch.
```go
func (c *Client) HistoricalPriceEODLight(ctx context.Context, symbol, from, to string) ([]EODLight, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EODLight](ctx, c.http, "/stable/historical-price-eod/light", eodParams(symbol, from, to))
}
// Full → /stable/historical-price-eod/full ([]EODFull)
// DividendAdjusted → /stable/historical-price-eod/dividend-adjusted ([]EODAdjusted)
// NonSplitAdjusted → /stable/historical-price-eod/non-split-adjusted ([]EODAdjusted)
```
- [ ] **Step 3:** fixtures(1~2건 배열, AAPL). light{symbol,date,price,volume}, full{...OHLCV,change,changePercent,vwap}, dividend-adjusted/non-split-adjusted{symbol,date,adjOpen,adjHigh,adjLow,adjClose,volume}.
- [ ] **Step 4:** `chart/eod_test.go` — 헬퍼(newTestClient/newCapturingClient/capturedReq) calendar 패턴 정의. 테스트: EODLight 파싱(Price!=0, Volume!=0) / EODFull 파싱(Close!=0, VWAP!=0) + delegation(path+from/to) / EODAdjusted 파싱(AdjClose!=0) 두 메서드 모두 / 빈 symbol 가드 1건.
- [ ] **Step 5:** `unset GOROOT && go test ./chart/ && go vet ./chart/ && gofmt -l chart/`. 커밋 `feat(chart): 패키지 기반 + EOD 4종`.

### Task 2: intraday 6종

**Files:** Create `chart/intraday.go`, `chart/intraday_test.go`, testdata `intraday-1min.json` (1개 fixture 공유).

- [ ] **Step 1:** `chart/intraday.go` — 스펙 `IntradayBar` struct + 6 메서드(공유 IntradayBar, path 만 다름):
```go
func (c *Client) Intraday1Min(ctx context.Context, symbol, from, to string, nonadjusted bool) ([]IntradayBar, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[IntradayBar](ctx, c.http, "/stable/historical-chart/1min", intradayParams(symbol, from, to, nonadjusted))
}
// Intraday5Min → 5min, Intraday15Min → 15min, Intraday30Min → 30min, Intraday1Hour → 1hour, Intraday4Hour → 4hour (동일 패턴)
```
- [ ] **Step 2:** fixture `chart/testdata/intraday-1min.json` — 2건 {date,open,low,high,close,volume} (symbol 키 없음).
- [ ] **Step 3:** `chart/intraday_test.go`(헬퍼 재사용): Intraday1Min 파싱(Close!=0, Volume!=0) / delegation Intraday1Min(...,true) path `/stable/historical-chart/1min` + nonadjusted=="true" / Intraday5Min(...,false) path `/stable/historical-chart/5min` + nonadjusted 키 부재 / 빈 symbol 가드 1건. (대표로 1min/5min 검증, 나머지 4개는 path 만 다른 동일 코드.)
- [ ] **Step 4:** `go test ./chart/ && go vet && gofmt -l`. 커밋 `feat(chart): intraday 6종(1min~4hour)`.

### Task 3: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/chart/main.go`.

- [ ] **Step 1:** `client.go` — import `chart`, struct 에 `Chart *chart.Client`, NewClient 에 `c.Chart = chart.New(hc)`.
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasChart`.
- [ ] **Step 3:** README 표 Chart 행 신규: `| Chart | \`client.Chart\` | HistoricalPriceEODLight, HistoricalPriceEODFull, HistoricalPriceEODDividendAdjusted, HistoricalPriceEODNonSplitAdjusted, Intraday1Min, Intraday5Min, Intraday15Min, Intraday30Min, Intraday1Hour, Intraday4Hour — 10 endpoint |`.
- [ ] **Step 4:** `examples/chart/main.go` — NewClientFromEnv → HistoricalPriceEODLight(AAPL,"","") 첫 행 출력 + Intraday1Hour(AAPL,"","",false) 건수 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Chart`: HistoricalPriceEODLight(AAPL,"","") len>0 / HistoricalPriceEODFull(AAPL,"","") rows[0].Close>0 / Intraday1Hour(AAPL,"","",false) len>0.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(chart): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 10 endpoint: EOD 4=T1, intraday 6=T2, 와이어/문서=T3.
- EODAdjusted 는 dividend/non-split 공유. IntradayBar 는 6 interval 공유, symbol 필드 없음.
- intraday nonadjusted=true 일 때만 쿼리 포함 — 테스트로 부재/존재 검증.
