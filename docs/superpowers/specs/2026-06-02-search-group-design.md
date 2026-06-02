# FMP Go SDK — Search 그룹 (v0.5.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/search-group`
- 토픽: FMP `search` 카테고리 7 endpoint 추가. 전체 API 커버리지 캠페인 3번째 그룹.

## 배경 / 목적

quote(v0.3.0)·company+internal/fetch(v0.4.0)로 템플릿·공유 helper 확립. 추천 순서(search → news → analyst → calendar → statements 확장)의 첫 그룹 `search` — 작고 foundational, 외부 개발자가 심볼/회사명/식별자(CIK/CUSIP/ISIN) 검색·스크리너로 가장 먼저 찾는 기능.

## 결정 사항 (브레인스토밍)

- **범위**: search 7 endpoint 전부. 신규 `search/` 패키지, `internal/fetch` 사용.
- **전부 list 반환**: 검색은 다건 → 모든 메서드 `[]T`(단일 `*T` 없음).
- **struct 재사용**: search-symbol/search-name → `SymbolSearchResult` 공용.
- **screener**: 19 필터 → `ScreenerParams` struct + `toMap()`(statements.Params 패턴). 숫자 0=생략, boolean 은 `*bool`(false/미지정 구분).
- **템플릿 계승**: 필드 한국어 주석, fixture + delegation 테스트, README/examples.
- **릴리스**: `v0.5.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | helper | 반환 |
|---|---|---|---|---|
| `search.go` | `SearchSymbol(ctx, query)` | `/stable/search-symbol` | List{query}+가드 | `[]SymbolSearchResult` |
| | `SearchName(ctx, query)` | `/stable/search-name` | List{query}+가드 | `[]SymbolSearchResult` |
| `identifiers.go` | `SearchCIK(ctx, cik)` | `/stable/search-cik` | List{cik}+가드 | `[]CIKSearchResult` |
| | `SearchCUSIP(ctx, cusip)` | `/stable/search-cusip` | List{cusip}+가드 | `[]CUSIPSearchResult` |
| | `SearchISIN(ctx, isin)` | `/stable/search-isin` | List{isin}+가드 | `[]ISINSearchResult` |
| `variants.go` | `SearchExchangeVariants(ctx, symbol)` | `/stable/search-exchange-variants` | ListBySymbol | `[]ExchangeVariant` |
| `screener.go` | `CompanyScreener(ctx, params)` | `/stable/company-screener` | List{params.toMap()} | `[]ScreenerResult` |
| `client.go` | `New(http)` | — | — | `*Client` |

## 응답 타입 (faithful, 필드 한국어 주석)

```go
// SymbolSearchResult — 심볼/회사명 검색 결과 (search-symbol / search-name 공용)
type SymbolSearchResult struct {
	Symbol           string `json:"symbol"`           // 종목 심볼
	Name             string `json:"name"`             // 종목/회사명
	Currency         string `json:"currency"`         // 통화
	ExchangeFullName string `json:"exchangeFullName"` // 거래소 전체명
	Exchange         string `json:"exchange"`         // 거래소 코드
}

// CIKSearchResult — CIK 검색 결과
type CIKSearchResult struct {
	Symbol           string `json:"symbol"`           // 종목 심볼
	CompanyName      string `json:"companyName"`      // 회사명
	CIK              string `json:"cik"`              // SEC CIK
	ExchangeFullName string `json:"exchangeFullName"` // 거래소 전체명
	Exchange         string `json:"exchange"`         // 거래소 코드
	Currency         string `json:"currency"`         // 통화
}

// CUSIPSearchResult — CUSIP 검색 결과
type CUSIPSearchResult struct {
	Symbol      string  `json:"symbol"`      // 종목 심볼
	CompanyName string  `json:"companyName"` // 회사명
	CUSIP       string  `json:"cusip"`       // CUSIP 코드
	MarketCap   float64 `json:"marketCap"`   // 시가총액
}

// ISINSearchResult — ISIN 검색 결과
type ISINSearchResult struct {
	Symbol    string  `json:"symbol"`    // 종목 심볼
	Name      string  `json:"name"`      // 회사명
	ISIN      string  `json:"isin"`      // ISIN 코드
	MarketCap int64   `json:"marketCap"` // 시가총액
}

// ExchangeVariant — 거래소별 심볼 변형 (profile 유사). 정확 필드는 fixture 로 확정.
type ExchangeVariant struct {
	Symbol      string  `json:"symbol"`      // 종목 심볼
	Price       float64 `json:"price"`       // 현재가
	Beta        float64 `json:"beta"`        // 베타
	VolAvg      int64   `json:"volAvg"`      // 평균 거래량
	MktCap      int64   `json:"mktCap"`      // 시가총액
	LastDiv     float64 `json:"lastDiv"`     // 최근 배당
	Range       string  `json:"range"`       // 52주 범위
	Changes     float64 `json:"changes"`     // 등락액
	CompanyName string  `json:"companyName"` // 회사명
	Currency    string  `json:"currency"`    // 통화
	// 그 외 필드 fixture 확인 후 추가(profile 유사 다수 가능)
}

