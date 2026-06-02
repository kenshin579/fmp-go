package news

import (
	"context"
	"strconv"

	"github.com/kenshin579/fmp-go/internal/fetch"
)

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

// FMPArticles 는 FMP 자체 작성 기사를 페이지 단위로 조회한다.
func (c *Client) FMPArticles(ctx context.Context, page int) ([]FMPArticle, error) {
	return fetch.List[FMPArticle](ctx, c.http, "/stable/fmp-articles", map[string]string{"page": strconv.Itoa(page)})
}
