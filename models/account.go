//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
  "time"
  "errors"
  "github.com/net-worth-server/services"  
)

type Account struct {
  Id uint `gorm:"primary_key" json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Name string `sql:"not null" json:"name"`
  Balance float64 `sql:"type:DECIMAL(12,2)" json:"balance"`
  Units float64 `gorm:"-" json:"units"`    
} 

//
// Get all accounts.
//
func (db *DB) GetAllAcounts() []Account {

  accounts := []Account{}

  db.Order("name asc").Find(&accounts)

  // Struct to capture the sum result.
  type Result struct {
    Sum float64
  }

  // Loop through the accounts and add extra voodoo
  for key, row := range accounts {
    accounts[key].Units = db.getUnits(row.Id)
  }

  // Return the accounts.
  return accounts
}

//
// Get account by id.
//
func (db *DB) GetAccountById(id uint) (Account, error) {

  account := Account{}

  // Find result or send 404
  if err := db.First(&account, id).Error; err != nil {
    return Account{}, err
  }

  // Add in units.
  account.Units = db.getUnits(id)

  // Return the account.
  return account, nil
}

//
// Get account by name.
//
func (db *DB) GetAccountByName(name string) (Account, error) {

  account := Account{}

  // Validate to make sure we do not already have this record.
  if ! db.Where("name = ?", name).First(&account).RecordNotFound() {
    return Account{}, errors.New("We already have an account with the name " + name)
  } 

  // Add in units.
  account.Units = db.getUnits(account.Id)

  // Return the account.
  return account, nil
}

//
// Create account
//
func (db *DB) CreateAccount(account *Account) error {

  // Validate to make sure we do not already have this record.
  _, err := db.GetAccountByName(account.Name)

  if err != nil {
    return err
  }

  // Install the units based on the balance we passed in. - WE have to do this before creation
  db.AddUnits(time.Now(), account.Balance, "Account Opening - " + account.Name) 

  // Store in database.
  if err :=  db.Create(account).Error; err != nil {
    return err
  }  

  // When we create our units are one unit per dollar lets map the units with the dollars  
  db.Create(&AccountUnits{ AccountId: account.Id, Amount: account.Balance, Units: account.Balance, PricePer: 1.00, Note: "Account Opening - " + account.Name })  

  // Mark the value of the account as of today.  
  db.Create(&AccountMarks{ Date: time.Now(), AccountId: account.Id, Balance: account.Balance })  

  // Add in units.
  account.Units = db.getUnits(account.Id)

  // Mark
  db.MarkByDate(time.Now())  

  // Log
  services.LogDebug("Created new accounts - " + account.Name)

  // Return happy.
  return nil
}

//
// Add units to the object.
//
func (db *DB) getUnits(id uint) float64 {

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

//
// Get the total price per unit of all accounts. AKA price per share
//
func (db *DB) GetAccountsPricePerUnit() float64 {

  total := db.GetAccountsTotalBalance()
  totalShares := db.GetUnitsTotalCount()

  if totalShares <= 0 {
    return 0.00
  }

  return total / totalShares
}

//
// Get the balance of all accounts. 
//
func (db *DB) GetAccountsTotalBalance() float64 {

  // Struct to capture the sum result.
  type Result struct {
    Sum float64
  }

  var u Result

  // Query and get unit count.
  db.Raw("SELECT SUM(balance) AS sum FROM accounts").Scan(&u)

  // Return count
  return u.Sum
}

/* End File */