# Crowdfunding By CIK

Access detailed information on all crowdfunding campaigns launched by a specific company with the FMP Crowdfunding By CIK API.

## Endpoint

`GET https://financialmodelingprep.com/stable/crowdfunding-offerings?cik=0001916078`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| cik* | string | 0001916078 |

## Description

The FMP Crowdfunding By CIK API provides a comprehensive list of crowdfunding campaigns launched by companies, identified by their Central Index Key (CIK). This endpoint is invaluable for investors and analysts who need to:

- Identify Company-Specific Campaigns: Discover all crowdfunding campaigns initiated by companies you are interested in investing in.

- Track Crowdfunding Activity Over Time: Monitor the crowdfunding activity of specific companies to identify trends, growth, and changes in their fundraising efforts.

- Spot Investment Opportunities: Use the data on crowdfunding campaigns to uncover potential investment opportunities based on the crowdfunding strategies of companies.

This API is essential for those looking to make informed decisions based on the crowdfunding activity of specific companies.

## Response (example)

```json
[
	{
		"cik": "0001916078",
		"companyName": "OYO Fitness, Inc",
		"date": "12-31-2021",
		"filingDate": "2022-07-21 00:00:00",
		"acceptedDate": "2022-07-21 17:28:54",
		"formType": "C-U",
		"formSignification": "Progress Update",
		"nameOfIssuer": "OYO Fitness, Inc",
		"legalStatusForm": "Corporation",
		"jurisdictionOrganization": "DE",
		"issuerStreet": "374 N. 750TH RD",
		"issuerCity": "OVERBROOK",
		"issuerStateOrCountry": "KS",
		"issuerZipCode": "66524",
		"issuerWebsite": "https://www.oyofitness.com/",
		"intermediaryCompanyName": "StartEngine Capital, LLC",
		"intermediaryCommissionCik": "0001665160",
		"intermediaryCommissionFileNumber": "007-00007",
		"compensationAmount": "7 - 13 percent",
		"financialInterest": "Two percent (2%) of securities of the total amount of investments raised in the offering, along the same terms as investors.",
		"securityOfferedType": "Other",
		"securityOfferedOtherDescription": "Non-Voting Common Stock",
		"numberOfSecurityOffered": 5000,
		"offeringPrice": 2,
		"offeringAmount": 10000,
		"overSubscriptionAccepted": "Y",
		"overSubscriptionAllocationType": "Other",
		"maximumOfferingAmount": 1070000,
		"offeringDeadlineDate": "07-19-2022",
		"currentNumberOfEmployees": 5,
		"totalAssetMostRecentFiscalYear": 497717,
		"totalAssetPriorFiscalYear": 248472,
		"cashAndCashEquiValentMostRecentFiscalYear": 150142,
		"cashAndCashEquiValentPriorFiscalYear": 54571,
		"accountsReceivableMostRecentFiscalYear": 0,
		"accountsReceivablePriorFiscalYear": 0,
		"shortTermDebtMostRecentFiscalYear": 3286745,
		"shortTermDebtPriorFiscalYear": 2214117,
		"longTermDebtMostRecentFiscalYear": 82243,
		"longTermDebtPriorFiscalYear": 105850,
		"revenueMostRecentFiscalYear": 4344154,
		"revenuePriorFiscalYear": 11078510,
		"costGoodsSoldMostRecentFiscalYear": 2445024,
		"costGoodsSoldPriorFiscalYear": 5737776,
		"taxesPaidMostRecentFiscalYear": 0,
		"taxesPaidPriorFiscalYear": 0,
		"netIncomeMostRecentFiscalYear": -964551,
		"netIncomePriorFiscalYear": -10860
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/crowdfunding-by-cik · 카테고리: Fundraisers
