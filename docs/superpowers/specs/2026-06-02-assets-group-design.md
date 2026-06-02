# FMP Go SDK — Asset Lists 그룹 (v0.28.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/assets-group`
- 토픽: crypto/forex/commodity 카테고리의 **목록 endpoint 3개**. 캠페인 26번째 그룹.

## 배경 / 중복 분석 (중요)
crypto/forex/commodity 각 9 endpoint 중 **6개(quote/quote-short/all-quotes/eod-light/eod-full/intraday×3)는 이미 구현된 범용 endpoint** 와 동일 path:
- quote → `quote.Quote`(`/stable/quote`), quote-short → `quote.QuoteShort`, all-quotes → `quote.{Crypto,Forex,Commodity}Quotes`(`/stable/batch-*-quotes`)
- eod-light/full → `chart.HistoricalPriceEOD{Light,Full}`, intraday → `chart.Intraday{1Min,5Min,1Hour}`

따라서 **새로 추가할 고유 endpoint 는 자산군별 list 3개뿐**. 나머지는 기존 클라이언트로 `client.Quote.Quote(ctx,"BTCUSD")` / `client.Chart.HistoricalPriceEODLight(ctx,"BTCUSD","","")` 처럼 사용(README 명시).

## 결정 사항
- 신규 `assets/` 패키지, internal/fetch. 3 구조체(자산군별 list 형태가 모두 다름).
- `CryptoListItem`, `ForexPair`, `CommodityListItem`. 무파라미터.
- 릴리스 `v0.28.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | 반환 |
|---|---|---|
| `CryptoList(ctx)` | `/stable/cryptocurrency-list` | `[]CryptoListItem` |
| `ForexList(ctx)` | `/stable/forex-list` | `[]ForexPair` |
| `CommodityList(ctx)` | `/stable/commodities-list` | `[]CommodityListItem` |

파일: `assets/client.go`(New), `assets/assets.go`(3 struct + 3 method).

## 루트 Client 와이어
```go
SECFilings *secfilings.Client
Assets     *assets.Client // 암호화폐/외환/원자재 목록
```
`c.Assets = assets.New(hc)`. `TestNewClient_HasAssets`.

## 응답 타입
```go
// CryptoListItem — 암호화폐 목록 (cryptocurrency-list)
type CryptoListItem struct {
	Symbol            string   `json:"symbol"`            // 심볼(예: BTCUSD)
	Name              string   `json:"name"`              // 명칭
	Exchange          string   `json:"exchange"`          // 거래소
	IcoDate           string   `json:"icoDate"`           // ICO 일자
	CirculatingSupply float64  `json:"circulatingSupply"` // 유통 공급량
	TotalSupply       *float64 `json:"totalSupply"`       // 총 공급량(null 가능)
}

// ForexPair — 외환 페어 목록 (forex-list)
type ForexPair struct {
	Symbol       string `json:"symbol"`       // 심볼(예: EURUSD)
	FromCurrency string `json:"fromCurrency"` // 기준 통화
	ToCurrency   string `json:"toCurrency"`   // 상대 통화
	FromName     string `json:"fromName"`     // 기준 통화명
	ToName       string `json:"toName"`       // 상대 통화명
}

// CommodityListItem — 원자재 목록 (commodities-list)
type CommodityListItem struct {
	Symbol     string  `json:"symbol"`     // 심볼
	Name       string  `json:"name"`       // 명칭
	Exchange   *string `json:"exchange"`   // 거래소(null 가능)
	TradeMonth string  `json:"tradeMonth"` // 거래월(예: Dec)
	Currency   string  `json:"currency"`   // 통화
}
```

## 시그니처 규칙
- 3 메서드 모두 `(ctx)` → `fetch.List[T](ctx, c.http, path, nil)`.

## 테스트
- fixture 단위: CryptoListItem(CirculatingSupply float, TotalSupply null→nil), ForexPair(FromCurrency/ToCurrency), CommodityListItem(Exchange null→nil, TradeMonth).
- delegation: CryptoList path / ForexList path / CommodityList path.
- 통합: CryptoList len>0 / ForexList len>0 / CommodityList len>0.

## 문서 / 릴리스
- README Assets 행(3 endpoint) + 비고: 자산군 시세/시계열은 `client.Quote`/`client.Chart` 사용.
- `examples/assets/main.go` — CryptoList + ForexList.
- 릴리스 `v0.28.0`.

## 범위 밖 / 위험
- 자산군 quote/quote-short/all-quotes/eod/intraday 는 기존 quote·chart 로 커버(중복 구현 안 함).
- CommodityListItem.Exchange / CryptoListItem.TotalSupply nullable(*).
- 다음: bulk(마지막 카테고리) 후 캠페인 마무리.
