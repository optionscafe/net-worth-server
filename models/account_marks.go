//
// Date: 10/20/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"time"

	"github.com/optionscafe/net-worth-server/services"
)

type AccountMarks struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Date      time.Time `gorm:"type:date" json:"date"`
	AccountId uint      `sql:"not null;index:UserId" json:"account_id"`
	Units     float64   `sql:"type:DECIMAL(12,2)" json:"units"`
	PricePer  float64   `sql:"type:DECIMAL(12,2)" json:"price_per"`
	Balance   float64   `sql:"type:DECIMAL(12,2)" json:"balance"`
}

//
// Get marks account by id.
//
func (db *DB) GetMarksByAccountById(accountId uint) []AccountMarks {

	m := []AccountMarks{}

	// Make query
	db.Where("account_id = ?", accountId).Order("date desc").Find(&m)

	// Return the accounts.
	return m
}

//
// Create Account Mark. We only allow one per date per account.
// If this mark is already in place we simply update it.
//
func (db *DB) MarkAccountByDate(accountId uint, date time.Time, balance float64) error {

	m := AccountMarks{}

	// Get total units in this account.
	units := db.GetMarkAccountUnitsByAccountId(accountId)

	// Figure out price per unit.
	perUnit := balance / units

	// Validate to make sure we do not already have this record.
	if err := db.Where("account_id = ? AND date = ?", accountId, date).First(&m).Error; err != nil {

		// Create new mark
		mark := AccountMarks{AccountId: accountId, Date: date.UTC(), Balance: balance, Units: units, PricePer: perUnit}

		// Insert new mark
		if err := db.Create(&mark).Error; err != nil {
			services.LogError(err, "MarkAccountByDate : Create new mark")
			return err
		}

	} else {

		// Update mark
		if err := db.Model(&m).Update("balance", balance).Update("units", units).Update("price_per", perUnit).Error; err != nil {
			services.LogError(err, "MarkAccountByDate : Update mark")
			return err
		}

	}

	// Update the balance on the account level
	account := Account{}
	db.First(&account, accountId)
	account.Balance = balance
	db.Save(&account)

	// Mark for all accounts.
	db.MarkByDate(date)

	// Return happy.
	return nil
}

//
// Get mark by date and account id.
//
func (db *DB) GetMarksByAccountByIdAndDate(accountId uint, date time.Time) (AccountMarks, error) {

	m := AccountMarks{}

	// Find result or send error
	if err := db.Where("account_id = ? AND date = ?", accountId, date).First(&m).Error; err != nil {
		return m, err
	}

	// Return the accounts.
	return m, nil
}

//
// Add units to the object.
//
func (db *DB) GetMarkAccountUnitsByAccountId(id uint) float64 {

	// Struct to capture the sum result.
	type Result struct {
		Sum float64
	}

	var u Result

	// Query and get unit count.
	db.Raw("SELECT SUM(units) AS sum FROM account_units WHERE account_id = ?", id).Scan(&u)

	// Return count
	return u.Sum
}

/* End File */
