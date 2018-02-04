//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package cron

import (
	"errors"
	"github.com/jasonlvhit/gocron"
	"github.com/optionscafe/net-worth-server/brokers/tradier"
	"github.com/optionscafe/net-worth-server/models"
	"github.com/optionscafe/net-worth-server/services"
	"net/http"
	"os"
	"time"
)

//
// Setup Cron Job
//
func TradierStart(db *models.DB) {

	go func() {

		// Start scheduling
		s := gocron.NewScheduler()

		// Lets get started
		services.LogInfo("Cron Tradier Started")

		// Setup jobs we need to run
		//s.Every(1).Second().Do(func () { MarkTradier(db) })
		s.Every(1).Day().At("22:00").Do(func() { MarkTradier(db) })

		// function Start all the pending jobs
		<-s.Start()

	}()

}

//
// Make a call to tradier get balance and then market it.
// This is called once a day at the end of the day.
//
func MarkTradier(db *models.DB) error {

	var targetAccountId uint = 0

	// Log
	services.LogInfo("Tradier starting account marked from cron")

	// Get accounts in our system
	accounts := db.GetAllAcounts()

	// Loop through and find the account Id
	for _, row := range accounts {
		if row.AccountNumber == os.Getenv("TRADIER_ACCOUNT") {
			targetAccountId = row.Id
		}
	}

	// Make sure we found an account
	if targetAccountId == 0 {
		return errors.New("No account found - " + os.Getenv("TRADIER_ACCOUNT"))
	}

	// Get all the balances from Tradier.
	balances, err := tradier.GetBalances()

	if err != nil {
		return err
	}

	// Format the date for today.
	ts := time.Now()
	date := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)

	// Process the Marked data.
	processMarkTradier(db, balances, targetAccountId, date, os.Getenv("TRADIER_ACCOUNT"))

	// Send health check notice.
	if len(os.Getenv("HEALTH_CHECK_MARKTRADIER_URL")) > 0 {

		resp, err := http.Get(os.Getenv("HEALTH_CHECK_MARKTRADIER_URL"))

		if err != nil {
			services.LogError(err, "Could send health check - "+os.Getenv("HEALTH_CHECK_MARKTRADIER_URL"))
		}

		defer resp.Body.Close()

	}

	// Log
	services.LogInfo("Tradier account marked from cron")

	// Return happy
	return nil
}

//
// Process marking from Tradier.
//
func processMarkTradier(db *models.DB, balances []tradier.Balance, targetAccountId uint, date time.Time, accountNumber string) error {

	// Loop through the different Tradier accounts.
	for _, row := range balances {

		// Find the tradier account we are after. Then mark the asset
		if row.AccountNumber == accountNumber {
			db.MarkAccountByDate(targetAccountId, date, row.AccountValue)
			break
		}

	}

	// Return happy
	return nil
}

/* End File */
