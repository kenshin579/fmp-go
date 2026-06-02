# FMP Go SDK — Insider Trades 그룹 (v0.17.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/insidertrades-group`
- 토픽: FMP `insiderTrades` 카테고리 6 endpoint. 캠페인 15번째 그룹.

## 결정 사항
- 신규 `insidertrades/` 패키지, internal/fetch. 5 구조체.
- `InsiderTrade`(latest + search 공유, 16필드 — directOrIndirect/formType 포함), `InsiderTransactionType`, `TradeStatistics`(13필드, cik 포함), `AcquisitionOwnership`(15필드 — 수치도 전부 string), `ReportingName`.
- **AcquisitionOwnership 수치 전부 string**: FMP 가 의결권/지분율을 문자열로 반환.
- search 는 옵션 6개 → `SearchParams` 구조체. statistics/acquisition/reporting-name 은 필수 인자 가드.
- 릴리스 `v0.17.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `LatestInsiderTrades(ctx, date, page, limit)` | `/stable/insider-trading/latest` | date,page,limit | `[]InsiderTrade` |
| `SearchInsiderTrades(ctx, p SearchParams)` | `/stable/insider-trading/search` | symbol,page,limit,reportingCik,companyCik,transactionType | `[]InsiderTrade` |
| `TransactionTypes(ctx)` | `/stable/insider-trading-transaction-type` | — | `[]InsiderTransactionType` |
| `Statistics(ctx, symbol)` | `/stable/insider-trading/statistics` | symbol(필수) | `[]TradeStatistics` |
| `AcquisitionOwnership(ctx, symbol, limit)` | `/stable/acquisition-of-beneficial-ownership` | symbol(필수),limit | `[]AcquisitionOwnership` |
| `SearchReportingName(ctx, name)` | `/stable/insider-trading/reporting-name` | name(필수) | `[]ReportingName` |

파일: `insidertrades/client.go`(New), `insidertrades/trades.go`(InsiderTrade + SearchParams + Latest/Search), `insidertrades/misc.go`(나머지 4 struct + 4 method).

> 메서드명 `AcquisitionOwnership` 가 타입명과 동일하나 Go 에서 메서드 식별자는 패키지 스코프가 아니라 합법(컴파일 정상).

## 루트 Client 와이어
```go
MarketHours  *markethours.Client
InsiderTrades *insidertrades.Client // 내부자 거래
```
`c.InsiderTrades = insidertrades.New(hc)`. `TestNewClient_HasInsiderTrades`.

## 응답 타입 (faithful, 필드 한국어 주석)
```go
// InsiderTrade — 내부자 거래 (insider-trading/latest, /search 공유)
type InsiderTrade struct {
	Symbol                   string  `json:"symbol"`                   // 종목 심볼
	FilingDate               string  `json:"filingDate"`               // 공시일
	TransactionDate          string  `json:"transactionDate"`          // 거래일
	ReportingCik             string  `json:"reportingCik"`             // 보고자 CIK
	CompanyCik               string  `json:"companyCik"`               // 회사 CIK
	TransactionType          string  `json:"transactionType"`          // 거래 유형(예: P-Purchase)
	SecuritiesOwned          int64   `json:"securitiesOwned"`          // 보유 증권 수
	ReportingName            string  `json:"reportingName"`            // 보고자명
	TypeOfOwner              string  `json:"typeOfOwner"`              // 소유자 유형
	AcquisitionOrDisposition string  `json:"acquisitionOrDisposition"` // 취득/처분(A/D)
	DirectOrIndirect         string  `json:"directOrIndirect"`         // 직접/간접
	FormType                 string  `json:"formType"`                 // 양식 유형(예: 4)
	SecuritiesTransacted     int64   `json:"securitiesTransacted"`     // 거래 증권 수
	Price                    float64 `json:"price"`                    // 단가
	SecurityName             string  `json:"securityName"`             // 증권명
	URL                      string  `json:"url"`                      // 공시 URL
}

// SearchParams — SearchInsiderTrades 옵션. 빈 값은 쿼리에서 제외.
type SearchParams struct {
	Symbol          string
	Page            int
	Limit           int
	ReportingCik    string
	CompanyCik      string
	TransactionType string
}

// InsiderTransactionType — 거래 유형 코드 목록 (insider-trading-transaction-type)
type InsiderTransactionType struct {
	TransactionType string `json:"transactionType"` // 거래 유형 코드
}

