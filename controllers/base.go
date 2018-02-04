//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/optionscafe/net-worth-server/models"
)

type Controller struct {
	DB models.Datastore
}

//
// JSON Decoder. This returns a populated object.
//
func (t *Controller) DecodePostedJson(m interface{}, w http.ResponseWriter, r *http.Request) error {

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(m); err != nil {
		return err
	}

	r.Body.Close()

	// Return happy.
	return nil
}

/* End File */
