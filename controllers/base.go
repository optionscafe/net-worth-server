//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
  "os"
  "strings"
  "net/http"
  "encoding/json"
  "github.com/jinzhu/gorm"
  "github.com/gorilla/mux"
)

type Controller struct {
  DB *gorm.DB
}

//
// Here we make sure we passed in a proper Bearer Access Token. 
//
func (t *Controller) AuthMiddleware(next http.Handler) http.Handler {

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    // Make sure we have a Bearer token.
    auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

    if len(auth) != 2 || auth[0] != "Bearer" {
      t.RespondError(w, http.StatusUnauthorized, "Authorization Failed (#001)")
      return
    }

    // Make sure we have a known access token.
    if auth[1] != os.Getenv("ACCESS_TOKEN") {
      t.RespondError(w, http.StatusUnauthorized, "Authorization Failed (#002)")
      return
    } 

    // On to next request in the Middleware chain.
    next.ServeHTTP(w, r)

  })

}

//
// Return one record based on the model we pass in. Sort of a generic function
// might not use this function in all API calls.
//
func (t *Controller) RespondById(m interface{}, w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)

  // Find result or send 404
  if err := t.DB.First(m, vars["id"]).Error; err != nil {
    t.RespondError(w, http.StatusNotFound, err.Error())
    return
  }

  // Return happy (json)
  t.RespondJSON(w, http.StatusOK, m)
}

//
// RespondJSON makes the response with payload as json format
//
func (t *Controller) RespondJSON(w http.ResponseWriter, status int, payload interface{}) {

  response, err := json.Marshal(payload)
  
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }
  
  // Return json.
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write([]byte(response))
}

//
// RespondError makes the error response with payload as json format
//
func (t *Controller) RespondError(w http.ResponseWriter, code int, message string) {
  t.RespondJSON(w, code, map[string]string{"error": message})
}

/* End File */