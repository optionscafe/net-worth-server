//
// Date: 10/20/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
  "time"
  "errors"
  "github.com/net-worth-server/services"
)

type Mark struct {
  Id uint `gorm:"primary_key" json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Date time.Time `gorm:"type:date" json:"date"`
  Units float64 `sql:"type:DECIMAL(12,2)" json:"units"`
  PricePer float64 `sql:"type:DECIMAL(12,2)" json:"price_per"`  
  Balance float64 `sql:"type:DECIMAL(12,2)" json:"balance"`     
} 

//
// Get all marks.
//
func (db *DB) GetAllMarks() []Mark {

  marks := []Mark{}

  db.Order("date desc").Find(&marks)

  return marks
}

//
// Mark account by date.
//
func (db *DB) MarkByDate(date time.Time) error {

  // Get the total balance of all accounts.
  balance := db.GetAccountsTotalBalance()

  // Get current share price for all accounts
  pricePer := db.GetAccountsPricePerUnit()

  // Get total number of units for all accounts.
  units := db.GetUnitsTotalCount()

  // Check to see if we have already marked today. 
  mark, err := db.GetMarkByDate(date)

  // Setup new mark obj & Store in database.
  if err != nil {

    u := Mark{ Date: date.UTC(), Units: units, PricePer: pricePer, Balance: balance }

    if err :=  db.Create(&u).Error; err != nil {
      services.LogError(err, "")     
      return err
    }

    services.LogDebug("Marked new mark for " + date.UTC().Format("2006-01-02"))

  } else {
    
    mark.Units = units
    mark.PricePer = pricePer
    mark.Balance = balance

    if err := db.Save(&mark).Error; err != nil {
      services.LogError(err, "")
      return err
    }

    services.LogDebug("Marked updated for " + date.UTC().Format("2006-01-02"))    

  }

  // Return happy
  return nil
}

//
// Get mark by date.
//
func (db *DB) GetMarkByDate(date time.Time) (*Mark, error) {

  m := Mark{}

  if db.Where("date = ?", date.UTC().Format("2006-01-02")).First(&m).RecordNotFound() {
    return &Mark{}, errors.New("No Mark record found for " + date.UTC().Format("2006-01-02"))
  }    

  // Return happy
  return &m, nil
}

/* End File */