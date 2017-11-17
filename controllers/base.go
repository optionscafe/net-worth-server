//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/net-worth-server/models"
)

type Controller struct {
	DB models.Datastore
}

//
// Here we make sure we passed in a proper Bearer Access Token.
//
func (t *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Make sure we have a Bearer token.
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization Failed (#001)"})
			return
		}

		// Make sure we have a known access token.
		if auth[1] != os.Getenv("ACCESS_TOKEN") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization Failed (#002)"})
			return
		}

		// On to next request in the Middleware chain.
		c.Next()
	}
}

//
// Convert too. If for some reason this fails we return zero, and abort.
// This is done this way to keep the code clean on the controller side.
// For example if someone passes in "asdfasdf" instead of an integer this returns
// zero. Then most likely zero would be passed to the function that calls a mysql query.
// Since this is mostly used for "ids" the query will return no results. Better to pass zero
// to a query than some random string.
//
func (t *Controller) ConvertUrlParamToUint(parm string, c *gin.Context) uint {

	// Convert string to int.
	p, err := strconv.ParseUint(c.Param(parm), 10, 32)

	if err != nil {
		return 0
	}

	return uint(p)
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

// //
// // RespondJSON makes the response with payload as json format
// //
// func (t *Controller) RespondJSON(w http.ResponseWriter, status int, payload interface{}) {

//   response, err := json.Marshal(payload)

//   if err != nil {
//     w.WriteHeader(http.StatusInternalServerError)
//     w.Write([]byte(err.Error()))
//     return
//   }

//   // Return json.
//   w.Header().Set("Content-Type", "application/json")
//   w.WriteHeader(status)
//   w.Write([]byte(response))
// }

// //
// // RespondError makes the error response with payload as json format
// //
// func (t *Controller) RespondError(w http.ResponseWriter, code int, message string) {
//   t.RespondJSON(w, code, map[string]string{"error": message})
// }

/* End File */
