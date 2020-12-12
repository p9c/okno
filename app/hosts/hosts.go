package hosts

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"net/http"
)

type Host struct {
	Name   string
	Slug   string
	Host   string
	Routes func(r *mux.Router)
	Db     *scribble.Driver
}

func GetHosts() []*Host {
	return []*Host{
		okno(),
	}
}

func okno() *Host {
	////////////////
	// OKNO.rS
	////////////////
	host := &Host{
		Name: "OKNO",
		Slug: "okno",
		Host: "okno.rs",
	}
	routes := func(r *mux.Router) {
		r.Headers("X-Requested-With", "XMLHttpRequest")
		r.Host(host.Host).Path("/").HandlerFunc(host.ApiData()).Name("comhttpus")
		r.Host(host.Host).Path("/{coin}").HandlerFunc(host.ApiData()).Name("comhttpus")
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
	return host
}

func (h *Host) ApiData() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Asas")
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Println("asasasassa")
		//fmt.Fprintf(w, "Category: %v\n", vars["coin"])

		//profile := Profile{"Alex", []string{"snowboarding", "programming"}}

		js, err := json.Marshal(vars)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
