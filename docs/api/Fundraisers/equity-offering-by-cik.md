# Equity Offering By CIK

Access detailed information on equity offerings announced by specific companies with the FMP Company Equity Offerings by CIK API. Track offering activity and identify potential investment opportunities.

## Endpoint

`GET https://financialmodelingprep.com/stable/fundraising?cik=0001547416`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| cik* | string | 0001547416 |

## Description

The FMP Company Equity Offerings by CIK API provides a comprehensive list of all equity offerings announced by a particular company, identified by its Central Index Key (CIK). This API is essential for:

- Identifying Company-Specific Offerings: Quickly find and track equity offerings announced by companies you are interested in by searching with their CIK.

- Tracking Offering Activity Over Time: Monitor the equity offering history of specific companies to gain insights into their financing activities and strategic moves.

- Spotting Investment Opportunities: Use equity offering data to identify potential investment opportunities, understanding how a company's offering activity might impact its stock price and market position.

Investors can leverage this API to stay informed about the equity offering activities of the companies they follow, allowing for more informed investment decisions.

## Response (example)

```json
[
	{
		"cik": "0001547416",
		"companyName": "NJOY INC",
		"date": "2014-02-28",
		"filingDate": "2014-02-28 00:00:00",
		"acceptedDate": "2014-02-28 16:00:25",
		"formType": "D",
		"formSignification": "Notice of Exempt Offering of Securities",
		"entityName": "NJOY INC",
		"issuerStreet": "15211 N. KIERLAND BLVD., SUITE 200",
		"issuerCity": "SCOTTSDALE",
		"issuerStateOrCountry": "AZ",
		"issuerStateOrCountryDescription": "ARIZONA",
		"issuerZipCode": "85254",
		"issuerPhoneNumber": "480-397-2300",
		"jurisdictionOfIncorporation": "DELAWARE",
		"entityType": "Corporation",
		"incorporatedWithinFiveYears": null,
		"yearOfIncorporation": "",
		"relatedPersonFirstName": "CRAIG",
		"relatedPersonLastName": "WEISS",
		"relatedPersonStreet": "c/o NJOY, INC.",
		"relatedPersonCity": "SCOTTSDALE",
		"relatedPersonStateOrCountry": "AZ",
		"relatedPersonStateOrCountryDescription": "ARIZONA",
		"relatedPersonZipCode": "85254",
		"relatedPersonRelationship": "Executive Officer, Director",
		"industryGroupType": "Other",
		"revenueRange": "Decline to Disclose",
		"federalExemptionsExclusions": "06b",
		"isAmendment": false,
		"dateOfFirstSale": "2014-02-14",
		"durationOfOfferingIsMoreThanYear": false,
		"securitiesOfferedAreOfEquityType": true,
		"isBusinessCombinationTransaction": false,
		"minimumInvestmentAccepted": 0,
		"totalOfferingAmount": 71999990,
		"totalAmountSold": 71999990,
		"totalAmountRemaining": 0,
		"hasNonAccreditedInvestors": false,
		"totalNumberAlreadyInvested": 24,
		"salesCommissions": 0,
		"findersFees": 0,
		"grossProceedsUsed": 0
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/equity-offering-by-cik · 카테고리: Fundraisers
