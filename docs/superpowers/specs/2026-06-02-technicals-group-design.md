# FMP Go SDK — Technical Indicators 그룹 (v0.18.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/technicals-group`
- 토픽: FMP `technicalIndicators` 카테고리 9 endpoint. 캠페인 16번째 그룹.

## 결정 사항
- 신규 `technicals/` 패키지, internal/fetch.
- 9종 응답이 공통 OHLCV(date/open/high/low/close/volume) + 단일 지표값 키만 다름 → **임베디드 `Bar` + 9 thin 구조체**(각자 실제 JSON 키로 타입드 필드). 커스텀 언마샬 불필요(임베디드 필드 승격).
- 9 메서드 시그니처 동일: `(ctx, symbol string, periodLength int, timeframe, from, to string)`. symbol/periodLength/timeframe 필수 가드.
- 메서드명 == 타입명(SMA 등) — Go 합법.
- path: `/stable/technical-indicators/{sma,ema,wma,dema,tema,rsi,standarddeviation,williams,adx}`. standard-deviation 만 path 소문자 `standarddeviation`, JSON 키 camelCase `standardDeviation`.
- 릴리스 `v0.18.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | 지표 키 | 반환 |
|---|---|---|---|
| `SMA(ctx, symbol, periodLength, timeframe, from, to)` | `/stable/technical-indicators/sma` | sma | `[]SMA` |
| `EMA(...)` | `.../ema` | ema | `[]EMA` |
| `WMA(...)` | `.../wma` | wma | `[]WMA` |
| `DEMA(...)` | `.../dema` | dema | `[]DEMA` |
| `TEMA(...)` | `.../tema` | tema | `[]TEMA` |
| `RSI(...)` | `.../rsi` | rsi | `[]RSI` |
| `StandardDeviation(...)` | `.../standarddeviation` | standardDeviation | `[]StandardDeviation` |
| `Williams(...)` | `.../williams` | williams | `[]Williams` |
| `ADX(...)` | `.../adx` | adx | `[]ADX` |

파일: `technicals/client.go`(New + validate + indicatorParams), `technicals/movingaverage.go`(Bar + SMA/EMA/WMA/DEMA/TEMA), `technicals/oscillators.go`(RSI/StandardDeviation/Williams/ADX).

helpers (client.go):
```go
func validate(symbol string, periodLength int, timeframe string) error {
	if strings.TrimSpace(symbol) == "" {
		return fmt.Errorf("fmp: symbol must not be empty")
	}
	if periodLength <= 0 {
		return fmt.Errorf("fmp: periodLength must be > 0")
	}
	if strings.TrimSpace(timeframe) == "" {
		return fmt.Errorf("fmp: timeframe must not be empty")
	}
	return nil
}
func indicatorParams(symbol string, periodLength int, timeframe, from, to string) map[string]string {
	q := map[string]string{
		"symbol":       symbol,
		"periodLength": strconv.Itoa(periodLength),
		"timeframe":    timeframe,
	}
	if from != "" {
		q["from"] = from
	}
	if to != "" {
		q["to"] = to
	}
	return q
}
```

## 루트 Client 와이어
```go
InsiderTrades      *insidertrades.Client
TechnicalIndicators *technicals.Client // 기술 지표(SMA/EMA/RSI/ADX 등)
```
`c.TechnicalIndicators = technicals.New(hc)`. `TestNewClient_HasTechnicalIndicators`.

## 응답 타입 (faithful)
```go
// Bar — 기술지표 응답 공통 OHLCV 블록(임베디드).
type Bar struct {
	Date   string  `json:"date"`   // 일시(YYYY-MM-DD HH:MM:SS)
	Open   float64 `json:"open"`   // 시가
	High   float64 `json:"high"`   // 고가
	Low    float64 `json:"low"`    // 저가
	Close  float64 `json:"close"`  // 종가
	Volume int64   `json:"volume"` // 거래량
}

// SMA — 단순 이동평균 (technical-indicators/sma)
type SMA struct {
	Bar
	SMA float64 `json:"sma"` // 단순 이동평균
}
// EMA{Bar; EMA float64 `json:"ema"`} — 지수 이동평균
// WMA{Bar; WMA float64 `json:"wma"`} — 가중 이동평균
// DEMA{Bar; DEMA float64 `json:"dema"`} — 이중 지수 이동평균
// TEMA{Bar; TEMA float64 `json:"tema"`} — 삼중 지수 이동평균
// RSI{Bar; RSI float64 `json:"rsi"`} — 상대강도지수
// StandardDeviation{Bar; StandardDeviation float64 `json:"standardDeviation"`} — 표준편차
// Williams{Bar; Williams float64 `json:"williams"`} — 윌리엄스 %R
// ADX{Bar; ADX float64 `json:"adx"`} — 평균방향지수
```

## 시그니처 규칙
- 9 메서드 전부: `(ctx, symbol string, periodLength int, timeframe, from, to string)` → `validate` 후 `fetch.List[T](ctx, c.http, path, indicatorParams(...))`.

## 테스트
- fixture 단위: SMA 파싱(Bar.Close!=0, SMA!=0 — 임베디드 승격 확인). RSI 파싱(RSI!=0). StandardDeviation 파싱(키 standardDeviation 매핑). 대표 3~4종.
- delegation: SMA(symbol,10,"1day",from,to) path+symbol/periodLength/timeframe/from/to / ADX path / StandardDeviation path(소문자 standarddeviation).
- 가드: 빈 symbol, periodLength 0, 빈 timeframe 각 대표 1건.
- 통합: SMA("AAPL",10,"1day","","") len>0 & Close>0 & SMA>0 / RSI("AAPL",14,"1day","","") / ADX("AAPL",14,"1day","","").

## 문서 / 릴리스
- README Technical Indicators 행(9 endpoint).
- `examples/technicals/main.go` — SMA + RSI.
- 릴리스 `v0.18.0`.

## 범위 밖 / 위험
- timeframe 유효값(1min~1day) 검증 안 함(빈값만 가드, FMP 위임).
- standard-deviation path/키 불일치 주의(테스트로 확인).
- 다음 그룹: crypto/forex/commodity 또는 secFilings/senate/form13F.
