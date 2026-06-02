# FMP Go SDK — Calendar 그룹 (v0.8.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/calendar-group`
- 토픽: FMP `calendar` 카테고리 9 endpoint 추가. 전체 API 커버리지 캠페인 6번째 그룹. 배당/실적/IPO/분할 캘린더.

## 결정 사항 (브레인스토밍)

- **범위**: calendar 9 endpoint 전부. 신규 `calendar/` 패키지, `internal/fetch` 사용.
- **전부 list**. calendar/ipos 7개 = `(from, to string)` 날짜범위, company 변형 3개 = `(symbol)`.
- **struct 재사용**: dividends/earnings/splits 의 calendar+company 변형이 각각 `Dividend`/`Earning`/`Split` 공유.
- **nullable**: Earning(eps/revenue), IPO(shares/priceRange/marketCap) → 포인터.
- **page 미노출**: 일부 calendar 지원하나 date range 가 주 필터(YAGNI).
- **템플릿 계승**: 필드 한국어 주석, fixture + delegation 테스트.
- **릴리스**: `v0.8.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | helper | 반환 |
|---|---|---|---|---|
| `dividends.go` | `DividendsCalendar(ctx, from, to)` | `/stable/dividends-calendar` | List{from,to} | `[]Dividend` |
| | `CompanyDividends(ctx, symbol)` | `/stable/dividends` | ListBySymbol | `[]Dividend` |
| `earnings.go` | `EarningsCalendar(ctx, from, to)` | `/stable/earnings-calendar` | List{from,to} | `[]Earning` |
| | `CompanyEarnings(ctx, symbol)` | `/stable/earnings` | ListBySymbol | `[]Earning` |
| `ipos.go` | `IPOsCalendar(ctx, from, to)` | `/stable/ipos-calendar` | List{from,to} | `[]IPO` |
| | `IPODisclosures(ctx, from, to)` | `/stable/ipos-disclosure` | List{from,to} | `[]IPODisclosure` |
| | `IPOProspectuses(ctx, from, to)` | `/stable/ipos-prospectus` | List{from,to} | `[]IPOProspectus` |
| `splits.go` | `SplitsCalendar(ctx, from, to)` | `/stable/splits-calendar` | List{from,to} | `[]Split` |
| | `CompanySplits(ctx, symbol)` | `/stable/splits` | ListBySymbol | `[]Split` |
| `client.go` | `New(http)` | — | — | `*Client` |

`dateRange(from, to string) map[string]string` — 패키지 헬퍼, 빈 값 제외.

## 응답 타입 (faithful, 필드 한국어 주석)

```go
// Dividend — 배당 (dividends-calendar / dividends 공용)
type Dividend struct {
	Symbol          string  `json:"symbol"`          // 종목 심볼
	Date            string  `json:"date"`            // 배당락일
	RecordDate      string  `json:"recordDate"`      // 기준일
	PaymentDate     string  `json:"paymentDate"`     // 지급일
	DeclarationDate string  `json:"declarationDate"` // 선언일
	AdjDividend     float64 `json:"adjDividend"`     // 수정 배당금
	Dividend        float64 `json:"dividend"`        // 배당금
	Yield           float64 `json:"yield"`           // 배당수익률 (%)
	Frequency       string  `json:"frequency"`       // 배당 주기 (예: Quarterly)
}

// Earning — 실적 (earnings-calendar / earnings 공용). 미래 실적은 actual/estimated null.
type Earning struct {
	Symbol           string   `json:"symbol"`           // 종목 심볼
	Date             string   `json:"date"`             // 실적 발표일
	EpsActual        *float64 `json:"epsActual"`        // 실제 EPS(결측 가능)
	EpsEstimated     *float64 `json:"epsEstimated"`     // 추정 EPS(결측 가능)
	RevenueActual    *int64   `json:"revenueActual"`    // 실제 매출(결측 가능)
	RevenueEstimated *int64   `json:"revenueEstimated"` // 추정 매출(결측 가능)
	LastUpdated      string   `json:"lastUpdated"`      // 최종 갱신일
}

// IPO — IPO 일정 (ipos-calendar). shares/priceRange/marketCap 결측 가능.
type IPO struct {
	Symbol     string  `json:"symbol"`     // 종목 심볼
	Date       string  `json:"date"`       // IPO 일자
	Daa        string  `json:"daa"`        // 공시 일시 (ISO8601)
	Company    string  `json:"company"`    // 회사명
	Exchange   string  `json:"exchange"`   // 거래소
	Actions    string  `json:"actions"`    // 상태 (예: Expected)
	Shares     *int64  `json:"shares"`     // 공모 주식 수(결측 가능)
	PriceRange *string `json:"priceRange"` // 공모가 범위(결측 가능)
	MarketCap  *int64  `json:"marketCap"`  // 시가총액(결측 가능)
}

