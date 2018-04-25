package datastore

import (
	"github.com/usefathom/fathom/pkg/models"
)

// SavePageview inserts a single pageview model into the connected database
func SavePageview(pv *models.Pageview) error {
	query := dbx.Rebind(`INSERT INTO pageviews(page_id, visitor_id, referrer_url, referrer_keyword, timestamp) VALUES( ?, ?, ?, ?, ?)`)
	result, err := dbx.Exec(query, pv.PageID, pv.VisitorID, pv.ReferrerUrl, pv.ReferrerKeyword, pv.Timestamp)
	if err != nil {
		return err
	}

	pv.ID, _ = result.LastInsertId()
	return nil
}

// SavePageviews inserts multiple pageview models into the connected database using a transaction
func SavePageviews(pvs []*models.Pageview) error {
	query := dbx.Rebind(`INSERT INTO pageviews(page_id, visitor_id, referrer_url, referrer_keyword, timestamp ) VALUES( ?, ?, ?, ?, ? )`)
	tx, err := dbx.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, pv := range pvs {
		result, err := stmt.Exec(pv.PageID, pv.VisitorID, pv.ReferrerUrl, pv.ReferrerKeyword, pv.Timestamp)
		if err != nil {
			return err
		}

		pv.ID, err = result.LastInsertId()
	}

	err = tx.Commit()
	return err
}
