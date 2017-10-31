//
// Date: 10/19/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
  "os"
  "flag" 
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

// Database interface
type Datastore interface {

  // Accounts
  GetAllAcounts() []Account
  GetAccountById(uint) (Account, error)
  GetAccountByName(string) (Account, error)
  CreateAccount(*Account) error
  GetAccountsTotalBalance() float64
  GetAccountsPricePerUnit() float64

  // Account Marks
  MarkAccountByDate(uint, time.Time, float64) error
  GetMarksByAccountById(uint) []AccountMarks
  GetMarksByAccountByIdAndDate(uint, time.Time) (AccountMarks, error)

  // Units
  AddUnits(time.Time, float64, string) error
  GetUnitsTotalCount() float64

  // Ledgers
  GetAllLedgers() []Ledger
  CreateLedger(uint, time.Time, float64, string, string) (*Ledger, error)   

  // LedgerCategory
  GetLedgerCategoryById(uint) (LedgerCategory, error)

  // Marks
  GetAllMarks() []Mark
  MarkByDate(time.Time) error
  GetMarkByDate(time.Time) (*Mark, error)
  GetMarkAccountUnitsByAccountId(uint) float64
  
}

type DB struct {
  *gorm.DB
}

//
// Setup the db connection.
//
func NewDB() (*DB, error) {

  dbName := os.Getenv("DB_DATABASE") 

  // Is this a testing run?
  if flag.Lookup("test.v") != nil {
    dbName = os.Getenv("DB_DATABASE_TESTING") 
  }

  // Connect to Mysql
  db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/" + dbName + "?charset=utf8&parseTime=True&loc=UTC")  

  
  if err != nil {
    return nil, err
  }

  // Ping make sure the server is up.
  if err = db.DB().Ping(); err != nil {
    return nil, err
  }  

  // Enable
  //db.LogMode(true)
  //db.SetLogger(log.New(os.Stdout, "\r\n", 0))  

  // Run migrations
  db.AutoMigrate(&Unit{})
  db.AutoMigrate(&Mark{}) 
  db.AutoMigrate(&Ledger{})    
  db.AutoMigrate(&Account{})
  db.AutoMigrate(&AccountMarks{})
  db.AutoMigrate(&AccountUnits{})
  db.AutoMigrate(&LedgerCategory{})

  // Is this a testing run? If so load testing data.
  if flag.Lookup("test.v") != nil {
    LoadTestingData(db)
  }

  // Return db connection.
  return &DB{db}, nil
}

//
// Load testing data.
//
func LoadTestingData(db *gorm.DB) {

  // Shared time we use.
  ts := time.Date(2017, 10, 29, 17, 20, 01, 507451, time.UTC)

  // Accounts
  db.Exec("TRUNCATE TABLE accounts;")
  db.Create(&Account{ Name: "Tradier", Balance: 14678.33, Units: 14678.33 })
  db.Create(&Account{ Name: "E*Trade", Balance: 85345.33, Units: 85345.33 })
  db.Create(&Account{ Name: "Lending Club", Balance: 5000.00, Units: 5000.00 })  

  // Ledger Categories
  db.Exec("TRUNCATE TABLE ledger_categories;")  
  db.Create(&LedgerCategory{ Name: "Dividends" })  
  db.Create(&LedgerCategory{ Name: "Rent Payment" })
  db.Create(&LedgerCategory{ Name: "Other Income" })

  // Ledgers
  db.Exec("TRUNCATE TABLE ledgers;")  
  db.Create(&Ledger{ AccountId: 1, Date: ts, Amount: 55.45, Note: "1st ledger test.", CategoryId: 1 })
  db.Create(&Ledger{ AccountId: 2, Date: ts, Amount: 1155.45, Note: "2nd ledger test.", CategoryId: 2 })
  db.Create(&Ledger{ AccountId: 3, Date: ts, Amount: 155.45, Note: "3rd ledger test.", CategoryId: 3 })
  db.Create(&Ledger{ AccountId: 3, Date: ts, Amount: 455.00, Note: "4th ledger test.", CategoryId: 1 })    

}

/* End File */