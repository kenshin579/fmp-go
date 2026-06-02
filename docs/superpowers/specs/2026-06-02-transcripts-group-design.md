# FMP Go SDK — Earnings Transcripts 그룹 (v0.22.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/transcripts-group`
- 토픽: FMP `earningsTranscript` 카테고리 3 endpoint(중복 1개 제외). 캠페인 20번째 그룹.

## 결정 사항
- 신규 `transcripts/` 패키지, internal/fetch. 3 구조체.
- **available-transcript-symbols 제외**: 실제 path `/stable/earnings-transcript-list` 가 이미 `directory.EarningsTranscriptList`(TranscriptSymbol)로 구현됨 → 중복, 제외.
- `EarningCallTranscript`(content, search-transcripts), `LatestEarningCallTranscript`(latest), `EarningCallTranscriptDate`(dates).
- **명명 함정**: search 는 `period`(string "Q3") + `year`(int), latest 는 `period`(string) + `fiscalYear`(int), dates 는 `quarter`(int 1) + `fiscalYear`(int). 정확히 구분.
- search-transcripts 쿼리: symbol/year/quarter 전부 필수(문자열) + limit. content 는 큰 문자열.
- 릴리스 `v0.22.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `Transcript(ctx, symbol, year, quarter, limit)` | `/stable/earning-call-transcript` | symbol/year/quarter(필수),limit | `[]EarningCallTranscript` |
| `Latest(ctx, page, limit)` | `/stable/earning-call-transcript-latest` | page,limit | `[]LatestEarningCallTranscript` |
| `Dates(ctx, symbol)` | `/stable/earning-call-transcript-dates` | symbol(필수) | `[]EarningCallTranscriptDate` |

파일: `transcripts/client.go`(New + pageParams helper), `transcripts/transcripts.go`(3 struct + 3 method).
- Transcript: symbol/year/quarter 빈값 가드 + {symbol, year, quarter, limit?}.
- Latest: pageParams(page, limit).
- Dates: symbol 빈값 가드 + {symbol}.

## 루트 Client 와이어
```go
ESG                 *esg.Client
EarningsTranscripts *transcripts.Client // 실적 발표 트랜스크립트
```
`c.EarningsTranscripts = transcripts.New(hc)`. `TestNewClient_HasEarningsTranscripts`.

## 응답 타입 (faithful, 필드 한국어 주석)
```go
// EarningCallTranscript — 실적 발표 트랜스크립트 본문 (earning-call-transcript)
type EarningCallTranscript struct {
	Symbol  string `json:"symbol"`  // 종목 심볼
	Period  string `json:"period"`  // 분기(예: Q3)
	Year    int    `json:"year"`    // 연도
	Date    string `json:"date"`    // 발표일
	Content string `json:"content"` // 전문(텍스트)
}

// LatestEarningCallTranscript — 최신 트랜스크립트 목록 (earning-call-transcript-latest)
type LatestEarningCallTranscript struct {
	Symbol     string `json:"symbol"`     // 종목 심볼
	Period     string `json:"period"`     // 분기(예: Q3)
	FiscalYear int    `json:"fiscalYear"` // 회계연도
	Date       string `json:"date"`       // 발표일
}

// EarningCallTranscriptDate — 종목별 트랜스크립트 가용 일자 (earning-call-transcript-dates)
type EarningCallTranscriptDate struct {
	Quarter    int    `json:"quarter"`    // 분기(숫자 1~4)
	FiscalYear int    `json:"fiscalYear"` // 회계연도
	Date       string `json:"date"`       // 발표일
}
```

## 시그니처 규칙
- Transcript: `(ctx, symbol, year, quarter string, limit int)` → symbol/year/quarter 가드 + params{symbol, year, quarter, limit?}.
- Latest: `(ctx, page, limit int)` → pageParams.
- Dates: `(ctx, symbol string)` → symbol 가드 + {symbol}.

## 테스트
- fixture 단위: EarningCallTranscript(Year int, Period "Q3", Content!=""), LatestEarningCallTranscript(FiscalYear int, Period), EarningCallTranscriptDate(Quarter int 1, FiscalYear int).
- delegation: Transcript("AAPL","2020","Q3",0) path+symbol/year/quarter / Latest(0,10) path+page/limit / Dates("AAPL") path+symbol.
- 가드: Transcript 빈 symbol/year/quarter 각 1건(대표), Dates 빈 symbol.
- 통합: Latest(0,10) len>0 / Dates("AAPL") len>0 / Transcript("AAPL","2023","1",0) len>0 & Content!="".

## 문서 / 릴리스
- README Earnings Transcripts 행(3 endpoint).
- `examples/transcripts/main.go` — Latest + Dates.
- 릴리스 `v0.22.0`.

## 범위 밖 / 위험
- available-transcript-symbols 제외(directory 중복).
- 명명(year/fiscalYear/quarter/period) 혼동 주의 — 각 struct 정확히.
- 다음 그룹: etfAndMutualFunds / form13F / commitmentOfTraders / crypto·forex·commodity.
