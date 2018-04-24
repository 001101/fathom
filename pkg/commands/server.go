package commands

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/usefathom/fathom/pkg/api"
	"log"
	"net/http"
	"os"
)

// Server starts the HTTP server, listening on the given port
func Server(port int) {

	// register routes
	r := mux.NewRouter()
	r.Handle("/collect", api.NewCollectHandler()).Methods("GET")
	r.Handle("/api/session", api.LoginHandler).Methods("POST")
	r.Handle("/api/session", api.LogoutHandler).Methods("DELETE")
	r.Handle("/api/visitors/count", api.Authorize(api.GetVisitorsCountHandler)).Methods("GET")
	r.Handle("/api/visitors/count/group/{period}", api.Authorize(api.GetVisitorsPeriodCountHandler)).Methods("GET")
	r.Handle("/api/visitors/count/realtime", api.Authorize(api.GetVisitorsRealtimeCountHandler)).Methods("GET")
	r.Handle("/api/pageviews/count", api.Authorize(api.GetPageviewsCountHandler)).Methods("GET")
	r.Handle("/api/pageviews/count/group/{period}", api.Authorize(api.GetPageviewsPeriodCountHandler)).Methods("GET")
	r.Handle("/api/pageviews", api.Authorize(api.GetPageviewsHandler)).Methods("GET")
	r.Handle("/api/languages", api.Authorize(api.GetLanguagesHandler)).Methods("GET")
	r.Handle("/api/referrers", api.Authorize(api.GetReferrersHandler)).Methods("GET")
	r.Handle("/api/screen-resolutions", api.Authorize(api.GetScreenResolutionsHandler)).Methods("GET")
	//r.Handle("/api/countries", api.Authorize(api.GetCountriesHandler)).Methods("GET")
	r.Handle("/api/browsers", api.Authorize(api.GetBrowsersHandler)).Methods("GET")

	r.Path("/tracker.js").Handler(http.FileServer(http.Dir("./assets/dist/js/")))

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./assets/dist/"))))

	log.Printf("HTTP server will now start listening on http://127.0.01:%d/\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.LoggingHandler(os.Stdout, r))
	log.Println(err)
}
