/**
 * @Author: Spicer Matthews
 * @Date:   03/14/2018
 * @Email:  spicer@cloudmanic.com
 * @Last modified by:   Spicer Matthews
 * @Last modified time: 03/14/2018
 * @Copyright: 2018 Cloudmanic Labs, LLC. All rights reserved.
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/optionscafe/net-worth-server/models"
)

//
// Return all ledger categories items in this system.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/ledger_categories
//
func (t *Controller) GetLedgersCategories(c *gin.Context) {
	// Place to store the results.
	var results = []models.LedgerCategory{}

	// Run the query
	err := t.DB.Query(&results, models.QueryParam{})

	// Return json based on if this was a good result or not.
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There was an error. Please contact your admin."})
	} else {
		c.JSON(http.StatusOK, results)
	}
}

/* End File */
