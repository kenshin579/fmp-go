# FMP Go SDK — Quote 그룹 (v0.3.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/quote-group`
- 토픽: FMP `quote` 카테고리 16 endpoint 를 SDK 에 추가. 전체 API 커버리지 캠페인의 첫 그룹(이후 모든 그룹 PR 의 템플릿).

## 배경 / 목적

fmp-go 는 현재 263 endpoint 중 4개(profile / income / balance / ratios)만 구현. 목표는 **전체 FMP API 커버리지**를 그룹 단위로 점진 구축해, (1) moneyflow 가 필요 시 즉시 소비하고 (2) 외부 개발자가 범용으로 쓸 수 있게 하는 것.

분해 전략(브레인스토밍 확정): **그룹 단위** — 한 그룹 = 한 PR + 한 minor 릴리스. 기존 v0.1.0(Company.Profile) / v0.2.0(Statements+Ratios) 패턴 연장. 첫 그룹으로 **`quote`** 선정 — FMP 에서 가장 많이 호출되는 그룹이라 외부 개발자 가치 1순위, 16 endpoint 가 유사 구조라 완결성 높음.

## 결정 사항 (브레인스토밍)

- **범위**: `quote` 그룹 16 endpoint **전부**. 5개 파일 + 5개 응답 타입.
- **필드 주석**: 모든 응답 struct 필드에 한국어 설명 주석. 이후 **모든 그룹 PR 의 표준**(외부 개발자 가시성).
- **시그니처 규칙**: 단일 심볼 → `*T`(빈 결과 ErrNotFound), 배치 → `symbols ...string` 가변 인자 `[]T`, 자산군 전체 → 무파라미터 `[]T`. (Company.Profile=`*T`, Statements=`[]T` 패턴 연장.)
- **faithful 1:1 매핑**: FMP 응답 필드를 그대로 노출. fixture 와 1:1 대조(statements 의 "39=39" 검증).
- **릴리스**: 완료 후 `./scripts/release.sh v0.3.0`.

## 패키지 구조 + endpoint 매핑

신규 `quote/` 패키지 (`statements/` 구조 미러링):

| 파일 | 메서드 | endpoint (path) | 반환 |
|---|---|---|---|
| `client.go` | `New(http)` | — | `*Client` |
| `quote.go` | `Quote(ctx, symbol)` | `/stable/quote` | `*Quote` |
| | `BatchQuote(ctx, symbols...)` | `/stable/batch-quote` | `[]Quote` |
| `short.go` | `QuoteShort(ctx, symbol)` | `/stable/quote-short` | `*QuoteShort` |
| | `BatchQuoteShort(ctx, symbols...)` | `/stable/batch-quote-short` | `[]QuoteShort` |
| `change.go` | `PriceChange(ctx, symbol)` | `/stable/stock-price-change` | `*PriceChange` |
| `aftermarket.go` | `AftermarketQuote(ctx, symbol)` | `/stable/aftermarket-quote` | `*AftermarketQuote` |
| | `AftermarketTrade(ctx, symbol)` | `/stable/aftermarket-trade` | `*AftermarketTrade` |
| | `BatchAftermarketQuote(ctx, symbols...)` | `/stable/batch-aftermarket-quote` | `[]AftermarketQuote` |
| | `BatchAftermarketTrade(ctx, symbols...)` | `/stable/batch-aftermarket-trade` | `[]AftermarketTrade` |
| `asset_class.go` | `ExchangeQuotes(ctx, exchange)` | `/stable/batch-exchange-quote` | `[]Quote` |
| | `IndexQuotes(ctx)` | `/stable/batch-index-quotes` | `[]QuoteShort` |
| | `CommodityQuotes(ctx)` | `/stable/batch-commodity-quotes` | `[]QuoteShort` |
| | `CryptoQuotes(ctx)` | `/stable/batch-crypto-quotes` | `[]QuoteShort` |
| | `ETFQuotes(ctx)` | `/stable/batch-etf-quotes` | `[]QuoteShort` |
| | `ForexQuotes(ctx)` | `/stable/batch-forex-quotes` | `[]QuoteShort` |
| | `MutualFundQuotes(ctx)` | `/stable/batch-mutualfund-quotes` | `[]QuoteShort` |

