package statements

import "context"

// IncomeStatementTTM 는 종목의 TTM(최근 12개월) 손익계산서를 조회한다. period 미지원.
func (c *Client) IncomeStatementTTM(ctx context.Context, p Params) ([]IncomeStatement, error) {
	return fetchList[IncomeStatement](ctx, c, "/stable/income-statement-ttm", p, p.ttmQueryParams())
}

// BalanceSheetStatementTTM 는 종목의 TTM 대차대조표를 조회한다. period 미지원.
func (c *Client) BalanceSheetStatementTTM(ctx context.Context, p Params) ([]BalanceSheetStatement, error) {
	return fetchList[BalanceSheetStatement](ctx, c, "/stable/balance-sheet-statement-ttm", p, p.ttmQueryParams())
}

// CashFlowStatementTTM 는 종목의 TTM 현금흐름표를 조회한다. period 미지원.
func (c *Client) CashFlowStatementTTM(ctx context.Context, p Params) ([]CashFlowStatement, error) {
	return fetchList[CashFlowStatement](ctx, c, "/stable/cash-flow-statement-ttm", p, p.ttmQueryParams())
}
