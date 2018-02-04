//
// Date: 10/27/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"strings"
	"time"

	"github.com/optionscafe/net-worth-server/services"
)

type Ledger struct {
	Id           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Date         time.Time `gorm:"type:date" json:"date"`
	AccountId    uint      `sql:"not null;index:UserId" json:"account_id"`
	AccountName  string    `gorm:"-" json:"account_name"`
	CategoryId   uint      `sql:"not null;index:CategoryId" json:"category_id"`
	CategoryName string    `gorm:"-" json:"category_name"`
	Amount       float64   `sql:"type:DECIMAL(12,2)" json:"amount"`
	Note         string    `sql:"not null" json:"note"`
}

//
// Get all marks.
//
func (db *DB) GetAllLedgers() []Ledger {

	ledgers := []Ledger{}

	db.Order("account_id asc").Find(&ledgers)

	// Add in our one to one look ups
	for i := range ledgers {
		account, _ := db.GetAccountById(ledgers[i].AccountId)
		ledgers[i].AccountName = account.Name

		category, _ := db.GetLedgerCategoryById(ledgers[i].CategoryId)
		ledgers[i].CategoryName = category.Name
	}

	return ledgers
}

//
// Insert Ledger
//
func (db *DB) CreateLedger(accountId uint, date time.Time, amount float64, category string, note string) (*Ledger, error) {

	// Get category name.
	catName := strings.Title(strings.ToLower(strings.Trim(category, " ")))

	// Get category id.
	c := LedgerCategory{}
	db.FirstOrCreate(&c, LedgerCategory{Name: catName})

	u := Ledger{AccountId: accountId, Date: date, Amount: amount, Note: note, CategoryId: c.Id}

	if err := db.Create(&u).Error; err != nil {
		services.LogError(err, "")
		return &Ledger{}, err
	}

	// Add in account name
	account, _ := db.GetAccountById(u.AccountId)
	u.AccountName = account.Name

	// Add category name
	cat, _ := db.GetLedgerCategoryById(u.CategoryId)
	u.CategoryName = cat.Name

	services.LogDebug("Created ledger entry on date " + date.UTC().Format("2006-01-02") + " " + account.Name + ".")

	// Return happy
	return &u, nil
}

/* End File */
