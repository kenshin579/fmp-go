# FMP Go SDK — 기반 + Company Profile Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Financial Modeling Prep(FMP) API의 Go 클라이언트 라이브러리 기반(`fmp.Client` + `internal/httpclient` + auth/errors)과 첫 카테고리(Company Profile)를 구현해 `v0.1.0` 릴리스 준비.

**Architecture:** opendart 패턴 차용 — 단일 진입점 `fmp.Client`가 도메인별 서비스 서브클라이언트(`company.Client`)를 필드로 노출. `internal/httpclient`가 FMP stable 엔드포인트(`https://financialmodelingprep.com/stable/...`)로 GET, `apikey` 쿼리 자동 주입, JSON 디코딩, 에러 매핑을 담당. FMP는 status envelope 없이 순수 JSON 배열/객체를 반환하고 에러는 HTTP 상태코드 + `{"Error Message": "..."}` 바디로 옴.

**Tech Stack:** Go 1.25+, 표준 라이브러리(`net/http`, `encoding/json`, `net/http/httptest`), git 태그 릴리스. 외부 의존성 없음.

**Spec:** `docs/superpowers/specs/2026-05-26-fmp-go-sdk-foundation-design.md`

**선행:** 서브프로젝트 A0(FMP API docs 카탈로그)가 `docs/api/company/profile-symbol.md`를 먼저 생성. Task 2는 그 카탈로그 + 실호출 fixture를 참조한다.

**Repo:** `github.com/kenshin579/fmp-go` — 워크스페이스 `/Users/frankoh/src/workspace_moneyflow/fmp-go`, branch `feature/sdk-foundation`.

**참고 템플릿:** `/Users/frankoh/src/workspace_moneyflow/opendart` (같은 저자의 동일 패턴 SDK).

---

## File Structure

- Create: `go.mod` — `module github.com/kenshin579/fmp-go`, go 1.25.
- Create: `internal/httpclient/client.go` — HTTP 계층: baseURL, apikey 주입, GET, JSON 디코딩, 에러 매핑.
- Create: `internal/httpclient/client_test.go` — httptest 스텁 단위테스트.
- Create: `errors.go` — `APIError`, `ErrNotFound` (httpclient 타입 alias/재노출).
- Create: `config.go` — functional Option (baseURL/timeout/httpClient).
- Create: `client.go` — `fmp.Client` 진입점 + `NewClient`.
- Create: `from_env.go` — `NewClientFromEnv()` (FMP_API_KEY).
- Create: `client_test.go` — NewClient/NewClientFromEnv 단위테스트.
- Create: `company/client.go` — `company.Client` 서브클라이언트.
- Create: `company/profile.go` — `Profile` 구조체 + `Profile(ctx, symbol)` 메서드.
- Create: `company/profile_test.go` — fixture 기반 테스트.
- Create: `company/testdata/profile_aapl.json` — 프로필 응답 fixture.
- Create: `examples/company/main.go` — 사용 예시.
- Create: `integration_test.go` — build tag `integration`, FMP_API_KEY 있을 때만.
- Create: `scripts/release.sh` — 릴리스 자동화(opendart에서 복사·적응).
- Modify: `README.md` — 설치/사용/커버리지.

---

## Task 1: 모듈 초기화

**Files:**
- Create: `go.mod`

- [ ] **Step 1: go module 초기화**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
go mod init github.com/kenshin579/fmp-go
```

생성된 `go.mod`:
```
module github.com/kenshin579/fmp-go

