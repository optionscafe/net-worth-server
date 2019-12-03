//
// Date: 10/27/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/optionscafe/net-worth-server/brokers/etrade"
	"github.com/optionscafe/net-worth-server/library/services"
)

// Signature struct represents the signature portion of the webhook POST body
type Signature struct {
	TimeStamp string `json:"timestamp"`
	Token     string `json:"token"`
	Signature string `json:"signature"`
}

//
// DoWebhookEmails will receive emails from mailgun and parse them.
//
func (t *Controller) DoWebhookEmails(c *gin.Context) {
	// Get POST vars
	html := c.PostForm("body-html")
	token := c.PostForm("token")
	timeStamp := c.PostForm("timestamp")
	signature := c.PostForm("signature")

	// Setup the signature
	sig := Signature{
		TimeStamp: timeStamp,
		Token:     token,
		Signature: signature,
	}

	// Validate email if we have a MG key on file. Mostly ignored for testing, I know lame testing strat.
	if len(os.Getenv("MG_MAILGUN_KEY")) > 0 {
		ok, err := VerifyWebhookSignature(sig)

		if err != nil {
			services.Error(err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			services.Error(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Email is not signed."})
			return
		}
	}

	// Since we only support Etrade at the moment we don't need to detect which type of email this is.
	// We just parse the email and get the balance of the account.
	parsed, err := etrade.ParseDigestEmail(html)

	if err != nil {
		services.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Make sure he balance is greater than zero
	if parsed.Balance <= 0 {
		services.Error(errors.New("Etrade balance is not greater than zero."))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Get the account id
	acctID, err := strconv.ParseFloat(os.Getenv("ETRADE_ACCOUNT_ID"), 64)

	if err != nil {
		services.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Look up the E*Trade account.
	acct, err := t.DB.GetAccountById(uint(acctID))

	if err != nil {
		services.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Format the date for today.
	ts := time.Now()
	date := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)

	// Mark account.
	t.DB.MarkAccountByDate(acct.Id, date, parsed.Balance)

	// Send health check notice.
	if len(os.Getenv("HEALTH_CHECK_ETRADE_URL")) > 0 {
		resp, err := http.Get(os.Getenv("HEALTH_CHECK_ETRADE_URL"))

		if err != nil {

			services.Error(errors.New(err.Error() + " Could send health check - " + os.Getenv("HEALTH_CHECK_ETRADE_URL")))
		}

		defer resp.Body.Close()
	}

	// Log success
	services.Info(`E*Trade account marked from email webhook. Balance: $` + fmt.Sprintf("%.2f", parsed.Balance))

	// Return Happy
	c.JSON(http.StatusOK, gin.H{})
}

//
// VerifyWebhookSignature to make sure the request is valid.
//
func VerifyWebhookSignature(sig Signature) (verified bool, err error) {
	apiKey := os.Getenv("MG_MAILGUN_KEY")
	h := hmac.New(sha256.New, []byte(apiKey))
	io.WriteString(h, sig.TimeStamp)
	io.WriteString(h, sig.Token)

	calculatedSignature := h.Sum(nil)
	signature, err := hex.DecodeString(sig.Signature)
	if err != nil {
		return false, err
	}
	if len(calculatedSignature) != len(signature) {
		return false, nil
	}

	return subtle.ConstantTimeCompare(signature, calculatedSignature) == 1, nil
}

/* End File */
