//
// Date: 10/19/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"database/sql/driver"
	"flag"
	"fmt"
	"go/build"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	env "github.com/jpfuentes2/go-env"
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

	// Account Units
	AccountUnitsAddFunds(uint, time.Time, float64, string) error

	// Units
	AddUnits(time.Time, float64, string) error
	GetUnitsTotalCount() float64

	// Ledgers
	GetAllLedgers() []Ledger
	CreateLedger(uint, Date, float64, string, string) (*Ledger, error)

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

// Used for date formatting with json conversion.
type Date struct {
	time.Time
}

//
// Start up the controller.
//
func init() {
	// Helpful for testing
	env.ReadEnv(build.Default.GOPATH + "/src/github.com/optionscafe/net-worth-server/.env")
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
	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+"/"+dbName+"?charset=utf8&parseTime=True&loc=UTC")

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
	ds := Date{time.Date(2017, 10, 29, 17, 20, 01, 507451, time.UTC)}
	totalUnits := 14678.33 + 85345.33 + 5000.00 + 4501.02

	// Accounts
	db.Exec("TRUNCATE TABLE accounts;")
	db.Create(&Account{Name: "Tradier", Balance: 14678.33, Units: 14678.33, AccountNumber: "7af234fS"})
	db.Create(&Account{Name: "Ally", Balance: 85345.33, Units: 85345.33, AccountNumber: "23423499"})
	db.Create(&Account{Name: "Lending Club", Balance: 5000.00, Units: 5000.00, AccountNumber: "1238888"})
	db.Create(&Account{Name: "Robinhood", Balance: 4501.02, Units: 4501.02, AccountNumber: "72294Daf33"})

	// Add  Units
	db.Exec("TRUNCATE TABLE units;")
	db.Create(&Unit{Date: ts, Amount: totalUnits, Units: totalUnits, PricePer: 1.00, Note: "Test Note #1"})

	// Add Account Units
	db.Exec("TRUNCATE TABLE account_units;")
	db.Create(&AccountUnits{AccountId: 1, Date: ts, Amount: 14678.33, Units: 14678.33, PricePer: 1.00, Note: "Test Note #1"})
	db.Create(&AccountUnits{AccountId: 2, Date: ts, Amount: 85345.33, Units: 85345.33, PricePer: 1.00, Note: "Test Note #2"})
	db.Create(&AccountUnits{AccountId: 3, Date: ts, Amount: 5000.00, Units: 5000.00, PricePer: 1.00, Note: "Test Note #3"})
	db.Create(&AccountUnits{AccountId: 4, Date: ts, Amount: 4501.02, Units: 4501.02, PricePer: 1.00, Note: "Test Note #4"})

	// Add Account Marks
	db.Exec("TRUNCATE TABLE account_marks;")
	db.Create(&AccountMarks{Date: ts, AccountId: 1, Units: 14678.33, PricePer: 1.00, Balance: 14678.33})
	db.Create(&AccountMarks{Date: ts, AccountId: 2, Units: 85345.33, PricePer: 1.00, Balance: 85345.33})
	db.Create(&AccountMarks{Date: ts, AccountId: 3, Units: 5000.00, PricePer: 1.00, Balance: 5000.00})
	db.Create(&AccountMarks{Date: ts, AccountId: 4, Units: 4501.02, PricePer: 1.00, Balance: 4501.02})

	// Ledger Categories
	db.Exec("TRUNCATE TABLE ledger_categories;")
	db.Create(&LedgerCategory{Name: "Dividends"})
	db.Create(&LedgerCategory{Name: "Rent Payment"})
	db.Create(&LedgerCategory{Name: "Other Income"})

	// Ledgers
	db.Exec("TRUNCATE TABLE ledgers;")
	db.Create(&Ledger{AccountId: 1, Date: ds, Amount: 55.45, Note: "1st ledger test.", CategoryId: 1})
	db.Create(&Ledger{AccountId: 2, Date: ds, Amount: 1155.45, Note: "2nd ledger test.", CategoryId: 2})
	db.Create(&Ledger{AccountId: 3, Date: ds, Amount: 155.45, Note: "3rd ledger test.", CategoryId: 3})
	db.Create(&Ledger{AccountId: 4, Date: ds, Amount: 455.00, Note: "4th ledger test.", CategoryId: 1})

}

//
// Convert JSON string to a date. Format XXXX-XX-XX
//
func (t *Date) UnmarshalJSON(b []byte) error {

	// Remove quotes
	str := strings.Replace(string(b), "\"", "", -1)

	// Parse string
	tt, _ := time.Parse("2006-01-02", str)

	// Return UTC
	*t = Date{tt.UTC()}
	return nil
}

//
// Convert JSON string to a date. Make it so when we create JSON we return this format XXXX-XX-XX
//
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))), nil
}

//
// Format time Date Going into the DB
//
func (t Date) Value() (driver.Value, error) {
	return t.Format("2006-01-02"), nil
}

//
// Convert to type Date Coming out of the DB
//
func (t *Date) Scan(value interface{}) error {

	// Parse string
	tt, _ := time.Parse("2006-01-02 03:04:05 -0700 MST", fmt.Sprintf("%s", value))

	// Return UTC
	*t = Date{tt.UTC()}
	return nil
}

/* End File */
