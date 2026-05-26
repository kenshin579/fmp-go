# Latest Crowdfunding Campaigns

Discover the most recent crowdfunding campaigns with the FMP Latest Crowdfunding Campaigns API. Stay informed on which companies and projects are actively raising funds, their financial details, and offering terms.

## Endpoint

`GET https://financialmodelingprep.com/stable/crowdfunding-offerings-latest?page=0&limit=100`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 100 |

## Description

The FMP Latest Crowdfunding Campaigns API provides detailed information on current crowdfunding campaigns, including the names of issuers, offering types, and financial data. This API is essential for investors, analysts, and platforms that want to track the latest crowdfunding activity.

- Track Crowdfunding Campaigns: Access the most up-to-date information on crowdfunding campaigns, including company names, funding goals, and offering types.

- Detailed Financial Information: View key financial metrics such as total assets, cash equivalents, debt, and net income for each company running a crowdfunding campaign.

- Company Backgrounds: Get insights into the legal status and jurisdiction of the companies, including the number of employees and other relevant organizational data.

This API is a valuable tool for those looking to follow new crowdfunding opportunities, assess potential investments, or stay up to date on market trends in the crowdfunding space.

Example Use Case
An investor can use the Crowdfunding Campaigns API to review the financial health and offering details of various crowdfunding campaigns, helping them evaluate potential opportunities and diversify their portfolio.

## Response (example)

```json
[
	{
		"cik": "0001532978",
		"companyName": "Gumroad, Inc.",
		"date": "09-22-2011",
		"filingDate": "2026-04-08 00:00:00",
		"acceptedDate": "2026-04-08 16:54:45",
		"formType": "C-AR",
		"formSignification": "Annual Report",
		"nameOfIssuer": "Gumroad, Inc.",
		"legalStatusForm": "Corporation",
		"jurisdictionOrganization": "DE",
		"issuerStreet": "548 Market St, #41309",
		"issuerCity": "San Francisco",
		"issuerStateOrCountry": "CA",
		"issuerZipCode": "94104",
		"issuerWebsite": "https://gumroad.com/",
		"intermediaryCompanyName": null,
		"intermediaryCommissionCik": "0001532978",
		"intermediaryCommissionFileNumber": null,
		"compensationAmount": null,
		"financialInterest": null,
		"securityOfferedType": null,
		"securityOfferedOtherDescription": null,
		"numberOfSecurityOffered": 0,
		"offeringPrice": 0,
		"offeringAmount": 0,
		"overSubscriptionAccepted": "N",
		"overSubscriptionAllocationType": null,
		"maximumOfferingAmount": 0,
		"offeringDeadlineDate": null,
		"currentNumberOfEmployees": 2,
		"totalAssetMostRecentFiscalYear": 11948947.05,
		"totalAssetPriorFiscalYear": 16720734.62,
		"cashAndCashEquiValentMostRecentFiscalYear": 6153268.63,
		"cashAndCashEquiValentPriorFiscalYear": 13821885.61,
		"accountsReceivableMostRecentFiscalYear": 0,
		"accountsReceivablePriorFiscalYear": 0,
		"shortTermDebtMostRecentFiscalYear": 4191955.58,
		"shortTermDebtPriorFiscalYear": 4635820.52,
		"longTermDebtMostRecentFiscalYear": 0,
		"longTermDebtPriorFiscalYear": 0,
		"revenueMostRecentFiscalYear": 17785730.09,
		"revenuePriorFiscalYear": 18951308.03,
		"costGoodsSoldMostRecentFiscalYear": 6078822.25,
		"costGoodsSoldPriorFiscalYear": 5923421.2,
		"taxesPaidMostRecentFiscalYear": 1815741.46,
		"taxesPaidPriorFiscalYear": 354538.04,
		"netIncomeMostRecentFiscalYear": 3792179.17,
		"netIncomePriorFiscalYear": 4290423.04
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/latest-crowdfunding · 카테고리: Fundraisers
