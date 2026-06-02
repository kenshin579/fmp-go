# FMP Go SDK — Form 13F 그룹 (v0.25.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/form13f-group`
- 토픽: FMP `form13F`(기관 13F 보유) 카테고리 8 endpoint. 캠페인 23번째 그룹.

## 결정 사항
- 신규 `form13f/` 패키지, internal/fetch. 8 구조체(공유 없음, 대형 다수).
- 대형 struct(HolderAnalytics 39, HolderPerformance 34, PositionSummary 39)는 **카탈로그 JSON 전수 매핑**(COT 선례).
- 공통 타입 규칙: 금액/주식/카운트/value/performance(절대) → int64, *Percentage/weight*/ratio/price/ownership*/turnover* → float64, 텍스트 → string, bool(isNew/isSoldOut/isCountedForPerformance) → bool, dates#3 의 year/quarter 는 숫자 int64.
- path 전부 `/stable/institutional-ownership/...`.
- 릴리스 `v0.25.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 | 카탈로그 |
|---|---|---|---|---|
| `LatestFilings(ctx, page, limit)` | `/stable/institutional-ownership/latest` | page,limit | `[]Filing` | latest-filings.md |
| `Extract(ctx, cik, year, quarter)` | `/stable/institutional-ownership/extract` | cik/year/quarter(필수) | `[]Holding` | filings-extract.md |
| `FilingDates(ctx, cik)` | `/stable/institutional-ownership/dates` | cik(필수) | `[]FilingDate` | form-13f-filings-dates.md |
| `ExtractAnalyticsByHolder(ctx, symbol, year, quarter, page, limit)` | `/stable/institutional-ownership/extract-analytics/holder` | symbol/year/quarter(필수),page,limit | `[]HolderAnalytics` | filings-extract-with-analytics-by-holder.md |
| `HolderPerformanceSummary(ctx, cik, page)` | `/stable/institutional-ownership/holder-performance-summary` | cik(필수),page | `[]HolderPerformance` | holder-performance-summary.md |
| `HoldersIndustryBreakdown(ctx, cik, year, quarter)` | `/stable/institutional-ownership/holder-industry-breakdown` | cik/year/quarter(필수) | `[]IndustryBreakdown` | holders-industry-breakdown.md |
| `PositionsSummary(ctx, symbol, year, quarter)` | `/stable/institutional-ownership/symbol-positions-summary` | symbol/year/quarter(필수) | `[]PositionSummary` | positions-summary.md |
| `IndustrySummary(ctx, year, quarter)` | `/stable/institutional-ownership/industry-summary` | year/quarter(필수) | `[]IndustrySummary` | industry-summary.md |

파일: `form13f/client.go`(New + pageParams), `form13f/filings.go`(Filing/Holding/FilingDate/IndustrySummary + 4 method), `form13f/analytics.go`(HolderAnalytics/IndustryBreakdown + 2 method), `form13f/performance.go`(HolderPerformance/PositionSummary + 2 method).

구현 시 각 struct 는 해당 카탈로그 파일 JSON 예시에서 **모든 키 전수 매핑**.

## 작은 4 struct (참고 — 구현은 카탈로그 확인)
- `Filing`(latest): cik,name,date,filingDate,acceptedDate,formType,link,finalLink (전부 string)
- `Holding`(extract): date,filingDate,acceptedDate,cik,securityCusip,symbol,nameOfIssuer,titleOfClass,sharesType,putCallShare,link,finalLink(string) + shares,value(int64)
- `FilingDate`(dates): date(string), year(int64), quarter(int64)
- `IndustrySummary`(industry-summary): industryTitle(string), industryValue(int64), date(string)

대형 4 struct(HolderAnalytics/HolderPerformance/IndustryBreakdown/PositionSummary)는 카탈로그 키 전수 매핑(타입 규칙 적용, bool 3종 주의).

## 루트 Client 와이어
```go
Fundraisers *fundraisers.Client
Form13F     *form13f.Client // 기관 13F 보유/분석
```
`c.Form13F = form13f.New(hc)`. `TestNewClient_HasForm13F`.

## 시그니처 규칙
- LatestFilings: pageParams. Extract/HoldersIndustryBreakdown/PositionsSummary: cik|symbol/year/quarter 가드 + {각 param}. FilingDates: cik 가드. ExtractAnalyticsByHolder: symbol/year/quarter 가드 + page/limit. HolderPerformanceSummary: cik 가드 + page. IndustrySummary: year/quarter 가드.
- 다중 필수 가드는 한 번에 검사(예: symbol/year/quarter 중 하나라도 빈값이면 에러).

## 테스트
- fixture 단위: 각 struct 카탈로그 예시 기반, 대표 필드 검증(특히 HolderAnalytics 의 bool isNew/isSoldOut, FilingDate 의 year/quarter int).
- delegation: LatestFilings(page/limit) / Extract(cik/year/quarter) / ExtractAnalyticsByHolder(symbol/year/quarter/page) / IndustrySummary(year/quarter) path+쿼리.
- 가드: Extract 빈 cik, PositionsSummary 빈 symbol, IndustrySummary 빈 year (대표).
- 통합: LatestFilings(0,5) / FilingDates("0001067983"=Berkshire) / PositionsSummary("AAPL","2023","3") / IndustrySummary("2023","3") err 체크.

## 문서 / 릴리스
- README Form 13F 행(8 endpoint).
- `examples/form13f/main.go` — LatestFilings + PositionsSummary.
- 릴리스 `v0.25.0`.

## 범위 밖 / 위험
- 대형 struct 카탈로그 전수 매핑(누락 키 없도록). bool 3종(#4) + year/quarter int(#3) 주의.
- nullable 없음(빈 문자열 가능 — plain string).
- 다음 그룹: etfAndMutualFunds / secFilings / crypto·forex·commodity / bulk.
