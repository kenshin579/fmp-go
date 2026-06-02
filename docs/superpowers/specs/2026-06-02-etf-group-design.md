# FMP Go SDK — ETF & Mutual Funds 그룹 (v0.26.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/etf-group`
- 토픽: FMP `etfAndMutualFunds` 카테고리 9 endpoint. 캠페인 24번째 그룹.

## 결정 사항
- 신규 `etf/` 패키지, internal/fetch. 10 구조체(9 + 중첩 1).
- **타입 함정 주의**: country-weighting `weightPercentage` 는 `string`("97.29%") 이고 symbol 없음; sector-weighting `weightPercentage` 는 `float64` 이고 symbol 있음 → 별도 struct. latest-disclosures 키는 `weightPercent`(Percentage 아님). marketValue 는 float64(소수). mutual-fund-disclosure 의 `is*`/`fairValLevel`/`entityOrgType` 은 "N"/"Y"/"2" string. cik 0-padded string. cur_cd snake_case.
- 모든 endpoint array 반환, page/limit 없음.
- 릴리스 `v0.26.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `Holdings(ctx, symbol)` | `/stable/etf/holdings` | symbol(필수) | `[]ETFHolding` |
| `Information(ctx, symbol)` | `/stable/etf/info` | symbol(필수) | `[]ETFInformation` |
| `CountryWeightings(ctx, symbol)` | `/stable/etf/country-weightings` | symbol(필수) | `[]ETFCountryWeighting` |
| `SectorWeightings(ctx, symbol)` | `/stable/etf/sector-weightings` | symbol(필수) | `[]ETFSectorWeighting` |
| `AssetExposure(ctx, symbol)` | `/stable/etf/asset-exposure` | symbol(필수) | `[]ETFAssetExposure` |
| `DisclosureHoldersSearch(ctx, name)` | `/stable/funds/disclosure-holders-search` | name(필수) | `[]FundDisclosureHolderSearch` |
| `DisclosureDates(ctx, symbol, cik)` | `/stable/funds/disclosure-dates` | symbol(필수),cik(선택) | `[]FundDisclosureDate` |
| `LatestDisclosureHolders(ctx, symbol)` | `/stable/funds/disclosure-holders-latest` | symbol(필수) | `[]FundDisclosureHolderLatest` |
| `Disclosure(ctx, symbol, year, quarter, cik)` | `/stable/funds/disclosure` | symbol/year/quarter(필수),cik(선택) | `[]MutualFundDisclosure` |

파일: `etf/client.go`(New), `etf/etf.go`(ETF 5종 struct+method), `etf/funds.go`(funds 4종 struct+method).

## 루트 Client 와이어
```go
Form13F *form13f.Client
ETF     *etf.Client // ETF/뮤추얼펀드 보유·정보
```
`c.ETF = etf.New(hc)`. `TestNewClient_HasETF`.

## 응답 타입
```go
// ETFHolding — ETF 보유 종목 (etf/holdings)
type ETFHolding struct {
	Symbol           string  `json:"symbol"`           // ETF 심볼
	Asset            string  `json:"asset"`            // 구성 종목 심볼
	Name             string  `json:"name"`             // 구성 종목명
	ISIN             string  `json:"isin"`             // ISIN(빈값 가능)
	SecurityCusip    string  `json:"securityCusip"`    // CUSIP(빈값 가능)
	SharesNumber     int64   `json:"sharesNumber"`     // 보유 주식 수
	WeightPercentage float64 `json:"weightPercentage"` // 비중(%)
	MarketValue      float64 `json:"marketValue"`      // 시장가치(소수)
	UpdatedAt        string  `json:"updatedAt"`        // 갱신 일시
	Updated          string  `json:"updated"`          // 갱신 일시2
}

// ETFSectorExposure — ETF info 의 sectorsList 중첩 원소.
type ETFSectorExposure struct {
	Industry string  `json:"industry"` // 산업
	Exposure float64 `json:"exposure"` // 노출 비중
}

