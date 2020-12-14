package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func whitepaperParallelcoinINFO(db *scribble.Driver) *Host {
	////////////////
	// whitepaper.parallelcoin.INFO
	////////////////
	host := &Host{
		Name: "ParallelCoin Whitepaper",
		Slug: "whitepaper_parallelcoin_info",
		Host: "whitepaper.parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
