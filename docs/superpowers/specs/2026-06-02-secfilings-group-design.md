# FMP Go SDK — SEC Filings 그룹 (v0.27.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/secfilings-group`
- 토픽: FMP `secFilings` 카테고리 12 endpoint. 캠페인 25번째 그룹.

## 결정 사항
- 신규 `secfilings/` 패키지, internal/fetch. 5 구조체.
- `LatestFiling`(financials-latest + 8k-latest 공유, hasFinancials bool 포함), `FilingSearchResult`(search-by-symbol/cik/form-type 공유, hasFinancials 없음), `CompanySearchResult`(search-by-name + company-search-by-symbol/cik + industry-classification-search + all-industry-classification 5개 공유), `IndustryClassification`(industry-classification-list), `CompanyProfile`(sec-profile, 35필드).
- **path family 주의**: 필링검색 `sec-filings-search/{symbol,cik,form-type}` vs 회사검색 `sec-filings-company-search/{symbol,cik,name}`. search-by-name 은 회사검색 계열(param 이름 `company`).
- sicCode/employees/fiftyTwoWeekRange 전부 string. securityType 만 nullable(*string). filingDate/acceptedDate datetime 문자열.
- 릴리스 `v0.27.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `LatestFinancials(ctx, from, to, page, limit)` | `/stable/sec-filings-financials` | from/to(필수),page,limit | `[]LatestFiling` |
| `Latest8K(ctx, from, to, page, limit)` | `/stable/sec-filings-8k` | from/to(필수),page,limit | `[]LatestFiling` |
| `SearchBySymbol(ctx, symbol, from, to, page, limit)` | `/stable/sec-filings-search/symbol` | symbol/from/to(필수),page,limit | `[]FilingSearchResult` |
| `SearchByCIK(ctx, cik, from, to, page, limit)` | `/stable/sec-filings-search/cik` | cik/from/to(필수),page,limit | `[]FilingSearchResult` |
| `SearchByFormType(ctx, formType, from, to, page, limit)` | `/stable/sec-filings-search/form-type` | formType/from/to(필수),page,limit | `[]FilingSearchResult` |
| `SearchByName(ctx, company)` | `/stable/sec-filings-company-search/name` | company(필수) | `[]CompanySearchResult` |
| `CompanySearchBySymbol(ctx, symbol)` | `/stable/sec-filings-company-search/symbol` | symbol(필수) | `[]CompanySearchResult` |
| `CompanySearchByCIK(ctx, cik)` | `/stable/sec-filings-company-search/cik` | cik(필수) | `[]CompanySearchResult` |
| `Profile(ctx, symbol, cik)` | `/stable/sec-profile` | symbol 또는 cik | `[]CompanyProfile` |
| `IndustryClassificationList(ctx, industryTitle, sicCode)` | `/stable/standard-industrial-classification-list` | 둘 다 선택 | `[]IndustryClassification` |
| `IndustryClassificationSearch(ctx, symbol, cik, sicCode)` | `/stable/industry-classification-search` | 전부 선택 | `[]CompanySearchResult` |
| `AllIndustryClassification(ctx, page, limit)` | `/stable/all-industry-classification` | page,limit | `[]CompanySearchResult` |

파일: `secfilings/client.go`(New + helpers), `secfilings/filings.go`(LatestFiling/FilingSearchResult + 5 method), `secfilings/company.go`(CompanySearchResult/IndustryClassification + 6 method), `secfilings/profile.go`(CompanyProfile + Profile method).

helpers:
- `fromToPage(from, to string, page, limit int)` — from/to 항상(검증은 메서드), page/limit>0 조건. 또는 메서드별 빌드.
- 필링 검색/latest: from/to 빈값 가드. 주 파라미터(symbol/cik/formType/company) 가드. Profile: symbol·cik 모두 비어있으면 에러.

## 루트 Client 와이어
```go
ETF        *etf.Client
SECFilings *secfilings.Client // SEC 공시 검색/분류/프로필
```
`c.SECFilings = secfilings.New(hc)`. `TestNewClient_HasSECFilings`.

## 응답 타입
```go
// LatestFiling — 최신 SEC 공시 (sec-filings-financials / sec-filings-8k 공유)
type LatestFiling struct {
	Symbol        string `json:"symbol"`        // 종목 심볼
	CIK           string `json:"cik"`           // CIK
	FilingDate    string `json:"filingDate"`    // 제출 일시
	AcceptedDate  string `json:"acceptedDate"`  // 수리 일시
	FormType      string `json:"formType"`      // 양식 유형
	HasFinancials bool   `json:"hasFinancials"` // 재무 포함 여부
	Link          string `json:"link"`          // 공시 링크
	FinalLink     string `json:"finalLink"`     // 최종 문서 링크
}

// FilingSearchResult — SEC 공시 검색 결과 (sec-filings-search/{symbol,cik,form-type} 공유)
type FilingSearchResult struct {
	Symbol       string `json:"symbol"`       // 종목 심볼
	CIK          string `json:"cik"`          // CIK
	FilingDate   string `json:"filingDate"`   // 제출 일시
	AcceptedDate string `json:"acceptedDate"` // 수리 일시
	FormType     string `json:"formType"`     // 양식 유형
	Link         string `json:"link"`         // 공시 링크
	FinalLink    string `json:"finalLink"`    // 최종 문서 링크
}

