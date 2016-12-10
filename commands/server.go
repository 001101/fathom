package commands

import(
  "net/http"
  "os"
  "github.com/dannyvankooten/ana/api"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "log"
)

func startServer() {
  // register routes
  r := mux.NewRouter()
  r.HandleFunc("/collect", api.CollectHandler).Methods("GET")
  r.Handle("/api/session", api.LoginHandler).Methods("POST")
  r.Handle("/api/session", api.LogoutHandler).Methods("DELETE")
  r.Handle("/api/visitors/count", api.Authorize(api.GetVisitorsCountHandler)).Methods("GET")
  r.Handle("/api/visitors/count/group/{period}", api.Authorize(api.GetVisitorsPeriodCountHandler)).Methods("GET")
  r.Handle("/api/visitors/count/realtime", api.Authorize(api.GetVisitorsRealtimeCountHandler)).Methods("GET")
  r.Handle("/api/pageviews/count", api.Authorize(api.GetPageviewsCountHandler)).Methods("GET")
  r.Handle("/api/pageviews/count/group/{period}", api.Authorize(api.GetPageviewsPeriodCountHandler)).Methods("GET")
  r.Handle("/api/pageviews", api.Authorize(api.GetPageviewsHandler)).Methods("GET")
  r.Handle("/api/languages", api.Authorize(api.GetLanguagesHandler)).Methods("GET")
  r.Handle("/api/screen-resolutions", api.Authorize(api.GetScreenResolutionsHandler)).Methods("GET")
  r.Handle("/api/countries", api.Authorize(api.GetCountriesHandler)).Methods("GET")
  r.Handle("/api/browsers", api.Authorize(api.GetBrowsersHandler)).Methods("GET")

  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  r.Path("/tracker.js").Handler(http.FileServer(http.Dir("./static/js/")))
  r.Handle("/", http.FileServer(http.Dir("./views/")))

  log.Print("API server is now listening on :8080")
  http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
