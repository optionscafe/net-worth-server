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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"github.com/optionscafe/net-worth-server/models"
)

//
// Test get a ledger categories. 01
//
func TestGetLedgerCategories01(t *testing.T) {

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	// Create controller
	c := &Controller{DB: db}

	// Setup request
	req, _ := http.NewRequest("GET", "/api/v1/ledger_categories", nil)

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()

	r := gin.New()
	r.GET("/api/v1/ledger_categories", c.GetLedgersCategories)
	r.ServeHTTP(w, req)

	// Grab result and convert to strut
	result := []models.LedgerCategory{}
	err := json.Unmarshal([]byte(w.Body.String()), &result)

	// Test results
	st.Expect(t, err, nil)
	st.Expect(t, result[0].Id, uint(1))
	st.Expect(t, result[0].Name, "Dividends")
	st.Expect(t, result[1].Id, uint(2))
	st.Expect(t, result[1].Name, "Rent Payment")
	st.Expect(t, result[2].Id, uint(3))
	st.Expect(t, result[2].Name, "Other Income")
}

/* End File */
