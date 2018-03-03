//
// Date: 3/1/2018
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"github.com/optionscafe/net-worth-server/models"
	"github.com/tidwall/gjson"
)

//
// Test logging in user.
//
func TestDoOauthToken01(t *testing.T) {

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	// Create controller
	c := &Controller{DB: db}

	// Create a test user.
	db.Create(&models.User{FirstName: "Spicer", LastName: "Matthews", Email: "spicer@cloudmanic.com", Password: "$2a$10$jP2rLOmWtIsMeQcCD..bye6kPADPwLMrz0aOaAPb/4Y4w68VYfJ2m"})

	// Post data
	var postStr = []byte(`{ "username": "spicer@cloudmanic.com", "password": "foobar", "client_id": "Vm4YwgHM2bweuzYeZ", "grant_type": "password" }`)

	// Make a mock request.
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(postStr))
	req.Header.Set("Accept", "application/json")

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()
	r := gin.New()
	r.POST("/login", c.DoOauthToken)
	r.ServeHTTP(w, req)

	// Parse json that returned.
	st.Expect(t, gjson.Get(w.Body.String(), "id").Int(), int64(4))
	st.Expect(t, gjson.Get(w.Body.String(), "first_name").String(), "Spicer")
	st.Expect(t, gjson.Get(w.Body.String(), "last_name").String(), "Matthews")
	st.Expect(t, gjson.Get(w.Body.String(), "email").String(), "spicer@cloudmanic.com")
	st.Expect(t, len(gjson.Get(w.Body.String(), "session.access_token").String()), 50)

	// ------------- Test the login fails for the wrong password. ----------- //

	// Post data
	postStr = []byte(`{ "username": "spicer@cloudmanic.com", "password": "abc123", "client_id": "Vm4YwgHM2bweuzYeZ", "grant_type": "password" }`)

	// Make a mock request.
	req2, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(postStr))
	req2.Header.Set("Accept", "application/json")

	// Setup writer.
	w = httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()
	r = gin.New()
	r.POST("/login", c.DoOauthToken)
	r.ServeHTTP(w, req2)

	// Parse json that returned.
	st.Expect(t, gjson.Get(w.Body.String(), "error").String(), "Sorry, we could not find your account.")
}

//
// Test logging in user. Failed client_id and grant type.
//
func TestDoOauthToken02(t *testing.T) {

	// Start the db connection.
	db, _ := models.NewDB()
	defer db.Close()

	// Create controller
	c := &Controller{DB: db}

	// Create a test user.
	db.Create(&models.User{FirstName: "Spicer", LastName: "Matthews", Email: "spicer@cloudmanic.com", Password: "$2a$10$jP2rLOmWtIsMeQcCD..bye6kPADPwLMrz0aOaAPb/4Y4w68VYfJ2m"})

	// Post data
	var postStr = []byte(`{ "username": "spicer@cloudmanic.com", "password": "foobar", "client_id": "wrong-client-id", "grant_type": "password" }`)

	// Make a mock request.
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(postStr))
	req.Header.Set("Accept", "application/json")

	// Setup writer.
	w := httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()
	r := gin.New()
	r.POST("/login", c.DoOauthToken)
	r.ServeHTTP(w, req)

	// Parse json that returned.
	st.Expect(t, gjson.Get(w.Body.String(), "error").String(), "Invalid client_id or grant type.")

	// ------------- Test the login fails for the wrong grant_type. ----------- //

	// Post data
	postStr = []byte(`{ "username": "spicer@cloudmanic.com", "password": "abc123", "client_id": "Vm4YwgHM2bweuzYeZ", "grant_type": "authorization_code" }`)

	// Make a mock request.
	req2, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(postStr))
	req2.Header.Set("Accept", "application/json")

	// Setup writer.
	w = httptest.NewRecorder()
	gin.SetMode("release")
	gin.DisableConsoleColor()
	r = gin.New()
	r.POST("/login", c.DoOauthToken)
	r.ServeHTTP(w, req2)

	// Parse json that returned.
	st.Expect(t, gjson.Get(w.Body.String(), "error").String(), "Invalid client_id or grant type.")
}

/* End File */
