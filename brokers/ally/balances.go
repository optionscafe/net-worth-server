//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package ally

import (
	"github.com/tidwall/gjson"
)

type Balance struct {
	AccountNumber string
	AccountValue  float64
	AccountName   string
}

//
// Get Balances
//
func GetBalances() ([]Balance, error) {

	// Get the JSON
	jsonRt, err := SendGetRequest("/accounts/balances")

	if err != nil {
		return []Balance{}, err
	}

	// Process the json response
	return processJsonResponse(jsonRt)
}

//
// Process JSON
//
func processJsonResponse(jsonRt string) ([]Balance, error) {

	var balances []Balance

	// Do we have only one account?
	vo := gjson.Get(jsonRt, "response.accountbalance")

	// Loop through the different accounts.
	vo.ForEach(func(key, value gjson.Result) bool {

		// Add to balances array
		balances = append(balances, Balance{
			AccountNumber: gjson.Get(value.String(), "account").String(),
			AccountValue:  gjson.Get(value.String(), "accountvalue").Float(),
			AccountName:   gjson.Get(value.String(), "accountname").String(),
		})

		// keep iterating
		return true

	})

	// Return happy
	return balances, nil
}

/* End File */
