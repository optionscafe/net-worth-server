//
// Date: 3/1/2018
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"github.com/gin-gonic/gin"
)

//
// Do Routes
//
func (t *Controller) DoRoutes(r *gin.Engine) {

	// --------- API V1 sub-routes ----------- //

	apiV1 := r.Group("/api/v1")

	apiV1.Use(t.AuthMiddleware())
	{
		// Mark routes
		apiV1.GET("/marks", t.GetMarks)

		// Ledgers routes
		apiV1.GET("/ledgers", t.GetLedgers)
		apiV1.POST("/ledgers", t.CreateLedger)

		// Account routes
		apiV1.GET("/accounts", t.GetAccounts)
		apiV1.POST("/accounts", t.CreateAccount)
		apiV1.GET("/accounts/:id", t.GetAccount)
		apiV1.GET("/accounts/:id/marks", t.GetAccountMarks)
		apiV1.POST("/accounts/:id/marks", t.CreateAccountMark)
		apiV1.POST("/accounts/:id/funds", t.AccountManageFunds)
	}

	// ------------ Non-Auth Routes ------ //

	// // Auth Routes
	r.POST("/oauth/token", t.DoOauthToken)
}
