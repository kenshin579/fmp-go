# FMP Go SDK — Market Performance 그룹 (v0.13.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/marketperf-group`
- 토픽: FMP `marketPerformance` 카테고리 11 endpoint. 캠페인 11번째 그룹.

## 배경 / 목적

시장 등락(gainers/losers/active) + 섹터/산업 성과·PE 스냅샷/시계열. moneyflow 시장 개요 화면에 활용.

## 결정 사항 (브레인스토밍)

- **범위**: marketPerformance 11 endpoint. 신규 `marketperf/` 패키지, internal/fetch.
- **5 구조체**: `MarketMover`(gainers/losers/most-actives 공유) / `SectorPerformance`(snapshot+historical 공유) / `IndustryPerformance` / `SectorPE` / `IndustryPE`. snapshot 과 historical 은 응답 shape 동일(쿼리만 다름) → 구조체 공유.
- **sector/industry 분리**: JSON 키가 `sector` vs `industry` 로 달라 별도 구조체(필드 충실).
- **쿼리 3종**: movers(파라미터 없음) / snapshot(date 필수 + exchange/dimension 선택) / historical(dimension 필수 + from/to/exchange 선택).
- **JSON 키 주의**: mover 변동률 키는 `changesPercentage`(s 포함). most-active path 는 `most-actives`(복수).
- **all list**, fetch.List. snapshot 은 date 빈값 가드, historical 은 dimension 빈값 가드, movers 가드 없음.
- **릴리스**: `v0.13.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | 반환 |
|---|---|---|---|
| `marketperf/movers.go` | `BiggestGainers(ctx)` | `/stable/biggest-gainers` | `[]MarketMover` |
| | `BiggestLosers(ctx)` | `/stable/biggest-losers` | `[]MarketMover` |
| | `MostActives(ctx)` | `/stable/most-actives` | `[]MarketMover` |
| `marketperf/performance.go` | `SectorPerformanceSnapshot(ctx, date, exchange, sector)` | `/stable/sector-performance-snapshot` | `[]SectorPerformance` |
| | `IndustryPerformanceSnapshot(ctx, date, exchange, industry)` | `/stable/industry-performance-snapshot` | `[]IndustryPerformance` |
| | `HistoricalSectorPerformance(ctx, sector, from, to, exchange)` | `/stable/historical-sector-performance` | `[]SectorPerformance` |
| | `HistoricalIndustryPerformance(ctx, industry, from, to, exchange)` | `/stable/historical-industry-performance` | `[]IndustryPerformance` |
| `marketperf/pe.go` | `SectorPESnapshot(ctx, date, exchange, sector)` | `/stable/sector-pe-snapshot` | `[]SectorPE` |
| | `IndustryPESnapshot(ctx, date, exchange, industry)` | `/stable/industry-pe-snapshot` | `[]IndustryPE` |
| | `HistoricalSectorPE(ctx, sector, from, to, exchange)` | `/stable/historical-sector-pe` | `[]SectorPE` |
| | `HistoricalIndustryPE(ctx, industry, from, to, exchange)` | `/stable/historical-industry-pe` | `[]IndustryPE` |
| `marketperf/client.go` | `New(http)` + snapshotParams/historicalParams | — | — |

helpers (client.go):
```go
// snapshotParams: date(필수) + exchange + dimension(sector|industry, 비어있지 않으면).
func snapshotParams(date, exchange, dimKey, dimVal string) map[string]string {
	q := map[string]string{"date": date}
	if exchange != "" { q["exchange"] = exchange }
	if dimVal != "" { q[dimKey] = dimVal }
	return q
}
// historicalParams: dimension(필수) + from + to + exchange(비어있지 않으면).
func historicalParams(dimKey, dimVal, from, to, exchange string) map[string]string {
	q := map[string]string{dimKey: dimVal}
	if from != "" { q["from"] = from }
	if to != "" { q["to"] = to }
	if exchange != "" { q["exchange"] = exchange }
	return q
}
```
snapshot 메서드는 date 빈값 가드, historical 메서드는 dimVal 빈값 가드(`fmp: ... must not be empty`).

## 루트 Client 와이어
```go
type Client struct {
	...
	Chart            *chart.Client
	MarketPerformance *marketperf.Client // 시장 성과(등락/섹터/산업/PE)
}
```
`NewClient` 에 `c.MarketPerformance = marketperf.New(hc)`. `client_test.go` 에 `TestNewClient_HasMarketPerformance`.

## 응답 타입 (faithful, 필드 한국어 주석)

```go
// MarketMover — 등락 상위 종목 (biggest-gainers/losers/most-actives 공유)
type MarketMover struct {
	Symbol            string  `json:"symbol"`            // 종목 심볼
	Price             float64 `json:"price"`             // 현재가
	Name              string  `json:"name"`              // 종목명
	Change            float64 `json:"change"`            // 변동액
	ChangesPercentage float64 `json:"changesPercentage"` // 변동률(%) — FMP 키 's' 포함
	Exchange          string  `json:"exchange"`          // 거래소
}

