package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func explorerParallelcoinINFO(db *scribble.Driver) *Host {
	////////////////
	// explorer.parallelcoin.INFO
	////////////////
	host := &Host{
		Name: "ParallelCoin explorer",
		Slug: "explorer_parallelcoin_info",
		Host: "explorer.parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
