//
// Date: 10/30/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/tidwall/gjson"  
  "github.com/jpfuentes2/go-env"  
  "github.com/net-worth-server/models"
)

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
  c := &Controller{ DB: db } 

  // Set the expected
  expectedAccountNames := []string{ "Tradier", "E*Trade", "Lending Club", "Lending Club" }
  expectedCategoryNames := []string{ "Dividends", "Rent Payment", "Other Income", "Dividends" } 
  expectedAmounts := []float64{ 55.45, 1155.45, 155.45, 455 }
  expectedNotes := []string{ "1st ledger test.", "2nd ledger test.", "3rd ledger test.", "4th ledger test." } 

  // Make a mock request.
  rec := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/api/v1/ledgers", nil)
  http.HandlerFunc(c.GetLedgers).ServeHTTP(rec, req)

  // Parse json that returned.
  result := gjson.Parse(rec.Body.String())  

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
    if "2017-10-29T00:00:00Z" != date {
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