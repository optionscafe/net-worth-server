//
// Date: 10/29/2017
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
//

package models_mock

import (
  "time"
  "github.com/net-worth-server/models"
)

//
// Get all marks.
//
func (db *DB) GetAllMarks() []models.Mark {
  ts := time.Date(2017, 10, 29, 17, 20, 01, 507451, time.UTC)
  objs := make([]models.Mark, 0)
  objs = append(objs, models.Mark{ 1, ts, ts, ts, 45.23, 16.78, 128.33 })
  objs = append(objs, models.Mark{ 2, ts, ts, ts, 45.23, 16.78, 128.33 })
  return objs
}

//
// Mark account by date.
//
func (db *DB) MarkByDate(date time.Time) error {
  return nil
}

//
// Get mark by date.
//
func (db *DB) GetMarkByDate(date time.Time) (*models.Mark, error) {
  return &models.Mark{}, nil
}

/* End File */