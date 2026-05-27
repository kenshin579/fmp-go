package fmp

import "github.com/kenshin579/fmp-go/internal/httpclient"

// APIError 는 FMP 에러 응답. errors.As 로 StatusCode/Message 접근.
type APIError = httpclient.APIError

// ErrNotFound 는 조회 결과가 없을 때(빈 배열 등) 서비스 계층이 반환한다.
var ErrNotFound = httpclient.ErrNotFound
