# FMP Go SDK — Chart 그룹 (v0.12.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/chart-group`
- 토픽: FMP `chart`(과거 시세) 카테고리 10 endpoint. 캠페인 10번째 그룹. statements 완료 후 첫 신규 카테고리.

## 배경 / 목적

과거 EOD/intraday 시세 — moneyflow 차트/백테스트에 직접 활용. 외부 개발자에게도 핵심.

## 결정 사항 (브레인스토밍)

- **범위**: chart 10 endpoint(EOD 4 + intraday 6). 신규 `chart/` 패키지, internal/fetch.
- **4 구조체**: `EODLight`(symbol,date,price,volume) / `EODFull`(OHLCV+change/changePercent/vwap) / `EODAdjusted`(adjOHLC, dividend-adjusted·non-split-adjusted 공유) / `IntradayBar`(date+OHLCV, symbol 없음, 6 interval 공유).
- **시그니처**: EOD 4개 `(ctx, symbol, from, to string)`. intraday 6개 `(ctx, symbol, from, to string, nonadjusted bool)`.
- **쿼리**: symbol 필수, from/to 비어있지 않으면 포함. intraday 는 nonadjusted=true 일 때만 `nonadjusted=true` 포함.
- **all list**, 빈 symbol 가드. fetch.List(빈 결과 ErrNotFound 안 던짐).
- **릴리스**: `v0.12.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | 반환 |
|---|---|---|---|
| `chart/eod.go` | `HistoricalPriceEODLight(ctx, symbol, from, to)` | `/stable/historical-price-eod/light` | `[]EODLight` |
| | `HistoricalPriceEODFull(ctx, symbol, from, to)` | `/stable/historical-price-eod/full` | `[]EODFull` |
| | `HistoricalPriceEODDividendAdjusted(ctx, symbol, from, to)` | `/stable/historical-price-eod/dividend-adjusted` | `[]EODAdjusted` |
| | `HistoricalPriceEODNonSplitAdjusted(ctx, symbol, from, to)` | `/stable/historical-price-eod/non-split-adjusted` | `[]EODAdjusted` |
| `chart/intraday.go` | `Intraday1Min(ctx, symbol, from, to, nonadjusted)` | `/stable/historical-chart/1min` | `[]IntradayBar` |
| | `Intraday5Min(...)` | `/stable/historical-chart/5min` | `[]IntradayBar` |
| | `Intraday15Min(...)` | `/stable/historical-chart/15min` | `[]IntradayBar` |
| | `Intraday30Min(...)` | `/stable/historical-chart/30min` | `[]IntradayBar` |
| | `Intraday1Hour(...)` | `/stable/historical-chart/1hour` | `[]IntradayBar` |
| | `Intraday4Hour(...)` | `/stable/historical-chart/4hour` | `[]IntradayBar` |
| `chart/client.go` | `New(http)` + eodParams/intradayParams | — | — |

helpers:
- `eodParams(symbol, from, to string) map[string]string` — symbol + from/to(비어있지 않으면).
- `intradayParams(symbol, from, to string, nonadjusted bool)` — eodParams + nonadjusted=="true"(true 일 때만).
- EOD/intraday 메서드 모두 빈 symbol 가드.

## 루트 Client 와이어
```go
type Client struct {
	...
	Reports *reports.Client
	Chart   *chart.Client // 과거 시세(EOD/intraday)
}
```
`NewClient` 에 `c.Chart = chart.New(hc)`. `client_test.go` 에 `TestNewClient_HasChart`.

## 응답 타입 (faithful, 필드 한국어 주석)

```go
// EODLight — EOD 경량 시세 (historical-price-eod/light)
type EODLight struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	Date   string  `json:"date"`   // 일자
	Price  float64 `json:"price"`  // 종가
	Volume int64   `json:"volume"` // 거래량
}

// EODFull — EOD 전체 시세 (historical-price-eod/full)
type EODFull struct {
	Symbol        string  `json:"symbol"`        // 종목 심볼
	Date          string  `json:"date"`          // 일자
	Open          float64 `json:"open"`          // 시가
	High          float64 `json:"high"`          // 고가
	Low           float64 `json:"low"`           // 저가
	Close         float64 `json:"close"`         // 종가
	Volume        int64   `json:"volume"`        // 거래량
	Change        float64 `json:"change"`        // 변동액
	ChangePercent float64 `json:"changePercent"` // 변동률(%)
	VWAP          float64 `json:"vwap"`          // 거래량가중평균가
}

// EODAdjusted — 조정 시세 (dividend-adjusted / non-split-adjusted 공유)
type EODAdjusted struct {
	Symbol   string  `json:"symbol"`   // 종목 심볼
	Date     string  `json:"date"`     // 일자
	AdjOpen  float64 `json:"adjOpen"`  // 조정 시가
	AdjHigh  float64 `json:"adjHigh"`  // 조정 고가
	AdjLow   float64 `json:"adjLow"`   // 조정 저가
	AdjClose float64 `json:"adjClose"` // 조정 종가
	Volume   int64   `json:"volume"`   // 거래량
}

// IntradayBar — 분/시간봉 (historical-chart/{interval} 6종 공유). symbol 없음.
type IntradayBar struct {
	Date   string  `json:"date"`   // 일시
	Open   float64 `json:"open"`   // 시가
	Low    float64 `json:"low"`    // 저가
	High   float64 `json:"high"`   // 고가
	Close  float64 `json:"close"`  // 종가
	Volume int64   `json:"volume"` // 거래량
}
```

## 시그니처 규칙
- EOD 4개: `(ctx, symbol, from, to string)` → 빈 symbol 가드 + `fetch.List[T](ctx, c.http, path, eodParams(symbol, from, to))`.
- intraday 6개: `(ctx, symbol, from, to string, nonadjusted bool)` → 빈 symbol 가드 + `fetch.List[IntradayBar](ctx, c.http, path, intradayParams(symbol, from, to, nonadjusted))`.

## 테스트
- fixture 단위: EODLight/EODFull/EODAdjusted/IntradayBar 각 파싱(가격 float, volume int, IntradayBar symbol 없음 확인).
- delegation: HistoricalPriceEODFull(symbol,from,to) path+쿼리 from/to / Intraday1Min(symbol,from,to,true) path+nonadjusted=true / Intraday5Min(...,false) nonadjusted 키 부재.
- 가드: EOD/intraday 빈 symbol 대표 각 1건.
- 통합(`//go:build integration`): HistoricalPriceEODLight(AAPL,"","") len>0 / HistoricalPriceEODFull(AAPL,from,to) Close>0 / Intraday1Hour(AAPL,"","",false) len>0.

## 문서 / 릴리스
- README 커버리지 표 Chart 행 신규(10 endpoint).
- `examples/chart/main.go` — HistoricalPriceEODLight + Intraday1Hour.
- 릴리스 `v0.12.0`.

## 범위 밖 / 위험
- nonadjusted 는 intraday 전용 — EOD 엔 없음.
- IntradayBar JSON 키 순서(open,low,high) 무관(언마샬).
- 대용량 응답(수년치 일봉/분봉) — 페이지네이션 없음(FMP from/to 로 제한). 호출자 책임.
- 다음 그룹: marketPerformance 또는 directory 등(미정).
