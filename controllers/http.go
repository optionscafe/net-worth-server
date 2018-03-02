//
// Date: 3/1/2018
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/optionscafe/net-worth-server/services"
)

//
// Start the webserver
//
func (t *Controller) StartWebServer() {

	// Set GIN Settings
	gin.SetMode("release")
	gin.DisableConsoleColor()

	// Set Router
	router := gin.New()

	// Logger - Global middleware
	if os.Getenv("HTTP_LOG_REQUESTS") == "true" {
		router.Use(gin.Logger())
	}

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// CORS Middleware - Global middleware
	router.Use(t.CorsMiddleware())

	// Register Routes
	t.DoRoutes(router)

	// Setup http server
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + os.Getenv("HTTP_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Log this start.
	services.LogInfo("Starting web server at http://localhost:" + os.Getenv("HTTP_PORT"))

	// Start server and log if fails
	log.Fatal(srv.ListenAndServe())
}

/* end File */
