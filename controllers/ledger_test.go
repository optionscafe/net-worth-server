//
// Date: 10/30/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jpfuentes2/go-env"
	"github.com/nbio/st"
	"github.com/optionscafe/net-worth-server/models"
	"github.com/tidwall/gjson"
)

//
// Test create a ledger. 01
//
func TestCreateLedger01(t *testing.T) {

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	// Create controller
	c := &Controller{DB: db}

	// Post data
	var postStr = []byte(`{"date": "2017-10-05","amount":1001.12,"account_id":1,"category_name":"Dividends","note":"This is a test note.", "symbol": "bac"}`)

	// Setup request
	req, _ := http.NewRequest("POST", "/api/v1/ledgers", bytes.NewBuffer(postStr))

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()

	r := gin.New()
	r.POST("/api/v1/ledgers", c.CreateLedger)
	r.ServeHTTP(w, req)

	// Grab result and convert to strut
	result := models.Ledger{}
	err := json.Unmarshal([]byte(w.Body.String()), &result)

	// Test results
	st.Expect(t, err, nil)
	st.Expect(t, result.Id, uint(5))
	st.Expect(t, result.Date.Format("2006-01-02"), "2017-10-05")
	st.Expect(t, result.AccountName, "Tradier")
	st.Expect(t, result.Symbol, "BAC")
	st.Expect(t, result.CategoryName, "Dividends")
	st.Expect(t, result.Amount, 1001.12)
	st.Expect(t, result.Note, "This is a test note.")
}

//
// Test create a ledger. 02 - Test in valid date.
//
func TestCreateLedger02(t *testing.T) {

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	// Create controller
	c := &Controller{DB: db}

	// Post data
	var postStr = []byte(`{"date": "10/5/2017","amount":1001.12,"account_id":1,"category_name":"Dividends","note":"This is a test note."}`)

	// Setup request
	req, _ := http.NewRequest("POST", "/api/v1/ledgers", bytes.NewBuffer(postStr))

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()

	r := gin.New()
	r.POST("/api/v1/ledgers", c.CreateLedger)
	r.ServeHTTP(w, req)

	// Get error string
	errorStr := gjson.Get(w.Body.String(), "errors.date").String()

	// Test results
	st.Expect(t, errorStr, "Date field must be in XXXX-XX-XX format.")
}

//
// Return all ledger items in this system.
//
func TestGetLedgers(t *testing.T) {

	// Load config file.
	env.ReadEnv("../.env")

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	// Create controller
	c := &Controller{DB: db}

	// Set the expected
	expectedAccountNames := []string{"Tradier", "Ally", "Lending Club", "Robinhood"}
	expectedCategoryNames := []string{"Dividends", "Rent Payment", "Other Income", "Dividends"}
	expectedAmounts := []float64{55.45, 1155.45, 155.45, 455}
	expectedNotes := []string{"1st ledger test.", "2nd ledger test.", "3rd ledger test.", "4th ledger test."}

	// Setup request
	req, _ := http.NewRequest("GET", "/api/v1/ledgers", nil)

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()
	r := gin.New()
	r.GET("/api/v1/ledgers", c.GetLedgers)
	r.ServeHTTP(w, req)

	// Parse json that returned.
	result := gjson.Parse(w.Body.String())

	// Index each loop
	loop := 0

	// Loop through and build rows of output table.
	result.ForEach(func(key, value gjson.Result) bool {

		// Get values from json
		id := gjson.Get(value.String(), "id").Int()
		account_name := gjson.Get(value.String(), "account_name").String()
		category_name := gjson.Get(value.String(), "category_name").String()
		amount := gjson.Get(value.String(), "amount").Float()
		date := gjson.Get(value.String(), "date").String()
		note := gjson.Get(value.String(), "note").String()

		// Test id.
		if (loop + 1) != int(id) {
			t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", (loop + 1), id)
		}

		// Test account name.
		if expectedAccountNames[loop] != account_name {
			t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expectedAccountNames[loop], account_name)
		}

		// Test category_name.
		if expectedCategoryNames[loop] != category_name {
			t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expectedCategoryNames[loop], category_name)
		}

		// Test amount.
		if expectedAmounts[loop] != amount {
			t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expectedAmounts[loop], amount)
		}

		// Test date.
		if "2017-10-29" != date {
			t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", "2017-10-29T00:00:00Z", date)
		}

		// Test note.
		if expectedNotes[loop] != note {
			t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expectedNotes[loop], note)
		}

		// Up the loop count.
		loop++

		// keep iterating
		return true
	})

}

/* End File */
