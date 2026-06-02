# Market Performance 그룹 (v0.13.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `marketperf/` 패키지 11 endpoint, 5 구조체.

**Architecture:** internal/fetch(List) + helper snapshotParams/historicalParams. MarketMover/SectorPerformance/IndustryPerformance/SectorPE/IndustryPE. snapshot·historical 은 구조체 공유, 쿼리만 다름. 구조체는 스펙 verbatim.

**Tech Stack:** Go 1.25 generics, internal/fetch.

참고: `unset GOROOT`. 커밋 한국어 `feat(marketperf): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의. nil params 는 GetJSON 의 `range` 에서 안전(0회).

---

### Task 1: 패키지 scaffold + movers 3종

**Files:** Create `marketperf/client.go`, `marketperf/movers.go`, `marketperf/movers_test.go`, testdata `biggest-gainers.json`.

- [ ] **Step 1:** `marketperf/client.go`:
```go
// Package marketperf 는 FMP 시장 성과 API sub-client (등락/섹터/산업/PE).
// fmp.Client.MarketPerformance 로 접근.
package marketperf

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 시장 성과 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

// snapshotParams: date(필수) + exchange + dimension(sector|industry, 비어있지 않으면).
func snapshotParams(date, exchange, dimKey, dimVal string) map[string]string {
	q := map[string]string{"date": date}
	if exchange != "" {
		q["exchange"] = exchange
	}
	if dimVal != "" {
		q[dimKey] = dimVal
	}
	return q
}

// historicalParams: dimension(필수) + from + to + exchange(비어있지 않으면).
func historicalParams(dimKey, dimVal, from, to, exchange string) map[string]string {
	q := map[string]string{dimKey: dimVal}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	if exchange != "" {
		q["exchange"] = exchange
	}
	return q
}
```
(snapshotParams/historicalParams 는 Task 2/3 에서 사용 — 패키지 레벨 미사용은 컴파일 에러 아님.)
- [ ] **Step 2:** `marketperf/movers.go`:
```go
package marketperf

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

// MarketMover — 등락 상위 종목 (biggest-gainers/losers/most-actives 공유)
type MarketMover struct {
	Symbol            string  `json:"symbol"`            // 종목 심볼
	Price             float64 `json:"price"`             // 현재가
	Name              string  `json:"name"`              // 종목명
	Change            float64 `json:"change"`            // 변동액
	ChangesPercentage float64 `json:"changesPercentage"` // 변동률(%) — FMP 키 's' 포함
	Exchange          string  `json:"exchange"`          // 거래소
}

// BiggestGainers 는 상승률 상위 종목을 조회한다.
func (c *Client) BiggestGainers(ctx context.Context) ([]MarketMover, error) {
	return fetch.List[MarketMover](ctx, c.http, "/stable/biggest-gainers", nil)
}

// BiggestLosers 는 하락률 상위 종목을 조회한다.
func (c *Client) BiggestLosers(ctx context.Context) ([]MarketMover, error) {
	return fetch.List[MarketMover](ctx, c.http, "/stable/biggest-losers", nil)
}

