package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func solutionsRS(db *scribble.Driver) *Host {
	////////////////
	// solutions.RS
	////////////////
	host := &Host{
		Name: "W-ing Solutions",
		Slug: "solutions_rs",
		Host: "solutions.rs",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
