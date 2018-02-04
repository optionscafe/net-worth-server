//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package cron

import (
	"github.com/jpfuentes2/go-env"
	"github.com/optionscafe/net-worth-server/brokers/tradier"
	"github.com/optionscafe/net-worth-server/models"
	"testing"
	"time"
)

//
// Make a call to tradier get balance and then market it.
// This is called once a day at the end of the day.
//
func TestMarkTradier(t *testing.T) {

	// Load config file.
	env.ReadEnv("../.env")

	// Shared time.
	ts := time.Date(2017, 10, 31, 0, 0, 0, 0, time.UTC)

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	var balances = []tradier.Balance{
		{
			AccountNumber:     "12312312",
			AccountValue:      100.00,
			TotalCash:         100.00,
			OptionBuyingPower: 100.00,
			StockBuyingPower:  100.00,
		},

		{
			AccountNumber:     "7af234fS",
			AccountValue:      66444.77,
			TotalCash:         66444.77,
			OptionBuyingPower: 66444.77,
			StockBuyingPower:  66444.77,
		},

		{
			AccountNumber:     "1ba4dk9d",
			AccountValue:      777.11,
			TotalCash:         777.11,
			OptionBuyingPower: 777.11,
			StockBuyingPower:  777.11,
		},
	}

	// Process the mark
	processMarkTradier(db, balances, 1, ts, "7af234fS")

	// Query to make sure this was stored in the database as it should.
	account, _ := db.GetAccountById(1)

	// Query and get the mark we just placed.
	mark, _ := db.GetMarksByAccountByIdAndDate(1, ts)

	// Did we store the account balance in the correct field?
	if account.Balance != 66444.77 {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 66444.77, account.Balance)
	}

	// Mark - Date
	if mark.Date != ts {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", ts.String(), mark.Date.String())
	}

	// Mark - Balance
	if mark.Balance != 66444.77 {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 66444.77, mark.Balance)
	}

	// Mark - PricePer
	if mark.PricePer != 4.53 {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 4.53, mark.PricePer)
	}

}

/* End File */
