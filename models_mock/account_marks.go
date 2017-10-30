//
// Date: 10/29/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models_mock

import (
  "time"
  "github.com/net-worth-server/models"
)

//
// Get marks account by id.
//
func (db *DB) GetMarksByAccountById(accountId uint) []models.AccountMarks {
  return []models.AccountMarks{}
}

//
// Create Account Mark. We only allow one per date per account.
// If this mark is already in place we simply update it.
//
func (db *DB) MarkAccountByDate(accountId uint, date time.Time, balance float64) error {
  return nil
}

//
// Get mark by date and account id.
//
func (db *DB) GetMarksByAccountByIdAndDate(accountId uint, date time.Time) (models.AccountMarks, error) {
  return models.AccountMarks{}, nil
}

//
// Add units to the object.
//
func (db *DB) GetMarkAccountUnitsByAccountId(id uint) float64 {
  return 20
}

/* End File */