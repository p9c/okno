package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"net/http"
)

func sokno(db *scribble.Driver) *Host {
	////////////////
	// s.okno.RS
	////////////////
	h := &Host{
		Name: "sokno",
		Slug: "sokno",
		Host: "s.okno.rs",
	}
	//hs := handlers.Handlers{jdb.NewJDB(db, h.Slug)}
	routes := func(r *mux.Router) {
		s := h.sub(r)
		s.Headers("X-Requested-With", "XMLHttpRequest")
		//l.Host("s.okno.rs").Path("/").HandlerFunc(hs.Homepage()).Name("okno")
		//l.Host("s.okno.rs").Path("/").PathPrefix("/").Handler(http.FileServer(http.Dir("static"))).Name("okno")
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	}
	h.Routes = routes
	//seed.Seed(jdb.NewJDB(db, h.Slug))
	return h
}
