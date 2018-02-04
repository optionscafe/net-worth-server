//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"time"

	"github.com/optionscafe/net-worth-server/services"
)

type AccountUnits struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AccountId uint      `sql:"not null;index:UserId" json:"account_id"`
	Date      time.Time `gorm:"type:date" json:"date"`
	Amount    float64   `sql:"type:DECIMAL(12,2)" json:"amount"`
	Units     float64   `sql:"type:DECIMAL(12,2)" json:"units"`
	PricePer  float64   `sql:"type:DECIMAL(12,2)" json:"price_per"`
	Note      string    `sql:"not null" json:"note"`
}

//
// Add / Subtract funds from an account.
//
func (db *DB) AccountUnitsAddFunds(accountId uint, date time.Time, amount float64, note string) error {

	// Get account.
	account, err := db.GetAccountById(accountId)

	if err != nil {
		services.LogError(err, "")
		return err
	}

	// Figure out units and price per.
	pricePer := account.Balance / account.Units
	units := amount / pricePer

	obj := AccountUnits{AccountId: accountId, Date: date, Amount: amount, Units: units, PricePer: pricePer, Note: note}

	if err := db.Create(&obj).Error; err != nil {
		services.LogError(err, "")
		return err
	}

	// Add the new units to the over all units.
	if err := db.AddUnits(date, amount, note); err != nil {
		services.LogError(err, "")
		return err
	}
	// Update the account with new balance.
	account.Balance = account.Balance + amount
	db.Save(&account)

	// Mark account with new value.
	err = db.MarkAccountByDate(accountId, date, account.Balance)

	if err != nil {
		services.LogError(err, "")
		return err
	}

	// Return success.
	return nil
}

/* End File */
