//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package ally

import (
  "log"
  "testing"
  "github.com/jpfuentes2/go-env"      
)

//
// Test processJsonResponse()
//
func TestProcessJsonResponse(t *testing.T) {

  // Load config file. 
  env.ReadEnv("../../.env")

  jsonStr := `{ "response":{ "@id":"8292134-3234-444e-af69-832913", "elapsedtime": "0", "accountbalance":[{"account":"8888888","accountname":"Traditional IRA","accountvalue":"8323.23"},{"account":"JA234fDD","accountname":"Trading Account","accountvalue":"10123.44"},{"account":"88823432DDB","accountname":"Roth IRA","accountvalue":"424305.42"}],"totalbalance":{"accountvalue":"12324305.42"}, "error": "Success" } }`

  // Make test request
  balances, err := processJsonResponse(jsonStr)

  if err != nil {
    log.Fatal(err)
  }

  // Expected 
  expectedBalances := []float64{ 8323.23, 10123.44, 424305.42 }
  expectedNumber := []string{ "8888888", "JA234fDD", "88823432DDB" }
  expectedNames := []string{ "Traditional IRA", "Trading Account", "Roth IRA" }

  // Loop over the results.
  for key, row := range balances {

    // AccountValue
    if row.AccountValue != expectedBalances[key] {
      t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expectedBalances[key], row.AccountValue)
    }

    // AccountName
    if row.AccountName != expectedNames[key] {
      t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expectedNames[key], row.AccountName)
    }

    // AccountNumber
    if row.AccountNumber != expectedNumber[key] {
      t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expectedNumber[key], row.AccountNumber)
    }

  }

}

/* End File */