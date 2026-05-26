# Financial Reports Form 10-K XLSX

Download detailed 10-K reports in XLSX format with the Financial Reports Form 10-K XLSX API. Effortlessly access and analyze annual financial data for companies in a spreadsheet-friendly format.

## Endpoint

`GET https://financialmodelingprep.com/stable/financial-reports-xlsx?symbol=AAPL&year=2022&period=FY`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| year* | number | 2022 |
| period* | string | Q1,Q2,Q3,Q4,FY |

## Description

The Financial Reports Form 10-K XLSX API provides users with the ability to download 10-K financial reports in a format that can be opened in Excel. This allows for:

- Detailed Financial Analysis: View comprehensive financial data, including income statements, balance sheets, and cash flow statements, with Excel's built-in analysis tools.

- Flexible Data Usage: Customize and manipulate the data for further analysis, enabling users to run financial models or track trends.

- Efficient Reporting: Create financial summaries, pivot tables, and other visualizations based on the data from 10-K reports.

- Historical Data Access: Download reports from previous fiscal years for detailed historical comparisons.

This API makes it simple to work with financial data in a spreadsheet, streamlining analysis and reporting workflows.

Example Use Case
A financial analyst can download Apple's 2022 10-K report in XLSX format, making it easier to import the data into their financial models and analyze trends over the fiscal year.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"period": "FY",
		"year": "2022",
		"Cover Page": [
			{
				"Cover Page - USD ($) shares in Thousands, $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Oct. 14, 2022",
					"Mar. 25, 2022"
				]
			},
			{
				"Entity Information [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Auditor Information": [
			{
				"Auditor Information": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Auditor Information [Abstract]": [
					" "
				]
			}
		],
		"CONSOLIDATED STATEMENTS OF OPER": [
			{
				"CONSOLIDATED STATEMENTS OF OPERATIONS - USD ($) shares in Thousands, $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Net sales": [
					394328,
					365817,
					274515
				]
			}
		],
		"CONSOLIDATED STATEMENTS OF COMP": [
			{
				"CONSOLIDATED STATEMENTS OF COMPREHENSIVE INCOME - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Statement of Comprehensive Income [Abstract]": [
					" ",
					" ",
					" "
				]
			}
		],
		"CONSOLIDATED BALANCE SHEETS": [
			{
				"CONSOLIDATED BALANCE SHEETS - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Current assets:": [
					" ",
					" "
				]
			},
			{
				"Cash and cash equivalents": [
					23646,
					34940
				]
			}
		],
		"CONSOLIDATED BALANCE SHEETS (Pa": [
			{
				"CONSOLIDATED BALANCE SHEETS (Parenthetical) - $ / shares": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Statement of Financial Position [Abstract]": [
					" ",
					" "
				]
			},
			{
				"Common stock, par value (in dollars per share)": [
					0.00001,
					0.00001
				]
			}
		],
		"CONSOLIDATED STATEMENTS OF SHAR": [
			{
				"CONSOLIDATED STATEMENTS OF SHAREHOLDERS' EQUITY - USD ($) $ in Millions": [
					"Total",
					"Common stock and additional paid-in capital",
					"Retained earnings/(Accumulated deficit)",
					"Retained earnings/(Accumulated deficit) Cumulative effect of change in accounting principle",
					"Accumulated other comprehensive income/(loss)",
					"Accumulated other comprehensive income/(loss) Cumulative effect of change in accounting principle"
				]
			},
			{
				"Beginning balances at Sep. 28, 2019": [
					90488,
					45174,
					45898,
					-136,
					-584,
					136
				]
			},
			{
				"Increase (Decrease) in Stockholders' Equity [Roll Forward]": [
					" ",
					" ",
					" ",
					" ",
					" ",
					" "
				]
			}
		],
		"CONSOLIDATED STATEMENTS OF CASH": [
			{
				"CONSOLIDATED STATEMENTS OF CASH FLOWS - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Statement of Cash Flows [Abstract]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Summary of Significant Accounti": [
			{
				"Summary of Significant Accounting Policies": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Accounting Policies [Abstract]": [
					" "
				]
			}
		],
		"Revenue": [
			{
				"Revenue": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Revenue from Contract with Customer [Abstract]": [
					" "
				]
			}
		],
		"Financial Instruments": [
			{
				"Financial Instruments": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Investments, All Other Investments [Abstract]": [
					" "
				]
			}
		],
		"Consolidated Financial Statemen": [
			{
				"Consolidated Financial Statement Details": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Organization, Consolidation and Presentation of Financial Statements [Abstract]": [
					" "
				]
			}
		],
		"Income Taxes": [
			{
				"Income Taxes": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Income Tax Disclosure [Abstract]": [
					" "
				]
			}
		],
		"Leases": [
			{
				"Leases": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Leases [Abstract]": [
					" "
				]
			}
		],
		"Debt": [
			{
				"Debt": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Debt Disclosure [Abstract]": [
					" "
				]
			}
		],
		"Shareholders' Equity": [
			{
				"Shareholders' Equity": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Equity [Abstract]": [
					" "
				]
			}
		],
		"Benefit Plans": [
			{
				"Benefit Plans": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Share-Based Payment Arrangement [Abstract]": [
					" "
				]
			}
		],
		"Commitments and Contingencies": [
			{
				"Commitments and Contingencies": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Commitments and Contingencies Disclosure [Abstract]": [
					" "
				]
			}
		],
		"Segment Information and Geograp": [
			{
				"Segment Information and Geographic Data": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Segment Reporting [Abstract]": [
					" "
				]
			}
		],
		"Summary of Significant Accoun_2": [
			{
				"Summary of Significant Accounting Policies (Policies)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Accounting Policies [Abstract]": [
					" "
				]
			}
		],
		"Summary of Significant Accoun_3": [
			{
				"Summary of Significant Accounting Policies (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Accounting Policies [Abstract]": [
					" "
				]
			}
		],
		"Revenue (Tables)": [
			{
				"Revenue (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Revenue from Contract with Customer [Abstract]": [
					" "
				]
			}
		],
		"Financial Instruments (Tables)": [
			{
				"Financial Instruments (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Investments, All Other Investments [Abstract]": [
					" "
				]
			}
		],
		"Consolidated Financial Statem_2": [
			{
				"Consolidated Financial Statement Details (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Organization, Consolidation and Presentation of Financial Statements [Abstract]": [
					" "
				]
			}
		],
		"Income Taxes (Tables)": [
			{
				"Income Taxes (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Income Tax Disclosure [Abstract]": [
					" "
				]
			}
		],
		"Leases (Tables)": [
			{
				"Leases (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Leases [Abstract]": [
					" "
				]
			}
		],
		"Debt (Tables)": [
			{
				"Debt (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Debt Disclosure [Abstract]": [
					" "
				]
			}
		],
		"Shareholders' Equity (Tables)": [
			{
				"Shareholders' Equity (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Equity [Abstract]": [
					" "
				]
			}
		],
		"Benefit Plans (Tables)": [
			{
				"Benefit Plans (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Share-Based Payment Arrangement [Abstract]": [
					" "
				]
			}
		],
		"Commitments and Contingencies (": [
			{
				"Commitments and Contingencies (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Commitments and Contingencies Disclosure [Abstract]": [
					" "
				]
			}
		],
		"Segment Information and Geogr_2": [
			{
				"Segment Information and Geographic Data (Tables)": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022"
				]
			},
			{
				"Segment Reporting [Abstract]": [
					" "
				]
			}
		],
		"Summary of Significant Accoun_4": [
			{
				"Summary of Significant Accounting Policies - Additional Information (Details) $ in Billions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022 USD ($) performanceObligation",
					"Sep. 25, 2021 USD ($)",
					"Sep. 26, 2020 USD ($)"
				]
			},
			{
				"Significant Accounting Policies [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Summary of Significant Accoun_5": [
			{
				"Summary of Significant Accounting Policies - Computation of Basic and Diluted Earnings Per Share (Details) - USD ($) $ / shares in Units, shares in Thousands, $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Numerator:": [
					" ",
					" ",
					" "
				]
			}
		],
		"Revenue - Net Sales Disaggregat": [
			{
				"Revenue - Net Sales Disaggregated by Significant Products and Services (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Disaggregation of Revenue [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Revenue - Additional Informatio": [
			{
				"Revenue - Additional Information (Details) - USD ($) $ in Billions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Revenue from Contract with Customer [Abstract]": [
					" ",
					" "
				]
			},
			{
				"Total deferred revenue": [
					12.4,
					11.9
				]
			}
		],
		"Revenue - Deferred Revenue, Exp": [
			{
				"Revenue - Deferred Revenue, Expected Timing of Realization (Details)": [
					"Sep. 24, 2022"
				]
			},
			{
				"Revenue, Remaining Performance Obligation, Expected Timing of Satisfaction, Start Date [Axis]: 2022-09-25": [
					" "
				]
			},
			{
				"Revenue, Remaining Performance Obligation, Expected Timing of Satisfaction [Line Items]": [
					" "
				]
			}
		],
		"Financial Instruments - Cash, C": [
			{
				"Financial Instruments - Cash, Cash Equivalents and Marketable Securities (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Debt Securities, Available-for-sale [Line Items]": [
					" ",
					" "
				]
			},
			{
				"Cash, Cash Equivalents and Marketable Securities, Adjusted Cost": [
					183061,
					189961
				]
			}
		],
		"Financial Instruments - Non-Cur": [
			{
				"Financial Instruments - Non-Current Marketable Debt Securities by Contractual Maturity (Details) $ in Millions": [
					"Sep. 24, 2022 USD ($)"
				]
			},
			{
				"Fair value of non-current marketable debt securities by contractual maturity": [
					" "
				]
			},
			{
				"Due after 1 year through 5 years": [
					87031
				]
			}
		],
		"Financial Instruments - Additio": [
			{
				"Financial Instruments - Additional Information (Details) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022 USD ($) Customer Vendor",
					"Sep. 25, 2021 Vendor"
				]
			},
			{
				"Financial Instruments [Line Items]": [
					" ",
					" "
				]
			}
		],
		"Financial Instruments - Notiona": [
			{
				"Financial Instruments - Notional Amounts Associated with Derivative Instruments (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Derivatives designated as accounting hedges | Foreign exchange contracts": [
					" ",
					" "
				]
			},
			{
				"Derivative [Line Items]": [
					" ",
					" "
				]
			}
		],
		"Financial Instruments - Gross F": [
			{
				"Financial Instruments - Gross Fair Values of Derivative Assets and Liabilities (Details) - Level 2 $ in Millions": [
					"Sep. 24, 2022 USD ($)"
				]
			},
			{
				"Other current assets and other non-current assets | Foreign exchange contracts": [
					" "
				]
			},
			{
				"Derivative assets:": [
					" "
				]
			}
		],
		"Financial Instruments - Derivat": [
			{
				"Financial Instruments - Derivative Instruments Designated as Fair Value Hedges and Related Hedged Items (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Current and non-current marketable securities": [
					" ",
					" "
				]
			},
			{
				"Derivatives, Fair Value [Line Items]": [
					" ",
					" "
				]
			}
		],
		"Consolidated Financial Statem_3": [
			{
				"Consolidated Financial Statement Details - Property, Plant and Equipment, Net (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Property, Plant and Equipment [Line Items]": [
					" ",
					" "
				]
			},
			{
				"Gross property, plant and equipment": [
					114457,
					109723
				]
			}
		],
		"Consolidated Financial Statem_4": [
			{
				"Consolidated Financial Statement Details - Other Non-Current Liabilities (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Organization, Consolidation and Presentation of Financial Statements [Abstract]": [
					" ",
					" "
				]
			},
			{
				"Long-term taxes payable": [
					16657,
					24689
				]
			}
		],
		"Consolidated Financial Statem_5": [
			{
				"Consolidated Financial Statement Details - Other Income/(Expense), Net (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Organization, Consolidation and Presentation of Financial Statements [Abstract]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Income Taxes - Provision for In": [
			{
				"Income Taxes - Provision for Income Taxes (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Federal:": [
					" ",
					" ",
					" "
				]
			}
		],
		"Income Taxes - Additional Infor": [
			{
				"Income Taxes - Additional Information (Details) $ in Millions, € in Billions": [
					null,
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Aug. 30, 2016 EUR (€) Subsidiary",
					"Sep. 24, 2022 USD ($)",
					"Sep. 25, 2021 USD ($)",
					"Sep. 26, 2020 USD ($)",
					"Sep. 24, 2022 EUR (€)",
					"Sep. 28, 2019 USD ($)"
				]
			},
			{
				"Income Tax Contingency [Line Items]": [
					" ",
					" ",
					" ",
					" ",
					" ",
					" "
				]
			}
		],
		"Income Taxes - Reconciliation o": [
			{
				"Income Taxes - Reconciliation of Provision for Income Taxes to Amount Computed by Applying the Statutory Federal Income Tax Rate to Income Before Provision for Income Taxes (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Income Tax Disclosure [Abstract]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Income Taxes - Significant Comp": [
			{
				"Income Taxes - Significant Components of Deferred Tax Assets and Liabilities (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Deferred tax assets:": [
					" ",
					" "
				]
			},
			{
				"Amortization and depreciation": [
					1496,
					5575
				]
			}
		],
		"Income Taxes - Aggregate Change": [
			{
				"Income Taxes - Aggregate Changes in Gross Unrecognized Tax Benefits (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Reconciliation of Unrecognized Tax Benefits, Excluding Amounts Pertaining to Examined Tax Returns [Roll Forward]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Leases - Additional Information": [
			{
				"Leases - Additional Information (Details) - USD ($) $ in Billions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Lessee, Lease, Description [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Leases - ROU Assets and Lease L": [
			{
				"Leases - ROU Assets and Lease Liabilities (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Lease-Related Assets and Liabilities": [
					" ",
					" "
				]
			},
			{
				"Operating lease right-of-use assets": [
					10417,
					10087
				]
			}
		],
		"Leases - Lease Liability Maturi": [
			{
				"Leases - Lease Liability Maturities (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Operating Leases": [
					" ",
					" "
				]
			},
			{
				"2023": [
					1758,
					" "
				]
			}
		],
		"Debt - Additional Information (": [
			{
				"Debt - Additional Information (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Debt Instrument [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Debt - Summary of Cash Flows As": [
			{
				"Debt - Summary of Cash Flows Associated with Commercial Paper (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Maturities 90 days or less:": [
					" ",
					" ",
					" "
				]
			}
		],
		"Debt - Summary of Term Debt (De": [
			{
				"Debt - Summary of Term Debt (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Debt Instrument [Line Items]": [
					" ",
					" "
				]
			}
		],
		"Debt - Future Principal Payment": [
			{
				"Debt - Future Principal Payments for Term Debt (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Debt Disclosure [Abstract]": [
					" ",
					" "
				]
			},
			{
				"2023": [
					11139,
					" "
				]
			}
		],
		"Shareholders' Equity - Addition": [
			{
				"Shareholders' Equity - Additional Information (Details) shares in Millions, $ in Billions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022 USD ($) shares"
				]
			},
			{
				"Stockholders' Equity Note [Abstract]": [
					" "
				]
			}
		],
		"Shareholders' Equity - Shares o": [
			{
				"Shareholders' Equity - Shares of Common Stock (Details) - shares shares in Thousands": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Increase (Decrease) in Shares of Common Stock Outstanding [Roll Forward]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Benefit Plans - Additional Info": [
			{
				"Benefit Plans - Additional Information (Details) shares in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022 USD ($) shares",
					"Sep. 25, 2021 USD ($) shares",
					"Sep. 26, 2020 USD ($) shares",
					"Mar. 04, 2022 shares",
					"Nov. 09, 2021 shares",
					"Mar. 10, 2015 shares"
				]
			},
			{
				"Share-based Compensation Arrangement by Share-based Payment Award [Line Items]": [
					" ",
					" ",
					" ",
					" ",
					" ",
					" "
				]
			}
		],
		"Benefit Plans - Restricted Stoc": [
			{
				"Benefit Plans - Restricted Stock Units Activity and Related Information (Details) - Restricted stock units - USD ($) $ / shares in Units, shares in Thousands, $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Number of Restricted Stock Units": [
					" ",
					" ",
					" "
				]
			}
		],
		"Benefit Plans - Summary of Shar": [
			{
				"Benefit Plans - Summary of Share-Based Compensation Expense and the Related Income Tax Benefit (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Share-Based Payment Arrangement [Abstract]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Commitments and Contingencies -": [
			{
				"Commitments and Contingencies - Future Payments Under Unconditional Purchase Obligations (Details) $ in Millions": [
					"Sep. 24, 2022 USD ($)"
				]
			},
			{
				"Unconditional Purchase Obligation, Fiscal Year Maturity [Abstract]": [
					" "
				]
			},
			{
				"2023": [
					13488
				]
			}
		],
		"Segment Information and Geogr_3": [
			{
				"Segment Information and Geographic Data - Information by Reportable Segment (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Segment Reporting Information [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Segment Information and Geogr_4": [
			{
				"Segment Information and Geographic Data - Reconciliation of Segment Operating Income to the Consolidated Statements of Operations (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Segment Reporting, Reconciling Item for Operating Profit (Loss) from Segment to Consolidated [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Segment Information and Geogr_5": [
			{
				"Segment Information and Geographic Data - Net Sales (Details) - USD ($) $ in Millions": [
					"12 Months Ended"
				]
			},
			{
				"items": [
					"Sep. 24, 2022",
					"Sep. 25, 2021",
					"Sep. 26, 2020"
				]
			},
			{
				"Revenues from External Customers and Long-Lived Assets [Line Items]": [
					" ",
					" ",
					" "
				]
			}
		],
		"Segment Information and Geogr_6": [
			{
				"Segment Information and Geographic Data - Long-Lived Assets (Details) - USD ($) $ in Millions": [
					"Sep. 24, 2022",
					"Sep. 25, 2021"
				]
			},
			{
				"Revenues from External Customers and Long-Lived Assets [Line Items]": [
					" ",
					" "
				]
			},
			{
				"Long-lived assets": [
					42117,
					39440
				]
			}
		]
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/financial-reports-form-10-k-xlsx · 카테고리: statements
