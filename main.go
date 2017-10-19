//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package main

import (
  "os"
  "log"
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"  
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql" 
  _ "github.com/jpfuentes2/go-env/autoload"
  "github.com/net-worth-server/models"
  "github.com/net-worth-server/services"
  "github.com/net-worth-server/controllers"
)

//
// Main...
//
func main() {
    
  // Start the db connection.
  db, err := SetupDb()

  if err != nil {
    services.LogFatal(err, "Failed to connect database")
  }

  // Close db when this app dies. (This might be useless)
  defer db.Close()

  // Startup controller
  c := &controllers.Controller{ DB: db }

  // Set Router
  r := mux.NewRouter()

  // Define routes
  r.HandleFunc("/api/v1/accounts", c.GetAccountsHandler).Methods("GET")
  r.HandleFunc("/api/v1/accounts/{id}", c.GetAccountHandler).Methods("GET")
  
  // Setup http server    
  srv := &http.Server{
    Handler: handlers.CombinedLoggingHandler(os.Stdout, c.AuthMiddleware(r)),
    Addr: ":" + os.Getenv("HTTP_PORT"),
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  // Log this start.
  services.LogInfo("Starting web server on port " + os.Getenv("HTTP_PORT"))

  // Start server and log if fails
  log.Fatal(srv.ListenAndServe())
}

//
// Setup the db connection.
//
func SetupDb() (*gorm.DB, error) {

  // Connect to Mysql
  db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_DATABASE") + "?charset=utf8&parseTime=True&loc=Local")
  
  if err != nil {
    return nil, err
  }

  // Enable
  //t.Connection.LogMode(true)
  //t.Connection.SetLogger(log.New(os.Stdout, "\r\n", 0))  

  // Run migrations
  db.AutoMigrate(&models.Account{})

  // Return db connection.
  return db, nil
}

/* End File */