go 1.25
```

- [ ] **Step 2: 빌드 확인**

Run: `go build ./...`
Expected: 에러 없음(소스가 아직 없으므로 no-op 성공).

- [ ] **Step 3: Commit**

```bash
git add go.mod
git commit -m "chore: go module 초기화 (github.com/kenshin579/fmp-go)"
```

---

## Task 2: Company Profile fixture 확보

**Files:**
- Create: `company/testdata/profile_aapl.json`

A0 카탈로그가 `docs/api/company/profile-symbol.md`(엔드포인트·파라미터·응답 예시)를 이미 생성했다. 본 태스크는 그 문서의 응답 예시(또는 실호출)로 fixture를 만든다. TDD 아님(준비).

- [ ] **Step 1: fixture 확보**

`FMP_API_KEY`가 설정돼 있으면 실제 응답을 저장(가장 정확):
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
mkdir -p company/testdata
if [ -n "$FMP_API_KEY" ]; then
  curl -sS "https://financialmodelingprep.com/stable/profile?symbol=AAPL&apikey=$FMP_API_KEY" \
    | python3 -m json.tool > company/testdata/profile_aapl.json
fi
```
키가 없으면 `docs/api/company/profile-symbol.md`의 응답 예시 JSON을 정리해 같은 파일로 만든다. **응답이 배열(`[ { ... } ]`)인지 단일 객체인지 확인**하고 결과를 보고한다(설계는 배열 가정). fixture는 배열 형태로 저장한다.

fixture에는 최소한 다음 필드가 채워져 있어야 한다(moneyflow 소비 필드): `symbol, companyName, ceo, ipoDate, website, description, sector, industry, country, exchange, image, currency`.

- [ ] **Step 2: Commit**

```bash
git add company/testdata/profile_aapl.json
git commit -m "test(company): AAPL 프로필 응답 fixture"
```

---

## Task 3: internal/httpclient (HTTP 계층)

**Files:**
- Create: `internal/httpclient/client.go`
- Create: `internal/httpclient/client_test.go`

FMP는 status envelope가 없다. GET → 본문 읽기 → 비-200이면 `APIError`, 200이지만 본문이 `{"Error Message": "..."}`면 `APIError`, 그 외엔 `out`으로 디코딩.

- [ ] **Step 1: 실패하는 테스트 작성**

Create `internal/httpclient/client_test.go`:
```go
package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON_InjectsAPIKeyAndDecodes(t *testing.T) {
	var gotKey, gotSymbol, gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKey = r.URL.Query().Get("apikey")
		gotSymbol = r.URL.Query().Get("symbol")
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"symbol":"AAPL","companyName":"Apple Inc."}]`))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "k123", BaseURL: srv.URL})
	var out []map[string]any
	err := c.GetJSON(context.Background(), "/stable/profile", map[string]string{"symbol": "AAPL"}, &out)
	if err != nil {
		t.Fatalf("GetJSON: %v", err)
	}
	if gotKey != "k123" {
		t.Errorf("apikey = %q, want k123", gotKey)
	}
	if gotSymbol != "AAPL" {
		t.Errorf("symbol = %q, want AAPL", gotSymbol)
	}
	if gotPath != "/stable/profile" {
		t.Errorf("path = %q", gotPath)
	}
	if len(out) != 1 || out[0]["symbol"] != "AAPL" {
		t.Errorf("decoded = %v", out)
	}
}

func TestGetJSON_HTTPErrorStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"Error Message":"Invalid API KEY."}`))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "bad", BaseURL: srv.URL})
	var out []map[string]any
	err := c.GetJSON(context.Background(), "/stable/profile", nil, &out)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("want *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 401 || apiErr.Message != "Invalid API KEY." {
		t.Errorf("APIError = %+v", apiErr)
	}
}

func TestGetJSON_ErrorBodyWith200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Error Message":"Limit Reach."}`))
	}))
	defer srv.Close()

	c := New(Config{APIKey: "k", BaseURL: srv.URL})
	var out []map[string]any
	err := c.GetJSON(context.Background(), "/stable/profile", nil, &out)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("want *APIError, got %T: %v", err, err)
	}
	if apiErr.Message != "Limit Reach." {
		t.Errorf("message = %q", apiErr.Message)
	}
}
```

- [ ] **Step 2: 테스트 실행 — 실패 확인**

Run: `cd /Users/frankoh/src/workspace_moneyflow/fmp-go && go test ./internal/httpclient/`
Expected: 컴파일 실패(`New`, `Config`, `APIError`, `GetJSON` 미정의).

- [ ] **Step 3: 구현**

Create `internal/httpclient/client.go`:
```go
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

