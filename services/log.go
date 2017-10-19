//
// Date: 10/18/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package services

import (
  "log" 
)

//
// Normal Log.
//
func LogInfo(message string) {
  log.Println(message)
}

//
// Debug Log.
//
func LogDebug(message string) {
  log.Println(message)
}

//
// Fatal Log.
//
func LogFatal(err error, message string) {
  log.Fatal(message + " (" + err.Error()  + ")")
}

//
// Error Log.
//
func LogError(err error, message string) {
  log.Println(message + " (" + err.Error()  + ")")
}

/* End File */