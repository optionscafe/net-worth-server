//
// Date: 10/19/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"go/build"
	"time"

	_ "github.com/go-sql-driver/mysql"
	env "github.com/jpfuentes2/go-env"
)

// Database interface
type Datastore interface {

	// Generic database functions
	CreateNewRecord(model interface{}, params InsertParam) error
	Count(model interface{}, params QueryParam) (uint, error)
	Query(model interface{}, params QueryParam) error
	QueryWithNoFilterCount(model interface{}, params QueryParam) (int, error)
	GetQueryMetaData(limitCount int, noLimitCount int, params QueryParam) QueryMetaData

	// Applications
	ValidateClientIdGrantType(clientId string, grantType string) (Application, error)

	// Users
	UpdateUser(user *User) error
	GetUserById(id uint) (User, error)
	ValidatePassword(password string) error
	ValidateEmailAddress(email string) error
	GetUserByEmail(email string) (User, error)
	ValidateUserLogin(email string, password string) error
	ValidateCreateUser(first string, last string, email string, password string) error
	LoginUserByEmailPass(email string, password string, appId uint, userAgent string, ipAddress string) (User, error)
	CreateUser(first string, last string, email string, password string, appId uint, userAgent string, ipAddress string) (User, error)

	// Sessions
	UpdateSession(session *Session) error
	GetByAccessToken(accessToken string) (Session, error)
	CreateSession(UserId uint, appId uint, UserAgent string, LastIpAddress string) (Session, error)

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

//
// Start up the controller.
//
func init() {
	// Helpful for testing
	env.ReadEnv(build.Default.GOPATH + "/src/github.com/optionscafe/net-worth-server/.env")
}

/* End File */
