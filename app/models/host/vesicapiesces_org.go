package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func vesicaPiescesORG(db *scribble.Driver) *Host {
	////////////////
	// VesicaPiescesORG
	////////////////
	host := &Host{
		Name: "Vesica Piesces",
		Slug: "vesicapiesces_org",
		Host: "vesicapiesces.org",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
