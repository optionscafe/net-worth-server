//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/optionscafe/net-worth-server/models"
)

type Controller struct {
	DB models.Datastore
}

type ValidateRequest interface {
	Validate() error
}

//
// Validate and Create object.
//
func (t *Controller) ValidateRequest(c *gin.Context, obj ValidateRequest) error {

	// Bind the JSON that got sent into an object and validate.
	if err := c.ShouldBindJSON(obj); err == nil {

		// Run validation
		err := obj.Validate()

		// If we had validation errors return them and do no more.
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err})
			return err
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	return nil
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
