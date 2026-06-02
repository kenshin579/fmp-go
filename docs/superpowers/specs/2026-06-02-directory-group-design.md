# FMP Go SDK — Directory 그룹 (v0.14.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/directory-group`
- 토픽: FMP `directory` 카테고리 11 endpoint. 캠페인 12번째 그룹. 심볼/거래소/섹터 등 기반 목록.

## 배경 / 목적

심볼/거래소/섹터/산업/국가 마스터 목록 — moneyflow 종목 enumerate·필터 기반. 외부 개발자 온보딩에도 필수.

## 결정 사항 (브레인스토밍)

- **범위**: directory 11 endpoint 전부. 신규 `directory/` 패키지, internal/fetch.
- **전부 객체 배열**(bare string 없음): available-sectors/industries/countries 도 `{sector}`/`{industry}`/`{country}` 단일필드 객체.
- **10 구조체**: `SymbolName`(etfs-list/actively-trading-list 공유) + `CompanySymbol` + `FinancialSymbol` + `CIKEntry` + `SymbolChange` + `TranscriptSymbol`(noOfTranscripts 문자열) + `Exchange` + `Sector` + `Industry` + `Country`.
- **쿼리**: 9개 무파라미터, CIKList(page,limit), SymbolChangesList(invalid bool, limit).
- **타입 주의**: cik 0-padded 문자열, noOfTranscripts 따옴표 문자열.
- **all list**, fetch.List. 무파라미터는 nil.
- **릴리스**: `v0.14.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | 반환 |
|---|---|---|---|
| `directory/symbols.go` | `CompanySymbolsList(ctx)` | `/stable/stock-list` | `[]CompanySymbol` |
| | `FinancialSymbolsList(ctx)` | `/stable/financial-statement-symbol-list` | `[]FinancialSymbol` |
| | `CIKList(ctx, page, limit)` | `/stable/cik-list` | `[]CIKEntry` |
| | `SymbolChangesList(ctx, invalid bool, limit int)` | `/stable/symbol-change` | `[]SymbolChange` |
| | `ETFsList(ctx)` | `/stable/etf-list` | `[]SymbolName` |
| | `ActivelyTradingList(ctx)` | `/stable/actively-trading-list` | `[]SymbolName` |
| | `EarningsTranscriptList(ctx)` | `/stable/earnings-transcript-list` | `[]TranscriptSymbol` |
| `directory/available.go` | `AvailableExchanges(ctx)` | `/stable/available-exchanges` | `[]Exchange` |
| | `AvailableSectors(ctx)` | `/stable/available-sectors` | `[]Sector` |
| | `AvailableIndustries(ctx)` | `/stable/available-industries` | `[]Industry` |
| | `AvailableCountries(ctx)` | `/stable/available-countries` | `[]Country` |
| `directory/client.go` | `New(http)` | — | — |

CIKList: `{page, limit}`(page 0 포함, limit>0). SymbolChangesList: `{invalid: "true"|"false", limit}`(limit>0).

## 루트 Client 와이어
```go
type Client struct {
	...
	MarketPerformance *marketperf.Client
	Directory         *directory.Client // 심볼/거래소/섹터 목록
}
```
`NewClient` 에 `c.Directory = directory.New(hc)`. `client_test.go` 에 `TestNewClient_HasDirectory`.

## 응답 타입 (faithful, 필드 한국어 주석)

```go
// SymbolName — 심볼+이름 (etf-list / actively-trading-list 공유)
type SymbolName struct {
	Symbol string `json:"symbol"` // 종목 심볼
	Name   string `json:"name"`   // 종목명
}

// CompanySymbol — 회사 심볼 목록 (stock-list)
type CompanySymbol struct {
	Symbol      string `json:"symbol"`      // 종목 심볼
	CompanyName string `json:"companyName"` // 회사명
}

// FinancialSymbol — 재무제표 제공 심볼 (financial-statement-symbol-list)
type FinancialSymbol struct {
	Symbol            string `json:"symbol"`            // 종목 심볼
	CompanyName       string `json:"companyName"`       // 회사명
	TradingCurrency   string `json:"tradingCurrency"`   // 거래 통화
	ReportingCurrency string `json:"reportingCurrency"` // 보고 통화
}

// CIKEntry — CIK 목록 (cik-list). cik 0-padded 문자열.
type CIKEntry struct {
	CIK         string `json:"cik"`         // SEC CIK(0-padded)
	CompanyName string `json:"companyName"` // 회사명
}

// SymbolChange — 심볼 변경 이력 (symbol-change)
type SymbolChange struct {
	Date        string `json:"date"`        // 변경일
	CompanyName string `json:"companyName"` // 회사명
	OldSymbol   string `json:"oldSymbol"`   // 이전 심볼
	NewSymbol   string `json:"newSymbol"`   // 신규 심볼
}

// TranscriptSymbol — 실적 발표 트랜스크립트 보유 심볼 (earnings-transcript-list).
// noOfTranscripts 는 FMP 가 문자열로 반환.
type TranscriptSymbol struct {
	Symbol         string `json:"symbol"`         // 종목 심볼
	CompanyName    string `json:"companyName"`    // 회사명
	NoOfTranscripts string `json:"noOfTranscripts"` // 트랜스크립트 수(문자열)
}

// Exchange — 거래소 목록 (available-exchanges)
type Exchange struct {
	Exchange     string `json:"exchange"`     // 거래소 코드
	Name         string `json:"name"`         // 거래소명
	CountryName  string `json:"countryName"`  // 국가명
	CountryCode  string `json:"countryCode"`  // 국가 코드
	SymbolSuffix string `json:"symbolSuffix"` // 심볼 접미사
	Delay        string `json:"delay"`        // 시세 지연(Real-time 등)
}

// Sector — 섹터 목록 (available-sectors)
type Sector struct {
	Sector string `json:"sector"` // 섹터명
}

// Industry — 산업 목록 (available-industries)
type Industry struct {
	Industry string `json:"industry"` // 산업명
}

// Country — 국가 목록 (available-countries)
type Country struct {
	Country string `json:"country"` // 국가 코드
}
```

## 시그니처 규칙
- 무파라미터 9개: `(ctx)` → `fetch.List[T](ctx, c.http, path, nil)`.
- CIKList: `(ctx, page, limit int)` → params {page: strconv, limit if >0}.
- SymbolChangesList: `(ctx, invalid bool, limit int)` → params {invalid: "true"|"false", limit if >0}.

## 테스트
- fixture 단위: 각 구조체 대표 파싱(SymbolName 공유 2 메서드, CIKEntry.CIK 0-padded 보존, TranscriptSymbol.NoOfTranscripts 문자열, Exchange 6필드, Sector/Industry/Country 단일필드).
- delegation: CompanySymbolsList path / CIKList(page,limit) path+page/limit / SymbolChangesList(true, 10) path+invalid=true/limit / AvailableExchanges path.
- 통합(`//go:build integration`): CompanySymbolsList len>0 / AvailableExchanges len>0 & 첫 행 Exchange!="" / AvailableSectors len>0.

## 문서 / 릴리스
- README 커버리지 표 Directory 행 신규(11 endpoint).
- `examples/directory/main.go` — AvailableExchanges + AvailableSectors.
- 릴리스 `v0.14.0`.

## 범위 밖 / 위험
- 무파라미터 목록은 대용량(stock-list 수만 건) — 페이지네이션 없음(FMP 미지원). 호출자 메모리 주의.
- noOfTranscripts/cik 문자열 유지(숫자 변환 호출자 책임).
- 다음 그룹: marketHours 또는 crypto/forex/commodity 등.
