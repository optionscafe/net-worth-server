//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

// https://github.com/mingrammer/go-todo-rest-api-example/blob/master/app/handler/projects.go (helpful)

package controllers

import (
  "net/http"
  "github.com/networth-server/models"  
)

//
// Return all accounts in this system.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/accounts
//
func (t *Controller) GetAccountsHandler(w http.ResponseWriter, r *http.Request) {

  t.RespondJSON(w, http.StatusOK, t.DB.Find(&[]models.Account{}))
}

//
// Return one account by id.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/accounts/1
//
func (t *Controller) GetAccountHandler(w http.ResponseWriter, r *http.Request) {

  t.RespondById(&models.Account{}, w, r) 
}

/* End File */