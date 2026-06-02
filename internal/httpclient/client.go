// Package httpclient 는 FMP REST 호출의 단일 GET 통로다.
// apikey 쿼리 주입, JSON 디코딩, 에러 매핑(HTTP 상태 + FMP "Error Message" 바디)을 담당한다.
package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// DefaultBaseURL 은 FMP API 베이스 URL.
const DefaultBaseURL = "https://financialmodelingprep.com"

// ErrNotFound 는 조회 결과가 없을 때 서비스 계층이 반환하는 sentinel.
var ErrNotFound = errors.New("fmp: not found")

// APIError 는 FMP 에러 응답(비-200 또는 "Error Message" 바디)을 나타낸다.
type APIError struct {
	StatusCode int    // HTTP 상태코드
	Message    string // FMP "Error Message" 또는 상태 텍스트
}

func (e *APIError) Error() string {
	return fmt.Sprintf("fmp: api error (status %d): %s", e.StatusCode, e.Message)
}

// Config 는 Client 생성 인자.
type Config struct {
	APIKey     string
	BaseURL    string        // 빈 값이면 DefaultBaseURL
	Timeout    time.Duration // 0이면 30s
	HTTPClient *http.Client  // nil이면 Timeout 적용 기본 클라이언트
}

// Client 는 FMP HTTP 계층.
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

// New 는 Config 로 Client 를 만든다.
func New(cfg Config) *Client {
	base := cfg.BaseURL
	if base == "" {
		base = DefaultBaseURL
	}
	hc := cfg.HTTPClient
	if hc == nil {
		timeout := cfg.Timeout
		if timeout == 0 {
			timeout = 30 * time.Second
		}
		hc = &http.Client{Timeout: timeout}
	}
	return &Client{apiKey: cfg.APIKey, baseURL: base, http: hc}
}

// errorEnvelope 는 FMP 에러 바디 형태 {"Error Message": "..."}.
type errorEnvelope struct {
	ErrorMessage string `json:"Error Message"`
}

// get 은 apikey 를 주입해 GET 후 응답 바디를 반환한다.
// 비-200 또는 "Error Message" 바디는 *APIError 로 매핑한다.
func (c *Client) get(ctx context.Context, path string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("apikey", c.apiKey)
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fmp: GET %s: %w", path, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fmp: read %s: %w", path, err)
	}

	if resp.StatusCode != http.StatusOK {
		msg := resp.Status
		var env errorEnvelope
		if json.Unmarshal(body, &env) == nil && env.ErrorMessage != "" {
			msg = env.ErrorMessage
		}
		return nil, &APIError{StatusCode: resp.StatusCode, Message: msg}
	}

	// 200 이지만 에러 envelope 일 수 있음(FMP 관용).
	var env errorEnvelope
	if json.Unmarshal(body, &env) == nil && env.ErrorMessage != "" {
		return nil, &APIError{StatusCode: resp.StatusCode, Message: env.ErrorMessage}
	}

	return body, nil
}

// GetJSON 은 apikey 를 주입해 GET 후 응답을 out 으로 디코딩한다.
// 비-200 또는 "Error Message" 바디는 *APIError 로 매핑한다.
func (c *Client) GetJSON(ctx context.Context, path string, params map[string]string, out any) error {
	body, err := c.get(ctx, path, params)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("fmp: decode %s: %w", path, err)
	}
	return nil
}

// GetRaw 는 apikey 주입 GET 후 응답 바디를 원시 바이트로 반환한다(CSV 등 비-JSON 용).
func (c *Client) GetRaw(ctx context.Context, path string, params map[string]string) ([]byte, error) {
	return c.get(ctx, path, params)
}
