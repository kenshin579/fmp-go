# Earnings Transcripts 그룹 (v0.22.0) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development.

**Goal:** 신규 `transcripts/` 패키지 3 endpoint, 3 구조체.

**Architecture:** internal/fetch(List) + pageParams. 구조체는 스펙 verbatim. available-transcript-symbols 는 directory 중복이라 제외.

참고: `unset GOROOT`. 커밋 한국어 `feat(transcripts): ...`. 테스트 헬퍼는 calendar 패턴 패키지 내부 정의.

---

### Task 1: 패키지 + 3 endpoint

**Files:** Create `transcripts/client.go`, `transcripts/transcripts.go`, `transcripts/transcripts_test.go`, testdata `transcript.json`, `latest.json`, `dates.json`.

- [ ] **Step 1:** `transcripts/client.go`:
```go
// Package transcripts 는 FMP 실적 발표 트랜스크립트 API sub-client.
// fmp.Client.EarningsTranscripts 로 접근.
package transcripts

import (
	"strconv"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

type Client struct {
	http *httpclient.Client
}

func New(http *httpclient.Client) *Client { return &Client{http: http} }

func pageParams(page, limit int) map[string]string {
	q := map[string]string{"page": strconv.Itoa(page)}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return q
}
```
- [ ] **Step 2:** `transcripts/transcripts.go` — 스펙 3 struct + 3 메서드. import: context, fmt, strconv, strings, internal/fetch.
```go
func (c *Client) Transcript(ctx context.Context, symbol, year, quarter string, limit int) ([]EarningCallTranscript, error) {
	if strings.TrimSpace(symbol) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(quarter) == "" {
		return nil, fmt.Errorf("fmp: symbol, year, quarter must not be empty")
	}
	q := map[string]string{"symbol": symbol, "year": year, "quarter": quarter}
	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}
	return fetch.List[EarningCallTranscript](ctx, c.http, "/stable/earning-call-transcript", q)
}
func (c *Client) Latest(ctx context.Context, page, limit int) ([]LatestEarningCallTranscript, error) {
	return fetch.List[LatestEarningCallTranscript](ctx, c.http, "/stable/earning-call-transcript-latest", pageParams(page, limit))
}
func (c *Client) Dates(ctx context.Context, symbol string) ([]EarningCallTranscriptDate, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("fmp: symbol must not be empty")
	}
	return fetch.List[EarningCallTranscriptDate](ctx, c.http, "/stable/earning-call-transcript-dates", map[string]string{"symbol": symbol})
}
```
- [ ] **Step 3:** fixtures: transcript.json `[{symbol:"AAPL",period:"Q3",year:2020,date:"2020-07-30",content:"Operator: Good afternoon..."}]`, latest.json `[{symbol:"AAPL",period:"Q3",fiscalYear:2025,date:"2025-02-04"}]`, dates.json `[{quarter:1,fiscalYear:2025,date:"2025-01-30"}]`.
- [ ] **Step 4:** `transcripts_test.go`(헬퍼 calendar 패턴 정의): Transcript("AAPL","2020","Q3",0) 파싱(Year==2020, Period=="Q3", Content!="") + delegation(path+symbol/year/quarter) / Latest(0,10) 파싱(FiscalYear==2025, Period!="") + delegation(path+page/limit) / Dates("AAPL") 파싱(Quarter==1, FiscalYear==2025) + delegation(path+symbol) / 가드 Transcript 빈 symbol + Dates 빈 symbol.
- [ ] **Step 5:** `unset GOROOT && go test ./transcripts/ && go vet ./transcripts/ && gofmt -l transcripts/`. 커밋 `feat(transcripts): 트랜스크립트 본문/최신/일자 3종`.

### Task 2: 루트 와이어 + README + examples + 통합테스트

**Files:** Modify `client.go`, `client_test.go`, `README.md`, `integration_test.go`; Create `examples/transcripts/main.go`.

- [ ] **Step 1:** `client.go` — import `transcripts`, struct 에 `EarningsTranscripts *transcripts.Client`, NewClient 에 `c.EarningsTranscripts = transcripts.New(hc)`. (gofmt -w client.go.)
- [ ] **Step 2:** `client_test.go` — `TestNewClient_HasEarningsTranscripts`.
- [ ] **Step 3:** README 표 행 신규: `| Earnings Transcripts | \`client.EarningsTranscripts\` | Transcript, Latest, Dates — 3 endpoint |`.
- [ ] **Step 4:** `examples/transcripts/main.go` — NewClientFromEnv → Latest(0,5) 건수 + Dates("AAPL") 가용 일자 출력.
- [ ] **Step 5:** `integration_test.go` — `TestIntegration_EarningsTranscripts`: Latest(0,10) len>0 / Dates("AAPL") len>0 / Transcript("AAPL","2023","1",0) err 체크.
- [ ] **Step 6:** `unset GOROOT && go build ./... && go vet ./... && go test ./... && go build -tags integration ./... && gofmt -l .` 통과. 커밋 `feat(transcripts): 루트 와이어 + README + examples + 통합테스트`.

## Self-Review 메모
- 3 endpoint, 3 struct. year(int)/fiscalYear(int)/quarter(int)/period(string) 구분.
- Transcript symbol/year/quarter 가드, Dates symbol 가드.
- available-transcript-symbols 제외(directory 중복).