// ETFInformation — ETF/펀드 프로필 (etf/info)
type ETFInformation struct {
	Symbol                string              `json:"symbol"`                // 심볼
	Name                  string              `json:"name"`                  // 명칭
	Description           string              `json:"description"`           // 설명(빈값 가능)
	ISIN                  string              `json:"isin"`                  // ISIN(빈값 가능)
	AssetClass            string              `json:"assetClass"`            // 자산군
	SecurityCusip         string              `json:"securityCusip"`         // CUSIP(빈값 가능)
	Domicile              string              `json:"domicile"`              // 소재지
	Website               string              `json:"website"`               // 웹사이트(빈값 가능)
	ETFCompany            string              `json:"etfCompany"`            // 운용사
	ExpenseRatio          float64             `json:"expenseRatio"`          // 보수율
	AssetsUnderManagement int64               `json:"assetsUnderManagement"` // 운용자산(AUM)
	AvgVolume             int64               `json:"avgVolume"`             // 평균 거래량
	InceptionDate         string              `json:"inceptionDate"`         // 설정일
	NAV                   float64             `json:"nav"`                   // 순자산가치
	NAVCurrency           string              `json:"navCurrency"`           // NAV 통화
	HoldingsCount         int                 `json:"holdingsCount"`         // 보유 종목 수
	UpdatedAt             string              `json:"updatedAt"`             // 갱신 일시(ISO)
	SectorsList           []ETFSectorExposure `json:"sectorsList"`           // 섹터 노출 목록
}

// ETFCountryWeighting — ETF 국가 비중 (etf/country-weightings). weightPercentage 는 "97.29%" 문자열.
type ETFCountryWeighting struct {
	Country          string `json:"country"`          // 국가
	WeightPercentage string `json:"weightPercentage"` // 비중 문자열("97.29%")
}

// ETFSectorWeighting — ETF 섹터 비중 (etf/sector-weightings). weightPercentage 는 숫자.
type ETFSectorWeighting struct {
	Symbol           string  `json:"symbol"`           // ETF 심볼
	Sector           string  `json:"sector"`           // 섹터
	WeightPercentage float64 `json:"weightPercentage"` // 비중(숫자)
}

// ETFAssetExposure — 특정 종목을 보유한 ETF 노출 (etf/asset-exposure)
type ETFAssetExposure struct {
	Symbol           string  `json:"symbol"`           // ETF 심볼
	Asset            string  `json:"asset"`            // 대상 종목
	SharesNumber     int64   `json:"sharesNumber"`     // 보유 주식 수
	WeightPercentage float64 `json:"weightPercentage"` // 비중(%)
	MarketValue      float64 `json:"marketValue"`      // 시장가치
}

// FundDisclosureHolderSearch — 펀드 공시 보유자 검색 (funds/disclosure-holders-search)
type FundDisclosureHolderSearch struct {
	Symbol              string `json:"symbol"`              // 심볼
	CIK                 string `json:"cik"`                 // CIK(0-padded)
	ClassID             string `json:"classId"`             // 클래스 ID
	SeriesID            string `json:"seriesId"`            // 시리즈 ID
	EntityName          string `json:"entityName"`          // 엔티티명
	EntityOrgType       string `json:"entityOrgType"`       // 엔티티 유형(문자열)
	SeriesName          string `json:"seriesName"`          // 시리즈명
	ClassName           string `json:"className"`           // 클래스명
	ReportingFileNumber string `json:"reportingFileNumber"` // 보고 파일 번호
	Address             string `json:"address"`             // 주소
	City                string `json:"city"`                // 시
	ZipCode             string `json:"zipCode"`             // 우편번호
	State               string `json:"state"`               // 주
}

// FundDisclosureDate — 펀드 공시 가용 일자 (funds/disclosure-dates)
type FundDisclosureDate struct {
	Date    string `json:"date"`    // 일자
	Year    int    `json:"year"`    // 연도
	Quarter int    `json:"quarter"` // 분기
}

// FundDisclosureHolderLatest — 최신 펀드 보유자 (funds/disclosure-holders-latest). 키 weightPercent.
type FundDisclosureHolderLatest struct {
	CIK           string  `json:"cik"`           // CIK(0-padded)
	Holder        string  `json:"holder"`        // 보유 펀드명
	Shares        int64   `json:"shares"`        // 보유 주식 수
	DateReported  string  `json:"dateReported"`  // 보고일
	Change        int64   `json:"change"`        // 변화량
	WeightPercent float64 `json:"weightPercent"` // 비중(키 weightPercent)
}

