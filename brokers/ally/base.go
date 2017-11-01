//
// Date: 10/31/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package ally

import (
  "os"
  "errors"
  "io/ioutil"
  "github.com/dghubble/oauth1" 
)

const (
  apiBaseUrl = "https://api.tradeking.com/v1"
)
//
// Send a GET request to Tradier. Returns the JSON string or an error
//
func SendGetRequest(urlStr string) (string, error) {

  // Read credentials from environment variables
  consumerKey := os.Getenv("ALLY_CONSUMER_KEY")
  consumerSecret := os.Getenv("ALLY_CONSUMER_SECRET")
  accessToken := os.Getenv("ALLY_ACCESS_TOKEN")
  accessSecret := os.Getenv("ALLY_ACCESS_SECRET")
  
  // Make sure there are not errors
  if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
    return "", errors.New("Missing configs.")
  }

  config := oauth1.NewConfig(consumerKey, consumerSecret)
  token := oauth1.NewToken(accessToken, accessSecret)

  // httpClient will automatically authorize http.Request's
  httpClient := config.Client(oauth1.NoContext, token)

  // Make request
  resp, err := httpClient.Get(apiBaseUrl + urlStr + ".json")
  
  if err != nil {
    return "", err  
  }

  defer resp.Body.Close()
  
  // Decode the body
  body, _ := ioutil.ReadAll(resp.Body)

  if err != nil {
    return "", err  
  }     
 
  // Return happy.
  return string(body), nil
}

/* End File */