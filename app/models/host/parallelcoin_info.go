package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"net/http"
)

func parallelcoinINFO(db *scribble.Driver) *Host {
	////////////////
	// parallelcoin.INFO
	////////////////
	host := &Host{
		Name: "ParallelCoin ",
		Slug: "parallelcoin_info",
		Host: "parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		s := r.Host(host.Host).Subrouter()
		s.StrictSlash(true)
		//s.Headers("X-Requested-With", "XMLHttpRequest")
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/public/parallelcoininfo")))
	}
	host.Routes = routes
	return host
}