// ScreenerResult — 스크리너 결과. nullable(MarketCap/Beta/LastAnnualDividend) → 포인터.
type ScreenerResult struct {
	Symbol             string   `json:"symbol"`             // 종목 심볼
	CompanyName        string   `json:"companyName"`        // 회사명
	MarketCap          *int64   `json:"marketCap"`          // 시가총액(결측 가능)
	Sector             string   `json:"sector"`             // 섹터
	Industry           string   `json:"industry"`           // 산업
	Beta               *float64 `json:"beta"`               // 베타(결측 가능)
	Price              float64  `json:"price"`              // 현재가
	LastAnnualDividend *float64 `json:"lastAnnualDividend"` // 최근 연간 배당(결측 가능)
	Volume             int64    `json:"volume"`             // 거래량
	Exchange           string   `json:"exchange"`           // 거래소
	ExchangeShortName  string   `json:"exchangeShortName"`  // 거래소 약칭
	Country            string   `json:"country"`            // 국가
	IsEtf              bool     `json:"isEtf"`              // ETF 여부
	IsFund             bool     `json:"isFund"`             // 펀드 여부
	IsActivelyTrading  bool     `json:"isActivelyTrading"`  // 거래 활성 여부
}
```

## ScreenerParams (다중 필터)

```go
// ScreenerParams — company-screener 필터. 빈 값/0/nil 은 쿼리에서 생략.
type ScreenerParams struct {
	MarketCapMoreThan      int64
	MarketCapLowerThan     int64
	Sector                 string
	Industry               string
	BetaMoreThan           float64
	BetaLowerThan          float64
	PriceMoreThan          float64
	PriceLowerThan         float64
	DividendMoreThan       float64
	DividendLowerThan      float64
	VolumeMoreThan         int64
	VolumeLowerThan        int64
	Exchange               string
	Country                string
	IsEtf                  *bool
	IsFund                 *bool
	IsActivelyTrading      *bool
	Limit                  int
	IncludeAllShareClasses *bool
}
```
- `toMap()` 가 비제로/non-nil 만 쿼리 포함. boolean 은 `strconv.FormatBool`.
- 숫자 0 = 미지정(생략) — 단순화. boolean 만 `*bool`(false/미지정 구분).
- 빈 params 도 유효(전체 스크리닝).

## 시그니처 규칙
- query/cik/cusip/isin: 빈 가드(`strings.TrimSpace`) + `fetch.List[T](ctx, c.http, path, map[string]string{"<key>": value})`.
- exchange-variants: `fetch.ListBySymbol[ExchangeVariant]`.
- screener: `fetch.List[ScreenerResult](ctx, c.http, path, params.toMap())`.

## 루트 Client 와이어
```go
type Client struct {
	...
	Quote  *quote.Client
	Search *search.Client // 검색 (신규)
}
```
`NewClient` 에 `c.Search = search.New(hc)`. `client_test.go` 에 `TestNewClient_HasSearch`.

## 테스트
- 7 endpoint fixture 단위테스트(faithful). symbol/name 공유 struct 양쪽 fixture.
- delegation: SearchSymbol(query) / SearchCIK(cik) / CompanyScreener(다중 param) path+param 매핑.
- 가드: 빈 query/cik/cusip/isin 대표 1건씩.
- `ScreenerParams.toMap()` 단위: 비제로/non-nil 만 포함, `*bool false`→`"false"`, 빈 params→빈 맵.
- nullable(ScreenerResult) null→nil fixture 검증.
- 통합(`//go:build integration`): SearchSymbol("AAPL")/SearchName("Apple")/CompanyScreener(Sector:Technology,Limit:5).

## 문서 / 릴리스
- README 커버리지 표 Search 행(7 endpoint).
- `examples/search/main.go` — SearchSymbol + CompanyScreener.
- 릴리스 `v0.5.0`.

## 범위 밖 / 위험
- 나머지 25 그룹 별도 PR(다음: news → analyst → calendar → statements 확장).
- ExchangeVariant 정확 필드(profile 유사 다수) 구현 시 fixture 확정 — 다르면 조정.
- screener 숫자 0=생략 단순화 한계(0 필터 필요한 드문 경우 미지원) — 후속 포인터화 가능.
- ISIN/CUSIP 의 marketCap 타입(int64 vs float64) fixture 로 확정 — search-cusip 예시는 float, search-isin 예시는 int 로 보임(개별 확인).
