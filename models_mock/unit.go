//
// Date: 10/29/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models_mock

import (
  "time"
)

//
// Add Units.
//
func (db *DB) AddUnits(date time.Time, amount float64, note string) error {
  return nil
}

//
// Get the total number of units.
//
func (db *DB) GetUnitsTotalCount() float64 {
  return 15
}

/* End File */