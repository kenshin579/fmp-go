# Equity Offering Updates

Stay informed about the latest equity offerings with the FMP Equity Offering Updates API. Track new shares being issued by companies and get insights into exempt offerings and amendments.

## Endpoint

`GET https://financialmodelingprep.com/stable/fundraising-latest?page=0&limit=10`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| page | number | 0 |
| limit | number | 10 |
| cik | string | 0002013736 |

## Description

The FMP Equity Offering Updates API provides detailed information on newly issued equity securities, including company details, offering amounts, and regulatory filings. This API is a crucial tool for investors, analysts, and market researchers who need to:

- Monitor New Equity Issuances: Track companies issuing new shares and stay informed about recent equity offerings.

- Analyze Offering Details: Access important data such as filing dates, form types, industry classifications, and the minimum investment required.

- Stay Compliant: Get information on exempt offerings under regulations like 06b, 3C, and 3C.1 to assess the legal status of an equity issue.

This API is invaluable for keeping up-to-date with the latest equity issuances, ensuring you never miss an important offering or amendment.

Example Use Case
An institutional investor could use the Equity Offering Updates API to identify new investment opportunities by tracking newly issued equity offerings from companies across various sectors.

## Response (example)

```json
[
	{
		"cik": "0002103666",
		"companyName": "Evolution Ventures Minerva Fund, LP - B4",
		"date": "2026-04-08",
		"filingDate": "2026-04-08 00:00:00",
		"acceptedDate": "2026-04-08 17:30:42",
		"formType": "D/A",
		"formSignification": "Notice of Exempt Offering of Securities Amendement",
		"entityName": "Evolution Ventures Minerva Fund, LP - B4",
		"issuerStreet": "2006 196TH ST SW",
		"issuerCity": "LYNNWOOD",
		"issuerStateOrCountry": "WA",
		"issuerStateOrCountryDescription": "WASHINGTON",
		"issuerZipCode": "98036",
		"issuerPhoneNumber": "206.801.6359",
		"jurisdictionOfIncorporation": "DELAWARE",
		"entityType": "Limited Partnership",
		"incorporatedWithinFiveYears": true,
		"yearOfIncorporation": "2025",
		"relatedPersonFirstName": "N/A",
		"relatedPersonLastName": "Fund GP, LLC",
		"relatedPersonStreet": "301 North Market Street, Suite 1414",
		"relatedPersonCity": "Wilmington",
		"relatedPersonStateOrCountry": "DE",
		"relatedPersonStateOrCountryDescription": "DELAWARE",
		"relatedPersonZipCode": "19801",
		"relatedPersonRelationship": "Director",
		"industryGroupType": "Pooled Investment Fund",
		"revenueRange": "Decline to Disclose",
		"federalExemptionsExclusions": "06b, 3C, 3C.1",
		"isAmendment": true,
		"dateOfFirstSale": "2026-01-01",
		"durationOfOfferingIsMoreThanYear": false,
		"securitiesOfferedAreOfEquityType": null,
		"isBusinessCombinationTransaction": false,
		"minimumInvestmentAccepted": 10000,
		"totalOfferingAmount": 186842,
		"totalAmountSold": 186842,
		"totalAmountRemaining": 0,
		"hasNonAccreditedInvestors": false,
		"totalNumberAlreadyInvested": 17,
		"salesCommissions": 0,
		"findersFees": 0,
		"grossProceedsUsed": 20000
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/latest-equity-offering · 카테고리: Fundraisers
