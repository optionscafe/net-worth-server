//
// Date: 10/29/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/net-worth-server/models_mock"
)

//
// Test GetMarks
//
func TestGetMarks(t *testing.T) {
    
  // Create controller
  c := &Controller{ DB: &models_mock.DB{} }

  // Make a mock request.
  rec := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/api/v1/marks", nil)
  http.HandlerFunc(c.GetMarks).ServeHTTP(rec, req)

  // Expected JSON response.
  expected := `[{"id":1,"created_at":"2017-10-29T17:20:01.000507451Z","updated_at":"2017-10-29T17:20:01.000507451Z","date":"2017-10-29T17:20:01.000507451Z","units":45.23,"price_per":16.78,"balance":128.33},{"id":2,"created_at":"2017-10-29T17:20:01.000507451Z","updated_at":"2017-10-29T17:20:01.000507451Z","date":"2017-10-29T17:20:01.000507451Z","units":45.23,"price_per":16.78,"balance":128.33}]`

  // Test that the json response is correct.
  if expected != rec.Body.String() {
    t.Errorf("\n\n...expected = %v\n\n...obtained = %v\n\n", expected, rec.Body.String())
  }
}

/* End File */