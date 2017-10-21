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
  "github.com/net-worth-server/models"
)

type Controller struct {
  DB models.Datastore
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
// JSON Decoder. This returns a populated object.
//
func (t *Controller) DecodePostedJson(m interface{}, w http.ResponseWriter, r *http.Request) {

  decoder := json.NewDecoder(r.Body)
  
  if err := decoder.Decode(m); err != nil {
    t.RespondError(w, http.StatusBadRequest, err.Error())
    return
  }
  
  r.Body.Close()
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