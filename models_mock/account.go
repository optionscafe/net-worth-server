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
// Get all accounts.
//
func (db *DB) GetAllAcounts() []models.Account {
  return []models.Account{}
}

//
// Get account by id.
//
func (db *DB) GetAccountById(id uint) (models.Account, error) {
  return models.Account{}, nil
}

//
// Get account by name.
//
func (db *DB) GetAccountByName(name string) (models.Account, error) {
  return models.Account{}, nil
}

//
// Create account
//
func (db *DB) CreateAccount(account *models.Account) error {
  return nil
}

//
// Get the total price per unit of all accounts. AKA price per share
//
func (db *DB) GetAccountsPricePerUnit() float64 {
  return 15
}

//
// Get the balance of all accounts. 
//
func (db *DB) GetAccountsTotalBalance() float64 {
  return 20
}

/* End File */