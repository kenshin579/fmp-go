# FMP Go SDK — Senate/House 그룹 (v0.20.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/senate-group`
- 토픽: FMP `senate` 카테고리 6 endpoint(상원/하원 의원 거래 공시). 캠페인 18번째 그룹.

## 결정 사항
- 신규 `senate/` 패키지, internal/fetch. **단일 구조체** `CongressTrade`(상원/하원 6 endpoint 공유, 15필드 전부 string).
- **amount 는 범위 문자열**("$1,001 - $15,000"), capitalGainsOver200USD 도 문자열("False"). null 없음(빈 문자열).
- 쿼리 3패턴: latest(page,limit), trades(symbol 필수+page,limit), trades-by-name(name 필수).
- 릴리스 `v0.20.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `SenateLatest(ctx, page, limit)` | `/stable/senate-latest` | page,limit | `[]CongressTrade` |
| `SenateTrades(ctx, symbol, page, limit)` | `/stable/senate-trades` | symbol(필수),page,limit | `[]CongressTrade` |
| `SenateTradesByName(ctx, name)` | `/stable/senate-trades-by-name` | name(필수) | `[]CongressTrade` |
| `HouseLatest(ctx, page, limit)` | `/stable/house-latest` | page,limit | `[]CongressTrade` |
| `HouseTrades(ctx, symbol, page, limit)` | `/stable/house-trades` | symbol(필수),page,limit | `[]CongressTrade` |
| `HouseTradesByName(ctx, name)` | `/stable/house-trades-by-name` | name(필수) | `[]CongressTrade` |

파일: `senate/client.go`(New + pageParams helper), `senate/senate.go`(CongressTrade + 6 method).
- `pageParams(page, limit int)` — page 항상(0 허용), limit>0.
- trades 2개: symbol 빈값 가드. by-name 2개: name 빈값 가드.

## 루트 Client 와이어
```go
DCF    *dcf.Client
Senate *senate.Client // 의회(상원/하원) 거래 공시
```
`c.Senate = senate.New(hc)`. `TestNewClient_HasSenate`.

## 응답 타입 (faithful, 필드 한국어 주석)
```go
// CongressTrade — 의회 의원 거래 공시 (senate/house latest·trades·by-name 공유).
// 모든 값이 문자열(amount 는 범위, capitalGainsOver200USD 는 "True"/"False").
type CongressTrade struct {
	Symbol                 string `json:"symbol"`                 // 종목 심볼
	DisclosureDate         string `json:"disclosureDate"`         // 공시일
	TransactionDate        string `json:"transactionDate"`        // 거래일
	FirstName              string `json:"firstName"`              // 이름
	LastName               string `json:"lastName"`               // 성
	Office                 string `json:"office"`                 // 의원명/직위
	District               string `json:"district"`               // 주(상원)/주+선거구(하원)
	Owner                  string `json:"owner"`                  // 소유자(Spouse/Joint 등)
	AssetDescription       string `json:"assetDescription"`       // 자산 설명
	AssetType              string `json:"assetType"`              // 자산 유형(Stock 등)
	Type                   string `json:"type"`                   // 거래 유형(Purchase/Sale)
	Amount                 string `json:"amount"`                 // 금액 범위("$1,001 - $15,000")
	CapitalGainsOver200USD string `json:"capitalGainsOver200USD"` // 200달러 초과 자본이득 여부(문자열)
	Comment                string `json:"comment"`                // 비고
	Link                   string `json:"link"`                   // 공시 원문 URL
}
```

## 시그니처 규칙
- SenateLatest/HouseLatest: `(ctx, page, limit int)` → `pageParams(page, limit)`.
- SenateTrades/HouseTrades: `(ctx, symbol string, page, limit int)` → symbol 가드 + pageParams + {symbol}.
- SenateTradesByName/HouseTradesByName: `(ctx, name string)` → name 가드 + {name}.

## 테스트
- fixture 단위: CongressTrade 파싱(Amount 범위 문자열, CapitalGainsOver200USD 문자열, Symbol/Type). senate fixture 1건(capitalGainsOver200USD 포함), latest fixture(해당 키 없음 → 빈 문자열) 검증.
- delegation: SenateLatest(0,10) path+page/limit / SenateTrades("AAPL",0,5) path+symbol / SenateTradesByName("Moran") path+name / HouseLatest path / HouseTrades path / HouseTradesByName path.
- 가드: SenateTrades 빈 symbol, SenateTradesByName 빈 name (대표).
- 통합: SenateLatest(0,10) len>0 / HouseLatest(0,10) len>0 / SenateTrades("AAPL",0,5) err 체크.

## 문서 / 릴리스
- README Senate/House 행(6 endpoint).
- `examples/senate/main.go` — SenateLatest + HouseLatest.
- 릴리스 `v0.20.0`.

## 범위 밖 / 위험
- amount/capitalGains 문자열 유지(파싱 호출자 책임).
- senate-latest 응답에 capitalGainsOver200USD 누락 → 빈 문자열(정상).
- 다음 그룹: ESG / earningsTranscript / form13F / crypto·forex·commodity.
