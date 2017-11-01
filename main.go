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
  _ "github.com/jpfuentes2/go-env/autoload"
  "github.com/net-worth-server/cron"
  "github.com/net-worth-server/models"
  "github.com/net-worth-server/services"
  "github.com/net-worth-server/controllers"
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

  // Startup controller
  c := &controllers.Controller{ DB: db }

  // Set Router
  r := mux.NewRouter()

  // Mark routes
  r.HandleFunc("/api/v1/marks", c.GetMarks).Methods("GET")

  // Ledgers routes
  r.HandleFunc("/api/v1/ledgers", c.GetLedgers).Methods("GET")
  r.HandleFunc("/api/v1/ledgers", c.CreateLedger).Methods("POST")

  // Account routes
  r.HandleFunc("/api/v1/accounts", c.GetAccounts).Methods("GET")
  r.HandleFunc("/api/v1/accounts", c.CreateAccount).Methods("POST")
  r.HandleFunc("/api/v1/accounts/{id}", c.GetAccount).Methods("GET")
  r.HandleFunc("/api/v1/accounts/{id}/marks", c.GetAccountMarks).Methods("GET")
  r.HandleFunc("/api/v1/accounts/{id}/marks", c.CreateAccountMark).Methods("POST")

  // Setup handler
  var handler = c.AuthMiddleware(r)

  if os.Getenv("LOG_REQUESTS") == "true" {
    handler = handlers.CombinedLoggingHandler(os.Stdout, c.AuthMiddleware(r))
  }
  
  // Setup http server    
  srv := &http.Server{
    Handler: handler,
    Addr: ":" + os.Getenv("HTTP_PORT"),
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  // Log this start.
  services.LogInfo("Starting web server on port " + os.Getenv("HTTP_PORT"))

  // Start server and log if fails
  log.Fatal(srv.ListenAndServe())
}

/* End File */