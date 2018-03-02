//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package main

import (
	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/optionscafe/net-worth-server/cmd"
	"github.com/optionscafe/net-worth-server/controllers"
	"github.com/optionscafe/net-worth-server/cron"
	"github.com/optionscafe/net-worth-server/models"
	"github.com/optionscafe/net-worth-server/services"
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

	// See if this a command. If so run the command and do not start the app.
	status := cmd.Run(db)

	if status == true {
		return
	}

	// Close db when this app dies. (This might be useless)
	defer db.Close()

	// Setup cron jobs
	cron.AllyStart(db)
	cron.TradierStart(db)
	cron.RobinhoodStart(db)

	// Startup controller
	c := &controllers.Controller{DB: db}

	// Start webserver & controllers
	c.StartWebServer()
}

/* End File */
