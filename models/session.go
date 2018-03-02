//
// Date: 3/1/2018
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"errors"
	"time"

	"github.com/optionscafe/net-worth-server/library/helpers"
	"github.com/optionscafe/net-worth-server/library/services"
)

type Session struct {
	Id            uint      `gorm:"primary_key" json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	UserId        uint      `sql:"not null;index:UserId"  json:"-"`
	UserAgent     string    `sql:"not null"  json:"-"`
	AccessToken   string    `sql:"not null"  json:"access_token"`
	LastIpAddress string    `sql:"not null"  json:"-"`
	LastActivity  time.Time ` json:"-"`
}

//
// Update Session.
//
func (t *DB) UpdateSession(session *Session) error {
	t.Save(session)
	return nil
}

//
// Get by Access token.
//
func (db *DB) GetByAccessToken(accessToken string) (Session, error) {

	session := Session{}

	if db.First(&session, "access_token = ?", accessToken).RecordNotFound() {
		return Session{}, errors.New("Access Token Not Found - Unable to Authenticate")
	}

	// Return happy
	return session, nil
}

//
// Create new session. A user can have more than one session. Typically it is one session per browser or device.
// We return the session object. The big thing here is we create the access token for this session.
//
func (db *DB) CreateSession(UserId uint, UserAgent string, LastIpAddress string) (Session, error) {

	// Create an access token.
	access_token, err := helpers.GenerateRandomString(50)

	if err != nil {
		services.Error(err)
		return Session{}, err
	}

	// Save the session into the database.
	session := Session{UserId: UserId, UserAgent: UserAgent, AccessToken: access_token, LastIpAddress: LastIpAddress}
	db.Create(&session)

	// Return the session.
	return session, nil
}

/* End File */