> **주의**: 카탈로그 파일명과 실제 path 불일치 — `quote-change.md` → `stock-price-change`, `full-*-quotes.md` → `batch-*-quotes`. 구현 시 각 `docs/api/quote/*.md` 의 `GET .../stable/...` 줄을 정확히 따른다.
> **주의**: asset-class 응답 shape(full `Quote` vs `QuoteShort`)은 카탈로그 기준 가정 — 구현 시 fixture 로 확정. 추가 필드 발견 시 그 그룹만 별도 struct.

## 응답 타입 (faithful, 필드 주석 포함)

```go
// Quote — 전체 시세 (quote / batch-quote / exchange-quote 공용)
type Quote struct {
	Symbol           string  `json:"symbol"`           // 종목 심볼 (예: AAPL)
	Name             string  `json:"name"`             // 종목명
	Price            float64 `json:"price"`            // 현재가
	ChangePercentage float64 `json:"changePercentage"` // 등락률 (%)
	Change           float64 `json:"change"`           // 전일 대비 등락액
	Volume           int64   `json:"volume"`           // 거래량
	DayLow           float64 `json:"dayLow"`           // 당일 저가
	DayHigh          float64 `json:"dayHigh"`          // 당일 고가
	YearHigh         float64 `json:"yearHigh"`         // 52주 최고가
	YearLow          float64 `json:"yearLow"`          // 52주 최저가
	MarketCap        int64   `json:"marketCap"`        // 시가총액
	PriceAvg50       float64 `json:"priceAvg50"`       // 50일 이동평균가
	PriceAvg200      float64 `json:"priceAvg200"`      // 200일 이동평균가
	Exchange         string  `json:"exchange"`         // 거래소 (예: NASDAQ)
	Open             float64 `json:"open"`             // 시가
	PreviousClose    float64 `json:"previousClose"`    // 전일 종가
	Timestamp        int64   `json:"timestamp"`        // 시세 시각 (Unix epoch sec)
}

// QuoteShort — 경량 시세 (quote-short / batch-quote-short / 자산군 배치 공용)
type QuoteShort struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	Price  float64 `json:"price"`  // 현재가
	Change float64 `json:"change"` // 전일 대비 등락액
	Volume int64   `json:"volume"` // 거래량
}

// PriceChange — 기간별 등락률(%) — stock-price-change
type PriceChange struct {
	Symbol string  `json:"symbol"` // 종목 심볼
	D1     float64 `json:"1D"`     // 1일 등락률 (%)
	D5     float64 `json:"5D"`     // 5일
	M1     float64 `json:"1M"`     // 1개월
	M3     float64 `json:"3M"`     // 3개월
	M6     float64 `json:"6M"`     // 6개월
	YTD    float64 `json:"ytd"`    // 연초 대비
	Y1     float64 `json:"1Y"`     // 1년
	Y3     float64 `json:"3Y"`     // 3년
	Y5     float64 `json:"5Y"`     // 5년
	Y10    float64 `json:"10Y"`    // 10년
	Max    float64 `json:"max"`    // 상장 이후 전체
}

// AftermarketQuote — 시간외 호가
type AftermarketQuote struct {
	Symbol    string  `json:"symbol"`    // 종목 심볼
	BidSize   int64   `json:"bidSize"`   // 매수 호가 수량
	BidPrice  float64 `json:"bidPrice"`  // 매수 호가
	AskSize   int64   `json:"askSize"`   // 매도 호가 수량
	AskPrice  float64 `json:"askPrice"`  // 매도 호가
	Volume    int64   `json:"volume"`    // 거래량
	Timestamp int64   `json:"timestamp"` // 시각 (Unix epoch ms)
}

// AftermarketTrade — 시간외 체결
type AftermarketTrade struct {
	Symbol    string  `json:"symbol"`    // 종목 심볼
	Price     float64 `json:"price"`     // 체결가
	TradeSize int64   `json:"tradeSize"` // 체결 수량
	Timestamp int64   `json:"timestamp"` // 시각 (Unix epoch ms)
}
```

