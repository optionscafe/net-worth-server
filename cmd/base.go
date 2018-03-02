//
// Date: 3/1/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package cmd

import (
	"flag"
	"fmt"

	"github.com/optionscafe/net-worth-server/cmd/actions"
	"github.com/optionscafe/net-worth-server/models"
)

//
// Run this and see if we have any commands to run.
//
func Run(db *models.DB) bool {

	// Grab flags
	action := flag.String("cmd", "none", "")
	first := flag.String("first", "", "")
	last := flag.String("last", "", "")
	email := flag.String("email", "", "")
	password := flag.String("password", "", "")
	flag.Parse()

	switch *action {

	// Create a new user from the CLI
	case "create-user":
		actions.CreateUserAccount(db, *first, *last, *email, *password)
		return true
		break

	// Just a test
	case "test":
		fmt.Println("CMD Works....")
		return true
		break

	}

	return false
}

/* End File */
