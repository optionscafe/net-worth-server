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
  "github.com/net-worth-server/brokers/robinhood"
)

//
// Setup Cron Job
//
func RobinhoodStart(db *models.DB) {

  go func() {
  
    // Start scheduling
    s := gocron.NewScheduler()

    // Lets get started
    services.LogInfo("Cron Robinhood Started")  

    // Setup jobs we need to run 
    //s.Every(1).Second().Do(func () { MarkRobinhood(db) })
    s.Every(1).Day().At("22:01").Do(func () { MarkRobinhood(db) })

    // function Start all the pending jobs
    <- s.Start()
  
  }()

}

//
// Make a call to Robinhood get balance and then market it.
// This is called once a day at the end of the day.
//
func MarkRobinhood(db *models.DB) error {

  var targetAccountId uint = 0

  // Log
  services.LogInfo("Robinhood starting account marked from cron")

  // Get accounts in our system
  accounts := db.GetAllAcounts()

  // Loop through and find the account Id
  for _, row := range accounts {
    if row.AccountNumber == os.Getenv("ROBINHOOD_ACCOUNT") {
      targetAccountId = row.Id   
    }
  }

  // Make sure we found an account
  if targetAccountId == 0 {
    return errors.New("No account found - " + os.Getenv("ROBINHOOD_ACCOUNT"))
  }

  // Get all the balances from Robinhood.
  balance, err := robinhood.GetBalances()

  if err != nil {
    return err
  } 

  // Format the date for today.
  ts := time.Now()
  date := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)

  // Process the Marked data.
  processMarkRobinhood(db, balance, targetAccountId, date)

  // Send health check notice.
  if len(os.Getenv("HEALTH_CHECK_ROBINHOOD_URL")) > 0 {

    resp, err := http.Get(os.Getenv("HEALTH_CHECK_ROBINHOOD_URL"))
    
    if err != nil {
      services.LogError(err, "Could send HEALTH_CHECK_ROBINHOOD_URL health check - " + os.Getenv("HEALTH_CHECK_ROBINHOOD_URL"))
    }
    
    defer resp.Body.Close()
    
  }

  // Log
  services.LogInfo("Robinhood account marked from cron")

  // Return happy
  return nil
}

//
// Process marking from Robinhood.
//
func processMarkRobinhood(db *models.DB, balance robinhood.Balance, targetAccountId uint, date time.Time) error {

  // Mark asset
  db.MarkAccountByDate(targetAccountId, date, balance.AccountValue)

  // Return happy
  return nil
}

/* End File */