// TradeStatistics — 종목 내부자 거래 통계 (insider-trading/statistics)
type TradeStatistics struct {
	Symbol               string  `json:"symbol"`               // 종목 심볼
	CIK                  string  `json:"cik"`                  // 회사 CIK
	Year                 int64   `json:"year"`                 // 연도
	Quarter              int64   `json:"quarter"`              // 분기
	AcquiredTransactions int64   `json:"acquiredTransactions"` // 취득 거래 수
	DisposedTransactions int64   `json:"disposedTransactions"` // 처분 거래 수
	AcquiredDisposedRatio float64 `json:"acquiredDisposedRatio"` // 취득/처분 비율
	TotalAcquired        int64   `json:"totalAcquired"`        // 총 취득
	TotalDisposed        int64   `json:"totalDisposed"`        // 총 처분
	AverageAcquired      float64 `json:"averageAcquired"`      // 평균 취득
	AverageDisposed      float64 `json:"averageDisposed"`      // 평균 처분
	TotalPurchases       int64   `json:"totalPurchases"`       // 총 매수
	TotalSales           int64   `json:"totalSales"`           // 총 매도
}

// AcquisitionOwnership — 수익적 소유 취득 (acquisition-of-beneficial-ownership).
// FMP 가 의결권/지분율 수치를 전부 문자열로 반환 → 전 필드 string.
type AcquisitionOwnership struct {
	CIK                              string `json:"cik"`                              // CIK
	Symbol                           string `json:"symbol"`                           // 종목 심볼
	FilingDate                       string `json:"filingDate"`                       // 공시일
	AcceptedDate                     string `json:"acceptedDate"`                     // 수리일
	CUSIP                            string `json:"cusip"`                            // CUSIP
	NameOfReportingPerson            string `json:"nameOfReportingPerson"`            // 보고자명
	CitizenshipOrPlaceOfOrganization string `json:"citizenshipOrPlaceOfOrganization"` // 시민권/설립지
	SoleVotingPower                  string `json:"soleVotingPower"`                  // 단독 의결권
	SharedVotingPower                string `json:"sharedVotingPower"`                // 공동 의결권
	SoleDispositivePower             string `json:"soleDispositivePower"`             // 단독 처분권
	SharedDispositivePower           string `json:"sharedDispositivePower"`           // 공동 처분권
	AmountBeneficiallyOwned          string `json:"amountBeneficiallyOwned"`          // 수익적 소유량
	PercentOfClass                   string `json:"percentOfClass"`                   // 클래스 비율(%)
	TypeOfReportingPerson            string `json:"typeOfReportingPerson"`            // 보고자 유형
	URL                              string `json:"url"`                              // 공시 URL
}

// ReportingName — 보고자명 검색 결과 (insider-trading/reporting-name)
type ReportingName struct {
	ReportingCik  string `json:"reportingCik"`  // 보고자 CIK
	ReportingName string `json:"reportingName"` // 보고자명
}
```

## 시그니처 규칙
- LatestInsiderTrades: `(ctx, date string, page, limit int)` → params{date?, page, limit?}. (date 비어있으면 제외, page 항상, limit>0).
- SearchInsiderTrades: `(ctx, p SearchParams)` → p.queryParams()(빈 문자열/page 처리). 최소 1개 조건 권장이나 강제 안 함.
- TransactionTypes: `(ctx)` → nil.
- Statistics: `(ctx, symbol)` → symbol 가드 + {symbol}.
- AcquisitionOwnership: `(ctx, symbol, limit)` → symbol 가드 + {symbol, limit?}.
- SearchReportingName: `(ctx, name)` → name 가드 + {name}.

## 테스트
- fixture 단위: InsiderTrade(SecuritiesTransacted int, Price float, directOrIndirect/formType 파싱), InsiderTransactionType, TradeStatistics(AverageAcquired float, CIK), AcquisitionOwnership(PercentOfClass string), ReportingName.
- delegation: LatestInsiderTrades(date,page,limit) path+쿼리 / SearchInsiderTrades(SearchParams{Symbol,TransactionType}) path+symbol/transactionType / Statistics(symbol) path+symbol.
- 가드: Statistics 빈 symbol, AcquisitionOwnership 빈 symbol, SearchReportingName 빈 name.
- 통합: LatestInsiderTrades("",0,5) len>0 / SearchInsiderTrades({Symbol:"AAPL",Limit:5}) / Statistics("AAPL") / TransactionTypes len>0.

## 문서 / 릴리스
- README Insider Trades 행(6 endpoint).
- `examples/insidertrades/main.go` — SearchInsiderTrades(AAPL) + Statistics(AAPL).
- 릴리스 `v0.17.0`.

## 범위 밖 / 위험
- AcquisitionOwnership 수치 string 유지(숫자 변환 호출자 책임).
- SearchParams 무조건 허용(빈 검색은 FMP 위임).
- 다음 그룹: technicalIndicators 또는 crypto/forex.
