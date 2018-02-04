//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package cron

import (
	"github.com/jpfuentes2/go-env"
	"github.com/optionscafe/net-worth-server/brokers/ally"
	"github.com/optionscafe/net-worth-server/models"
	"testing"
	"time"
)

//
// Make a call to ally get balance and then market it.
// This is called once a day at the end of the day.
//
func TestMarkAlly(t *testing.T) {

	// Load config file.
	env.ReadEnv("../.env")

	// Shared time.
	ts := time.Date(2017, 10, 31, 0, 0, 0, 0, time.UTC)

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	var balances = []ally.Balance{
		{
			AccountNumber: "12312312",
			AccountValue:  100.00,
			AccountName:   "Account ABC",
		},

		{
			AccountNumber: "23423499",
			AccountValue:  66444.77,
			AccountName:   "Account 123",
		},

		{
			AccountNumber: "1ba4dk9d",
			AccountValue:  777.11,
			AccountName:   "Account XYZ",
		},
	}

	// Process the mark
	processMarkAlly(db, balances, 2, ts, "23423499")

	// Query to make sure this was stored in the database as it should.
	account, _ := db.GetAccountById(2)

	// Query and get the mark we just placed.
	mark, _ := db.GetMarksByAccountByIdAndDate(2, ts)

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
	if mark.PricePer != 0.78 {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 0.78, mark.PricePer)
	}

}

/* End File */
