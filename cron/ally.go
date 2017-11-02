//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package cron

import (
  "os"
  "time"
  "errors"
  "net/http"  
  "github.com/jasonlvhit/gocron" 
  "github.com/net-worth-server/models"
  "github.com/net-worth-server/services"    
  "github.com/net-worth-server/brokers/ally"
)

//
// Setup Cron Job
//
func AllyStart(db *models.DB) {

  go func() {

    // Start scheduling
    s := gocron.NewScheduler()
  
    // Lets get started
    services.LogInfo("Cron Ally Started")  

    // Setup jobs we need to run 
    //s.Every(1).Second().Do(func () { MarkAlly(db) })
    s.Every(1).Day().At("22:02").Do(func () { MarkAlly(db) })

    // function Start all the pending jobs
    <- s.Start()
  
  }()

}

//
// Make a call to ally get balance and then market it.
// This is called once a day at the end of the day.
//
func MarkAlly(db *models.DB) error {

  var targetAccountId uint = 0

  // Log
  services.LogInfo("Ally starting account marked from cron")

  // Get accounts in our system
  accounts := db.GetAllAcounts()

  // Loop through and find the account Id
  for _, row := range accounts {
    if row.AccountNumber == os.Getenv("ALLY_ACCOUNT") {
      targetAccountId = row.Id   
    }
  }

  // Make sure we found an account
  if targetAccountId == 0 {
    return errors.New("No account found - " + os.Getenv("ALLY_ACCOUNT"))
  }

  // Get all the balances from Tradier.
  balances, err := ally.GetBalances()

  if err != nil {
    return err
  } 

  // Format the date for today.
  ts := time.Now()
  date := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)

  // Process the Marked data.
  processMarkAlly(db, balances, targetAccountId, date, os.Getenv("ALLY_ACCOUNT"))

  // Send health check notice.
  if len(os.Getenv("HEALTH_CHECK_ALLY_URL")) > 0 {

    resp, err := http.Get(os.Getenv("HEALTH_CHECK_ALLY_URL"))
    
    if err != nil {
      services.LogError(err, "Could send HEALTH_CHECK_ALLY_URL health check - " + os.Getenv("HEALTH_CHECK_ALLY_URL"))
    }
    
    defer resp.Body.Close()
    
  }

  // Log
  services.LogInfo("Ally account marked from cron")

  // Return happy
  return nil
}

//
// Process marking from Ally.
//
func processMarkAlly(db *models.DB, balances []ally.Balance, targetAccountId uint, date time.Time, AccountNumber string) error {

  // Loop through the different Tradier accounts.
  for _, row := range balances {

    // Find the ally account we are after. Then mark the asset
    if row.AccountNumber == AccountNumber {
      db.MarkAccountByDate(targetAccountId, date, row.AccountValue)
      break
    }
    
  }

  // Return happy
  return nil
}

/* End File */