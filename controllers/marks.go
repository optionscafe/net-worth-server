//
// Date: 10/20/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
  "net/http"
)

//
// Return all marks in this system.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/marks
//
func (t *Controller) GetMarks(w http.ResponseWriter, r *http.Request) {

  // Return Happy
  t.RespondJSON(w, http.StatusOK, t.DB.GetAllMarks())
}

/* End File */