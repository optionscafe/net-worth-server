//
// Date: 10/27/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/optionscafe/net-worth-server/models"
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

	ledger := models.Ledger{}

	// Here we parse the JSON sent in, assign it to a struct, set validation errors if any.
	if t.ValidateRequest(c, &ledger) != nil {
		return
	}

	// Store in database & return json.
	newLedger, err := t.DB.CreateLedger(ledger.AccountId, ledger.Date, ledger.Amount, ledger.CategoryName, ledger.Note)

	// Return json based on if this was a good result or not.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, newLedger)
	}
}

/* End File */
