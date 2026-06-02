package news

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
