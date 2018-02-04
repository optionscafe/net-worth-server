//
// Date: 10/27/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"time"
	//"github.com/optionscafe/net-worth-server/services"
)

type LedgerCategory struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `sql:"not null" json:"Name"`
}

//
// Get LedgerCategory by id.
//
func (db *DB) GetLedgerCategoryById(id uint) (LedgerCategory, error) {

	r := LedgerCategory{}

	// Find result or send error
	if err := db.First(&r, id).Error; err != nil {
		return LedgerCategory{}, err
	}

	// Return the LedgerCategory.
	return r, nil
}

/* End File */
