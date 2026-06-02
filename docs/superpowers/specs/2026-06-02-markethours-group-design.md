# FMP Go SDK — Market Hours 그룹 (v0.16.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/markethours-group`
- 토픽: FMP `marketHours` 카테고리 3 endpoint. 캠페인 14번째 그룹.

## 결정 사항
- 신규 `markethours/` 패키지, internal/fetch. 2 구조체.
- `ExchangeHours`(exchange-market-hours / all-exchange-market-hours 공유), `ExchangeHoliday`(adjOpenTime/adjCloseTime nullable).
- ExchangeMarketHours/HolidaysByExchange 는 exchange 필수 가드. AllExchangeMarketHours 무파라미터(timestamp 미노출).
- 릴리스 `v0.16.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `ExchangeMarketHours(ctx, exchange)` | `/stable/exchange-market-hours` | exchange(필수) | `[]ExchangeHours` |
| `AllExchangeMarketHours(ctx)` | `/stable/all-exchange-market-hours` | — | `[]ExchangeHours` |
| `HolidaysByExchange(ctx, exchange, from, to)` | `/stable/holidays-by-exchange` | exchange(필수),from,to | `[]ExchangeHoliday` |

파일: `markethours/client.go`(New), `markethours/markethours.go`(2 struct + 3 method).

## 루트 Client 와이어
```go
Economics  *economics.Client
MarketHours *markethours.Client // 거래소 운영시간/휴장일
```
`c.MarketHours = markethours.New(hc)`. `TestNewClient_HasMarketHours`.

## 응답 타입 (faithful, 필드 한국어 주석)
```go
// ExchangeHours — 거래소 운영시간 (exchange-market-hours / all-exchange-market-hours 공유)
type ExchangeHours struct {
	Exchange     string `json:"exchange"`     // 거래소 코드
	Name         string `json:"name"`         // 거래소명
	OpeningHour  string `json:"openingHour"`  // 개장 시각(UTC offset 포함, 예 "09:30 AM -04:00")
	ClosingHour  string `json:"closingHour"`  // 폐장 시각
	Timezone     string `json:"timezone"`     // 타임존(예 America/New_York)
	IsMarketOpen bool   `json:"isMarketOpen"` // 현재 개장 여부
}

// ExchangeHoliday — 거래소 휴장일 (holidays-by-exchange). adj 시각은 null 가능.
type ExchangeHoliday struct {
	Exchange     string  `json:"exchange"`     // 거래소 코드
	Date         string  `json:"date"`         // 일자
	Name         string  `json:"name"`         // 휴일명
	IsClosed     bool    `json:"isClosed"`     // 휴장 여부
	AdjOpenTime  *string `json:"adjOpenTime"`  // 조정 개장 시각(null 가능)
	AdjCloseTime *string `json:"adjCloseTime"` // 조정 폐장 시각(null 가능)
}
```

## 시그니처 규칙
- ExchangeMarketHours: `(ctx, exchange string)` → exchange 빈값 가드 + `fetch.List[ExchangeHours](..., {"exchange": exchange})`.
- AllExchangeMarketHours: `(ctx)` → `fetch.List[ExchangeHours](..., nil)`.
- HolidaysByExchange: `(ctx, exchange, from, to string)` → exchange 빈값 가드 + params{exchange, from?, to?}.

## 테스트
- fixture 단위: ExchangeHours(IsMarketOpen bool, Timezone), ExchangeHoliday(IsClosed bool, AdjOpenTime/AdjCloseTime null→nil).
- delegation: ExchangeMarketHours("NASDAQ") path+exchange / AllExchangeMarketHours path / HolidaysByExchange("NASDAQ",from,to) path+exchange/from/to.
- 가드: ExchangeMarketHours 빈 exchange, HolidaysByExchange 빈 exchange.
- 통합: AllExchangeMarketHours len>0 / ExchangeMarketHours("NASDAQ") len>0 / HolidaysByExchange("NASDAQ","","") .

## 문서 / 릴리스
- README Market Hours 행(3 endpoint).
- `examples/markethours/main.go` — ExchangeMarketHours("NASDAQ") + AllExchangeMarketHours.
- 릴리스 `v0.16.0`.

## 범위 밖 / 위험
- timestamp 쿼리 미노출(현재 시각 기준, 후속 가능).
- 다음 그룹: crypto/forex/commodity 또는 insiderTrades.
