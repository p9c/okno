package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"net/http"
)

func admin(db *scribble.Driver) *Host {
	////////////////
	// admin.okno.rs
	////////////////
	h := &Host{
		Name: "Admin",
		Slug: "admin",
		Host: "admin.okno.rs",
	}
	//h := handlers.Handlers{jdb.NewJDB(db, host.Slug)}
	routes := func(r *mux.Router) {
		s := h.sub(r)
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/admin/public")))
	}
	h.Routes = routes
	return h
}
