# FMP Go SDK — ESG 그룹 (v0.21.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/esg-group`
- 토픽: FMP `ESG` 카테고리 3 endpoint. 캠페인 19번째 그룹.

## 결정 사항
- 신규 `esg/` 패키지, internal/fetch. 3 구조체(공유 안 함).
- `ESGRating`(esg-ratings), `ESGDisclosure`(esg-disclosures), `ESGBenchmark`(esg-benchmark).
- **path 주의**: esg-search.md 의 실제 path 는 `/stable/esg-disclosures`.
- 점수 키 casing: environmentalScore/socialScore/governanceScore + 합성 `ESGScore`(대문자 ESG). ratings 등급 `ESGRiskRating`.
- fiscalYear 는 숫자(int). cik 0-padded 문자열.
- Ratings/Disclosures 는 symbol 필수, Benchmark 는 year 필수.
- 릴리스 `v0.21.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `Ratings(ctx, symbol)` | `/stable/esg-ratings` | symbol(필수) | `[]ESGRating` |
| `Disclosures(ctx, symbol)` | `/stable/esg-disclosures` | symbol(필수) | `[]ESGDisclosure` |
| `Benchmark(ctx, year)` | `/stable/esg-benchmark` | year(필수) | `[]ESGBenchmark` |

파일: `esg/client.go`(New), `esg/esg.go`(3 struct + 3 method).
- Ratings/Disclosures: symbol 빈값 가드 + {symbol}. Benchmark: year 빈값 가드 + {year}.

## 루트 Client 와이어
```go
Senate *senate.Client
ESG    *esg.Client // ESG 평가/공시/벤치마크
```
`c.ESG = esg.New(hc)`. `TestNewClient_HasESG`.

## 응답 타입 (faithful, 필드 한국어 주석)
```go
// ESGRating — ESG 등급 (esg-ratings)
type ESGRating struct {
	Symbol        string `json:"symbol"`        // 종목 심볼
	CIK           string `json:"cik"`           // SEC CIK(0-padded)
	CompanyName   string `json:"companyName"`   // 회사명
	Industry      string `json:"industry"`      // 산업
	FiscalYear    int    `json:"fiscalYear"`    // 회계연도
	ESGRiskRating string `json:"ESGRiskRating"` // ESG 리스크 등급(예: B)
	IndustryRank  string `json:"industryRank"`  // 산업 내 순위("4 out of 5")
}

// ESGDisclosure — ESG 공시 (esg-disclosures)
type ESGDisclosure struct {
	Date               string  `json:"date"`               // 공시일
	AcceptedDate       string  `json:"acceptedDate"`       // 수리일
	Symbol             string  `json:"symbol"`             // 종목 심볼
	CIK                string  `json:"cik"`                // SEC CIK(0-padded)
	CompanyName        string  `json:"companyName"`        // 회사명
	FormType           string  `json:"formType"`           // 공시 양식(8-K 등)
	EnvironmentalScore float64 `json:"environmentalScore"` // 환경 점수
	SocialScore        float64 `json:"socialScore"`        // 사회 점수
	GovernanceScore    float64 `json:"governanceScore"`    // 지배구조 점수
	ESGScore           float64 `json:"ESGScore"`           // 종합 ESG 점수
	URL                string  `json:"url"`                // 공시 원문 URL
}

// ESGBenchmark — 섹터별 ESG 벤치마크 (esg-benchmark)
type ESGBenchmark struct {
	FiscalYear         int     `json:"fiscalYear"`         // 회계연도
	Sector             string  `json:"sector"`             // 섹터
	EnvironmentalScore float64 `json:"environmentalScore"` // 환경 점수
	SocialScore        float64 `json:"socialScore"`        // 사회 점수
	GovernanceScore    float64 `json:"governanceScore"`    // 지배구조 점수
	ESGScore           float64 `json:"ESGScore"`           // 종합 ESG 점수
}
```

## 시그니처 규칙
- Ratings/Disclosures: `(ctx, symbol)` → symbol 가드 + `fetch.List[T](..., {"symbol": symbol})`.
- Benchmark: `(ctx, year)` → year 가드 + `fetch.List[ESGBenchmark](..., {"year": year})`.

## 테스트
- fixture 단위: ESGRating(FiscalYear int, ESGRiskRating, IndustryRank 문자열), ESGDisclosure(ESGScore 키 매핑, scores), ESGBenchmark(Sector, ESGScore).
- delegation: Ratings("AAPL") path+symbol / Disclosures("AAPL") path `/stable/esg-disclosures`+symbol / Benchmark("2023") path+year.
- 가드: Ratings 빈 symbol, Benchmark 빈 year.
- 통합: Ratings("AAPL") len>0 / Disclosures("AAPL") err 체크 / Benchmark("2023") len>0.

## 문서 / 릴리스
- README ESG 행(3 endpoint).
- `examples/esg/main.go` — Ratings + Benchmark.
- 릴리스 `v0.21.0`.

## 범위 밖 / 위험
- 점수 필드 null 가능성(저커버 종목) — 카탈로그 비-null 이라 float64 유지.
- esg-disclosures path(파일명 esg-search 와 다름) 주의.
- 다음 그룹: earningsTranscript / etfAndMutualFunds / form13F / crypto·forex·commodity.