// CompanySearchResult — 회사 검색/산업분류 결과 (company-search + industry-classification 공유).
// businessAddress 는 endpoint 에 따라 평문 또는 리스트형 문자열.
type CompanySearchResult struct {
	Symbol          string `json:"symbol"`          // 종목 심볼(없으면 "None")
	Name            string `json:"name"`            // 회사명
	CIK             string `json:"cik"`             // CIK
	SICCode         string `json:"sicCode"`         // SIC 코드(문자열)
	IndustryTitle   string `json:"industryTitle"`   // 산업명
	BusinessAddress string `json:"businessAddress"` // 사업장 주소
	PhoneNumber     string `json:"phoneNumber"`     // 전화번호
}

// IndustryClassification — SIC 산업분류 목록 (standard-industrial-classification-list)
type IndustryClassification struct {
	Office        string `json:"office"`        // SEC 담당 office
	SICCode       string `json:"sicCode"`       // SIC 코드(문자열)
	IndustryTitle string `json:"industryTitle"` // 산업명
}

// CompanyProfile — SEC 회사 전체 프로필 (sec-profile)
type CompanyProfile struct {
	Symbol                  string  `json:"symbol"`
	CIK                     string  `json:"cik"`
	RegistrantName          string  `json:"registrantName"`
	SICCode                 string  `json:"sicCode"`
	SICDescription          string  `json:"sicDescription"`
	SICGroup                string  `json:"sicGroup"`
	ISIN                    string  `json:"isin"`
	BusinessAddress         string  `json:"businessAddress"`
	MailingAddress          string  `json:"mailingAddress"`
	PhoneNumber             string  `json:"phoneNumber"`
	PostalCode              string  `json:"postalCode"`
	City                    string  `json:"city"`
	State                   string  `json:"state"`
	Country                 string  `json:"country"`
	Description             string  `json:"description"`
	CEO                     string  `json:"ceo"`
	Website                 string  `json:"website"`
	Exchange                string  `json:"exchange"`
	StateLocation           string  `json:"stateLocation"`
	StateOfIncorporation    string  `json:"stateOfIncorporation"`
	FiscalYearEnd           string  `json:"fiscalYearEnd"`
	IPODate                 string  `json:"ipoDate"`
	Employees               string  `json:"employees"`
	SECFilingsURL           string  `json:"secFilingsUrl"`
	TaxIdentificationNumber string  `json:"taxIdentificationNumber"`
	FiftyTwoWeekRange       string  `json:"fiftyTwoWeekRange"`
	IsActive                bool    `json:"isActive"`
	AssetType               string  `json:"assetType"`
	OpenFigiComposite       string  `json:"openFigiComposite"`
	PriceCurrency           string  `json:"priceCurrency"`
	MarketSector            string  `json:"marketSector"`
	SecurityType            *string `json:"securityType"` // null 가능
	IsEtf                   bool    `json:"isEtf"`
	IsAdr                   bool    `json:"isAdr"`
	IsFund                  bool    `json:"isFund"`
}
```

## 시그니처 규칙
- LatestFinancials/Latest8K: `(ctx, from, to string, page, limit int)` → from/to 빈값 가드 + {from,to,page,limit?}.
- SearchBySymbol/ByCIK/ByFormType: `(ctx, <주param>, from, to string, page, limit int)` → 주param·from·to 빈값 가드 + {param,from,to,page,limit?}.
- SearchByName: `(ctx, company string)` → company 가드 + {company}.
- CompanySearchBySymbol/ByCIK: `(ctx, <param>)` → 가드 + {param}.
- Profile: `(ctx, symbol, cik string)` → symbol·cik 모두 빈값이면 에러 + {symbol?,cik?}.
- IndustryClassificationList: `(ctx, industryTitle, sicCode string)` → 빈값 제외 맵(가드 없음).
- IndustryClassificationSearch: `(ctx, symbol, cik, sicCode string)` → 빈값 제외 맵(가드 없음).
- AllIndustryClassification: `(ctx, page, limit int)` → {page,limit?}.

## 테스트
- fixture 단위: LatestFiling(HasFinancials bool) / FilingSearchResult(hasFinancials 없음) / CompanySearchResult(SICCode string) / IndustryClassification(Office) / CompanyProfile(IsActive/IsEtf bool, SecurityType null→nil, Employees string).
- delegation: LatestFinancials(from/to/page) / SearchBySymbol(symbol/from/to) / SearchByName(company) / Profile(symbol) / IndustryClassificationList(industryTitle) path+쿼리.
- 가드: LatestFinancials 빈 from, SearchBySymbol 빈 symbol, SearchByName 빈 company, Profile 빈 symbol+빈 cik (대표).
- 통합: Latest8K(from,to,0,5) / SearchBySymbol("AAPL",from,to,0,5) / Profile("AAPL","") / IndustryClassificationList("","") err 체크.

## 문서 / 릴리스
- README SEC Filings 행(12 endpoint).
- `examples/secfilings/main.go` — Profile + Latest8K.
- 릴리스 `v0.27.0`.

## 범위 밖 / 위험
- path family(sec-filings-search vs sec-filings-company-search) 혼동 주의.
- businessAddress 평문/리스트형 변형(둘 다 string).
- securityType 만 nullable(*string). sicCode/employees string.
- 다음 그룹: crypto·forex·commodity / bulk.
