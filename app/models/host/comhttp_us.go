package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
)

func comhttpUS(db *scribble.Driver) *Host {
	////////////////
	// com-httpUS
	////////////////
	host := &Host{
		Name: "COM-HTTP.US",
		Slug: "comhttp_us",
		Host: "com-http.us",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
