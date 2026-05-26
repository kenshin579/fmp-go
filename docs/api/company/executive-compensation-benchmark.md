# Executive Compensation Benchmark

Gain access to average executive compensation data across various industries with the FMP Executive Compensation Benchmark API. This API provides essential insights for comparing executive pay by industry, helping you understand compensation trends and benchmarks.

## Endpoint

`GET https://financialmodelingprep.com/stable/executive-compensation-benchmark`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| year | string | 2024 |

## Description

The FMP Executive Compensation Benchmark API is designed to help businesses, analysts, and compensation consultants assess how executive pay compares across industries. It's ideal for:

- Industry Benchmarking: Evaluate average executive compensation within specific industries to determine market rates.

- Compensation Trends: Understand how executive pay varies across different sectors, providing valuable insights for salary negotiations or organizational planning.

- Competitive Analysis: Compare compensation data by industry to ensure your company remains competitive in attracting top talent.

This API provides a valuable resource for HR professionals, compensation analysts, and business leaders seeking to align executive pay with industry standards.

Example Use Case
An HR professional can use the Executive Compensation Benchmark API to compare the average pay for executives in the technology sector against those in the consumer goods sector, helping to determine competitive salary packages for their company's leadership team.

## Response (example)

```json
[
	{
		"industryTitle": "ABRASIVE, ASBESTOS & MISC NONMETALLIC MINERAL PRODS",
		"year": 2024,
		"averageCompensation": 784407.5555555555
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/executive-compensation-benchmark · 카테고리: company
