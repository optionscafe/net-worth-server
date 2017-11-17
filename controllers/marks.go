//
// Date: 10/20/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//
// Return all marks in this system.
//
// curl -H "Authorization: Bearer XXXXX" http://localhost:9090/api/v1/marks
//
func (t *Controller) GetMarks(c *gin.Context) {

	// Return Happy
	c.JSON(http.StatusOK, t.DB.GetAllMarks())
}

/* End File */
