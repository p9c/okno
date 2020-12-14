package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/auth"
	"github.com/p9c/okno/app/handlers"
	"github.com/p9c/okno/app/jdb"
	"net/http"
)

func api(db *scribble.Driver) *Host {
	////////////////
	// api.okno.RS
	////////////////
	host := &Host{
		Name: "Api",
		Slug: "api",
		Host: "api.okno.rs",
	}
	h := handlers.Handlers{jdb.NewJDB(db, host.Slug)}
	routes := func(r *mux.Router) {
		s := r.Host(host.Host).Subrouter()
		s.Use(CommonMiddleware)
		host.testRoutes(r)
		s.HandleFunc("/register", h.APICreateUser).Methods("POST")
		s.HandleFunc("/login", h.APILogin()).Methods("POST")
		//Auth route
		a := r.PathPrefix("/auth").Subrouter()
		a.Use(auth.JwtVerify)
		a.HandleFunc("/user", h.APIFetchUsers).Methods("GET")
		a.HandleFunc("/user/{id}", h.APIGetUser).Methods("GET")
		a.HandleFunc("/user/{id}", h.APIUpdateUser).Methods("PUT")
		a.HandleFunc("/user/{id}", h.APIDeleteUser).Methods("DELETE")

		a.HandleFunc("/post", h.APIFetchUsers).Methods("GET")
		a.HandleFunc("/post/{id}", h.APIGetUser).Methods("GET")
		a.HandleFunc("/post/{id}", h.APIUpdateUser).Methods("PUT")
		a.HandleFunc("/post/{id}", h.APIDeleteUser).Methods("DELETE")
	}
	host.Routes = routes

	return host
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
