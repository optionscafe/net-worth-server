//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package robinhood

import (
	"log"
	"testing"

	"github.com/jpfuentes2/go-env"
)

// //
// // TestGetBalances01
// //
// func TestGetBalances01(t *testing.T) {
// 	// Load config file.
// 	env.ReadEnv("../../.env")
//
// 	// Make API call
// 	balance, err := GetBalances()
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	spew.Dump(balance)
// }

//
// Test processJsonResponse()
//
func TestProcessJsonResponse(t *testing.T) {
	// Load config file.
	env.ReadEnv("../../.env")

	// Json string
	jsonStr := `{"unwithdrawable_grants":"0.0000","account":"https://api.robinhood.com/accounts/FF234afadr/","excess_maintenance_with_uncleared_deposits":"1826517.4175","url":"https://api.robinhood.com/portfolios/FF234afadr/","excess_maintenance":"1826517.4175","market_value":"1828664.8900","withdrawable_amount":"24318.7500","last_core_market_value":"1828664.8900","unwithdrawable_deposits":"0.0000","extended_hours_equity":"1828687.0100","excess_margin":"1824351.1950","excess_margin_with_uncleared_deposits":"1824351.1950","equity":"6524.6400","last_core_equity":"128683.6400","adjusted_equity_previous_close":"1828665.7300","equity_previous_close":"1828665.7300","start_date":"2014-12-22","extended_hours_market_value":"1828668.2600"}`

	// Make test request
	balance, err := processJsonResponse(jsonStr)

	if err != nil {
		log.Fatal(err)
	}

	// AccountValue
	if balance.AccountValue != 6524.64 {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 6524.64, balance.AccountValue)
	}
}

/* End File */
