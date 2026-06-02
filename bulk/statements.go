package bulk

import "context"

func (c *Client) IncomeStatement(ctx context.Context, year, period string) ([]byte, error) {
	return c.yearPeriod(ctx, "/stable/income-statement-bulk", year, period)
}
func (c *Client) IncomeStatementGrowth(ctx context.Context, year, period string) ([]byte, error) {
	return c.yearPeriod(ctx, "/stable/income-statement-growth-bulk", year, period)
}
func (c *Client) BalanceSheetStatement(ctx context.Context, year, period string) ([]byte, error) {
	return c.yearPeriod(ctx, "/stable/balance-sheet-statement-bulk", year, period)
}
func (c *Client) BalanceSheetStatementGrowth(ctx context.Context, year, period string) ([]byte, error) {
	return c.yearPeriod(ctx, "/stable/balance-sheet-statement-growth-bulk", year, period)
}
func (c *Client) CashFlowStatement(ctx context.Context, year, period string) ([]byte, error) {
	return c.yearPeriod(ctx, "/stable/cash-flow-statement-bulk", year, period)
}
func (c *Client) CashFlowStatementGrowth(ctx context.Context, year, period string) ([]byte, error) {
	return c.yearPeriod(ctx, "/stable/cash-flow-statement-growth-bulk", year, period)
}
