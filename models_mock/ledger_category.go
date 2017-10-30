//
// Date: 10/29/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models_mock

import (
  "github.com/net-worth-server/models"
)

//
// Get LedgerCategory by id.
//
func (db *DB) GetLedgerCategoryById(id uint) (models.LedgerCategory, error) {
  return models.LedgerCategory{}, nil
}

/* End File */