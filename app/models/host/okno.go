package host

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/handlers"
	"github.com/p9c/okno/app/jdb"
	"github.com/p9c/okno/seed"
)

func okno(db *scribble.Driver) *Host {
	////////////////
	// okno.RS
	////////////////
	host := &Host{
		Name: "okno",
		Slug: "okno",
		Host: "okno.rs",
	}

	h := handlers.Handlers{jdb.NewJDB(db, host.Slug)}

	routes := func(r *mux.Router) {
		r.Headers("X-Requested-With", "XMLHttpRequest")
		r.Host(host.Host).Path("/").HandlerFunc(h.Homepage()).Name("comhttpus")
		r.Host(host.Host).Path("/{coin}").HandlerFunc(h.Homepage()).Name("comhttpus")
		//r.Host(okno).Path("/coins/").HandlerFunc(rts.Coins).Methods("GET")
		//r.Host(okno).Path("/bitnodes/").HandlerFunc(rts.BitNodes).Methods("GET")
		//r.Host(okno).Path("/bitnodes/{coin}/{nodeid}").HandlerFunc(rts.BitNode).Methods("GET")
		//
		//r.Host("{coin}." + okno).Path("/").HandlerFunc(rts.Coin).Methods("GET")
		//r.Host("{coin}." + okno).Path("/favicon.ico").HandlerFunc(rts.Ico).Name("ico")
		//
		//r.Host("api." + okno).Path("/{coin}/i").HandlerFunc(rts.ApiInfo).Name("info")
	}
	host.Routes = routes

	seed.Seed(jdb.NewJDB(db, host.Slug))
	return host
}
