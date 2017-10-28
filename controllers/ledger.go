//
// Date: 10/27/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
  "time"
  "net/http"
  "io/ioutil"
  "github.com/tidwall/gjson" 
)

//
// Return all ledger items in this system.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/ledgers
//
func (t *Controller) GetLedgers(w http.ResponseWriter, r *http.Request) {

  // Return Happy
  t.RespondJSON(w, http.StatusOK, t.DB.GetAllLedgers())
}

//
// Create a new record from the data passed in.
//
// curl -H "Content-Type: application/json" -X POST -d '{"date": "2017-10-05","amount":1001.12,"account_id":1,"note":"This is a test note."}' -H "Authorization: Bearer XXXXXX" http://localhost:9090/api/v1/ledgers
//
func (t *Controller) CreateLedger(w http.ResponseWriter, r *http.Request) {

  // Grab data from json string.
  body, _ := ioutil.ReadAll(r.Body)
  date := gjson.Get(string(body), "date").String()
  amount := gjson.Get(string(body), "amount").Float()
  account_id := gjson.Get(string(body), "account_id").Uint()
  note := gjson.Get(string(body), "note").String()  

  // Reformat date.
  pDate, err := time.Parse("2006-01-02", date)

  if err != nil {
    t.RespondError(w, http.StatusBadRequest, "Unable to parse date.")
    return
  }

  // Store in database & return json.
  ledger, err := t.DB.CreateLedger(uint(account_id), pDate, amount, note); 
  
  if err != nil {
    t.RespondError(w, http.StatusBadRequest, err.Error())
  } else {
    t.RespondJSON(w, http.StatusCreated, ledger)
  }
}

/* End File */