- 금액/주가 `float64`, 거래량/시총/timestamp `int64`.
- 숫자 시작 JSON 키(`1D` 등)는 Go 필드명 불가 → `D1`/`M1`/`Y1` + 명시 태그.
- 각 endpoint fixture 와 1:1 대조로 누락/추가 필드 검증.

## 메서드 시그니처 규칙

- **단일 심볼** → `*T`. FMP 가 배열 반환하지만 1건이므로 첫 요소(Company.Profile 패턴). 빈 배열 → `httpclient.ErrNotFound`. 빈 symbol → 호출 전 가드 에러(`strings.TrimSpace`).
- **배치** → `symbols ...string` 가변 인자, 내부 `strings.Join(symbols, ",")`. 인자 0개 → 가드 에러.
- **자산군 전체** → 무파라미터.
- 모든 메서드 `c.http.GetJSON(ctx, "/stable/...", params, &out)` 위임.

## 루트 Client 와이어

```go
type Client struct {
	http       *httpclient.Client
	Company    *company.Client
	Statements *statements.Client
	Ratios     *ratios.Client
	Quote      *quote.Client // 시세 (신규)
}
```
`NewClient` 에 `c.Quote = quote.New(hc)` 추가. `client_test.go` 에 `TestNewClient_HasQuote` (non-nil).

## 에러
- 단일 심볼 빈 배열 → `httpclient.ErrNotFound`.
- 배치/자산군 빈 배열 → 그대로 반환(결과 없음이 정상일 수 있음).
- 빈 symbol / 빈 symbols → 호출 전 가드 에러.
- 비-200 / `Error Message` → `httpclient.APIError` (기존 매핑 재사용).

## 테스트
- **단위(fixture)**: 16 endpoint 각 `testdata/*.json`(카탈로그 응답 예시 추출) + faithful 디코딩 검증.
- **Delegation**: stub backend 로 path/params 위임 1~2 sample.
- **에러 경로**: 단일 빈 배열 → ErrNotFound 1건, 빈 symbol 가드 1건.
- **통합(`//go:build integration`)**: `FMP_API_KEY` 있으면 실 AAPL quote/quote-short/price-change + batch + crypto sample. 없으면 Skip.

## 문서 / examples
- `README.md` 커버리지 표에 `Quote` 행 추가(16 endpoint).
- `examples/quote/main.go` 신규 — Quote + batch + asset-class 사용 예시.

## 릴리스
- main 머지 → `./scripts/release.sh v0.3.0` → 태그 + GitHub Release.
- moneyflow 소비는 별도(필요 시점).

## 범위 밖 / 후속
- **나머지 27 그룹** — 각자 별도 spec → plan → PR → minor 릴리스. 본 작업이 템플릿(필드 주석/시그니처/fixture/README/examples).
- 다음 그룹 우선순위: company 완성 → search → news → analyst → calendar → statements 확장 → ...
- **moneyflow 통합** — 별도 작업.
- **WebSocket 실시간 스트림** — FMP 별도 상품, REST quote 그룹과 무관.

## 위험 / 주의
- asset-class 응답 shape 가정(exchange=full / 나머지=short) — fixture 로 확정, 다르면 별도 struct.
- 카탈로그 파일명 ≠ 실제 path — 각 .md 의 `GET` 줄 정확히 따름.
- 큰 정수(marketCap >2^53) 지수표기 시 디코딩 정밀도 — fixture 확인.
- rate limit — SDK 는 호출만, 한도 관리는 소비자 책임(README 명시).
