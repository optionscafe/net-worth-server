//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/net-worth-server/controllers"
	"github.com/net-worth-server/cron"
	"github.com/net-worth-server/models"
	"github.com/net-worth-server/services"
)

//
// Main...
//
func main() {

	// Start the db connection.
	db, err := models.NewDB()

	if err != nil {
		services.LogFatal(err, "Failed to connect database")
	}

	// Close db when this app dies. (This might be useless)
	defer db.Close()

	// Setup cron jobs
	cron.AllyStart(db)
	cron.TradierStart(db)
	cron.RobinhoodStart(db)

	// Startup controller
	c := &controllers.Controller{DB: db}

	// Set GIN Settings
	gin.SetMode("release")
	gin.DisableConsoleColor()

	// Set Router
	//r := mux.NewRouter()
	r := gin.Default()

	// Auth middleware
	r.Use(c.AuthMiddleware())
	r.Use(c.ParamValidateMiddleware())

	// Mark routes
	r.GET("/api/v1/marks", c.GetMarks)

	// Ledgers routes
	r.GET("/api/v1/ledgers", c.GetLedgers)
	r.POST("/api/v1/ledgers", c.CreateLedger)

	// Account routes
	r.GET("/api/v1/accounts", c.GetAccounts)
	r.POST("/api/v1/accounts", c.CreateAccount)
	r.GET("/api/v1/accounts/:id", c.GetAccount)
	r.GET("/api/v1/accounts/:id/marks", c.GetAccountMarks)
	r.POST("/api/v1/accounts/:id/marks", c.CreateAccountMark)
	r.POST("/api/v1/accounts/:id/funds", c.AccountManageFunds)

	// Setup http server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("HTTP_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Log this start.
	services.LogInfo("Starting web server on port " + os.Getenv("HTTP_PORT"))

	// Start server and log if fails
	log.Fatal(srv.ListenAndServe())
}

/* End File */
