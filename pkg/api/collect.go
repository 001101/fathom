package api

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mssola/user_agent"
	"github.com/usefathom/fathom/pkg/aggregator"
	"github.com/usefathom/fathom/pkg/datastore"
	"github.com/usefathom/fathom/pkg/models"
)

func ShouldCollect(r *http.Request) bool {
	// abort if this is a bot.
	ua := user_agent.New(r.UserAgent())
	if ua.Bot() {
		return false
	}

	if r.Referer() == "" {
		return false
	}

	return true
}

func (api *API) NewCollectHandler() http.Handler {
	go aggregate(api.database)

	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		if !ShouldCollect(r) {
			return nil
		}

		u, err := url.Parse(r.Referer())
		if err != nil {
			return err
		}

		q := r.URL.Query()
		now := time.Now()

		// get pageview details
		pageview := &models.Pageview{
			SessionID:    q.Get("sid"),
			Hostname:     u.Scheme + "://" + u.Host,
			Pathname:     "/" + strings.TrimLeft(q.Get("p"), "/"),
			IsNewVisitor: q.Get("nv") == "1",
			IsNewSession: q.Get("ns") == "1",
			IsUnique:     q.Get("u") == "1",
			IsBounce:     q.Get("b") != "0",
			Referrer:     q.Get("r"),
			Duration:     0,
			Timestamp:    now,
		}

		// find previous pageview by same visitor
		if !pageview.IsNewSession {
			previousPageview, err := api.database.GetMostRecentPageviewBySessionID(pageview.SessionID)
			if err != nil && err != datastore.ErrNoResults {
				return err
			}

			// if we have a recent pageview that is less than 30 minutes old
			if previousPageview != nil && previousPageview.Timestamp.After(now.Add(-30*time.Minute)) {
				previousPageview.Duration = (now.Unix() - previousPageview.Timestamp.Unix())
				previousPageview.IsBounce = false
				err := api.database.UpdatePageview(previousPageview)
				if err != nil {
					return err
				}
			}
		}

		// save new pageview
		err = api.database.SavePageview(pageview)
		if err != nil {
			return err
		}

		// don't you cache this
		w.Header().Set("Content-Type", "image/gif")
		w.Header().Set("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.WriteHeader(http.StatusOK)

		// 1x1 px transparent GIF
		b, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
		w.Write(b)
		return nil
	})
}

// runs the aggregate func every minute
func aggregate(db datastore.Datastore) {
	agg := aggregator.New(db)
	agg.Run()

	timeout := 1 * time.Minute

	for {
		select {
		case <-time.After(timeout):
			agg.Run()
		}
	}
}
