# FMP Go SDK — Commitment of Traders 그룹 (v0.23.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/cot-group`
- 토픽: FMP `commitmentOfTraders` 카테고리 3 endpoint. 캠페인 21번째 그룹.

## 결정 사항
- 신규 `cot/` 패키지, internal/fetch. 3 구조체.
- `COTReport`(전체 포지션 리포트, **124필드** — 카탈로그 JSON 에서 직접 도출), `COTAnalysis`(16필드, bool 1개), `COTList`(2필드).
- **COTReport 는 필드 수가 많아 구현 시 `docs/api/commitmentOfTraders/cot-report.md` 의 JSON 예시에서 모든 키를 그대로 매핑**(faithful). 타입 규칙: 포지션/미결제약정/트레이더 수/change-* → int64, pct*/conc* → float64, 텍스트 → string. FMP 오타/접미사 그대로 보존(`changeInNoncommSpeadAll`, `tradersNoncommSpeadOl`, pct·trader "old" 블록은 `Ol` 접미사, 포지션 블록은 `Old`).
- Report/Analysis 는 symbol/from/to **선택 필터**(빈값 제외, 가드 없음). List 는 무파라미터.
- 릴리스 `v0.23.0`.

## 패키지 구조 + endpoint 매핑
| 메서드 | path | query | 반환 |
|---|---|---|---|
| `Report(ctx, symbol, from, to)` | `/stable/commitment-of-traders-report` | symbol,from,to(선택) | `[]COTReport` |
| `Analysis(ctx, symbol, from, to)` | `/stable/commitment-of-traders-analysis` | symbol,from,to(선택) | `[]COTAnalysis` |
| `List(ctx)` | `/stable/commitment-of-traders-list` | — | `[]COTList` |

파일: `cot/client.go`(New + filterParams helper), `cot/cot.go`(3 struct + 3 method).
- `filterParams(symbol, from, to string)` — 빈값 제외 맵.
- Report/Analysis: filterParams. List: nil.

## 루트 Client 와이어
```go
EarningsTranscripts *transcripts.Client
COT                 *cot.Client // 상품선물 COT 리포트
```
`c.COT = cot.New(hc)`. `TestNewClient_HasCOT`.

## 응답 타입

### COTAnalysis (commitment-of-traders-analysis) — 16필드
```go
// COTAnalysis — COT 분석/요약 (commitment-of-traders-analysis).
type COTAnalysis struct {
	Symbol                       string  `json:"symbol"`                       // 심볼
	Date                         string  `json:"date"`                         // 일자
	Name                         string  `json:"name"`                         // 상품명
	Sector                       string  `json:"sector"`                       // 섹터
	Exchange                     string  `json:"exchange"`                     // 거래소
	CurrentLongMarketSituation   float64 `json:"currentLongMarketSituation"`   // 현재 롱 비중
	CurrentShortMarketSituation  float64 `json:"currentShortMarketSituation"`  // 현재 숏 비중
	MarketSituation              string  `json:"marketSituation"`              // 현재 시장 상황
	PreviousLongMarketSituation  float64 `json:"previousLongMarketSituation"`  // 이전 롱 비중
	PreviousShortMarketSituation float64 `json:"previousShortMarketSituation"` // 이전 숏 비중
	PreviousMarketSituation      string  `json:"previousMarketSituation"`      // 이전 시장 상황
	NetPostion                   int64   `json:"netPostion"`                   // 순포지션(FMP 오타 netPostion)
	PreviousNetPosition          int64   `json:"previousNetPosition"`          // 이전 순포지션
	ChangeInNetPosition          float64 `json:"changeInNetPosition"`          // 순포지션 변화
	MarketSentiment              string  `json:"marketSentiment"`              // 시장 심리
	ReversalTrend                bool    `json:"reversalTrend"`                // 반전 추세 여부
}
```

### COTList (commitment-of-traders-list) — 2필드
```go
// COTList — COT 보고 대상 상품 목록 (commitment-of-traders-list).
type COTList struct {
	Symbol string `json:"symbol"` // 심볼
	Name   string `json:"name"`   // 상품명
}
```

### COTReport (commitment-of-traders-report) — 124필드 (카탈로그에서 도출)
구현 시 `docs/api/commitmentOfTraders/cot-report.md` JSON 예시의 모든 키를 struct 로 매핑.
- 식별/텍스트(string): symbol, date, name, sector, marketAndExchangeNames, cftcContractMarketCode, cftcMarketCode, cftcRegionCode, cftcCommodityCode, contractUnits
- 포지션 블록 All/Old/Other(int64): openInterest*, noncommPositionsLong/Short/Spread*, commPositionsLong/Short*, totReptPositionsLong/Short*, nonreptPositionsLong/Short*
- change(int64): changeInOpenInterestAll, changeInNoncommLong/ShortAll, changeInNoncommSpeadAll(오타), changeInCommLong/ShortAll, changeInTotReptLong/ShortAll, changeInNonreptLong/ShortAll
- pct 블록 All/Ol/Other(float64): pctOfOpenInterest*, pctOfOiNoncommLong/Short/Spread*, pctOfOiCommLong/Short*, pctOfOiTotReptLong/Short*, pctOfOiNonreptLong/Short* (※ old 블록 접미사는 `Ol`)
- trader 블록 All/Ol/Other(int64): tradersTot*, tradersNoncommLong/Short/Spread*, tradersCommLong/Short*, tradersTotReptLong/Short* (※ Ol 블록에 오타 `tradersNoncommSpeadOl`)
- concentration 블록 All/Ol/Other(float64): concGrossLe4/8TdrLong/Short*, concNetLe4/8TdrLong/Short*

Go 필드명은 관례대로(예: OpenInterestAll, PctOfOiNoncommLongAll), JSON 태그는 카탈로그 키 그대로(오타/Ol 접미사 보존).

## 시그니처 규칙
- Report/Analysis: `(ctx, symbol, from, to string)` → `fetch.List[T](..., filterParams(symbol, from, to))`.
- List: `(ctx)` → `fetch.List[COTList](..., nil)`.

## 테스트
- fixture 단위: COTReport(대표 필드 — OpenInterestAll int, PctOfOpenInterestAll float, ChangeInNoncommSpeadAll 오타 키, 텍스트 Name), COTAnalysis(NetPostion 오타 키, ReversalTrend bool, MarketSentiment), COTList(Symbol/Name).
  - COTReport fixture 는 카탈로그 예시를 그대로 사용(전체 키) — 파싱 후 대표 6~8개 필드 비0/존재 검증.
- delegation: Report("ES",from,to) path+symbol/from/to / Analysis("ES",from,to) path+symbol / List path.
- 통합: List len>0 / Report("",from,to) 또는 Report("ES","","") err 체크 / Analysis("ES","","") err 체크.

## 문서 / 릴리스
- README Commitment of Traders 행(3 endpoint).
- `examples/cot/main.go` — List + Analysis.
- 릴리스 `v0.23.0`.

## 범위 밖 / 위험
- COTReport 124필드 — 카탈로그에서 정확히 도출(오타/접미사 보존). 누락 키 없도록 카탈로그 JSON 전수 매핑.
- symbol/from/to 선택 필터(가드 없음).
- 다음 그룹: etfAndMutualFunds / form13F / Fundraisers / crypto·forex·commodity / secFilings.
