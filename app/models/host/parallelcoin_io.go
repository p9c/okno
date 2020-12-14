package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func parallelcoinIO(db *scribble.Driver) *Host {
	////////////////
	// parallelcoin.IO
	////////////////
	host := &Host{
		Name: "ParallelCoin",
		Slug: "parallelcoin",
		Host: "parallelcoin.io",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
