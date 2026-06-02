# FMP Go SDK — News 그룹 (v0.6.0) 설계

- 작성일: 2026-06-02
- 상태: 확정 (브레인스토밍 완료)
- 레포: `github.com/kenshin579/fmp-go`, branch `feature/news-group`
- 토픽: FMP `news` 카테고리 10 endpoint 추가. 전체 API 커버리지 캠페인 4번째 그룹. moneyflow 뉴스 탭 데이터 소스로 직결.

## 배경 / 목적

추천 순서(search → news → analyst → ...)의 두 번째. `news` 는 외부 개발자 가치 + moneyflow 의 placeholder "뉴스" 탭 직결. 10 endpoint 중 9개가 동일 `Article` shape 라 완결성 높음.

## 결정 사항 (브레인스토밍)

- **범위**: news 10 endpoint 전부. 신규 `news/` 패키지, `internal/fetch` 사용.
- **전부 list 반환**. latest(page) 6개 + search(symbols) 4개.
- **struct 재사용**: 9 endpoint → `Article` 공용. fmp-articles 만 `FMPArticle`.
- **네이밍**: latest 는 `XNewsLatest`/`PressReleasesLatest`, search 는 `SearchXNews`/`SearchPressReleases`.
- **템플릿 계승**: 필드 한국어 주석, fixture + delegation 테스트, README/examples.
- **릴리스**: `v0.6.0`.

## 패키지 구조 + endpoint 매핑

| 파일 | 메서드 | path | helper | 반환 |
|---|---|---|---|---|
| `latest.go` | `StockNewsLatest(ctx, page)` | `/stable/news/stock-latest` | List{page} | `[]Article` |
| | `CryptoNewsLatest(ctx, page)` | `/stable/news/crypto-latest` | List{page} | `[]Article` |
| | `ForexNewsLatest(ctx, page)` | `/stable/news/forex-latest` | List{page} | `[]Article` |
| | `GeneralNewsLatest(ctx, page)` | `/stable/news/general-latest` | List{page} | `[]Article` |
| | `PressReleasesLatest(ctx, page)` | `/stable/news/press-releases-latest` | List{page} | `[]Article` |
| `search.go` | `SearchStockNews(ctx, symbols...)` | `/stable/news/stock` | ListBySymbols | `[]Article` |
| | `SearchCryptoNews(ctx, symbols...)` | `/stable/news/crypto` | ListBySymbols | `[]Article` |
| | `SearchForexNews(ctx, symbols...)` | `/stable/news/forex` | ListBySymbols | `[]Article` |
| | `SearchPressReleases(ctx, symbols...)` | `/stable/news/press-releases` | ListBySymbols | `[]Article` |
| `fmp_articles.go` | `FMPArticles(ctx, page)` | `/stable/fmp-articles` | List{page} | `[]FMPArticle` |
| `client.go` | `New(http)` | — | — | `*Client` |

## 응답 타입 (faithful, 필드 한국어 주석)

```go
// Article — 뉴스/보도자료 기사 (stock/crypto/forex/press-releases/general + 각 search 공용)
type Article struct {
	Symbol        string `json:"symbol"`        // 관련 종목 심볼 (general 뉴스는 빈 문자열)
	PublishedDate string `json:"publishedDate"` // 게시 일시 (YYYY-MM-DD HH:MM:SS)
	Publisher     string `json:"publisher"`     // 발행처 (예: Seeking Alpha)
	Title         string `json:"title"`         // 기사 제목
	Image         string `json:"image"`         // 대표 이미지 URL
	Site          string `json:"site"`          // 출처 사이트 도메인
	Text          string `json:"text"`          // 기사 본문 요약
	URL           string `json:"url"`           // 원문 URL
}

// FMPArticle — FMP 자체 작성 기사 (fmp-articles)
type FMPArticle struct {
	Title   string `json:"title"`   // 제목
	Date    string `json:"date"`    // 작성 일시
	Content string `json:"content"` // 본문 (HTML)
	Tickers string `json:"tickers"` // 관련 티커 (예: "NYSE:MRK")
	Image   string `json:"image"`   // 이미지 URL
	Link    string `json:"link"`    // FMP 기사 링크
	Author  string `json:"author"`  // 작성자
	Site    string `json:"site"`    // 출처 (Financial Modeling Prep)
}
```
- 전 필드 string(뉴스 메타 텍스트). `Article.Symbol` 은 general-news `null` → `""` 안전 디코딩. `FMPArticle.Content` 는 HTML 포함.

## 시그니처 규칙
- latest 6개: `(ctx, page int)` → `fetch.List[T](ctx, c.http, path, map[string]string{"page": strconv.Itoa(page)})`.
- search 4개: `(ctx, symbols ...string)` → `fetch.ListBySymbols[Article](ctx, c.http, path, symbols)` (쉼표 join + 빈 가드 내장).

## 루트 Client 와이어
```go
type Client struct {
	...
	Search *search.Client
	News   *news.Client // 뉴스 (신규)
}
```
`NewClient` 에 `c.News = news.New(hc)`. `client_test.go` 에 `TestNewClient_HasNews`.

## 테스트
- fixture 단위: stock-latest(Article) + search-stock(Article) + general(symbol null→"") + fmp-articles(FMPArticle). 동일 shape·helper 는 delegation 으로 커버.
- delegation: `StockNewsLatest(page)` path+page / `SearchStockNews(symbols)` path+symbols join / `FMPArticles(page)` path. 대표 3개.
- 가드: search 빈 symbols 대표 1건.
- 통합(`//go:build integration`): StockNewsLatest(0) / SearchStockNews("AAPL") / FMPArticles(0).

## 문서 / 릴리스
- README 커버리지 표 News 행(10 endpoint).
- `examples/news/main.go` — SearchStockNews + StockNewsLatest.
- 릴리스 `v0.6.0`.

## 범위 밖 / 위험
- 나머지 24 그룹 별도 PR(다음: analyst → calendar → statements 확장).
- general-news `symbol: null` → `""` 디코딩 fixture 검증(string null 안전성).
- latest 의 `limit` 파라미터는 FMP 기본값 사용(page 만 노출) — 후속 필요 시 추가.
- moneyflow 뉴스 탭 통합은 별도 작업(필요 시 `go get ...@v0.6.0`).
