package models

import (
	"time"
)

type ReferrerStats struct {
	URL         string    `db:"url"`
	Visitors    int64     `db:"visitors"`
	Pageviews   int64     `db:"pageviews"`
	BounceRate  float64   `db:"bounce_rate"`
	AvgDuration int64     `db:"avg_duration"`
	Date        time.Time `db:"date" json:"omitempty"`
}
