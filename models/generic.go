//
// Date: 2/10/2018
// Author(s): Spicer Matthews (spicer@cloudmanic.com)
// Copyright: 2018 Cloudmanic Labs, LLC. All rights reserved.
//

package models

import (
	"errors"
	"math"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type InsertParam struct {
	Debug bool
}

type QueryParam struct {
	UserId           uint
	AccountId        uint
	Limit            int
	Offset           int
	Page             int
	Order            string
	Sort             string
	SearchCols       []string
	SearchTerm       string
	Debug            bool
	Wheres           []KeyValue
	PreLoads         []string
	AllowedOrderCols []string
}

type KeyValue struct {
	Key      string
	Value    string
	ValueInt int
}

type QueryMetaData struct {
	Page         int
	Limit        int
	Offset       int
	PageCount    int
	LastPage     bool
	LimitCount   int
	NoLimitCount int
}

//
// Create a new entry.
//
func (t *DB) CreateNewRecord(model interface{}, params InsertParam) error {

	if params.Debug {

		if err := t.Debug().Create(model).Error; err != nil {
			return err
		}

	} else {

		if err := t.Create(model).Error; err != nil {
			return err
		}

	}

	return nil
}

//
// A generic way to query any model we want.
//
func (t *DB) Query(model interface{}, params QueryParam) error {

	// Build the query.
	query, err := t.buildGenericQuery(params)

	if err != nil {
		return err
	}

	// Run the query.
	if err := query.Find(model).Error; err != nil {
		return err
	}

	// If we made it this far no errors.
	return nil
}

//
// A generic way to query any model we want with meta data.
//
func (t *DB) QueryWithNoFilterCount(model interface{}, params QueryParam) (int, error) {

	var noFilterCount int = 0

	// Build the query.
	query, err := t.buildGenericQuery(params)

	if err != nil {
		return noFilterCount, err
	}

	// Run the query.
	if err := query.Find(model).Offset(-1).Limit(-1).Count(&noFilterCount).Error; err != nil {
		return noFilterCount, err
	}

	// If we made it this far no errors.
	return noFilterCount, nil

}

//
// Return all the meta data we need about a query.
//
func (t *DB) GetQueryMetaData(limitCount int, noLimitCount int, params QueryParam) QueryMetaData {

	// Start meta data object
	meta := QueryMetaData{
		LastPage:     false,
		Limit:        int(params.Limit),
		LimitCount:   limitCount,
		NoLimitCount: noLimitCount,
	}

	// Need a limit value.
	if meta.Limit <= 0 {
		return meta
	}

	// Add the page.
	if params.Page > 0 {
		meta.Page = params.Page
	}

	// Set offset
	if meta.Page > 0 {
		meta.Offset = (meta.Page * meta.Limit) - meta.Limit
	}

	// Get page count
	meta.PageCount = int(math.Ceil(float64(noLimitCount) / float64(params.Limit)))

	// Figure out if we are on the last page.
	if meta.Page == meta.PageCount {
		meta.LastPage = true
	}

	// Return meta data object.
	return meta
}

//
// A generic way find out how many rows are in a table
//
func (t *DB) Count(model interface{}, params QueryParam) (uint, error) {

	var count uint = 0
	var query *gorm.DB

	// Useful just to kick this off
	query = t.Order("id ASC")

	// Are we debugging this?
	if params.Debug {
		query = query.Debug()
	}

	// Add in Where clauses
	for _, row := range params.Wheres {
		if len(row.Value) > 0 {
			query = query.Where(row.Key+" = ?", row.Value)
		}

		if row.ValueInt > 0 {
			query = query.Where(row.Key+" = ?", row.ValueInt)
		}
	}

	// Run the query
	query.Model(model).Count(&count)

	// If we made it this far no errors.
	return count, nil
}

//
// Build generic query.
//
func (t *DB) buildGenericQuery(params QueryParam) (*gorm.DB, error) {
	var query *gorm.DB

	// Validate order column
	if (len(params.Order) > 0) && (len(params.AllowedOrderCols) > 0) {
		var found = false

		for _, row := range params.AllowedOrderCols {
			if params.Order == row {
				found = true
			}
		}

		if !found {
			return query, errors.New("Invalid order parameter. - " + params.Order)
		}
	}

	// Do some quick filtering - Think injections
	var sortText = strings.ToUpper(params.Sort)
	if len(sortText) > 0 && ((sortText != "ASC") && (sortText != "DESC")) {
		return query, errors.New("Invalid sort parameter. - " + params.Sort)
	}

	// Set order and get query object
	if (len(params.Order) > 0) && (len(params.Sort) > 0) {
		query = t.Order(params.Order + " " + params.Sort)
	} else if len(params.Order) > 0 {
		query = t.Order(params.Order + " ASC")
	} else {
		query = t.Order("id ASC")
	}

	// Are we debugging this?
	if params.Debug {
		query = query.Debug()
	}

	// If we passed in a page we figure out the offset from the page.
	if params.Page > 0 {
		if (params.Page > 0) && (params.Limit > 0) {
			params.Offset = (params.Page * params.Limit) - params.Limit
		}
	}

	// Offset
	if params.Offset > 0 {
		query = query.Offset(params.Offset)
	}

	// Limit
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	// Add in user id (almost every table has this column)
	if params.UserId > 0 {
		query = query.Where("user_id = ?", params.UserId)
	}

	// Add preloads
	for _, row := range params.PreLoads {
		query = query.Preload(row)
	}

	// Add in Where clauses
	for _, row := range params.Wheres {
		if len(row.Value) > 0 {
			query = query.Where(row.Key+" = ?", row.Value)
		}

		if row.ValueInt > 0 {
			query = query.Where(row.Key+" = ?", row.ValueInt)
		}
	}

	// Search a particular column
	if (len(params.SearchTerm) > 0) && (len(params.SearchCols) > 0) {
		var likes []string
		var terms []interface{}

		for _, row := range params.SearchCols {
			str := row + " LIKE ?"
			likes = append(likes, str)
			terms = append(terms, "%"+params.SearchTerm+"%")
		}

		// Built where query.
		query = query.Where(strings.Join(likes, " OR "), terms...)
	}

	// Return query.
	return query, nil
}

/* End File */
