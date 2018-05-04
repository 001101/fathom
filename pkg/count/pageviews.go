package count

import (
	"log"
	"time"

	"github.com/usefathom/fathom/pkg/datastore"
	"github.com/usefathom/fathom/pkg/models"
)

// Pageviews returns a point slice containing language data per language
func Pageviews(before int64, after int64, limit int64) ([]*models.Total, error) {
	points, err := datastore.TotalPageviewsPerPage(before, after, limit)
	if err != nil {
		return nil, err
	}

	total, err := datastore.TotalPageviews(before, after)
	if err != nil {
		return nil, err
	}

	points = calculatePercentagesOfTotal(points, total)
	return points, nil
}

// CreatePageviewTotals aggregates pageview data for each page into daily totals
func CreatePageviewTotals(since string) {
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	totals, err := datastore.PageviewCountPerPageAndDay(tomorrow, since)
	if err != nil {
		log.Fatal(err)
	}

	err = datastore.SavePageTotals("pageviews", totals)
	if err != nil {
		log.Fatal(err)
	}
}