// MutualFundDisclosure — 뮤추얼펀드 보유 공시 (funds/disclosure). is* 는 "N"/"Y" 문자열.
type MutualFundDisclosure struct {
	CIK                 string  `json:"cik"`                 // CIK
	Date                string  `json:"date"`                // 일자
	AcceptedDate        string  `json:"acceptedDate"`        // 수리 일시
	Symbol              string  `json:"symbol"`              // 심볼
	Name                string  `json:"name"`                // 명칭
	LEI                 string  `json:"lei"`                 // LEI
	Title               string  `json:"title"`               // 제목
	CUSIP               string  `json:"cusip"`               // CUSIP(N/A 가능)
	ISIN                string  `json:"isin"`                // ISIN
	Balance             int64   `json:"balance"`             // 잔량
	Units               string  `json:"units"`               // 단위(NS 등)
	CurCd               string  `json:"cur_cd"`              // 통화 코드(snake_case 키)
	ValUsd              float64 `json:"valUsd"`              // USD 가치
	PctVal              float64 `json:"pctVal"`              // 비중(소수)
	PayoffProfile       string  `json:"payoffProfile"`       // 페이오프(Long 등)
	AssetCat            string  `json:"assetCat"`            // 자산 분류(EC 등)
	IssuerCat           string  `json:"issuerCat"`           // 발행인 분류(CORP 등)
	InvCountry          string  `json:"invCountry"`          // 투자 국가
	IsRestrictedSec     string  `json:"isRestrictedSec"`     // 제한 증권 여부("N"/"Y")
	FairValLevel        string  `json:"fairValLevel"`        // 공정가치 레벨("2" 등)
	IsCashCollateral    string  `json:"isCashCollateral"`    // 현금담보 여부("N"/"Y")
	IsNonCashCollateral string  `json:"isNonCashCollateral"` // 비현금담보 여부("N"/"Y")
	IsLoanByFund        string  `json:"isLoanByFund"`        // 펀드 대여 여부("N"/"Y")
}
```

## 시그니처 규칙
- ETF 5종 + LatestDisclosureHolders: `(ctx, symbol)` → symbol 가드 + {symbol}.
- DisclosureHoldersSearch: `(ctx, name)` → name 가드 + {name}.
- DisclosureDates: `(ctx, symbol, cik)` → symbol 가드 + {symbol, cik?}.
- Disclosure: `(ctx, symbol, year, quarter, cik)` → symbol/year/quarter 가드 + {symbol, year, quarter, cik?}.

## 테스트
- fixture 단위: ETFHolding(MarketValue float, SharesNumber int), ETFInformation(SectorsList 중첩 파싱, NAV float, HoldingsCount int), ETFCountryWeighting(WeightPercentage "97.29%" string), ETFSectorWeighting(WeightPercentage float), ETFAssetExposure, FundDisclosureHolderSearch, FundDisclosureDate(Year/Quarter int), FundDisclosureHolderLatest(WeightPercent 키), MutualFundDisclosure(IsRestrictedSec "N" string, CurCd cur_cd 키, PctVal float).
- delegation: Holdings(symbol) / DisclosureHoldersSearch(name) / Disclosure(symbol,year,quarter,cik) path+쿼리.
- 가드: Holdings 빈 symbol, DisclosureHoldersSearch 빈 name, Disclosure 빈 symbol/year/quarter (대표).
- 통합: Holdings("SPY") / Information("SPY") / SectorWeightings("SPY") / Disclosure("VFIAX","2023","4","") err 체크.

## 문서 / 릴리스
- README ETF & Mutual Funds 행(9 endpoint).
- `examples/etf/main.go` — Holdings + Information.
- 릴리스 `v0.26.0`.

## 범위 밖 / 위험
- country vs sector weightPercentage 타입 차이(string vs float) — 별도 struct 필수.
- is* "N"/"Y" string(bool 아님), cur_cd snake_case, weightPercent vs weightPercentage 키 차이.
- 다음 그룹: secFilings / crypto·forex·commodity / bulk.
