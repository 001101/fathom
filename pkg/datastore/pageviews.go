package datastore

import (
	"github.com/usefathom/fathom/pkg/models"
)

//var pv models.Pageview

// SavePageview ...
func SavePageview(pv *models.Pageview) error {
	// prepare statement for inserting data
	stmt, err := DB.Prepare(`INSERT INTO pageviews (
     page_id,
     visitor_id,
     referrer_url,
     referrer_keyword,
     timestamp
   ) VALUES( ?, ?, ?, ?, ? )`)
	defer stmt.Close()
	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		pv.PageID,
		pv.VisitorID,
		pv.ReferrerUrl,
		pv.ReferrerKeyword,
		pv.Timestamp,
	)

	if err != nil {
		return err
	}

	pv.ID, err = result.LastInsertId()
	return err
}

// SavePageviews ...
func SavePageviews(pvs []*models.Pageview) error {
	tx, err := DB.Begin()
	stmt, err := tx.Prepare(`INSERT INTO pageviews(
     page_id,
     visitor_id,
     referrer_url,
     referrer_keyword,
     timestamp
   ) VALUES( ?, ?, ?, ?, ? )`)
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, pv := range pvs {
		result, err := stmt.Exec(
			pv.PageID,
			pv.VisitorID,
			pv.ReferrerUrl,
			pv.ReferrerKeyword,
			pv.Timestamp,
		)

		if err != nil {
			return err
		}

		pv.ID, err = result.LastInsertId()
	}

	err = tx.Commit()
	return err
}