// GetJSON 은 apikey 를 주입해 GET 후 응답을 out 으로 디코딩한다.
// 비-200 또는 "Error Message" 바디는 *APIError 로 매핑한다.
func (c *Client) GetJSON(ctx context.Context, path string, params map[string]string, out any) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("apikey", c.apiKey)
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("fmp: GET %s: %w", path, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("fmp: read %s: %w", path, err)
	}

	if resp.StatusCode != http.StatusOK {
		msg := resp.Status
		var env errorEnvelope
		if json.Unmarshal(body, &env) == nil && env.ErrorMessage != "" {
			msg = env.ErrorMessage
		}
		return &APIError{StatusCode: resp.StatusCode, Message: msg}
	}

	// 200 이지만 에러 envelope 일 수 있음(FMP 관용).
	var env errorEnvelope
	if json.Unmarshal(body, &env) == nil && env.ErrorMessage != "" {
		return &APIError{StatusCode: resp.StatusCode, Message: env.ErrorMessage}
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("fmp: decode %s: %w", path, err)
	}
	return nil
}
```

- [ ] **Step 4: 테스트 통과 확인**

Run: `go test ./internal/httpclient/ -v` → 3개 모두 PASS. 그리고 `go vet ./internal/httpclient/`.

- [ ] **Step 5: Commit**

```bash
git add internal/httpclient/
git commit -m "feat(httpclient): FMP HTTP 계층 — apikey 주입 + JSON 디코딩 + 에러 매핑"
```

---

## Task 4: root client / config / from_env / errors

**Files:**
- Create: `config.go`, `errors.go`, `client.go`, `from_env.go`
- Create: `client_test.go`
- Create: `company/client.go` (컴파일용 골격)

- [ ] **Step 1: 실패하는 테스트 작성**

Create `client_test.go`:
```go
package fmp_test

import (
	"testing"

	fmp "github.com/kenshin579/fmp-go"
)

func TestNewClient_EmptyKey(t *testing.T) {
	_, err := fmp.NewClient("")
	if err == nil {
		t.Fatal("expected error for empty apiKey")
	}
}

func TestNewClient_HasCompany(t *testing.T) {
	c, err := fmp.NewClient("k123")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.Company == nil {
		t.Fatal("Company sub-client is nil")
	}
}

func TestNewClientFromEnv_MissingEnv(t *testing.T) {
	t.Setenv("FMP_API_KEY", "")
	if _, err := fmp.NewClientFromEnv(); err == nil {
		t.Fatal("expected error when FMP_API_KEY unset")
	}
}

func TestNewClientFromEnv_Reads(t *testing.T) {
	t.Setenv("FMP_API_KEY", "envkey")
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	if c.Company == nil {
		t.Fatal("Company nil")
	}
}
```

- [ ] **Step 2: 테스트 실행 — 실패 확인**

Run: `go test . -run TestNewClient` → 컴파일 실패(`fmp.NewClient` 등 미정의).

- [ ] **Step 3: company 패키지 최소 골격**

Create `company/client.go`:
```go
// Package company 는 FMP 회사 정보 API sub-client 다. fmp.Client.Company 로 접근한다.
package company

import "github.com/kenshin579/fmp-go/internal/httpclient"

// Client 는 회사 정보 sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }
```

- [ ] **Step 4: root 파일 구현**

Create `config.go`:
```go
package fmp

import (
	"net/http"
	"time"
)

type clientOptions struct {
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
}

// Option 은 functional option.
type Option func(*clientOptions)

// WithBaseURL 은 API 베이스 URL 을 지정한다(테스트/프록시용).
func WithBaseURL(u string) Option { return func(o *clientOptions) { o.baseURL = u } }

// WithTimeout 은 HTTP 타임아웃을 지정한다(기본 30s).
func WithTimeout(d time.Duration) Option { return func(o *clientOptions) { o.timeout = d } }

