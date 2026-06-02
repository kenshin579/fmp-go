package bulk

import (
	"context"
	"fmt"
	"strings"
)

func (c *Client) Profile(ctx context.Context, part string) ([]byte, error) {
	if strings.TrimSpace(part) == "" {
		return nil, fmt.Errorf("fmp: part must not be empty")
	}
	return c.http.GetRaw(ctx, "/stable/profile-bulk", map[string]string{"part": part})
}
func (c *Client) ETFHolder(ctx context.Context, part string) ([]byte, error) {
	if strings.TrimSpace(part) == "" {
		return nil, fmt.Errorf("fmp: part must not be empty")
	}
	return c.http.GetRaw(ctx, "/stable/etf-holder-bulk", map[string]string{"part": part})
}
func (c *Client) EOD(ctx context.Context, date string) ([]byte, error) {
	if strings.TrimSpace(date) == "" {
		return nil, fmt.Errorf("fmp: date must not be empty")
	}
	return c.http.GetRaw(ctx, "/stable/eod-bulk", map[string]string{"date": date})
}
func (c *Client) EarningsSurprises(ctx context.Context, year string) ([]byte, error) {
	if strings.TrimSpace(year) == "" {
		return nil, fmt.Errorf("fmp: year must not be empty")
	}
	return c.http.GetRaw(ctx, "/stable/earnings-surprises-bulk", map[string]string{"year": year})
}
func (c *Client) RatiosTTM(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/ratios-ttm-bulk", nil)
}
func (c *Client) KeyMetricsTTM(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/key-metrics-ttm-bulk", nil)
}
func (c *Client) Scores(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/scores-bulk", nil)
}
func (c *Client) DCF(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/dcf-bulk", nil)
}
func (c *Client) Peers(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/peers-bulk", nil)
}
func (c *Client) PriceTargetSummary(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/price-target-summary-bulk", nil)
}
func (c *Client) Rating(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/rating-bulk", nil)
}
func (c *Client) UpgradesDowngradesConsensus(ctx context.Context) ([]byte, error) {
	return c.http.GetRaw(ctx, "/stable/upgrades-downgrades-consensus-bulk", nil)
}
