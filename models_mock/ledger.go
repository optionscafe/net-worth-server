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
// Get all marks.
//
func (db *DB) GetAllLedgers() []models.Ledger {
  return []models.Ledger{}
}

//
// Insert Ledger
//
func (db *DB) CreateLedger(accountId uint, date time.Time, amount float64, category string, note string) (*models.Ledger, error) {
  return &models.Ledger{}, nil
}

/* End File */