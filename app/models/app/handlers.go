package app

import (
	"github.com/gorilla/mux"
	"github.com/p9c/okno/app/models/host"
	"net/http"
)

func (o *OKNO) Handlers() *mux.Router {

	//r := mux.NewRouter().StrictSlash(true)
	//r.Use(CommonMiddleware)

	//r.HandleFunc("/", controllers.TestAPI).Methods("GET")

	// Auth route
	//s := r.PathPrefix("/auth").Subrouter()
	//s.Use(auth.JwtVerify)
	//s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")

	//r := mux.NewRouter()
	//r.Headers("X-Requested-With", "XMLHttpRequest")

	r := mux.NewRouter().StrictSlash(true)
	for _, host := range host.GetHosts(o.Database) {
		host.Routes(r)
	}
	//r.NotFoundHandler = r.NewRoute().HandlerFunc(http.FileServer(http.Dir("js/errors/404"))).GetHandler()

	//r.NotFoundHandler = r.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	//r.Headers("X-Requested-With", "XMLHttpRequest")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
