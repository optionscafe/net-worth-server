//
// Date: 3/1/2018
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/optionscafe/net-worth-server/library/realip"
	"github.com/optionscafe/net-worth-server/library/services"
	"github.com/tidwall/gjson"
)

//
// Login to account.
//
func (t *Controller) DoLogin(c *gin.Context) {

	// Get vars we post in.
	body, _ := ioutil.ReadAll(c.Request.Body)
	email := gjson.Get(string(body), "email").String()
	password := gjson.Get(string(body), "password").String()

	// Close body.
	defer c.Request.Body.Close()

	// Validate user.
	if err := t.DB.ValidateUserLogin(email, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login user in by email and password
	user, err := t.DB.LoginUserByEmailPass(email, password, c.Request.UserAgent(), realip.RealIP(c.Request))

	if err != nil {
		services.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sorry, we could not find your account."})
		return
	}

	// Return success json.
	c.JSON(200, user)
}

/* End File */