// WithHTTPClient 는 사용자 정의 *http.Client 를 주입한다.
func WithHTTPClient(c *http.Client) Option { return func(o *clientOptions) { o.httpClient = c } }
```

Create `errors.go`:
```go
package fmp

import "github.com/kenshin579/fmp-go/internal/httpclient"

// APIError 는 FMP 에러 응답. errors.As 로 StatusCode/Message 접근.
type APIError = httpclient.APIError

// ErrNotFound 는 조회 결과가 없을 때(빈 배열 등) 서비스 계층이 반환한다.
var ErrNotFound = httpclient.ErrNotFound
```

Create `client.go`:
```go
// Package fmp 는 Financial Modeling Prep(FMP) API 의 Go 클라이언트다.
package fmp

import (
	"errors"
	"time"

	"github.com/kenshin579/fmp-go/company"
	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 fmp-go 라이브러리의 단일 진입점.
type Client struct {
	http *httpclient.Client

	Company *company.Client // 회사 정보(프로필 등)
}

// NewClient 는 API 키로 Client 를 만든다.
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("fmp: apiKey is required")
	}
	cfg := clientOptions{timeout: 30 * time.Second}
	for _, opt := range opts {
		opt(&cfg)
	}
	hc := httpclient.New(httpclient.Config{
		APIKey:     apiKey,
		BaseURL:    cfg.baseURL,
		Timeout:    cfg.timeout,
		HTTPClient: cfg.httpClient,
	})
	c := &Client{http: hc}
	c.Company = company.New(hc)
	return c, nil
}
```

Create `from_env.go`:
```go
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
```

- [ ] **Step 5: 테스트 통과 확인**

Run: `go build ./... && go test . -run TestNewClient -v && go vet ./...`
Expected: PASS.

- [ ] **Step 6: Commit**

```bash
git add client.go config.go from_env.go errors.go client_test.go company/client.go
git commit -m "feat: fmp.Client 진입점 + config/from_env/errors + company sub-client 골격"
```

---

## Task 5: company.Profile

**Files:**
- Create: `company/profile.go`
- Create: `company/profile_test.go`
- (uses `company/testdata/profile_aapl.json` from Task 2)

- [ ] **Step 1: 실패하는 테스트 작성**

Create `company/profile_test.go`:
```go
package company

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

func newTestClient(t *testing.T, status int, body string) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	c := New(httpclient.New(httpclient.Config{APIKey: "k", BaseURL: srv.URL}))
	return c, srv.Close
}

func TestProfile_ParsesFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/profile_aapl.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil {
		t.Fatalf("fixture is not a JSON array: %v", err)
	}
	if len(arr) == 0 {
		t.Fatal("fixture array empty")
	}

	c, cleanup := newTestClient(t, http.StatusOK, string(raw))
	defer cleanup()

	p, err := c.Profile(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("Profile: %v", err)
	}
	if p.Symbol != "AAPL" {
		t.Errorf("Symbol = %q, want AAPL", p.Symbol)
	}
	if p.CompanyName == "" {
		t.Error("CompanyName empty")
	}
	if p.CEO == "" {
		t.Error("CEO empty")
	}
	if p.Website == "" {
		t.Error("Website empty")
	}
}

