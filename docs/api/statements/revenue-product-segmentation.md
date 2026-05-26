# Revenue Product Segmentation

Access detailed revenue breakdowns by product line with the Revenue Product Segmentation API. Understand which products drive a company's earnings and get insights into the performance of individual product segments.

## Endpoint

`GET https://financialmodelingprep.com/stable/revenue-product-segmentation?symbol=AAPL`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| symbol* | string | AAPL |
| period | string | annual,quarter |
| structure | string | flat |

## Description

The Revenue Product Segmentation API provides a comprehensive breakdown of a company's revenue by product, making it easy to analyze performance across different product categories. This API is ideal for:

- Product-Specific Revenue Analysis: Understand how much each product contributes to the company's total earnings.

- Strategic Insights: Gain insights into the growth or decline of specific product segments to inform investment decisions or corporate strategy.

- Competitive Benchmarking: Compare product segment revenues across different companies in the same industry to gauge market position.

This API offers a detailed view of product-level revenue, helping users identify growth drivers and track the financial health of specific product lines.

Example Use Case
An investor can use the Revenue Product Segmentation API to see how much of Apple's earnings come from iPhone sales compared to other products, such as Macs or wearables.

## Response (example)

```json
[
	{
		"symbol": "AAPL",
		"fiscalYear": 2024,
		"period": "FY",
		"reportedCurrency": null,
		"date": "2024-09-28",
		"data": {
			"Mac": 29984000000,
			"Service": 96169000000,
			"Wearables, Home and Accessories": 37005000000,
			"iPad": 26694000000,
			"iPhone": 201183000000
		}
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/revenue-product-segmentation · 카테고리: statements
