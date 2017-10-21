//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
  "time"
)

type AccountUnits struct {
  Id uint `gorm:"primary_key" json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  AccountId uint `sql:"not null;index:UserId" json:"account_id"`
  Amount float64 `sql:"type:DECIMAL(12,2)" json:"amount"`
  Units float64 `sql:"type:DECIMAL(12,2)" json:"units"`
  PricePer float64 `sql:"type:DECIMAL(12,2)" json:"price_per"`  
  Note string `sql:"not null" json:"note"`    
} 

/* End File */