// IPODisclosure — IPO 공시 서류 (ipos-disclosure)
type IPODisclosure struct {
	Symbol            string `json:"symbol"`            // 종목 심볼
	FilingDate        string `json:"filingDate"`        // 제출일
	AcceptedDate      string `json:"acceptedDate"`      // 수리일
	EffectivenessDate string `json:"effectivenessDate"` // 효력 발생일
	CIK               string `json:"cik"`               // SEC CIK
	Form              string `json:"form"`              // 공시 양식 (예: CERT)
	URL               string `json:"url"`               // 원문 URL
}

// IPOProspectus — IPO 투자설명서 (ipos-prospectus)
type IPOProspectus struct {
	Symbol                          string  `json:"symbol"`                          // 종목 심볼
	AcceptedDate                    string  `json:"acceptedDate"`                    // 수리일
	FilingDate                      string  `json:"filingDate"`                      // 제출일
	IPODate                         string  `json:"ipoDate"`                         // IPO 일자
	CIK                             string  `json:"cik"`                             // SEC CIK
	PricePublicPerShare             float64 `json:"pricePublicPerShare"`             // 주당 공모가
	PricePublicTotal                float64 `json:"pricePublicTotal"`                // 총 공모금액
	DiscountsAndCommissionsPerShare float64 `json:"discountsAndCommissionsPerShare"` // 주당 인수수수료
	DiscountsAndCommissionsTotal    float64 `json:"discountsAndCommissionsTotal"`    // 총 인수수수료
	ProceedsBeforeExpensesPerShare  float64 `json:"proceedsBeforeExpensesPerShare"`  // 주당 순수취금(비용 전)
	ProceedsBeforeExpensesTotal     float64 `json:"proceedsBeforeExpensesTotal"`     // 총 순수취금(비용 전)
	Form                            string  `json:"form"`                            // 공시 양식 (예: 424B4)
	URL                             string  `json:"url"`                             // 원문 URL
}

// Split — 주식 분할 (splits-calendar / splits 공용)
type Split struct {
	Symbol      string `json:"symbol"`      // 종목 심볼
	Date        string `json:"date"`        // 분할 기준일
	Numerator   int    `json:"numerator"`   // 분할 비율 분자
	Denominator int    `json:"denominator"` // 분할 비율 분모
}
```

## 시그니처 규칙
- calendar/ipos 7개: `(ctx, from, to string)` → `fetch.List[T](ctx, c.http, path, dateRange(from, to))`. `dateRange` 가 빈 from/to 제외.
- company 3개: `(ctx, symbol string)` → `fetch.ListBySymbol`(가드 내장).

## 루트 Client 와이어
```go
type Client struct {
	...
	Analyst  *analyst.Client
	Calendar *calendar.Client // 캘린더 (신규)
}
```
`NewClient` 에 `c.Calendar = calendar.New(hc)`. `client_test.go` 에 `TestNewClient_HasCalendar`.

## 테스트
- fixture 단위: 6 struct. Earning null+값 양쪽 → 포인터 null→nil. IPO shares/priceRange/marketCap null 검증.
- delegation: DividendsCalendar(from,to) path+from/to / CompanyDividends(symbol) path+symbol. `dateRange` 빈값 생략 단위테스트.
- 가드: company 빈 symbol 대표 1건.
- 통합(`//go:build integration`): EarningsCalendar(범위) / CompanyDividends("AAPL") / SplitsCalendar(범위) / IPOsCalendar(범위).

## 문서 / 릴리스
- README 커버리지 표 Calendar 행(9 endpoint).
- `examples/calendar/main.go` — EarningsCalendar + CompanyDividends.
- 릴리스 `v0.8.0`.

## 범위 밖 / 위험
- 나머지 22 그룹 별도 PR(다음: statements 확장).
- page 파라미터 미노출(date range 주 필터, 후속 가능).
- IPO.PriceRange 실제 값 타입(string 가정) — 카탈로그 예시 null 뿐 → 통합테스트 확인, 다르면 조정.
- 날짜 형식 `YYYY-MM-DD` 는 호출자 책임(SDK 는 문자열 그대로 전달).
