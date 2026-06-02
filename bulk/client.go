// Package bulk 는 FMP 대량 CSV export API sub-client. 모든 메서드는 원시 CSV 바이트를 반환한다.
// fmp.Client.Bulk 로 접근.
package bulk

import (
	"context"
	"fmt"
	"strings"

	"github.com/kenshin579/fmp-go/internal/httpclient"
)

// Client 는 bulk sub-client.
type Client struct {
	http *httpclient.Client
}

// New 는 internal 용도 — root fmp.NewClient 가 호출한다.
func New(http *httpclient.Client) *Client { return &Client{http: http} }

func (c *Client) yearPeriod(ctx context.Context, path, year, period string) ([]byte, error) {
	if strings.TrimSpace(year) == "" || strings.TrimSpace(period) == "" {
		return nil, fmt.Errorf("fmp: year, period must not be empty")
	}
	return c.http.GetRaw(ctx, path, map[string]string{"year": year, "period": period})
}
