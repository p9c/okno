package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func beliRS(db *scribble.Driver) *Host {
	////////////////
	// parallelcoin.IO
	////////////////
	host := &Host{
		Name: "Beli",
		Slug: "beli_rs",
		Host: "beli.rs",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
