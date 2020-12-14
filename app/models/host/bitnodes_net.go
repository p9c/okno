package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func bitNodesNET(db *scribble.Driver) *Host {
	////////////////
	// bitnodes.NET
	////////////////
	host := &Host{
		Name: "BitNodes",
		Slug: "bitnodes_net",
		Host: "bitnodes.net",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
