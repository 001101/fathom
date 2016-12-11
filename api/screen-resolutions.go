package api

import (
  "net/http"
  "github.com/dannyvankooten/ana/db"
  "github.com/dannyvankooten/ana/count"
  "encoding/json"
)

// URL: /api/screen-resolutions
var GetScreenResolutionsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  before, after := getRequestedPeriods(r)

  // get total
  total := count.TotalVisitors(before, after)

  // get rows
  stmt, err := db.Conn.Prepare(`
    SELECT
    v.screen_resolution,
    COUNT(DISTINCT(pv.visitor_id)) AS count
    FROM pageviews pv
    LEFT JOIN visitors v ON v.id = pv.visitor_id
    WHERE UNIX_TIMESTAMP(pv.timestamp) <= ? AND UNIX_TIMESTAMP(pv.timestamp) >= ?
    GROUP BY v.screen_resolution
    ORDER BY count DESC
    LIMIT ?`)
  checkError(err)
  defer stmt.Close()
  rows, err := stmt.Query(before, after, getRequestedLimit(r))
  checkError(err)
  defer rows.Close()

  results := make([]Datapoint, 0)
  for rows.Next() {
    var d Datapoint
    err = rows.Scan(&d.Label, &d.Count);
    checkError(err)

    d.Percentage = float64(d.Count) / total * 100
    results = append(results, d)
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(results)
})
