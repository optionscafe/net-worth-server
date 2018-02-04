//
// Date: 11/8/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

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
// Capture parms out of URL and save them to context. This is useful because we validate integers
// in urls instead of passing what could be a string into an SQL function.
//
func (t *Controller) ParamValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Capture id...
		if len(c.Param("id")) > 0 {

			// Validate input and cast it to a uint.
			id, err := strconv.ParseUint(c.Param("id"), 10, 32)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "The id passed in via the URL is not an integer."})
				return
			}

			c.Set("id", uint(id))
		}

		// On to the next middleware or the controller.
		c.Next()
	}
}

/* End File */
