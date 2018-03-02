//
// Date: 3/1/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package actions

import (
	"fmt"

	"github.com/optionscafe/net-worth-server/models"
)

//
// Create a new user account.
//
// go run main.go -cmd=create-user -first=Spicer -last=Matthews -email=spicer@cloudmanic.com -password=foobar
//
func CreateUserAccount(db *models.DB, first string, last string, email string, password string) {

	// Create user. (note we do not do any validation know what your doing....)
	user, err := db.CreateUser(first, last, email, password, "-cmd=create-user", "127.0.0.1")

	if err != nil {
		panic(err)
	}

	fmt.Println("Success UserId: ", user.Id, " AccessToken: "+user.Session.AccessToken)
}

/* End File */
