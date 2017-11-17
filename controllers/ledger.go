//
// Date: 10/27/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

//
// Return all ledger items in this system.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/ledgers
//
func (t *Controller) GetLedgers(c *gin.Context) {

	// Return Happy
	c.JSON(http.StatusOK, t.DB.GetAllLedgers())
}

//
// Create a new record from the data passed in.
//
// curl -H "Content-Type: application/json" -X POST -d '{"date": "2017-10-05","amount":1001.12,"account_id":1,"category_name":"Dividends","note":"This is a test note."}' -H "Authorization: Bearer XXXXXX" http://localhost:9090/api/v1/ledgers
//
func (t *Controller) CreateLedger(c *gin.Context) {

	// Grab data from json string.
	body, _ := ioutil.ReadAll(c.Request.Body)
	date := gjson.Get(string(body), "date").String()
	amount := gjson.Get(string(body), "amount").Float()
	account_id := gjson.Get(string(body), "account_id").Uint()
	category_name := gjson.Get(string(body), "category_name").String()
	note := gjson.Get(string(body), "note").String()

	// Reformat date.
	pDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse date."})
		return
	}

	// Store in database & return json.
	ledger, err := t.DB.CreateLedger(uint(account_id), pDate, amount, category_name, note)

	// Return json based on if this was a good result or not.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, ledger)
	}
}

/* End File */
