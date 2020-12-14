package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func punqRS(db *scribble.Driver) *Host {
	////////////////
	// punq.RS
	////////////////
	host := &Host{
		Name: "Pun Q",
		Slug: "punq_rs",
		Host: "punq.rs",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
