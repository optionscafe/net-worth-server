//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/optionscafe/net-worth-server/models"
	"github.com/tidwall/gjson"
)

//
// Return all accounts in this system.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/accounts
//
func (t *Controller) GetAccounts(c *gin.Context) {

	// Return Happy
	c.JSON(http.StatusOK, t.DB.GetAllAcounts())
}

//
// Return one account by id.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/accounts/1
//
func (t *Controller) GetAccount(c *gin.Context) {

	// Get the account id.
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse account id."})
		return
	}

	// Get the account by id.
	account, err := t.DB.GetAccountById(uint(idInt))

	// Return json based on if this was a good result or not.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, account)
	}
}

//
// Create a new record from the data passed in.
//
// curl -H "Content-Type: application/json" -X POST -d '{"name":"Lending Club","balance":70123.66}' -H "Authorization: Bearer XXXXXX" http://localhost:9090/api/v1/accounts
//
func (t *Controller) CreateAccount(c *gin.Context) {

	account := models.Account{}

	// Decode the json we posted in.
	if err := t.DecodePostedJson(&account, c.Writer, c.Request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Store in database & return json.
	if err := t.DB.CreateAccount(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// Get fresh account because of timezone issues, and add in all the other magic.
		account, _ := t.DB.GetAccountById(account.Id)
		c.JSON(http.StatusCreated, account)
	}
}

//
// Get all the marks for an account.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/accounts/XX/marks
//
func (t *Controller) GetAccountMarks(c *gin.Context) {

	// Get the account id.
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse account id."})
		return
	}

	// Return Happy
	c.JSON(http.StatusOK, t.DB.GetMarksByAccountById(uint(idInt)))
}

//
// Mark account value for today.
//
// curl -H "Content-Type: application/json" -X POST -d '{"balance":1000.00, "date": "2017-10-05"}' -H "Authorization: Bearer XXXXXX" http://localhost:9090/api/v1/accounts/XX/marks
//
func (t *Controller) CreateAccountMark(c *gin.Context) {

	// Get the account id.
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse account id."})
		return
	}

	id := uint(idInt)

	body, _ := ioutil.ReadAll(c.Request.Body)
	date := gjson.Get(string(body), "date").String()
	balance := gjson.Get(string(body), "balance").Float()

	// Reformat date.
	pDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse date."})
		return
	}

	// Store in database & return json.
	err = t.DB.MarkAccountByDate(id, pDate.UTC(), balance)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong please email support"})
		return
	}

	// Get the account by id.
	account, err := t.DB.GetAccountById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, account)
	}
}

//
// Manage funds going into an account. Add funds to an account, or remove.
// Negative values will be removing funds from an account.
//
// curl -H "Content-Type: application/json" -X POST -d '{"amount":500.12, "date": "2017-11-12", "note": "Test note."}' -H "Authorization: Bearer XXXXXX" http://localhost:9090/api/v1/accounts/XX/funds
//
func (t *Controller) AccountManageFunds(c *gin.Context) {

	// Get the account id.
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse account id."})
		return
	}

	body, _ := ioutil.ReadAll(c.Request.Body)
	date := gjson.Get(string(body), "date").String()
	amount := gjson.Get(string(body), "amount").Float()
	note := gjson.Get(string(body), "note").String()

	// Reformat date.
	pDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse date."})
		return
	}

	// Add / Subtract money to an account units
	t.DB.AccountUnitsAddFunds(uint(idInt), pDate, amount, note)

	// Get the account by id.
	account, err := t.DB.GetAccountById(uint(idInt))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, account)
	}
}

/* End File */
