//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package cron

import (
  "time"
  "testing"
  "github.com/jpfuentes2/go-env"  
  "github.com/net-worth-server/models"
  "github.com/net-worth-server/brokers/robinhood"        
)

//
// Make a call to Robinhood get balance and then market it.
// This is called once a day at the end of the day.
//
func TestMarkRobinhood(t *testing.T) {

  // Load config file. 
  env.ReadEnv("../.env")

  // Shared time.
  ts := time.Date(2017, 10, 31, 0, 0, 0, 0, time.UTC)

  // Start the db connection.
  db, _ := models.NewDB()
  defer db.Close()  

  var balance = robinhood.Balance{ AccountValue: 5501.02 }

  // Process the mark
  processMarkRobinhood(db, balance, 4, ts)

  // Query to make sure this was stored in the database as it should.
  account, _ := db.GetAccountById(4)

  // Query and get the mark we just placed.
  mark, _ := db.GetMarksByAccountByIdAndDate(4, ts)

  // Did we store the account balance in the correct field?
  if account.Balance != 5501.02 {
    t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 5501.02, account.Balance)
  }

  // Mark - Date
  if mark.Date != ts {
    t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", ts.String(), mark.Date.String())
  }

  // Mark - Balance
  if mark.Balance != 5501.02 {
    t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 5501.02, mark.Balance)
  }

  // Mark - PricePer
  if mark.PricePer != 1.22 {
    t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 1.22, mark.PricePer)
  }

}

/* End File */