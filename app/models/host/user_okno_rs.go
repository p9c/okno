package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/auth"
	"github.com/p9c/okno/app/handlers"
	"github.com/p9c/okno/app/jdb"
)

func users(db *scribble.Driver) *Host {
	////////////////
	// user.okno.RS
	////////////////
	host := &Host{
		Name: "Users",
		Slug: "users",
		Host: "user.okno.rs",
	}
	h := handlers.Handlers{jdb.NewJDB(db, host.Slug)}

	routes := func(r *mux.Router) {
		r.Use(CommonMiddleware)
		s := r.PathPrefix("/auth").Subrouter()
		s.Use(auth.JwtVerify)
		s.HandleFunc("/user", h.APIFetchUsers).Methods("GET")
		s.HandleFunc("/user/{id}", h.APIGetUser).Methods("GET")
		s.HandleFunc("/user/{id}", h.APIUpdateUser).Methods("PUT")
		s.HandleFunc("/user/{id}", h.APIDeleteUser).Methods("DELETE")
	}
	host.Routes = routes
	return host
}