// MostActives 는 거래 활발 종목을 조회한다.
func (c *Client) MostActives(ctx context.Context) ([]MarketMover, error) {
	return fetch.List[MarketMover](ctx, c.http, "/stable/most-actives", nil)
}
```
- [ ] **Step 3:** fixture `biggest-gainers.json` — 1~2건 `{symbol, price, name, change, changesPercentage, exchange}`.
- [ ] **Step 4:** `movers_test.go`(헬퍼 calendar 패턴 정의): BiggestGainers 파싱(Symbol!="", ChangesPercentage!=0, Price!=0) + delegation(path `/stable/biggest-gainers`). BiggestLosers/MostActives 는 path 만 다른 동일 코드 — delegation path 각각 1줄 검증(capturing) 권장.
- [ ] **Step 5:** `unset GOROOT && go test ./marketperf/ && go vet ./marketperf/ && gofmt -l marketperf/`. 커밋 `feat(marketperf): 패키지 기반 + 등락 상위 3종`.

### Task 2: 섹터/산업 성과 4종

**Files:** Create `marketperf/performance.go`, `marketperf/performance_test.go`, testdata `sector-performance-snapshot.json`, `industry-performance-snapshot.json`.

- [ ] **Step 1:** `performance.go` — 스펙 `SectorPerformance` + `IndustryPerformance` struct + 4 메서드:
```go
// SectorPerformanceSnapshot(ctx, date, exchange, sector) → date 가드 + snapshotParams(date, exchange, "sector", sector), path /stable/sector-performance-snapshot, []SectorPerformance
// IndustryPerformanceSnapshot(ctx, date, exchange, industry) → date 가드 + snapshotParams(..., "industry", industry), path /stable/industry-performance-snapshot, []IndustryPerformance
// HistoricalSectorPerformance(ctx, sector, from, to, exchange) → sector 가드 + historicalParams("sector", sector, from, to, exchange), path /stable/historical-sector-performance, []SectorPerformance
// HistoricalIndustryPerformance(ctx, industry, from, to, exchange) → industry 가드 + historicalParams("industry", industry, from, to, exchange), path /stable/historical-industry-performance, []IndustryPerformance
```
date/dimension 빈값 가드는 `if strings.TrimSpace(x) == "" { return nil, fmt.Errorf("fmp: ... must not be empty") }`. import: context, fmt, strings, internal/fetch.
- [ ] **Step 2:** fixtures: sector-performance-snapshot `[{date,sector,exchange,averageChange}]`, industry-performance-snapshot `[{date,industry,exchange,averageChange}]`.
- [ ] **Step 3:** `performance_test.go`(헬퍼 재사용): SectorPerformanceSnapshot 파싱(Sector!="", AverageChange 파싱) + delegation(path+date+sector) / IndustryPerformanceSnapshot 파싱 / HistoricalSectorPerformance delegation(path `/stable/historical-sector-performance`+sector+from+to) / 빈 date 가드 1건 + 빈 sector(historical) 가드 1건.
- [ ] **Step 4:** `go test ./marketperf/ && go vet && gofmt -l`. 커밋 `feat(marketperf): 섹터/산업 성과 4종`.

### Task 3: 섹터/산업 PER 4종

**Files:** Create `marketperf/pe.go`, `marketperf/pe_test.go`, testdata `sector-pe-snapshot.json`, `industry-pe-snapshot.json`.

- [ ] **Step 1:** `pe.go` — 스펙 `SectorPE` + `IndustryPE` struct + 4 메서드:
```go
// SectorPESnapshot(ctx, date, exchange, sector) → date 가드 + snapshotParams(date, exchange, "sector", sector), path /stable/sector-pe-snapshot, []SectorPE
// IndustryPESnapshot(ctx, date, exchange, industry) → date 가드, path /stable/industry-pe-snapshot, []IndustryPE
// HistoricalSectorPE(ctx, sector, from, to, exchange) → sector 가드 + historicalParams("sector", ...), path /stable/historical-sector-pe, []SectorPE
// HistoricalIndustryPE(ctx, industry, from, to, exchange) → industry 가드, path /stable/historical-industry-pe, []IndustryPE
```
- [ ] **Step 2:** fixtures: sector-pe-snapshot `[{date,sector,exchange,pe}]`, industry-pe-snapshot `[{date,industry,exchange,pe}]`.
- [ ] **Step 3:** `pe_test.go`(헬퍼 재사용): SectorPESnapshot 파싱(PE!=0) + delegation(path+date+sector) / IndustryPESnapshot 파싱 / HistoricalSectorPE delegation(path `/stable/historical-sector-pe`) / 빈 date 가드 1건.
- [ ] **Step 4:** `go test ./marketperf/ && go vet && gofmt -l`. 커밋 `feat(marketperf): 섹터/산업 PER 4종`.

### Task 4: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/marketperf/main.go`.

- [ ] **Step 1:** `client.go` — import `marketperf`, struct 에 `MarketPerformance *marketperf.Client`, NewClient 에 `c.MarketPerformance = marketperf.New(hc)`.
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasMarketPerformance`.
- [ ] **Step 3:** README 표 행 신규: `| Market Performance | \`client.MarketPerformance\` | BiggestGainers, BiggestLosers, MostActives, SectorPerformanceSnapshot, IndustryPerformanceSnapshot, HistoricalSectorPerformance, HistoricalIndustryPerformance, SectorPESnapshot, IndustryPESnapshot, HistoricalSectorPE, HistoricalIndustryPE — 11 endpoint |`.
- [ ] **Step 4:** `examples/marketperf/main.go` — NewClientFromEnv → BiggestGainers 상위 3개 출력 + SectorPerformanceSnapshot(time.Now().Format("2006-01-02"),"","") 건수 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_MarketPerformance`: BiggestGainers len>0 / SectorPerformanceSnapshot(time.Now().Format("2006-01-02"),"","") (주말 빈 결과 가능하니 err 만 체크) / SectorPESnapshot(동일 날짜,"",""). import time 필요.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(marketperf): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 11 endpoint: movers 3=T1, performance 4=T2, PE 4=T3, 와이어/문서=T4.
- snapshot/historical 구조체 공유(쿼리만 다름). sector/industry 별도 구조체.
- ChangesPercentage 키 's' 포함. most-actives 복수형 path.
- snapshot date 가드, historical dimension 가드, movers 가드 없음.
