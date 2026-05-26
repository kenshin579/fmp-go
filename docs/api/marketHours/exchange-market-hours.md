# Global Exchange Market Hours

Retrieve trading hours for specific stock exchanges using the Global Exchange Market Hours API. Find out the opening and closing times of global exchanges to plan your trading strategies effectively.

## Endpoint

`GET https://financialmodelingprep.com/stable/exchange-market-hours?exchange=NASDAQ`

## Parameters

| Query Parameter | Type | Example |
| --- | --- | --- |
| exchange* | string | NASDAQ |
| timestamp | string | 1769527402 |

## Description

The FMP Global Exchange Market Hours API provides essential information about the opening and closing hours of various stock exchanges around the world. This API helps users track when exchanges like NASDAQ, NYSE, and others are open for trading, along with information about the time zone and whether the market is currently open. Key features include:

- Trading Hours by Exchange: Access the opening and closing times for specific stock exchanges worldwide.

- Real-Time Market Status: Find out if the market is currently open or closed for trading.

- Time Zone Support: View exchange market hours in the local time zone of each exchange for accurate planning.

- Global Exchange Coverage: Get information on major stock exchanges, including NASDAQ, NYSE, and others.

This API is ideal for traders, analysts, and investors who need to stay informed about market hours to manage their trading strategies across different regions.

## Response (example)

```json
[
	{
		"exchange": "NASDAQ",
		"name": "NASDAQ",
		"openingHour": "09:30 AM -04:00",
		"closingHour": "04:00 PM -04:00",
		"timezone": "America/New_York",
		"isMarketOpen": false
	}
]
```

> 출처: https://site.financialmodelingprep.com/developer/docs/stable/exchange-market-hours · 카테고리: marketHours
