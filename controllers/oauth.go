//
// Date: 3/1/2018
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/optionscafe/net-worth-server/library/realip"
	"github.com/optionscafe/net-worth-server/library/services"
	"github.com/tidwall/gjson"
)

//
// Login to account and get a token.
//
func (t *Controller) DoOauthToken(c *gin.Context) {
	var username string
	var password string
	var grantType string
	var clientId string

	// A special case to handle clients that do not post in via JSON (looking at you PAW)
	if strings.Contains(c.Request.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		username = c.PostForm("username")
		password = c.PostForm("password")
		grantType = c.PostForm("grant_type")
		clientId = c.PostForm("client_id")
	} else {
		body, _ := ioutil.ReadAll(c.Request.Body)
		username = gjson.Get(string(body), "username").String()
		password = gjson.Get(string(body), "password").String()
		grantType = gjson.Get(string(body), "grant_type").String()
		clientId = gjson.Get(string(body), "client_id").String()
	}

	// Close body.
	defer c.Request.Body.Close()

	// First we validate the grant type and client id. Make sure this is a known application.
	app, err := t.DB.ValidateClientIdGrantType(clientId, grantType)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client_id or grant type."})
		return
	}

	// Validate user.
	if err := t.DB.ValidateUserLogin(username, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login user in by email and password
	user, err := t.DB.LoginUserByEmailPass(username, password, app.Id, c.Request.UserAgent(), realip.RealIP(c.Request))

	if err != nil {
		services.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sorry, we could not find your account."})
		return
	}

	// Return success json.
	c.JSON(200, gin.H{"access_token": user.Session.AccessToken, "token_type": "bearer"})
}

/* End File */