func TestProfile_EmptyArrayIsNotFound(t *testing.T) {
	c, cleanup := newTestClient(t, http.StatusOK, `[]`)
	defer cleanup()

	_, err := c.Profile(context.Background(), "NOPE")
	if !errors.Is(err, httpclient.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
```

- [ ] **Step 2: 테스트 실행 — 실패 확인**

Run: `go test ./company/ -run TestProfile` → 컴파일 실패(`Profile` 메서드/`Profile` 타입 미정의).

- [ ] **Step 3: 구현**

Create `company/profile.go`:
```go
package company

import (
	"context"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Profile 은 FMP /stable/profile 응답 한 종목.
type Profile struct {
	Symbol            string  `json:"symbol"`
	CompanyName       string  `json:"companyName"`
	Price             float64 `json:"price"`
	MarketCap         int64   `json:"marketCap"`
	Currency          string  `json:"currency"`
	CIK               string  `json:"cik"`
	ISIN              string  `json:"isin"`
	CUSIP             string  `json:"cusip"`
	Exchange          string  `json:"exchange"`
	ExchangeFullName  string  `json:"exchangeFullName"`
	Industry          string  `json:"industry"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	CEO               string  `json:"ceo"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	Image             string  `json:"image"`
	IPODate           string  `json:"ipoDate"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsAdr             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
}

// Profile 은 종목의 회사 프로필을 조회한다. 결과가 없으면 httpclient.ErrNotFound.
func (c *Client) Profile(ctx context.Context, symbol string) (*Profile, error) {
	var out []Profile
	if err := c.http.GetJSON(ctx, "/stable/profile", map[string]string{"symbol": symbol}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, httpclient.ErrNotFound
	}
	return &out[0], nil
}
```

> Task 2에서 응답이 **단일 객체**로 확인됐다면, `var out []Profile` 대신 `var out Profile`로 디코딩하고 빈 Symbol을 ErrNotFound로 판정하도록 조정한다. fixture도 그에 맞춘다. (설계 기본 가정은 배열.)
> A0 카탈로그의 `profile-symbol.md` 실제 필드와 위 struct 필드를 대조해 누락 필드를 추가한다(JSON 태그는 FMP 키와 정확히 일치해야 함).

- [ ] **Step 4: 테스트 통과 확인**

Run: `go test ./company/ -v && go vet ./company/`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add company/profile.go company/profile_test.go
git commit -m "feat(company): Profile 엔드포인트 — /stable/profile 조회 + 매핑"
```

---

## Task 6: examples / README / release.sh / integration test

**Files:**
- Create: `examples/company/main.go`
- Create: `integration_test.go`
- Create: `scripts/release.sh`
- Modify: `README.md`

- [ ] **Step 1: 사용 예시**

Create `examples/company/main.go`:
```go
// examples/company — FMP 회사 프로필 조회 예제.
//
// 실행: FMP_API_KEY=... go run ./examples/company AAPL
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	fmp "github.com/kenshin579/fmp-go"
)

func main() {
	symbol := "AAPL"
	if len(os.Args) > 1 {
		symbol = os.Args[1]
	}
	client, err := fmp.NewClientFromEnv()
	if err != nil {
		log.Fatalf("NewClientFromEnv: %v", err)
	}
	p, err := client.Company.Profile(context.Background(), symbol)
	if err != nil {
		log.Fatalf("Profile(%s): %v", symbol, err)
	}
	fmt.Printf("%s (%s)\nCEO: %s\nIPO: %s\nWeb: %s\nSector: %s / %s\n%s\n",
		p.CompanyName, p.Symbol, p.CEO, p.IPODate, p.Website, p.Sector, p.Industry, p.Description)
}
```

- [ ] **Step 2: 통합 테스트(build tag)**

Create `integration_test.go`:
```go
//go:build integration

package fmp_test

import (
	"context"
	"os"
	"testing"

	fmp "github.com/kenshin579/fmp-go"
)

// 실행: FMP_API_KEY=... go test -tags integration -run TestIntegration ./...
func TestIntegration_CompanyProfile(t *testing.T) {
	if os.Getenv("FMP_API_KEY") == "" {
		t.Skip("FMP_API_KEY not set")
	}
	c, err := fmp.NewClientFromEnv()
	if err != nil {
		t.Fatalf("NewClientFromEnv: %v", err)
	}
	p, err := c.Company.Profile(context.Background(), "AAPL")
	if err != nil {
		t.Fatalf("Profile: %v", err)
	}
	if p.Symbol != "AAPL" || p.CompanyName == "" {
		t.Errorf("unexpected profile: %+v", p)
	}
}
```

- [ ] **Step 3: release.sh (opendart에서 복사·적응)**

Run:
```bash
mkdir -p scripts
cp /Users/frankoh/src/workspace_moneyflow/opendart/scripts/release.sh scripts/release.sh
```
그런 다음 `scripts/release.sh` 안의 문자열 `opendart`를 `fmp-go`로 치환(주석/echo 메시지만 — 로직은 프로젝트 무관). 실행 권한 확인:
```bash
chmod +x scripts/release.sh
bash -n scripts/release.sh   # 문법 검사
```
Expected: 문법 에러 없음.

- [ ] **Step 4: README 작성**

Overwrite `README.md`:
```markdown
# fmp-go

Financial Modeling Prep(FMP) API 의 Go 클라이언트 라이브러리.

## 설치

​```bash
go get github.com/kenshin579/fmp-go@v0.1.0
​```

## 사용

​```go
client, _ := fmp.NewClientFromEnv() // FMP_API_KEY
ctx := context.Background()

profile, _ := client.Company.Profile(ctx, "AAPL") // 회사 프로필
fmt.Println(profile.CompanyName, profile.CEO, profile.Website)
​```

## 인증

발급받은 API 키를 `FMP_API_KEY` 환경변수로 두거나 `fmp.NewClient(apiKey)` 로 전달한다.
모든 요청에 `apikey` 쿼리로 자동 주입된다. FMP stable 엔드포인트
(`https://financialmodelingprep.com/stable/...`)를 사용한다.

## 커버리지

| 카테고리 | 서비스 | 엔드포인트 |
|----------|--------|-----------|
| Company | `client.Company` | Profile (`/stable/profile`) |

> 전체 FMP API 커버리지를 목표로 카테고리 단위로 점진 확장한다.
> 전체 API 문서 카탈로그: `docs/api/`.

## 개발

​```bash
go build ./...
go vet ./...
go test ./...                          # 단위 테스트
go test -tags integration ./...        # 통합(FMP_API_KEY 필요)
​```
```

- [ ] **Step 5: 빌드/검증**

Run:
```bash
go build ./... && go vet ./... && go test ./... && go build ./examples/...
```
Expected: 모두 PASS, examples 컴파일 성공.

- [ ] **Step 6: Commit**

```bash
git add examples/ integration_test.go scripts/release.sh README.md
git commit -m "docs: README + 사용 예제 + 통합 테스트 + release.sh"
```

---

## Task 7: 전체 검증

**Files:** 없음(검증 전용)

- [ ] **Step 1: 전체 빌드/vet/test**

Run:
```bash
cd /Users/frankoh/src/workspace_moneyflow/fmp-go
go build ./... && go vet ./... && go test ./...
```
Expected: 전부 통과.

- [ ] **Step 2: 통합 테스트(가능 시)**

`FMP_API_KEY`가 있으면:
```bash
go test -tags integration -run TestIntegration ./...
```
Expected: PASS(또는 키 없으면 Skip). 실패 시 실제 응답 형태(배열/객체, 필드명)를 fixture·struct와 대조해 Task 5에서 조정.

- [ ] **Step 3: 커밋 상태 확인**

Run: `git log --oneline -8 && git status --short`
Expected: Task 1~6 커밋이 `feature/sdk-foundation`에 순서대로, 워킹트리 클린.

---

## 자기 점검 메모 (작성자용)

- **응답 배열 vs 객체**: 설계 기본은 배열. Task 2에서 실측 확인하고 Task 5에서 필요 시 조정(계획에 분기 명시).
- **Profile 필드**: Task 5 struct는 알려진 stable 필드 기준. A0 카탈로그 + Task 2 fixture와 대조해 누락 필드 추가(JSON 태그는 FMP 키와 정확 일치).
- **에러 의미**: 빈 배열 → `ErrNotFound`(company 계층). HTTP 비-200/`Error Message` → `APIError`(httpclient). 402/429/401은 StatusCode로 식별.
- **릴리스**: v0.1.0 태그는 본 계획 범위 밖(머지 후 별도). 서브프로젝트 B는 그 태그를 소비.
- **외부 의존성 없음**: 표준 라이브러리만. go.sum 불필요.
```
