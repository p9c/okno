package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func beliRS(db *scribble.Driver) *Host {
	////////////////
	// parallelcoin.IO
	////////////////
	h := &Host{
		Name: "Beli",
		Slug: "beli_rs",
		Host: "beli.rs",
	}
	routes := func(r *mux.Router) {
		h.testRoutes(r)
	}
	h.Routes = routes
	return h
}
