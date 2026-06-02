# Asset Lists 그룹 (v0.28.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `assets/` 패키지 3 endpoint(암호화폐/외환/원자재 목록).

**Architecture:** internal/fetch(List). 3 구조체(스펙 verbatim). 자산군 시세/시계열은 기존 quote·chart 가 커버하므로 중복 구현 안 함.

참고: `unset GOROOT`. 커밋 한국어 `feat(assets): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 + 3 list endpoint

**Files:** Create `assets/client.go`, `assets/assets.go`, `assets/assets_test.go`, testdata `crypto-list.json`, `forex-list.json`, `commodities-list.json`.

- [ ] **Step 1:** `assets/client.go`:
```go
// Package assets 는 FMP 암호화폐/외환/원자재 목록 API sub-client.
// fmp.Client.Assets 로 접근. (시세/시계열은 client.Quote, client.Chart 사용)
package assets

import "github.com/kenshin579/fmp-go/internal/httpclient"

type Client struct {
	http *httpclient.Client
}

func New(http *httpclient.Client) *Client { return &Client{http: http} }
```
- [ ] **Step 2:** `assets/assets.go` — 스펙 `CryptoListItem`/`ForexPair`/`CommodityListItem` struct + 3 메서드:
```go
func (c *Client) CryptoList(ctx context.Context) ([]CryptoListItem, error) {
	return fetch.List[CryptoListItem](ctx, c.http, "/stable/cryptocurrency-list", nil)
}
func (c *Client) ForexList(ctx context.Context) ([]ForexPair, error) {
	return fetch.List[ForexPair](ctx, c.http, "/stable/forex-list", nil)
}
func (c *Client) CommodityList(ctx context.Context) ([]CommodityListItem, error) {
	return fetch.List[CommodityListItem](ctx, c.http, "/stable/commodities-list", nil)
}
```
import: context, internal/fetch.
- [ ] **Step 3:** fixtures: crypto-list.json `[{symbol:"BTCUSD",name:"Bitcoin",exchange:"CRYPTO",icoDate:"2009-01-03",circulatingSupply:19600000,totalSupply:21000000}]`, forex-list.json `[{symbol:"EURUSD",fromCurrency:"EUR",toCurrency:"USD",fromName:"Euro",toName:"US Dollar"}]`, commodities-list.json `[{symbol:"GCUSD",name:"Gold",exchange:null,tradeMonth:"Dec",currency:"USD"}]`.
- [ ] **Step 4:** `assets_test.go`(헬퍼 calendar 패턴 정의): CryptoList 파싱(CirculatingSupply!=0, TotalSupply!=nil 값) + delegation(path) / ForexList 파싱(FromCurrency=="EUR") + path / CommodityList 파싱(Exchange==nil null, TradeMonth=="Dec") + path. (TotalSupply null 케이스도 한 행 추가해 nil 검증하면 더 좋음.)
- [ ] **Step 5:** `unset GOROOT && go test ./assets/ && go vet ./assets/ && gofmt -l assets/`. 커밋 `feat(assets): 암호화폐/외환/원자재 목록 3종`.

### Task 2: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/assets/main.go`.

- [ ] **Step 1:** `client.go` — import `assets`, struct 에 `Assets *assets.Client`, NewClient 에 `c.Assets = assets.New(hc)`. (gofmt -w client.go.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasAssets`.
- [ ] **Step 3:** README 표 행 신규: `| Assets | \`client.Assets\` | CryptoList, ForexList, CommodityList — 3 endpoint (시세/시계열은 client.Quote·client.Chart 사용) |`.
- [ ] **Step 4:** `examples/assets/main.go` — NewClientFromEnv → CryptoList 건수 + ForexList 건수 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_Assets`: CryptoList len>0 / ForexList len>0 / CommodityList len>0.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(assets): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 3 endpoint, 3 struct. 자산군 시세/시계열 중복 구현 안 함(기존 quote·chart).
- CryptoListItem.TotalSupply / CommodityListItem.Exchange nullable(*).
