//
// Date: 11/12/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	env "github.com/jpfuentes2/go-env"
	"github.com/optionscafe/net-worth-server/models"
	"github.com/tidwall/gjson"
)

//
// Manage funds going into an account. Add funds to an account, or remove.
// Negative values will be removing funds from an account.
//
func TestAccountManageFunds(t *testing.T) {

	// Load config file.
	env.ReadEnv("../.env")

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	// Create controller
	c := &Controller{DB: db}

	// Post data
	var postStr = []byte(`{"amount":500.12, "date": "2017-11-12", "note": "Test Note from TestAccountManageFunds"}`)

	// Make a mock request.
	req, _ := http.NewRequest("POST", "/api/v1/accounts/1/funds", bytes.NewBuffer(postStr))
	req.Header.Set("Accept", "application/json")

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()
	r := gin.New()
	r.Use(func(c *gin.Context) { c.Set("id", uint(1)) })

	r.POST("/api/v1/accounts/:id/funds", c.AccountManageFunds)
	r.ServeHTTP(w, req)

	// Parse json that returned.
	units := gjson.Get(w.Body.String(), "units").Float()
	balance := gjson.Get(w.Body.String(), "balance").Float()
	accountNumber := gjson.Get(w.Body.String(), "account_number").String()

	// Test accountNumber.
	if accountNumber != "7af234fS" {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", "7af234fS", accountNumber)
	}

	// Test units.
	if units != 15178.45 {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 15178.45, units)
	}

	// Test balance.
	if balance != 15178.45 {
		t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", 15178.45, balance)
	}
}

/* End File */
