package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func docParallelcoinINFO(db *scribble.Driver) *Host {
	////////////////
	// parallelcoin.IO
	////////////////
	host := &Host{
		Name: "ParallelCoin Documentation",
		Slug: "doc_parallelcoin_info",
		Host: "doc.parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
