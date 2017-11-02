//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package robinhood

import (
  "os"
  "github.com/tidwall/gjson"   
)

type Balance struct {
  AccountValue float64
}

//
// Get Balances
//
func GetBalances() (Balance, error) {
    
  // Get the JSON
  jsonRt, err := SendGetRequest("/accounts/" + os.Getenv("ROBINHOOD_ACCOUNT") + "/portfolio/")

  if err != nil {
    return Balance{}, err  
  }

  // Process the json response
  return processJsonResponse(jsonRt) 
}

//
// Process JSON
//
func processJsonResponse(jsonRt string) (Balance, error) {

  balance := Balance{ AccountValue: gjson.Get(jsonRt, "equity").Float() }
  
  // Return happy
  return balance, nil  
}

/* End File */