//
// Date: 12/2/2019
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2019 Cloudmanic Labs, LLC. All rights reserved.
//

package etrade

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParsedEmail struct
type ParsedEmail struct {
	Balance float64 `sql:"type:DECIMAL(12,2)" json:"balance"`
}

//
// ParseDigestEmail will parse the HTML of your daily etrade digest email.
// We do this because etrade's API sucks and makes users login every day.
//
func ParseDigestEmail(html string) (parsed ParsedEmail, err error) {
	// Convert html to reader
	reader := strings.NewReader(html)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return parsed, err
	}

	// Find the review items
	doc.Find(".previewHeader").Each(func(i int, s *goquery.Selection) {
		// Parse HTML
		parts1 := strings.Split(s.Text(), "Market Value:")
		parts2 := strings.Split(parts1[1], "(")
		parts3 := strings.Split(parts2[0], "$")
		cleanedNumber := strings.Replace(strings.Trim(parts3[1], " "), ",", "", -1)

		// Convert to real
		if s, err := strconv.ParseFloat(cleanedNumber, 64); err == nil {
			parsed.Balance = s
		}
	})

	// Return happy
	return parsed, err
}

/* End File */
