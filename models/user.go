/**
 * @Author: Spicer Matthews
 * @Date:   03/01/2018
 * @Email:  spicer@cloudmanic.com
 * @Last modified by:   Spicer Matthews
 * @Last modified time: 03/13/2018
 * @Copyright: 2018 Cloudmanic Labs, LLC. All rights reserved.
 */

//
// Date: 3/1/2018
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"errors"
	"html/template"
	"time"

	"github.com/optionscafe/net-worth-server/library/checkmail"
	"github.com/optionscafe/net-worth-server/library/services"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	FirstName string    `sql:"not null" json:"first_name"`
	LastName  string    `sql:"not null" json:"last_name"`
	Email     string    `sql:"not null" json:"email"`
	Password  string    `sql:"not null" json:"-"`
	Session   Session   `json:"session"`
}

//
// Update user.
//
func (t *DB) UpdateUser(user *User) error {
	t.Save(user)
	return nil
}

//
// Get a user by Id.
//
func (t *DB) GetUserById(id uint) (User, error) {

	var u User

	if t.Where("Id = ?", id).First(&u).RecordNotFound() {
		return u, errors.New("Record not found")
	}

	// Return the user.
	return u, nil
}

//
// Get a user by email.
//
func (db *DB) GetUserByEmail(email string) (User, error) {

	var u User

	if db.Where("email = ?", email).First(&u).RecordNotFound() {
		return u, errors.New("Record not found")
	}

	// Return the user.
	return u, nil
}

//
// Login a user in by email and password. The userAgent is a way to marking what device this
// login request came from. Same with ipAddress.
//
func (db *DB) LoginUserByEmailPass(email string, password string, appId uint, userAgent string, ipAddress string) (User, error) {

	var user User

	// See if we already have this user.
	user, err := db.GetUserByEmail(email)

	if err != nil {
		return user, errors.New("Sorry, we could not find your account.")
	}

	// Validate password here by comparing hashes nil means success
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	// Create a session so we get an access_token
	session, err := db.CreateSession(user.Id, appId, userAgent, ipAddress)

	if err != nil {
		services.Error(err)
		return User{}, err
	}

	// Add the session to the user object.
	user.Session = session

	return user, nil
}

//
// Create a new user.
//
func (db *DB) CreateUser(first string, last string, email string, password string, appId uint, userAgent string, ipAddress string) (User, error) {

	// Lets do some validation
	if err := db.ValidateCreateUser(first, last, email, password); err != nil {
		return User{}, err
	}

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		services.Error(err)
		return User{}, err
	}

	// Install user into the database
	var _first = template.HTMLEscapeString(first)
	var _last = template.HTMLEscapeString(last)

	user := User{FirstName: _first, LastName: _last, Email: email, Password: string(hash)}
	db.Create(&user)

	// Log user creation.
	services.Info("CreateUser - Created a new user account - " + first + " " + last + " " + email)

	// Create a session so we get an access_token (if we passed in an appId)
	if appId > 0 {
		session, err := db.CreateSession(user.Id, appId, userAgent, ipAddress)

		if err != nil {
			services.Error(err)
			return User{}, err
		}

		// Add the session to the user object.
		user.Session = session
	}

	// Return the user.
	return user, nil
}

//
// Validate a login user action.
//
func (db *DB) ValidateUserLogin(email string, password string) error {

	// Make sure the password is at least 6 chars long
	if len(password) < 6 {
		return errors.New("The password filed must be at least 6 characters long.")
	}

	// Lets validate the email address
	if err := db.ValidateEmailAddress(email); err != nil {
		return err
	}

	// See if we already have this user.
	_, err := db.GetUserByEmail(email)

	if err != nil {
		return errors.New("Sorry, we were unable to find our account.")
	}

	// Return happy.
	return nil
}

//
// Validate a create user action.
//
func (db *DB) ValidateCreateUser(first string, last string, email string, password string) error {

	// Are first and last name fields empty
	if (len(first) == 0) && (len(last) == 0) {
		return errors.New("First name and last name fields are required.")
	}

	// Are first name empty
	if len(first) == 0 {
		return errors.New("First name field is required.")
	}

	// Are last name empty
	if len(last) == 0 {
		return errors.New("Last name field is required.")
	}

	// Make sure the password is at least 6 chars long
	err := db.ValidatePassword(password)

	if err != nil {
		return err
	}

	// Lets validate the email address
	if err := db.ValidateEmailAddress(email); err != nil {
		return err
	}

	// See if we already have this user.
	_, err = db.GetUserByEmail(email)

	if err == nil {
		return errors.New("Looks like you already have an account.")
	}

	// Return happy.
	return nil
}

//
// Validate password.
//
func (db *DB) ValidatePassword(password string) error {

	// Make sure the password is at least 6 chars long
	if len(password) < 6 {
		return errors.New("The password filed must be at least 6 characters long.")
	}

	// Return happy.
	return nil
}

//
// Validate an email address
//
func (db *DB) ValidateEmailAddress(email string) error {

	// Check length
	if len(email) == 0 {
		return errors.New("Email address field is required.")
	}

	// Check format
	if err := checkmail.ValidateFormat(email); err != nil {
		return errors.New("Email address is not a valid format.")
	}

	// Return happy.
	return nil
}

/* End File */
