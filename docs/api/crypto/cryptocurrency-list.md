# Cryptocurrency List

Access a comprehensive list of all cryptocurrencies traded on exchanges worldwide with the FMP Cryptocurrencies Overview API. Get detailed information on each cryptocurrency to inform your investment strategies.

## Endpoint

`GET https://financialmodelingprep.com/stable/cryptocurrency-list`

## Description

The FMP Cryptocurrencies Overview API provides detailed information on all cryptocurrencies that are actively traded on global exchanges. This API is essential for:

- Cryptocurrency Identification: Access a list of all traded cryptocurrencies, including their symbols, names, and the fiat currency they are paired with. This data helps investors identify different cryptocurrencies and understand their market presence.

- Exchange Details: The API also provides information about the exchange where the cryptocurrency is listed, including the exchange name and a short name identifier. This allows investors to track where each cryptocurrency is traded.

- Informed Decision-Making: Use the detailed data provided by this API to track cryptocurrency performance, monitor market trends, and make informed investment decisions.

Example

Market Analysis: A crypto trader might use the Cryptocurrencies Overview API to compile a list of all cryptocurrencies paired with USD across different exchanges. By analyzing this data, the trader can identify which cryptocurrencies are gaining popularity and may present investment opportunities.

## Response (example)

```json
[
	{
		"symbol": "ALIENUSD",
		"name": "Alien Inu USD",
		"exchange": "CCC",
		"icoDate": "2021-11-22",
		"circulatingSupply": 0,
		"totalSupply": null
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/cryptocurrency-list · 카테고리: crypto