// SectorPerformance — 섹터 성과 (sector-performance-snapshot / historical-sector-performance 공유)
type SectorPerformance struct {
	Date          string  `json:"date"`          // 일자
	Sector        string  `json:"sector"`        // 섹터명
	Exchange      string  `json:"exchange"`      // 거래소
	AverageChange float64 `json:"averageChange"` // 평균 변동률
}

// IndustryPerformance — 산업 성과 (industry-performance-snapshot / historical-industry-performance 공유)
type IndustryPerformance struct {
	Date          string  `json:"date"`          // 일자
	Industry      string  `json:"industry"`      // 산업명
	Exchange      string  `json:"exchange"`      // 거래소
	AverageChange float64 `json:"averageChange"` // 평균 변동률
}

// SectorPE — 섹터 PER (sector-pe-snapshot / historical-sector-pe 공유)
type SectorPE struct {
	Date     string  `json:"date"`     // 일자
	Sector   string  `json:"sector"`   // 섹터명
	Exchange string  `json:"exchange"` // 거래소
	PE       float64 `json:"pe"`       // 섹터 평균 PER
}

// IndustryPE — 산업 PER (industry-pe-snapshot / historical-industry-pe 공유)
type IndustryPE struct {
	Date     string  `json:"date"`     // 일자
	Industry string  `json:"industry"` // 산업명
	Exchange string  `json:"exchange"` // 거래소
	PE       float64 `json:"pe"`       // 산업 평균 PER
}
```

## 시그니처 규칙
- movers 3개: `(ctx)` → `fetch.List[MarketMover](ctx, c.http, path, nil)`. (params nil 허용 — httpclient 가 빈 쿼리 처리. 만약 nil 미허용이면 `map[string]string{}`.)
- snapshot 4개: `(ctx, date, exchange, dim string)` → date 빈값 가드 + `fetch.List[T](..., snapshotParams(date, exchange, "sector"|"industry", dim))`.
- historical 4개: `(ctx, dim, from, to, exchange string)` → dim 빈값 가드 + `fetch.List[T](..., historicalParams("sector"|"industry", dim, from, to, exchange))`.

## 테스트
- fixture 단위: MarketMover(ChangesPercentage 키 파싱), SectorPerformance/IndustryPerformance(AverageChange), SectorPE/IndustryPE(PE).
- delegation: BiggestGainers path / SectorPerformanceSnapshot(date,exchange,sector) path+date/sector / HistoricalSectorPE(sector,from,to,"") path+sector/from/to.
- 가드: snapshot 빈 date 1건, historical 빈 dimension 1건.
- 통합(`//go:build integration`): BiggestGainers len>0 / SectorPerformanceSnapshot(오늘 날짜,"","") len>0 / SectorPESnapshot(날짜,"","") .
  - 통합 날짜는 최근 영업일 문자열 하드코딩 대신 time.Now().Format("2006-01-02") 사용.

## 문서 / 릴리스
- README 커버리지 표 Market Performance 행 신규(11 endpoint).
- `examples/marketperf/main.go` — BiggestGainers + SectorPerformanceSnapshot.
- 릴리스 `v0.13.0`.

## 범위 밖 / 위험
- movers 무파라미터 — nil params 가 httpclient.GetJSON 에서 동작하는지 확인(미동작 시 빈 맵).
- snapshot date 형식 YYYY-MM-DD 는 호출자 책임. 주말/휴일 빈 결과 가능.
- 다음 그룹: directory 또는 marketHours 등.
