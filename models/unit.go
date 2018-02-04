//
// Date: 10/20/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"time"
)

type Unit struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Date      time.Time `gorm:"type:date" json:"date"`
	Amount    float64   `sql:"type:DECIMAL(12,2)" json:"amount"`
	Units     float64   `sql:"type:DECIMAL(12,2)" json:"units"`
	PricePer  float64   `sql:"type:DECIMAL(12,2)" json:"price_per"`
	Note      string    `sql:"not null" json:"note"`
}

//
// Add Units.
//
func (db *DB) AddUnits(date time.Time, amount float64, note string) error {

	var pp float64
	var units float64

	// Get the per unit price.
	pricePerUnit := db.GetAccountsPricePerUnit()

	// Figure out how many units we are adding here.
	if pricePerUnit <= 0 {
		units = amount
		pp = 1.00
	} else {
		pp = pricePerUnit
		units = amount / pricePerUnit
	}

	// Setup the unit object
	u := Unit{Date: date, Amount: amount, Units: units, PricePer: pp, Note: note}

	// Store in database.
	if err := db.Create(&u).Error; err != nil {
		return err
	}

	// Return happy
	return nil
}

//
// Get the total number of units.
//
func (db *DB) GetUnitsTotalCount() float64 {

	// Struct to capture the sum result.
	type Result struct {
		Sum float64
	}

	var u Result

	// Query and get count.
	db.Raw("SELECT SUM(units) AS sum FROM units").Scan(&u)

	// Return count
	return u.Sum
}

/* End File */
