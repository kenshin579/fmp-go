package fmp

import (
	"errors"
	"os"
)

// NewClientFromEnv 는 FMP_API_KEY 환경변수로 Client 를 만든다.
func NewClientFromEnv(opts ...Option) (*Client, error) {
	key := os.Getenv("FMP_API_KEY")
	if key == "" {
		return nil, errors.New("fmp: FMP_API_KEY is not set")
	}
	return NewClient(key, opts...)
}
